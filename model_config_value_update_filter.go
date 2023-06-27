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

// ConfigValueUpdateFilter the model 'ConfigValueUpdateFilter'
type ConfigValueUpdateFilter string

// List of ConfigValueUpdateFilter
const (
	IS_DEMO_SITE ConfigValueUpdateFilter = "configuration.is_demo_site"
	PERMISSION_UPDATE_CHECK ConfigValueUpdateFilter = "configuration.permission_update_check"
	LAST_UPDATE_CHECK ConfigValueUpdateFilter = "configuration.last_update_check"
	SINGLE_USER_MODE ConfigValueUpdateFilter = "configuration.single_user_mode"
)

// All allowed values of ConfigValueUpdateFilter enum
var AllowedConfigValueUpdateFilterEnumValues = []ConfigValueUpdateFilter{
	"configuration.is_demo_site",
	"configuration.permission_update_check",
	"configuration.last_update_check",
	"configuration.single_user_mode",
}

func (v *ConfigValueUpdateFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConfigValueUpdateFilter(value)
	for _, existing := range AllowedConfigValueUpdateFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConfigValueUpdateFilter", value)
}

// NewConfigValueUpdateFilterFromValue returns a pointer to a valid ConfigValueUpdateFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConfigValueUpdateFilterFromValue(v string) (*ConfigValueUpdateFilter, error) {
	ev := ConfigValueUpdateFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConfigValueUpdateFilter: valid values are %v", v, AllowedConfigValueUpdateFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConfigValueUpdateFilter) IsValid() bool {
	for _, existing := range AllowedConfigValueUpdateFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConfigValueUpdateFilter value
func (v ConfigValueUpdateFilter) Ptr() *ConfigValueUpdateFilter {
	return &v
}

type NullableConfigValueUpdateFilter struct {
	value *ConfigValueUpdateFilter
	isSet bool
}

func (v NullableConfigValueUpdateFilter) Get() *ConfigValueUpdateFilter {
	return v.value
}

func (v *NullableConfigValueUpdateFilter) Set(val *ConfigValueUpdateFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigValueUpdateFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigValueUpdateFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigValueUpdateFilter(val *ConfigValueUpdateFilter) *NullableConfigValueUpdateFilter {
	return &NullableConfigValueUpdateFilter{value: val, isSet: true}
}

func (v NullableConfigValueUpdateFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigValueUpdateFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

