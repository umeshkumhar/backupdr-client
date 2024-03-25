# DiskRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Eligible** | Pointer to **bool** |  | [optional] 
**Freespace** | Pointer to **string** |  | [optional] 
**Requiredsize** | Pointer to **float64** |  | [optional] 

## Methods

### NewDiskRest

`func NewDiskRest() *DiskRest`

NewDiskRest instantiates a new DiskRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskRestWithDefaults

`func NewDiskRestWithDefaults() *DiskRest`

NewDiskRestWithDefaults instantiates a new DiskRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DiskRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiskRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiskRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiskRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEligible

`func (o *DiskRest) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *DiskRest) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *DiskRest) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *DiskRest) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetFreespace

`func (o *DiskRest) GetFreespace() string`

GetFreespace returns the Freespace field if non-nil, zero value otherwise.

### GetFreespaceOk

`func (o *DiskRest) GetFreespaceOk() (*string, bool)`

GetFreespaceOk returns a tuple with the Freespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreespace

`func (o *DiskRest) SetFreespace(v string)`

SetFreespace sets Freespace field to given value.

### HasFreespace

`func (o *DiskRest) HasFreespace() bool`

HasFreespace returns a boolean if a field has been set.

### GetRequiredsize

`func (o *DiskRest) GetRequiredsize() float64`

GetRequiredsize returns the Requiredsize field if non-nil, zero value otherwise.

### GetRequiredsizeOk

`func (o *DiskRest) GetRequiredsizeOk() (*float64, bool)`

GetRequiredsizeOk returns a tuple with the Requiredsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredsize

`func (o *DiskRest) SetRequiredsize(v float64)`

SetRequiredsize sets Requiredsize field to given value.

### HasRequiredsize

`func (o *DiskRest) HasRequiredsize() bool`

HasRequiredsize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


