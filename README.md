# discovery

[Godoc](http://godoc.org/github.com/Clever/discovery-go)

## import

```go
import "gopkg.in/Clever/discovery-go.v1"
```

## Motivation

This library was designed to have a simple way to expose dependencies to a service in an extendable way that allows services to expose multiple interfaces.

## Description

Discovery Go is a standardized library for obtaining a URL to another service given the name of that service and the interface you are connecting to.
This library is meant to be used in an environment where the service dependencies needed have been provided in environment variables.
The URL for a service is expected to be in the form `SERVICE_{ SERVICE_NAME }_{ INTERFACE }_URL` in an environment variable.

## Example

Let's imagine a service has a dependency on influxdb and needs to send data to the [protobuf TCP interface](https://github.com/Clever/influxdb-service/blob/master/launch/influxdb-service.yml#L20-L23).
This service can obtain the proper URL to access this interface by specifying `influxdb-service` and `protobuf` in their call.

```go
import discovery "gopkg.in/Clever/discovery-go.v1"

var influxDbURL string

func init() {
    var err error
    influxDbURL, err = discovery.URLString("influxdb-service", "protobuf")
    if err != nil {
        log.Fatalf("Failed to obtain URL for influxdb, err: %s", err)
    }
}
```
