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

// RuleTriggerKeyword The type of thing this trigger responds to. A limited set is possible
type RuleTriggerKeyword string

// List of RuleTriggerKeyword
const (
	FROM_ACCOUNT_STARTS RuleTriggerKeyword = "from_account_starts"
	FROM_ACCOUNT_ENDS RuleTriggerKeyword = "from_account_ends"
	FROM_ACCOUNT_IS RuleTriggerKeyword = "from_account_is"
	FROM_ACCOUNT_CONTAINS RuleTriggerKeyword = "from_account_contains"
	TO_ACCOUNT_STARTS RuleTriggerKeyword = "to_account_starts"
	TO_ACCOUNT_ENDS RuleTriggerKeyword = "to_account_ends"
	TO_ACCOUNT_IS RuleTriggerKeyword = "to_account_is"
	TO_ACCOUNT_CONTAINS RuleTriggerKeyword = "to_account_contains"
	AMOUNT_LESS RuleTriggerKeyword = "amount_less"
	AMOUNT_EXACTLY RuleTriggerKeyword = "amount_exactly"
	AMOUNT_MORE RuleTriggerKeyword = "amount_more"
	DESCRIPTION_STARTS RuleTriggerKeyword = "description_starts"
	DESCRIPTION_ENDS RuleTriggerKeyword = "description_ends"
	DESCRIPTION_CONTAINS RuleTriggerKeyword = "description_contains"
	DESCRIPTION_IS RuleTriggerKeyword = "description_is"
	TRANSACTION_TYPE RuleTriggerKeyword = "transaction_type"
	CATEGORY_IS RuleTriggerKeyword = "category_is"
	BUDGET_IS RuleTriggerKeyword = "budget_is"
	TAG_IS RuleTriggerKeyword = "tag_is"
	CURRENCY_IS RuleTriggerKeyword = "currency_is"
	HAS_ATTACHMENTS RuleTriggerKeyword = "has_attachments"
	HAS_NO_CATEGORY RuleTriggerKeyword = "has_no_category"
	HAS_ANY_CATEGORY RuleTriggerKeyword = "has_any_category"
	HAS_NO_BUDGET RuleTriggerKeyword = "has_no_budget"
	HAS_ANY_BUDGET RuleTriggerKeyword = "has_any_budget"
	HAS_NO_TAG RuleTriggerKeyword = "has_no_tag"
	HAS_ANY_TAG RuleTriggerKeyword = "has_any_tag"
	NOTES_CONTAINS RuleTriggerKeyword = "notes_contains"
	NOTES_START RuleTriggerKeyword = "notes_start"
	NOTES_END RuleTriggerKeyword = "notes_end"
	NOTES_ARE RuleTriggerKeyword = "notes_are"
	NO_NOTES RuleTriggerKeyword = "no_notes"
	ANY_NOTES RuleTriggerKeyword = "any_notes"
	SOURCE_ACCOUNT_IS RuleTriggerKeyword = "source_account_is"
	DESTINATION_ACCOUNT_IS RuleTriggerKeyword = "destination_account_is"
	SOURCE_ACCOUNT_STARTS RuleTriggerKeyword = "source_account_starts"
)

// All allowed values of RuleTriggerKeyword enum
var AllowedRuleTriggerKeywordEnumValues = []RuleTriggerKeyword{
	"from_account_starts",
	"from_account_ends",
	"from_account_is",
	"from_account_contains",
	"to_account_starts",
	"to_account_ends",
	"to_account_is",
	"to_account_contains",
	"amount_less",
	"amount_exactly",
	"amount_more",
	"description_starts",
	"description_ends",
	"description_contains",
	"description_is",
	"transaction_type",
	"category_is",
	"budget_is",
	"tag_is",
	"currency_is",
	"has_attachments",
	"has_no_category",
	"has_any_category",
	"has_no_budget",
	"has_any_budget",
	"has_no_tag",
	"has_any_tag",
	"notes_contains",
	"notes_start",
	"notes_end",
	"notes_are",
	"no_notes",
	"any_notes",
	"source_account_is",
	"destination_account_is",
	"source_account_starts",
}

func (v *RuleTriggerKeyword) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RuleTriggerKeyword(value)
	for _, existing := range AllowedRuleTriggerKeywordEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RuleTriggerKeyword", value)
}

// NewRuleTriggerKeywordFromValue returns a pointer to a valid RuleTriggerKeyword
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRuleTriggerKeywordFromValue(v string) (*RuleTriggerKeyword, error) {
	ev := RuleTriggerKeyword(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RuleTriggerKeyword: valid values are %v", v, AllowedRuleTriggerKeywordEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RuleTriggerKeyword) IsValid() bool {
	for _, existing := range AllowedRuleTriggerKeywordEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleTriggerKeyword value
func (v RuleTriggerKeyword) Ptr() *RuleTriggerKeyword {
	return &v
}

type NullableRuleTriggerKeyword struct {
	value *RuleTriggerKeyword
	isSet bool
}

func (v NullableRuleTriggerKeyword) Get() *RuleTriggerKeyword {
	return v.value
}

func (v *NullableRuleTriggerKeyword) Set(val *RuleTriggerKeyword) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleTriggerKeyword) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleTriggerKeyword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleTriggerKeyword(val *RuleTriggerKeyword) *NullableRuleTriggerKeyword {
	return &NullableRuleTriggerKeyword{value: val, isSet: true}
}

func (v NullableRuleTriggerKeyword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleTriggerKeyword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

