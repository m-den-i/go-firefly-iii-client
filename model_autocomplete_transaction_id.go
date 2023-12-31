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

// checks if the AutocompleteTransactionID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutocompleteTransactionID{}

// AutocompleteTransactionID struct for AutocompleteTransactionID
type AutocompleteTransactionID struct {
	// The ID of a transaction journal (basically a single split).
	Id string `json:"id"`
	// The ID of the underlying transaction group.
	TransactionGroupId *string `json:"transaction_group_id,omitempty"`
	// Transaction description with ID in the name.
	Name string `json:"name"`
	// Transaction description with ID in the name.
	Description string `json:"description"`
}

// NewAutocompleteTransactionID instantiates a new AutocompleteTransactionID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteTransactionID(id string, name string, description string) *AutocompleteTransactionID {
	this := AutocompleteTransactionID{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewAutocompleteTransactionIDWithDefaults instantiates a new AutocompleteTransactionID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteTransactionIDWithDefaults() *AutocompleteTransactionID {
	this := AutocompleteTransactionID{}
	return &this
}

// GetId returns the Id field value
func (o *AutocompleteTransactionID) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutocompleteTransactionID) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutocompleteTransactionID) SetId(v string) {
	o.Id = v
}

// GetTransactionGroupId returns the TransactionGroupId field value if set, zero value otherwise.
func (o *AutocompleteTransactionID) GetTransactionGroupId() string {
	if o == nil || IsNil(o.TransactionGroupId) {
		var ret string
		return ret
	}
	return *o.TransactionGroupId
}

// GetTransactionGroupIdOk returns a tuple with the TransactionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteTransactionID) GetTransactionGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionGroupId) {
		return nil, false
	}
	return o.TransactionGroupId, true
}

// HasTransactionGroupId returns a boolean if a field has been set.
func (o *AutocompleteTransactionID) HasTransactionGroupId() bool {
	if o != nil && !IsNil(o.TransactionGroupId) {
		return true
	}

	return false
}

// SetTransactionGroupId gets a reference to the given string and assigns it to the TransactionGroupId field.
func (o *AutocompleteTransactionID) SetTransactionGroupId(v string) {
	o.TransactionGroupId = &v
}

// GetName returns the Name field value
func (o *AutocompleteTransactionID) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutocompleteTransactionID) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutocompleteTransactionID) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AutocompleteTransactionID) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AutocompleteTransactionID) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AutocompleteTransactionID) SetDescription(v string) {
	o.Description = v
}

func (o AutocompleteTransactionID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutocompleteTransactionID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.TransactionGroupId) {
		toSerialize["transaction_group_id"] = o.TransactionGroupId
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableAutocompleteTransactionID struct {
	value *AutocompleteTransactionID
	isSet bool
}

func (v NullableAutocompleteTransactionID) Get() *AutocompleteTransactionID {
	return v.value
}

func (v *NullableAutocompleteTransactionID) Set(val *AutocompleteTransactionID) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteTransactionID) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteTransactionID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteTransactionID(val *AutocompleteTransactionID) *NullableAutocompleteTransactionID {
	return &NullableAutocompleteTransactionID{value: val, isSet: true}
}

func (v NullableAutocompleteTransactionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteTransactionID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


