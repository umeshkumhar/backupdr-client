# AgentRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecret** | Pointer to **string** | Secret key generated at the host and used by managing backup appliance to authenticate itself to host while bootstrapping PKI on the host. | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**Agenttype** | Pointer to **string** |  | [optional] 
**Haspassword** | Pointer to **bool** |  | [optional] 
**Alternatekey** | Pointer to **string** |  | [optional] 
**Hasalternatekey** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | Unique ID for this object | [optional] 
**Href** | Pointer to **string** | URL to access this object | [optional] 
**Syncdate** | Pointer to **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] 
**Stale** | Pointer to **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] 

## Methods

### NewAgentRest

`func NewAgentRest() *AgentRest`

NewAgentRest instantiates a new AgentRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRestWithDefaults

`func NewAgentRestWithDefaults() *AgentRest`

NewAgentRestWithDefaults instantiates a new AgentRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSecret

`func (o *AgentRest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *AgentRest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *AgentRest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *AgentRest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetPassword

`func (o *AgentRest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AgentRest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AgentRest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AgentRest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *AgentRest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AgentRest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AgentRest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *AgentRest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *AgentRest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AgentRest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AgentRest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AgentRest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAgentVersion

`func (o *AgentRest) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *AgentRest) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *AgentRest) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *AgentRest) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetAgenttype

`func (o *AgentRest) GetAgenttype() string`

GetAgenttype returns the Agenttype field if non-nil, zero value otherwise.

### GetAgenttypeOk

`func (o *AgentRest) GetAgenttypeOk() (*string, bool)`

GetAgenttypeOk returns a tuple with the Agenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgenttype

`func (o *AgentRest) SetAgenttype(v string)`

SetAgenttype sets Agenttype field to given value.

### HasAgenttype

`func (o *AgentRest) HasAgenttype() bool`

HasAgenttype returns a boolean if a field has been set.

### GetHaspassword

`func (o *AgentRest) GetHaspassword() bool`

GetHaspassword returns the Haspassword field if non-nil, zero value otherwise.

### GetHaspasswordOk

`func (o *AgentRest) GetHaspasswordOk() (*bool, bool)`

GetHaspasswordOk returns a tuple with the Haspassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaspassword

`func (o *AgentRest) SetHaspassword(v bool)`

SetHaspassword sets Haspassword field to given value.

### HasHaspassword

`func (o *AgentRest) HasHaspassword() bool`

HasHaspassword returns a boolean if a field has been set.

### GetAlternatekey

`func (o *AgentRest) GetAlternatekey() string`

GetAlternatekey returns the Alternatekey field if non-nil, zero value otherwise.

### GetAlternatekeyOk

`func (o *AgentRest) GetAlternatekeyOk() (*string, bool)`

GetAlternatekeyOk returns a tuple with the Alternatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatekey

`func (o *AgentRest) SetAlternatekey(v string)`

SetAlternatekey sets Alternatekey field to given value.

### HasAlternatekey

`func (o *AgentRest) HasAlternatekey() bool`

HasAlternatekey returns a boolean if a field has been set.

### GetHasalternatekey

`func (o *AgentRest) GetHasalternatekey() bool`

GetHasalternatekey returns the Hasalternatekey field if non-nil, zero value otherwise.

### GetHasalternatekeyOk

`func (o *AgentRest) GetHasalternatekeyOk() (*bool, bool)`

GetHasalternatekeyOk returns a tuple with the Hasalternatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasalternatekey

`func (o *AgentRest) SetHasalternatekey(v bool)`

SetHasalternatekey sets Hasalternatekey field to given value.

### HasHasalternatekey

`func (o *AgentRest) HasHasalternatekey() bool`

HasHasalternatekey returns a boolean if a field has been set.

### GetId

`func (o *AgentRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentRest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentRest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *AgentRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AgentRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AgentRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AgentRest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSyncdate

`func (o *AgentRest) GetSyncdate() int64`

GetSyncdate returns the Syncdate field if non-nil, zero value otherwise.

### GetSyncdateOk

`func (o *AgentRest) GetSyncdateOk() (*int64, bool)`

GetSyncdateOk returns a tuple with the Syncdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncdate

`func (o *AgentRest) SetSyncdate(v int64)`

SetSyncdate sets Syncdate field to given value.

### HasSyncdate

`func (o *AgentRest) HasSyncdate() bool`

HasSyncdate returns a boolean if a field has been set.

### GetStale

`func (o *AgentRest) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *AgentRest) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *AgentRest) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *AgentRest) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


