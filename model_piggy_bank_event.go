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

// checks if the PiggyBankEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PiggyBankEvent{}

// PiggyBankEvent struct for PiggyBankEvent
type PiggyBankEvent struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CurrencyId *string `json:"currency_id,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
	Amount *string `json:"amount,omitempty"`
	// The journal associated with the event.
	TransactionJournalId *string `json:"transaction_journal_id,omitempty"`
	// The transaction group associated with the event.
	TransactionGroupId *string `json:"transaction_group_id,omitempty"`
}

// NewPiggyBankEvent instantiates a new PiggyBankEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPiggyBankEvent() *PiggyBankEvent {
	this := PiggyBankEvent{}
	return &this
}

// NewPiggyBankEventWithDefaults instantiates a new PiggyBankEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPiggyBankEventWithDefaults() *PiggyBankEvent {
	this := PiggyBankEvent{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PiggyBankEvent) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PiggyBankEvent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PiggyBankEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PiggyBankEvent) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankEvent) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PiggyBankEvent) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PiggyBankEvent) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *PiggyBankEvent) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankEvent) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *PiggyBankEvent) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *PiggyBankEvent) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PiggyBankEvent) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankEvent) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PiggyBankEvent) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PiggyBankEvent) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *PiggyBankEvent) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankEvent) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencySymbol) {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *PiggyBankEvent) HasCurrencySymbol() bool {
	if o != nil && !IsNil(o.CurrencySymbol) {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *PiggyBankEvent) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *PiggyBankEvent) GetCurrencyDecimalPlaces() int32 {
	if o == nil || IsNil(o.CurrencyDecimalPlaces) {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankEvent) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrencyDecimalPlaces) {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *PiggyBankEvent) HasCurrencyDecimalPlaces() bool {
	if o != nil && !IsNil(o.CurrencyDecimalPlaces) {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *PiggyBankEvent) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PiggyBankEvent) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankEvent) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PiggyBankEvent) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *PiggyBankEvent) SetAmount(v string) {
	o.Amount = &v
}

// GetTransactionJournalId returns the TransactionJournalId field value if set, zero value otherwise.
func (o *PiggyBankEvent) GetTransactionJournalId() string {
	if o == nil || IsNil(o.TransactionJournalId) {
		var ret string
		return ret
	}
	return *o.TransactionJournalId
}

// GetTransactionJournalIdOk returns a tuple with the TransactionJournalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankEvent) GetTransactionJournalIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionJournalId) {
		return nil, false
	}
	return o.TransactionJournalId, true
}

// HasTransactionJournalId returns a boolean if a field has been set.
func (o *PiggyBankEvent) HasTransactionJournalId() bool {
	if o != nil && !IsNil(o.TransactionJournalId) {
		return true
	}

	return false
}

// SetTransactionJournalId gets a reference to the given string and assigns it to the TransactionJournalId field.
func (o *PiggyBankEvent) SetTransactionJournalId(v string) {
	o.TransactionJournalId = &v
}

// GetTransactionGroupId returns the TransactionGroupId field value if set, zero value otherwise.
func (o *PiggyBankEvent) GetTransactionGroupId() string {
	if o == nil || IsNil(o.TransactionGroupId) {
		var ret string
		return ret
	}
	return *o.TransactionGroupId
}

// GetTransactionGroupIdOk returns a tuple with the TransactionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBankEvent) GetTransactionGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionGroupId) {
		return nil, false
	}
	return o.TransactionGroupId, true
}

// HasTransactionGroupId returns a boolean if a field has been set.
func (o *PiggyBankEvent) HasTransactionGroupId() bool {
	if o != nil && !IsNil(o.TransactionGroupId) {
		return true
	}

	return false
}

// SetTransactionGroupId gets a reference to the given string and assigns it to the TransactionGroupId field.
func (o *PiggyBankEvent) SetTransactionGroupId(v string) {
	o.TransactionGroupId = &v
}

func (o PiggyBankEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PiggyBankEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.CurrencySymbol) {
		toSerialize["currency_symbol"] = o.CurrencySymbol
	}
	if !IsNil(o.CurrencyDecimalPlaces) {
		toSerialize["currency_decimal_places"] = o.CurrencyDecimalPlaces
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.TransactionJournalId) {
		toSerialize["transaction_journal_id"] = o.TransactionJournalId
	}
	if !IsNil(o.TransactionGroupId) {
		toSerialize["transaction_group_id"] = o.TransactionGroupId
	}
	return toSerialize, nil
}

type NullablePiggyBankEvent struct {
	value *PiggyBankEvent
	isSet bool
}

func (v NullablePiggyBankEvent) Get() *PiggyBankEvent {
	return v.value
}

func (v *NullablePiggyBankEvent) Set(val *PiggyBankEvent) {
	v.value = val
	v.isSet = true
}

func (v NullablePiggyBankEvent) IsSet() bool {
	return v.isSet
}

func (v *NullablePiggyBankEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePiggyBankEvent(val *PiggyBankEvent) *NullablePiggyBankEvent {
	return &NullablePiggyBankEvent{value: val, isSet: true}
}

func (v NullablePiggyBankEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePiggyBankEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

