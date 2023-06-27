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

// checks if the WebhookMessageArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookMessageArray{}

// WebhookMessageArray struct for WebhookMessageArray
type WebhookMessageArray struct {
	Data []WebhookMessageRead `json:"data"`
	Meta Meta `json:"meta"`
}

// NewWebhookMessageArray instantiates a new WebhookMessageArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookMessageArray(data []WebhookMessageRead, meta Meta) *WebhookMessageArray {
	this := WebhookMessageArray{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewWebhookMessageArrayWithDefaults instantiates a new WebhookMessageArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookMessageArrayWithDefaults() *WebhookMessageArray {
	this := WebhookMessageArray{}
	return &this
}

// GetData returns the Data field value
func (o *WebhookMessageArray) GetData() []WebhookMessageRead {
	if o == nil {
		var ret []WebhookMessageRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WebhookMessageArray) GetDataOk() ([]WebhookMessageRead, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *WebhookMessageArray) SetData(v []WebhookMessageRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *WebhookMessageArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *WebhookMessageArray) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *WebhookMessageArray) SetMeta(v Meta) {
	o.Meta = v
}

func (o WebhookMessageArray) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookMessageArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

type NullableWebhookMessageArray struct {
	value *WebhookMessageArray
	isSet bool
}

func (v NullableWebhookMessageArray) Get() *WebhookMessageArray {
	return v.value
}

func (v *NullableWebhookMessageArray) Set(val *WebhookMessageArray) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookMessageArray) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookMessageArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookMessageArray(val *WebhookMessageArray) *NullableWebhookMessageArray {
	return &NullableWebhookMessageArray{value: val, isSet: true}
}

func (v NullableWebhookMessageArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookMessageArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


