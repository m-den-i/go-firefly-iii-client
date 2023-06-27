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

// checks if the RuleTriggerUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleTriggerUpdate{}

// RuleTriggerUpdate struct for RuleTriggerUpdate
type RuleTriggerUpdate struct {
	Type *RuleTriggerKeyword `json:"type,omitempty"`
	// The accompanying value the trigger responds to. This value is often mandatory, but this depends on the trigger.
	Value *string `json:"value,omitempty"`
	// Order of the trigger
	Order *int32 `json:"order,omitempty"`
	// If the trigger is active.
	Active *bool `json:"active,omitempty"`
	// When true, other triggers will not be checked if this trigger was triggered.
	StopProcessing *bool `json:"stop_processing,omitempty"`
}

// NewRuleTriggerUpdate instantiates a new RuleTriggerUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleTriggerUpdate() *RuleTriggerUpdate {
	this := RuleTriggerUpdate{}
	return &this
}

// NewRuleTriggerUpdateWithDefaults instantiates a new RuleTriggerUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleTriggerUpdateWithDefaults() *RuleTriggerUpdate {
	this := RuleTriggerUpdate{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RuleTriggerUpdate) GetType() RuleTriggerKeyword {
	if o == nil || IsNil(o.Type) {
		var ret RuleTriggerKeyword
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleTriggerUpdate) GetTypeOk() (*RuleTriggerKeyword, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RuleTriggerUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RuleTriggerKeyword and assigns it to the Type field.
func (o *RuleTriggerUpdate) SetType(v RuleTriggerKeyword) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RuleTriggerUpdate) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleTriggerUpdate) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RuleTriggerUpdate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RuleTriggerUpdate) SetValue(v string) {
	o.Value = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RuleTriggerUpdate) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleTriggerUpdate) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RuleTriggerUpdate) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *RuleTriggerUpdate) SetOrder(v int32) {
	o.Order = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RuleTriggerUpdate) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleTriggerUpdate) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RuleTriggerUpdate) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *RuleTriggerUpdate) SetActive(v bool) {
	o.Active = &v
}

// GetStopProcessing returns the StopProcessing field value if set, zero value otherwise.
func (o *RuleTriggerUpdate) GetStopProcessing() bool {
	if o == nil || IsNil(o.StopProcessing) {
		var ret bool
		return ret
	}
	return *o.StopProcessing
}

// GetStopProcessingOk returns a tuple with the StopProcessing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleTriggerUpdate) GetStopProcessingOk() (*bool, bool) {
	if o == nil || IsNil(o.StopProcessing) {
		return nil, false
	}
	return o.StopProcessing, true
}

// HasStopProcessing returns a boolean if a field has been set.
func (o *RuleTriggerUpdate) HasStopProcessing() bool {
	if o != nil && !IsNil(o.StopProcessing) {
		return true
	}

	return false
}

// SetStopProcessing gets a reference to the given bool and assigns it to the StopProcessing field.
func (o *RuleTriggerUpdate) SetStopProcessing(v bool) {
	o.StopProcessing = &v
}

func (o RuleTriggerUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleTriggerUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.StopProcessing) {
		toSerialize["stop_processing"] = o.StopProcessing
	}
	return toSerialize, nil
}

type NullableRuleTriggerUpdate struct {
	value *RuleTriggerUpdate
	isSet bool
}

func (v NullableRuleTriggerUpdate) Get() *RuleTriggerUpdate {
	return v.value
}

func (v *NullableRuleTriggerUpdate) Set(val *RuleTriggerUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleTriggerUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleTriggerUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleTriggerUpdate(val *RuleTriggerUpdate) *NullableRuleTriggerUpdate {
	return &NullableRuleTriggerUpdate{value: val, isSet: true}
}

func (v NullableRuleTriggerUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleTriggerUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

