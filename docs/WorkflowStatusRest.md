# WorkflowStatusRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to [**WorkflowStatusDetailsRest**](WorkflowStatusDetailsRest.md) |  | [optional] 
**Next** | Pointer to [**WorkflowStatusDetailsRest**](WorkflowStatusDetailsRest.md) |  | [optional] 
**Prev** | Pointer to [**WorkflowStatusDetailsRest**](WorkflowStatusDetailsRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewWorkflowStatusRest

`func NewWorkflowStatusRest() *WorkflowStatusRest`

NewWorkflowStatusRest instantiates a new WorkflowStatusRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStatusRestWithDefaults

`func NewWorkflowStatusRestWithDefaults() *WorkflowStatusRest`

NewWorkflowStatusRestWithDefaults instantiates a new WorkflowStatusRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *WorkflowStatusRest) GetCurrent() WorkflowStatusDetailsRest`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *WorkflowStatusRest) GetCurrentOk() (*WorkflowStatusDetailsRest, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *WorkflowStatusRest) SetCurrent(v WorkflowStatusDetailsRest)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *WorkflowStatusRest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetNext

`func (o *WorkflowStatusRest) GetNext() WorkflowStatusDetailsRest`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *WorkflowStatusRest) GetNextOk() (*WorkflowStatusDetailsRest, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *WorkflowStatusRest) SetNext(v WorkflowStatusDetailsRest)`

SetNext sets Next field to given value.

### HasNext

`func (o *WorkflowStatusRest) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *WorkflowStatusRest) GetPrev() WorkflowStatusDetailsRest`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *WorkflowStatusRest) GetPrevOk() (*WorkflowStatusDetailsRest, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *WorkflowStatusRest) SetPrev(v WorkflowStatusDetailsRest)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *WorkflowStatusRest) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetId

`func (o *WorkflowStatusRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowStatusRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowStatusRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowStatusRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *WorkflowStatusRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WorkflowStatusRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WorkflowStatusRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *WorkflowStatusRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *WorkflowStatusRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *WorkflowStatusRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *WorkflowStatusRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *WorkflowStatusRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *WorkflowStatusRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *WorkflowStatusRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *WorkflowStatusRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *WorkflowStatusRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


