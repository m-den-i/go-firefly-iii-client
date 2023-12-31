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

// checks if the LinkTypeStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkTypeStore{}

// LinkTypeStore struct for LinkTypeStore
type LinkTypeStore struct {
	Name string `json:"name"`
	Inward string `json:"inward"`
	Outward string `json:"outward"`
}

// NewLinkTypeStore instantiates a new LinkTypeStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTypeStore(name string, inward string, outward string) *LinkTypeStore {
	this := LinkTypeStore{}
	this.Name = name
	this.Inward = inward
	this.Outward = outward
	return &this
}

// NewLinkTypeStoreWithDefaults instantiates a new LinkTypeStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTypeStoreWithDefaults() *LinkTypeStore {
	this := LinkTypeStore{}
	return &this
}

// GetName returns the Name field value
func (o *LinkTypeStore) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LinkTypeStore) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LinkTypeStore) SetName(v string) {
	o.Name = v
}

// GetInward returns the Inward field value
func (o *LinkTypeStore) GetInward() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Inward
}

// GetInwardOk returns a tuple with the Inward field value
// and a boolean to check if the value has been set.
func (o *LinkTypeStore) GetInwardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inward, true
}

// SetInward sets field value
func (o *LinkTypeStore) SetInward(v string) {
	o.Inward = v
}

// GetOutward returns the Outward field value
func (o *LinkTypeStore) GetOutward() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outward
}

// GetOutwardOk returns a tuple with the Outward field value
// and a boolean to check if the value has been set.
func (o *LinkTypeStore) GetOutwardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outward, true
}

// SetOutward sets field value
func (o *LinkTypeStore) SetOutward(v string) {
	o.Outward = v
}

func (o LinkTypeStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkTypeStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["inward"] = o.Inward
	toSerialize["outward"] = o.Outward
	return toSerialize, nil
}

type NullableLinkTypeStore struct {
	value *LinkTypeStore
	isSet bool
}

func (v NullableLinkTypeStore) Get() *LinkTypeStore {
	return v.value
}

func (v *NullableLinkTypeStore) Set(val *LinkTypeStore) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTypeStore) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTypeStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTypeStore(val *LinkTypeStore) *NullableLinkTypeStore {
	return &NullableLinkTypeStore{value: val, isSet: true}
}

func (v NullableLinkTypeStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTypeStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


