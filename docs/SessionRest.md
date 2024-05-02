# SessionRest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | Session id of the current login | [optional] [default to null]
**Region** | **string** | region | [optional] [default to null]
**User** | [***UserRest**](UserRest.md) |  | [optional] [default to null]
**Timezone** | **string** | User&#x27;s timezone, if set. | [optional] [default to null]
**Username** | **string** |  | [optional] [default to null]
**Userpref** | **string** | This object contains user preferences that client programs can utilize. The content is opaque to API service. | [optional] [default to null]
**Rights** | [**[]RightRest**](RightRest.md) | Effective access rights that this login session contains. | [optional] [default to null]
**Authconfig** | [***AuthConfigRest**](AuthConfigRest.md) |  | [optional] [default to null]
**Href** | **string** | URL to access this object | [optional] [default to null]
**Syncdate** | **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] [default to null]
**Stale** | **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] [default to null]
**Id** | **string** | Unique ID for this object | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

