# AuditRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Jobname** | Pointer to **string** |  | [optional] 
**Issuedate** | Pointer to **int64** |  | [optional] 
**Proxy** | Pointer to **string** |  | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewAuditRest

`func NewAuditRest() *AuditRest`

NewAuditRest instantiates a new AuditRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditRestWithDefaults

`func NewAuditRestWithDefaults() *AuditRest`

NewAuditRestWithDefaults instantiates a new AuditRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *AuditRest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuditRest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuditRest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuditRest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCommand

`func (o *AuditRest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *AuditRest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *AuditRest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *AuditRest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetUsername

`func (o *AuditRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuditRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuditRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AuditRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIpaddress

`func (o *AuditRest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *AuditRest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *AuditRest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *AuditRest) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetStatus

`func (o *AuditRest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuditRest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuditRest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuditRest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJobname

`func (o *AuditRest) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *AuditRest) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *AuditRest) SetJobname(v string)`

SetJobname sets Jobname field to given value.

### HasJobname

`func (o *AuditRest) HasJobname() bool`

HasJobname returns a boolean if a field has been set.

### GetIssuedate

`func (o *AuditRest) GetIssuedate() int64`

GetIssuedate returns the Issuedate field if non-nil, zero value otherwise.

### GetIssuedateOk

`func (o *AuditRest) GetIssuedateOk() (*int64, bool)`

GetIssuedateOk returns a tuple with the Issuedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedate

`func (o *AuditRest) SetIssuedate(v int64)`

SetIssuedate sets Issuedate field to given value.

### HasIssuedate

`func (o *AuditRest) HasIssuedate() bool`

HasIssuedate returns a boolean if a field has been set.

### GetProxy

`func (o *AuditRest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AuditRest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AuditRest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AuditRest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetPrivileged

`func (o *AuditRest) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *AuditRest) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *AuditRest) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *AuditRest) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetId

`func (o *AuditRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *AuditRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AuditRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AuditRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AuditRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *AuditRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *AuditRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *AuditRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *AuditRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *AuditRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *AuditRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *AuditRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *AuditRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


