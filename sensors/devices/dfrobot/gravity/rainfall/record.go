package rainfall

import "time"

type RainFall struct {
	Time   time.Time `json:"time"`
	Record Record    `json:"record"`
	Device Device    `json:"device"`
}

type Device struct {
	Device       string        `json:"device"`
	Manufacturer string        `json:"manufacturer"`
	Version      uint16        `json:"version"`
	Uptime       time.Duration `json:"uptime"`
}

type Record struct {
	Hour  float64 `json:"hour"`
	Day   float64 `json:"day"`
	Total float64 `json:"total"`
}

func newRainFall(version uint16) RainFall {
	return RainFall{
		Time: time.Now(),
		Device: Device{
			Device:       "SEN0575",
			Manufacturer: "DF Robot",
			Version:      version,
		},
	}
}
