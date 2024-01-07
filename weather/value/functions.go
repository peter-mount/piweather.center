package value

import (
	"math"
)

func init() {
	NewCalculator("max", Basic2ArgCalculator(math.Max))
	NewCalculator("min", Basic2ArgCalculator(math.Min))
	NewCalculator("ceil", Basic1ArgCalculator(math.Ceil))
	NewCalculator("floor", Basic1ArgCalculator(math.Floor))
}
