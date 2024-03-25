# ComponentDetailRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Installed** | Pointer to **int64** | Install time in UNIX Epoch time (microseconds since Jan 1, 1970) | [optional] 

## Methods

### NewComponentDetailRest

`func NewComponentDetailRest() *ComponentDetailRest`

NewComponentDetailRest instantiates a new ComponentDetailRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentDetailRestWithDefaults

`func NewComponentDetailRestWithDefaults() *ComponentDetailRest`

NewComponentDetailRestWithDefaults instantiates a new ComponentDetailRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ComponentDetailRest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ComponentDetailRest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ComponentDetailRest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ComponentDetailRest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetSummary

`func (o *ComponentDetailRest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ComponentDetailRest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ComponentDetailRest) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ComponentDetailRest) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetInstalled

`func (o *ComponentDetailRest) GetInstalled() int64`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *ComponentDetailRest) GetInstalledOk() (*int64, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *ComponentDetailRest) SetInstalled(v int64)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *ComponentDetailRest) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


