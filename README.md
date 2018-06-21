# Qlik Core Licensing Examples

[![CircleCI](https://circleci.com/gh/qlik-oss/core-using-licenses.svg?style=shield)](https://circleci.com/gh/qlik-oss/core-using-licenses)

This repository contains several examples that show you how to set up the Qlik Licenses service with or without a license, and how to use licenses service metrics to monitor your license usage. Each example includes a runnable test to verify the setup. To run these tests, you must have Go installed.

## Using Qlik Core community version without a license

The [docker-compose.only-engine.yml](./docker-compose.only-engine) file contains an example configuration of the Qlik Associative Engine without a license.

To start it, run the following command:

```bash
ACCEPT_EULA=<yes/no> docker-compose -f docker-compose.only-engine.yml up -d
```

You can verify that the Qlik Associtaive Engine only allows five concurrent sessions by running the following command:

```bash
go test test/no_license_test.go test/utils_test.go
```

## Using Qlik Core with a license

The [docker-compose.engine-and-license-service.yml](./docker-compose.engine-and-license-service.yml) file contains an example configuration of the Qlik Associative Engine and the Qlik Licenses service.

The `yml` file contains an address parameter: `-S LicenseServiceUrl=http://licenses:9200`. This tells Qlik Associative Engine where to find the licenses service.

The `yml` file also contains two environment variables: `LICENSES_SERIAL_NBR` and `LICENSES_CONTROL_NBR`. You need to replace these variables with your license serial number and license control number.

To start it, run the following command:

```bash
ACCEPT_EULA=<yes/no> docker-compose -f docker-compose.engine-and-license-service.yml up -d
```

A valid license allows you to run more than five concurrent sessions. You can verify the license by running the follow command:

```bash
go test test/with_license_test.go test/utils_test.go
```

## Using the Qlik Licenses service metrics to monitor your license usage

Both the Qlik Associative Engine and the Qlik Licenses service expose metrics endpoints that can you can use to monitor the current license status.

The [docker-compose.metrics.yaml](./docker-compose.metrics.yml) sets up [Prometheus](https://prometheus.io/) and [Grafana](https://grafana.com/). In this example, these two services will be configured on startup. Prometheus is set up to scrape the relevant license metrics and Grafana is set up with a preconfigured dashboard for license monitoring.

To start the monitoring example, run the following command:

```bash
ACCEPT_EULA=<yes/no> docker-compose -f docker-compose.engine-and-license-service.yml -f docker-compose.metrics.yml up -d
```

You should now be able to monitor the current license consumption in the preconfigured Grafana dashboard [here](http://localhost:3000/d/license_monitoring/qlik-core-licensing-monitoring?refresh=5s&orgId=1).

By default, the dashboard updates every 5 seconds. You can try out the monitoring by either using the test case mentioned in previous section, or by opening sessions using [enigma-go](https://github.com/qlik-oss/enigma-go) or [enigma.js](https://github.com/qlik-oss/enigma.js).
