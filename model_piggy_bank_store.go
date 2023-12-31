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

// checks if the PiggyBankStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PiggyBankStore{}

// PiggyBankStore struct for PiggyBankStore
type PiggyBankStore struct {
	Name string `json:"name"`
	// The ID of the asset account this piggy bank is connected to.
	AccountId string `json:"account_id"`
	TargetAmount NullableString `json:"target_amount"`
	CurrentAmount *string `json:"current_amount,omitempty"`
	// The date you started with this piggy bank.
	StartDate *string `json:"start_date,omitempty"`
	// The date you intend to finish saving money.
	TargetDate NullableString `json:"target_date,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Active *bool `json:"active,omitempty"`
	Notes NullableString `json:"notes,omitempty"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupId NullableString `json:"object_group_id,omitempty"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle NullableString `json:"object_group_title,omitempty"`
}

// NewPiggyBankStore instantiates a new PiggyBankStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPiggyBankStore(name string, accountId string, targetAmount NullableString) *PiggyBankStore {
	this := PiggyBankStore{}
	this.Name = name
	this.AccountId = accountId
	this.TargetAmount = targetAmount
	return &this
}

// NewPiggyBankStoreWithDefaults instantiates a new PiggyBankStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPiggyBankStoreWithDefaults() *PiggyBankStore {
	this := PiggyBankStore{}
	return &this
}

// GetName returns the Name field value
func (o *PiggyBankStore) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PiggyBankStore) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PiggyBankStore) SetName(v string) {
	o.Name = v
}

// GetAccountId returns the AccountId field value
func (o *PiggyBankStore) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *PiggyBankStore) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *PiggyBankStore) SetAccountId(v string) {
	o.AccountId = v
}

// GetTargetAmount returns the TargetAmount field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PiggyBankStore) GetTargetAmount() string {
	if o == nil || o.TargetAmount.Get() == nil {
		var ret string
		return ret
	}

	return *o.TargetAmount.Get()
}

// GetTargetAmountOk returns a tuple with the TargetAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBankStore) GetTargetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetAmount.Get(), o.TargetAmount.IsSet()
}

// SetTargetAmount sets field value
func (o *PiggyBankStore) SetTargetAmount(v string) {
	o.TargetAmount.Set(&v)
}

// GetCurrentAmount returns the CurrentAmount field value if set, zero value otherwise.
func (o *PiggyBankStore) GetCurrentAmount() string {
	if o == nil || IsNil(o.CurrentAmount) {
		var ret string
		return ret
	}
	return *o.CurrentAmount
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankStore) GetCurrentAmountOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentAmount) {
		return nil, false
	}
	return o.CurrentAmount, true
}

// HasCurrentAmount returns a boolean if a field has been set.
func (o *PiggyBankStore) HasCurrentAmount() bool {
	if o != nil && !IsNil(o.CurrentAmount) {
		return true
	}

	return false
}

// SetCurrentAmount gets a reference to the given string and assigns it to the CurrentAmount field.
func (o *PiggyBankStore) SetCurrentAmount(v string) {
	o.CurrentAmount = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PiggyBankStore) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankStore) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PiggyBankStore) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *PiggyBankStore) SetStartDate(v string) {
	o.StartDate = &v
}

// GetTargetDate returns the TargetDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PiggyBankStore) GetTargetDate() string {
	if o == nil || IsNil(o.TargetDate.Get()) {
		var ret string
		return ret
	}
	return *o.TargetDate.Get()
}

// GetTargetDateOk returns a tuple with the TargetDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBankStore) GetTargetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetDate.Get(), o.TargetDate.IsSet()
}

// HasTargetDate returns a boolean if a field has been set.
func (o *PiggyBankStore) HasTargetDate() bool {
	if o != nil && o.TargetDate.IsSet() {
		return true
	}

	return false
}

// SetTargetDate gets a reference to the given NullableString and assigns it to the TargetDate field.
func (o *PiggyBankStore) SetTargetDate(v string) {
	o.TargetDate.Set(&v)
}
// SetTargetDateNil sets the value for TargetDate to be an explicit nil
func (o *PiggyBankStore) SetTargetDateNil() {
	o.TargetDate.Set(nil)
}

// UnsetTargetDate ensures that no value is present for TargetDate, not even an explicit nil
func (o *PiggyBankStore) UnsetTargetDate() {
	o.TargetDate.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PiggyBankStore) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankStore) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PiggyBankStore) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *PiggyBankStore) SetOrder(v int32) {
	o.Order = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PiggyBankStore) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankStore) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PiggyBankStore) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PiggyBankStore) SetActive(v bool) {
	o.Active = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PiggyBankStore) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBankStore) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *PiggyBankStore) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *PiggyBankStore) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *PiggyBankStore) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *PiggyBankStore) UnsetNotes() {
	o.Notes.Unset()
}

// GetObjectGroupId returns the ObjectGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PiggyBankStore) GetObjectGroupId() string {
	if o == nil || IsNil(o.ObjectGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.ObjectGroupId.Get()
}

// GetObjectGroupIdOk returns a tuple with the ObjectGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBankStore) GetObjectGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectGroupId.Get(), o.ObjectGroupId.IsSet()
}

// HasObjectGroupId returns a boolean if a field has been set.
func (o *PiggyBankStore) HasObjectGroupId() bool {
	if o != nil && o.ObjectGroupId.IsSet() {
		return true
	}

	return false
}

// SetObjectGroupId gets a reference to the given NullableString and assigns it to the ObjectGroupId field.
func (o *PiggyBankStore) SetObjectGroupId(v string) {
	o.ObjectGroupId.Set(&v)
}
// SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil
func (o *PiggyBankStore) SetObjectGroupIdNil() {
	o.ObjectGroupId.Set(nil)
}

// UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
func (o *PiggyBankStore) UnsetObjectGroupId() {
	o.ObjectGroupId.Unset()
}

// GetObjectGroupTitle returns the ObjectGroupTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PiggyBankStore) GetObjectGroupTitle() string {
	if o == nil || IsNil(o.ObjectGroupTitle.Get()) {
		var ret string
		return ret
	}
	return *o.ObjectGroupTitle.Get()
}

// GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBankStore) GetObjectGroupTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectGroupTitle.Get(), o.ObjectGroupTitle.IsSet()
}

// HasObjectGroupTitle returns a boolean if a field has been set.
func (o *PiggyBankStore) HasObjectGroupTitle() bool {
	if o != nil && o.ObjectGroupTitle.IsSet() {
		return true
	}

	return false
}

// SetObjectGroupTitle gets a reference to the given NullableString and assigns it to the ObjectGroupTitle field.
func (o *PiggyBankStore) SetObjectGroupTitle(v string) {
	o.ObjectGroupTitle.Set(&v)
}
// SetObjectGroupTitleNil sets the value for ObjectGroupTitle to be an explicit nil
func (o *PiggyBankStore) SetObjectGroupTitleNil() {
	o.ObjectGroupTitle.Set(nil)
}

// UnsetObjectGroupTitle ensures that no value is present for ObjectGroupTitle, not even an explicit nil
func (o *PiggyBankStore) UnsetObjectGroupTitle() {
	o.ObjectGroupTitle.Unset()
}

func (o PiggyBankStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PiggyBankStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["account_id"] = o.AccountId
	toSerialize["target_amount"] = o.TargetAmount.Get()
	if !IsNil(o.CurrentAmount) {
		toSerialize["current_amount"] = o.CurrentAmount
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if o.TargetDate.IsSet() {
		toSerialize["target_date"] = o.TargetDate.Get()
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	// skip: active is readOnly
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

type NullablePiggyBankStore struct {
	value *PiggyBankStore
	isSet bool
}

func (v NullablePiggyBankStore) Get() *PiggyBankStore {
	return v.value
}

func (v *NullablePiggyBankStore) Set(val *PiggyBankStore) {
	v.value = val
	v.isSet = true
}

func (v NullablePiggyBankStore) IsSet() bool {
	return v.isSet
}

func (v *NullablePiggyBankStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePiggyBankStore(val *PiggyBankStore) *NullablePiggyBankStore {
	return &NullablePiggyBankStore{value: val, isSet: true}
}

func (v NullablePiggyBankStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePiggyBankStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


