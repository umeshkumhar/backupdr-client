# EventRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **map[string]interface{}** |  | [optional] 
**Errorcode** | Pointer to **string** |  | [optional] 
**Srcid** | Pointer to **string** |  | [optional] 
**Component** | Pointer to **map[string]interface{}** |  | [optional] 
**Appname** | Pointer to **string** |  | [optional] 
**Apptype** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to [**ClusterRest**](ClusterRest.md) |  | [optional] 
**Eventdate** | Pointer to **int64** |  | [optional] 
**Eventid** | Pointer to **int32** |  | [optional] 
**Sourceeventdate** | Pointer to **string** |  | [optional] 
**Jobname** | Pointer to **string** |  | [optional] 
**Errormessage** | Pointer to **string** |  | [optional] 
**Messagetext** | Pointer to **string** |  | [optional] 
**Sequenceid** | Pointer to **int64** |  | [optional] 
**Requiresclearing** | Pointer to **bool** |  | [optional] 
**Eventstatus** | Pointer to **string** |  | [optional] 
**Notification** | Pointer to **map[string]interface{}** |  | [optional] 
**Errormessages** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewEventRest

`func NewEventRest() *EventRest`

NewEventRest instantiates a new EventRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRestWithDefaults

`func NewEventRestWithDefaults() *EventRest`

NewEventRestWithDefaults instantiates a new EventRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *EventRest) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EventRest) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EventRest) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *EventRest) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetErrorcode

`func (o *EventRest) GetErrorcode() string`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *EventRest) GetErrorcodeOk() (*string, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *EventRest) SetErrorcode(v string)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *EventRest) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetSrcid

`func (o *EventRest) GetSrcid() string`

GetSrcid returns the Srcid field if non-nil, zero value otherwise.

### GetSrcidOk

`func (o *EventRest) GetSrcidOk() (*string, bool)`

GetSrcidOk returns a tuple with the Srcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcid

`func (o *EventRest) SetSrcid(v string)`

SetSrcid sets Srcid field to given value.

### HasSrcid

`func (o *EventRest) HasSrcid() bool`

HasSrcid returns a boolean if a field has been set.

### GetComponent

`func (o *EventRest) GetComponent() map[string]interface{}`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *EventRest) GetComponentOk() (*map[string]interface{}, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *EventRest) SetComponent(v map[string]interface{})`

SetComponent sets Component field to given value.

### HasComponent

`func (o *EventRest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetAppname

`func (o *EventRest) GetAppname() string`

GetAppname returns the Appname field if non-nil, zero value otherwise.

### GetAppnameOk

`func (o *EventRest) GetAppnameOk() (*string, bool)`

GetAppnameOk returns a tuple with the Appname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppname

`func (o *EventRest) SetAppname(v string)`

SetAppname sets Appname field to given value.

### HasAppname

`func (o *EventRest) HasAppname() bool`

HasAppname returns a boolean if a field has been set.

### GetApptype

`func (o *EventRest) GetApptype() string`

GetApptype returns the Apptype field if non-nil, zero value otherwise.

### GetApptypeOk

`func (o *EventRest) GetApptypeOk() (*string, bool)`

GetApptypeOk returns a tuple with the Apptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApptype

`func (o *EventRest) SetApptype(v string)`

SetApptype sets Apptype field to given value.

### HasApptype

`func (o *EventRest) HasApptype() bool`

HasApptype returns a boolean if a field has been set.

### GetCluster

`func (o *EventRest) GetCluster() ClusterRest`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *EventRest) GetClusterOk() (*ClusterRest, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *EventRest) SetCluster(v ClusterRest)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *EventRest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetEventdate

`func (o *EventRest) GetEventdate() int64`

GetEventdate returns the Eventdate field if non-nil, zero value otherwise.

### GetEventdateOk

`func (o *EventRest) GetEventdateOk() (*int64, bool)`

GetEventdateOk returns a tuple with the Eventdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventdate

`func (o *EventRest) SetEventdate(v int64)`

SetEventdate sets Eventdate field to given value.

### HasEventdate

`func (o *EventRest) HasEventdate() bool`

HasEventdate returns a boolean if a field has been set.

### GetEventid

`func (o *EventRest) GetEventid() int32`

GetEventid returns the Eventid field if non-nil, zero value otherwise.

### GetEventidOk

`func (o *EventRest) GetEventidOk() (*int32, bool)`

GetEventidOk returns a tuple with the Eventid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventid

`func (o *EventRest) SetEventid(v int32)`

SetEventid sets Eventid field to given value.

### HasEventid

`func (o *EventRest) HasEventid() bool`

HasEventid returns a boolean if a field has been set.

### GetSourceeventdate

`func (o *EventRest) GetSourceeventdate() string`

GetSourceeventdate returns the Sourceeventdate field if non-nil, zero value otherwise.

### GetSourceeventdateOk

`func (o *EventRest) GetSourceeventdateOk() (*string, bool)`

GetSourceeventdateOk returns a tuple with the Sourceeventdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceeventdate

`func (o *EventRest) SetSourceeventdate(v string)`

SetSourceeventdate sets Sourceeventdate field to given value.

### HasSourceeventdate

`func (o *EventRest) HasSourceeventdate() bool`

HasSourceeventdate returns a boolean if a field has been set.

### GetJobname

`func (o *EventRest) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *EventRest) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *EventRest) SetJobname(v string)`

SetJobname sets Jobname field to given value.

### HasJobname

`func (o *EventRest) HasJobname() bool`

HasJobname returns a boolean if a field has been set.

### GetErrormessage

`func (o *EventRest) GetErrormessage() string`

GetErrormessage returns the Errormessage field if non-nil, zero value otherwise.

### GetErrormessageOk

`func (o *EventRest) GetErrormessageOk() (*string, bool)`

GetErrormessageOk returns a tuple with the Errormessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrormessage

`func (o *EventRest) SetErrormessage(v string)`

SetErrormessage sets Errormessage field to given value.

### HasErrormessage

`func (o *EventRest) HasErrormessage() bool`

HasErrormessage returns a boolean if a field has been set.

### GetMessagetext

`func (o *EventRest) GetMessagetext() string`

GetMessagetext returns the Messagetext field if non-nil, zero value otherwise.

### GetMessagetextOk

`func (o *EventRest) GetMessagetextOk() (*string, bool)`

GetMessagetextOk returns a tuple with the Messagetext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagetext

`func (o *EventRest) SetMessagetext(v string)`

SetMessagetext sets Messagetext field to given value.

### HasMessagetext

`func (o *EventRest) HasMessagetext() bool`

HasMessagetext returns a boolean if a field has been set.

### GetSequenceid

`func (o *EventRest) GetSequenceid() int64`

GetSequenceid returns the Sequenceid field if non-nil, zero value otherwise.

### GetSequenceidOk

`func (o *EventRest) GetSequenceidOk() (*int64, bool)`

GetSequenceidOk returns a tuple with the Sequenceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceid

`func (o *EventRest) SetSequenceid(v int64)`

SetSequenceid sets Sequenceid field to given value.

### HasSequenceid

`func (o *EventRest) HasSequenceid() bool`

HasSequenceid returns a boolean if a field has been set.

### GetRequiresclearing

`func (o *EventRest) GetRequiresclearing() bool`

GetRequiresclearing returns the Requiresclearing field if non-nil, zero value otherwise.

### GetRequiresclearingOk

`func (o *EventRest) GetRequiresclearingOk() (*bool, bool)`

GetRequiresclearingOk returns a tuple with the Requiresclearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresclearing

`func (o *EventRest) SetRequiresclearing(v bool)`

SetRequiresclearing sets Requiresclearing field to given value.

### HasRequiresclearing

`func (o *EventRest) HasRequiresclearing() bool`

HasRequiresclearing returns a boolean if a field has been set.

### GetEventstatus

`func (o *EventRest) GetEventstatus() string`

GetEventstatus returns the Eventstatus field if non-nil, zero value otherwise.

### GetEventstatusOk

`func (o *EventRest) GetEventstatusOk() (*string, bool)`

GetEventstatusOk returns a tuple with the Eventstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventstatus

`func (o *EventRest) SetEventstatus(v string)`

SetEventstatus sets Eventstatus field to given value.

### HasEventstatus

`func (o *EventRest) HasEventstatus() bool`

HasEventstatus returns a boolean if a field has been set.

### GetNotification

`func (o *EventRest) GetNotification() map[string]interface{}`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *EventRest) GetNotificationOk() (*map[string]interface{}, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *EventRest) SetNotification(v map[string]interface{})`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *EventRest) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetErrormessages

`func (o *EventRest) GetErrormessages() []string`

GetErrormessages returns the Errormessages field if non-nil, zero value otherwise.

### GetErrormessagesOk

`func (o *EventRest) GetErrormessagesOk() (*[]string, bool)`

GetErrormessagesOk returns a tuple with the Errormessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrormessages

`func (o *EventRest) SetErrormessages(v []string)`

SetErrormessages sets Errormessages field to given value.

### HasErrormessages

`func (o *EventRest) HasErrormessages() bool`

HasErrormessages returns a boolean if a field has been set.

### GetId

`func (o *EventRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *EventRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *EventRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *EventRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *EventRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *EventRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *EventRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *EventRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *EventRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *EventRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *EventRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *EventRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *EventRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


