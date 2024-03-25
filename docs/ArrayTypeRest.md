# ArrayTypeRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**[]ArrayOptionRest**](ArrayOptionRest.md) |  | [optional] 
**Appliances** | Pointer to [**[]ClusterRest**](ClusterRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewArrayTypeRest

`func NewArrayTypeRest() *ArrayTypeRest`

NewArrayTypeRest instantiates a new ArrayTypeRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArrayTypeRestWithDefaults

`func NewArrayTypeRestWithDefaults() *ArrayTypeRest`

NewArrayTypeRestWithDefaults instantiates a new ArrayTypeRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArrayTypeRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArrayTypeRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArrayTypeRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArrayTypeRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *ArrayTypeRest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ArrayTypeRest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ArrayTypeRest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ArrayTypeRest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOptions

`func (o *ArrayTypeRest) GetOptions() []ArrayOptionRest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ArrayTypeRest) GetOptionsOk() (*[]ArrayOptionRest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ArrayTypeRest) SetOptions(v []ArrayOptionRest)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ArrayTypeRest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetAppliances

`func (o *ArrayTypeRest) GetAppliances() []ClusterRest`

GetAppliances returns the Appliances field if non-nil, zero value otherwise.

### GetAppliancesOk

`func (o *ArrayTypeRest) GetAppliancesOk() (*[]ClusterRest, bool)`

GetAppliancesOk returns a tuple with the Appliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliances

`func (o *ArrayTypeRest) SetAppliances(v []ClusterRest)`

SetAppliances sets Appliances field to given value.

### HasAppliances

`func (o *ArrayTypeRest) HasAppliances() bool`

HasAppliances returns a boolean if a field has been set.

### GetId

`func (o *ArrayTypeRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArrayTypeRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArrayTypeRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArrayTypeRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ArrayTypeRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArrayTypeRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArrayTypeRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ArrayTypeRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ArrayTypeRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ArrayTypeRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ArrayTypeRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ArrayTypeRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ArrayTypeRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ArrayTypeRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ArrayTypeRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ArrayTypeRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


