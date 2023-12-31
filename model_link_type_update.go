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

// checks if the LinkTypeUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkTypeUpdate{}

// LinkTypeUpdate struct for LinkTypeUpdate
type LinkTypeUpdate struct {
	Name *string `json:"name,omitempty"`
	Inward *string `json:"inward,omitempty"`
	Outward *string `json:"outward,omitempty"`
}

// NewLinkTypeUpdate instantiates a new LinkTypeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTypeUpdate() *LinkTypeUpdate {
	this := LinkTypeUpdate{}
	return &this
}

// NewLinkTypeUpdateWithDefaults instantiates a new LinkTypeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTypeUpdateWithDefaults() *LinkTypeUpdate {
	this := LinkTypeUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LinkTypeUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTypeUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LinkTypeUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LinkTypeUpdate) SetName(v string) {
	o.Name = &v
}

// GetInward returns the Inward field value if set, zero value otherwise.
func (o *LinkTypeUpdate) GetInward() string {
	if o == nil || IsNil(o.Inward) {
		var ret string
		return ret
	}
	return *o.Inward
}

// GetInwardOk returns a tuple with the Inward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTypeUpdate) GetInwardOk() (*string, bool) {
	if o == nil || IsNil(o.Inward) {
		return nil, false
	}
	return o.Inward, true
}

// HasInward returns a boolean if a field has been set.
func (o *LinkTypeUpdate) HasInward() bool {
	if o != nil && !IsNil(o.Inward) {
		return true
	}

	return false
}

// SetInward gets a reference to the given string and assigns it to the Inward field.
func (o *LinkTypeUpdate) SetInward(v string) {
	o.Inward = &v
}

// GetOutward returns the Outward field value if set, zero value otherwise.
func (o *LinkTypeUpdate) GetOutward() string {
	if o == nil || IsNil(o.Outward) {
		var ret string
		return ret
	}
	return *o.Outward
}

// GetOutwardOk returns a tuple with the Outward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTypeUpdate) GetOutwardOk() (*string, bool) {
	if o == nil || IsNil(o.Outward) {
		return nil, false
	}
	return o.Outward, true
}

// HasOutward returns a boolean if a field has been set.
func (o *LinkTypeUpdate) HasOutward() bool {
	if o != nil && !IsNil(o.Outward) {
		return true
	}

	return false
}

// SetOutward gets a reference to the given string and assigns it to the Outward field.
func (o *LinkTypeUpdate) SetOutward(v string) {
	o.Outward = &v
}

func (o LinkTypeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkTypeUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Inward) {
		toSerialize["inward"] = o.Inward
	}
	if !IsNil(o.Outward) {
		toSerialize["outward"] = o.Outward
	}
	return toSerialize, nil
}

type NullableLinkTypeUpdate struct {
	value *LinkTypeUpdate
	isSet bool
}

func (v NullableLinkTypeUpdate) Get() *LinkTypeUpdate {
	return v.value
}

func (v *NullableLinkTypeUpdate) Set(val *LinkTypeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTypeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTypeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTypeUpdate(val *LinkTypeUpdate) *NullableLinkTypeUpdate {
	return &NullableLinkTypeUpdate{value: val, isSet: true}
}

func (v NullableLinkTypeUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTypeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


