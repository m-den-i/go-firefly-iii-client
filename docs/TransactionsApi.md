# \TransactionsApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTransaction**](TransactionsApi.md#DeleteTransaction) | **Delete** /v1/transactions/{id} | Delete a transaction.
[**DeleteTransactionJournal**](TransactionsApi.md#DeleteTransactionJournal) | **Delete** /v1/transaction-journals/{id} | Delete split from transaction
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /v1/transactions/{id} | Get a single transaction.
[**GetTransactionByJournal**](TransactionsApi.md#GetTransactionByJournal) | **Get** /v1/transaction-journals/{id} | Get a single transaction, based on one of the underlying transaction journals (transaction splits).
[**ListAttachmentByTransaction**](TransactionsApi.md#ListAttachmentByTransaction) | **Get** /v1/transactions/{id}/attachments | Lists all attachments.
[**ListEventByTransaction**](TransactionsApi.md#ListEventByTransaction) | **Get** /v1/transactions/{id}/piggy-bank-events | Lists all piggy bank events.
[**ListLinksByJournal**](TransactionsApi.md#ListLinksByJournal) | **Get** /v1/transaction-journals/{id}/links | Lists all the transaction links for an individual journal (individual split).
[**ListTransaction**](TransactionsApi.md#ListTransaction) | **Get** /v1/transactions | List all the user&#39;s transactions. 
[**StoreTransaction**](TransactionsApi.md#StoreTransaction) | **Post** /v1/transactions | Store a new transaction
[**UpdateTransaction**](TransactionsApi.md#UpdateTransaction) | **Put** /v1/transactions/{id} | Update existing transaction. For more information, see https://docs.firefly-iii.org/firefly-iii/api/specials



## DeleteTransaction

> DeleteTransaction(ctx, id).XTraceId(xTraceId).Execute()

Delete a transaction.



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
    id := "123" // string | The ID of the transaction.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransactionsApi.DeleteTransaction(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.DeleteTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransactionRequest struct via the builder pattern


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


## DeleteTransactionJournal

> DeleteTransactionJournal(ctx, id).XTraceId(xTraceId).Execute()

Delete split from transaction



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
    id := "123" // string | The ID of the transaction journal (the split) you wish to delete.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransactionsApi.DeleteTransactionJournal(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.DeleteTransactionJournal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the transaction journal (the split) you wish to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransactionJournalRequest struct via the builder pattern


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


## GetTransaction

> TransactionSingle GetTransaction(ctx, id).XTraceId(xTraceId).Execute()

Get a single transaction.



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
    id := "123" // string | The ID of the transaction.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransaction(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: TransactionSingle
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**TransactionSingle**](TransactionSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionByJournal

> TransactionSingle GetTransactionByJournal(ctx, id).XTraceId(xTraceId).Execute()

Get a single transaction, based on one of the underlying transaction journals (transaction splits).



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
    id := "123" // string | The ID of the transaction journal (split).
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransactionByJournal(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactionByJournal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionByJournal`: TransactionSingle
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactionByJournal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the transaction journal (split). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByJournalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**TransactionSingle**](TransactionSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttachmentByTransaction

> AttachmentArray ListAttachmentByTransaction(ctx, id).XTraceId(xTraceId).Page(page).Execute()

Lists all attachments.



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
    id := "123" // string | The ID of the transaction.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.ListAttachmentByTransaction(context.Background(), id).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListAttachmentByTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttachmentByTransaction`: AttachmentArray
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListAttachmentByTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAttachmentByTransactionRequest struct via the builder pattern


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


## ListEventByTransaction

> PiggyBankEventArray ListEventByTransaction(ctx, id).XTraceId(xTraceId).Page(page).Execute()

Lists all piggy bank events.



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
    id := "123" // string | The ID of the transaction.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.ListEventByTransaction(context.Background(), id).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListEventByTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventByTransaction`: PiggyBankEventArray
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListEventByTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEventByTransactionRequest struct via the builder pattern


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


## ListLinksByJournal

> TransactionLinkArray ListLinksByJournal(ctx, id).XTraceId(xTraceId).Page(page).Execute()

Lists all the transaction links for an individual journal (individual split).



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
    id := "123" // string | The ID of the transaction journal / the split.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.ListLinksByJournal(context.Background(), id).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListLinksByJournal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLinksByJournal`: TransactionLinkArray
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListLinksByJournal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the transaction journal / the split. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLinksByJournalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**TransactionLinkArray**](TransactionLinkArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransaction

> TransactionArray ListTransaction(ctx).XTraceId(xTraceId).Page(page).Start(start).End(end).Type_(type_).Execute()

List all the user's transactions. 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/m-den-i/go-firefly-iii-client"
)

func main() {
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD. This is the start date of the selected range (inclusive).  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD. This is the end date of the selected range (inclusive).  (optional)
    type_ := openapiclient.TransactionTypeFilter("all") // TransactionTypeFilter | Optional filter on the transaction type(s) returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.ListTransaction(context.Background()).XTraceId(xTraceId).Page(page).Start(start).End(end).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransaction`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 
 **start** | **string** | A date formatted YYYY-MM-DD. This is the start date of the selected range (inclusive).  | 
 **end** | **string** | A date formatted YYYY-MM-DD. This is the end date of the selected range (inclusive).  | 
 **type_** | [**TransactionTypeFilter**](TransactionTypeFilter.md) | Optional filter on the transaction type(s) returned. | 

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


## StoreTransaction

> TransactionSingle StoreTransaction(ctx).TransactionStore(transactionStore).XTraceId(xTraceId).Execute()

Store a new transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/m-den-i/go-firefly-iii-client"
)

func main() {
    transactionStore := *openapiclient.NewTransactionStore([]openapiclient.TransactionSplitStore{*openapiclient.NewTransactionSplitStore(openapiclient.TransactionTypeProperty("withdrawal"), time.Now(), "123.45", "Vegetables")}) // TransactionStore | JSON array or key=value pairs with the necessary transaction information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.StoreTransaction(context.Background()).TransactionStore(transactionStore).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.StoreTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreTransaction`: TransactionSingle
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.StoreTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionStore** | [**TransactionStore**](TransactionStore.md) | JSON array or key&#x3D;value pairs with the necessary transaction information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**TransactionSingle**](TransactionSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransaction

> TransactionSingle UpdateTransaction(ctx, id).TransactionUpdate(transactionUpdate).XTraceId(xTraceId).Execute()

Update existing transaction. For more information, see https://docs.firefly-iii.org/firefly-iii/api/specials



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
    id := "123" // string | The ID of the transaction.
    transactionUpdate := *openapiclient.NewTransactionUpdate() // TransactionUpdate | JSON array with updated transaction information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.UpdateTransaction(context.Background(), id).TransactionUpdate(transactionUpdate).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.UpdateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransaction`: TransactionSingle
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.UpdateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionUpdate** | [**TransactionUpdate**](TransactionUpdate.md) | JSON array with updated transaction information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**TransactionSingle**](TransactionSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

