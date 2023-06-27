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
	"time"
)

// checks if the Webhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Webhook{}

// Webhook struct for Webhook
type Webhook struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Boolean to indicate if the webhook is active
	Active *bool `json:"active,omitempty"`
	// A title for the webhook for easy recognition.
	Title string `json:"title"`
	// A 24-character secret for the webhook. It's generated by Firefly III when saving a new webhook. If you submit a new secret through the PUT endpoint it will generate a new secret for the selected webhook, a new secret bearing no relation to whatever you just submitted.
	Secret *string `json:"secret,omitempty"`
	Trigger WebhookTrigger `json:"trigger"`
	Response WebhookResponse `json:"response"`
	Delivery WebhookDelivery `json:"delivery"`
	// The URL of the webhook. Has to start with `https`.
	Url string `json:"url"`
}

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook(title string, trigger WebhookTrigger, response WebhookResponse, delivery WebhookDelivery, url string) *Webhook {
	this := Webhook{}
	this.Title = title
	this.Trigger = trigger
	this.Response = response
	this.Delivery = delivery
	this.Url = url
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Webhook) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Webhook) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Webhook) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Webhook) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Webhook) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Webhook) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Webhook) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Webhook) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Webhook) SetActive(v bool) {
	o.Active = &v
}

// GetTitle returns the Title field value
func (o *Webhook) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Webhook) SetTitle(v string) {
	o.Title = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *Webhook) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *Webhook) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *Webhook) SetSecret(v string) {
	o.Secret = &v
}

// GetTrigger returns the Trigger field value
func (o *Webhook) GetTrigger() WebhookTrigger {
	if o == nil {
		var ret WebhookTrigger
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetTriggerOk() (*WebhookTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *Webhook) SetTrigger(v WebhookTrigger) {
	o.Trigger = v
}

// GetResponse returns the Response field value
func (o *Webhook) GetResponse() WebhookResponse {
	if o == nil {
		var ret WebhookResponse
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetResponseOk() (*WebhookResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *Webhook) SetResponse(v WebhookResponse) {
	o.Response = v
}

// GetDelivery returns the Delivery field value
func (o *Webhook) GetDelivery() WebhookDelivery {
	if o == nil {
		var ret WebhookDelivery
		return ret
	}

	return o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetDeliveryOk() (*WebhookDelivery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivery, true
}

// SetDelivery sets field value
func (o *Webhook) SetDelivery(v WebhookDelivery) {
	o.Delivery = v
}

// GetUrl returns the Url field value
func (o *Webhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Webhook) SetUrl(v string) {
	o.Url = v
}

func (o Webhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Webhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["title"] = o.Title
	// skip: secret is readOnly
	toSerialize["trigger"] = o.Trigger
	toSerialize["response"] = o.Response
	toSerialize["delivery"] = o.Delivery
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

