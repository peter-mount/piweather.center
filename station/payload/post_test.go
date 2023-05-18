package payload

import "testing"

func TestUnmarshalPost(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		wantErr bool
		want    map[string]interface{}
	}{
		{
			name:    "ecowitt",
			content: []byte("PASSKEY=A50960B3AC048F573FCA89960A56F661&stationtype=GW2000A_V2.2.3&runtime=519509&dateutc=2023-05-18+17:47:20&tempinf=76.28&humidityin=38&baromrelin=29.976&baromabsin=29.976&tempf=73.94&humidity=36&winddir=71&windspeedmph=1.12&windgustmph=2.24&maxdailygust=5.82&solarradiation=85.71&uv=0&rrain_piezo=0.000&erain_piezo=0.000&hrain_piezo=0.000&drain_piezo=0.000&wrain_piezo=0.000&mrain_piezo=0.000&yrain_piezo=0.000&ws90cap_volt=4.2&ws90_ver=126&lightning_num=0&lightning=&lightning_time=&wh57batt=5&wh90batt=3.18&freq=868M&model=GW2000A&interval=60"),
			wantErr: false,
			want: map[string]interface{}{
				"PASSKEY":        "A50960B3AC048F573FCA89960A56F661",
				"stationtype":    "GW2000A_V2.2.3",
				"runtime":        "519509",
				"dateutc":        "2023-05-18 17:47:20",
				"tempinf":        "76.28",
				"humidityin":     "38",
				"baromrelin":     "29.976",
				"baromabsin":     "29.976",
				"tempf":          "73.94",
				"humidity":       "36",
				"winddir":        "71",
				"windspeedmph":   "1.12",
				"windgustmph":    "2.24",
				"maxdailygust":   "5.82",
				"solarradiation": "85.71",
				"uv":             "0",
				"rrain_piezo":    "0.000",
				"erain_piezo":    "0.000",
				"hrain_piezo":    "0.000",
				"drain_piezo":    "0.000",
				"wrain_piezo":    "0.000",
				"mrain_piezo":    "0.000",
				"yrain_piezo":    "0.000",
				"ws90cap_volt":   "4.2",
				"ws90_ver":       "126",
				"lightning_num":  "0",
				"lightning":      "",
				"lightning_time": "",
				"wh57batt":       "5",
				"wh90batt":       "3.18",
				"freq":           "868M",
				"model":          "GW2000A",
				"interval":       "60",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := make(map[string]interface{})
			err := UnmarshalPost(tt.content, &m)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("UnmarshalPost() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Check we got what we expected
			for k, v := range tt.want {
				e, exists := m[k]
				switch {
				case !exists:
					t.Errorf("Expected %q but missing from response", k)
				case v != e:
					t.Errorf("Expected %q for %q got %q", v, k, e)
				}
			}

			// Report if we get something we shouldnt have
			for k, v := range m {
				if _, exists := tt.want[k]; !exists {
					t.Errorf("Unexpected result %q got %q", k, v)
				}
			}
		})
	}
}
