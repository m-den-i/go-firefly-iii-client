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

// checks if the AutocompleteBill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutocompleteBill{}

// AutocompleteBill struct for AutocompleteBill
type AutocompleteBill struct {
	Id string `json:"id"`
	// Name of the bill found by an auto-complete search.
	Name string `json:"name"`
	// Is the bill active or not?
	Active *bool `json:"active,omitempty"`
}

// NewAutocompleteBill instantiates a new AutocompleteBill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteBill(id string, name string) *AutocompleteBill {
	this := AutocompleteBill{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAutocompleteBillWithDefaults instantiates a new AutocompleteBill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteBillWithDefaults() *AutocompleteBill {
	this := AutocompleteBill{}
	return &this
}

// GetId returns the Id field value
func (o *AutocompleteBill) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutocompleteBill) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutocompleteBill) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutocompleteBill) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutocompleteBill) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutocompleteBill) SetName(v string) {
	o.Name = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AutocompleteBill) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteBill) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AutocompleteBill) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AutocompleteBill) SetActive(v bool) {
	o.Active = &v
}

func (o AutocompleteBill) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutocompleteBill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableAutocompleteBill struct {
	value *AutocompleteBill
	isSet bool
}

func (v NullableAutocompleteBill) Get() *AutocompleteBill {
	return v.value
}

func (v *NullableAutocompleteBill) Set(val *AutocompleteBill) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteBill) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteBill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteBill(val *AutocompleteBill) *NullableAutocompleteBill {
	return &NullableAutocompleteBill{value: val, isSet: true}
}

func (v NullableAutocompleteBill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteBill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

