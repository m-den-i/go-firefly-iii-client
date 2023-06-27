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

// checks if the AccountSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountSingle{}

// AccountSingle struct for AccountSingle
type AccountSingle struct {
	Data AccountRead `json:"data"`
}

// NewAccountSingle instantiates a new AccountSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSingle(data AccountRead) *AccountSingle {
	this := AccountSingle{}
	this.Data = data
	return &this
}

// NewAccountSingleWithDefaults instantiates a new AccountSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSingleWithDefaults() *AccountSingle {
	this := AccountSingle{}
	return &this
}

// GetData returns the Data field value
func (o *AccountSingle) GetData() AccountRead {
	if o == nil {
		var ret AccountRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccountSingle) GetDataOk() (*AccountRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AccountSingle) SetData(v AccountRead) {
	o.Data = v
}

func (o AccountSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAccountSingle struct {
	value *AccountSingle
	isSet bool
}

func (v NullableAccountSingle) Get() *AccountSingle {
	return v.value
}

func (v *NullableAccountSingle) Set(val *AccountSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSingle(val *AccountSingle) *NullableAccountSingle {
	return &NullableAccountSingle{value: val, isSet: true}
}

func (v NullableAccountSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


