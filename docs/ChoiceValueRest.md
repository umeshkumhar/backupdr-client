# ChoiceValueRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Selected** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewChoiceValueRest

`func NewChoiceValueRest() *ChoiceValueRest`

NewChoiceValueRest instantiates a new ChoiceValueRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChoiceValueRestWithDefaults

`func NewChoiceValueRestWithDefaults() *ChoiceValueRest`

NewChoiceValueRestWithDefaults instantiates a new ChoiceValueRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ChoiceValueRest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ChoiceValueRest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ChoiceValueRest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ChoiceValueRest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *ChoiceValueRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChoiceValueRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChoiceValueRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChoiceValueRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ChoiceValueRest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ChoiceValueRest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ChoiceValueRest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ChoiceValueRest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetKey

`func (o *ChoiceValueRest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ChoiceValueRest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ChoiceValueRest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ChoiceValueRest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *ChoiceValueRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChoiceValueRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChoiceValueRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChoiceValueRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisabled

`func (o *ChoiceValueRest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ChoiceValueRest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ChoiceValueRest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ChoiceValueRest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetSelected

`func (o *ChoiceValueRest) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *ChoiceValueRest) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *ChoiceValueRest) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *ChoiceValueRest) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### GetId

`func (o *ChoiceValueRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChoiceValueRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChoiceValueRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChoiceValueRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ChoiceValueRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ChoiceValueRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ChoiceValueRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ChoiceValueRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ChoiceValueRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ChoiceValueRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ChoiceValueRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ChoiceValueRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ChoiceValueRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ChoiceValueRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ChoiceValueRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ChoiceValueRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


