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

// checks if the BillStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillStore{}

// BillStore struct for BillStore
type BillStore struct {
	// Use either currency_id or currency_code
	CurrencyId *string `json:"currency_id,omitempty"`
	// Use either currency_id or currency_code
	CurrencyCode *string `json:"currency_code,omitempty"`
	Name string `json:"name"`
	AmountMin string `json:"amount_min"`
	AmountMax string `json:"amount_max"`
	Date time.Time `json:"date"`
	// The date after which this bill is no longer valid or applicable
	EndDate *time.Time `json:"end_date,omitempty"`
	// The date before which the bill must be renewed (or cancelled)
	ExtensionDate *time.Time `json:"extension_date,omitempty"`
	RepeatFreq BillRepeatFrequency `json:"repeat_freq"`
	// How often the bill must be skipped. 1 means a bi-monthly bill.
	Skip *int32 `json:"skip,omitempty"`
	// If the bill is active.
	Active *bool `json:"active,omitempty"`
	Notes NullableString `json:"notes,omitempty"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupId NullableString `json:"object_group_id,omitempty"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle NullableString `json:"object_group_title,omitempty"`
}

// NewBillStore instantiates a new BillStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillStore(name string, amountMin string, amountMax string, date time.Time, repeatFreq BillRepeatFrequency) *BillStore {
	this := BillStore{}
	this.Name = name
	this.AmountMin = amountMin
	this.AmountMax = amountMax
	this.Date = date
	this.RepeatFreq = repeatFreq
	return &this
}

// NewBillStoreWithDefaults instantiates a new BillStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillStoreWithDefaults() *BillStore {
	this := BillStore{}
	return &this
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *BillStore) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillStore) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *BillStore) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *BillStore) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BillStore) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillStore) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BillStore) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BillStore) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetName returns the Name field value
func (o *BillStore) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BillStore) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BillStore) SetName(v string) {
	o.Name = v
}

// GetAmountMin returns the AmountMin field value
func (o *BillStore) GetAmountMin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountMin
}

// GetAmountMinOk returns a tuple with the AmountMin field value
// and a boolean to check if the value has been set.
func (o *BillStore) GetAmountMinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMin, true
}

// SetAmountMin sets field value
func (o *BillStore) SetAmountMin(v string) {
	o.AmountMin = v
}

// GetAmountMax returns the AmountMax field value
func (o *BillStore) GetAmountMax() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountMax
}

// GetAmountMaxOk returns a tuple with the AmountMax field value
// and a boolean to check if the value has been set.
func (o *BillStore) GetAmountMaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMax, true
}

// SetAmountMax sets field value
func (o *BillStore) SetAmountMax(v string) {
	o.AmountMax = v
}

// GetDate returns the Date field value
func (o *BillStore) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *BillStore) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *BillStore) SetDate(v time.Time) {
	o.Date = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillStore) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillStore) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillStore) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillStore) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetExtensionDate returns the ExtensionDate field value if set, zero value otherwise.
func (o *BillStore) GetExtensionDate() time.Time {
	if o == nil || IsNil(o.ExtensionDate) {
		var ret time.Time
		return ret
	}
	return *o.ExtensionDate
}

// GetExtensionDateOk returns a tuple with the ExtensionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillStore) GetExtensionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExtensionDate) {
		return nil, false
	}
	return o.ExtensionDate, true
}

// HasExtensionDate returns a boolean if a field has been set.
func (o *BillStore) HasExtensionDate() bool {
	if o != nil && !IsNil(o.ExtensionDate) {
		return true
	}

	return false
}

// SetExtensionDate gets a reference to the given time.Time and assigns it to the ExtensionDate field.
func (o *BillStore) SetExtensionDate(v time.Time) {
	o.ExtensionDate = &v
}

// GetRepeatFreq returns the RepeatFreq field value
func (o *BillStore) GetRepeatFreq() BillRepeatFrequency {
	if o == nil {
		var ret BillRepeatFrequency
		return ret
	}

	return o.RepeatFreq
}

// GetRepeatFreqOk returns a tuple with the RepeatFreq field value
// and a boolean to check if the value has been set.
func (o *BillStore) GetRepeatFreqOk() (*BillRepeatFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepeatFreq, true
}

// SetRepeatFreq sets field value
func (o *BillStore) SetRepeatFreq(v BillRepeatFrequency) {
	o.RepeatFreq = v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *BillStore) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillStore) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *BillStore) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *BillStore) SetSkip(v int32) {
	o.Skip = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *BillStore) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillStore) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *BillStore) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *BillStore) SetActive(v bool) {
	o.Active = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillStore) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillStore) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *BillStore) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *BillStore) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *BillStore) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *BillStore) UnsetNotes() {
	o.Notes.Unset()
}

// GetObjectGroupId returns the ObjectGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillStore) GetObjectGroupId() string {
	if o == nil || IsNil(o.ObjectGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.ObjectGroupId.Get()
}

// GetObjectGroupIdOk returns a tuple with the ObjectGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillStore) GetObjectGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectGroupId.Get(), o.ObjectGroupId.IsSet()
}

// HasObjectGroupId returns a boolean if a field has been set.
func (o *BillStore) HasObjectGroupId() bool {
	if o != nil && o.ObjectGroupId.IsSet() {
		return true
	}

	return false
}

// SetObjectGroupId gets a reference to the given NullableString and assigns it to the ObjectGroupId field.
func (o *BillStore) SetObjectGroupId(v string) {
	o.ObjectGroupId.Set(&v)
}
// SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil
func (o *BillStore) SetObjectGroupIdNil() {
	o.ObjectGroupId.Set(nil)
}

// UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
func (o *BillStore) UnsetObjectGroupId() {
	o.ObjectGroupId.Unset()
}

// GetObjectGroupTitle returns the ObjectGroupTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillStore) GetObjectGroupTitle() string {
	if o == nil || IsNil(o.ObjectGroupTitle.Get()) {
		var ret string
		return ret
	}
	return *o.ObjectGroupTitle.Get()
}

// GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillStore) GetObjectGroupTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectGroupTitle.Get(), o.ObjectGroupTitle.IsSet()
}

// HasObjectGroupTitle returns a boolean if a field has been set.
func (o *BillStore) HasObjectGroupTitle() bool {
	if o != nil && o.ObjectGroupTitle.IsSet() {
		return true
	}

	return false
}

// SetObjectGroupTitle gets a reference to the given NullableString and assigns it to the ObjectGroupTitle field.
func (o *BillStore) SetObjectGroupTitle(v string) {
	o.ObjectGroupTitle.Set(&v)
}
// SetObjectGroupTitleNil sets the value for ObjectGroupTitle to be an explicit nil
func (o *BillStore) SetObjectGroupTitleNil() {
	o.ObjectGroupTitle.Set(nil)
}

// UnsetObjectGroupTitle ensures that no value is present for ObjectGroupTitle, not even an explicit nil
func (o *BillStore) UnsetObjectGroupTitle() {
	o.ObjectGroupTitle.Unset()
}

func (o BillStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	toSerialize["name"] = o.Name
	toSerialize["amount_min"] = o.AmountMin
	toSerialize["amount_max"] = o.AmountMax
	toSerialize["date"] = o.Date
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.ExtensionDate) {
		toSerialize["extension_date"] = o.ExtensionDate
	}
	toSerialize["repeat_freq"] = o.RepeatFreq
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.ObjectGroupId.IsSet() {
		toSerialize["object_group_id"] = o.ObjectGroupId.Get()
	}
	if o.ObjectGroupTitle.IsSet() {
		toSerialize["object_group_title"] = o.ObjectGroupTitle.Get()
	}
	return toSerialize, nil
}

type NullableBillStore struct {
	value *BillStore
	isSet bool
}

func (v NullableBillStore) Get() *BillStore {
	return v.value
}

func (v *NullableBillStore) Set(val *BillStore) {
	v.value = val
	v.isSet = true
}

func (v NullableBillStore) IsSet() bool {
	return v.isSet
}

func (v *NullableBillStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillStore(val *BillStore) *NullableBillStore {
	return &NullableBillStore{value: val, isSet: true}
}

func (v NullableBillStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


