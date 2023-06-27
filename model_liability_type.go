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

// LiabilityType Mandatory when type is liability. Specifies the exact type.
type LiabilityType string

// List of LiabilityType
const (
	LOAN LiabilityType = "loan"
	DEBT LiabilityType = "debt"
	MORTGAGE LiabilityType = "mortgage"
	NULL LiabilityType = "null"
)

// All allowed values of LiabilityType enum
var AllowedLiabilityTypeEnumValues = []LiabilityType{
	"loan",
	"debt",
	"mortgage",
	"null",
}

func (v *LiabilityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LiabilityType(value)
	for _, existing := range AllowedLiabilityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LiabilityType", value)
}

// NewLiabilityTypeFromValue returns a pointer to a valid LiabilityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLiabilityTypeFromValue(v string) (*LiabilityType, error) {
	ev := LiabilityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LiabilityType: valid values are %v", v, AllowedLiabilityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LiabilityType) IsValid() bool {
	for _, existing := range AllowedLiabilityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LiabilityType value
func (v LiabilityType) Ptr() *LiabilityType {
	return &v
}

type NullableLiabilityType struct {
	value *LiabilityType
	isSet bool
}

func (v NullableLiabilityType) Get() *LiabilityType {
	return v.value
}

func (v *NullableLiabilityType) Set(val *LiabilityType) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilityType) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilityType(val *LiabilityType) *NullableLiabilityType {
	return &NullableLiabilityType{value: val, isSet: true}
}

func (v NullableLiabilityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

