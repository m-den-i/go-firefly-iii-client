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

// checks if the RecurrenceRepetitionStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecurrenceRepetitionStore{}

// RecurrenceRepetitionStore struct for RecurrenceRepetitionStore
type RecurrenceRepetitionStore struct {
	Type RecurrenceRepetitionType `json:"type"`
	// Information that defined the type of repetition. - For 'daily', this is empty. - For 'weekly', it is day of the week between 1 and 7 (Monday - Sunday). - For 'ndom', it is '1,2' or '4,5' or something else, where the first number is the week in the month, and the second number is the day in the week (between 1 and 7). '2,3' means: the 2nd Wednesday of the month - For 'monthly' it is the day of the month (1 - 31) - For yearly, it is a full date, ie '2018-09-17'. The year you use does not matter. 
	Moment string `json:"moment"`
	// How many occurrences to skip. 0 means skip nothing. 1 means every other.
	Skip *int32 `json:"skip,omitempty"`
	// How to respond when the recurring transaction falls in the weekend. Possible values: 1. Do nothing, just create it 2. Create no transaction. 3. Skip to the previous Friday. 4. Skip to the next Monday. 
	Weekend *int32 `json:"weekend,omitempty"`
}

// NewRecurrenceRepetitionStore instantiates a new RecurrenceRepetitionStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurrenceRepetitionStore(type_ RecurrenceRepetitionType, moment string) *RecurrenceRepetitionStore {
	this := RecurrenceRepetitionStore{}
	this.Type = type_
	this.Moment = moment
	return &this
}

// NewRecurrenceRepetitionStoreWithDefaults instantiates a new RecurrenceRepetitionStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurrenceRepetitionStoreWithDefaults() *RecurrenceRepetitionStore {
	this := RecurrenceRepetitionStore{}
	return &this
}

// GetType returns the Type field value
func (o *RecurrenceRepetitionStore) GetType() RecurrenceRepetitionType {
	if o == nil {
		var ret RecurrenceRepetitionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecurrenceRepetitionStore) GetTypeOk() (*RecurrenceRepetitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RecurrenceRepetitionStore) SetType(v RecurrenceRepetitionType) {
	o.Type = v
}

// GetMoment returns the Moment field value
func (o *RecurrenceRepetitionStore) GetMoment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Moment
}

// GetMomentOk returns a tuple with the Moment field value
// and a boolean to check if the value has been set.
func (o *RecurrenceRepetitionStore) GetMomentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Moment, true
}

// SetMoment sets field value
func (o *RecurrenceRepetitionStore) SetMoment(v string) {
	o.Moment = v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *RecurrenceRepetitionStore) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceRepetitionStore) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *RecurrenceRepetitionStore) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *RecurrenceRepetitionStore) SetSkip(v int32) {
	o.Skip = &v
}

// GetWeekend returns the Weekend field value if set, zero value otherwise.
func (o *RecurrenceRepetitionStore) GetWeekend() int32 {
	if o == nil || IsNil(o.Weekend) {
		var ret int32
		return ret
	}
	return *o.Weekend
}

// GetWeekendOk returns a tuple with the Weekend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceRepetitionStore) GetWeekendOk() (*int32, bool) {
	if o == nil || IsNil(o.Weekend) {
		return nil, false
	}
	return o.Weekend, true
}

// HasWeekend returns a boolean if a field has been set.
func (o *RecurrenceRepetitionStore) HasWeekend() bool {
	if o != nil && !IsNil(o.Weekend) {
		return true
	}

	return false
}

// SetWeekend gets a reference to the given int32 and assigns it to the Weekend field.
func (o *RecurrenceRepetitionStore) SetWeekend(v int32) {
	o.Weekend = &v
}

func (o RecurrenceRepetitionStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecurrenceRepetitionStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["moment"] = o.Moment
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Weekend) {
		toSerialize["weekend"] = o.Weekend
	}
	return toSerialize, nil
}

type NullableRecurrenceRepetitionStore struct {
	value *RecurrenceRepetitionStore
	isSet bool
}

func (v NullableRecurrenceRepetitionStore) Get() *RecurrenceRepetitionStore {
	return v.value
}

func (v *NullableRecurrenceRepetitionStore) Set(val *RecurrenceRepetitionStore) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurrenceRepetitionStore) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurrenceRepetitionStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurrenceRepetitionStore(val *RecurrenceRepetitionStore) *NullableRecurrenceRepetitionStore {
	return &NullableRecurrenceRepetitionStore{value: val, isSet: true}
}

func (v NullableRecurrenceRepetitionStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurrenceRepetitionStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


