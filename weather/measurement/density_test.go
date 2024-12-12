package measurement

import "testing"

func Test_density(t *testing.T) {
	testConversions(t, []conversionTest{
		newConversionTest(GramsPerCubicMeter.Value(1), PoundsPerCubitFoot.Value(16018.463374), false),
		newConversionTest(PoundsPerCubitFoot.Value(1), GramsPerCubicMeter.Value(0.0000624279605761), false),
	})
}
