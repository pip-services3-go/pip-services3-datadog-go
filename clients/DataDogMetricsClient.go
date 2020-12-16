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

type DataDogMetricsClient struct {
	*rpcclient.RestClient
	defaultConfig      *cconf.ConfigParams
	credentialResolver *cauth.CredentialResolver
}

func NewDataDogMetricsClient(config *cconf.ConfigParams) *DataDogMetricsClient {
	c := &DataDogMetricsClient{
		RestClient:         rpcclient.NewRestClient(),
		credentialResolver: cauth.NewEmptyCredentialResolver(),
	}

	c.defaultConfig = cconf.NewConfigParamsFromTuples(
		"connection.protocol", "https",
		"connection.host", "api.datadoghq.com",
		"connection.port", 443,
		"credential.internal_network", "true",
	)

	if config != nil {
		c.Configure(config)
	}
	c.BaseRoute = "api/v1"
	return c
}

func (c *DataDogMetricsClient) Configure(config *cconf.ConfigParams) {
	config = c.defaultConfig.Override(config)
	c.RestClient.Configure(config)
	c.credentialResolver.Configure(config)
}

func (c *DataDogMetricsClient) SetReferences(refs cref.IReferences) {
	c.RestClient.SetReferences(refs)
	c.credentialResolver.SetReferences(refs)
}

func (c *DataDogMetricsClient) Open(correlationId string) error {
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

func (c *DataDogMetricsClient) convertTags(tags map[string]string) string {
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

func (c *DataDogMetricsClient) convertPoints(points []DataDogMetricPoint) []interface{} {

	result := make([]interface{}, 0)
	for _, p := range points {
		point := make([]string, 2)
		var tm int64 = 0
		if p.Time.IsZero() {
			tm = time.Now().UTC().UnixNano() / (int64)(time.Second)
		} else {
			tm = p.Time.UnixNano() / (int64)(time.Second)
		}
		point[0] = cconv.StringConverter.ToString(tm)
		point[1] = cconv.StringConverter.ToString(p.Value)
		result = append(result, point)
	}

	return result
}

func (c *DataDogMetricsClient) convertMetric(metric DataDogMetric) map[string]interface{} {
	tags := metric.Tags

	if metric.Service != "" {
		if tags == nil {
			tags = make(map[string]string, 0)
		}
		tags["service"] = metric.Service
	}

	result := map[string]interface{}{
		"metric": metric.Metric,
		"points": c.convertPoints(metric.Points),
	}

	if metric.Type != "" {
		result["type"] = metric.Type
	} else {
		result["type"] = "gauge"
	}

	if tags != nil {
		result["tags"] = c.convertTags(tags)
	}
	if metric.Host != "" {
		result["host"] = metric.Host
	}
	if metric.Interval > 0 {
		result["interval"] = metric.Interval
	}

	return result
}

func (c *DataDogMetricsClient) convertMetrics(metrics []DataDogMetric) map[string]interface{} {

	series := make([]interface{}, 0)
	for _, metric := range metrics {
		series = append(series, c.convertMetric(metric))
	}
	return map[string]interface{}{
		"series": series,
	}
}

func (c *DataDogMetricsClient) SendMetrics(correlationId string, metrics []DataDogMetric) error {
	data := c.convertMetrics(metrics)

	// Commented instrumentation because otherwise it will never stop sending logs...
	// timing := c.Instrument(correlationId, "datadog.send_metrics");
	result, err := c.Call(nil, "post", "series", correlationId, nil, data)
	//timing.EndTiming();
	_, err = c.InstrumentError(correlationId, "datadog.send_metrics", err, result)
	return err
}
