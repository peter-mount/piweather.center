package measurement

import (
	"testing"
)

func Test_concentration(t *testing.T) {
	testConversions(t, []conversionTest{
		newConversionTest(PartsPerMillion.Value(1), MicrogramsPerCubicMeter.Value(1000), false),
		newConversionTest(MicrogramsPerCubicMeter.Value(1), PartsPerMillion.Value(0.001), false),
	})
}
