# MountRest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | [optional] [default to null]
**Hostname** | **string** | Mutually exclusive with host, when mounting as a brand new host(VM) | [optional] [default to null]
**Container** | **bool** |  | [optional] [default to null]
**Diskpool** | [***DiskPoolRest**](DiskPoolRest.md) |  | [optional] [default to null]
**Cloudvmoptions** | [***CloudVmMountRest**](CloudVmMountRest.md) |  | [optional] [default to null]
**Rdmmode** | **string** |  | [optional] [default to null]
**Rehydrationmode** | **string** |  | [optional] [default to null]
**Appaware** | **bool** |  | [optional] [default to null]
**Systemstateoptions** | [**[]SystemStateOptionRest**](SystemStateOptionRest.md) |  | [optional] [default to null]
**Migratevm** | **bool** |  | [optional] [default to null]
**Datastore** | **string** | Mutually exclusive with diskpool | [optional] [default to null]
**Recoverytime** | **int64** | Can be used when log backup is available. Microseconds since Epoch in UTC | [optional] [default to null]
**Poweronvm** | **bool** |  | [optional] [default to null]
**Scripts** | [**[]ScriptRest**](ScriptRest.md) |  | [optional] [default to null]
**Restoreobjectmappings** | [**[]RestoreObjectMappingRest**](RestoreObjectMappingRest.md) | A customized mapping from the volumes to their mount points | [optional] [default to null]
**Restoreoptions** | [**[]RestoreOptionRest**](RestoreOptionRest.md) |  | [optional] [default to null]
**Provisioningoptions** | [**[]ProvisioningOptionRest**](ProvisioningOptionRest.md) | List of provisioning options when appaware is true | [optional] [default to null]
**Selectedobjects** | [**[]SelectedObjectRest**](SelectedObjectRest.md) | A subset of restorable objects to be mounted | [optional] [default to null]
**Nfsoptions** | [***NfsOptionsRest**](NfsOptionsRest.md) |  | [optional] [default to null]
**Instantmount** | **bool** |  | [optional] [default to null]
**Mgmtserver** | [***HostRest**](HostRest.md) |  | [optional] [default to null]
**Hypervisor** | [***HostRest**](HostRest.md) |  | [optional] [default to null]
**Allowedips** | **[]string** |  | [optional] [default to null]
**Path** | **string** |  | [optional] [default to null]
**Host** | [***HostRest**](HostRest.md) |  | [optional] [default to null]
**Href** | **string** | URL to access this object | [optional] [default to null]
**Syncdate** | **int64** | When this object was last synced from appliances (UNIX Epoch time in microseconds). It does not apply to local resources. | [optional] [default to null]
**Stale** | **bool** | Optional flag to indicate if the information is out-of-date due to communication problems with appliances. It does not apply to local resources. | [optional] [default to null]
**Id** | **string** | Unique ID for this object | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

