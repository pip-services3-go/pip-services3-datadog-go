import { ConfigParams } from 'pip-services3-commons-node';
import { References } from 'pip-services3-commons-node';

import { DataDogCounters } from '../../src/count/DataDogCounters';
import { CountersFixture } from '../fixtures/CountersFixture';

suite('DataDogCounters', ()=> {
    let _counters: DataDogCounters;
    let _fixture: CountersFixture;

    setup((done) => {
        let apiKey = process.env['DATADOG_API_KEY'] || '3eb3355caf628d4689a72084425177ac';

        _counters = new DataDogCounters();
        _fixture = new CountersFixture(_counters);

        let config = ConfigParams.fromTuples(
            'source', 'test',
            'credential.access_key', apiKey
        );
        _counters.configure(config);

        _counters.open(null, (err) => {
             done(err);
        });
    });

    teardown((done) => {
        _counters.close(null, done);
    });

    test('Simple Counters', (done) => {
        _fixture.testSimpleCounters(done);
    });

    test('Measure Elapsed Time', (done) => {
        _fixture.testMeasureElapsedTime(done);
    });

});