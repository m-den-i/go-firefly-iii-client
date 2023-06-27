# \RulesApi

All URIs are relative to *https://demo.firefly-iii.org/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRule**](RulesApi.md#DeleteRule) | **Delete** /v1/rules/{id} | Delete an rule.
[**FireRule**](RulesApi.md#FireRule) | **Post** /v1/rules/{id}/trigger | Fire the rule on your transactions.
[**GetRule**](RulesApi.md#GetRule) | **Get** /v1/rules/{id} | Get a single rule.
[**ListRule**](RulesApi.md#ListRule) | **Get** /v1/rules | List all rules.
[**StoreRule**](RulesApi.md#StoreRule) | **Post** /v1/rules | Store a new rule
[**TestRule**](RulesApi.md#TestRule) | **Get** /v1/rules/{id}/test | Test which transactions would be hit by the rule. No changes will be made.
[**UpdateRule**](RulesApi.md#UpdateRule) | **Put** /v1/rules/{id} | Update existing rule.



## DeleteRule

> DeleteRule(ctx, id).XTraceId(xTraceId).Execute()

Delete an rule.



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
    id := "123" // string | The ID of the rule.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RulesApi.DeleteRule(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.DeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## FireRule

> FireRule(ctx, id).XTraceId(xTraceId).Start(start).End(end).Accounts(accounts).Execute()

Fire the rule on your transactions.



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
    id := "123" // string | The ID of the rule.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. If the start date is not present, it will be set to one year ago. If you use this field, both the start date and the end date must be present.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. If the end date is not present, it will be set to today. If you use this field, both the start date and the end date must be present.  (optional)
    accounts := []int64{int64(123)} // []int64 | Limit the triggering of the rule to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RulesApi.FireRule(context.Background(), id).XTraceId(xTraceId).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.FireRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFireRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **start** | **string** | A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. If the start date is not present, it will be set to one year ago. If you use this field, both the start date and the end date must be present.  | 
 **end** | **string** | A date formatted YYYY-MM-DD, to limit the transactions the actions will be applied to. If the end date is not present, it will be set to today. If you use this field, both the start date and the end date must be present.  | 
 **accounts** | **[]int64** | Limit the triggering of the rule to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  | 

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


## GetRule

> RuleSingle GetRule(ctx, id).XTraceId(xTraceId).Execute()

Get a single rule.



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
    id := "123" // string | The ID of the object.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.GetRule(context.Background(), id).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRule`: RuleSingle
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**RuleSingle**](RuleSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRule

> RuleArray ListRule(ctx).XTraceId(xTraceId).Page(page).Execute()

List all rules.



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
    resp, r, err := apiClient.RulesApi.ListRule(context.Background()).XTraceId(xTraceId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.ListRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRule`: RuleArray
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.ListRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRuleRequest struct via the builder pattern


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


## StoreRule

> RuleSingle StoreRule(ctx).RuleStore(ruleStore).XTraceId(xTraceId).Execute()

Store a new rule



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
    ruleStore := *openapiclient.NewRuleStore("First rule title.", "81", openapiclient.RuleTriggerType("store-journal"), []openapiclient.RuleTriggerStore{*openapiclient.NewRuleTriggerStore(openapiclient.RuleTriggerKeyword("from_account_starts"), "tag1")}, []openapiclient.RuleActionStore{*openapiclient.NewRuleActionStore(openapiclient.RuleActionKeyword("user_action"), "Daily groceries")}) // RuleStore | JSON array or key=value pairs with the necessary rule information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.StoreRule(context.Background()).RuleStore(ruleStore).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.StoreRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreRule`: RuleSingle
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.StoreRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleStore** | [**RuleStore**](RuleStore.md) | JSON array or key&#x3D;value pairs with the necessary rule information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**RuleSingle**](RuleSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRule

> TransactionArray TestRule(ctx, id).XTraceId(xTraceId).Start(start).End(end).Accounts(accounts).Execute()

Test which transactions would be hit by the rule. No changes will be made.



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
    id := "123" // string | The ID of the rule.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  (optional)
    accounts := []int64{int64(123)} // []int64 | Limit the testing of the rule to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.TestRule(context.Background(), id).XTraceId(xTraceId).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.TestRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestRule`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.TestRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTraceId** | **string** | Unique identifier associated with this request. | 
 **start** | **string** | A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  | 
 **end** | **string** | A date formatted YYYY-MM-DD, to limit the transactions the test will be applied to. Both the start date and the end date must be present.  | 
 **accounts** | **[]int64** | Limit the testing of the rule to these asset accounts or liabilities. Only asset accounts and liabilities will be accepted. Other types will be silently dropped.  | 

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


## UpdateRule

> RuleSingle UpdateRule(ctx, id).RuleUpdate(ruleUpdate).XTraceId(xTraceId).Execute()

Update existing rule.



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
    id := "123" // string | The ID of the object.
    ruleUpdate := *openapiclient.NewRuleUpdate() // RuleUpdate | JSON array with updated rule information. See the model for the exact specifications.
    xTraceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier associated with this request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.UpdateRule(context.Background(), id).RuleUpdate(ruleUpdate).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRule`: RuleSingle
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.UpdateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleUpdate** | [**RuleUpdate**](RuleUpdate.md) | JSON array with updated rule information. See the model for the exact specifications. | 
 **xTraceId** | **string** | Unique identifier associated with this request. | 

### Return type

[**RuleSingle**](RuleSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

