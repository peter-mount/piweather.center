package strings

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

func TestDegDMSStringExt(t *testing.T) {
	type args struct {
		d         float64
		sign      bool
		degDigits int
		precision int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// This tests the output where the value is close to the upper limit.
		// e.g When d is very close to 360.0 (which is out of bounds as it's equivalent to 0)
		// the traditional code returns "359:59:60" which is invalid for seconds as the
		// usual FormatFloat (or %f in sprintf) rounds the value up which is incorrect
		//
		// Precision 0
		{name: "dd:mm:ss", args: args{d: 359.9, degDigits: 3, precision: 0}, want: "359:53:59"},
		{name: "dd:mm:ss", args: args{d: 359.99, degDigits: 3, precision: 0}, want: "359:59:24"},
		{name: "dd:mm:ss", args: args{d: 359.999, degDigits: 3, precision: 0}, want: "359:59:56"},
		{name: "dd:mm:ss", args: args{d: 359.9999, degDigits: 3, precision: 0}, want: "359:59:59"},
		{name: "dd:mm:ss", args: args{d: 359.99999, degDigits: 3, precision: 0}, want: "359:59:59"},
		{name: "dd:mm:ss", args: args{d: 359.999999, degDigits: 3, precision: 0}, want: "359:59:59"},
		{name: "dd:mm:ss", args: args{d: 359.9999999, degDigits: 3, precision: 0}, want: "359:59:59"},
		// Precision 1
		{name: "dd:mm:ss.s", args: args{d: 359.9, degDigits: 3, precision: 1}, want: "359:53:59.9"},
		{name: "dd:mm:ss.s", args: args{d: 359.99, degDigits: 3, precision: 1}, want: "359:59:24.0"},
		{name: "dd:mm:ss.s", args: args{d: 359.999, degDigits: 3, precision: 1}, want: "359:59:56.4"},
		{name: "dd:mm:ss.s", args: args{d: 359.9999, degDigits: 3, precision: 1}, want: "359:59:59.6"},
		{name: "dd:mm:ss.s", args: args{d: 359.99999, degDigits: 3, precision: 1}, want: "359:59:59.9"},
		{name: "dd:mm:ss.s", args: args{d: 359.999999, degDigits: 3, precision: 1}, want: "359:59:59.9"},
		{name: "dd:mm:ss.s", args: args{d: 359.9999999, degDigits: 3, precision: 1}, want: "359:59:59.9"},
		// Precision 2
		{name: "dd:mm:ss.ss", args: args{d: 359.9, degDigits: 3, precision: 2}, want: "359:53:59.99"},
		{name: "dd:mm:ss.ss", args: args{d: 359.99, degDigits: 3, precision: 2}, want: "359:59:24.00"},
		{name: "dd:mm:ss.ss", args: args{d: 359.999, degDigits: 3, precision: 2}, want: "359:59:56.40"},
		{name: "dd:mm:ss.ss", args: args{d: 359.9999, degDigits: 3, precision: 2}, want: "359:59:59.64"},
		{name: "dd:mm:ss.ss", args: args{d: 359.99999, degDigits: 3, precision: 2}, want: "359:59:59.96"},
		{name: "dd:mm:ss.ss", args: args{d: 359.999999, degDigits: 3, precision: 2}, want: "359:59:59.99"},
		{name: "dd:mm:ss.ss", args: args{d: 359.9999999, degDigits: 3, precision: 2}, want: "359:59:59.99"},
		// Precision 3
		{name: "dd:mm:ss.sss", args: args{d: 359.9, degDigits: 3, precision: 3}, want: "359:53:59.999"},
		{name: "dd:mm:ss.sss", args: args{d: 359.99, degDigits: 3, precision: 3}, want: "359:59:24.000"},
		{name: "dd:mm:ss.sss", args: args{d: 359.999, degDigits: 3, precision: 3}, want: "359:59:56.400"},
		{name: "dd:mm:ss.sss", args: args{d: 359.9999, degDigits: 3, precision: 3}, want: "359:59:59.640"},
		{name: "dd:mm:ss.sss", args: args{d: 359.99999, degDigits: 3, precision: 3}, want: "359:59:59.964"},
		{name: "dd:mm:ss.sss", args: args{d: 359.999999, degDigits: 3, precision: 3}, want: "359:59:59.996"},
		{name: "dd:mm:ss.sss", args: args{d: 359.9999999, degDigits: 3, precision: 3}, want: "359:59:59.999"},
		// Precision 4
		{name: "dd:mm:ss.ssss", args: args{d: 359.9, degDigits: 3, precision: 4}, want: "359:53:59.9999"},
		{name: "dd:mm:ss.ssss", args: args{d: 359.99, degDigits: 3, precision: 4}, want: "359:59:24.0000"},
		{name: "dd:mm:ss.ssss", args: args{d: 359.999, degDigits: 3, precision: 4}, want: "359:59:56.4000"},
		{name: "dd:mm:ss.ssss", args: args{d: 359.9999, degDigits: 3, precision: 4}, want: "359:59:59.6400"},
		{name: "dd:mm:ss.ssss", args: args{d: 359.99999, degDigits: 3, precision: 4}, want: "359:59:59.9640"},
		{name: "dd:mm:ss.ssss", args: args{d: 359.999999, degDigits: 3, precision: 4}, want: "359:59:59.9964"},
		{name: "dd:mm:ss.ssss", args: args{d: 359.9999999, degDigits: 3, precision: 4}, want: "359:59:59.9996"},
		// Precision 5
		{name: "dd:mm:ss.sssss", args: args{d: 359.9, degDigits: 3, precision: 5}, want: "359:53:59.99999"},
		{name: "dd:mm:ss.sssss", args: args{d: 359.99, degDigits: 3, precision: 5}, want: "359:59:24.00000"},
		{name: "dd:mm:ss.sssss", args: args{d: 359.999, degDigits: 3, precision: 5}, want: "359:59:56.40000"},
		{name: "dd:mm:ss.sssss", args: args{d: 359.9999, degDigits: 3, precision: 5}, want: "359:59:59.64000"},
		{name: "dd:mm:ss.sssss", args: args{d: 359.99999, degDigits: 3, precision: 5}, want: "359:59:59.96400"},
		{name: "dd:mm:ss.sssss", args: args{d: 359.999999, degDigits: 3, precision: 5}, want: "359:59:59.99640"},
		{name: "dd:mm:ss.sssss", args: args{d: 359.9999999, degDigits: 3, precision: 5}, want: "359:59:59.99963"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DegDMSStringExt(tt.args.d, tt.args.sign, "+", "-", tt.args.degDigits, tt.args.precision); got != tt.want {
				t.Errorf("DegDMSStringExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
