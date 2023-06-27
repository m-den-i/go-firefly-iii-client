# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Active** | Pointer to **bool** | Boolean to indicate if the webhook is active | [optional] 
**Title** | **string** | A title for the webhook for easy recognition. | 
**Secret** | Pointer to **string** | A 24-character secret for the webhook. It&#39;s generated by Firefly III when saving a new webhook. If you submit a new secret through the PUT endpoint it will generate a new secret for the selected webhook, a new secret bearing no relation to whatever you just submitted. | [optional] [readonly] 
**Trigger** | [**WebhookTrigger**](WebhookTrigger.md) |  | 
**Response** | [**WebhookResponse**](WebhookResponse.md) |  | 
**Delivery** | [**WebhookDelivery**](WebhookDelivery.md) |  | 
**Url** | **string** | The URL of the webhook. Has to start with &#x60;https&#x60;. | 

## Methods

### NewWebhook

`func NewWebhook(title string, trigger WebhookTrigger, response WebhookResponse, delivery WebhookDelivery, url string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Webhook) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Webhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Webhook) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Webhook) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Webhook) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Webhook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetActive

`func (o *Webhook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Webhook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Webhook) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Webhook) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTitle

`func (o *Webhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Webhook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Webhook) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSecret

`func (o *Webhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Webhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Webhook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Webhook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTrigger

`func (o *Webhook) GetTrigger() WebhookTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Webhook) GetTriggerOk() (*WebhookTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Webhook) SetTrigger(v WebhookTrigger)`

SetTrigger sets Trigger field to given value.


### GetResponse

`func (o *Webhook) GetResponse() WebhookResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Webhook) GetResponseOk() (*WebhookResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Webhook) SetResponse(v WebhookResponse)`

SetResponse sets Response field to given value.


### GetDelivery

`func (o *Webhook) GetDelivery() WebhookDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *Webhook) GetDeliveryOk() (*WebhookDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *Webhook) SetDelivery(v WebhookDelivery)`

SetDelivery sets Delivery field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

