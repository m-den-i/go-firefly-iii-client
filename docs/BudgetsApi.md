# \BudgetsApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBudget**](BudgetsApi.md#DeleteBudget) | **Delete** /v1/budgets/{id} | Delete a budget.
[**DeleteBudgetLimit**](BudgetsApi.md#DeleteBudgetLimit) | **Delete** /v1/budgets/{id}/limits/{limitId} | Delete a budget limit.
[**GetBudget**](BudgetsApi.md#GetBudget) | **Get** /v1/budgets/{id} | Get a single budget.
[**GetBudgetLimit**](BudgetsApi.md#GetBudgetLimit) | **Get** /v1/budgets/{id}/limits/{limitId} | Get single budget limit.
[**ListAttachmentByBudget**](BudgetsApi.md#ListAttachmentByBudget) | **Get** /v1/budgets/{id}/attachments | Lists all attachments of a budget.
[**ListBudget**](BudgetsApi.md#ListBudget) | **Get** /v1/budgets | List all budgets.
[**ListBudgetLimit**](BudgetsApi.md#ListBudgetLimit) | **Get** /v1/budget-limits | Get list of budget limits by date
[**ListBudgetLimitByBudget**](BudgetsApi.md#ListBudgetLimitByBudget) | **Get** /v1/budgets/{id}/limits | Get all limits for a budget.
[**ListTransactionByBudget**](BudgetsApi.md#ListTransactionByBudget) | **Get** /v1/budgets/{id}/transactions | All transactions to a budget.
[**ListTransactionByBudgetLimit**](BudgetsApi.md#ListTransactionByBudgetLimit) | **Get** /v1/budgets/{id}/limits/{limitId}/transactions | List all transactions by a budget limit ID.
[**StoreBudget**](BudgetsApi.md#StoreBudget) | **Post** /v1/budgets | Store a new budget
[**StoreBudgetLimit**](BudgetsApi.md#StoreBudgetLimit) | **Post** /v1/budgets/{id}/limits | Store new budget limit.
[**UpdateBudget**](BudgetsApi.md#UpdateBudget) | **Put** /v1/budgets/{id} | Update existing budget.
[**UpdateBudgetLimit**](BudgetsApi.md#UpdateBudgetLimit) | **Put** /v1/budgets/{id}/limits/{limitId} | Update existing budget limit.



## DeleteBudget

> DeleteBudget(ctx, id).XTraceId(xTraceId).Execute()

Delete a budget.



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
    id := "123" // string | The ID of the budget.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BudgetsApi.DeleteBudget(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.DeleteBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBudgetRequest struct via the builder pattern


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


## DeleteBudgetLimit

> DeleteBudgetLimit(ctx, id, limitId).XTraceId(xTraceId).Execute()

Delete a budget limit.



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
    id := "123" // string | The ID of the budget. The budget limit MUST be associated to the budget ID.
    limitId := "123" // string | The ID of the budget limit. The budget limit MUST be associated to the budget ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BudgetsApi.DeleteBudgetLimit(context.Background(), id, limitId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.DeleteBudgetLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the budget. The budget limit MUST be associated to the budget ID. | 
**limitId** | **string** | The ID of the budget limit. The budget limit MUST be associated to the budget ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBudgetLimitRequest struct via the builder pattern


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


## GetBudget

> BudgetSingle GetBudget(ctx, id).XTraceId(xTraceId).Start(start).End(end).Execute()

Get a single budget.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the requested budget.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD, to get info on how much the user has spent.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD, to get info on how much the user has spent.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.GetBudget(context.Background(), id).XTraceId(xTraceId).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.GetBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBudget`: BudgetSingle
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.GetBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the requested budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **start** | **string** | A date formatted YYYY-MM-DD, to get info on how much the user has spent.  | 
 **end** | **string** | A date formatted YYYY-MM-DD, to get info on how much the user has spent.  | 

### Return type

[**BudgetSingle**](BudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetLimit

> BudgetLimitSingle GetBudgetLimit(ctx, id, limitId).XTraceId(xTraceId).Execute()

Get single budget limit.

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
    id := "123" // string | The ID of the budget. The budget limit MUST be associated to the budget ID.
    limitId := int32(1) // int32 | The ID of the budget limit. The budget limit MUST be associated to the budget ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.GetBudgetLimit(context.Background(), id, limitId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.GetBudgetLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBudgetLimit`: BudgetLimitSingle
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.GetBudgetLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the budget. The budget limit MUST be associated to the budget ID. | 
**limitId** | **int32** | The ID of the budget limit. The budget limit MUST be associated to the budget ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**BudgetLimitSingle**](BudgetLimitSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttachmentByBudget

> AttachmentArray ListAttachmentByBudget(ctx, id).XTraceId(xTraceId).Page(page).Execute()

Lists all attachments of a budget.



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
    id := "123" // string | The ID of the budget.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.ListAttachmentByBudget(context.Background(), id).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.ListAttachmentByBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttachmentByBudget`: AttachmentArray
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.ListAttachmentByBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAttachmentByBudgetRequest struct via the builder pattern


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


## ListBudget

> BudgetArray ListBudget(ctx).XTraceId(xTraceId).Page(page).Start(start).End(end).Execute()

List all budgets.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD, to get info on how much the user has spent. You must submit both start and end.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD, to get info on how much the user has spent. You must submit both start and end.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.ListBudget(context.Background()).XTraceId(xTraceId).Page(page).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.ListBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBudget`: BudgetArray
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.ListBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 
 **start** | **string** | A date formatted YYYY-MM-DD, to get info on how much the user has spent. You must submit both start and end.  | 
 **end** | **string** | A date formatted YYYY-MM-DD, to get info on how much the user has spent. You must submit both start and end.  | 

### Return type

[**BudgetArray**](BudgetArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBudgetLimit

> BudgetLimitArray ListBudgetLimit(ctx).Start(start).End(end).XTraceId(xTraceId).Execute()

Get list of budget limits by date



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.ListBudgetLimit(context.Background()).Start(start).End(end).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.ListBudgetLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBudgetLimit`: BudgetLimitArray
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.ListBudgetLimit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBudgetLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**BudgetLimitArray**](BudgetLimitArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBudgetLimitByBudget

> BudgetLimitArray ListBudgetLimitByBudget(ctx, id).XTraceId(xTraceId).Start(start).End(end).Execute()

Get all limits for a budget.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the requested budget.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.ListBudgetLimitByBudget(context.Background(), id).XTraceId(xTraceId).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.ListBudgetLimitByBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBudgetLimitByBudget`: BudgetLimitArray
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.ListBudgetLimitByBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the requested budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBudgetLimitByBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 

### Return type

[**BudgetLimitArray**](BudgetLimitArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionByBudget

> TransactionArray ListTransactionByBudget(ctx, id).XTraceId(xTraceId).Limit(limit).Page(page).Start(start).End(end).Type_(type_).Execute()

All transactions to a budget.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the budget.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    limit := int32(5) // int32 | Limits the number of results on one page. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD.  (optional)
    type_ := openapiclient.TransactionTypeFilter("all") // TransactionTypeFilter | Optional filter on the transaction type(s) returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.ListTransactionByBudget(context.Background(), id).XTraceId(xTraceId).Limit(limit).Page(page).Start(start).End(end).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.ListTransactionByBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionByBudget`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.ListTransactionByBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionByBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **limit** | **int32** | Limits the number of results on one page. | 
 **page** | **int32** | Page number. The default pagination is 50. | 
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **type_** | [**TransactionTypeFilter**](TransactionTypeFilter.md) | Optional filter on the transaction type(s) returned | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionByBudgetLimit

> TransactionArray ListTransactionByBudgetLimit(ctx, id, limitId).XTraceId(xTraceId).Page(page).Type_(type_).Execute()

List all transactions by a budget limit ID.



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
    id := "123" // string | The ID of the budget. The budget limit MUST be associated to the budget ID.
    limitId := "123" // string | The ID of the budget limit. The budget limit MUST be associated to the budget ID.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    type_ := openapiclient.TransactionTypeFilter("all") // TransactionTypeFilter | Optional filter on the transaction type(s) returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.ListTransactionByBudgetLimit(context.Background(), id, limitId).XTraceId(xTraceId).Page(page).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.ListTransactionByBudgetLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionByBudgetLimit`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.ListTransactionByBudgetLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the budget. The budget limit MUST be associated to the budget ID. | 
**limitId** | **string** | The ID of the budget limit. The budget limit MUST be associated to the budget ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionByBudgetLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 
 **type_** | [**TransactionTypeFilter**](TransactionTypeFilter.md) | Optional filter on the transaction type(s) returned | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreBudget

> BudgetSingle StoreBudget(ctx).BudgetStore(budgetStore).XTraceId(xTraceId).Execute()

Store a new budget



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
    budgetStore := *openapiclient.NewBudgetStore("Bills") // BudgetStore | JSON array or key=value pairs with the necessary budget information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.StoreBudget(context.Background()).BudgetStore(budgetStore).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.StoreBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreBudget`: BudgetSingle
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.StoreBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **budgetStore** | [**BudgetStore**](BudgetStore.md) | JSON array or key&#x3D;value pairs with the necessary budget information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**BudgetSingle**](BudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreBudgetLimit

> BudgetLimitSingle StoreBudgetLimit(ctx, id).BudgetLimitStore(budgetLimitStore).XTraceId(xTraceId).Execute()

Store new budget limit.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the budget.
    budgetLimitStore := *openapiclient.NewBudgetLimitStore("23", time.Now(), time.Now(), "123.45") // BudgetLimitStore | JSON array or key=value pairs with the necessary budget information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.StoreBudgetLimit(context.Background(), id).BudgetLimitStore(budgetLimitStore).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.StoreBudgetLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreBudgetLimit`: BudgetLimitSingle
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.StoreBudgetLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreBudgetLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **budgetLimitStore** | [**BudgetLimitStore**](BudgetLimitStore.md) | JSON array or key&#x3D;value pairs with the necessary budget information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**BudgetLimitSingle**](BudgetLimitSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBudget

> BudgetSingle UpdateBudget(ctx, id).BudgetUpdate(budgetUpdate).XTraceId(xTraceId).Execute()

Update existing budget.



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
    id := "123" // string | The ID of the budget.
    budgetUpdate := *openapiclient.NewBudgetUpdate() // BudgetUpdate | JSON array with updated budget information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.UpdateBudget(context.Background(), id).BudgetUpdate(budgetUpdate).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.UpdateBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBudget`: BudgetSingle
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.UpdateBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **budgetUpdate** | [**BudgetUpdate**](BudgetUpdate.md) | JSON array with updated budget information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**BudgetSingle**](BudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBudgetLimit

> BudgetLimitSingle UpdateBudgetLimit(ctx, id, limitId).BudgetLimit(budgetLimit).XTraceId(xTraceId).Execute()

Update existing budget limit.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "123" // string | The ID of the budget. The budget limit MUST be associated to the budget ID.
    limitId := "123" // string | The ID of the budget limit. The budget limit MUST be associated to the budget ID.
    budgetLimit := *openapiclient.NewBudgetLimit(time.Now(), time.Now(), "23", "123.45") // BudgetLimit | JSON array with updated budget limit information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BudgetsApi.UpdateBudgetLimit(context.Background(), id, limitId).BudgetLimit(budgetLimit).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BudgetsApi.UpdateBudgetLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBudgetLimit`: BudgetLimitSingle
    fmt.Fprintf(os.Stdout, "Response from `BudgetsApi.UpdateBudgetLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the budget. The budget limit MUST be associated to the budget ID. | 
**limitId** | **string** | The ID of the budget limit. The budget limit MUST be associated to the budget ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBudgetLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **budgetLimit** | [**BudgetLimit**](BudgetLimit.md) | JSON array with updated budget limit information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**BudgetLimitSingle**](BudgetLimitSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

