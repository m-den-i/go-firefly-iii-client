# \DataApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateTransactions**](DataApi.md#BulkUpdateTransactions) | **Post** /v1/data/bulk/transactions | Bulk update transaction properties. For more information, see https://docs.firefly-iii.org/firefly-iii/api/specials
[**DestroyData**](DataApi.md#DestroyData) | **Delete** /v1/data/destroy | Endpoint to destroy user data
[**ExportAccounts**](DataApi.md#ExportAccounts) | **Get** /v1/data/export/accounts | Export account data from Firefly III
[**ExportBills**](DataApi.md#ExportBills) | **Get** /v1/data/export/bills | Export bills from Firefly III
[**ExportBudgets**](DataApi.md#ExportBudgets) | **Get** /v1/data/export/budgets | Export budgets and budget amount data from Firefly III
[**ExportCategories**](DataApi.md#ExportCategories) | **Get** /v1/data/export/categories | Export category data from Firefly III
[**ExportPiggies**](DataApi.md#ExportPiggies) | **Get** /v1/data/export/piggy-banks | Export piggy banks from Firefly III
[**ExportRecurring**](DataApi.md#ExportRecurring) | **Get** /v1/data/export/recurring | Export recurring transaction data from Firefly III
[**ExportRules**](DataApi.md#ExportRules) | **Get** /v1/data/export/rules | Export rule groups and rule data from Firefly III
[**ExportTags**](DataApi.md#ExportTags) | **Get** /v1/data/export/tags | Export tag data from Firefly III
[**ExportTransactions**](DataApi.md#ExportTransactions) | **Get** /v1/data/export/transactions | Export transaction data from Firefly III
[**PurgeData**](DataApi.md#PurgeData) | **Delete** /v1/data/purge | Endpoint to purge user data



## BulkUpdateTransactions

> BulkUpdateTransactions(ctx).Query(query).Execute()

Bulk update transaction properties. For more information, see https://docs.firefly-iii.org/firefly-iii/api/specials



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
    query := "query_example" // string | The JSON query.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataApi.BulkUpdateTransactions(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.BulkUpdateTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The JSON query. | 

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


## DestroyData

> DestroyData(ctx).Objects(objects).XTraceId(xTraceId).Execute()

Endpoint to destroy user data



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
    objects := openapiclient.DataDestroyObject("not_assets_liabilities") // DataDestroyObject | The type of data that you wish to destroy. You can only use one at a time.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataApi.DestroyData(context.Background()).Objects(objects).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.DestroyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objects** | [**DataDestroyObject**](DataDestroyObject.md) | The type of data that you wish to destroy. You can only use one at a time. | 
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


## ExportAccounts

> *os.File ExportAccounts(ctx).XTraceId(xTraceId).Type_(type_).Execute()

Export account data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataApi.ExportAccounts(context.Background()).XTraceId(xTraceId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAccounts`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportBills

> *os.File ExportBills(ctx).XTraceId(xTraceId).Type_(type_).Execute()

Export bills from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataApi.ExportBills(context.Background()).XTraceId(xTraceId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportBills``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportBills`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportBills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportBillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportBudgets

> *os.File ExportBudgets(ctx).XTraceId(xTraceId).Type_(type_).Execute()

Export budgets and budget amount data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataApi.ExportBudgets(context.Background()).XTraceId(xTraceId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportBudgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportBudgets`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCategories

> *os.File ExportCategories(ctx).XTraceId(xTraceId).Type_(type_).Execute()

Export category data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataApi.ExportCategories(context.Background()).XTraceId(xTraceId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportCategories`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPiggies

> *os.File ExportPiggies(ctx).XTraceId(xTraceId).Type_(type_).Execute()

Export piggy banks from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataApi.ExportPiggies(context.Background()).XTraceId(xTraceId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportPiggies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPiggies`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportPiggies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportPiggiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRecurring

> *os.File ExportRecurring(ctx).XTraceId(xTraceId).Type_(type_).Execute()

Export recurring transaction data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataApi.ExportRecurring(context.Background()).XTraceId(xTraceId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportRecurring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRecurring`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportRecurring`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRecurringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRules

> *os.File ExportRules(ctx).XTraceId(xTraceId).Type_(type_).Execute()

Export rule groups and rule data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataApi.ExportRules(context.Background()).XTraceId(xTraceId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRules`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTags

> *os.File ExportTags(ctx).XTraceId(xTraceId).Type_(type_).Execute()

Export tag data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataApi.ExportTags(context.Background()).XTraceId(xTraceId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTags`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTransactions

> *os.File ExportTransactions(ctx).Start(start).End(end).XTraceId(xTraceId).Accounts(accounts).Type_(type_).Execute()

Export transaction data from Firefly III



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
    accounts := "1,2,3" // string | Limit the export of transactions to these accounts only. Only asset accounts will be accepted. Other types will be silently dropped.  (optional)
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataApi.ExportTransactions(context.Background()).Start(start).End(end).XTraceId(xTraceId).Accounts(accounts).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTransactions`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **accounts** | **string** | Limit the export of transactions to these accounts only. Only asset accounts will be accepted. Other types will be silently dropped.  | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeData

> PurgeData(ctx).XTraceId(xTraceId).Execute()

Endpoint to purge user data



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataApi.PurgeData(context.Background()).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.PurgeData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurgeDataRequest struct via the builder pattern


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

