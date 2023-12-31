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

// checks if the AutocompleteRuleGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutocompleteRuleGroup{}

// AutocompleteRuleGroup struct for AutocompleteRuleGroup
type AutocompleteRuleGroup struct {
	Id string `json:"id"`
	// Name of the rule group found by an auto-complete search.
	Name string `json:"name"`
	// Description of the rule group found by auto-complete.
	Description *string `json:"description,omitempty"`
}

// NewAutocompleteRuleGroup instantiates a new AutocompleteRuleGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteRuleGroup(id string, name string) *AutocompleteRuleGroup {
	this := AutocompleteRuleGroup{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAutocompleteRuleGroupWithDefaults instantiates a new AutocompleteRuleGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteRuleGroupWithDefaults() *AutocompleteRuleGroup {
	this := AutocompleteRuleGroup{}
	return &this
}

// GetId returns the Id field value
func (o *AutocompleteRuleGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutocompleteRuleGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutocompleteRuleGroup) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutocompleteRuleGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutocompleteRuleGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutocompleteRuleGroup) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AutocompleteRuleGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteRuleGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AutocompleteRuleGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AutocompleteRuleGroup) SetDescription(v string) {
	o.Description = &v
}

func (o AutocompleteRuleGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutocompleteRuleGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAutocompleteRuleGroup struct {
	value *AutocompleteRuleGroup
	isSet bool
}

func (v NullableAutocompleteRuleGroup) Get() *AutocompleteRuleGroup {
	return v.value
}

func (v *NullableAutocompleteRuleGroup) Set(val *AutocompleteRuleGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteRuleGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteRuleGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteRuleGroup(val *AutocompleteRuleGroup) *NullableAutocompleteRuleGroup {
	return &NullableAutocompleteRuleGroup{value: val, isSet: true}
}

func (v NullableAutocompleteRuleGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteRuleGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


