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

// checks if the ObjectGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectGroup{}

// ObjectGroup struct for ObjectGroup
type ObjectGroup struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Title string `json:"title"`
	// Order of the object group
	Order int32 `json:"order"`
}

// NewObjectGroup instantiates a new ObjectGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectGroup(title string, order int32) *ObjectGroup {
	this := ObjectGroup{}
	this.Title = title
	this.Order = order
	return &this
}

// NewObjectGroupWithDefaults instantiates a new ObjectGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectGroupWithDefaults() *ObjectGroup {
	this := ObjectGroup{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ObjectGroup) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectGroup) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ObjectGroup) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ObjectGroup) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ObjectGroup) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectGroup) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ObjectGroup) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ObjectGroup) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetTitle returns the Title field value
func (o *ObjectGroup) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ObjectGroup) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ObjectGroup) SetTitle(v string) {
	o.Title = v
}

// GetOrder returns the Order field value
func (o *ObjectGroup) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ObjectGroup) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ObjectGroup) SetOrder(v int32) {
	o.Order = v
}

func (o ObjectGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	toSerialize["title"] = o.Title
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

type NullableObjectGroup struct {
	value *ObjectGroup
	isSet bool
}

func (v NullableObjectGroup) Get() *ObjectGroup {
	return v.value
}

func (v *NullableObjectGroup) Set(val *ObjectGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectGroup(val *ObjectGroup) *NullableObjectGroup {
	return &NullableObjectGroup{value: val, isSet: true}
}

func (v NullableObjectGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


