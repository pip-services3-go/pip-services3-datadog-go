package clients_test

import (
	"os"
	"testing"
	"time"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	rnd "github.com/pip-services3-go/pip-services3-commons-go/random"
	clients1 "github.com/pip-services3-go/pip-services3-datadog-go/clients"
	"github.com/stretchr/testify/assert"
)

func TestDataDogMetricClient(t *testing.T) {
	var client *clients1.DataDogMetricsClient

	apiKey := os.Getenv("DATADOG_API_KEY")
	if apiKey == "" {
		apiKey = "3eb3355caf628d4689a72084425177ac"
	}

	client = clients1.NewDataDogMetricsClient(nil)

	config := cconf.NewConfigParamsFromTuples(
		"source", "test",
		"credential.access_key", apiKey,
	)
	client.Configure(config)

	err := client.Open("")
	assert.Nil(t, err)

	defer client.Close("")

	t.Run("Send Metrics", func(t *testing.T) {
		metrics := []clients1.DataDogMetric{
			{
				Metric:  "test.metric.1",
				Service: "TestService Golang",
				Host:    "TestHost",
				Type:    clients1.Gauge,
				Points: []clients1.DataDogMetricPoint{
					{
						Time:  time.Now().UTC(),
						Value: rnd.RandomDouble.NextDouble(0, 100),
					},
				},
			},
			{
				Metric:   "test.metric.2",
				Service:  "TestService Golang",
				Host:     "TestHost",
				Type:     clients1.Rate,
				Interval: 100,
				Points: []clients1.DataDogMetricPoint{
					{
						Time:  time.Now().UTC(),
						Value: rnd.RandomDouble.NextDouble(0, 100),
					},
				},
			},
			{
				Metric:   "test.metric.3",
				Service:  "TestService Golang",
				Host:     "TestHost",
				Type:     clients1.Count,
				Interval: 100,
				Points: []clients1.DataDogMetricPoint{
					{
						Time:  time.Now().UTC(),
						Value: rnd.RandomDouble.NextDouble(0, 100),
					},
				},
			},
		}

		err := client.SendMetrics("", metrics)
		assert.Nil(t, err)

	})

}
