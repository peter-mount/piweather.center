package rainfall

import "time"

type RainFall struct {
	Time   string `json:"time"`
	Record Record `json:"record"`
	Device Device `json:"device"`
}

type Device struct {
	Device       string `json:"device"`
	Manufacturer string `json:"manufacturer"`
	Version      uint16 `json:"version"`
	Uptime       int    `json:"uptime"`
}

type Record struct {
	Hour        float64 `json:"hour"`
	Day         float64 `json:"day"`
	Total       float64 `json:"total"`
	BucketCount uint32  `json:"bucket_count"`
}

func newRainFall(version uint16) RainFall {
	return RainFall{
		Time: time.Now().Format(time.RFC3339),
		Device: Device{
			Device:       "SEN0575",
			Manufacturer: "DF Robot",
			Version:      version,
		},
	}
}
