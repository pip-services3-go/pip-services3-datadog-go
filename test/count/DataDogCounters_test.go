package count_test

import (
	"os"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	ddcount "github.com/pip-services3-go/pip-services3-datadog-go/count"
	ddfixture "github.com/pip-services3-go/pip-services3-datadog-go/test/fixtures"

	"github.com/stretchr/testify/assert"
)

func TestDataDogCounters(t *testing.T) {
	var counters *ddcount.DataDogCounters
	var fixture *ddfixture.CountersFixture

	apiKey := os.Getenv("DATADOG_API_KEY")
	if apiKey == "" {
		apiKey = "3eb3355caf628d4689a72084425177ac"
	}

	counters = ddcount.NewDataDogCounters()
	fixture = ddfixture.NewCountersFixture(counters.CachedCounters)

	config := cconf.NewConfigParamsFromTuples(
		"source", "test",
		"credential.access_key", apiKey,
	)
	counters.Configure(config)

	err := counters.Open("")
	assert.Nil(t, err)

	defer counters.Close("")

	t.Run("Simple Counters", func(t *testing.T) {
		fixture.TestSimpleCounters(t)
	})

	t.Run("Measure Elapsed Time", func(t *testing.T) {
		fixture.TestMeasureElapsedTime(t)
	})

}
