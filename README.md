# Core-using-licenses

This repo contains examples on how to set up the Qlik Licenses service as well as runnable tests to verify the setups.

It also contains an example of how you could use the metrics of the Qlik Licenses service to monitor your license usage.

## Using the community version of Qlik Core without a license

The [docker-compose.only-engine.yml](./docker-compose.only-engine)` file contains an example setup of only the Qlik Associative Engine.

You start it by running the command `docker-compose -f docker-compose.only-engine.yml up -d`

You could then if you have go installed verify that the Qlik Associtaive Engine only allows five concurent sessions by running the command `go test test/no_license_test.go test/utils_test.go`

## Using Qlik Core with a license

The [docker-compose.engine-and-license-service.yml](./docker-compose.engine-and-license-service.yml) file contains an example setup of the Qlik Associative Engine and the Qlik Licenses service.

It contains the parameter `-S LicenseServiceUrl=http://licenses:9200` to the Qlik Associative Engine with the adress of the Qlik Licenses service as well as the environment variables `LICENSES_SERIAL_NBR` and `LICENSES_CONTROL_NBR` for the Qlik Licenses service.

You can start it with the command `docker-compose -f docker-compose.engine-and-license-service up -d` provided that you have populated the `LICENSES_SERIAL_NBR` and `LICENSES_CONTROL_NBR` environment variables with your license.

By running the command `go test test/with_license_test.go test/utils_test.go` you could then verify that with a license more than five concurent sessions could be created.

## Using the Qlik Licenses metrics to monitor your license usage
