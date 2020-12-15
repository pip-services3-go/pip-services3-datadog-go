// import { ConfigParams } from 'pip-services3-commons-node';

// import { DataDogLogger } from '../../src/log/DataDogLogger';
// import { LoggerFixture } from '../fixtures/LoggerFixture';

// let assert = require('chai').assert;
// let async = require('async');

// suite('DataDogLogger', () => {
//     let _logger: DataDogLogger;
//     let _fixture: LoggerFixture;

//     setup((done) => {
//         let apiKey = process.env['DATADOG_API_KEY'] || '3eb3355caf628d4689a72084425177ac';

//         _logger = new DataDogLogger();
//         _fixture = new LoggerFixture(_logger);

//         let config = ConfigParams.fromTuples(
//             'source', 'test',
//             'credential.access_key', apiKey
//         );
//         _logger.configure(config);

//         _logger.open(null, (err) => {
//             done(err);
//         });
//     });

//     teardown((done) => {
//         _logger.close(null, done);
//     });

//     test('Log Level', () => {
//         _fixture.testLogLevel();
//     });

//     test('Simple Logging', (done) => {
//         _fixture.testSimpleLogging(done);
//     });

//     test('Error Logging', (done) => {
//         _fixture.testErrorLogging(done);
//     });

// });