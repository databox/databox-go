# Go API client for databox

Push API resources Open API documentation

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.4.0
- Package version: 0.4.0
- Generator version: 7.6.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import databox "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `databox.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), databox.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `databox.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), databox.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `databox.ContextOperationServerIndices` and `databox.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), databox.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), databox.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://push.databox.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**DataDelete**](docs/DefaultAPI.md#datadelete) | **Delete** /data | 
*DefaultAPI* | [**DataMetricKeyDelete**](docs/DefaultAPI.md#datametrickeydelete) | **Delete** /data/{metricKey} | 
*DefaultAPI* | [**DataPost**](docs/DefaultAPI.md#datapost) | **Post** /data | 
*DefaultAPI* | [**MetrickeysGet**](docs/DefaultAPI.md#metrickeysget) | **Get** /metrickeys | 
*DefaultAPI* | [**MetrickeysPost**](docs/DefaultAPI.md#metrickeyspost) | **Post** /metrickeys | 
*DefaultAPI* | [**PingGet**](docs/DefaultAPI.md#pingget) | **Get** /ping | 


## Documentation For Models

 - [ApiResponse](docs/ApiResponse.md)
 - [PushData](docs/PushData.md)
 - [PushDataAttribute](docs/PushDataAttribute.md)
 - [State](docs/State.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### basicAuth

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), databox.ContextBasicAuth, databox.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



