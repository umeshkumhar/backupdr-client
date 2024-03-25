# RefreshRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to **[]string** |  | [optional] 
**Sourceimage** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewRefreshRest

`func NewRefreshRest() *RefreshRest`

NewRefreshRest instantiates a new RefreshRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshRestWithDefaults

`func NewRefreshRestWithDefaults() *RefreshRest`

NewRefreshRestWithDefaults instantiates a new RefreshRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *RefreshRest) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RefreshRest) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RefreshRest) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RefreshRest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSourceimage

`func (o *RefreshRest) GetSourceimage() string`

GetSourceimage returns the Sourceimage field if non-nil, zero value otherwise.

### GetSourceimageOk

`func (o *RefreshRest) GetSourceimageOk() (*string, bool)`

GetSourceimageOk returns a tuple with the Sourceimage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceimage

`func (o *RefreshRest) SetSourceimage(v string)`

SetSourceimage sets Sourceimage field to given value.

### HasSourceimage

`func (o *RefreshRest) HasSourceimage() bool`

HasSourceimage returns a boolean if a field has been set.

### GetId

`func (o *RefreshRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RefreshRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RefreshRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RefreshRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *RefreshRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RefreshRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RefreshRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RefreshRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *RefreshRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *RefreshRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *RefreshRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *RefreshRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *RefreshRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *RefreshRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *RefreshRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *RefreshRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


