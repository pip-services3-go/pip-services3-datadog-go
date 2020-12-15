package clients1

import "time"

type DataDogMetricPoint struct {
	Time  time.Time `json:"time"`
	value float64   `json:"value"`
}
