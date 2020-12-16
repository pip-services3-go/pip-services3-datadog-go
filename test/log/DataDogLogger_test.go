package log_test

import (
	"os"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	ddlog "github.com/pip-services3-go/pip-services3-datadog-go/log"
	ddfixture "github.com/pip-services3-go/pip-services3-datadog-go/test/fixtures"

	"github.com/stretchr/testify/assert"
)

func TestDataDogLogger(t *testing.T) {
	var logger *ddlog.DataDogLogger
	var fixture *ddfixture.LoggerFixture

	apiKey := os.Getenv("DATADOG_API_KEY")
	if apiKey == "" {
		apiKey = "3eb3355caf628d4689a72084425177ac"
	}

	logger = ddlog.NewDataDogLogger()
	fixture = ddfixture.NewLoggerFixture(logger.CachedLogger)

	config := cconf.NewConfigParamsFromTuples(
		"source", "test",
		"credential.access_key", apiKey,
	)
	logger.Configure(config)

	err := logger.Open("")
	assert.Nil(t, err)

	defer logger.Close("")

	t.Run("Log Level", func(t *testing.T) {
		fixture.TestLogLevel(t)
	})

	t.Run("Simple Logging", func(t *testing.T) {
		fixture.TestSimpleLogging(t)
	})

	t.Run("Error Logging", func(t *testing.T) {
		fixture.TestErrorLogging(t)
	})

}
