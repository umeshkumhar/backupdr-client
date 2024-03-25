# WorkflowRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Application** | Pointer to [**ApplicationRest**](ApplicationRest.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Status** | Pointer to [**WorkflowStatusRest**](WorkflowStatusRest.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Props** | Pointer to [**[]KeyValueRest**](KeyValueRest.md) |  | [optional] 
**Schedule** | Pointer to [**WorkScheduleRest**](WorkScheduleRest.md) |  | [optional] 
**Items** | Pointer to [**[]WorkflowItemRest**](WorkflowItemRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewWorkflowRest

`func NewWorkflowRest() *WorkflowRest`

NewWorkflowRest instantiates a new WorkflowRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRestWithDefaults

`func NewWorkflowRestWithDefaults() *WorkflowRest`

NewWorkflowRestWithDefaults instantiates a new WorkflowRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApplication

`func (o *WorkflowRest) GetApplication() ApplicationRest`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *WorkflowRest) GetApplicationOk() (*ApplicationRest, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *WorkflowRest) SetApplication(v ApplicationRest)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *WorkflowRest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCluster

`func (o *WorkflowRest) GetCluster() ClusterRest`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *WorkflowRest) GetClusterOk() (*ClusterRest, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *WorkflowRest) SetCluster(v ClusterRest)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *WorkflowRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowRest) GetStatus() WorkflowStatusRest`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRest) GetStatusOk() (*WorkflowStatusRest, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRest) SetStatus(v WorkflowStatusRest)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDisabled

`func (o *WorkflowRest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *WorkflowRest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *WorkflowRest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *WorkflowRest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetProps

`func (o *WorkflowRest) GetProps() []KeyValueRest`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *WorkflowRest) GetPropsOk() (*[]KeyValueRest, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *WorkflowRest) SetProps(v []KeyValueRest)`

SetProps sets Props field to given value.

### HasProps

`func (o *WorkflowRest) HasProps() bool`

HasProps returns a boolean if a field has been set.

### GetSchedule

`func (o *WorkflowRest) GetSchedule() WorkScheduleRest`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *WorkflowRest) GetScheduleOk() (*WorkScheduleRest, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *WorkflowRest) SetSchedule(v WorkScheduleRest)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *WorkflowRest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetItems

`func (o *WorkflowRest) GetItems() []WorkflowItemRest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WorkflowRest) GetItemsOk() (*[]WorkflowItemRest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WorkflowRest) SetItems(v []WorkflowItemRest)`

SetItems sets Items field to given value.

### HasItems

`func (o *WorkflowRest) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetId

`func (o *WorkflowRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *WorkflowRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WorkflowRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WorkflowRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *WorkflowRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *WorkflowRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *WorkflowRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *WorkflowRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *WorkflowRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *WorkflowRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *WorkflowRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *WorkflowRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *WorkflowRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


