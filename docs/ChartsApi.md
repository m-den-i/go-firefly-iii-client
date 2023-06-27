# \ChartsApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChartAccountOverview**](ChartsApi.md#GetChartAccountOverview) | **Get** /v1/chart/account/overview | Dashboard chart with asset account balance information.



## GetChartAccountOverview

> []ChartDataSet GetChartAccountOverview(ctx).Start(start).End(end).XTraceId(xTraceId).Execute()

Dashboard chart with asset account balance information.



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
    resp, r, err := apiClient.ChartsApi.GetChartAccountOverview(context.Background()).Start(start).End(end).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChartsApi.GetChartAccountOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChartAccountOverview`: []ChartDataSet
    fmt.Fprintf(os.Stdout, "Response from `ChartsApi.GetChartAccountOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChartAccountOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**[]ChartDataSet**](ChartDataSet.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

