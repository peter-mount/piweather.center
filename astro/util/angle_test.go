package util

import (
	"github.com/soniakeys/unit"
	"testing"
)

func TestParseAngle(t *testing.T) {
	tests := []struct {
		name    string
		want    unit.Angle
		wantErr bool
	}{
		// int
		{name: "12", want: unit.AngleFromDeg(12)},
		{name: "+12", want: unit.AngleFromDeg(12)},
		{name: "N12", want: unit.AngleFromDeg(12)},
		{name: "n12", want: unit.AngleFromDeg(12)},
		{name: "-12", want: unit.AngleFromDeg(-12)},
		{name: "S12", want: unit.AngleFromDeg(-12)},
		{name: "s12", want: unit.AngleFromDeg(-12)},
		// float64
		{name: "12.3456", want: unit.AngleFromDeg(12.3456)},
		{name: "+12.3456", want: unit.AngleFromDeg(12.3456)},
		{name: "N12.3456", want: unit.AngleFromDeg(12.3456)},
		{name: "n12.3456", want: unit.AngleFromDeg(12.3456)},
		{name: "-12.3456", want: unit.AngleFromDeg(-12.3456)},
		{name: "S12.3456", want: unit.AngleFromDeg(-12.3456)},
		{name: "s12.3456", want: unit.AngleFromDeg(-12.3456)},
		// DD:MM:SS
		{name: "12:34:56", want: unit.NewAngle('+', 12, 34, 56)},
		{name: "12:34:56.78", want: unit.NewAngle('+', 12, 34, 56.78)},
		{name: "+12:34:56", want: unit.NewAngle('+', 12, 34, 56)},
		{name: "+12:34:56.78", want: unit.NewAngle('+', 12, 34, 56.78)},
		{name: "N12:34:56", want: unit.NewAngle('+', 12, 34, 56)},
		{name: "N12:34:56.78", want: unit.NewAngle('+', 12, 34, 56.78)},
		{name: "n12:34:56", want: unit.NewAngle('+', 12, 34, 56)},
		{name: "n12:34:56.78", want: unit.NewAngle('+', 12, 34, 56.78)},
		{name: "-12:34:56", want: unit.NewAngle('-', 12, 34, 56)},
		{name: "-12:34:56.78", want: unit.NewAngle('-', 12, 34, 56.78)},
		{name: "S12:34:56", want: unit.NewAngle('-', 12, 34, 56)},
		{name: "S12:34:56.78", want: unit.NewAngle('-', 12, 34, 56.78)},
		{name: "s12:34:56", want: unit.NewAngle('-', 12, 34, 56)},
		{name: "s12:34:56.78", want: unit.NewAngle('-', 12, 34, 56.78)},
		// DD:MM
		{name: "12:34", want: unit.NewAngle('+', 12, 34, 0)},
		{name: "12:34.8", want: unit.NewAngle('+', 12, 34, 48)},
		{name: "+12:34", want: unit.NewAngle('+', 12, 34, 0)},
		{name: "+12:34.8", want: unit.NewAngle('+', 12, 34, 48)},
		{name: "N12:34", want: unit.NewAngle('+', 12, 34, 0)},
		{name: "N12:34.8", want: unit.NewAngle('+', 12, 34, 48)},
		{name: "n12:34", want: unit.NewAngle('+', 12, 34, 0)},
		{name: "n12:34.8", want: unit.NewAngle('+', 12, 34, 48)},
		{name: "-12:34", want: unit.NewAngle('-', 12, 34, 0)},
		{name: "-12:34.8", want: unit.NewAngle('-', 12, 34, 48)},
		{name: "S12:34", want: unit.NewAngle('-', 12, 34, 0)},
		{name: "S12:34.8", want: unit.NewAngle('-', 12, 34, 48)},
		{name: "s12:34", want: unit.NewAngle('-', 12, 34, 0)},
		{name: "s12:34.8", want: unit.NewAngle('-', 12, 34, 48)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAngle(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAngle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseAngle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
