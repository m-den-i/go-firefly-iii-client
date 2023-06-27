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

// checks if the CurrencySingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrencySingle{}

// CurrencySingle struct for CurrencySingle
type CurrencySingle struct {
	Data CurrencyRead `json:"data"`
}

// NewCurrencySingle instantiates a new CurrencySingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencySingle(data CurrencyRead) *CurrencySingle {
	this := CurrencySingle{}
	this.Data = data
	return &this
}

// NewCurrencySingleWithDefaults instantiates a new CurrencySingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencySingleWithDefaults() *CurrencySingle {
	this := CurrencySingle{}
	return &this
}

// GetData returns the Data field value
func (o *CurrencySingle) GetData() CurrencyRead {
	if o == nil {
		var ret CurrencyRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CurrencySingle) GetDataOk() (*CurrencyRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CurrencySingle) SetData(v CurrencyRead) {
	o.Data = v
}

func (o CurrencySingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrencySingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCurrencySingle struct {
	value *CurrencySingle
	isSet bool
}

func (v NullableCurrencySingle) Get() *CurrencySingle {
	return v.value
}

func (v *NullableCurrencySingle) Set(val *CurrencySingle) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencySingle) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencySingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencySingle(val *CurrencySingle) *NullableCurrencySingle {
	return &NullableCurrencySingle{value: val, isSet: true}
}

func (v NullableCurrencySingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencySingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


