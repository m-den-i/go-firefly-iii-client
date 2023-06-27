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

// checks if the TransactionLinkSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionLinkSingle{}

// TransactionLinkSingle struct for TransactionLinkSingle
type TransactionLinkSingle struct {
	Data TransactionLinkRead `json:"data"`
}

// NewTransactionLinkSingle instantiates a new TransactionLinkSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionLinkSingle(data TransactionLinkRead) *TransactionLinkSingle {
	this := TransactionLinkSingle{}
	this.Data = data
	return &this
}

// NewTransactionLinkSingleWithDefaults instantiates a new TransactionLinkSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionLinkSingleWithDefaults() *TransactionLinkSingle {
	this := TransactionLinkSingle{}
	return &this
}

// GetData returns the Data field value
func (o *TransactionLinkSingle) GetData() TransactionLinkRead {
	if o == nil {
		var ret TransactionLinkRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionLinkSingle) GetDataOk() (*TransactionLinkRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionLinkSingle) SetData(v TransactionLinkRead) {
	o.Data = v
}

func (o TransactionLinkSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionLinkSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTransactionLinkSingle struct {
	value *TransactionLinkSingle
	isSet bool
}

func (v NullableTransactionLinkSingle) Get() *TransactionLinkSingle {
	return v.value
}

func (v *NullableTransactionLinkSingle) Set(val *TransactionLinkSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionLinkSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionLinkSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionLinkSingle(val *TransactionLinkSingle) *NullableTransactionLinkSingle {
	return &NullableTransactionLinkSingle{value: val, isSet: true}
}

func (v NullableTransactionLinkSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionLinkSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


