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

// checks if the UserArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserArray{}

// UserArray struct for UserArray
type UserArray struct {
	Data []UserRead `json:"data"`
	Meta Meta `json:"meta"`
	Links PageLink `json:"links"`
}

// NewUserArray instantiates a new UserArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserArray(data []UserRead, meta Meta, links PageLink) *UserArray {
	this := UserArray{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewUserArrayWithDefaults instantiates a new UserArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserArrayWithDefaults() *UserArray {
	this := UserArray{}
	return &this
}

// GetData returns the Data field value
func (o *UserArray) GetData() []UserRead {
	if o == nil {
		var ret []UserRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserArray) GetDataOk() ([]UserRead, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *UserArray) SetData(v []UserRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *UserArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *UserArray) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *UserArray) SetMeta(v Meta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *UserArray) GetLinks() PageLink {
	if o == nil {
		var ret PageLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UserArray) GetLinksOk() (*PageLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *UserArray) SetLinks(v PageLink) {
	o.Links = v
}

func (o UserArray) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableUserArray struct {
	value *UserArray
	isSet bool
}

func (v NullableUserArray) Get() *UserArray {
	return v.value
}

func (v *NullableUserArray) Set(val *UserArray) {
	v.value = val
	v.isSet = true
}

func (v NullableUserArray) IsSet() bool {
	return v.isSet
}

func (v *NullableUserArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserArray(val *UserArray) *NullableUserArray {
	return &NullableUserArray{value: val, isSet: true}
}

func (v NullableUserArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


