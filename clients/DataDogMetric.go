package clients1

type DataDogMetric struct {
	Metric   string               `json:"metric"`
	Service  string               `json:"service"`
	Host     string               `json:"host"`
	Tags     interface{}          `json:"tags"`
	Type     string               `json:"type"`
	Interval int64                `json:"interval"`
	Points   []DataDogMetricPoint `json:"points"`
}
