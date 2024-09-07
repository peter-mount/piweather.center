package rainfall

type Record struct {
	// Tips number of bucket tips since last reading.
	Tips uint32 `json:"tips"`

	// Hour mm of rain in last hour
	Hour float64 `json:"hour"`

	// Day mm of rain in last 24 hours
	Day float64 `json:"day"`

	// Total mm of rain since the sensor was powered on
	Total float64 `json:"total"`

	// BucketCount number of bucket tips since sensor was powered on
	BucketCount uint32 `json:"bucket_count"`

	// Duration time in seconds since last reading.
	Duration uint32 `json:"duration"`
}
