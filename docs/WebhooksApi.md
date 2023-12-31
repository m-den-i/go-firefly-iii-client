# \WebhooksApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /v1/webhooks/{id} | Delete a webhook.
[**DeleteWebhookMessage**](WebhooksApi.md#DeleteWebhookMessage) | **Delete** /v1/webhooks/{id}/messages/{messageId} | Delete a webhook message.
[**DeleteWebhookMessageAttempt**](WebhooksApi.md#DeleteWebhookMessageAttempt) | **Delete** /v1/webhooks/{id}/messages/{messageId}/attempts/{attemptId} | Delete a webhook attempt.
[**GetSingleWebhookMessage**](WebhooksApi.md#GetSingleWebhookMessage) | **Get** /v1/webhooks/{id}/messages/{messageId} | Get a single message from a webhook.
[**GetSingleWebhookMessageAttempt**](WebhooksApi.md#GetSingleWebhookMessageAttempt) | **Get** /v1/webhooks/{id}/messages/{messageId}/attempts/{attemptId} | Get a single failed attempt from a single webhook message.
[**GetWebhook**](WebhooksApi.md#GetWebhook) | **Get** /v1/webhooks/{id} | Get a single webhook.
[**GetWebhookMessageAttempts**](WebhooksApi.md#GetWebhookMessageAttempts) | **Get** /v1/webhooks/{id}/messages/{messageId}/attempts | Get all the failed attempts of a single webhook message.
[**GetWebhookMessages**](WebhooksApi.md#GetWebhookMessages) | **Get** /v1/webhooks/{id}/messages | Get all the messages of a single webhook.
[**ListWebhook**](WebhooksApi.md#ListWebhook) | **Get** /v1/webhooks | List all webhooks.
[**StoreWebhook**](WebhooksApi.md#StoreWebhook) | **Post** /v1/webhooks | Store a new webhook
[**SubmitWebook**](WebhooksApi.md#SubmitWebook) | **Post** /v1/webhooks/{id}/submit | Submit messages for a webhook.
[**UpdateWebhook**](WebhooksApi.md#UpdateWebhook) | **Put** /v1/webhooks/{id} | Update existing webhook.



## DeleteWebhook

> DeleteWebhook(ctx, id).XTraceId(xTraceId).Execute()

Delete a webhook.



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
    id := "123" // string | The webhook ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.DeleteWebhook(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## DeleteWebhookMessage

> DeleteWebhookMessage(ctx, id, messageId).XTraceId(xTraceId).Execute()

Delete a webhook message.



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
    id := "123" // string | The webhook ID.
    messageId := int32(1) // int32 | The webhook message ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.DeleteWebhookMessage(context.Background(), id, messageId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteWebhookMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 
**messageId** | **int32** | The webhook message ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookMessageRequest struct via the builder pattern


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


## DeleteWebhookMessageAttempt

> DeleteWebhookMessageAttempt(ctx, id, messageId, attemptId).XTraceId(xTraceId).Execute()

Delete a webhook attempt.



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
    id := "123" // string | The webhook ID.
    messageId := int32(1) // int32 | The webhook message ID.
    attemptId := int32(1) // int32 | The webhook message attempt ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.DeleteWebhookMessageAttempt(context.Background(), id, messageId, attemptId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteWebhookMessageAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 
**messageId** | **int32** | The webhook message ID. | 
**attemptId** | **int32** | The webhook message attempt ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookMessageAttemptRequest struct via the builder pattern


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


## GetSingleWebhookMessage

> WebhookMessageSingle GetSingleWebhookMessage(ctx, id, messageId).XTraceId(xTraceId).Execute()

Get a single message from a webhook.



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
    id := "123" // string | The webhook ID.
    messageId := int32(1) // int32 | The webhook message ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetSingleWebhookMessage(context.Background(), id, messageId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetSingleWebhookMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleWebhookMessage`: WebhookMessageSingle
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetSingleWebhookMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 
**messageId** | **int32** | The webhook message ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleWebhookMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**WebhookMessageSingle**](WebhookMessageSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleWebhookMessageAttempt

> WebhookAttemptSingle GetSingleWebhookMessageAttempt(ctx, id, messageId, attemptId).XTraceId(xTraceId).Execute()

Get a single failed attempt from a single webhook message.



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
    id := "123" // string | The webhook ID.
    messageId := int32(1) // int32 | The webhook message ID.
    attemptId := int32(1) // int32 | The webhook attempt ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetSingleWebhookMessageAttempt(context.Background(), id, messageId, attemptId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetSingleWebhookMessageAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleWebhookMessageAttempt`: WebhookAttemptSingle
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetSingleWebhookMessageAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 
**messageId** | **int32** | The webhook message ID. | 
**attemptId** | **int32** | The webhook attempt ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleWebhookMessageAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**WebhookAttemptSingle**](WebhookAttemptSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> WebhookSingle GetWebhook(ctx, id).XTraceId(xTraceId).Execute()

Get a single webhook.



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
    id := "123" // string | The webhook ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetWebhook(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhook`: WebhookSingle
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**WebhookSingle**](WebhookSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookMessageAttempts

> WebhookAttemptArray GetWebhookMessageAttempts(ctx, id, messageId).XTraceId(xTraceId).Page(page).Execute()

Get all the failed attempts of a single webhook message.



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
    id := "123" // string | The webhook ID.
    messageId := int32(1) // int32 | The webhook message ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is per 50 items. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetWebhookMessageAttempts(context.Background(), id, messageId).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhookMessageAttempts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookMessageAttempts`: WebhookAttemptArray
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhookMessageAttempts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 
**messageId** | **int32** | The webhook message ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookMessageAttemptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is per 50 items. | 

### Return type

[**WebhookAttemptArray**](WebhookAttemptArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookMessages

> WebhookMessageArray GetWebhookMessages(ctx, id).XTraceId(xTraceId).Execute()

Get all the messages of a single webhook.



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
    id := "123" // string | The webhook ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetWebhookMessages(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhookMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookMessages`: WebhookMessageArray
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhookMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**WebhookMessageArray**](WebhookMessageArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhook

> WebhookArray ListWebhook(ctx).XTraceId(xTraceId).Page(page).Execute()

List all webhooks.



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
    page := int32(1) // int32 | The page number, if necessary. The default pagination is 50, so 50 webhooks per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ListWebhook(context.Background()).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ListWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhook`: WebhookArray
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ListWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | The page number, if necessary. The default pagination is 50, so 50 webhooks per page. | 

### Return type

[**WebhookArray**](WebhookArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreWebhook

> WebhookSingle StoreWebhook(ctx).WebhookStore(webhookStore).XTraceId(xTraceId).Execute()

Store a new webhook



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
    webhookStore := *openapiclient.NewWebhookStore("Update magic mirror on new transaction", openapiclient.WebhookTrigger("STORE_TRANSACTION"), openapiclient.WebhookResponse("TRANSACTIONS"), openapiclient.WebhookDelivery("JSON"), "https://example.com") // WebhookStore | JSON array or key=value pairs with the necessary webhook information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.StoreWebhook(context.Background()).WebhookStore(webhookStore).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.StoreWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreWebhook`: WebhookSingle
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.StoreWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookStore** | [**WebhookStore**](WebhookStore.md) | JSON array or key&#x3D;value pairs with the necessary webhook information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**WebhookSingle**](WebhookSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitWebook

> SubmitWebook(ctx, id).XTraceId(xTraceId).Execute()

Submit messages for a webhook.



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
    id := "123" // string | The webhook ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.SubmitWebook(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.SubmitWebook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitWebookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> WebhookSingle UpdateWebhook(ctx, id).WebhookUpdate(webhookUpdate).XTraceId(xTraceId).Execute()

Update existing webhook.



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
    id := "123" // string | The webhook ID.
    webhookUpdate := *openapiclient.NewWebhookUpdate() // WebhookUpdate | JSON array with updated webhook information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.UpdateWebhook(context.Background(), id).WebhookUpdate(webhookUpdate).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhook`: WebhookSingle
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The webhook ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookUpdate** | [**WebhookUpdate**](WebhookUpdate.md) | JSON array with updated webhook information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**WebhookSingle**](WebhookSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

