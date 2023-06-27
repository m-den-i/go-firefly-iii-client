# \ObjectGroupsApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObjectGroup**](ObjectGroupsApi.md#DeleteObjectGroup) | **Delete** /v1/object-groups/{id} | Delete a object group.
[**GetObjectGroup**](ObjectGroupsApi.md#GetObjectGroup) | **Get** /v1/object-groups/{id} | Get a single object group.
[**ListBillByObjectGroup**](ObjectGroupsApi.md#ListBillByObjectGroup) | **Get** /v1/object-groups/{id}/bills | List all bills with this object group.
[**ListObjectGroups**](ObjectGroupsApi.md#ListObjectGroups) | **Get** /v1/object-groups | List all oject groups.
[**ListPiggyBankByObjectGroup**](ObjectGroupsApi.md#ListPiggyBankByObjectGroup) | **Get** /v1/object-groups/{id}/piggy-banks | List all piggy banks related to the object group.
[**UpdateObjectGroup**](ObjectGroupsApi.md#UpdateObjectGroup) | **Put** /v1/object-groups/{id} | Update existing object group.



## DeleteObjectGroup

> DeleteObjectGroup(ctx, id).XTraceId(xTraceId).Execute()

Delete a object group.



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
    id := "123" // string | The ID of the object group.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ObjectGroupsApi.DeleteObjectGroup(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.DeleteObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the object group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectGroupRequest struct via the builder pattern


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


## GetObjectGroup

> ObjectGroupSingle GetObjectGroup(ctx, id).XTraceId(xTraceId).Execute()

Get a single object group.



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
    id := "123" // string | The ID of the object group.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectGroupsApi.GetObjectGroup(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.GetObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectGroup`: ObjectGroupSingle
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.GetObjectGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the object group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**ObjectGroupSingle**](ObjectGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillByObjectGroup

> BillArray ListBillByObjectGroup(ctx, id).XTraceId(xTraceId).Page(page).Execute()

List all bills with this object group.



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
    id := "123" // string | The ID of the account.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(56) // int32 | Page number. The default pagination is per 50 items. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectGroupsApi.ListBillByObjectGroup(context.Background(), id).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.ListBillByObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillByObjectGroup`: BillArray
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.ListBillByObjectGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBillByObjectGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is per 50 items. | 

### Return type

[**BillArray**](BillArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjectGroups

> ObjectGroupArray ListObjectGroups(ctx).XTraceId(xTraceId).Page(page).Execute()

List all oject groups.



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
    resp, r, err := apiClient.ObjectGroupsApi.ListObjectGroups(context.Background()).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.ListObjectGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObjectGroups`: ObjectGroupArray
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.ListObjectGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListObjectGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**ObjectGroupArray**](ObjectGroupArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPiggyBankByObjectGroup

> PiggyBankArray ListPiggyBankByObjectGroup(ctx, id).XTraceId(xTraceId).Page(page).Execute()

List all piggy banks related to the object group.



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
    id := "123" // string | The ID of the account.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(56) // int32 | Page number. The default pagination is per 50 items. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectGroupsApi.ListPiggyBankByObjectGroup(context.Background(), id).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.ListPiggyBankByObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPiggyBankByObjectGroup`: PiggyBankArray
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.ListPiggyBankByObjectGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPiggyBankByObjectGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is per 50 items. | 

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


## UpdateObjectGroup

> ObjectGroupSingle UpdateObjectGroup(ctx, id).ObjectGroupUpdate(objectGroupUpdate).XTraceId(xTraceId).Execute()

Update existing object group.



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
    id := "123" // string | The ID of the object group
    objectGroupUpdate := *openapiclient.NewObjectGroupUpdate() // ObjectGroupUpdate | JSON array with updated piggy bank information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectGroupsApi.UpdateObjectGroup(context.Background(), id).ObjectGroupUpdate(objectGroupUpdate).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.UpdateObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObjectGroup`: ObjectGroupSingle
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.UpdateObjectGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the object group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectGroupUpdate** | [**ObjectGroupUpdate**](ObjectGroupUpdate.md) | JSON array with updated piggy bank information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**ObjectGroupSingle**](ObjectGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

