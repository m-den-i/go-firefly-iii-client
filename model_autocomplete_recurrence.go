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

// checks if the AutocompleteRecurrence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutocompleteRecurrence{}

// AutocompleteRecurrence struct for AutocompleteRecurrence
type AutocompleteRecurrence struct {
	Id string `json:"id"`
	// Name of the recurrence found by an auto-complete search.
	Name string `json:"name"`
	// Description of the recurrence found by auto-complete.
	Description *string `json:"description,omitempty"`
}

// NewAutocompleteRecurrence instantiates a new AutocompleteRecurrence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteRecurrence(id string, name string) *AutocompleteRecurrence {
	this := AutocompleteRecurrence{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAutocompleteRecurrenceWithDefaults instantiates a new AutocompleteRecurrence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteRecurrenceWithDefaults() *AutocompleteRecurrence {
	this := AutocompleteRecurrence{}
	return &this
}

// GetId returns the Id field value
func (o *AutocompleteRecurrence) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutocompleteRecurrence) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutocompleteRecurrence) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutocompleteRecurrence) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutocompleteRecurrence) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutocompleteRecurrence) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AutocompleteRecurrence) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteRecurrence) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AutocompleteRecurrence) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AutocompleteRecurrence) SetDescription(v string) {
	o.Description = &v
}

func (o AutocompleteRecurrence) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutocompleteRecurrence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAutocompleteRecurrence struct {
	value *AutocompleteRecurrence
	isSet bool
}

func (v NullableAutocompleteRecurrence) Get() *AutocompleteRecurrence {
	return v.value
}

func (v *NullableAutocompleteRecurrence) Set(val *AutocompleteRecurrence) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteRecurrence) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteRecurrence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteRecurrence(val *AutocompleteRecurrence) *NullableAutocompleteRecurrence {
	return &NullableAutocompleteRecurrence{value: val, isSet: true}
}

func (v NullableAutocompleteRecurrence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteRecurrence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


