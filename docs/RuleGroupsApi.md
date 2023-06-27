# \RuleGroupsApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRuleGroup**](RuleGroupsApi.md#DeleteRuleGroup) | **Delete** /v1/rule-groups/{id} | Delete a rule group.
[**FireRuleGroup**](RuleGroupsApi.md#FireRuleGroup) | **Post** /v1/rule-groups/{id}/trigger | Fire the rule group on your transactions.
[**GetRuleGroup**](RuleGroupsApi.md#GetRuleGroup) | **Get** /v1/rule-groups/{id} | Get a single rule group.
[**ListRuleByGroup**](RuleGroupsApi.md#ListRuleByGroup) | **Get** /v1/rule-groups/{id}/rules | List rules in this rule group.
[**ListRuleGroup**](RuleGroupsApi.md#ListRuleGroup) | **Get** /v1/rule-groups | List all rule groups.
[**StoreRuleGroup**](RuleGroupsApi.md#StoreRuleGroup) | **Post** /v1/rule-groups | Store a new rule group.
[**TestRuleGroup**](RuleGroupsApi.md#TestRuleGroup) | **Get** /v1/rule-groups/{id}/test | Test which transactions would be hit by the rule group. No changes will be made.
[**UpdateRuleGroup**](RuleGroupsApi.md#UpdateRuleGroup) | **Put** /v1/rule-groups/{id} | Update existing rule group.



## DeleteRuleGroup

> DeleteRuleGroup(ctx, id).XTraceId(xTraceId).Execute()

Delete a rule group.



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
    id := "123" // string | The ID of the rule group.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RuleGroupsApi.DeleteRuleGroup(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.DeleteRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleGroupRequest struct via the builder pattern


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


## FireRuleGroup

> FireRuleGroup(ctx, id).XTraceId(xTraceId).Start(start).End(end).Accounts(accounts).Execute()

Fire the rule group on your transactions.



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
    id := "123" // string | The ID of the rule group.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. Both the start date and the end date must be present.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. Both the start date and the end date must be present.  (optional)
    accounts := []int64{int64(123)} // []int64 | Limit the triggering of the rule group to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RuleGroupsApi.FireRuleGroup(context.Background(), id).XTraceId(xTraceId).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.FireRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFireRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **start** | **string** | A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. Both the start date and the end date must be present.  | 
 **end** | **string** | A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. Both the start date and the end date must be present.  | 
 **accounts** | **[]int64** | Limit the triggering of the rule group to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  | 

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


## GetRuleGroup

> RuleGroupSingle GetRuleGroup(ctx, id).XTraceId(xTraceId).Execute()

Get a single rule group.



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
    id := "123" // string | The ID of the rule group.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.GetRuleGroup(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.GetRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleGroup`: RuleGroupSingle
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.GetRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**RuleGroupSingle**](RuleGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRuleByGroup

> RuleArray ListRuleByGroup(ctx, id).XTraceId(xTraceId).Page(page).Execute()

List rules in this rule group.



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
    id := "123" // string | The ID of the rule group.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.ListRuleByGroup(context.Background(), id).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.ListRuleByGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRuleByGroup`: RuleArray
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.ListRuleByGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRuleByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**RuleArray**](RuleArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRuleGroup

> RuleGroupArray ListRuleGroup(ctx).XTraceId(xTraceId).Page(page).Execute()

List all rule groups.



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
    page := int32(1) // int32 | Page number. The default pagination is 50 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.ListRuleGroup(context.Background()).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.ListRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRuleGroup`: RuleGroupArray
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.ListRuleGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50 | 

### Return type

[**RuleGroupArray**](RuleGroupArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreRuleGroup

> RuleGroupSingle StoreRuleGroup(ctx).RuleGroupStore(ruleGroupStore).XTraceId(xTraceId).Execute()

Store a new rule group.



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
    ruleGroupStore := *openapiclient.NewRuleGroupStore("Default rule group") // RuleGroupStore | JSON array or key=value pairs with the necessary rule group information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.StoreRuleGroup(context.Background()).RuleGroupStore(ruleGroupStore).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.StoreRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreRuleGroup`: RuleGroupSingle
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.StoreRuleGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleGroupStore** | [**RuleGroupStore**](RuleGroupStore.md) | JSON array or key&#x3D;value pairs with the necessary rule group information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**RuleGroupSingle**](RuleGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRuleGroup

> TransactionArray TestRuleGroup(ctx, id).XTraceId(xTraceId).Page(page).Start(start).End(end).SearchLimit(searchLimit).TriggeredLimit(triggeredLimit).Accounts(accounts).Execute()

Test which transactions would be hit by the rule group. No changes will be made.



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
    id := "123" // string | The ID of the rule group.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    page := int32(1) // int32 | Page number. The default pagination is 50 items. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  (optional)
    searchLimit := int32(56) // int32 | Maximum number of transactions Firefly III will try. Don't set this too high, or it will take Firefly III very long to run the test. I suggest a max of 200.  (optional)
    triggeredLimit := int32(56) // int32 | Maximum number of transactions the rule group can actually trigger on, before Firefly III stops. I would suggest setting this to 10 or 15. Don't go above the user's page size, because browsing to page 2 or 3 of a test result would fire the test again, making any navigation efforts very slow.  (optional)
    accounts := []int64{int64(123)} // []int64 | Limit the testing of the rule group to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.TestRuleGroup(context.Background(), id).XTraceId(xTraceId).Page(page).Start(start).End(end).SearchLimit(searchLimit).TriggeredLimit(triggeredLimit).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.TestRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestRuleGroup`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.TestRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **page** | **int32** | Page number. The default pagination is 50 items. | 
 **start** | **string** | A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  | 
 **end** | **string** | A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  | 
 **searchLimit** | **int32** | Maximum number of transactions Firefly III will try. Don&#39;t set this too high, or it will take Firefly III very long to run the test. I suggest a max of 200.  | 
 **triggeredLimit** | **int32** | Maximum number of transactions the rule group can actually trigger on, before Firefly III stops. I would suggest setting this to 10 or 15. Don&#39;t go above the user&#39;s page size, because browsing to page 2 or 3 of a test result would fire the test again, making any navigation efforts very slow.  | 
 **accounts** | **[]int64** | Limit the testing of the rule group to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  | 

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


## UpdateRuleGroup

> RuleGroupSingle UpdateRuleGroup(ctx, id).RuleGroupUpdate(ruleGroupUpdate).XTraceId(xTraceId).Execute()

Update existing rule group.



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
    id := "123" // string | The ID of the rule group.
    ruleGroupUpdate := *openapiclient.NewRuleGroupUpdate() // RuleGroupUpdate | JSON array with updated rule group information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.UpdateRuleGroup(context.Background(), id).RuleGroupUpdate(ruleGroupUpdate).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.UpdateRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRuleGroup`: RuleGroupSingle
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.UpdateRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleGroupUpdate** | [**RuleGroupUpdate**](RuleGroupUpdate.md) | JSON array with updated rule group information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**RuleGroupSingle**](RuleGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

