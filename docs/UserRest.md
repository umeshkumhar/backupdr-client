# UserRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Rightlist** | Pointer to [**[]RightRest**](RightRest.md) |  | [optional] 
**Rolelist** | Pointer to [**[]RoleRest**](RoleRest.md) |  | [optional] 
**Modifydate** | Pointer to **int64** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Dataaccesslevel** | Pointer to **int32** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Createdate** | Pointer to **int64** |  | [optional] 
**Lastpasswordchangedate** | Pointer to **int64** |  | [optional] 
**Userpref** | Pointer to **string** |  | [optional] 
**Localonly** | Pointer to **bool** |  | [optional] 
**Orglist** | Pointer to [**[]OrganizationRest**](OrganizationRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewUserRest

`func NewUserRest() *UserRest`

NewUserRest instantiates a new UserRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRestWithDefaults

`func NewUserRestWithDefaults() *UserRest`

NewUserRestWithDefaults instantiates a new UserRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UserRest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserRest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserRest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserRest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVersion

`func (o *UserRest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserRest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserRest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UserRest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRightlist

`func (o *UserRest) GetRightlist() []RightRest`

GetRightlist returns the Rightlist field if non-nil, zero value otherwise.

### GetRightlistOk

`func (o *UserRest) GetRightlistOk() (*[]RightRest, bool)`

GetRightlistOk returns a tuple with the Rightlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightlist

`func (o *UserRest) SetRightlist(v []RightRest)`

SetRightlist sets Rightlist field to given value.

### HasRightlist

`func (o *UserRest) HasRightlist() bool`

HasRightlist returns a boolean if a field has been set.

### GetRolelist

`func (o *UserRest) GetRolelist() []RoleRest`

GetRolelist returns the Rolelist field if non-nil, zero value otherwise.

### GetRolelistOk

`func (o *UserRest) GetRolelistOk() (*[]RoleRest, bool)`

GetRolelistOk returns a tuple with the Rolelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolelist

`func (o *UserRest) SetRolelist(v []RoleRest)`

SetRolelist sets Rolelist field to given value.

### HasRolelist

`func (o *UserRest) HasRolelist() bool`

HasRolelist returns a boolean if a field has been set.

### GetModifydate

`func (o *UserRest) GetModifydate() int64`

GetModifydate returns the Modifydate field if non-nil, zero value otherwise.

### GetModifydateOk

`func (o *UserRest) GetModifydateOk() (*int64, bool)`

GetModifydateOk returns a tuple with the Modifydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifydate

`func (o *UserRest) SetModifydate(v int64)`

SetModifydate sets Modifydate field to given value.

### HasModifydate

`func (o *UserRest) HasModifydate() bool`

HasModifydate returns a boolean if a field has been set.

### GetEmail

`func (o *UserRest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataaccesslevel

`func (o *UserRest) GetDataaccesslevel() int32`

GetDataaccesslevel returns the Dataaccesslevel field if non-nil, zero value otherwise.

### GetDataaccesslevelOk

`func (o *UserRest) GetDataaccesslevelOk() (*int32, bool)`

GetDataaccesslevelOk returns a tuple with the Dataaccesslevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataaccesslevel

`func (o *UserRest) SetDataaccesslevel(v int32)`

SetDataaccesslevel sets Dataaccesslevel field to given value.

### HasDataaccesslevel

`func (o *UserRest) HasDataaccesslevel() bool`

HasDataaccesslevel returns a boolean if a field has been set.

### GetTimezone

`func (o *UserRest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserRest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserRest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserRest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetFirstname

`func (o *UserRest) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserRest) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserRest) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserRest) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserRest) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserRest) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserRest) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserRest) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetComments

`func (o *UserRest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *UserRest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *UserRest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *UserRest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedate

`func (o *UserRest) GetCreatedate() int64`

GetCreatedate returns the Createdate field if non-nil, zero value otherwise.

### GetCreatedateOk

`func (o *UserRest) GetCreatedateOk() (*int64, bool)`

GetCreatedateOk returns a tuple with the Createdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedate

`func (o *UserRest) SetCreatedate(v int64)`

SetCreatedate sets Createdate field to given value.

### HasCreatedate

`func (o *UserRest) HasCreatedate() bool`

HasCreatedate returns a boolean if a field has been set.

### GetLastpasswordchangedate

`func (o *UserRest) GetLastpasswordchangedate() int64`

GetLastpasswordchangedate returns the Lastpasswordchangedate field if non-nil, zero value otherwise.

### GetLastpasswordchangedateOk

`func (o *UserRest) GetLastpasswordchangedateOk() (*int64, bool)`

GetLastpasswordchangedateOk returns a tuple with the Lastpasswordchangedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastpasswordchangedate

`func (o *UserRest) SetLastpasswordchangedate(v int64)`

SetLastpasswordchangedate sets Lastpasswordchangedate field to given value.

### HasLastpasswordchangedate

`func (o *UserRest) HasLastpasswordchangedate() bool`

HasLastpasswordchangedate returns a boolean if a field has been set.

### GetUserpref

`func (o *UserRest) GetUserpref() string`

GetUserpref returns the Userpref field if non-nil, zero value otherwise.

### GetUserprefOk

`func (o *UserRest) GetUserprefOk() (*string, bool)`

GetUserprefOk returns a tuple with the Userpref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserpref

`func (o *UserRest) SetUserpref(v string)`

SetUserpref sets Userpref field to given value.

### HasUserpref

`func (o *UserRest) HasUserpref() bool`

HasUserpref returns a boolean if a field has been set.

### GetLocalonly

`func (o *UserRest) GetLocalonly() bool`

GetLocalonly returns the Localonly field if non-nil, zero value otherwise.

### GetLocalonlyOk

`func (o *UserRest) GetLocalonlyOk() (*bool, bool)`

GetLocalonlyOk returns a tuple with the Localonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalonly

`func (o *UserRest) SetLocalonly(v bool)`

SetLocalonly sets Localonly field to given value.

### HasLocalonly

`func (o *UserRest) HasLocalonly() bool`

HasLocalonly returns a boolean if a field has been set.

### GetOrglist

`func (o *UserRest) GetOrglist() []OrganizationRest`

GetOrglist returns the Orglist field if non-nil, zero value otherwise.

### GetOrglistOk

`func (o *UserRest) GetOrglistOk() (*[]OrganizationRest, bool)`

GetOrglistOk returns a tuple with the Orglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrglist

`func (o *UserRest) SetOrglist(v []OrganizationRest)`

SetOrglist sets Orglist field to given value.

### HasOrglist

`func (o *UserRest) HasOrglist() bool`

HasOrglist returns a boolean if a field has been set.

### GetId

`func (o *UserRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *UserRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UserRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UserRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *UserRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *UserRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *UserRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *UserRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *UserRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *UserRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *UserRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *UserRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *UserRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


