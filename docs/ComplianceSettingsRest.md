# ComplianceSettingsRest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarnThresholdType** | **string** |  | [optional] [default to null]
**WarnThresholdCustom** | **int32** |  | [optional] [default to null]
**ErrorThresholdType** | **string** |  | [optional] [default to null]
**ErrorThresholdCustom** | **int32** |  | [optional] [default to null]
**Policy** | [***PolicyRest**](PolicyRest.md) |  | [optional] [default to null]
**Href** | **string** | URL to access this object | [optional] [default to null]
**Syncdate** | **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] [default to null]
**Stale** | **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] [default to null]
**Id** | **string** | Unique ID for this object | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

