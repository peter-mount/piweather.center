package station

import (
	"context"
)

// Sensor is the common interface for Reading and CalculatedValue
type Sensor interface {
	WithContext(ctx context.Context) (context.Context, error)
	Sensors() *Sensors
	Accept(v Visitor) error
	GetID() string
	IsCalculated() bool
	IsPseudo() bool
}
