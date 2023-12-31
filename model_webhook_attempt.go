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

// checks if the WebhookAttempt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookAttempt{}

// WebhookAttempt struct for WebhookAttempt
type WebhookAttempt struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The ID of the webhook message this attempt belongs to.
	WebhookMessageId *string `json:"webhook_message_id,omitempty"`
	// The HTTP status code of the error, if any.
	StatusCode NullableInt32 `json:"status_code,omitempty"`
	// Internal log for this attempt. May contain sensitive user data.
	Logs NullableString `json:"logs,omitempty"`
	// Webhook receiver response for this attempt, if any. May contain sensitive user data.
	Response NullableString `json:"response,omitempty"`
}

// NewWebhookAttempt instantiates a new WebhookAttempt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookAttempt() *WebhookAttempt {
	this := WebhookAttempt{}
	return &this
}

// NewWebhookAttemptWithDefaults instantiates a new WebhookAttempt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookAttemptWithDefaults() *WebhookAttempt {
	this := WebhookAttempt{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WebhookAttempt) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAttempt) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WebhookAttempt) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WebhookAttempt) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WebhookAttempt) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAttempt) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WebhookAttempt) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WebhookAttempt) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetWebhookMessageId returns the WebhookMessageId field value if set, zero value otherwise.
func (o *WebhookAttempt) GetWebhookMessageId() string {
	if o == nil || IsNil(o.WebhookMessageId) {
		var ret string
		return ret
	}
	return *o.WebhookMessageId
}

// GetWebhookMessageIdOk returns a tuple with the WebhookMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookAttempt) GetWebhookMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookMessageId) {
		return nil, false
	}
	return o.WebhookMessageId, true
}

// HasWebhookMessageId returns a boolean if a field has been set.
func (o *WebhookAttempt) HasWebhookMessageId() bool {
	if o != nil && !IsNil(o.WebhookMessageId) {
		return true
	}

	return false
}

// SetWebhookMessageId gets a reference to the given string and assigns it to the WebhookMessageId field.
func (o *WebhookAttempt) SetWebhookMessageId(v string) {
	o.WebhookMessageId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookAttempt) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode.Get()) {
		var ret int32
		return ret
	}
	return *o.StatusCode.Get()
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookAttempt) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusCode.Get(), o.StatusCode.IsSet()
}

// HasStatusCode returns a boolean if a field has been set.
func (o *WebhookAttempt) HasStatusCode() bool {
	if o != nil && o.StatusCode.IsSet() {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given NullableInt32 and assigns it to the StatusCode field.
func (o *WebhookAttempt) SetStatusCode(v int32) {
	o.StatusCode.Set(&v)
}
// SetStatusCodeNil sets the value for StatusCode to be an explicit nil
func (o *WebhookAttempt) SetStatusCodeNil() {
	o.StatusCode.Set(nil)
}

// UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
func (o *WebhookAttempt) UnsetStatusCode() {
	o.StatusCode.Unset()
}

// GetLogs returns the Logs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookAttempt) GetLogs() string {
	if o == nil || IsNil(o.Logs.Get()) {
		var ret string
		return ret
	}
	return *o.Logs.Get()
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookAttempt) GetLogsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs.Get(), o.Logs.IsSet()
}

// HasLogs returns a boolean if a field has been set.
func (o *WebhookAttempt) HasLogs() bool {
	if o != nil && o.Logs.IsSet() {
		return true
	}

	return false
}

// SetLogs gets a reference to the given NullableString and assigns it to the Logs field.
func (o *WebhookAttempt) SetLogs(v string) {
	o.Logs.Set(&v)
}
// SetLogsNil sets the value for Logs to be an explicit nil
func (o *WebhookAttempt) SetLogsNil() {
	o.Logs.Set(nil)
}

// UnsetLogs ensures that no value is present for Logs, not even an explicit nil
func (o *WebhookAttempt) UnsetLogs() {
	o.Logs.Unset()
}

// GetResponse returns the Response field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookAttempt) GetResponse() string {
	if o == nil || IsNil(o.Response.Get()) {
		var ret string
		return ret
	}
	return *o.Response.Get()
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookAttempt) GetResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Response.Get(), o.Response.IsSet()
}

// HasResponse returns a boolean if a field has been set.
func (o *WebhookAttempt) HasResponse() bool {
	if o != nil && o.Response.IsSet() {
		return true
	}

	return false
}

// SetResponse gets a reference to the given NullableString and assigns it to the Response field.
func (o *WebhookAttempt) SetResponse(v string) {
	o.Response.Set(&v)
}
// SetResponseNil sets the value for Response to be an explicit nil
func (o *WebhookAttempt) SetResponseNil() {
	o.Response.Set(nil)
}

// UnsetResponse ensures that no value is present for Response, not even an explicit nil
func (o *WebhookAttempt) UnsetResponse() {
	o.Response.Unset()
}

func (o WebhookAttempt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookAttempt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	if !IsNil(o.WebhookMessageId) {
		toSerialize["webhook_message_id"] = o.WebhookMessageId
	}
	if o.StatusCode.IsSet() {
		toSerialize["status_code"] = o.StatusCode.Get()
	}
	if o.Logs.IsSet() {
		toSerialize["logs"] = o.Logs.Get()
	}
	if o.Response.IsSet() {
		toSerialize["response"] = o.Response.Get()
	}
	return toSerialize, nil
}

type NullableWebhookAttempt struct {
	value *WebhookAttempt
	isSet bool
}

func (v NullableWebhookAttempt) Get() *WebhookAttempt {
	return v.value
}

func (v *NullableWebhookAttempt) Set(val *WebhookAttempt) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookAttempt) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookAttempt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookAttempt(val *WebhookAttempt) *NullableWebhookAttempt {
	return &NullableWebhookAttempt{value: val, isSet: true}
}

func (v NullableWebhookAttempt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookAttempt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


