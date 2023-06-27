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

// checks if the WebhookRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookRead{}

// WebhookRead struct for WebhookRead
type WebhookRead struct {
	// Immutable value
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes Webhook `json:"attributes"`
	Links ObjectLink `json:"links"`
}

// NewWebhookRead instantiates a new WebhookRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookRead(type_ string, id string, attributes Webhook, links ObjectLink) *WebhookRead {
	this := WebhookRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewWebhookReadWithDefaults instantiates a new WebhookRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookReadWithDefaults() *WebhookRead {
	this := WebhookRead{}
	return &this
}

// GetType returns the Type field value
func (o *WebhookRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *WebhookRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *WebhookRead) GetAttributes() Webhook {
	if o == nil {
		var ret Webhook
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhookRead) GetAttributesOk() (*Webhook, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WebhookRead) SetAttributes(v Webhook) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *WebhookRead) GetLinks() ObjectLink {
	if o == nil {
		var ret ObjectLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *WebhookRead) GetLinksOk() (*ObjectLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *WebhookRead) SetLinks(v ObjectLink) {
	o.Links = v
}

func (o WebhookRead) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableWebhookRead struct {
	value *WebhookRead
	isSet bool
}

func (v NullableWebhookRead) Get() *WebhookRead {
	return v.value
}

func (v *NullableWebhookRead) Set(val *WebhookRead) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookRead) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookRead(val *WebhookRead) *NullableWebhookRead {
	return &NullableWebhookRead{value: val, isSet: true}
}

func (v NullableWebhookRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


