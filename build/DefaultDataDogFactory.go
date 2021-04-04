package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	count "github.com/pip-services3-go/pip-services3-datadog-go/count"
	log "github.com/pip-services3-go/pip-services3-datadog-go/log"
)

/*
DefaultDataDogFactory are creates DataDog components by their descriptors.
See DataDogLogger
*/
type DefaultDataDogFactory struct {
	cbuild.Factory
}

// NewDefaultDataDogFactory create a new instance of the factory.
// Retruns *DefaultDataDogFactory
// pointer on new factory
func NewDefaultDataDogFactory() *DefaultDataDogFactory {
	c := DefaultDataDogFactory{}
	c.Factory = *cbuild.NewFactory()
	dataDogLoggerDescriptor := cref.NewDescriptor("pip-services", "logger", "datadog", "*", "1.0")
	dataDogCountersDescriptor := cref.NewDescriptor("pip-services", "counters", "datadog", "*", "1.0")

	c.RegisterType(dataDogLoggerDescriptor, log.NewDataDogLogger)
	c.RegisterType(dataDogCountersDescriptor, count.NewDataDogCounters)

	return &c
}
