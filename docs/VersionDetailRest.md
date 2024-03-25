# VersionDetailRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** |  | [optional] 
**Components** | Pointer to [**[]ComponentDetailRest**](ComponentDetailRest.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Installed** | Pointer to **int64** | Install time in UNIX Epoch time (microseconds since Jan 1, 1970) | [optional] 

## Methods

### NewVersionDetailRest

`func NewVersionDetailRest() *VersionDetailRest`

NewVersionDetailRest instantiates a new VersionDetailRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionDetailRestWithDefaults

`func NewVersionDetailRestWithDefaults() *VersionDetailRest`

NewVersionDetailRestWithDefaults instantiates a new VersionDetailRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *VersionDetailRest) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *VersionDetailRest) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *VersionDetailRest) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *VersionDetailRest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetComponents

`func (o *VersionDetailRest) GetComponents() []ComponentDetailRest`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *VersionDetailRest) GetComponentsOk() (*[]ComponentDetailRest, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *VersionDetailRest) SetComponents(v []ComponentDetailRest)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *VersionDetailRest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetSummary

`func (o *VersionDetailRest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *VersionDetailRest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *VersionDetailRest) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *VersionDetailRest) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetInstalled

`func (o *VersionDetailRest) GetInstalled() int64`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *VersionDetailRest) GetInstalledOk() (*int64, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *VersionDetailRest) SetInstalled(v int64)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *VersionDetailRest) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


