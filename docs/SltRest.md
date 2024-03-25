# SltRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Immutable** | Pointer to **bool** |  | [optional] 
**OptionHref** | Pointer to **string** |  | [optional] 
**PolicyHref** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**[]SourceRest**](SourceRest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Sourcename** | Pointer to **string** |  | [optional] 
**Override** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to [**[]PolicyRest**](PolicyRest.md) |  | [optional] 
**Options** | Pointer to [**[]AdvancedOptionRest**](AdvancedOptionRest.md) |  | [optional] 
**Orglist** | Pointer to [**[]OrganizationRest**](OrganizationRest.md) |  | [optional] 
**Managedbyagm** | Pointer to **bool** |  | [optional] 
**Usedbycloudapp** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewSltRest

`func NewSltRest() *SltRest`

NewSltRest instantiates a new SltRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSltRestWithDefaults

`func NewSltRestWithDefaults() *SltRest`

NewSltRestWithDefaults instantiates a new SltRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImmutable

`func (o *SltRest) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *SltRest) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *SltRest) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *SltRest) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### GetOptionHref

`func (o *SltRest) GetOptionHref() string`

GetOptionHref returns the OptionHref field if non-nil, zero value otherwise.

### GetOptionHrefOk

`func (o *SltRest) GetOptionHrefOk() (*string, bool)`

GetOptionHrefOk returns a tuple with the OptionHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionHref

`func (o *SltRest) SetOptionHref(v string)`

SetOptionHref sets OptionHref field to given value.

### HasOptionHref

`func (o *SltRest) HasOptionHref() bool`

HasOptionHref returns a boolean if a field has been set.

### GetPolicyHref

`func (o *SltRest) GetPolicyHref() string`

GetPolicyHref returns the PolicyHref field if non-nil, zero value otherwise.

### GetPolicyHrefOk

`func (o *SltRest) GetPolicyHrefOk() (*string, bool)`

GetPolicyHrefOk returns a tuple with the PolicyHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyHref

`func (o *SltRest) SetPolicyHref(v string)`

SetPolicyHref sets PolicyHref field to given value.

### HasPolicyHref

`func (o *SltRest) HasPolicyHref() bool`

HasPolicyHref returns a boolean if a field has been set.

### GetSource

`func (o *SltRest) GetSource() []SourceRest`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SltRest) GetSourceOk() (*[]SourceRest, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SltRest) SetSource(v []SourceRest)`

SetSource sets Source field to given value.

### HasSource

`func (o *SltRest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDescription

`func (o *SltRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SltRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SltRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SltRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SltRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SltRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SltRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SltRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourcename

`func (o *SltRest) GetSourcename() string`

GetSourcename returns the Sourcename field if non-nil, zero value otherwise.

### GetSourcenameOk

`func (o *SltRest) GetSourcenameOk() (*string, bool)`

GetSourcenameOk returns a tuple with the Sourcename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcename

`func (o *SltRest) SetSourcename(v string)`

SetSourcename sets Sourcename field to given value.

### HasSourcename

`func (o *SltRest) HasSourcename() bool`

HasSourcename returns a boolean if a field has been set.

### GetOverride

`func (o *SltRest) GetOverride() string`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *SltRest) GetOverrideOk() (*string, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *SltRest) SetOverride(v string)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *SltRest) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### GetPolicies

`func (o *SltRest) GetPolicies() []PolicyRest`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *SltRest) GetPoliciesOk() (*[]PolicyRest, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *SltRest) SetPolicies(v []PolicyRest)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *SltRest) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetOptions

`func (o *SltRest) GetOptions() []AdvancedOptionRest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SltRest) GetOptionsOk() (*[]AdvancedOptionRest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SltRest) SetOptions(v []AdvancedOptionRest)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SltRest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOrglist

`func (o *SltRest) GetOrglist() []OrganizationRest`

GetOrglist returns the Orglist field if non-nil, zero value otherwise.

### GetOrglistOk

`func (o *SltRest) GetOrglistOk() (*[]OrganizationRest, bool)`

GetOrglistOk returns a tuple with the Orglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrglist

`func (o *SltRest) SetOrglist(v []OrganizationRest)`

SetOrglist sets Orglist field to given value.

### HasOrglist

`func (o *SltRest) HasOrglist() bool`

HasOrglist returns a boolean if a field has been set.

### GetManagedbyagm

`func (o *SltRest) GetManagedbyagm() bool`

GetManagedbyagm returns the Managedbyagm field if non-nil, zero value otherwise.

### GetManagedbyagmOk

`func (o *SltRest) GetManagedbyagmOk() (*bool, bool)`

GetManagedbyagmOk returns a tuple with the Managedbyagm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedbyagm

`func (o *SltRest) SetManagedbyagm(v bool)`

SetManagedbyagm sets Managedbyagm field to given value.

### HasManagedbyagm

`func (o *SltRest) HasManagedbyagm() bool`

HasManagedbyagm returns a boolean if a field has been set.

### GetUsedbycloudapp

`func (o *SltRest) GetUsedbycloudapp() bool`

GetUsedbycloudapp returns the Usedbycloudapp field if non-nil, zero value otherwise.

### GetUsedbycloudappOk

`func (o *SltRest) GetUsedbycloudappOk() (*bool, bool)`

GetUsedbycloudappOk returns a tuple with the Usedbycloudapp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedbycloudapp

`func (o *SltRest) SetUsedbycloudapp(v bool)`

SetUsedbycloudapp sets Usedbycloudapp field to given value.

### HasUsedbycloudapp

`func (o *SltRest) HasUsedbycloudapp() bool`

HasUsedbycloudapp returns a boolean if a field has been set.

### GetId

`func (o *SltRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SltRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SltRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SltRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *SltRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SltRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SltRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SltRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *SltRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *SltRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *SltRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *SltRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *SltRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *SltRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *SltRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *SltRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


