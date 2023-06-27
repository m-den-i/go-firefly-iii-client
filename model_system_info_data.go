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

// checks if the SystemInfoData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInfoData{}

// SystemInfoData struct for SystemInfoData
type SystemInfoData struct {
	Version *string `json:"version,omitempty"`
	ApiVersion *string `json:"api_version,omitempty"`
	PhpVersion *string `json:"php_version,omitempty"`
	Os *string `json:"os,omitempty"`
	Driver *string `json:"driver,omitempty"`
}

// NewSystemInfoData instantiates a new SystemInfoData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInfoData() *SystemInfoData {
	this := SystemInfoData{}
	return &this
}

// NewSystemInfoDataWithDefaults instantiates a new SystemInfoData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInfoDataWithDefaults() *SystemInfoData {
	this := SystemInfoData{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SystemInfoData) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoData) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SystemInfoData) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SystemInfoData) SetVersion(v string) {
	o.Version = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *SystemInfoData) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoData) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *SystemInfoData) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *SystemInfoData) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetPhpVersion returns the PhpVersion field value if set, zero value otherwise.
func (o *SystemInfoData) GetPhpVersion() string {
	if o == nil || IsNil(o.PhpVersion) {
		var ret string
		return ret
	}
	return *o.PhpVersion
}

// GetPhpVersionOk returns a tuple with the PhpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoData) GetPhpVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PhpVersion) {
		return nil, false
	}
	return o.PhpVersion, true
}

// HasPhpVersion returns a boolean if a field has been set.
func (o *SystemInfoData) HasPhpVersion() bool {
	if o != nil && !IsNil(o.PhpVersion) {
		return true
	}

	return false
}

// SetPhpVersion gets a reference to the given string and assigns it to the PhpVersion field.
func (o *SystemInfoData) SetPhpVersion(v string) {
	o.PhpVersion = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *SystemInfoData) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoData) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *SystemInfoData) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *SystemInfoData) SetOs(v string) {
	o.Os = &v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *SystemInfoData) GetDriver() string {
	if o == nil || IsNil(o.Driver) {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfoData) GetDriverOk() (*string, bool) {
	if o == nil || IsNil(o.Driver) {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *SystemInfoData) HasDriver() bool {
	if o != nil && !IsNil(o.Driver) {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *SystemInfoData) SetDriver(v string) {
	o.Driver = &v
}

func (o SystemInfoData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInfoData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["api_version"] = o.ApiVersion
	}
	if !IsNil(o.PhpVersion) {
		toSerialize["php_version"] = o.PhpVersion
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.Driver) {
		toSerialize["driver"] = o.Driver
	}
	return toSerialize, nil
}

type NullableSystemInfoData struct {
	value *SystemInfoData
	isSet bool
}

func (v NullableSystemInfoData) Get() *SystemInfoData {
	return v.value
}

func (v *NullableSystemInfoData) Set(val *SystemInfoData) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInfoData) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInfoData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInfoData(val *SystemInfoData) *NullableSystemInfoData {
	return &NullableSystemInfoData{value: val, isSet: true}
}

func (v NullableSystemInfoData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInfoData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

