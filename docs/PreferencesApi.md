# \PreferencesApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPreference**](PreferencesApi.md#GetPreference) | **Get** /v1/preferences/{name} | Return a single preference.
[**ListPreference**](PreferencesApi.md#ListPreference) | **Get** /v1/preferences | List all users preferences.
[**StorePreference**](PreferencesApi.md#StorePreference) | **Post** /v1/preferences | Store a new preference for this user.
[**UpdatePreference**](PreferencesApi.md#UpdatePreference) | **Put** /v1/preferences/{name} | Update preference



## GetPreference

> PreferenceSingle GetPreference(ctx, name).XTraceId(xTraceId).Execute()

Return a single preference.



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
    name := "currencyPreference" // string | The name of the preference.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PreferencesApi.GetPreference(context.Background(), name).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesApi.GetPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreference`: PreferenceSingle
    fmt.Fprintf(os.Stdout, "Response from `PreferencesApi.GetPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the preference. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**PreferenceSingle**](PreferenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPreference

> PreferenceArray ListPreference(ctx).XTraceId(xTraceId).Page(page).Execute()

List all users preferences.



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
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PreferencesApi.ListPreference(context.Background()).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesApi.ListPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPreference`: PreferenceArray
    fmt.Fprintf(os.Stdout, "Response from `PreferencesApi.ListPreference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**PreferenceArray**](PreferenceArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorePreference

> PreferenceSingle StorePreference(ctx).Preference(preference).XTraceId(xTraceId).Execute()

Store a new preference for this user.



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
    preference := *openapiclient.NewPreference("currencyPreference", openapiclient.PolymorphicProperty{ArrayOfString: new([]string)}) // Preference | JSON array with the necessary preference information or key=value pairs. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PreferencesApi.StorePreference(context.Background()).Preference(preference).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesApi.StorePreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorePreference`: PreferenceSingle
    fmt.Fprintf(os.Stdout, "Response from `PreferencesApi.StorePreference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorePreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preference** | [**Preference**](Preference.md) | JSON array with the necessary preference information or key&#x3D;value pairs. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**PreferenceSingle**](PreferenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePreference

> PreferenceSingle UpdatePreference(ctx, name).PreferenceUpdate(preferenceUpdate).XTraceId(xTraceId).Execute()

Update preference



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
    name := "currencyPreference" // string | The name of the preference. Will always overwrite. Will be created if it does not exist.
    preferenceUpdate := *openapiclient.NewPreferenceUpdate(openapiclient.PolymorphicProperty{ArrayOfString: new([]string)}) // PreferenceUpdate | JSON array or key=value pairs with the necessary preference information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PreferencesApi.UpdatePreference(context.Background(), name).PreferenceUpdate(preferenceUpdate).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesApi.UpdatePreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePreference`: PreferenceSingle
    fmt.Fprintf(os.Stdout, "Response from `PreferencesApi.UpdatePreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the preference. Will always overwrite. Will be created if it does not exist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preferenceUpdate** | [**PreferenceUpdate**](PreferenceUpdate.md) | JSON array or key&#x3D;value pairs with the necessary preference information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**PreferenceSingle**](PreferenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

