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

// PolymorphicProperty - struct for PolymorphicProperty
type PolymorphicProperty struct {
	ArrayOfString *[]string
	Bool *bool
	MapmapOfStringinterface{} *map[string]interface{}
	String *string
}

// []stringAsPolymorphicProperty is a convenience function that returns []string wrapped in PolymorphicProperty
func ArrayOfStringAsPolymorphicProperty(v *[]string) PolymorphicProperty {
	return PolymorphicProperty{
		ArrayOfString: v,
	}
}

// boolAsPolymorphicProperty is a convenience function that returns bool wrapped in PolymorphicProperty
func BoolAsPolymorphicProperty(v *bool) PolymorphicProperty {
	return PolymorphicProperty{
		Bool: v,
	}
}

// map[string]interface{}AsPolymorphicProperty is a convenience function that returns map[string]interface{} wrapped in PolymorphicProperty
func MapmapOfStringinterface{}AsPolymorphicProperty(v *map[string]interface{}) PolymorphicProperty {
	return PolymorphicProperty{
		MapmapOfStringinterface{}: v,
	}
}

// stringAsPolymorphicProperty is a convenience function that returns string wrapped in PolymorphicProperty
func StringAsPolymorphicProperty(v *string) PolymorphicProperty {
	return PolymorphicProperty{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PolymorphicProperty) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into MapmapOfStringinterface{}
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringinterface{})
	if err == nil {
		jsonMapmapOfStringinterface{}, _ := json.Marshal(dst.MapmapOfStringinterface{})
		if string(jsonMapmapOfStringinterface{}) == "{}" { // empty struct
			dst.MapmapOfStringinterface{} = nil
		} else {
			match++
		}
	} else {
		dst.MapmapOfStringinterface{} = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.Bool = nil
		dst.MapmapOfStringinterface{} = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PolymorphicProperty)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PolymorphicProperty)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PolymorphicProperty) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.MapmapOfStringinterface{} != nil {
		return json.Marshal(&src.MapmapOfStringinterface{})
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PolymorphicProperty) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.MapmapOfStringinterface{} != nil {
		return obj.MapmapOfStringinterface{}
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullablePolymorphicProperty struct {
	value *PolymorphicProperty
	isSet bool
}

func (v NullablePolymorphicProperty) Get() *PolymorphicProperty {
	return v.value
}

func (v *NullablePolymorphicProperty) Set(val *PolymorphicProperty) {
	v.value = val
	v.isSet = true
}

func (v NullablePolymorphicProperty) IsSet() bool {
	return v.isSet
}

func (v *NullablePolymorphicProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolymorphicProperty(val *PolymorphicProperty) *NullablePolymorphicProperty {
	return &NullablePolymorphicProperty{value: val, isSet: true}
}

func (v NullablePolymorphicProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolymorphicProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

