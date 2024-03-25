# FormFieldRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Header** | Pointer to [**[]ChoiceValueRest**](ChoiceValueRest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Maximum** | Pointer to **int32** |  | [optional] 
**Minimum** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Choices** | Pointer to [**[]ChoiceValueRest**](ChoiceValueRest.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Warning** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]VolumeSelectionRowRest**](VolumeSelectionRowRest.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**CurrentValue** | Pointer to **string** |  | [optional] 
**Validation** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**HelpId** | Pointer to **string** |  | [optional] 
**Optiongroup** | Pointer to **bool** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**Modified** | Pointer to **bool** |  | [optional] 
**Dynamic** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to [**[]ChoiceValueRest**](ChoiceValueRest.md) |  | [optional] 
**Networktags** | Pointer to [**[]ChoiceValueRest**](ChoiceValueRest.md) |  | [optional] 
**GetGetchoices** | Pointer to **string** |  | [optional] 
**GetGetDefault** | Pointer to **string** |  | [optional] 
**GetDependent** | Pointer to **[]string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**GroupBy** | Pointer to **bool** |  | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**Invalid** | Pointer to **string** |  | [optional] 
**InvalidStr** | Pointer to **string** |  | [optional] 
**GetDefault** | Pointer to **string** |  | [optional] 
**Checked** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewFormFieldRest

`func NewFormFieldRest() *FormFieldRest`

NewFormFieldRest instantiates a new FormFieldRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldRestWithDefaults

`func NewFormFieldRestWithDefaults() *FormFieldRest`

NewFormFieldRestWithDefaults instantiates a new FormFieldRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *FormFieldRest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldRest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldRest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormFieldRest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDisplayName

`func (o *FormFieldRest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FormFieldRest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FormFieldRest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FormFieldRest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHeader

`func (o *FormFieldRest) GetHeader() []ChoiceValueRest`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *FormFieldRest) GetHeaderOk() (*[]ChoiceValueRest, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *FormFieldRest) SetHeader(v []ChoiceValueRest)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *FormFieldRest) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetDescription

`func (o *FormFieldRest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormFieldRest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormFieldRest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FormFieldRest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *FormFieldRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormFieldRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormFieldRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormFieldRest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FormFieldRest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldRest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldRest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormFieldRest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *FormFieldRest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FormFieldRest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FormFieldRest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FormFieldRest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMaximum

`func (o *FormFieldRest) GetMaximum() int32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *FormFieldRest) GetMaximumOk() (*int32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *FormFieldRest) SetMaximum(v int32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *FormFieldRest) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMinimum

`func (o *FormFieldRest) GetMinimum() int32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *FormFieldRest) GetMinimumOk() (*int32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *FormFieldRest) SetMinimum(v int32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *FormFieldRest) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetErrorMessage

`func (o *FormFieldRest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FormFieldRest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FormFieldRest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FormFieldRest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetChoices

`func (o *FormFieldRest) GetChoices() []ChoiceValueRest`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *FormFieldRest) GetChoicesOk() (*[]ChoiceValueRest, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *FormFieldRest) SetChoices(v []ChoiceValueRest)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *FormFieldRest) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetDisabled

`func (o *FormFieldRest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *FormFieldRest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *FormFieldRest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *FormFieldRest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetWarning

`func (o *FormFieldRest) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *FormFieldRest) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *FormFieldRest) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *FormFieldRest) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetRows

`func (o *FormFieldRest) GetRows() []VolumeSelectionRowRest`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *FormFieldRest) GetRowsOk() (*[]VolumeSelectionRowRest, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *FormFieldRest) SetRows(v []VolumeSelectionRowRest)`

SetRows sets Rows field to given value.

### HasRows

`func (o *FormFieldRest) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetValues

`func (o *FormFieldRest) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FormFieldRest) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FormFieldRest) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *FormFieldRest) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetCurrentValue

`func (o *FormFieldRest) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *FormFieldRest) GetCurrentValueOk() (*string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *FormFieldRest) SetCurrentValue(v string)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *FormFieldRest) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetValidation

`func (o *FormFieldRest) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *FormFieldRest) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *FormFieldRest) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *FormFieldRest) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### GetTitle

`func (o *FormFieldRest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormFieldRest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormFieldRest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormFieldRest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetHelpId

`func (o *FormFieldRest) GetHelpId() string`

GetHelpId returns the HelpId field if non-nil, zero value otherwise.

### GetHelpIdOk

`func (o *FormFieldRest) GetHelpIdOk() (*string, bool)`

GetHelpIdOk returns a tuple with the HelpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpId

`func (o *FormFieldRest) SetHelpId(v string)`

SetHelpId sets HelpId field to given value.

### HasHelpId

`func (o *FormFieldRest) HasHelpId() bool`

HasHelpId returns a boolean if a field has been set.

### GetOptiongroup

`func (o *FormFieldRest) GetOptiongroup() bool`

GetOptiongroup returns the Optiongroup field if non-nil, zero value otherwise.

### GetOptiongroupOk

`func (o *FormFieldRest) GetOptiongroupOk() (*bool, bool)`

GetOptiongroupOk returns a tuple with the Optiongroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptiongroup

`func (o *FormFieldRest) SetOptiongroup(v bool)`

SetOptiongroup sets Optiongroup field to given value.

### HasOptiongroup

`func (o *FormFieldRest) HasOptiongroup() bool`

HasOptiongroup returns a boolean if a field has been set.

### GetReadonly

`func (o *FormFieldRest) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *FormFieldRest) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *FormFieldRest) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *FormFieldRest) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetModified

`func (o *FormFieldRest) GetModified() bool`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *FormFieldRest) GetModifiedOk() (*bool, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *FormFieldRest) SetModified(v bool)`

SetModified sets Modified field to given value.

### HasModified

`func (o *FormFieldRest) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDynamic

`func (o *FormFieldRest) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *FormFieldRest) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *FormFieldRest) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *FormFieldRest) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.

### GetTags

`func (o *FormFieldRest) GetTags() []ChoiceValueRest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FormFieldRest) GetTagsOk() (*[]ChoiceValueRest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FormFieldRest) SetTags(v []ChoiceValueRest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FormFieldRest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworktags

`func (o *FormFieldRest) GetNetworktags() []ChoiceValueRest`

GetNetworktags returns the Networktags field if non-nil, zero value otherwise.

### GetNetworktagsOk

`func (o *FormFieldRest) GetNetworktagsOk() (*[]ChoiceValueRest, bool)`

GetNetworktagsOk returns a tuple with the Networktags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworktags

`func (o *FormFieldRest) SetNetworktags(v []ChoiceValueRest)`

SetNetworktags sets Networktags field to given value.

### HasNetworktags

`func (o *FormFieldRest) HasNetworktags() bool`

HasNetworktags returns a boolean if a field has been set.

### GetGetGetchoices

`func (o *FormFieldRest) GetGetGetchoices() string`

GetGetGetchoices returns the GetGetchoices field if non-nil, zero value otherwise.

### GetGetGetchoicesOk

`func (o *FormFieldRest) GetGetGetchoicesOk() (*string, bool)`

GetGetGetchoicesOk returns a tuple with the GetGetchoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetGetchoices

`func (o *FormFieldRest) SetGetGetchoices(v string)`

SetGetGetchoices sets GetGetchoices field to given value.

### HasGetGetchoices

`func (o *FormFieldRest) HasGetGetchoices() bool`

HasGetGetchoices returns a boolean if a field has been set.

### GetGetGetDefault

`func (o *FormFieldRest) GetGetGetDefault() string`

GetGetGetDefault returns the GetGetDefault field if non-nil, zero value otherwise.

### GetGetGetDefaultOk

`func (o *FormFieldRest) GetGetGetDefaultOk() (*string, bool)`

GetGetGetDefaultOk returns a tuple with the GetGetDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetGetDefault

`func (o *FormFieldRest) SetGetGetDefault(v string)`

SetGetGetDefault sets GetGetDefault field to given value.

### HasGetGetDefault

`func (o *FormFieldRest) HasGetGetDefault() bool`

HasGetGetDefault returns a boolean if a field has been set.

### GetGetDependent

`func (o *FormFieldRest) GetGetDependent() []string`

GetGetDependent returns the GetDependent field if non-nil, zero value otherwise.

### GetGetDependentOk

`func (o *FormFieldRest) GetGetDependentOk() (*[]string, bool)`

GetGetDependentOk returns a tuple with the GetDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetDependent

`func (o *FormFieldRest) SetGetDependent(v []string)`

SetGetDependent sets GetDependent field to given value.

### HasGetDependent

`func (o *FormFieldRest) HasGetDependent() bool`

HasGetDependent returns a boolean if a field has been set.

### GetHidden

`func (o *FormFieldRest) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FormFieldRest) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FormFieldRest) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FormFieldRest) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetGroupBy

`func (o *FormFieldRest) GetGroupBy() bool`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *FormFieldRest) GetGroupByOk() (*bool, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *FormFieldRest) SetGroupBy(v bool)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *FormFieldRest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupType

`func (o *FormFieldRest) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *FormFieldRest) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *FormFieldRest) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *FormFieldRest) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetInvalid

`func (o *FormFieldRest) GetInvalid() string`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *FormFieldRest) GetInvalidOk() (*string, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *FormFieldRest) SetInvalid(v string)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *FormFieldRest) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.

### GetInvalidStr

`func (o *FormFieldRest) GetInvalidStr() string`

GetInvalidStr returns the InvalidStr field if non-nil, zero value otherwise.

### GetInvalidStrOk

`func (o *FormFieldRest) GetInvalidStrOk() (*string, bool)`

GetInvalidStrOk returns a tuple with the InvalidStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidStr

`func (o *FormFieldRest) SetInvalidStr(v string)`

SetInvalidStr sets InvalidStr field to given value.

### HasInvalidStr

`func (o *FormFieldRest) HasInvalidStr() bool`

HasInvalidStr returns a boolean if a field has been set.

### GetGetDefault

`func (o *FormFieldRest) GetGetDefault() string`

GetGetDefault returns the GetDefault field if non-nil, zero value otherwise.

### GetGetDefaultOk

`func (o *FormFieldRest) GetGetDefaultOk() (*string, bool)`

GetGetDefaultOk returns a tuple with the GetDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetDefault

`func (o *FormFieldRest) SetGetDefault(v string)`

SetGetDefault sets GetDefault field to given value.

### HasGetDefault

`func (o *FormFieldRest) HasGetDefault() bool`

HasGetDefault returns a boolean if a field has been set.

### GetChecked

`func (o *FormFieldRest) GetChecked() bool`

GetChecked returns the Checked field if non-nil, zero value otherwise.

### GetCheckedOk

`func (o *FormFieldRest) GetCheckedOk() (*bool, bool)`

GetCheckedOk returns a tuple with the Checked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecked

`func (o *FormFieldRest) SetChecked(v bool)`

SetChecked sets Checked field to given value.

### HasChecked

`func (o *FormFieldRest) HasChecked() bool`

HasChecked returns a boolean if a field has been set.

### GetId

`func (o *FormFieldRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormFieldRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormFieldRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormFieldRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *FormFieldRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FormFieldRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FormFieldRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FormFieldRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *FormFieldRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *FormFieldRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *FormFieldRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *FormFieldRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *FormFieldRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *FormFieldRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *FormFieldRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *FormFieldRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


