# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> DataDog components for Golang

This module is a part of the [Pip.Services](http://pipservices.org) polyglot microservices toolkit.
It contains the DataDog logger and performance counters components.

The module contains the following packages:
- [**Build**](build) - contains a class used to create DataDog components by their descriptors.
- [**Clients**](clients) - contains constants and classes used to define REST clients for DataDog
- [**Count**](count) - contains a class used to create performance counters that send their metrics to a DataDog service
- [**Log**](log) - contains a class used to create loggers that dump execution logs to a DataDog service.

<a name="links"></a> Quick links:

* [Configuration](https://www.pipservices.org/recipies/configuration)
* [API Reference](https://pip-services3-node.github.io/pip-services3-datadog-node/globals.html)
* [Change Log](CHANGELOG.md)
* [Get Help](https://www.pipservices.org/community/help)
* [Contribute](https://www.pipservices.org/community/contribute)

## Use

Install the NPM package as
```bash
npm install pip-services-datadog-node --save
```

## Develop

For development you shall install the following prerequisites:
* Node.js 8+
* Visual Studio Code or another IDE of your choice
* Docker
* Typescript

Install dependencies:
```bash
npm install
```

Compile the code:
```bash
tsc
```

Run automated tests:
```bash
npm test
```

Generate API documentation:
```bash
./docgen.ps1
```

Before committing changes run dockerized build and test as:
```bash
./build.ps1
./test.ps1
./clear.ps1
```

## Contacts

The library is created and maintained by **Sergey Seroukhov**.

The documentation is written by:
- **Mark Makarychev**