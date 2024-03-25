# ConsolidatedApplianceUpdatesRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appliancedetails** | Pointer to [**[]ApplianceUpdateRest**](ApplianceUpdateRest.md) |  | [optional] 
**Updateinformation** | Pointer to [**UpdateInformationRest**](UpdateInformationRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewConsolidatedApplianceUpdatesRest

`func NewConsolidatedApplianceUpdatesRest() *ConsolidatedApplianceUpdatesRest`

NewConsolidatedApplianceUpdatesRest instantiates a new ConsolidatedApplianceUpdatesRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsolidatedApplianceUpdatesRestWithDefaults

`func NewConsolidatedApplianceUpdatesRestWithDefaults() *ConsolidatedApplianceUpdatesRest`

NewConsolidatedApplianceUpdatesRestWithDefaults instantiates a new ConsolidatedApplianceUpdatesRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliancedetails

`func (o *ConsolidatedApplianceUpdatesRest) GetAppliancedetails() []ApplianceUpdateRest`

GetAppliancedetails returns the Appliancedetails field if non-nil, zero value otherwise.

### GetAppliancedetailsOk

`func (o *ConsolidatedApplianceUpdatesRest) GetAppliancedetailsOk() (*[]ApplianceUpdateRest, bool)`

GetAppliancedetailsOk returns a tuple with the Appliancedetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliancedetails

`func (o *ConsolidatedApplianceUpdatesRest) SetAppliancedetails(v []ApplianceUpdateRest)`

SetAppliancedetails sets Appliancedetails field to given value.

### HasAppliancedetails

`func (o *ConsolidatedApplianceUpdatesRest) HasAppliancedetails() bool`

HasAppliancedetails returns a boolean if a field has been set.

### GetUpdateinformation

`func (o *ConsolidatedApplianceUpdatesRest) GetUpdateinformation() UpdateInformationRest`

GetUpdateinformation returns the Updateinformation field if non-nil, zero value otherwise.

### GetUpdateinformationOk

`func (o *ConsolidatedApplianceUpdatesRest) GetUpdateinformationOk() (*UpdateInformationRest, bool)`

GetUpdateinformationOk returns a tuple with the Updateinformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateinformation

`func (o *ConsolidatedApplianceUpdatesRest) SetUpdateinformation(v UpdateInformationRest)`

SetUpdateinformation sets Updateinformation field to given value.

### HasUpdateinformation

`func (o *ConsolidatedApplianceUpdatesRest) HasUpdateinformation() bool`

HasUpdateinformation returns a boolean if a field has been set.

### GetId

`func (o *ConsolidatedApplianceUpdatesRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsolidatedApplianceUpdatesRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsolidatedApplianceUpdatesRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsolidatedApplianceUpdatesRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ConsolidatedApplianceUpdatesRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConsolidatedApplianceUpdatesRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConsolidatedApplianceUpdatesRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConsolidatedApplianceUpdatesRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ConsolidatedApplianceUpdatesRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ConsolidatedApplianceUpdatesRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ConsolidatedApplianceUpdatesRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ConsolidatedApplianceUpdatesRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ConsolidatedApplianceUpdatesRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ConsolidatedApplianceUpdatesRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ConsolidatedApplianceUpdatesRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ConsolidatedApplianceUpdatesRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


