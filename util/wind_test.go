package util

import (
	"testing"
)

func TestWindCompassDirection(t *testing.T) {
	t.Run("Test no index errors", func(t *testing.T) {
		for d := 0; d <= 360; d++ {
			_ = WindCompassDirection(float64(d))
		}
	})

	t.Run("Test negative angles", func(t *testing.T) {
		for d := -45; d <= 45; d++ {
			_ = WindCompassDirection(float64(d))
		}
	})

	t.Run("Test >360 angles", func(t *testing.T) {
		for d := 315; d <= 405; d++ {
			_ = WindCompassDirection(float64(d))
		}
	})
}
