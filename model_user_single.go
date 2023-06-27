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

// checks if the UserSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSingle{}

// UserSingle struct for UserSingle
type UserSingle struct {
	Data UserRead `json:"data"`
}

// NewUserSingle instantiates a new UserSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSingle(data UserRead) *UserSingle {
	this := UserSingle{}
	this.Data = data
	return &this
}

// NewUserSingleWithDefaults instantiates a new UserSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSingleWithDefaults() *UserSingle {
	this := UserSingle{}
	return &this
}

// GetData returns the Data field value
func (o *UserSingle) GetData() UserRead {
	if o == nil {
		var ret UserRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserSingle) GetDataOk() (*UserRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UserSingle) SetData(v UserRead) {
	o.Data = v
}

func (o UserSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableUserSingle struct {
	value *UserSingle
	isSet bool
}

func (v NullableUserSingle) Get() *UserSingle {
	return v.value
}

func (v *NullableUserSingle) Set(val *UserSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSingle(val *UserSingle) *NullableUserSingle {
	return &NullableUserSingle{value: val, isSet: true}
}

func (v NullableUserSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


