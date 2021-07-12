# Qlik Core Licensing Examples

*As of 1 July 2020, Qlik Core is no longer available to new customers. No further maintenance will be done in this repository.*

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
go test test/no_license_test.go test/utils_test.go -count=1
```

## Using Qlik Core with a license

The [docker-compose.engine-and-license-service.yml](./docker-compose.engine-and-license-service.yml) file contains an example configuration of the Qlik Associative Engine and the Qlik Licenses service.

The `yml` file contains an address parameter: `-S LicenseServiceUrl=http://licenses:9200`. This tells Qlik Associative Engine where to find the licenses service.

The `yml` file also contains an environment variable `LICENSE_KEY`. This variable should be set to your license key. You also need to specify your cost per session and your total license time as specified in your license on line [74 and 75](https://github.com/qlik-oss/core-using-licenses/blob/26c32f497c0973f69ad6122ffa2bdd0ce6e4b531/test/with_license_test.go#L74) in the file `test/with_license_test.go`

To start it, run the following command:

```bash
ACCEPT_EULA=<yes/no> docker-compose -f docker-compose.engine-and-license-service.yml up -d
```

A valid license allows you to run more than five concurrent sessions. You can verify the license and that the correct amount of the license has been consumed by running the follow command:

```bash
go test test/with_license_test.go test/utils_test.go -count=1
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
