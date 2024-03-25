# DiskMappingRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sourceinfo** | Pointer to [**[]DiskRest**](DiskRest.md) |  | [optional] 
**Targetlist** | Pointer to [**[]DiskRest**](DiskRest.md) |  | [optional] 
**Asmracnodelist** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDiskMappingRest

`func NewDiskMappingRest() *DiskMappingRest`

NewDiskMappingRest instantiates a new DiskMappingRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskMappingRestWithDefaults

`func NewDiskMappingRestWithDefaults() *DiskMappingRest`

NewDiskMappingRestWithDefaults instantiates a new DiskMappingRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceinfo

`func (o *DiskMappingRest) GetSourceinfo() []DiskRest`

GetSourceinfo returns the Sourceinfo field if non-nil, zero value otherwise.

### GetSourceinfoOk

`func (o *DiskMappingRest) GetSourceinfoOk() (*[]DiskRest, bool)`

GetSourceinfoOk returns a tuple with the Sourceinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceinfo

`func (o *DiskMappingRest) SetSourceinfo(v []DiskRest)`

SetSourceinfo sets Sourceinfo field to given value.

### HasSourceinfo

`func (o *DiskMappingRest) HasSourceinfo() bool`

HasSourceinfo returns a boolean if a field has been set.

### GetTargetlist

`func (o *DiskMappingRest) GetTargetlist() []DiskRest`

GetTargetlist returns the Targetlist field if non-nil, zero value otherwise.

### GetTargetlistOk

`func (o *DiskMappingRest) GetTargetlistOk() (*[]DiskRest, bool)`

GetTargetlistOk returns a tuple with the Targetlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetlist

`func (o *DiskMappingRest) SetTargetlist(v []DiskRest)`

SetTargetlist sets Targetlist field to given value.

### HasTargetlist

`func (o *DiskMappingRest) HasTargetlist() bool`

HasTargetlist returns a boolean if a field has been set.

### GetAsmracnodelist

`func (o *DiskMappingRest) GetAsmracnodelist() []string`

GetAsmracnodelist returns the Asmracnodelist field if non-nil, zero value otherwise.

### GetAsmracnodelistOk

`func (o *DiskMappingRest) GetAsmracnodelistOk() (*[]string, bool)`

GetAsmracnodelistOk returns a tuple with the Asmracnodelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmracnodelist

`func (o *DiskMappingRest) SetAsmracnodelist(v []string)`

SetAsmracnodelist sets Asmracnodelist field to given value.

### HasAsmracnodelist

`func (o *DiskMappingRest) HasAsmracnodelist() bool`

HasAsmracnodelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


