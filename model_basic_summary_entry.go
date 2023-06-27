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

// checks if the BasicSummaryEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicSummaryEntry{}

// BasicSummaryEntry struct for BasicSummaryEntry
type BasicSummaryEntry struct {
	// This is a reference to the type of info shared, not influenced by translations or user preferences. The EUR value is a reference to the currency code. Possibilities are: balance-in-ABC, spent-in-ABC, earned-in-ABC, bills-paid-in-ABC, bills-unpaid-in-ABC, left-to-spend-in-ABC and net-worth-in-ABC.
	Key *string `json:"key,omitempty"`
	// A translated title for the information shared.
	Title *string `json:"title,omitempty"`
	// The amount as a float.
	MonetaryValue *float64 `json:"monetary_value,omitempty"`
	// The currency ID of the associated currency.
	CurrencyId *string `json:"currency_id,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	// Number of decimals for the associated currency.
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
	// The amount formatted according to the users locale
	ValueParsed *string `json:"value_parsed,omitempty"`
	// Reference to a font-awesome icon without the fa- part.
	LocalIcon *string `json:"local_icon,omitempty"`
	// A short explanation of the amounts origin. Already formatted according to the locale of the user or translated, if relevant.
	SubTitle *string `json:"sub_title,omitempty"`
}

// NewBasicSummaryEntry instantiates a new BasicSummaryEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicSummaryEntry() *BasicSummaryEntry {
	this := BasicSummaryEntry{}
	return &this
}

// NewBasicSummaryEntryWithDefaults instantiates a new BasicSummaryEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicSummaryEntryWithDefaults() *BasicSummaryEntry {
	this := BasicSummaryEntry{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *BasicSummaryEntry) SetKey(v string) {
	o.Key = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BasicSummaryEntry) SetTitle(v string) {
	o.Title = &v
}

// GetMonetaryValue returns the MonetaryValue field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetMonetaryValue() float64 {
	if o == nil || IsNil(o.MonetaryValue) {
		var ret float64
		return ret
	}
	return *o.MonetaryValue
}

// GetMonetaryValueOk returns a tuple with the MonetaryValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetMonetaryValueOk() (*float64, bool) {
	if o == nil || IsNil(o.MonetaryValue) {
		return nil, false
	}
	return o.MonetaryValue, true
}

// HasMonetaryValue returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasMonetaryValue() bool {
	if o != nil && !IsNil(o.MonetaryValue) {
		return true
	}

	return false
}

// SetMonetaryValue gets a reference to the given float64 and assigns it to the MonetaryValue field.
func (o *BasicSummaryEntry) SetMonetaryValue(v float64) {
	o.MonetaryValue = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *BasicSummaryEntry) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BasicSummaryEntry) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencySymbol) {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasCurrencySymbol() bool {
	if o != nil && !IsNil(o.CurrencySymbol) {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *BasicSummaryEntry) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetCurrencyDecimalPlaces() int32 {
	if o == nil || IsNil(o.CurrencyDecimalPlaces) {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrencyDecimalPlaces) {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasCurrencyDecimalPlaces() bool {
	if o != nil && !IsNil(o.CurrencyDecimalPlaces) {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *BasicSummaryEntry) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

// GetValueParsed returns the ValueParsed field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetValueParsed() string {
	if o == nil || IsNil(o.ValueParsed) {
		var ret string
		return ret
	}
	return *o.ValueParsed
}

// GetValueParsedOk returns a tuple with the ValueParsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetValueParsedOk() (*string, bool) {
	if o == nil || IsNil(o.ValueParsed) {
		return nil, false
	}
	return o.ValueParsed, true
}

// HasValueParsed returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasValueParsed() bool {
	if o != nil && !IsNil(o.ValueParsed) {
		return true
	}

	return false
}

// SetValueParsed gets a reference to the given string and assigns it to the ValueParsed field.
func (o *BasicSummaryEntry) SetValueParsed(v string) {
	o.ValueParsed = &v
}

// GetLocalIcon returns the LocalIcon field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetLocalIcon() string {
	if o == nil || IsNil(o.LocalIcon) {
		var ret string
		return ret
	}
	return *o.LocalIcon
}

// GetLocalIconOk returns a tuple with the LocalIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetLocalIconOk() (*string, bool) {
	if o == nil || IsNil(o.LocalIcon) {
		return nil, false
	}
	return o.LocalIcon, true
}

// HasLocalIcon returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasLocalIcon() bool {
	if o != nil && !IsNil(o.LocalIcon) {
		return true
	}

	return false
}

// SetLocalIcon gets a reference to the given string and assigns it to the LocalIcon field.
func (o *BasicSummaryEntry) SetLocalIcon(v string) {
	o.LocalIcon = &v
}

// GetSubTitle returns the SubTitle field value if set, zero value otherwise.
func (o *BasicSummaryEntry) GetSubTitle() string {
	if o == nil || IsNil(o.SubTitle) {
		var ret string
		return ret
	}
	return *o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicSummaryEntry) GetSubTitleOk() (*string, bool) {
	if o == nil || IsNil(o.SubTitle) {
		return nil, false
	}
	return o.SubTitle, true
}

// HasSubTitle returns a boolean if a field has been set.
func (o *BasicSummaryEntry) HasSubTitle() bool {
	if o != nil && !IsNil(o.SubTitle) {
		return true
	}

	return false
}

// SetSubTitle gets a reference to the given string and assigns it to the SubTitle field.
func (o *BasicSummaryEntry) SetSubTitle(v string) {
	o.SubTitle = &v
}

func (o BasicSummaryEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicSummaryEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.MonetaryValue) {
		toSerialize["monetary_value"] = o.MonetaryValue
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
	if !IsNil(o.ValueParsed) {
		toSerialize["value_parsed"] = o.ValueParsed
	}
	if !IsNil(o.LocalIcon) {
		toSerialize["local_icon"] = o.LocalIcon
	}
	if !IsNil(o.SubTitle) {
		toSerialize["sub_title"] = o.SubTitle
	}
	return toSerialize, nil
}

type NullableBasicSummaryEntry struct {
	value *BasicSummaryEntry
	isSet bool
}

func (v NullableBasicSummaryEntry) Get() *BasicSummaryEntry {
	return v.value
}

func (v *NullableBasicSummaryEntry) Set(val *BasicSummaryEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicSummaryEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicSummaryEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicSummaryEntry(val *BasicSummaryEntry) *NullableBasicSummaryEntry {
	return &NullableBasicSummaryEntry{value: val, isSet: true}
}

func (v NullableBasicSummaryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicSummaryEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


