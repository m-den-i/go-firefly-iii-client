# \PiggyBanksApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePiggyBank**](PiggyBanksApi.md#DeletePiggyBank) | **Delete** /v1/piggy-banks/{id} | Delete a piggy bank.
[**GetPiggyBank**](PiggyBanksApi.md#GetPiggyBank) | **Get** /v1/piggy-banks/{id} | Get a single piggy bank.
[**ListAttachmentByPiggyBank**](PiggyBanksApi.md#ListAttachmentByPiggyBank) | **Get** /v1/piggy-banks/{id}/attachments | Lists all attachments.
[**ListEventByPiggyBank**](PiggyBanksApi.md#ListEventByPiggyBank) | **Get** /v1/piggy-banks/{id}/events | List all events linked to a piggy bank.
[**ListPiggyBank**](PiggyBanksApi.md#ListPiggyBank) | **Get** /v1/piggy-banks | List all piggy banks.
[**StorePiggyBank**](PiggyBanksApi.md#StorePiggyBank) | **Post** /v1/piggy-banks | Store a new piggy bank
[**UpdatePiggyBank**](PiggyBanksApi.md#UpdatePiggyBank) | **Put** /v1/piggy-banks/{id} | Update existing piggy bank.



## DeletePiggyBank

> DeletePiggyBank(ctx, id).XTraceId(xTraceId).Execute()

Delete a piggy bank.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the piggy bank.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PiggyBanksApi.DeletePiggyBank(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiggyBanksApi.DeletePiggyBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the piggy bank. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePiggyBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPiggyBank

> PiggyBankSingle GetPiggyBank(ctx, id).XTraceId(xTraceId).Execute()

Get a single piggy bank.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the piggy bank.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiggyBanksApi.GetPiggyBank(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiggyBanksApi.GetPiggyBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPiggyBank`: PiggyBankSingle
    fmt.Fprintf(os.Stdout, "Response from `PiggyBanksApi.GetPiggyBank`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the piggy bank. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPiggyBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**PiggyBankSingle**](PiggyBankSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttachmentByPiggyBank

> AttachmentArray ListAttachmentByPiggyBank(ctx, id).XTraceId(xTraceId).Page(page).Execute()

Lists all attachments.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the piggy bank.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiggyBanksApi.ListAttachmentByPiggyBank(context.Background(), id).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiggyBanksApi.ListAttachmentByPiggyBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttachmentByPiggyBank`: AttachmentArray
    fmt.Fprintf(os.Stdout, "Response from `PiggyBanksApi.ListAttachmentByPiggyBank`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the piggy bank. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAttachmentByPiggyBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**AttachmentArray**](AttachmentArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventByPiggyBank

> PiggyBankEventArray ListEventByPiggyBank(ctx, id).XTraceId(xTraceId).Page(page).Execute()

List all events linked to a piggy bank.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the piggy bank
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiggyBanksApi.ListEventByPiggyBank(context.Background(), id).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiggyBanksApi.ListEventByPiggyBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventByPiggyBank`: PiggyBankEventArray
    fmt.Fprintf(os.Stdout, "Response from `PiggyBanksApi.ListEventByPiggyBank`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the piggy bank | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEventByPiggyBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**PiggyBankEventArray**](PiggyBankEventArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPiggyBank

> PiggyBankArray ListPiggyBank(ctx).XTraceId(xTraceId).Page(page).Execute()

List all piggy banks.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiggyBanksApi.ListPiggyBank(context.Background()).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiggyBanksApi.ListPiggyBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPiggyBank`: PiggyBankArray
    fmt.Fprintf(os.Stdout, "Response from `PiggyBanksApi.ListPiggyBank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPiggyBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**PiggyBankArray**](PiggyBankArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorePiggyBank

> PiggyBankSingle StorePiggyBank(ctx).PiggyBankStore(piggyBankStore).XTraceId(xTraceId).Execute()

Store a new piggy bank



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    piggyBankStore := *openapiclient.NewPiggyBankStore("New digital camera", "13", "123.45") // PiggyBankStore | JSON array or key=value pairs with the necessary piggy bank information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiggyBanksApi.StorePiggyBank(context.Background()).PiggyBankStore(piggyBankStore).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiggyBanksApi.StorePiggyBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorePiggyBank`: PiggyBankSingle
    fmt.Fprintf(os.Stdout, "Response from `PiggyBanksApi.StorePiggyBank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorePiggyBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **piggyBankStore** | [**PiggyBankStore**](PiggyBankStore.md) | JSON array or key&#x3D;value pairs with the necessary piggy bank information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**PiggyBankSingle**](PiggyBankSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePiggyBank

> PiggyBankSingle UpdatePiggyBank(ctx, id).PiggyBankUpdate(piggyBankUpdate).XTraceId(xTraceId).Execute()

Update existing piggy bank.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the piggy bank
    piggyBankUpdate := *openapiclient.NewPiggyBankUpdate() // PiggyBankUpdate | JSON array with updated piggy bank information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiggyBanksApi.UpdatePiggyBank(context.Background(), id).PiggyBankUpdate(piggyBankUpdate).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiggyBanksApi.UpdatePiggyBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePiggyBank`: PiggyBankSingle
    fmt.Fprintf(os.Stdout, "Response from `PiggyBanksApi.UpdatePiggyBank`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the piggy bank | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePiggyBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **piggyBankUpdate** | [**PiggyBankUpdate**](PiggyBankUpdate.md) | JSON array with updated piggy bank information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**PiggyBankSingle**](PiggyBankSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

