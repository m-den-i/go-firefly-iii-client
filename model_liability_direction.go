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

// LiabilityDirection 'credit' indicates somebody owes you the liability. 'debit' Indicates you owe this debt yourself. Works only for liabiltiies.
type LiabilityDirection string

// List of LiabilityDirection
const (
	CREDIT LiabilityDirection = "credit"
	DEBIT LiabilityDirection = "debit"
	NULL LiabilityDirection = "null"
)

// All allowed values of LiabilityDirection enum
var AllowedLiabilityDirectionEnumValues = []LiabilityDirection{
	"credit",
	"debit",
	"null",
}

func (v *LiabilityDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LiabilityDirection(value)
	for _, existing := range AllowedLiabilityDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LiabilityDirection", value)
}

// NewLiabilityDirectionFromValue returns a pointer to a valid LiabilityDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLiabilityDirectionFromValue(v string) (*LiabilityDirection, error) {
	ev := LiabilityDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LiabilityDirection: valid values are %v", v, AllowedLiabilityDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LiabilityDirection) IsValid() bool {
	for _, existing := range AllowedLiabilityDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LiabilityDirection value
func (v LiabilityDirection) Ptr() *LiabilityDirection {
	return &v
}

type NullableLiabilityDirection struct {
	value *LiabilityDirection
	isSet bool
}

func (v NullableLiabilityDirection) Get() *LiabilityDirection {
	return v.value
}

func (v *NullableLiabilityDirection) Set(val *LiabilityDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilityDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilityDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilityDirection(val *LiabilityDirection) *NullableLiabilityDirection {
	return &NullableLiabilityDirection{value: val, isSet: true}
}

func (v NullableLiabilityDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilityDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

