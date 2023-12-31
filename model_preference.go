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
	"time"
)

// checks if the Preference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Preference{}

// Preference struct for Preference
type Preference struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Name string `json:"name"`
	Data PolymorphicProperty `json:"data"`
}

// NewPreference instantiates a new Preference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreference(name string, data PolymorphicProperty) *Preference {
	this := Preference{}
	this.Name = name
	this.Data = data
	return &this
}

// NewPreferenceWithDefaults instantiates a new Preference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferenceWithDefaults() *Preference {
	this := Preference{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Preference) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preference) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Preference) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Preference) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Preference) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Preference) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Preference) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Preference) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value
func (o *Preference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Preference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Preference) SetName(v string) {
	o.Name = v
}

// GetData returns the Data field value
func (o *Preference) GetData() PolymorphicProperty {
	if o == nil {
		var ret PolymorphicProperty
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Preference) GetDataOk() (*PolymorphicProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Preference) SetData(v PolymorphicProperty) {
	o.Data = v
}

func (o Preference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Preference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	toSerialize["name"] = o.Name
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePreference struct {
	value *Preference
	isSet bool
}

func (v NullablePreference) Get() *Preference {
	return v.value
}

func (v *NullablePreference) Set(val *Preference) {
	v.value = val
	v.isSet = true
}

func (v NullablePreference) IsSet() bool {
	return v.isSet
}

func (v *NullablePreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreference(val *Preference) *NullablePreference {
	return &NullablePreference{value: val, isSet: true}
}

func (v NullablePreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


