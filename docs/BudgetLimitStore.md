# BudgetLimitStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **string** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**CurrencyCode** | Pointer to **string** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**BudgetId** | **string** | The budget ID of the associated budget. | [readonly] 
**Start** | **string** | Start date of the budget limit. | 
**Period** | Pointer to **NullableString** | Period of the budget limit. Only used when auto-generated by auto-budget. | [optional] [readonly] 
**End** | **string** | End date of the budget limit. | 
**Amount** | **string** |  | 

## Methods

### NewBudgetLimitStore

`func NewBudgetLimitStore(budgetId string, start string, end string, amount string, ) *BudgetLimitStore`

NewBudgetLimitStore instantiates a new BudgetLimitStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetLimitStoreWithDefaults

`func NewBudgetLimitStoreWithDefaults() *BudgetLimitStore`

NewBudgetLimitStoreWithDefaults instantiates a new BudgetLimitStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *BudgetLimitStore) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BudgetLimitStore) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BudgetLimitStore) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BudgetLimitStore) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BudgetLimitStore) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BudgetLimitStore) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BudgetLimitStore) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *BudgetLimitStore) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetBudgetId

`func (o *BudgetLimitStore) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *BudgetLimitStore) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *BudgetLimitStore) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetStart

`func (o *BudgetLimitStore) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *BudgetLimitStore) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *BudgetLimitStore) SetStart(v string)`

SetStart sets Start field to given value.


### GetPeriod

`func (o *BudgetLimitStore) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BudgetLimitStore) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BudgetLimitStore) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *BudgetLimitStore) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *BudgetLimitStore) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *BudgetLimitStore) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetEnd

`func (o *BudgetLimitStore) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *BudgetLimitStore) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *BudgetLimitStore) SetEnd(v string)`

SetEnd sets End field to given value.


### GetAmount

`func (o *BudgetLimitStore) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BudgetLimitStore) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BudgetLimitStore) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

