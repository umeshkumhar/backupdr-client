# CollectionRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Userlist** | Pointer to **[]string** |  | [optional] 
**Sltlist** | Pointer to **[]string** |  | [optional] 
**Slplist** | Pointer to **[]string** |  | [optional] 
**Hostlist** | Pointer to **[]string** |  | [optional] 
**Poollist** | Pointer to **[]string** |  | [optional] 
**Applist** | Pointer to **[]string** |  | [optional] 
**Arraylist** | Pointer to **[]string** |  | [optional] 
**Lglist** | Pointer to **[]string** |  | [optional] 
**Cloudcredentiallist** | Pointer to **[]string** |  | [optional] 
**Sltlistcount** | Pointer to **int64** |  | [optional] 
**Hostlistcount** | Pointer to **int64** |  | [optional] 
**Slplistcount** | Pointer to **int64** |  | [optional] 
**Userlistcount** | Pointer to **int64** |  | [optional] 
**Poollistcount** | Pointer to **int64** |  | [optional] 
**Applistcount** | Pointer to **int64** |  | [optional] 
**Arraylistcount** | Pointer to **int64** |  | [optional] 
**Lglistcount** | Pointer to **int64** |  | [optional] 
**Cloudcredentiallistcount** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewCollectionRest

`func NewCollectionRest() *CollectionRest`

NewCollectionRest instantiates a new CollectionRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionRestWithDefaults

`func NewCollectionRestWithDefaults() *CollectionRest`

NewCollectionRestWithDefaults instantiates a new CollectionRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserlist

`func (o *CollectionRest) GetUserlist() []string`

GetUserlist returns the Userlist field if non-nil, zero value otherwise.

### GetUserlistOk

`func (o *CollectionRest) GetUserlistOk() (*[]string, bool)`

GetUserlistOk returns a tuple with the Userlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserlist

`func (o *CollectionRest) SetUserlist(v []string)`

SetUserlist sets Userlist field to given value.

### HasUserlist

`func (o *CollectionRest) HasUserlist() bool`

HasUserlist returns a boolean if a field has been set.

### GetSltlist

`func (o *CollectionRest) GetSltlist() []string`

GetSltlist returns the Sltlist field if non-nil, zero value otherwise.

### GetSltlistOk

`func (o *CollectionRest) GetSltlistOk() (*[]string, bool)`

GetSltlistOk returns a tuple with the Sltlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSltlist

`func (o *CollectionRest) SetSltlist(v []string)`

SetSltlist sets Sltlist field to given value.

### HasSltlist

`func (o *CollectionRest) HasSltlist() bool`

HasSltlist returns a boolean if a field has been set.

### GetSlplist

`func (o *CollectionRest) GetSlplist() []string`

GetSlplist returns the Slplist field if non-nil, zero value otherwise.

### GetSlplistOk

`func (o *CollectionRest) GetSlplistOk() (*[]string, bool)`

GetSlplistOk returns a tuple with the Slplist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlplist

`func (o *CollectionRest) SetSlplist(v []string)`

SetSlplist sets Slplist field to given value.

### HasSlplist

`func (o *CollectionRest) HasSlplist() bool`

HasSlplist returns a boolean if a field has been set.

### GetHostlist

`func (o *CollectionRest) GetHostlist() []string`

GetHostlist returns the Hostlist field if non-nil, zero value otherwise.

### GetHostlistOk

`func (o *CollectionRest) GetHostlistOk() (*[]string, bool)`

GetHostlistOk returns a tuple with the Hostlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostlist

`func (o *CollectionRest) SetHostlist(v []string)`

SetHostlist sets Hostlist field to given value.

### HasHostlist

`func (o *CollectionRest) HasHostlist() bool`

HasHostlist returns a boolean if a field has been set.

### GetPoollist

`func (o *CollectionRest) GetPoollist() []string`

GetPoollist returns the Poollist field if non-nil, zero value otherwise.

### GetPoollistOk

`func (o *CollectionRest) GetPoollistOk() (*[]string, bool)`

GetPoollistOk returns a tuple with the Poollist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoollist

`func (o *CollectionRest) SetPoollist(v []string)`

SetPoollist sets Poollist field to given value.

### HasPoollist

`func (o *CollectionRest) HasPoollist() bool`

HasPoollist returns a boolean if a field has been set.

### GetApplist

`func (o *CollectionRest) GetApplist() []string`

GetApplist returns the Applist field if non-nil, zero value otherwise.

### GetApplistOk

`func (o *CollectionRest) GetApplistOk() (*[]string, bool)`

GetApplistOk returns a tuple with the Applist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplist

`func (o *CollectionRest) SetApplist(v []string)`

SetApplist sets Applist field to given value.

### HasApplist

`func (o *CollectionRest) HasApplist() bool`

HasApplist returns a boolean if a field has been set.

### GetArraylist

`func (o *CollectionRest) GetArraylist() []string`

GetArraylist returns the Arraylist field if non-nil, zero value otherwise.

### GetArraylistOk

`func (o *CollectionRest) GetArraylistOk() (*[]string, bool)`

GetArraylistOk returns a tuple with the Arraylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArraylist

`func (o *CollectionRest) SetArraylist(v []string)`

SetArraylist sets Arraylist field to given value.

### HasArraylist

`func (o *CollectionRest) HasArraylist() bool`

HasArraylist returns a boolean if a field has been set.

### GetLglist

`func (o *CollectionRest) GetLglist() []string`

GetLglist returns the Lglist field if non-nil, zero value otherwise.

### GetLglistOk

`func (o *CollectionRest) GetLglistOk() (*[]string, bool)`

GetLglistOk returns a tuple with the Lglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLglist

`func (o *CollectionRest) SetLglist(v []string)`

SetLglist sets Lglist field to given value.

### HasLglist

`func (o *CollectionRest) HasLglist() bool`

HasLglist returns a boolean if a field has been set.

### GetCloudcredentiallist

`func (o *CollectionRest) GetCloudcredentiallist() []string`

GetCloudcredentiallist returns the Cloudcredentiallist field if non-nil, zero value otherwise.

### GetCloudcredentiallistOk

`func (o *CollectionRest) GetCloudcredentiallistOk() (*[]string, bool)`

GetCloudcredentiallistOk returns a tuple with the Cloudcredentiallist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudcredentiallist

`func (o *CollectionRest) SetCloudcredentiallist(v []string)`

SetCloudcredentiallist sets Cloudcredentiallist field to given value.

### HasCloudcredentiallist

`func (o *CollectionRest) HasCloudcredentiallist() bool`

HasCloudcredentiallist returns a boolean if a field has been set.

### GetSltlistcount

`func (o *CollectionRest) GetSltlistcount() int64`

GetSltlistcount returns the Sltlistcount field if non-nil, zero value otherwise.

### GetSltlistcountOk

`func (o *CollectionRest) GetSltlistcountOk() (*int64, bool)`

GetSltlistcountOk returns a tuple with the Sltlistcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSltlistcount

`func (o *CollectionRest) SetSltlistcount(v int64)`

SetSltlistcount sets Sltlistcount field to given value.

### HasSltlistcount

`func (o *CollectionRest) HasSltlistcount() bool`

HasSltlistcount returns a boolean if a field has been set.

### GetHostlistcount

`func (o *CollectionRest) GetHostlistcount() int64`

GetHostlistcount returns the Hostlistcount field if non-nil, zero value otherwise.

### GetHostlistcountOk

`func (o *CollectionRest) GetHostlistcountOk() (*int64, bool)`

GetHostlistcountOk returns a tuple with the Hostlistcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostlistcount

`func (o *CollectionRest) SetHostlistcount(v int64)`

SetHostlistcount sets Hostlistcount field to given value.

### HasHostlistcount

`func (o *CollectionRest) HasHostlistcount() bool`

HasHostlistcount returns a boolean if a field has been set.

### GetSlplistcount

`func (o *CollectionRest) GetSlplistcount() int64`

GetSlplistcount returns the Slplistcount field if non-nil, zero value otherwise.

### GetSlplistcountOk

`func (o *CollectionRest) GetSlplistcountOk() (*int64, bool)`

GetSlplistcountOk returns a tuple with the Slplistcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlplistcount

`func (o *CollectionRest) SetSlplistcount(v int64)`

SetSlplistcount sets Slplistcount field to given value.

### HasSlplistcount

`func (o *CollectionRest) HasSlplistcount() bool`

HasSlplistcount returns a boolean if a field has been set.

### GetUserlistcount

`func (o *CollectionRest) GetUserlistcount() int64`

GetUserlistcount returns the Userlistcount field if non-nil, zero value otherwise.

### GetUserlistcountOk

`func (o *CollectionRest) GetUserlistcountOk() (*int64, bool)`

GetUserlistcountOk returns a tuple with the Userlistcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserlistcount

`func (o *CollectionRest) SetUserlistcount(v int64)`

SetUserlistcount sets Userlistcount field to given value.

### HasUserlistcount

`func (o *CollectionRest) HasUserlistcount() bool`

HasUserlistcount returns a boolean if a field has been set.

### GetPoollistcount

`func (o *CollectionRest) GetPoollistcount() int64`

GetPoollistcount returns the Poollistcount field if non-nil, zero value otherwise.

### GetPoollistcountOk

`func (o *CollectionRest) GetPoollistcountOk() (*int64, bool)`

GetPoollistcountOk returns a tuple with the Poollistcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoollistcount

`func (o *CollectionRest) SetPoollistcount(v int64)`

SetPoollistcount sets Poollistcount field to given value.

### HasPoollistcount

`func (o *CollectionRest) HasPoollistcount() bool`

HasPoollistcount returns a boolean if a field has been set.

### GetApplistcount

`func (o *CollectionRest) GetApplistcount() int64`

GetApplistcount returns the Applistcount field if non-nil, zero value otherwise.

### GetApplistcountOk

`func (o *CollectionRest) GetApplistcountOk() (*int64, bool)`

GetApplistcountOk returns a tuple with the Applistcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplistcount

`func (o *CollectionRest) SetApplistcount(v int64)`

SetApplistcount sets Applistcount field to given value.

### HasApplistcount

`func (o *CollectionRest) HasApplistcount() bool`

HasApplistcount returns a boolean if a field has been set.

### GetArraylistcount

`func (o *CollectionRest) GetArraylistcount() int64`

GetArraylistcount returns the Arraylistcount field if non-nil, zero value otherwise.

### GetArraylistcountOk

`func (o *CollectionRest) GetArraylistcountOk() (*int64, bool)`

GetArraylistcountOk returns a tuple with the Arraylistcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArraylistcount

`func (o *CollectionRest) SetArraylistcount(v int64)`

SetArraylistcount sets Arraylistcount field to given value.

### HasArraylistcount

`func (o *CollectionRest) HasArraylistcount() bool`

HasArraylistcount returns a boolean if a field has been set.

### GetLglistcount

`func (o *CollectionRest) GetLglistcount() int64`

GetLglistcount returns the Lglistcount field if non-nil, zero value otherwise.

### GetLglistcountOk

`func (o *CollectionRest) GetLglistcountOk() (*int64, bool)`

GetLglistcountOk returns a tuple with the Lglistcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLglistcount

`func (o *CollectionRest) SetLglistcount(v int64)`

SetLglistcount sets Lglistcount field to given value.

### HasLglistcount

`func (o *CollectionRest) HasLglistcount() bool`

HasLglistcount returns a boolean if a field has been set.

### GetCloudcredentiallistcount

`func (o *CollectionRest) GetCloudcredentiallistcount() int64`

GetCloudcredentiallistcount returns the Cloudcredentiallistcount field if non-nil, zero value otherwise.

### GetCloudcredentiallistcountOk

`func (o *CollectionRest) GetCloudcredentiallistcountOk() (*int64, bool)`

GetCloudcredentiallistcountOk returns a tuple with the Cloudcredentiallistcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudcredentiallistcount

`func (o *CollectionRest) SetCloudcredentiallistcount(v int64)`

SetCloudcredentiallistcount sets Cloudcredentiallistcount field to given value.

### HasCloudcredentiallistcount

`func (o *CollectionRest) HasCloudcredentiallistcount() bool`

HasCloudcredentiallistcount returns a boolean if a field has been set.

### GetId

`func (o *CollectionRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CollectionRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *CollectionRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CollectionRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CollectionRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CollectionRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *CollectionRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *CollectionRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *CollectionRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *CollectionRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *CollectionRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *CollectionRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *CollectionRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *CollectionRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


