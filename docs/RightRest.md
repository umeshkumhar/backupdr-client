# RightRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]RightRest**](RightRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Parents** | Pointer to [**[]RightRest**](RightRest.md) |  | [optional] 
**Name** | Pointer to **string** | Name of access right | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewRightRest

`func NewRightRest() *RightRest`

NewRightRest instantiates a new RightRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRightRestWithDefaults

`func NewRightRestWithDefaults() *RightRest`

NewRightRestWithDefaults instantiates a new RightRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *RightRest) GetChildren() []RightRest`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *RightRest) GetChildrenOk() (*[]RightRest, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *RightRest) SetChildren(v []RightRest)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *RightRest) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetId

`func (o *RightRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RightRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RightRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RightRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParents

`func (o *RightRest) GetParents() []RightRest`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *RightRest) GetParentsOk() (*[]RightRest, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *RightRest) SetParents(v []RightRest)`

SetParents sets Parents field to given value.

### HasParents

`func (o *RightRest) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetName

`func (o *RightRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RightRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RightRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RightRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHref

`func (o *RightRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RightRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RightRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RightRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *RightRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *RightRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *RightRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *RightRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *RightRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *RightRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *RightRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *RightRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


