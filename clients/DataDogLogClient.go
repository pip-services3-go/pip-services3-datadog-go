package clients1

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
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

//     func (c*DataDogLogClient) Open(correlationId string) error {
//         c.credentialResolver.Lookup(correlationId, (err, credential) => {
//             if (err) {
//                 callback(err);
//                 return;
//             }

//             if (credential == null || credential.getAccessKey() == null) {
//                 err = new ConfigException(correlationId, "NO_ACCESS_KEY", "Missing access key in credentials");
//                 callback(err);
//                 return;
//             }

//             c.headers = c.headers || {};
//             c.headers["DD-API-KEY"] = credential.getAccessKey();

//             super.open(correlationId, callback);
//         });
//     }

//     func (c*DataDogLogClient) convertTags(tags []interface{}) string {
//         if (tags == null) return null;

//         let builder: string = "";

//         for (let key in tags) {
//             if (builder != "")
//                 builder += ",";
//             builder += key + ":" + tags[key];
//         }
//         return builder;
//     }

//     func (c*DataDogLogClient) convertMessage(message DataDogLogMessage) interface{} {
//          result := map[string]interface{}{
//             "timestamp": StringConverter.ToString(message.time || new Date()),
//             "status": message.status || "INFO",
//             "ddsource": message.source || "pip-services",
// //            "source": message.source || "pip-services",
//             "service": message.service,
//             "message": message.message,
//         };

//         if (message.tags)
//             result["ddtags"] = c.convertTags(message.tags);
//         if (message.host)
//             result["host"] = message.host;
//         if (message.logger_name)
//             result["logger.name"] = message.logger_name;
//         if (message.thread_name)
//             result["logger.thread_name"] = message.thread_name;
//         if (message.error_message)
//             result["error.message"] = message.error_message;
//         if (message.error_kind)
//             result["error.kind"] = message.error_kind;
//         if (message.error_stack)
//             result["error.stack"] = message.error_stack;

//         return result;
//     }

//     func (c*DataDogLogClient) convertMessages(messages []DataDogLogMessage)[] interface{} {
//         return _.map(messages, (m) => {return c.convertMessage(m);});
//     }

//     func (c*DataDogLogClient) SendLogs(correlationId string, messages []DataDogLogMessage) error {
//         let data = c.convertMessages(messages);

//         // Commented instrumentation because otherwise it will never stop sending logs...
//         //let timing = c.instrument(correlationId, "datadog.send_logs");
//         c.Call("post", "input", null, null, data, (err, result) => {
//             //timing.endTiming();
//             c.instrumentError(correlationId, "datadog.send_logs", err, result, callback);
//         });
//     }
