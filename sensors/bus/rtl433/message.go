package rtl433

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/tools/weathersensor/payload"
)

type Message struct {
	Payload *payload.Payload
	Model   string
	ID      string
	SubType string
}

var (
	rtlTimestamp = &station.SourcePath{Path: []string{"time"}}
)

func NewMessage(msg []byte) (*Message, error) {
	p, err := payload.FromBytes("rtl433", station.HttpFormatJson, rtlTimestamp, msg)
	if err == nil {
		return &Message{
			Payload: p,
			Model:   p.GetString("model"),
			ID:      p.GetIntString("id"),
			SubType: p.GetIntString("subtype"),
		}, nil
	}
	return nil, err
}
