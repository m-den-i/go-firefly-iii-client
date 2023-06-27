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

// checks if the PiggyBankEventRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PiggyBankEventRead{}

// PiggyBankEventRead struct for PiggyBankEventRead
type PiggyBankEventRead struct {
	// Immutable value
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes PiggyBankEvent `json:"attributes"`
	Links ObjectLink `json:"links"`
}

// NewPiggyBankEventRead instantiates a new PiggyBankEventRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPiggyBankEventRead(type_ string, id string, attributes PiggyBankEvent, links ObjectLink) *PiggyBankEventRead {
	this := PiggyBankEventRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewPiggyBankEventReadWithDefaults instantiates a new PiggyBankEventRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPiggyBankEventReadWithDefaults() *PiggyBankEventRead {
	this := PiggyBankEventRead{}
	return &this
}

// GetType returns the Type field value
func (o *PiggyBankEventRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PiggyBankEventRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PiggyBankEventRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PiggyBankEventRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PiggyBankEventRead) GetAttributes() PiggyBankEvent {
	if o == nil {
		var ret PiggyBankEvent
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventRead) GetAttributesOk() (*PiggyBankEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PiggyBankEventRead) SetAttributes(v PiggyBankEvent) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *PiggyBankEventRead) GetLinks() ObjectLink {
	if o == nil {
		var ret ObjectLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventRead) GetLinksOk() (*ObjectLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PiggyBankEventRead) SetLinks(v ObjectLink) {
	o.Links = v
}

func (o PiggyBankEventRead) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PiggyBankEventRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullablePiggyBankEventRead struct {
	value *PiggyBankEventRead
	isSet bool
}

func (v NullablePiggyBankEventRead) Get() *PiggyBankEventRead {
	return v.value
}

func (v *NullablePiggyBankEventRead) Set(val *PiggyBankEventRead) {
	v.value = val
	v.isSet = true
}

func (v NullablePiggyBankEventRead) IsSet() bool {
	return v.isSet
}

func (v *NullablePiggyBankEventRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePiggyBankEventRead(val *PiggyBankEventRead) *NullablePiggyBankEventRead {
	return &NullablePiggyBankEventRead{value: val, isSet: true}
}

func (v NullablePiggyBankEventRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePiggyBankEventRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

