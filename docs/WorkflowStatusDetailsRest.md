# WorkflowStatusDetailsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Startdate** | Pointer to **int64** |  | [optional] 
**Enddate** | Pointer to **int64** |  | [optional] 
**Jobtag** | Pointer to **string** |  | [optional] 
**Current** | Pointer to **string** |  | [optional] 
**Pending** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewWorkflowStatusDetailsRest

`func NewWorkflowStatusDetailsRest() *WorkflowStatusDetailsRest`

NewWorkflowStatusDetailsRest instantiates a new WorkflowStatusDetailsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStatusDetailsRestWithDefaults

`func NewWorkflowStatusDetailsRestWithDefaults() *WorkflowStatusDetailsRest`

NewWorkflowStatusDetailsRestWithDefaults instantiates a new WorkflowStatusDetailsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *WorkflowStatusDetailsRest) GetCompleted() string`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *WorkflowStatusDetailsRest) GetCompletedOk() (*string, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *WorkflowStatusDetailsRest) SetCompleted(v string)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *WorkflowStatusDetailsRest) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetResult

`func (o *WorkflowStatusDetailsRest) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowStatusDetailsRest) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowStatusDetailsRest) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *WorkflowStatusDetailsRest) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowStatusDetailsRest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowStatusDetailsRest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowStatusDetailsRest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowStatusDetailsRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartdate

`func (o *WorkflowStatusDetailsRest) GetStartdate() int64`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *WorkflowStatusDetailsRest) GetStartdateOk() (*int64, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *WorkflowStatusDetailsRest) SetStartdate(v int64)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *WorkflowStatusDetailsRest) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetEnddate

`func (o *WorkflowStatusDetailsRest) GetEnddate() int64`

GetEnddate returns the Enddate field if non-nil, zero value otherwise.

### GetEnddateOk

`func (o *WorkflowStatusDetailsRest) GetEnddateOk() (*int64, bool)`

GetEnddateOk returns a tuple with the Enddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnddate

`func (o *WorkflowStatusDetailsRest) SetEnddate(v int64)`

SetEnddate sets Enddate field to given value.

### HasEnddate

`func (o *WorkflowStatusDetailsRest) HasEnddate() bool`

HasEnddate returns a boolean if a field has been set.

### GetJobtag

`func (o *WorkflowStatusDetailsRest) GetJobtag() string`

GetJobtag returns the Jobtag field if non-nil, zero value otherwise.

### GetJobtagOk

`func (o *WorkflowStatusDetailsRest) GetJobtagOk() (*string, bool)`

GetJobtagOk returns a tuple with the Jobtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobtag

`func (o *WorkflowStatusDetailsRest) SetJobtag(v string)`

SetJobtag sets Jobtag field to given value.

### HasJobtag

`func (o *WorkflowStatusDetailsRest) HasJobtag() bool`

HasJobtag returns a boolean if a field has been set.

### GetCurrent

`func (o *WorkflowStatusDetailsRest) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *WorkflowStatusDetailsRest) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *WorkflowStatusDetailsRest) SetCurrent(v string)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *WorkflowStatusDetailsRest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetPending

`func (o *WorkflowStatusDetailsRest) GetPending() string`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *WorkflowStatusDetailsRest) GetPendingOk() (*string, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *WorkflowStatusDetailsRest) SetPending(v string)`

SetPending sets Pending field to given value.

### HasPending

`func (o *WorkflowStatusDetailsRest) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetId

`func (o *WorkflowStatusDetailsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowStatusDetailsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowStatusDetailsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowStatusDetailsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *WorkflowStatusDetailsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WorkflowStatusDetailsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WorkflowStatusDetailsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *WorkflowStatusDetailsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *WorkflowStatusDetailsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *WorkflowStatusDetailsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *WorkflowStatusDetailsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *WorkflowStatusDetailsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *WorkflowStatusDetailsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *WorkflowStatusDetailsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *WorkflowStatusDetailsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *WorkflowStatusDetailsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


