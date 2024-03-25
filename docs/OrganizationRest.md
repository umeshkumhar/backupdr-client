# OrganizationRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Createdate** | Pointer to **int64** |  | [optional] 
**Resourcecollection** | Pointer to [**CollectionRest**](CollectionRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewOrganizationRest

`func NewOrganizationRest() *OrganizationRest`

NewOrganizationRest instantiates a new OrganizationRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationRestWithDefaults

`func NewOrganizationRestWithDefaults() *OrganizationRest`

NewOrganizationRestWithDefaults instantiates a new OrganizationRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *OrganizationRest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OrganizationRest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OrganizationRest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OrganizationRest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OrganizationRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModifydate

`func (o *OrganizationRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *OrganizationRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *OrganizationRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *OrganizationRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetCreatedate

`func (o *OrganizationRest) GetCreatedate() int64`

GetCreatedate returns the Createdate field if non-nil, zero value otherwise.

### GetCreatedateOk

`func (o *OrganizationRest) GetCreatedateOk() (*int64, bool)`

GetCreatedateOk returns a tuple with the Createdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedate

`func (o *OrganizationRest) SetCreatedate(v int64)`

SetCreatedate sets Createdate field to given value.

### HasCreatedate

`func (o *OrganizationRest) HasCreatedate() bool`

HasCreatedate returns a boolean if a field has been set.

### GetResourcecollection

`func (o *OrganizationRest) GetResourcecollection() CollectionRest`

GetResourcecollection returns the Resourcecollection field if non-nil, zero value otherwise.

### GetResourcecollectionOk

`func (o *OrganizationRest) GetResourcecollectionOk() (*CollectionRest, bool)`

GetResourcecollectionOk returns a tuple with the Resourcecollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcecollection

`func (o *OrganizationRest) SetResourcecollection(v CollectionRest)`

SetResourcecollection sets Resourcecollection field to given value.

### HasResourcecollection

`func (o *OrganizationRest) HasResourcecollection() bool`

HasResourcecollection returns a boolean if a field has been set.

### GetId

`func (o *OrganizationRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *OrganizationRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *OrganizationRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *OrganizationRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *OrganizationRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *OrganizationRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *OrganizationRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *OrganizationRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *OrganizationRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *OrganizationRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *OrganizationRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *OrganizationRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *OrganizationRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


