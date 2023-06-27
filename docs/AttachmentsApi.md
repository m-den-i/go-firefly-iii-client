# \AttachmentsApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAttachment**](AttachmentsApi.md#DeleteAttachment) | **Delete** /v1/attachments/{id} | Delete an attachment.
[**DownloadAttachment**](AttachmentsApi.md#DownloadAttachment) | **Get** /v1/attachments/{id}/download | Download a single attachment.
[**GetAttachment**](AttachmentsApi.md#GetAttachment) | **Get** /v1/attachments/{id} | Get a single attachment.
[**ListAttachment**](AttachmentsApi.md#ListAttachment) | **Get** /v1/attachments | List all attachments.
[**StoreAttachment**](AttachmentsApi.md#StoreAttachment) | **Post** /v1/attachments | Store a new attachment.
[**UpdateAttachment**](AttachmentsApi.md#UpdateAttachment) | **Put** /v1/attachments/{id} | Update existing attachment.
[**UploadAttachment**](AttachmentsApi.md#UploadAttachment) | **Post** /v1/attachments/{id}/upload | Upload an attachment.



## DeleteAttachment

> DeleteAttachment(ctx, id).XTraceId(xTraceId).Execute()

Delete an attachment.



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
    id := "123" // string | The ID of the single attachment.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AttachmentsApi.DeleteAttachment(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.DeleteAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the single attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentRequest struct via the builder pattern


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


## DownloadAttachment

> *os.File DownloadAttachment(ctx, id).XTraceId(xTraceId).Execute()

Download a single attachment.



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
    id := "123" // string | The ID of the attachment.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentsApi.DownloadAttachment(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.DownloadAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadAttachment`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AttachmentsApi.DownloadAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachment

> AttachmentSingle GetAttachment(ctx, id).XTraceId(xTraceId).Execute()

Get a single attachment.



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
    id := "123" // string | The ID of the attachment.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentsApi.GetAttachment(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.GetAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttachment`: AttachmentSingle
    fmt.Fprintf(os.Stdout, "Response from `AttachmentsApi.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**AttachmentSingle**](AttachmentSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttachment

> AttachmentArray ListAttachment(ctx).XTraceId(xTraceId).Page(page).Execute()

List all attachments.



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
    resp, r, err := apiClient.AttachmentsApi.ListAttachment(context.Background()).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.ListAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttachment`: AttachmentArray
    fmt.Fprintf(os.Stdout, "Response from `AttachmentsApi.ListAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAttachmentRequest struct via the builder pattern


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
- **Accept**: application/json, application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreAttachment

> AttachmentSingle StoreAttachment(ctx).AttachmentStore(attachmentStore).XTraceId(xTraceId).Execute()

Store a new attachment.



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
    attachmentStore := *openapiclient.NewAttachmentStore("file.pdf", openapiclient.AttachableType("Account"), "134") // AttachmentStore | JSON array or key=value pairs with the necessary attachment information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentsApi.StoreAttachment(context.Background()).AttachmentStore(attachmentStore).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.StoreAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreAttachment`: AttachmentSingle
    fmt.Fprintf(os.Stdout, "Response from `AttachmentsApi.StoreAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachmentStore** | [**AttachmentStore**](AttachmentStore.md) | JSON array or key&#x3D;value pairs with the necessary attachment information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**AttachmentSingle**](AttachmentSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json, application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttachment

> AttachmentSingle UpdateAttachment(ctx, id).AttachmentUpdate(attachmentUpdate).XTraceId(xTraceId).Execute()

Update existing attachment.



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
    id := "123" // string | The ID of the attachment.
    attachmentUpdate := *openapiclient.NewAttachmentUpdate() // AttachmentUpdate | JSON array with updated attachment information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentsApi.UpdateAttachment(context.Background(), id).AttachmentUpdate(attachmentUpdate).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.UpdateAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttachment`: AttachmentSingle
    fmt.Fprintf(os.Stdout, "Response from `AttachmentsApi.UpdateAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachmentUpdate** | [**AttachmentUpdate**](AttachmentUpdate.md) | JSON array with updated attachment information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**AttachmentSingle**](AttachmentSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json, application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadAttachment

> UploadAttachment(ctx, id).XTraceId(xTraceId).Body(body).Execute()

Upload an attachment.



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
    id := "123" // string | The ID of the attachment.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    body := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AttachmentsApi.UploadAttachment(context.Background(), id).XTraceId(xTraceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.UploadAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

