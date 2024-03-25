# ChoiceMetaRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isvm** | Pointer to **bool** |  | [optional] 
**Pathname** | Pointer to **string** |  | [optional] 
**Isclustered** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewChoiceMetaRest

`func NewChoiceMetaRest() *ChoiceMetaRest`

NewChoiceMetaRest instantiates a new ChoiceMetaRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChoiceMetaRestWithDefaults

`func NewChoiceMetaRestWithDefaults() *ChoiceMetaRest`

NewChoiceMetaRestWithDefaults instantiates a new ChoiceMetaRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsvm

`func (o *ChoiceMetaRest) GetIsvm() bool`

GetIsvm returns the Isvm field if non-nil, zero value otherwise.

### GetIsvmOk

`func (o *ChoiceMetaRest) GetIsvmOk() (*bool, bool)`

GetIsvmOk returns a tuple with the Isvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvm

`func (o *ChoiceMetaRest) SetIsvm(v bool)`

SetIsvm sets Isvm field to given value.

### HasIsvm

`func (o *ChoiceMetaRest) HasIsvm() bool`

HasIsvm returns a boolean if a field has been set.

### GetPathname

`func (o *ChoiceMetaRest) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *ChoiceMetaRest) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *ChoiceMetaRest) SetPathname(v string)`

SetPathname sets Pathname field to given value.

### HasPathname

`func (o *ChoiceMetaRest) HasPathname() bool`

HasPathname returns a boolean if a field has been set.

### GetIsclustered

`func (o *ChoiceMetaRest) GetIsclustered() bool`

GetIsclustered returns the Isclustered field if non-nil, zero value otherwise.

### GetIsclusteredOk

`func (o *ChoiceMetaRest) GetIsclusteredOk() (*bool, bool)`

GetIsclusteredOk returns a tuple with the Isclustered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsclustered

`func (o *ChoiceMetaRest) SetIsclustered(v bool)`

SetIsclustered sets Isclustered field to given value.

### HasIsclustered

`func (o *ChoiceMetaRest) HasIsclustered() bool`

HasIsclustered returns a boolean if a field has been set.

### GetId

`func (o *ChoiceMetaRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChoiceMetaRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChoiceMetaRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChoiceMetaRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ChoiceMetaRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ChoiceMetaRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ChoiceMetaRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ChoiceMetaRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ChoiceMetaRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ChoiceMetaRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ChoiceMetaRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ChoiceMetaRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ChoiceMetaRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ChoiceMetaRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ChoiceMetaRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ChoiceMetaRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


