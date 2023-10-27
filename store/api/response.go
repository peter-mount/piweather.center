package api

type Response struct {
	Status  int           `json:"status,omitempty" xml:"status,attr,omitempty"`
	Message string        `json:"message,omitempty" xml:"message,omitempty"`
	Source  string        `json:"source,omitempty" xml:"source,omitempty"`
	Metric  string        `json:"metric,omitempty" xml:"metric,omitempty"`
	Results []MetricValue `json:"results,omitempty" xml:"results,omitempty"`
	Result  *MetricValue  `json:"result,omitempty" xml:"result,omitempty"`
	Metrics []Metric      `json:"metrics,omitempty" xml:"metrics,omitempty"`
}
