# \ConfigurationApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](ConfigurationApi.md#GetConfiguration) | **Get** /v1/configuration | Get Firefly III system configuration values.
[**GetSingleConfiguration**](ConfigurationApi.md#GetSingleConfiguration) | **Get** /v1/configuration/{name} | Get a single Firefly III system configuration value
[**SetConfiguration**](ConfigurationApi.md#SetConfiguration) | **Put** /v1/configuration/{name} | Update configuration value



## GetConfiguration

> []Configuration GetConfiguration(ctx).XTraceId(xTraceId).Execute()

Get Firefly III system configuration values.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/m-den-i/go-firefly-iii-client"
)

func main() {
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.GetConfiguration(context.Background()).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration`: []Configuration
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**[]Configuration**](Configuration.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleConfiguration

> ConfigurationSingle GetSingleConfiguration(ctx, name).XTraceId(xTraceId).Execute()

Get a single Firefly III system configuration value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/m-den-i/go-firefly-iii-client"
)

func main() {
    name := openapiclient.ConfigValueFilter("configuration.is_demo_site") // ConfigValueFilter | The name of the configuration value you want to know.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.GetSingleConfiguration(context.Background(), name).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetSingleConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleConfiguration`: ConfigurationSingle
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetSingleConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | [**ConfigValueFilter**](.md) | The name of the configuration value you want to know. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**ConfigurationSingle**](ConfigurationSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration

> ConfigurationSingle SetConfiguration(ctx, name).Value(value).XTraceId(xTraceId).Execute()

Update configuration value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/m-den-i/go-firefly-iii-client"
)

func main() {
    name := openapiclient.ConfigValueUpdateFilter("configuration.is_demo_site") // ConfigValueUpdateFilter | The name of the configuration value you want to update.
    value := openapiclient.PolymorphicProperty{ArrayOfString: new([]string)} // PolymorphicProperty | 
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.SetConfiguration(context.Background(), name).Value(value).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.SetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetConfiguration`: ConfigurationSingle
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.SetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | [**ConfigValueUpdateFilter**](.md) | The name of the configuration value you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **value** | [**PolymorphicProperty**](PolymorphicProperty.md) |  | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**ConfigurationSingle**](ConfigurationSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

