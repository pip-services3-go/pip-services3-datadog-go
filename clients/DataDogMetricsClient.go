package clients1

// /** @module clients */

// const _ = require('lodash');

// import { ConfigParams } from 'pip-services3-commons-node';
// import { ConfigException } from 'pip-services3-commons-node';
// import { IReferences } from 'pip-services3-commons-node';
// import { StringConverter } from 'pip-services3-commons-node';
// import { CredentialResolver } from 'pip-services3-components-node';
// import { RestClient, } from 'pip-services3-rpc-node';

// import { DataDogMetric } from './DataDogMetric';
// import { DataDogMetricPoint} from './DataDogMetricPoint';

// export class DataDogMetricsClient extends RestClient {
//     private _defaultConfig: ConfigParams = ConfigParams.fromTuples(
//         "connection.protocol", "https",
//         "connection.host", "api.datadoghq.com",
//         "connection.port", 443,
//         "credential.internal_network", "true"
//     );
//     private _credentialResolver: CredentialResolver = new CredentialResolver();

//     public constructor(config?: ConfigParams) {
//         super();

//         if (config) this.configure(config);
//         this._baseRoute = "api/v1";
//     }   

//     public configure(config: ConfigParams): void {
//         config = this._defaultConfig.override(config);
//         super.configure(config);
//         this._credentialResolver.configure(config);
//     }

//     public setReferences(refs: IReferences): void {
//         super.setReferences(refs);
//         this._credentialResolver.setReferences(refs);
//     }

//     public open(correlationId: string, callback: (err: any) => void): void {
//         this._credentialResolver.lookup(correlationId, (err, credential) => {
//             if (err) {
//                 callback(err);
//                 return;
//             }

//             if (credential == null || credential.getAccessKey() == null) {
//                 err = new ConfigException(correlationId, "NO_ACCESS_KEY", "Missing access key in credentials");
//                 callback(err);
//                 return;
//             }

//             this._headers = this._headers || {};
//             this._headers['DD-API-KEY'] = credential.getAccessKey();

//             super.open(correlationId, callback);
//         });    
//     }

//     private convertTags(tags: any[]): string {
//         if (tags == null) return null;

//         let builder: string = "";

//         for (let key in tags) {
//             if (builder != "")
//                 builder += ",";
//             builder += key + ":" + tags[key];
//         }
//         return builder;
//     }

//     private convertPoints(points: DataDogMetricPoint[]) {
//         let result = _.map(
//             points,
//             (p) => {
//                 let time = (p.time || new Date()).getTime() / 1000;
//                 return [
//                     StringConverter.toString(time),
//                     StringConverter.toString(p.value)
//                 ]; 
//             }
//         );
//         return result;
//     }

//     private convertMetric(metric: DataDogMetric): any {
//         let tags = metric.tags;

//         if (metric.service) {
//             tags = tags || {};
//             tags.service = metric.service;
//         }

//         let result = {
//             "metric": metric.metric,
//             "type": metric.type || 'gauge',
//             "points": this.convertPoints(metric.points),
//         };

//         if (tags)
//             result['tags'] = this.convertTags(tags);
//         if (metric.host)
//             result['host'] = metric.host;
//         if (metric.interval)
//             result['interval'] = metric.interval;

//         return result;
//     }

//     private convertMetrics(metrics: DataDogMetric[]): any {
//         let series = _.map(metrics, (m) => {return this.convertMetric(m);});
//         return {
//             series: series
//         };
//     }

//     public sendMetrics(correlationId: string, metrics: DataDogMetric[],
//         callback: (err: any) => void): void {
//         let data = this.convertMetrics(metrics);

//         // Commented instrumentation because otherwise it will never stop sending logs...
//         //let timing = this.instrument(correlationId, "datadog.send_metrics");
//         this.call("post", "series", null, null, data, (err, result) => {
//             //timing.endTiming();
//             this.instrumentError(correlationId, "datadog.send_metrics", err, result, callback);
//         });
//     }
// }