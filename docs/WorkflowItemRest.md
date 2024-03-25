# WorkflowItemRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Props** | Pointer to [**[]KeyValueRest**](KeyValueRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewWorkflowItemRest

`func NewWorkflowItemRest() *WorkflowItemRest`

NewWorkflowItemRest instantiates a new WorkflowItemRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowItemRestWithDefaults

`func NewWorkflowItemRestWithDefaults() *WorkflowItemRest`

NewWorkflowItemRestWithDefaults instantiates a new WorkflowItemRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowItemRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowItemRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowItemRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowItemRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowItemRest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowItemRest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowItemRest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowItemRest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetProps

`func (o *WorkflowItemRest) GetProps() []KeyValueRest`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *WorkflowItemRest) GetPropsOk() (*[]KeyValueRest, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *WorkflowItemRest) SetProps(v []KeyValueRest)`

SetProps sets Props field to given value.

### HasProps

`func (o *WorkflowItemRest) HasProps() bool`

HasProps returns a boolean if a field has been set.

### GetId

`func (o *WorkflowItemRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowItemRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowItemRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowItemRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *WorkflowItemRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WorkflowItemRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WorkflowItemRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *WorkflowItemRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *WorkflowItemRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *WorkflowItemRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *WorkflowItemRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *WorkflowItemRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *WorkflowItemRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *WorkflowItemRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *WorkflowItemRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *WorkflowItemRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


