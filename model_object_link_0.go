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

// checks if the ObjectLink0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectLink0{}

// ObjectLink0 struct for ObjectLink0
type ObjectLink0 struct {
	Rel *string `json:"rel,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewObjectLink0 instantiates a new ObjectLink0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectLink0() *ObjectLink0 {
	this := ObjectLink0{}
	return &this
}

// NewObjectLink0WithDefaults instantiates a new ObjectLink0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectLink0WithDefaults() *ObjectLink0 {
	this := ObjectLink0{}
	return &this
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *ObjectLink0) GetRel() string {
	if o == nil || IsNil(o.Rel) {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectLink0) GetRelOk() (*string, bool) {
	if o == nil || IsNil(o.Rel) {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *ObjectLink0) HasRel() bool {
	if o != nil && !IsNil(o.Rel) {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *ObjectLink0) SetRel(v string) {
	o.Rel = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ObjectLink0) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectLink0) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ObjectLink0) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ObjectLink0) SetUri(v string) {
	o.Uri = &v
}

func (o ObjectLink0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectLink0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rel) {
		toSerialize["rel"] = o.Rel
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableObjectLink0 struct {
	value *ObjectLink0
	isSet bool
}

func (v NullableObjectLink0) Get() *ObjectLink0 {
	return v.value
}

func (v *NullableObjectLink0) Set(val *ObjectLink0) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectLink0) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectLink0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectLink0(val *ObjectLink0) *NullableObjectLink0 {
	return &NullableObjectLink0{value: val, isSet: true}
}

func (v NullableObjectLink0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectLink0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


