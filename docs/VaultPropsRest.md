# VaultPropsRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vaulttype** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **bool** |  | [optional] 
**Objectsize** | Pointer to **int64** |  | [optional] 
**Accesskey** | Pointer to **string** |  | [optional] 
**Accessid** | Pointer to **string** |  | [optional] 
**Baseurl** | Pointer to **string** |  | [optional] 
**Authversion** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewVaultPropsRest

`func NewVaultPropsRest() *VaultPropsRest`

NewVaultPropsRest instantiates a new VaultPropsRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultPropsRestWithDefaults

`func NewVaultPropsRestWithDefaults() *VaultPropsRest`

NewVaultPropsRestWithDefaults instantiates a new VaultPropsRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVaulttype

`func (o *VaultPropsRest) GetVaulttype() string`

GetVaulttype returns the Vaulttype field if non-nil, zero value otherwise.

### GetVaulttypeOk

`func (o *VaultPropsRest) GetVaulttypeOk() (*string, bool)`

GetVaulttypeOk returns a tuple with the Vaulttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaulttype

`func (o *VaultPropsRest) SetVaulttype(v string)`

SetVaulttype sets Vaulttype field to given value.

### HasVaulttype

`func (o *VaultPropsRest) HasVaulttype() bool`

HasVaulttype returns a boolean if a field has been set.

### GetBucket

`func (o *VaultPropsRest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *VaultPropsRest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *VaultPropsRest) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *VaultPropsRest) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCompression

`func (o *VaultPropsRest) GetCompression() bool`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *VaultPropsRest) GetCompressionOk() (*bool, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *VaultPropsRest) SetCompression(v bool)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *VaultPropsRest) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetObjectsize

`func (o *VaultPropsRest) GetObjectsize() int64`

GetObjectsize returns the Objectsize field if non-nil, zero value otherwise.

### GetObjectsizeOk

`func (o *VaultPropsRest) GetObjectsizeOk() (*int64, bool)`

GetObjectsizeOk returns a tuple with the Objectsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsize

`func (o *VaultPropsRest) SetObjectsize(v int64)`

SetObjectsize sets Objectsize field to given value.

### HasObjectsize

`func (o *VaultPropsRest) HasObjectsize() bool`

HasObjectsize returns a boolean if a field has been set.

### GetAccesskey

`func (o *VaultPropsRest) GetAccesskey() string`

GetAccesskey returns the Accesskey field if non-nil, zero value otherwise.

### GetAccesskeyOk

`func (o *VaultPropsRest) GetAccesskeyOk() (*string, bool)`

GetAccesskeyOk returns a tuple with the Accesskey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccesskey

`func (o *VaultPropsRest) SetAccesskey(v string)`

SetAccesskey sets Accesskey field to given value.

### HasAccesskey

`func (o *VaultPropsRest) HasAccesskey() bool`

HasAccesskey returns a boolean if a field has been set.

### GetAccessid

`func (o *VaultPropsRest) GetAccessid() string`

GetAccessid returns the Accessid field if non-nil, zero value otherwise.

### GetAccessidOk

`func (o *VaultPropsRest) GetAccessidOk() (*string, bool)`

GetAccessidOk returns a tuple with the Accessid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessid

`func (o *VaultPropsRest) SetAccessid(v string)`

SetAccessid sets Accessid field to given value.

### HasAccessid

`func (o *VaultPropsRest) HasAccessid() bool`

HasAccessid returns a boolean if a field has been set.

### GetBaseurl

`func (o *VaultPropsRest) GetBaseurl() string`

GetBaseurl returns the Baseurl field if non-nil, zero value otherwise.

### GetBaseurlOk

`func (o *VaultPropsRest) GetBaseurlOk() (*string, bool)`

GetBaseurlOk returns a tuple with the Baseurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseurl

`func (o *VaultPropsRest) SetBaseurl(v string)`

SetBaseurl sets Baseurl field to given value.

### HasBaseurl

`func (o *VaultPropsRest) HasBaseurl() bool`

HasBaseurl returns a boolean if a field has been set.

### GetAuthversion

`func (o *VaultPropsRest) GetAuthversion() string`

GetAuthversion returns the Authversion field if non-nil, zero value otherwise.

### GetAuthversionOk

`func (o *VaultPropsRest) GetAuthversionOk() (*string, bool)`

GetAuthversionOk returns a tuple with the Authversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthversion

`func (o *VaultPropsRest) SetAuthversion(v string)`

SetAuthversion sets Authversion field to given value.

### HasAuthversion

`func (o *VaultPropsRest) HasAuthversion() bool`

HasAuthversion returns a boolean if a field has been set.

### GetRegion

`func (o *VaultPropsRest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VaultPropsRest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VaultPropsRest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VaultPropsRest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetId

`func (o *VaultPropsRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VaultPropsRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VaultPropsRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VaultPropsRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *VaultPropsRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VaultPropsRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VaultPropsRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VaultPropsRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *VaultPropsRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *VaultPropsRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *VaultPropsRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *VaultPropsRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *VaultPropsRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *VaultPropsRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *VaultPropsRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *VaultPropsRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


