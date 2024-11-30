package state

import (
	"fmt"
	"github.com/peter-mount/piweather.center/store/api"
	"strconv"
	"testing"
	"time"
)

func TestResponse_Add(t *testing.T) {
	type args struct {
		t  string
		id string
		i  int
		s  string
		m  api.Metric
	}
	tests := []struct {
		args args
	}{
		{
			args: args{
				t:  "rain",
				id: "1234",
				i:  0, s: "",
				m: api.Metric{
					Metric: "test.metric",
					Time:   time.Now(),
					Unit:   "celsius",
					Value:  10.5,
				},
			},
		},
	}
	for ti, tt := range tests {
		t.Run(strconv.Itoa(ti), func(t *testing.T) {
			r := &Response{}
			r.Set(tt.args.t, tt.args.id, tt.args.i, tt.args.s, tt.args.m)

			b, _ := r.Json()
			fmt.Println(string(b))
		})
	}
}
