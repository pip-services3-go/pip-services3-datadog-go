package clients1

import (
	"time"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cauth "github.com/pip-services3-go/pip-services3-components-go/auth"
	rpcclient "github.com/pip-services3-go/pip-services3-rpc-go/clients"
)

type DataDogLogClient struct {
	*rpcclient.RestClient
	defaultConfig      *cconf.ConfigParams
	credentialResolver *cauth.CredentialResolver
}

func NewDataDogLogClient(config *cconf.ConfigParams) *DataDogLogClient {

	c := &DataDogLogClient{
		RestClient:         rpcclient.NewRestClient(),
		credentialResolver: cauth.NewEmptyCredentialResolver(),
	}
	c.defaultConfig = cconf.NewConfigParamsFromTuples(
		"connection.protocol", "https",
		"connection.host", "http-intake.logs.datadoghq.com",
		"connection.port", 443,
		"credential.internal_network", "true",
	)

	if config != nil {
		c.Configure(config)
	}
	c.BaseRoute = "v1"
	return c
}

func (c *DataDogLogClient) Configure(config *cconf.ConfigParams) {
	config = c.defaultConfig.Override(config)
	c.RestClient.Configure(config)
	c.credentialResolver.Configure(config)
}

func (c *DataDogLogClient) SetReferences(refs cref.IReferences) {
	c.RestClient.SetReferences(refs)
	c.credentialResolver.SetReferences(refs)
}

func (c *DataDogLogClient) Open(correlationId string) error {
	credential, err := c.credentialResolver.Lookup(correlationId)
	if err != nil {
		return err
	}

	if credential == nil || credential.AccessKey() == "" {
		err = cerr.NewConfigError(correlationId, "NO_ACCESS_KEY", "Missing access key in credentials")
		return err
	}
	if c.Headers.Value() == nil {
		c.Headers = *cdata.NewEmptyStringValueMap()
	}
	c.Headers.SetAsObject("DD-API-KEY", credential.AccessKey())
	return c.RestClient.Open(correlationId)
}

func (c *DataDogLogClient) convertTags(tags map[string]string) string {
	if tags == nil {
		return ""
	}

	builder := ""

	for key, val := range tags {
		if builder != "" {
			builder += ","
		}
		builder += key + ":" + val
	}
	return builder
}

func (c *DataDogLogClient) convertMessage(message DataDogLogMessage) interface{} {

	timestamp := message.Time
	if timestamp.IsZero() {
		timestamp = time.Now().UTC()
	}
	result := map[string]interface{}{
		"timestamp": cconv.StringConverter.ToString(timestamp),
		"service":   message.Service,
		"message":   message.Message,
	}

	if message.Status != "" {
		result["status"] = message.Status
	} else {
		result["status"] = "INFO"
	}

	if message.Source != "" {
		result["ddsource"] = message.Source
	} else {
		result["ddsource"] = "pip-services"
	}

	if message.Tags != nil {
		result["ddtags"] = c.convertTags(message.Tags)
	}
	if message.Host != "" {
		result["host"] = message.Host
	}
	if message.LoggerName != "" {
		result["logger.name"] = message.LoggerName
	}
	if message.ThreadName != "" {
		result["logger.thread_name"] = message.ThreadName
	}
	if message.ErrorMessage != "" {
		result["error.message"] = message.ErrorMessage
	}
	if message.ErrorKind != "" {
		result["error.kind"] = message.ErrorKind
	}
	if message.ErrorStack != "" {
		result["error.stack"] = message.ErrorStack
	}

	return result
}

func (c *DataDogLogClient) convertMessages(messages []DataDogLogMessage) []interface{} {
	result := make([]interface{}, 0)

	for _, msg := range messages {
		result = append(result, c.convertMessage(msg))
	}
	return result
}

func (c *DataDogLogClient) SendLogs(correlationId string, messages []DataDogLogMessage) error {
	data := c.convertMessages(messages)

	// Commented instrumentation because otherwise it will never stop sending logs...
	//let timing = c.instrument(correlationId, "datadog.send_logs");
	result, err := c.Call(nil, "post", "input", correlationId, nil, data)
	//timing.endTiming();
	_, err = c.InstrumentError(correlationId, "datadog.send_logs", err, result)
	return err

}
