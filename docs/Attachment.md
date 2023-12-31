# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**AttachableType** | [**AttachableType**](AttachableType.md) |  | 
**AttachableId** | **string** | ID of the model this attachment is linked to. | 
**Md5** | Pointer to **string** | MD5 hash of the file for basic duplicate detection. | [optional] 
**Filename** | **string** |  | 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**UploadUrl** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Mime** | Pointer to **string** |  | [optional] [readonly] 
**Size** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewAttachment

`func NewAttachment(attachableType AttachableType, attachableId string, filename string, ) *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Attachment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Attachment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Attachment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Attachment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Attachment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Attachment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Attachment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Attachment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAttachableType

`func (o *Attachment) GetAttachableType() AttachableType`

GetAttachableType returns the AttachableType field if non-nil, zero value otherwise.

### GetAttachableTypeOk

`func (o *Attachment) GetAttachableTypeOk() (*AttachableType, bool)`

GetAttachableTypeOk returns a tuple with the AttachableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableType

`func (o *Attachment) SetAttachableType(v AttachableType)`

SetAttachableType sets AttachableType field to given value.


### GetAttachableId

`func (o *Attachment) GetAttachableId() string`

GetAttachableId returns the AttachableId field if non-nil, zero value otherwise.

### GetAttachableIdOk

`func (o *Attachment) GetAttachableIdOk() (*string, bool)`

GetAttachableIdOk returns a tuple with the AttachableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableId

`func (o *Attachment) SetAttachableId(v string)`

SetAttachableId sets AttachableId field to given value.


### GetMd5

`func (o *Attachment) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *Attachment) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *Attachment) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *Attachment) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetFilename

`func (o *Attachment) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Attachment) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Attachment) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetDownloadUrl

`func (o *Attachment) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *Attachment) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *Attachment) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *Attachment) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetUploadUrl

`func (o *Attachment) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *Attachment) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *Attachment) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *Attachment) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetTitle

`func (o *Attachment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Attachment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Attachment) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Attachment) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetNotes

`func (o *Attachment) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Attachment) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Attachment) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Attachment) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Attachment) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Attachment) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetMime

`func (o *Attachment) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *Attachment) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *Attachment) SetMime(v string)`

SetMime sets Mime field to given value.

### HasMime

`func (o *Attachment) HasMime() bool`

HasMime returns a boolean if a field has been set.

### GetSize

`func (o *Attachment) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Attachment) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Attachment) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Attachment) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


