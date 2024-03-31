# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DelegateGetCallDownloadLog**](ApplianceDelegationApi.md#DelegateGetCallDownloadLog) | **Get** /appliancedelegation/{cluster_id}/config/download/log | Download logs from backup/recovery appliance
[**DownloadConnector**](ApplianceDelegationApi.md#DownloadConnector) | **Get** /appliancedelegation/{cluster_id}/connectorbinary/{connectorname} | Download connectors from backup/recovery appliance
[**DownloadOssNotice**](ApplianceDelegationApi.md#DownloadOssNotice) | **Get** /appliancedelegation/{cluster_id}/config/download/ossnotice | Download zip file containing licenses and notices for open-source components from backup/recovery appliance
[**UploadSoftwareUpgradeToAppliance**](ApplianceDelegationApi.md#UploadSoftwareUpgradeToAppliance) | **Post** /appliancedelegation/{cluster_id}/cluster/uploadupdate | Upload software upgrade packages

# **DelegateGetCallDownloadLog**
> interface{} DelegateGetCallDownloadLog(ctx, clusterId)
Download logs from backup/recovery appliance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int64**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadConnector**
> interface{} DownloadConnector(ctx, connectorname, clusterId)
Download connectors from backup/recovery appliance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorname** | **string**|  | 
  **clusterId** | **int64**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadOssNotice**
> interface{} DownloadOssNotice(ctx, clusterId)
Download zip file containing licenses and notices for open-source components from backup/recovery appliance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int64**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadSoftwareUpgradeToAppliance**
> string UploadSoftwareUpgradeToAppliance(ctx, clusterId, optional)
Upload software upgrade packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int64**|  | 
 **optional** | ***ApplianceDelegationApiUploadSoftwareUpgradeToApplianceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceDelegationApiUploadSoftwareUpgradeToApplianceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**optional.Interface of FormDataContentDisposition**](.md)|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

