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

// checks if the CurrencyUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrencyUpdate{}

// CurrencyUpdate struct for CurrencyUpdate
type CurrencyUpdate struct {
	// If the currency is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// If the currency must be the default for the user. You can only submit TRUE.
	Default *bool `json:"default,omitempty"`
	// The currency code
	Code *string `json:"code,omitempty"`
	// The currency name
	Name *string `json:"name,omitempty"`
	// The currency symbol
	Symbol *string `json:"symbol,omitempty"`
	// How many decimals to use when displaying this currency. Between 0 and 16.
	DecimalPlaces *int32 `json:"decimal_places,omitempty"`
}

// NewCurrencyUpdate instantiates a new CurrencyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencyUpdate() *CurrencyUpdate {
	this := CurrencyUpdate{}
	return &this
}

// NewCurrencyUpdateWithDefaults instantiates a new CurrencyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyUpdateWithDefaults() *CurrencyUpdate {
	this := CurrencyUpdate{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CurrencyUpdate) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CurrencyUpdate) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CurrencyUpdate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CurrencyUpdate) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyUpdate) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CurrencyUpdate) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CurrencyUpdate) SetDefault(v bool) {
	o.Default = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CurrencyUpdate) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyUpdate) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CurrencyUpdate) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CurrencyUpdate) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CurrencyUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CurrencyUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CurrencyUpdate) SetName(v string) {
	o.Name = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CurrencyUpdate) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyUpdate) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CurrencyUpdate) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CurrencyUpdate) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDecimalPlaces returns the DecimalPlaces field value if set, zero value otherwise.
func (o *CurrencyUpdate) GetDecimalPlaces() int32 {
	if o == nil || IsNil(o.DecimalPlaces) {
		var ret int32
		return ret
	}
	return *o.DecimalPlaces
}

// GetDecimalPlacesOk returns a tuple with the DecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyUpdate) GetDecimalPlacesOk() (*int32, bool) {
	if o == nil || IsNil(o.DecimalPlaces) {
		return nil, false
	}
	return o.DecimalPlaces, true
}

// HasDecimalPlaces returns a boolean if a field has been set.
func (o *CurrencyUpdate) HasDecimalPlaces() bool {
	if o != nil && !IsNil(o.DecimalPlaces) {
		return true
	}

	return false
}

// SetDecimalPlaces gets a reference to the given int32 and assigns it to the DecimalPlaces field.
func (o *CurrencyUpdate) SetDecimalPlaces(v int32) {
	o.DecimalPlaces = &v
}

func (o CurrencyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrencyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.DecimalPlaces) {
		toSerialize["decimal_places"] = o.DecimalPlaces
	}
	return toSerialize, nil
}

type NullableCurrencyUpdate struct {
	value *CurrencyUpdate
	isSet bool
}

func (v NullableCurrencyUpdate) Get() *CurrencyUpdate {
	return v.value
}

func (v *NullableCurrencyUpdate) Set(val *CurrencyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyUpdate(val *CurrencyUpdate) *NullableCurrencyUpdate {
	return &NullableCurrencyUpdate{value: val, isSet: true}
}

func (v NullableCurrencyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


