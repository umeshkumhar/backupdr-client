# WorkflowRest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]WorkflowItemRest**](WorkflowItemRest.md) |  | [optional] [default to null]
**Status** | [***WorkflowStatusRest**](WorkflowStatusRest.md) |  | [optional] [default to null]
**Application** | [***ApplicationRest**](ApplicationRest.md) |  | [optional] [default to null]
**Cluster** | [***ClusterRest**](ClusterRest.md) |  | [optional] [default to null]
**Disabled** | **bool** |  | [optional] [default to null]
**Schedule** | [***WorkScheduleRest**](WorkScheduleRest.md) |  | [optional] [default to null]
**Props** | [**[]KeyValueRest**](KeyValueRest.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Href** | **string** | URL to access this object | [optional] [default to null]
**Syncdate** | **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] [default to null]
**Stale** | **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] [default to null]
**Id** | **string** | Unique ID for this object | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

