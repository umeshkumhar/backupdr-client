# AgentRest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecret** | **string** | Secret key generated at the host and used by managing backup appliance to authenticate itself to host while bootstrapping PKI on the host. | [optional] [default to null]
**Password** | **string** |  | [optional] [default to null]
**Username** | **string** |  | [optional] [default to null]
**AgentVersion** | **string** |  | [optional] [default to null]
**Agenttype** | **string** |  | [optional] [default to null]
**Haspassword** | **bool** |  | [optional] [default to null]
**Alternatekey** | **string** |  | [optional] [default to null]
**Hasalternatekey** | **bool** |  | [optional] [default to null]
**Port** | **int32** |  | [optional] [default to null]
**Href** | **string** | URL to access this object | [optional] [default to null]
**Syncdate** | **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] [default to null]
**Stale** | **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] [default to null]
**Id** | **string** | Unique ID for this object | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

