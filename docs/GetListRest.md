# GetListRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sortablefields** | Pointer to **[]string** |  | [optional] 
**Sortable** | Pointer to **bool** |  | [optional] 
**Filterable** | Pointer to **bool** |  | [optional] 
**Filterablefields** | Pointer to [**[]FilterFieldClassMapping**](FilterFieldClassMapping.md) |  | [optional] 
**Pageable** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetListRest

`func NewGetListRest() *GetListRest`

NewGetListRest instantiates a new GetListRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListRestWithDefaults

`func NewGetListRestWithDefaults() *GetListRest`

NewGetListRestWithDefaults instantiates a new GetListRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSortablefields

`func (o *GetListRest) GetSortablefields() []string`

GetSortablefields returns the Sortablefields field if non-nil, zero value otherwise.

### GetSortablefieldsOk

`func (o *GetListRest) GetSortablefieldsOk() (*[]string, bool)`

GetSortablefieldsOk returns a tuple with the Sortablefields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortablefields

`func (o *GetListRest) SetSortablefields(v []string)`

SetSortablefields sets Sortablefields field to given value.

### HasSortablefields

`func (o *GetListRest) HasSortablefields() bool`

HasSortablefields returns a boolean if a field has been set.

### GetSortable

`func (o *GetListRest) GetSortable() bool`

GetSortable returns the Sortable field if non-nil, zero value otherwise.

### GetSortableOk

`func (o *GetListRest) GetSortableOk() (*bool, bool)`

GetSortableOk returns a tuple with the Sortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortable

`func (o *GetListRest) SetSortable(v bool)`

SetSortable sets Sortable field to given value.

### HasSortable

`func (o *GetListRest) HasSortable() bool`

HasSortable returns a boolean if a field has been set.

### GetFilterable

`func (o *GetListRest) GetFilterable() bool`

GetFilterable returns the Filterable field if non-nil, zero value otherwise.

### GetFilterableOk

`func (o *GetListRest) GetFilterableOk() (*bool, bool)`

GetFilterableOk returns a tuple with the Filterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterable

`func (o *GetListRest) SetFilterable(v bool)`

SetFilterable sets Filterable field to given value.

### HasFilterable

`func (o *GetListRest) HasFilterable() bool`

HasFilterable returns a boolean if a field has been set.

### GetFilterablefields

`func (o *GetListRest) GetFilterablefields() []FilterFieldClassMapping`

GetFilterablefields returns the Filterablefields field if non-nil, zero value otherwise.

### GetFilterablefieldsOk

`func (o *GetListRest) GetFilterablefieldsOk() (*[]FilterFieldClassMapping, bool)`

GetFilterablefieldsOk returns a tuple with the Filterablefields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterablefields

`func (o *GetListRest) SetFilterablefields(v []FilterFieldClassMapping)`

SetFilterablefields sets Filterablefields field to given value.

### HasFilterablefields

`func (o *GetListRest) HasFilterablefields() bool`

HasFilterablefields returns a boolean if a field has been set.

### GetPageable

`func (o *GetListRest) GetPageable() bool`

GetPageable returns the Pageable field if non-nil, zero value otherwise.

### GetPageableOk

`func (o *GetListRest) GetPageableOk() (*bool, bool)`

GetPageableOk returns a tuple with the Pageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageable

`func (o *GetListRest) SetPageable(v bool)`

SetPageable sets Pageable field to given value.

### HasPageable

`func (o *GetListRest) HasPageable() bool`

HasPageable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


