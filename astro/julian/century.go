package julian

type Century float64

func (c Century) Float() float64 {
	return float64(c)
}

func (c Century) ToDay() Day {
	return c.ToYear().ToDay()
}

func (c Century) ToYear() Year {
	return Year(float64(c) * JulianCentury)
}

func (c Century) ToCentury() Century {
	return c
}
