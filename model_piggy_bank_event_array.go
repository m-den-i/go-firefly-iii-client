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

// checks if the PiggyBankEventArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PiggyBankEventArray{}

// PiggyBankEventArray struct for PiggyBankEventArray
type PiggyBankEventArray struct {
	Data []PiggyBankEventRead `json:"data"`
	Meta Meta `json:"meta"`
	Links PageLink `json:"links"`
}

// NewPiggyBankEventArray instantiates a new PiggyBankEventArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPiggyBankEventArray(data []PiggyBankEventRead, meta Meta, links PageLink) *PiggyBankEventArray {
	this := PiggyBankEventArray{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewPiggyBankEventArrayWithDefaults instantiates a new PiggyBankEventArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPiggyBankEventArrayWithDefaults() *PiggyBankEventArray {
	this := PiggyBankEventArray{}
	return &this
}

// GetData returns the Data field value
func (o *PiggyBankEventArray) GetData() []PiggyBankEventRead {
	if o == nil {
		var ret []PiggyBankEventRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventArray) GetDataOk() ([]PiggyBankEventRead, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PiggyBankEventArray) SetData(v []PiggyBankEventRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *PiggyBankEventArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventArray) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *PiggyBankEventArray) SetMeta(v Meta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *PiggyBankEventArray) GetLinks() PageLink {
	if o == nil {
		var ret PageLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventArray) GetLinksOk() (*PageLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PiggyBankEventArray) SetLinks(v PageLink) {
	o.Links = v
}

func (o PiggyBankEventArray) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PiggyBankEventArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullablePiggyBankEventArray struct {
	value *PiggyBankEventArray
	isSet bool
}

func (v NullablePiggyBankEventArray) Get() *PiggyBankEventArray {
	return v.value
}

func (v *NullablePiggyBankEventArray) Set(val *PiggyBankEventArray) {
	v.value = val
	v.isSet = true
}

func (v NullablePiggyBankEventArray) IsSet() bool {
	return v.isSet
}

func (v *NullablePiggyBankEventArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePiggyBankEventArray(val *PiggyBankEventArray) *NullablePiggyBankEventArray {
	return &NullablePiggyBankEventArray{value: val, isSet: true}
}

func (v NullablePiggyBankEventArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePiggyBankEventArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


