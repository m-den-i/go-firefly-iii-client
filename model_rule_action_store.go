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

// checks if the RuleActionStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleActionStore{}

// RuleActionStore struct for RuleActionStore
type RuleActionStore struct {
	Type RuleActionKeyword `json:"type"`
	// The accompanying value the action will set, change or update. Can be empty, but for some types this value is mandatory.
	Value NullableString `json:"value"`
	// Order of the action
	Order *int32 `json:"order,omitempty"`
	// If the action is active. Defaults to true.
	Active *bool `json:"active,omitempty"`
	// When true, other actions will not be fired after this action has fired. Defaults to false.
	StopProcessing *bool `json:"stop_processing,omitempty"`
}

// NewRuleActionStore instantiates a new RuleActionStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleActionStore(type_ RuleActionKeyword, value NullableString) *RuleActionStore {
	this := RuleActionStore{}
	this.Type = type_
	this.Value = value
	var active bool = true
	this.Active = &active
	var stopProcessing bool = false
	this.StopProcessing = &stopProcessing
	return &this
}

// NewRuleActionStoreWithDefaults instantiates a new RuleActionStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleActionStoreWithDefaults() *RuleActionStore {
	this := RuleActionStore{}
	var active bool = true
	this.Active = &active
	var stopProcessing bool = false
	this.StopProcessing = &stopProcessing
	return &this
}

// GetType returns the Type field value
func (o *RuleActionStore) GetType() RuleActionKeyword {
	if o == nil {
		var ret RuleActionKeyword
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RuleActionStore) GetTypeOk() (*RuleActionKeyword, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RuleActionStore) SetType(v RuleActionKeyword) {
	o.Type = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RuleActionStore) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleActionStore) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *RuleActionStore) SetValue(v string) {
	o.Value.Set(&v)
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RuleActionStore) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleActionStore) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RuleActionStore) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *RuleActionStore) SetOrder(v int32) {
	o.Order = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RuleActionStore) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleActionStore) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RuleActionStore) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *RuleActionStore) SetActive(v bool) {
	o.Active = &v
}

// GetStopProcessing returns the StopProcessing field value if set, zero value otherwise.
func (o *RuleActionStore) GetStopProcessing() bool {
	if o == nil || IsNil(o.StopProcessing) {
		var ret bool
		return ret
	}
	return *o.StopProcessing
}

// GetStopProcessingOk returns a tuple with the StopProcessing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleActionStore) GetStopProcessingOk() (*bool, bool) {
	if o == nil || IsNil(o.StopProcessing) {
		return nil, false
	}
	return o.StopProcessing, true
}

// HasStopProcessing returns a boolean if a field has been set.
func (o *RuleActionStore) HasStopProcessing() bool {
	if o != nil && !IsNil(o.StopProcessing) {
		return true
	}

	return false
}

// SetStopProcessing gets a reference to the given bool and assigns it to the StopProcessing field.
func (o *RuleActionStore) SetStopProcessing(v bool) {
	o.StopProcessing = &v
}

func (o RuleActionStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleActionStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value.Get()
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

type NullableRuleActionStore struct {
	value *RuleActionStore
	isSet bool
}

func (v NullableRuleActionStore) Get() *RuleActionStore {
	return v.value
}

func (v *NullableRuleActionStore) Set(val *RuleActionStore) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleActionStore) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleActionStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleActionStore(val *RuleActionStore) *NullableRuleActionStore {
	return &NullableRuleActionStore{value: val, isSet: true}
}

func (v NullableRuleActionStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleActionStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


