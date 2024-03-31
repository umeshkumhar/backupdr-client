# PrepmountRest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | [***HostRest**](HostRest.md) |  | [optional] [default to null]
**Rdmmode** | **string** |  | [optional] [default to null]
**Physicalrdm** | **string** |  | [optional] [default to null]
**Appaware** | **bool** |  | [optional] [default to null]
**Recoverytime** | **int64** |  | [optional] [default to null]
**Scripts** | [**[]ScriptRest**](ScriptRest.md) |  | [optional] [default to null]
**Restoreobjectmappings** | [**[]RestoreObjectMappingRest**](RestoreObjectMappingRest.md) |  | [optional] [default to null]
**Restoreoptions** | [**[]RestoreOptionRest**](RestoreOptionRest.md) |  | [optional] [default to null]
**Provisioningoptions** | [**[]ProvisioningOptionRest**](ProvisioningOptionRest.md) |  | [optional] [default to null]
**Selectedobjects** | [**[]SelectedObjectRest**](SelectedObjectRest.md) |  | [optional] [default to null]
**Nfsoptions** | [***NfsOptionsRest**](NfsOptionsRest.md) |  | [optional] [default to null]
**Instantmount** | **bool** |  | [optional] [default to null]
**Container** | **bool** |  | [optional] [default to null]
**Allowedips** | **[]string** |  | [optional] [default to null]
**Id** | **string** | Unique ID for this object | [optional] [default to null]
**Href** | **string** | URL to access this object | [optional] [default to null]
**Syncdate** | **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] [default to null]
**Stale** | **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

