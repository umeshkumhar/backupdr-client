# PolicyRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**[]SourceRest**](SourceRest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Slt** | Pointer to [**SltRest**](SltRest.md) |  | [optional] 
**Endtime** | Pointer to **string** |  | [optional] 
**Rpom** | Pointer to **string** |  | [optional] 
**Rpo** | Pointer to **string** |  | [optional] 
**Predecessor** | Pointer to [**PolicyRest**](PolicyRest.md) |  | [optional] 
**Exclusiontype** | Pointer to **string** |  | [optional] 
**Iscontinuous** | Pointer to **bool** |  | [optional] 
**Starttime** | Pointer to **string** |  | [optional] 
**Targetvault** | Pointer to **int32** |  | [optional] 
**Sourcevault** | Pointer to **int32** |  | [optional] 
**Selection** | Pointer to **string** |  | [optional] 
**Scheduletype** | Pointer to **string** |  | [optional] 
**Exclusion** | Pointer to **string** |  | [optional] 
**Reptype** | Pointer to **string** |  | [optional] 
**Retention** | Pointer to **string** |  | [optional] 
**Retentionm** | Pointer to **string** |  | [optional] 
**Encrypt** | Pointer to **string** |  | [optional] 
**Repeatinterval** | Pointer to **string** |  | [optional] 
**Exclusioninterval** | Pointer to **string** |  | [optional] 
**Remoteretention** | Pointer to **int32** |  | [optional] 
**Compliancesettings** | Pointer to [**ComplianceSettingsRest**](ComplianceSettingsRest.md) |  | [optional] 
**Options** | Pointer to [**[]AdvancedOptionRest**](AdvancedOptionRest.md) |  | [optional] 
**PolicyType** | Pointer to **string** |  | [optional] 
**Truncatelog** | Pointer to **string** |  | [optional] 
**Verifychoice** | Pointer to **string** |  | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Verification** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewPolicyRest

`func NewPolicyRest() *PolicyRest`

NewPolicyRest instantiates a new PolicyRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRestWithDefaults

`func NewPolicyRestWithDefaults() *PolicyRest`

NewPolicyRestWithDefaults instantiates a new PolicyRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PolicyRest) GetSource() []SourceRest`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PolicyRest) GetSourceOk() (*[]SourceRest, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PolicyRest) SetSource(v []SourceRest)`

SetSource sets Source field to given value.

### HasSource

`func (o *PolicyRest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *PolicyRest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PolicyRest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PolicyRest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PolicyRest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSlt

`func (o *PolicyRest) GetSlt() SltRest`

GetSlt returns the Slt field if non-nil, zero value otherwise.

### GetSltOk

`func (o *PolicyRest) GetSltOk() (*SltRest, bool)`

GetSltOk returns a tuple with the Slt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlt

`func (o *PolicyRest) SetSlt(v SltRest)`

SetSlt sets Slt field to given value.

### HasSlt

`func (o *PolicyRest) HasSlt() bool`

HasSlt returns a boolean if a field has been set.

### GetEndtime

`func (o *PolicyRest) GetEndtime() string`

GetEndtime returns the Endtime field if non-nil, zero value otherwise.

### GetEndtimeOk

`func (o *PolicyRest) GetEndtimeOk() (*string, bool)`

GetEndtimeOk returns a tuple with the Endtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndtime

`func (o *PolicyRest) SetEndtime(v string)`

SetEndtime sets Endtime field to given value.

### HasEndtime

`func (o *PolicyRest) HasEndtime() bool`

HasEndtime returns a boolean if a field has been set.

### GetRpom

`func (o *PolicyRest) GetRpom() string`

GetRpom returns the Rpom field if non-nil, zero value otherwise.

### GetRpomOk

`func (o *PolicyRest) GetRpomOk() (*string, bool)`

GetRpomOk returns a tuple with the Rpom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpom

`func (o *PolicyRest) SetRpom(v string)`

SetRpom sets Rpom field to given value.

### HasRpom

`func (o *PolicyRest) HasRpom() bool`

HasRpom returns a boolean if a field has been set.

### GetRpo

`func (o *PolicyRest) GetRpo() string`

GetRpo returns the Rpo field if non-nil, zero value otherwise.

### GetRpoOk

`func (o *PolicyRest) GetRpoOk() (*string, bool)`

GetRpoOk returns a tuple with the Rpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpo

`func (o *PolicyRest) SetRpo(v string)`

SetRpo sets Rpo field to given value.

### HasRpo

`func (o *PolicyRest) HasRpo() bool`

HasRpo returns a boolean if a field has been set.

### GetPredecessor

`func (o *PolicyRest) GetPredecessor() PolicyRest`

GetPredecessor returns the Predecessor field if non-nil, zero value otherwise.

### GetPredecessorOk

`func (o *PolicyRest) GetPredecessorOk() (*PolicyRest, bool)`

GetPredecessorOk returns a tuple with the Predecessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessor

`func (o *PolicyRest) SetPredecessor(v PolicyRest)`

SetPredecessor sets Predecessor field to given value.

### HasPredecessor

`func (o *PolicyRest) HasPredecessor() bool`

HasPredecessor returns a boolean if a field has been set.

### GetExclusiontype

`func (o *PolicyRest) GetExclusiontype() string`

GetExclusiontype returns the Exclusiontype field if non-nil, zero value otherwise.

### GetExclusiontypeOk

`func (o *PolicyRest) GetExclusiontypeOk() (*string, bool)`

GetExclusiontypeOk returns a tuple with the Exclusiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiontype

`func (o *PolicyRest) SetExclusiontype(v string)`

SetExclusiontype sets Exclusiontype field to given value.

### HasExclusiontype

`func (o *PolicyRest) HasExclusiontype() bool`

HasExclusiontype returns a boolean if a field has been set.

### GetIscontinuous

`func (o *PolicyRest) GetIscontinuous() bool`

GetIscontinuous returns the Iscontinuous field if non-nil, zero value otherwise.

### GetIscontinuousOk

`func (o *PolicyRest) GetIscontinuousOk() (*bool, bool)`

GetIscontinuousOk returns a tuple with the Iscontinuous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscontinuous

`func (o *PolicyRest) SetIscontinuous(v bool)`

SetIscontinuous sets Iscontinuous field to given value.

### HasIscontinuous

`func (o *PolicyRest) HasIscontinuous() bool`

HasIscontinuous returns a boolean if a field has been set.

### GetStarttime

`func (o *PolicyRest) GetStarttime() string`

GetStarttime returns the Starttime field if non-nil, zero value otherwise.

### GetStarttimeOk

`func (o *PolicyRest) GetStarttimeOk() (*string, bool)`

GetStarttimeOk returns a tuple with the Starttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttime

`func (o *PolicyRest) SetStarttime(v string)`

SetStarttime sets Starttime field to given value.

### HasStarttime

`func (o *PolicyRest) HasStarttime() bool`

HasStarttime returns a boolean if a field has been set.

### GetTargetvault

`func (o *PolicyRest) GetTargetvault() int32`

GetTargetvault returns the Targetvault field if non-nil, zero value otherwise.

### GetTargetvaultOk

`func (o *PolicyRest) GetTargetvaultOk() (*int32, bool)`

GetTargetvaultOk returns a tuple with the Targetvault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetvault

`func (o *PolicyRest) SetTargetvault(v int32)`

SetTargetvault sets Targetvault field to given value.

### HasTargetvault

`func (o *PolicyRest) HasTargetvault() bool`

HasTargetvault returns a boolean if a field has been set.

### GetSourcevault

`func (o *PolicyRest) GetSourcevault() int32`

GetSourcevault returns the Sourcevault field if non-nil, zero value otherwise.

### GetSourcevaultOk

`func (o *PolicyRest) GetSourcevaultOk() (*int32, bool)`

GetSourcevaultOk returns a tuple with the Sourcevault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcevault

`func (o *PolicyRest) SetSourcevault(v int32)`

SetSourcevault sets Sourcevault field to given value.

### HasSourcevault

`func (o *PolicyRest) HasSourcevault() bool`

HasSourcevault returns a boolean if a field has been set.

### GetSelection

`func (o *PolicyRest) GetSelection() string`

GetSelection returns the Selection field if non-nil, zero value otherwise.

### GetSelectionOk

`func (o *PolicyRest) GetSelectionOk() (*string, bool)`

GetSelectionOk returns a tuple with the Selection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelection

`func (o *PolicyRest) SetSelection(v string)`

SetSelection sets Selection field to given value.

### HasSelection

`func (o *PolicyRest) HasSelection() bool`

HasSelection returns a boolean if a field has been set.

### GetScheduletype

`func (o *PolicyRest) GetScheduletype() string`

GetScheduletype returns the Scheduletype field if non-nil, zero value otherwise.

### GetScheduletypeOk

`func (o *PolicyRest) GetScheduletypeOk() (*string, bool)`

GetScheduletypeOk returns a tuple with the Scheduletype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduletype

`func (o *PolicyRest) SetScheduletype(v string)`

SetScheduletype sets Scheduletype field to given value.

### HasScheduletype

`func (o *PolicyRest) HasScheduletype() bool`

HasScheduletype returns a boolean if a field has been set.

### GetExclusion

`func (o *PolicyRest) GetExclusion() string`

GetExclusion returns the Exclusion field if non-nil, zero value otherwise.

### GetExclusionOk

`func (o *PolicyRest) GetExclusionOk() (*string, bool)`

GetExclusionOk returns a tuple with the Exclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusion

`func (o *PolicyRest) SetExclusion(v string)`

SetExclusion sets Exclusion field to given value.

### HasExclusion

`func (o *PolicyRest) HasExclusion() bool`

HasExclusion returns a boolean if a field has been set.

### GetReptype

`func (o *PolicyRest) GetReptype() string`

GetReptype returns the Reptype field if non-nil, zero value otherwise.

### GetReptypeOk

`func (o *PolicyRest) GetReptypeOk() (*string, bool)`

GetReptypeOk returns a tuple with the Reptype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReptype

`func (o *PolicyRest) SetReptype(v string)`

SetReptype sets Reptype field to given value.

### HasReptype

`func (o *PolicyRest) HasReptype() bool`

HasReptype returns a boolean if a field has been set.

### GetRetention

`func (o *PolicyRest) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *PolicyRest) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *PolicyRest) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *PolicyRest) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetRetentionm

`func (o *PolicyRest) GetRetentionm() string`

GetRetentionm returns the Retentionm field if non-nil, zero value otherwise.

### GetRetentionmOk

`func (o *PolicyRest) GetRetentionmOk() (*string, bool)`

GetRetentionmOk returns a tuple with the Retentionm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionm

`func (o *PolicyRest) SetRetentionm(v string)`

SetRetentionm sets Retentionm field to given value.

### HasRetentionm

`func (o *PolicyRest) HasRetentionm() bool`

HasRetentionm returns a boolean if a field has been set.

### GetEncrypt

`func (o *PolicyRest) GetEncrypt() string`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *PolicyRest) GetEncryptOk() (*string, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *PolicyRest) SetEncrypt(v string)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *PolicyRest) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetRepeatinterval

`func (o *PolicyRest) GetRepeatinterval() string`

GetRepeatinterval returns the Repeatinterval field if non-nil, zero value otherwise.

### GetRepeatintervalOk

`func (o *PolicyRest) GetRepeatintervalOk() (*string, bool)`

GetRepeatintervalOk returns a tuple with the Repeatinterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatinterval

`func (o *PolicyRest) SetRepeatinterval(v string)`

SetRepeatinterval sets Repeatinterval field to given value.

### HasRepeatinterval

`func (o *PolicyRest) HasRepeatinterval() bool`

HasRepeatinterval returns a boolean if a field has been set.

### GetExclusioninterval

`func (o *PolicyRest) GetExclusioninterval() string`

GetExclusioninterval returns the Exclusioninterval field if non-nil, zero value otherwise.

### GetExclusionintervalOk

`func (o *PolicyRest) GetExclusionintervalOk() (*string, bool)`

GetExclusionintervalOk returns a tuple with the Exclusioninterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusioninterval

`func (o *PolicyRest) SetExclusioninterval(v string)`

SetExclusioninterval sets Exclusioninterval field to given value.

### HasExclusioninterval

`func (o *PolicyRest) HasExclusioninterval() bool`

HasExclusioninterval returns a boolean if a field has been set.

### GetRemoteretention

`func (o *PolicyRest) GetRemoteretention() int32`

GetRemoteretention returns the Remoteretention field if non-nil, zero value otherwise.

### GetRemoteretentionOk

`func (o *PolicyRest) GetRemoteretentionOk() (*int32, bool)`

GetRemoteretentionOk returns a tuple with the Remoteretention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteretention

`func (o *PolicyRest) SetRemoteretention(v int32)`

SetRemoteretention sets Remoteretention field to given value.

### HasRemoteretention

`func (o *PolicyRest) HasRemoteretention() bool`

HasRemoteretention returns a boolean if a field has been set.

### GetCompliancesettings

`func (o *PolicyRest) GetCompliancesettings() ComplianceSettingsRest`

GetCompliancesettings returns the Compliancesettings field if non-nil, zero value otherwise.

### GetCompliancesettingsOk

`func (o *PolicyRest) GetCompliancesettingsOk() (*ComplianceSettingsRest, bool)`

GetCompliancesettingsOk returns a tuple with the Compliancesettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliancesettings

`func (o *PolicyRest) SetCompliancesettings(v ComplianceSettingsRest)`

SetCompliancesettings sets Compliancesettings field to given value.

### HasCompliancesettings

`func (o *PolicyRest) HasCompliancesettings() bool`

HasCompliancesettings returns a boolean if a field has been set.

### GetOptions

`func (o *PolicyRest) GetOptions() []AdvancedOptionRest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PolicyRest) GetOptionsOk() (*[]AdvancedOptionRest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PolicyRest) SetOptions(v []AdvancedOptionRest)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PolicyRest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPolicyType

`func (o *PolicyRest) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *PolicyRest) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *PolicyRest) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *PolicyRest) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetTruncatelog

`func (o *PolicyRest) GetTruncatelog() string`

GetTruncatelog returns the Truncatelog field if non-nil, zero value otherwise.

### GetTruncatelogOk

`func (o *PolicyRest) GetTruncatelogOk() (*string, bool)`

GetTruncatelogOk returns a tuple with the Truncatelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatelog

`func (o *PolicyRest) SetTruncatelog(v string)`

SetTruncatelog sets Truncatelog field to given value.

### HasTruncatelog

`func (o *PolicyRest) HasTruncatelog() bool`

HasTruncatelog returns a boolean if a field has been set.

### GetVerifychoice

`func (o *PolicyRest) GetVerifychoice() string`

GetVerifychoice returns the Verifychoice field if non-nil, zero value otherwise.

### GetVerifychoiceOk

`func (o *PolicyRest) GetVerifychoiceOk() (*string, bool)`

GetVerifychoiceOk returns a tuple with the Verifychoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifychoice

`func (o *PolicyRest) SetVerifychoice(v string)`

SetVerifychoice sets Verifychoice field to given value.

### HasVerifychoice

`func (o *PolicyRest) HasVerifychoice() bool`

HasVerifychoice returns a boolean if a field has been set.

### GetOp

`func (o *PolicyRest) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PolicyRest) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PolicyRest) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *PolicyRest) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetVerification

`func (o *PolicyRest) GetVerification() bool`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *PolicyRest) GetVerificationOk() (*bool, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *PolicyRest) SetVerification(v bool)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *PolicyRest) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetId

`func (o *PolicyRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *PolicyRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PolicyRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PolicyRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PolicyRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *PolicyRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *PolicyRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *PolicyRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *PolicyRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *PolicyRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *PolicyRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *PolicyRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *PolicyRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


