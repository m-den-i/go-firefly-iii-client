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
	"fmt"
)

// AccountTypeProperty the model 'AccountTypeProperty'
type AccountTypeProperty string

// List of AccountTypeProperty
const (
	DEFAULT_ACCOUNT AccountTypeProperty = "Default account"
	CASH_ACCOUNT AccountTypeProperty = "Cash account"
	ASSET_ACCOUNT AccountTypeProperty = "Asset account"
	EXPENSE_ACCOUNT AccountTypeProperty = "Expense account"
	REVENUE_ACCOUNT AccountTypeProperty = "Revenue account"
	INITIAL_BALANCE_ACCOUNT AccountTypeProperty = "Initial balance account"
	BENEFICIARY_ACCOUNT AccountTypeProperty = "Beneficiary account"
	IMPORT_ACCOUNT AccountTypeProperty = "Import account"
	RECONCILIATION_ACCOUNT AccountTypeProperty = "Reconciliation account"
	LOAN AccountTypeProperty = "Loan"
	DEBT AccountTypeProperty = "Debt"
	MORTGAGE AccountTypeProperty = "Mortgage"
)

// All allowed values of AccountTypeProperty enum
var AllowedAccountTypePropertyEnumValues = []AccountTypeProperty{
	"Default account",
	"Cash account",
	"Asset account",
	"Expense account",
	"Revenue account",
	"Initial balance account",
	"Beneficiary account",
	"Import account",
	"Reconciliation account",
	"Loan",
	"Debt",
	"Mortgage",
}

func (v *AccountTypeProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountTypeProperty(value)
	for _, existing := range AllowedAccountTypePropertyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountTypeProperty", value)
}

// NewAccountTypePropertyFromValue returns a pointer to a valid AccountTypeProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountTypePropertyFromValue(v string) (*AccountTypeProperty, error) {
	ev := AccountTypeProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountTypeProperty: valid values are %v", v, AllowedAccountTypePropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountTypeProperty) IsValid() bool {
	for _, existing := range AllowedAccountTypePropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountTypeProperty value
func (v AccountTypeProperty) Ptr() *AccountTypeProperty {
	return &v
}

type NullableAccountTypeProperty struct {
	value *AccountTypeProperty
	isSet bool
}

func (v NullableAccountTypeProperty) Get() *AccountTypeProperty {
	return v.value
}

func (v *NullableAccountTypeProperty) Set(val *AccountTypeProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTypeProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTypeProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTypeProperty(val *AccountTypeProperty) *NullableAccountTypeProperty {
	return &NullableAccountTypeProperty{value: val, isSet: true}
}

func (v NullableAccountTypeProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTypeProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

