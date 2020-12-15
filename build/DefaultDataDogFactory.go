// /** @module build */
// import { Factory } from 'pip-services3-components-node';
// import { Descriptor } from 'pip-services3-commons-node';

// import { DataDogLogger } from '../log/DataDogLogger';

// /**
//  * Creates DataDog components by their descriptors.
//  * 
//  * @see [[DataDogLogger]]
//  */
// export class DefaultDataDogFactory extends Factory {
// 	public static readonly Descriptor = new Descriptor("pip-services", "factory", "datadog", "default", "1.0");
// 	public static readonly DataDogLoggerDescriptor = new Descriptor("pip-services", "logger", "datadog", "*", "1.0");

// 	/**
// 	 * Create a new instance of the factory.
// 	 */
// 	public constructor() {
//         super();
// 		this.registerAsType(DefaultDataDogFactory.DataDogLoggerDescriptor, DataDogLogger);
// 	}
// }