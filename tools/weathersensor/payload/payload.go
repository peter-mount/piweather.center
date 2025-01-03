package payload

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/peter-mount/piweather.center/config/station"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"gopkg.in/yaml.v3"
	"strconv"
	"strings"
	"time"
)

type Payload struct {
	id        string
	format    station.HttpFormatType
	timeField *station.SourcePath
	time      time.Time
	msg       []byte
	data      map[string]interface{}
}

func GetPayload(ctx context.Context) *Payload {
	return ctx.Value("payload.key").(*Payload)
}

func (p *Payload) AddContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "payload.key", p)
}

func (p *Payload) Id() string {
	return p.id
}

func (p *Payload) Format() station.HttpFormatType {
	return p.format
}

func (p *Payload) TimeField() *station.SourcePath {
	return p.timeField
}

func (p *Payload) Time() time.Time {
	return p.time
}

func (p *Payload) Msg() []byte {
	return p.msg
}

func (p *Payload) Data() map[string]interface{} {
	return p.data
}

func (p *Payload) Get(path *station.SourcePath) (interface{}, bool) {
	m := p.data

	l := len(path.Path) - 1
	for i, k := range path.Path {
		v, ok := m[k]
		if !ok {
			return nil, false
		}
		if i == l {
			return v, true
		}

		if nm, ok := v.(map[string]interface{}); ok {
			m = nm
		} else {
			return nil, false
		}
	}
	return nil, false
}

func FromBytes(id string, format station.HttpFormatType, timestamp *station.SourcePath, msg []byte) (*Payload, error) {
	// Payload needs a copy of msg in case the provider reuses that slice.
	// Time defaults to now in UTC before being overridden by the payload
	// as that allows for those messages without a time
	p := &Payload{
		id:        id,
		format:    format,
		timeField: timestamp,
		time:      time.Now().UTC(),
		msg:       make([]byte, len(msg)),
		data:      make(map[string]interface{}),
	}
	copy(p.msg, msg)

	var err error
	switch format {
	case station.HttpFormatJson:
		err = json.Unmarshal(p.msg, &p.data)

	case station.HttpFormatXml:
		err = xml.Unmarshal(p.msg, &p.data)

	case station.HttpFormatYaml:
		err = yaml.Unmarshal(p.msg, &p.data)

	case station.HttpFormatPost:
		err = UnmarshalPost(p.msg, &p.data)

	// carbon is formatted string: "key value timestamp" where timestamp is in unix seconds
	case station.HttpFormatCarbon:
		s := strings.SplitN(string(p.msg), " ", 3)
		if len(s) != 3 {
			err = fmt.Errorf("invalid carbon record %q", string(p.msg))
		} else if us, err1 := strconv.ParseInt(s[2], 10, 64); err1 != nil {
			err = err1
		} else {
			p.time = time.Unix(us, 0)
			p.data = map[string]interface{}{
				"timestamp": p.time,
				s[0]:        s[1],
			}
		}

	default:
		err = fmt.Errorf("unsupported format %q", format)
	}

	if err != nil {
		// uncomment and return nil for err if there's issues
		//log.Printf("Payload: %v", err)
		return nil, err
	}

	if timestamp != nil {
		if ts, ok := p.Get(timestamp); ok {
			if st, ok := ts.(string); ok {
				if t := time2.ParseTime(st); !t.IsZero() {
					p.time = t
				}
			}
		}
	}

	return p, nil
}

func FromLog(s string) (*Payload, error) {
	// Logs are timestamp,id,timestampField,format,{payload}
	a := strings.SplitN(s, ",", 5)
	if len(a) < 5 {
		return nil, nil
	}
	return FromBytes(
		a[1],
		station.ParseHttpFormatType(a[3]),
		&station.SourcePath{Path: strings.Split(a[2], ".")},
		[]byte(a[4]))
}

func (p *Payload) ToLog() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s\n",
		p.time.UTC().Format(time.RFC3339),
		p.id,
		p.timeField,
		p.format.String(),
		string(p.Msg()),
	)
}
