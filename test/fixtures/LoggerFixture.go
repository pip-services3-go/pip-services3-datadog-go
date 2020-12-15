// let assert = require('chai').assert;

// import { LogLevel } from 'pip-services3-components-node';
// import { CachedLogger } from 'pip-services3-components-node';

// export class LoggerFixture {
//     private _logger: CachedLogger;

//     public constructor(logger: CachedLogger) {
//         this._logger = logger;
//     }

//     public testLogLevel() {
//         assert.isTrue(this._logger.getLevel() >= LogLevel.None);
//         assert.isTrue(this._logger.getLevel() <= LogLevel.Trace);
//     }

//     public testSimpleLogging(done) {
//         this._logger.setLevel(LogLevel.Trace);

//         this._logger.fatal("987", null, "Fatal error message");
//         this._logger.error("987", null, "Error message");
//         this._logger.warn("987", "Warning message");
//         this._logger.info("987", "Information message");
//         this._logger.debug("987", "Debug message");
//         this._logger.trace("987", "Trace message");

//         this._logger.dump();
//         setTimeout(done, 1000);
//     }

//     public testErrorLogging(done) {
//         try {
//             // Raise an exception
//             throw new Error();
//         } catch (ex) {
//             this._logger.fatal("123", ex, "Fatal error");
//             this._logger.error("123", ex, "Recoverable error");

//             assert.isNotNull(ex);
//         }

//         this._logger.dump();
//         setTimeout(done, 1000);
//     }
    
// }