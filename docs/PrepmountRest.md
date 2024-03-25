# PrepmountRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to [**HostRest**](HostRest.md) |  | [optional] 
**Rdmmode** | Pointer to **string** |  | [optional] 
**Physicalrdm** | Pointer to **string** |  | [optional] 
**Appaware** | Pointer to **bool** |  | [optional] 
**Recoverytime** | Pointer to **int64** |  | [optional] 
**Scripts** | Pointer to [**[]ScriptRest**](ScriptRest.md) |  | [optional] 
**Restoreobjectmappings** | Pointer to [**[]RestoreObjectMappingRest**](RestoreObjectMappingRest.md) |  | [optional] 
**Restoreoptions** | Pointer to [**[]RestoreOptionRest**](RestoreOptionRest.md) |  | [optional] 
**Provisioningoptions** | Pointer to [**[]ProvisioningOptionRest**](ProvisioningOptionRest.md) |  | [optional] 
**Selectedobjects** | Pointer to [**[]SelectedObjectRest**](SelectedObjectRest.md) |  | [optional] 
**Nfsoptions** | Pointer to [**NfsOptionsRest**](NfsOptionsRest.md) |  | [optional] 
**Instantmount** | Pointer to **bool** |  | [optional] 
**Container** | Pointer to **bool** |  | [optional] 
**Allowedips** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewPrepmountRest

`func NewPrepmountRest() *PrepmountRest`

NewPrepmountRest instantiates a new PrepmountRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepmountRestWithDefaults

`func NewPrepmountRestWithDefaults() *PrepmountRest`

NewPrepmountRestWithDefaults instantiates a new PrepmountRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PrepmountRest) GetHost() HostRest`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PrepmountRest) GetHostOk() (*HostRest, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PrepmountRest) SetHost(v HostRest)`

SetHost sets Host field to given value.

### HasHost

`func (o *PrepmountRest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetRdmmode

`func (o *PrepmountRest) GetRdmmode() string`

GetRdmmode returns the Rdmmode field if non-nil, zero value otherwise.

### GetRdmmodeOk

`func (o *PrepmountRest) GetRdmmodeOk() (*string, bool)`

GetRdmmodeOk returns a tuple with the Rdmmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdmmode

`func (o *PrepmountRest) SetRdmmode(v string)`

SetRdmmode sets Rdmmode field to given value.

### HasRdmmode

`func (o *PrepmountRest) HasRdmmode() bool`

HasRdmmode returns a boolean if a field has been set.

### GetPhysicalrdm

`func (o *PrepmountRest) GetPhysicalrdm() string`

GetPhysicalrdm returns the Physicalrdm field if non-nil, zero value otherwise.

### GetPhysicalrdmOk

`func (o *PrepmountRest) GetPhysicalrdmOk() (*string, bool)`

GetPhysicalrdmOk returns a tuple with the Physicalrdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalrdm

`func (o *PrepmountRest) SetPhysicalrdm(v string)`

SetPhysicalrdm sets Physicalrdm field to given value.

### HasPhysicalrdm

`func (o *PrepmountRest) HasPhysicalrdm() bool`

HasPhysicalrdm returns a boolean if a field has been set.

### GetAppaware

`func (o *PrepmountRest) GetAppaware() bool`

GetAppaware returns the Appaware field if non-nil, zero value otherwise.

### GetAppawareOk

`func (o *PrepmountRest) GetAppawareOk() (*bool, bool)`

GetAppawareOk returns a tuple with the Appaware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppaware

`func (o *PrepmountRest) SetAppaware(v bool)`

SetAppaware sets Appaware field to given value.

### HasAppaware

`func (o *PrepmountRest) HasAppaware() bool`

HasAppaware returns a boolean if a field has been set.

### GetRecoverytime

`func (o *PrepmountRest) GetRecoverytime() int64`

GetRecoverytime returns the Recoverytime field if non-nil, zero value otherwise.

### GetRecoverytimeOk

`func (o *PrepmountRest) GetRecoverytimeOk() (*int64, bool)`

GetRecoverytimeOk returns a tuple with the Recoverytime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverytime

`func (o *PrepmountRest) SetRecoverytime(v int64)`

SetRecoverytime sets Recoverytime field to given value.

### HasRecoverytime

`func (o *PrepmountRest) HasRecoverytime() bool`

HasRecoverytime returns a boolean if a field has been set.

### GetScripts

`func (o *PrepmountRest) GetScripts() []ScriptRest`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *PrepmountRest) GetScriptsOk() (*[]ScriptRest, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *PrepmountRest) SetScripts(v []ScriptRest)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *PrepmountRest) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetRestoreobjectmappings

`func (o *PrepmountRest) GetRestoreobjectmappings() []RestoreObjectMappingRest`

GetRestoreobjectmappings returns the Restoreobjectmappings field if non-nil, zero value otherwise.

### GetRestoreobjectmappingsOk

`func (o *PrepmountRest) GetRestoreobjectmappingsOk() (*[]RestoreObjectMappingRest, bool)`

GetRestoreobjectmappingsOk returns a tuple with the Restoreobjectmappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreobjectmappings

`func (o *PrepmountRest) SetRestoreobjectmappings(v []RestoreObjectMappingRest)`

SetRestoreobjectmappings sets Restoreobjectmappings field to given value.

### HasRestoreobjectmappings

`func (o *PrepmountRest) HasRestoreobjectmappings() bool`

HasRestoreobjectmappings returns a boolean if a field has been set.

### GetRestoreoptions

`func (o *PrepmountRest) GetRestoreoptions() []RestoreOptionRest`

GetRestoreoptions returns the Restoreoptions field if non-nil, zero value otherwise.

### GetRestoreoptionsOk

`func (o *PrepmountRest) GetRestoreoptionsOk() (*[]RestoreOptionRest, bool)`

GetRestoreoptionsOk returns a tuple with the Restoreoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreoptions

`func (o *PrepmountRest) SetRestoreoptions(v []RestoreOptionRest)`

SetRestoreoptions sets Restoreoptions field to given value.

### HasRestoreoptions

`func (o *PrepmountRest) HasRestoreoptions() bool`

HasRestoreoptions returns a boolean if a field has been set.

### GetProvisioningoptions

`func (o *PrepmountRest) GetProvisioningoptions() []ProvisioningOptionRest`

GetProvisioningoptions returns the Provisioningoptions field if non-nil, zero value otherwise.

### GetProvisioningoptionsOk

`func (o *PrepmountRest) GetProvisioningoptionsOk() (*[]ProvisioningOptionRest, bool)`

GetProvisioningoptionsOk returns a tuple with the Provisioningoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningoptions

`func (o *PrepmountRest) SetProvisioningoptions(v []ProvisioningOptionRest)`

SetProvisioningoptions sets Provisioningoptions field to given value.

### HasProvisioningoptions

`func (o *PrepmountRest) HasProvisioningoptions() bool`

HasProvisioningoptions returns a boolean if a field has been set.

### GetSelectedobjects

`func (o *PrepmountRest) GetSelectedobjects() []SelectedObjectRest`

GetSelectedobjects returns the Selectedobjects field if non-nil, zero value otherwise.

### GetSelectedobjectsOk

`func (o *PrepmountRest) GetSelectedobjectsOk() (*[]SelectedObjectRest, bool)`

GetSelectedobjectsOk returns a tuple with the Selectedobjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedobjects

`func (o *PrepmountRest) SetSelectedobjects(v []SelectedObjectRest)`

SetSelectedobjects sets Selectedobjects field to given value.

### HasSelectedobjects

`func (o *PrepmountRest) HasSelectedobjects() bool`

HasSelectedobjects returns a boolean if a field has been set.

### GetNfsoptions

`func (o *PrepmountRest) GetNfsoptions() NfsOptionsRest`

GetNfsoptions returns the Nfsoptions field if non-nil, zero value otherwise.

### GetNfsoptionsOk

`func (o *PrepmountRest) GetNfsoptionsOk() (*NfsOptionsRest, bool)`

GetNfsoptionsOk returns a tuple with the Nfsoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsoptions

`func (o *PrepmountRest) SetNfsoptions(v NfsOptionsRest)`

SetNfsoptions sets Nfsoptions field to given value.

### HasNfsoptions

`func (o *PrepmountRest) HasNfsoptions() bool`

HasNfsoptions returns a boolean if a field has been set.

### GetInstantmount

`func (o *PrepmountRest) GetInstantmount() bool`

GetInstantmount returns the Instantmount field if non-nil, zero value otherwise.

### GetInstantmountOk

`func (o *PrepmountRest) GetInstantmountOk() (*bool, bool)`

GetInstantmountOk returns a tuple with the Instantmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantmount

`func (o *PrepmountRest) SetInstantmount(v bool)`

SetInstantmount sets Instantmount field to given value.

### HasInstantmount

`func (o *PrepmountRest) HasInstantmount() bool`

HasInstantmount returns a boolean if a field has been set.

### GetContainer

`func (o *PrepmountRest) GetContainer() bool`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *PrepmountRest) GetContainerOk() (*bool, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *PrepmountRest) SetContainer(v bool)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *PrepmountRest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetAllowedips

`func (o *PrepmountRest) GetAllowedips() []string`

GetAllowedips returns the Allowedips field if non-nil, zero value otherwise.

### GetAllowedipsOk

`func (o *PrepmountRest) GetAllowedipsOk() (*[]string, bool)`

GetAllowedipsOk returns a tuple with the Allowedips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedips

`func (o *PrepmountRest) SetAllowedips(v []string)`

SetAllowedips sets Allowedips field to given value.

### HasAllowedips

`func (o *PrepmountRest) HasAllowedips() bool`

HasAllowedips returns a boolean if a field has been set.

### GetId

`func (o *PrepmountRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrepmountRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrepmountRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrepmountRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *PrepmountRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PrepmountRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PrepmountRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PrepmountRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *PrepmountRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *PrepmountRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *PrepmountRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *PrepmountRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *PrepmountRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *PrepmountRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *PrepmountRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *PrepmountRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


