# ApplianceUpdateInstallationJobRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **int64** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**Clusterid** | Pointer to **int64** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Appliancename** | Pointer to **string** |  | [optional] 
**Updatestatus** | Pointer to **string** |  | [optional] 
**Installduration** | Pointer to **int64** |  | [optional] 
**Logslink** | Pointer to **string** |  | [optional] 
**Updateprogress** | Pointer to **int32** |  | [optional] 
**Updateinformation** | Pointer to [**UpdateInformationRest**](UpdateInformationRest.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewApplianceUpdateInstallationJobRest

`func NewApplianceUpdateInstallationJobRest() *ApplianceUpdateInstallationJobRest`

NewApplianceUpdateInstallationJobRest instantiates a new ApplianceUpdateInstallationJobRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceUpdateInstallationJobRestWithDefaults

`func NewApplianceUpdateInstallationJobRestWithDefaults() *ApplianceUpdateInstallationJobRest`

NewApplianceUpdateInstallationJobRestWithDefaults instantiates a new ApplianceUpdateInstallationJobRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *ApplianceUpdateInstallationJobRest) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceUpdateInstallationJobRest) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceUpdateInstallationJobRest) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceUpdateInstallationJobRest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceUpdateInstallationJobRest) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceUpdateInstallationJobRest) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceUpdateInstallationJobRest) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceUpdateInstallationJobRest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetClusterid

`func (o *ApplianceUpdateInstallationJobRest) GetClusterid() int64`

GetClusterid returns the Clusterid field if non-nil, zero value otherwise.

### GetClusteridOk

`func (o *ApplianceUpdateInstallationJobRest) GetClusteridOk() (*int64, bool)`

GetClusteridOk returns a tuple with the Clusterid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterid

`func (o *ApplianceUpdateInstallationJobRest) SetClusterid(v int64)`

SetClusterid sets Clusterid field to given value.

### HasClusterid

`func (o *ApplianceUpdateInstallationJobRest) HasClusterid() bool`

HasClusterid returns a boolean if a field has been set.

### GetDetails

`func (o *ApplianceUpdateInstallationJobRest) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ApplianceUpdateInstallationJobRest) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ApplianceUpdateInstallationJobRest) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ApplianceUpdateInstallationJobRest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetAppliancename

`func (o *ApplianceUpdateInstallationJobRest) GetAppliancename() string`

GetAppliancename returns the Appliancename field if non-nil, zero value otherwise.

### GetAppliancenameOk

`func (o *ApplianceUpdateInstallationJobRest) GetAppliancenameOk() (*string, bool)`

GetAppliancenameOk returns a tuple with the Appliancename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliancename

`func (o *ApplianceUpdateInstallationJobRest) SetAppliancename(v string)`

SetAppliancename sets Appliancename field to given value.

### HasAppliancename

`func (o *ApplianceUpdateInstallationJobRest) HasAppliancename() bool`

HasAppliancename returns a boolean if a field has been set.

### GetUpdatestatus

`func (o *ApplianceUpdateInstallationJobRest) GetUpdatestatus() string`

GetUpdatestatus returns the Updatestatus field if non-nil, zero value otherwise.

### GetUpdatestatusOk

`func (o *ApplianceUpdateInstallationJobRest) GetUpdatestatusOk() (*string, bool)`

GetUpdatestatusOk returns a tuple with the Updatestatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatestatus

`func (o *ApplianceUpdateInstallationJobRest) SetUpdatestatus(v string)`

SetUpdatestatus sets Updatestatus field to given value.

### HasUpdatestatus

`func (o *ApplianceUpdateInstallationJobRest) HasUpdatestatus() bool`

HasUpdatestatus returns a boolean if a field has been set.

### GetInstallduration

`func (o *ApplianceUpdateInstallationJobRest) GetInstallduration() int64`

GetInstallduration returns the Installduration field if non-nil, zero value otherwise.

### GetInstalldurationOk

`func (o *ApplianceUpdateInstallationJobRest) GetInstalldurationOk() (*int64, bool)`

GetInstalldurationOk returns a tuple with the Installduration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallduration

`func (o *ApplianceUpdateInstallationJobRest) SetInstallduration(v int64)`

SetInstallduration sets Installduration field to given value.

### HasInstallduration

`func (o *ApplianceUpdateInstallationJobRest) HasInstallduration() bool`

HasInstallduration returns a boolean if a field has been set.

### GetLogslink

`func (o *ApplianceUpdateInstallationJobRest) GetLogslink() string`

GetLogslink returns the Logslink field if non-nil, zero value otherwise.

### GetLogslinkOk

`func (o *ApplianceUpdateInstallationJobRest) GetLogslinkOk() (*string, bool)`

GetLogslinkOk returns a tuple with the Logslink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogslink

`func (o *ApplianceUpdateInstallationJobRest) SetLogslink(v string)`

SetLogslink sets Logslink field to given value.

### HasLogslink

`func (o *ApplianceUpdateInstallationJobRest) HasLogslink() bool`

HasLogslink returns a boolean if a field has been set.

### GetUpdateprogress

`func (o *ApplianceUpdateInstallationJobRest) GetUpdateprogress() int32`

GetUpdateprogress returns the Updateprogress field if non-nil, zero value otherwise.

### GetUpdateprogressOk

`func (o *ApplianceUpdateInstallationJobRest) GetUpdateprogressOk() (*int32, bool)`

GetUpdateprogressOk returns a tuple with the Updateprogress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateprogress

`func (o *ApplianceUpdateInstallationJobRest) SetUpdateprogress(v int32)`

SetUpdateprogress sets Updateprogress field to given value.

### HasUpdateprogress

`func (o *ApplianceUpdateInstallationJobRest) HasUpdateprogress() bool`

HasUpdateprogress returns a boolean if a field has been set.

### GetUpdateinformation

`func (o *ApplianceUpdateInstallationJobRest) GetUpdateinformation() UpdateInformationRest`

GetUpdateinformation returns the Updateinformation field if non-nil, zero value otherwise.

### GetUpdateinformationOk

`func (o *ApplianceUpdateInstallationJobRest) GetUpdateinformationOk() (*UpdateInformationRest, bool)`

GetUpdateinformationOk returns a tuple with the Updateinformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateinformation

`func (o *ApplianceUpdateInstallationJobRest) SetUpdateinformation(v UpdateInformationRest)`

SetUpdateinformation sets Updateinformation field to given value.

### HasUpdateinformation

`func (o *ApplianceUpdateInstallationJobRest) HasUpdateinformation() bool`

HasUpdateinformation returns a boolean if a field has been set.

### GetId

`func (o *ApplianceUpdateInstallationJobRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplianceUpdateInstallationJobRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplianceUpdateInstallationJobRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplianceUpdateInstallationJobRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *ApplianceUpdateInstallationJobRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplianceUpdateInstallationJobRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplianceUpdateInstallationJobRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplianceUpdateInstallationJobRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *ApplianceUpdateInstallationJobRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *ApplianceUpdateInstallationJobRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *ApplianceUpdateInstallationJobRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *ApplianceUpdateInstallationJobRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *ApplianceUpdateInstallationJobRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ApplianceUpdateInstallationJobRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ApplianceUpdateInstallationJobRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ApplianceUpdateInstallationJobRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


