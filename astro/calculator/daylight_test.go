package calculator

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/unit"
	"testing"
	"time"
)

func Test_calculator_SolarEphemeris(t *testing.T) {
	calc := &calculator{}
	err := calc.Start()
	if err != nil {
		t.Fatal(err)
	}
	// FIXME I need to find out a way to reference a local file/directory from a test
	calc.rootDir = "/home/peter/area51/piweather.center/builds/linux/amd64/lib/"

	londonLocation, err := time.LoadLocation("Europe/London")
	if err != nil {
		t.Fatal(err)
	}

	tm := value.BasicTime(
		time.Date(2023, 6, 10, 21, 0, 0, 0, londonLocation),
		// London
		&globe.Coord{
			Lat: unit.AngleFromDeg(51.5),
			// remember Meeus uses +ve west
			Lon: unit.AngleFromDeg(8.0 / 60.0),
		},
		0.0)

	tests := []struct {
		t       value.Time
		wantErr bool
	}{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got, err := calc.SolarEphemeris(tm.Clone())
			if (err != nil) != tt.wantErr {
				t.Errorf("SolarEphemeris() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println("Astronomical", got.AstronomicalDawn, got.AstronomicalDawn.IsValid())
			fmt.Println("    Nautical", got.NauticalDawn, got.NauticalDawn.IsValid())
			fmt.Println("       Civil", got.CivilDawn, got.CivilDawn.IsValid())
			fmt.Println("        Rise", got.SunRise, got.SunRise.IsValid())
			fmt.Println("         Set", got.SunSet, got.SunSet.IsValid())
			fmt.Println("       Civil", got.CivilDusk, got.CivilDusk.IsValid())
			fmt.Println("    Nautical", got.NauticalDusk, got.NauticalDusk.IsValid())
			fmt.Println("Astronomical", got.AstronomicalDusk, got.AstronomicalDusk.IsValid())
			fmt.Println("  Day Length", got.DayLength)
			fmt.Println("UpperTransit", got.UpperTransit, got.UpperTransit.IsValid())
			fmt.Println("LowerTransit", got.LowerTransit, got.LowerTransit.IsValid())
		})
	}
}
