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

// checks if the BillRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillRead{}

// BillRead struct for BillRead
type BillRead struct {
	// Immutable value
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes Bill `json:"attributes"`
}

// NewBillRead instantiates a new BillRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillRead(type_ string, id string, attributes Bill) *BillRead {
	this := BillRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewBillReadWithDefaults instantiates a new BillRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillReadWithDefaults() *BillRead {
	this := BillRead{}
	return &this
}

// GetType returns the Type field value
func (o *BillRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BillRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BillRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BillRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BillRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BillRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *BillRead) GetAttributes() Bill {
	if o == nil {
		var ret Bill
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BillRead) GetAttributesOk() (*Bill, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BillRead) SetAttributes(v Bill) {
	o.Attributes = v
}

func (o BillRead) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

type NullableBillRead struct {
	value *BillRead
	isSet bool
}

func (v NullableBillRead) Get() *BillRead {
	return v.value
}

func (v *NullableBillRead) Set(val *BillRead) {
	v.value = val
	v.isSet = true
}

func (v NullableBillRead) IsSet() bool {
	return v.isSet
}

func (v *NullableBillRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillRead(val *BillRead) *NullableBillRead {
	return &NullableBillRead{value: val, isSet: true}
}

func (v NullableBillRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


