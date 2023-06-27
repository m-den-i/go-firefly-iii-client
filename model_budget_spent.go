/*
Firefly III API v2.0.4

This is the documentation of the Firefly III API. You can find accompanying documentation on the website of Firefly III itself (see below). Please report any bugs or issues. You may use the \"Authorize\" button to try the API below. This file was last generated on 2023-06-11T09:14:35+00:00 

API version: 2.0.4
Contact: james@firefly-iii.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BudgetSpent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetSpent{}

// BudgetSpent struct for BudgetSpent
type BudgetSpent struct {
	// The amount spent.
	Sum *string `json:"sum,omitempty"`
	CurrencyId *string `json:"currency_id,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	// Number of decimals supported by the currency
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
}

// NewBudgetSpent instantiates a new BudgetSpent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetSpent() *BudgetSpent {
	this := BudgetSpent{}
	return &this
}

// NewBudgetSpentWithDefaults instantiates a new BudgetSpent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetSpentWithDefaults() *BudgetSpent {
	this := BudgetSpent{}
	return &this
}

// GetSum returns the Sum field value if set, zero value otherwise.
func (o *BudgetSpent) GetSum() string {
	if o == nil || IsNil(o.Sum) {
		var ret string
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetSpent) GetSumOk() (*string, bool) {
	if o == nil || IsNil(o.Sum) {
		return nil, false
	}
	return o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *BudgetSpent) HasSum() bool {
	if o != nil && !IsNil(o.Sum) {
		return true
	}

	return false
}

// SetSum gets a reference to the given string and assigns it to the Sum field.
func (o *BudgetSpent) SetSum(v string) {
	o.Sum = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *BudgetSpent) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetSpent) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *BudgetSpent) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *BudgetSpent) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BudgetSpent) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetSpent) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BudgetSpent) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BudgetSpent) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *BudgetSpent) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetSpent) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencySymbol) {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *BudgetSpent) HasCurrencySymbol() bool {
	if o != nil && !IsNil(o.CurrencySymbol) {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *BudgetSpent) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *BudgetSpent) GetCurrencyDecimalPlaces() int32 {
	if o == nil || IsNil(o.CurrencyDecimalPlaces) {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetSpent) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrencyDecimalPlaces) {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *BudgetSpent) HasCurrencyDecimalPlaces() bool {
	if o != nil && !IsNil(o.CurrencyDecimalPlaces) {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *BudgetSpent) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

func (o BudgetSpent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BudgetSpent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sum) {
		toSerialize["sum"] = o.Sum
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.CurrencySymbol) {
		toSerialize["currency_symbol"] = o.CurrencySymbol
	}
	if !IsNil(o.CurrencyDecimalPlaces) {
		toSerialize["currency_decimal_places"] = o.CurrencyDecimalPlaces
	}
	return toSerialize, nil
}

type NullableBudgetSpent struct {
	value *BudgetSpent
	isSet bool
}

func (v NullableBudgetSpent) Get() *BudgetSpent {
	return v.value
}

func (v *NullableBudgetSpent) Set(val *BudgetSpent) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetSpent) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetSpent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetSpent(val *BudgetSpent) *NullableBudgetSpent {
	return &NullableBudgetSpent{value: val, isSet: true}
}

func (v NullableBudgetSpent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetSpent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

