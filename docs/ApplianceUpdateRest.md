# ApplianceUpdateRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusterid** | Pointer to **int64** |  | [optional] 
**Appliancename** | Pointer to **string** |  | [optional] 
**Updatestatus** | Pointer to **string** |  | [optional] 
**Updateinformation** | Pointer to [**UpdateInformationRest**](UpdateInformationRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewApplianceUpdateRest

`func NewApplianceUpdateRest() *ApplianceUpdateRest`

NewApplianceUpdateRest instantiates a new ApplianceUpdateRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceUpdateRestWithDefaults

`func NewApplianceUpdateRestWithDefaults() *ApplianceUpdateRest`

NewApplianceUpdateRestWithDefaults instantiates a new ApplianceUpdateRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterid

`func (o *ApplianceUpdateRest) GetClusterid() int64`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *ApplianceUpdateRest) GetClusteridOk() (*int64, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *ApplianceUpdateRest) SetClusterid(v int64)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *ApplianceUpdateRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.

### GetAppliancename

`func (o *ApplianceUpdateRest) GetAppliancename() string`

GetAppliancename returns the Appliancename field if non-nil, zero value otherwise.

### GetAppliancenameOk

`func (o *ApplianceUpdateRest) GetAppliancenameOk() (*string, bool)`

GetAppliancenameOk returns a tuple with the Appliancename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliancename

`func (o *ApplianceUpdateRest) SetAppliancename(v string)`

SetAppliancename sets Appliancename field to given value.

### HasAppliancename

`func (o *ApplianceUpdateRest) HasAppliancename() bool`

HasAppliancename returns a boolean if a field has been set.

### GetUpdatestatus

`func (o *ApplianceUpdateRest) GetUpdatestatus() string`

GetUpdatestatus returns the Updatestatus field if non-nil, zero value otherwise.

### GetUpdatestatusOk

`func (o *ApplianceUpdateRest) GetUpdatestatusOk() (*string, bool)`

GetUpdatestatusOk returns a tuple with the Updatestatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatestatus

`func (o *ApplianceUpdateRest) SetUpdatestatus(v string)`

SetUpdatestatus sets Updatestatus field to given value.

### HasUpdatestatus

`func (o *ApplianceUpdateRest) HasUpdatestatus() bool`

HasUpdatestatus returns a boolean if a field has been set.

### GetUpdateinformation

`func (o *ApplianceUpdateRest) GetUpdateinformation() UpdateInformationRest`

GetUpdateinformation returns the Updateinformation field if non-nil, zero value otherwise.

### GetUpdateinformationOk

`func (o *ApplianceUpdateRest) GetUpdateinformationOk() (*UpdateInformationRest, bool)`

GetUpdateinformationOk returns a tuple with the Updateinformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateinformation

`func (o *ApplianceUpdateRest) SetUpdateinformation(v UpdateInformationRest)`

SetUpdateinformation sets Updateinformation field to given value.

### HasUpdateinformation

`func (o *ApplianceUpdateRest) HasUpdateinformation() bool`

HasUpdateinformation returns a boolean if a field has been set.

### GetId

`func (o *ApplianceUpdateRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplianceUpdateRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplianceUpdateRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplianceUpdateRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ApplianceUpdateRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplianceUpdateRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplianceUpdateRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplianceUpdateRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ApplianceUpdateRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ApplianceUpdateRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ApplianceUpdateRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ApplianceUpdateRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ApplianceUpdateRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ApplianceUpdateRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ApplianceUpdateRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ApplianceUpdateRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


