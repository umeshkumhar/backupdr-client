# SourceRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slps** | Pointer to [**[]SlpRest**](SlpRest.md) |  | [optional] 
**Host** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Srcid** | Pointer to **string** |  | [optional] 
**Clusterid** | Pointer to **string** |  | [optional] 

## Methods

### NewSourceRest

`func NewSourceRest() *SourceRest`

NewSourceRest instantiates a new SourceRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceRestWithDefaults

`func NewSourceRestWithDefaults() *SourceRest`

NewSourceRestWithDefaults instantiates a new SourceRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlps

`func (o *SourceRest) GetSlps() []SlpRest`

GetSlps returns the Slps field if non-nil, zero value otherwise.

### GetSlpsOk

`func (o *SourceRest) GetSlpsOk() (*[]SlpRest, bool)`

GetSlpsOk returns a tuple with the Slps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlps

`func (o *SourceRest) SetSlps(v []SlpRest)`

SetSlps sets Slps field to given value.

### HasSlps

`func (o *SourceRest) HasSlps() bool`

HasSlps returns a boolean if a field has been set.

### GetHost

`func (o *SourceRest) GetHost() HostRest`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SourceRest) GetHostOk() (*HostRest, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SourceRest) SetHost(v HostRest)`

SetHost sets Host field to given value.

### HasHost

`func (o *SourceRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetSrcid

`func (o *SourceRest) GetSrcid() string`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *SourceRest) GetSrcidOk() (*string, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *SourceRest) SetSrcid(v string)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *SourceRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetClusterid

`func (o *SourceRest) GetClusterid() string`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *SourceRest) GetClusteridOk() (*string, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *SourceRest) SetClusterid(v string)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *SourceRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


