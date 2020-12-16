package clients_test

import (
	"os"
	"testing"
	"time"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	clients1 "github.com/pip-services3-go/pip-services3-datadog-go/clients"
	"github.com/stretchr/testify/assert"
)

func TestDataDogLogClient(t *testing.T) {
	var client *clients1.DataDogLogClient

	apiKey := os.Getenv("DATADOG_API_KEY")
	if apiKey == "" {
		apiKey = "3eb3355caf628d4689a72084425177ac"
	}

	client = clients1.NewDataDogLogClient(nil)

	config := cconf.NewConfigParamsFromTuples(
		"source", "test",
		"credential.access_key", apiKey,
	)
	client.Configure(config)

	err := client.Open("")
	assert.Nil(t, err)

	defer client.Close("")

	t.Run("Send Logs", func(t *testing.T) {
		messages := []clients1.DataDogLogMessage{
			clients1.DataDogLogMessage{
				Time:    time.Now().UTC(),
				Service: "TestService",
				Host:    "TestHost",
				Status:  clients1.Debug,
				Message: "Test trace message",
			},
			clients1.DataDogLogMessage{
				Time:    time.Now().UTC(),
				Service: "TestService",
				Host:    "TestHost",
				Status:  clients1.Info,
				Message: "Test info message",
			},
			clients1.DataDogLogMessage{
				Time:       time.Now().UTC(),
				Service:    "TestService",
				Host:       "TestHost",
				Status:     clients1.Error,
				Message:    "Test error message",
				ErrorKind:  "Exception",
				ErrorStack: "Stack trace...",
			},
			clients1.DataDogLogMessage{
				Time:       time.Now().UTC(),
				Service:    "TestService",
				Host:       "TestHost",
				Status:     clients1.Emergency,
				Message:    "Test fatal message",
				ErrorKind:  "Exception",
				ErrorStack: "Stack trace...",
			},
		}

		err := client.SendLogs("", messages)
		assert.Nil(t, err)

	})

}
