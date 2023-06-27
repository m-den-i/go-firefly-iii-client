# \SummaryApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBasicSummary**](SummaryApi.md#GetBasicSummary) | **Get** /v1/summary/basic | Returns basic sums of the users data.



## GetBasicSummary

> []BasicSummaryEntry GetBasicSummary(ctx).Start(start).End(end).XTraceId(xTraceId).CurrencyCode(currencyCode).Execute()

Returns basic sums of the users data.



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
    currencyCode := "currencyCode_example" // string | A currency code like EUR or USD, to filter the result.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SummaryApi.GetBasicSummary(context.Background()).Start(start).End(end).XTraceId(xTraceId).CurrencyCode(currencyCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SummaryApi.GetBasicSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicSummary`: []BasicSummaryEntry
    fmt.Fprintf(os.Stdout, "Response from `SummaryApi.GetBasicSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **currencyCode** | **string** | A currency code like EUR or USD, to filter the result.  | 

### Return type

[**[]BasicSummaryEntry**](BasicSummaryEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

