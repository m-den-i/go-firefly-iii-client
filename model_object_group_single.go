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

// checks if the ObjectGroupSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectGroupSingle{}

// ObjectGroupSingle struct for ObjectGroupSingle
type ObjectGroupSingle struct {
	Data ObjectGroupRead `json:"data"`
}

// NewObjectGroupSingle instantiates a new ObjectGroupSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectGroupSingle(data ObjectGroupRead) *ObjectGroupSingle {
	this := ObjectGroupSingle{}
	this.Data = data
	return &this
}

// NewObjectGroupSingleWithDefaults instantiates a new ObjectGroupSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectGroupSingleWithDefaults() *ObjectGroupSingle {
	this := ObjectGroupSingle{}
	return &this
}

// GetData returns the Data field value
func (o *ObjectGroupSingle) GetData() ObjectGroupRead {
	if o == nil {
		var ret ObjectGroupRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ObjectGroupSingle) GetDataOk() (*ObjectGroupRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ObjectGroupSingle) SetData(v ObjectGroupRead) {
	o.Data = v
}

func (o ObjectGroupSingle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectGroupSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableObjectGroupSingle struct {
	value *ObjectGroupSingle
	isSet bool
}

func (v NullableObjectGroupSingle) Get() *ObjectGroupSingle {
	return v.value
}

func (v *NullableObjectGroupSingle) Set(val *ObjectGroupSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectGroupSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectGroupSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectGroupSingle(val *ObjectGroupSingle) *NullableObjectGroupSingle {
	return &NullableObjectGroupSingle{value: val, isSet: true}
}

func (v NullableObjectGroupSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectGroupSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


