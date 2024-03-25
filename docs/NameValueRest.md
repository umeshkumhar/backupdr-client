# NameValueRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Selection** | Pointer to [**[]DescValueRest**](DescValueRest.md) |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Constant** | Pointer to **string** |  | [optional] 
**Select** | Pointer to **bool** |  | [optional] 
**Editableonmount** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewNameValueRest

`func NewNameValueRest() *NameValueRest`

NewNameValueRest instantiates a new NameValueRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameValueRestWithDefaults

`func NewNameValueRestWithDefaults() *NameValueRest`

NewNameValueRestWithDefaults instantiates a new NameValueRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *NameValueRest) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *NameValueRest) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *NameValueRest) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *NameValueRest) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetDescription

`func (o *NameValueRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NameValueRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NameValueRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NameValueRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *NameValueRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NameValueRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NameValueRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NameValueRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *NameValueRest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NameValueRest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NameValueRest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NameValueRest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDefault

`func (o *NameValueRest) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *NameValueRest) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *NameValueRest) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *NameValueRest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetType

`func (o *NameValueRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NameValueRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NameValueRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NameValueRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSelection

`func (o *NameValueRest) GetSelection() []DescValueRest`

GetSelection returns the Selection field if non-nil, zero value otherwise.

### GetSelectionOk

`func (o *NameValueRest) GetSelectionOk() (*[]DescValueRest, bool)`

GetSelectionOk returns a tuple with the Selection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelection

`func (o *NameValueRest) SetSelection(v []DescValueRest)`

SetSelection sets Selection field to given value.

### HasSelection

`func (o *NameValueRest) HasSelection() bool`

HasSelection returns a boolean if a field has been set.

### GetAlias

`func (o *NameValueRest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *NameValueRest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *NameValueRest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *NameValueRest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetConstant

`func (o *NameValueRest) GetConstant() string`

GetConstant returns the Constant field if non-nil, zero value otherwise.

### GetConstantOk

`func (o *NameValueRest) GetConstantOk() (*string, bool)`

GetConstantOk returns a tuple with the Constant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstant

`func (o *NameValueRest) SetConstant(v string)`

SetConstant sets Constant field to given value.

### HasConstant

`func (o *NameValueRest) HasConstant() bool`

HasConstant returns a boolean if a field has been set.

### GetSelect

`func (o *NameValueRest) GetSelect() bool`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *NameValueRest) GetSelectOk() (*bool, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *NameValueRest) SetSelect(v bool)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *NameValueRest) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetEditableonmount

`func (o *NameValueRest) GetEditableonmount() bool`

GetEditableonmount returns the Editableonmount field if non-nil, zero value otherwise.

### GetEditableonmountOk

`func (o *NameValueRest) GetEditableonmountOk() (*bool, bool)`

GetEditableonmountOk returns a tuple with the Editableonmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableonmount

`func (o *NameValueRest) SetEditableonmount(v bool)`

SetEditableonmount sets Editableonmount field to given value.

### HasEditableonmount

`func (o *NameValueRest) HasEditableonmount() bool`

HasEditableonmount returns a boolean if a field has been set.

### GetId

`func (o *NameValueRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NameValueRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NameValueRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NameValueRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *NameValueRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NameValueRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NameValueRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NameValueRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *NameValueRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *NameValueRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *NameValueRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *NameValueRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *NameValueRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *NameValueRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *NameValueRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *NameValueRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


