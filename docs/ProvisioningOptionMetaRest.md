# ProvisioningOptionMetaRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choice** | Pointer to [**[]ChoiceMetaRest**](ChoiceMetaRest.md) |  | [optional] 
**ValidationRegex** | Pointer to **string** |  | [optional] 
**Unique** | Pointer to **bool** |  | [optional] 
**Restoreimmutable** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**[]NameValueRest**](NameValueRest.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Select** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewProvisioningOptionMetaRest

`func NewProvisioningOptionMetaRest() *ProvisioningOptionMetaRest`

NewProvisioningOptionMetaRest instantiates a new ProvisioningOptionMetaRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningOptionMetaRestWithDefaults

`func NewProvisioningOptionMetaRestWithDefaults() *ProvisioningOptionMetaRest`

NewProvisioningOptionMetaRestWithDefaults instantiates a new ProvisioningOptionMetaRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoice

`func (o *ProvisioningOptionMetaRest) GetChoice() []ChoiceMetaRest`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *ProvisioningOptionMetaRest) GetChoiceOk() (*[]ChoiceMetaRest, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *ProvisioningOptionMetaRest) SetChoice(v []ChoiceMetaRest)`

SetChoice sets Choice field to given value.

### HasChoice

`func (o *ProvisioningOptionMetaRest) HasChoice() bool`

HasChoice returns a boolean if a field has been set.

### GetValidationRegex

`func (o *ProvisioningOptionMetaRest) GetValidationRegex() string`

GetValidationRegex returns the ValidationRegex field if non-nil, zero value otherwise.

### GetValidationRegexOk

`func (o *ProvisioningOptionMetaRest) GetValidationRegexOk() (*string, bool)`

GetValidationRegexOk returns a tuple with the ValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegex

`func (o *ProvisioningOptionMetaRest) SetValidationRegex(v string)`

SetValidationRegex sets ValidationRegex field to given value.

### HasValidationRegex

`func (o *ProvisioningOptionMetaRest) HasValidationRegex() bool`

HasValidationRegex returns a boolean if a field has been set.

### GetUnique

`func (o *ProvisioningOptionMetaRest) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *ProvisioningOptionMetaRest) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *ProvisioningOptionMetaRest) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *ProvisioningOptionMetaRest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetRestoreimmutable

`func (o *ProvisioningOptionMetaRest) GetRestoreimmutable() bool`

GetRestoreimmutable returns the Restoreimmutable field if non-nil, zero value otherwise.

### GetRestoreimmutableOk

`func (o *ProvisioningOptionMetaRest) GetRestoreimmutableOk() (*bool, bool)`

GetRestoreimmutableOk returns a tuple with the Restoreimmutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreimmutable

`func (o *ProvisioningOptionMetaRest) SetRestoreimmutable(v bool)`

SetRestoreimmutable sets Restoreimmutable field to given value.

### HasRestoreimmutable

`func (o *ProvisioningOptionMetaRest) HasRestoreimmutable() bool`

HasRestoreimmutable returns a boolean if a field has been set.

### GetRequired

`func (o *ProvisioningOptionMetaRest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ProvisioningOptionMetaRest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ProvisioningOptionMetaRest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ProvisioningOptionMetaRest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDescription

`func (o *ProvisioningOptionMetaRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProvisioningOptionMetaRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProvisioningOptionMetaRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProvisioningOptionMetaRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ProvisioningOptionMetaRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningOptionMetaRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningOptionMetaRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisioningOptionMetaRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ProvisioningOptionMetaRest) GetValue() []NameValueRest`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProvisioningOptionMetaRest) GetValueOk() (*[]NameValueRest, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProvisioningOptionMetaRest) SetValue(v []NameValueRest)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProvisioningOptionMetaRest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *ProvisioningOptionMetaRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProvisioningOptionMetaRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProvisioningOptionMetaRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProvisioningOptionMetaRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ProvisioningOptionMetaRest) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ProvisioningOptionMetaRest) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ProvisioningOptionMetaRest) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ProvisioningOptionMetaRest) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetAlias

`func (o *ProvisioningOptionMetaRest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ProvisioningOptionMetaRest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ProvisioningOptionMetaRest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ProvisioningOptionMetaRest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetSelect

`func (o *ProvisioningOptionMetaRest) GetSelect() bool`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *ProvisioningOptionMetaRest) GetSelectOk() (*bool, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *ProvisioningOptionMetaRest) SetSelect(v bool)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *ProvisioningOptionMetaRest) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetId

`func (o *ProvisioningOptionMetaRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningOptionMetaRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningOptionMetaRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisioningOptionMetaRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ProvisioningOptionMetaRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProvisioningOptionMetaRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProvisioningOptionMetaRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ProvisioningOptionMetaRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ProvisioningOptionMetaRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ProvisioningOptionMetaRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ProvisioningOptionMetaRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ProvisioningOptionMetaRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ProvisioningOptionMetaRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ProvisioningOptionMetaRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ProvisioningOptionMetaRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ProvisioningOptionMetaRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


