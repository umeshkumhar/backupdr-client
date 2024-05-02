# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCloudPool**](DefaultApi.md#AddCloudPool) | **Post** /cloudcredential/{cloudcredential_id}/cloudpool | 
[**AddVm**](DefaultApi.md#AddVm) | **Post** /cloudcredential/{cloudcredential_id}/discovervm/addvm | 
[**CountCloudCredentials**](DefaultApi.md#CountCloudCredentials) | **Head** /cloudcredential | 
[**CountDynamicProtections**](DefaultApi.md#CountDynamicProtections) | **Head** /dynamicprotection | Get a count of total dynamic protection entries. It requires backupdr.managementServers.getDynamicProtection IAM permission
[**CreateCredential**](DefaultApi.md#CreateCredential) | **Post** /cloudcredential | 
[**CreateDynamicProtection**](DefaultApi.md#CreateDynamicProtection) | **Post** /dynamicprotection | Create a new dynamic protection entry. It requires backupdr.managementServers.createDynamicProtection IAM permission
[**DeleteCloudPool**](DefaultApi.md#DeleteCloudPool) | **Delete** /cloudcredential/{cloudcredential_id}/cloudpool | 
[**DeleteCredential**](DefaultApi.md#DeleteCredential) | **Delete** /cloudcredential/{cloudcredential_id} | 
[**DeleteDynamicProtection**](DefaultApi.md#DeleteDynamicProtection) | **Delete** /dynamicprotection/{dynamicprotection_id} | Remove a dynamic protection entry. It requires backupdr.managementServers.deleteDynamicProtection IAM permission
[**DiscoverCloudVm**](DefaultApi.md#DiscoverCloudVm) | **Post** /cloudcredential/{cloudcredential_id}/discovervm/vm | 
[**GetAllZones**](DefaultApi.md#GetAllZones) | **Get** /cloudcredential/{cloudcredential_id}/region/{region}/zone | 
[**GetCloudCredentialMetaInfo**](DefaultApi.md#GetCloudCredentialMetaInfo) | **Get** /cloudcredential/cloudtype/{type}/field | 
[**GetCloudTypes**](DefaultApi.md#GetCloudTypes) | **Get** /cloudcredential/cloudtype | 
[**GetCloudVmMetaInfo**](DefaultApi.md#GetCloudVmMetaInfo) | **Get** /cloudcredential/cloudtype/{type}/vminfo | 
[**GetCredential**](DefaultApi.md#GetCredential) | **Get** /cloudcredential/{cloudcredential_id} | 
[**GetDynamicProtection**](DefaultApi.md#GetDynamicProtection) | **Get** /dynamicprotection/{dynamicprotection_id} | Get individual dynamic protection entry details. It requires backupdr.managementServers.getDynamicProtection IAM permission
[**GetExternalGrammar**](DefaultApi.md#GetExternalGrammar) | **Get** /application.wadl/{path} | 
[**GetGCPProjects**](DefaultApi.md#GetGCPProjects) | **Get** /cloudcredential/{cloudcredential_id}/projects | 
[**GetMountVmParams**](DefaultApi.md#GetMountVmParams) | **Get** /cloudcredential/{cloudcredential_id}/region/{region}/mountinfo | 
[**GetRegions**](DefaultApi.md#GetRegions) | **Get** /cloudcredential/cloudtype/{type}/region | 
[**GetWadl**](DefaultApi.md#GetWadl) | **Get** /application.wadl | 
[**ListCredentials**](DefaultApi.md#ListCredentials) | **Get** /cloudcredential | 
[**ListDiskTypes**](DefaultApi.md#ListDiskTypes) | **Get** /cloudcredential/{cloudcredential_id}/disktypes | 
[**ListDynamicProtections**](DefaultApi.md#ListDynamicProtections) | **Get** /dynamicprotection | Get a list of dynamic protection entries. It requires backupdr.managementServers.listDynamicProtection IAM
[**PromoteExistingCredential**](DefaultApi.md#PromoteExistingCredential) | **Post** /cloudcredential/{cloudcredential_id}/promote | 
[**RunDynamicProtectionJob**](DefaultApi.md#RunDynamicProtectionJob) | **Post** /dynamicprotection/job | Run dynamic protection job. It requires backupdr.managementServers.assignBackupPlans IAM permission
[**SetDynamicProtectionJobConfig**](DefaultApi.md#SetDynamicProtectionJobConfig) | **Put** /dynamicprotection/jobconfig | Update dynamic protection job config. It requires backupdr.managementServers.assignBackupPlans IAM permission
[**TestCredential**](DefaultApi.md#TestCredential) | **Post** /cloudcredential/testconnection | 
[**UpdateCloudPool**](DefaultApi.md#UpdateCloudPool) | **Put** /cloudcredential/{cloudcredential_id}/cloudpool | 
[**UpdateCredential**](DefaultApi.md#UpdateCredential) | **Put** /cloudcredential/{cloudcredential_id} | 

# **AddCloudPool**
> AddCloudPool(ctx, cloudcredentialId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
 **optional** | ***DefaultApiAddCloudPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiAddCloudPoolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DiskPoolRest**](DiskPoolRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVm**
> AddVm(ctx, cloudcredentialId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
 **optional** | ***DefaultApiAddVmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiAddVmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CloudVmDiscoveryRest**](CloudVmDiscoveryRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountCloudCredentials**
> CountCloudCredentials(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountDynamicProtections**
> CountDynamicProtections(ctx, optional)
Get a count of total dynamic protection entries. It requires backupdr.managementServers.getDynamicProtection IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiCountDynamicProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiCountDynamicProtectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCredential**
> CreateCredential(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiCreateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiCreateCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CloudCredentialRest**](CloudCredentialRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDynamicProtection**
> DynamicProtectionRest CreateDynamicProtection(ctx, optional)
Create a new dynamic protection entry. It requires backupdr.managementServers.createDynamicProtection IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiCreateDynamicProtectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiCreateDynamicProtectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DynamicProtectionRest**](DynamicProtectionRest.md)|  | 

### Return type

[**DynamicProtectionRest**](DynamicProtectionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudPool**
> DeleteCloudPool(ctx, cloudcredentialId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
 **optional** | ***DefaultApiDeleteCloudPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiDeleteCloudPoolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DiskPoolRest**](DiskPoolRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCredential**
> DeleteCredential(ctx, cloudcredentialId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
 **optional** | ***DefaultApiDeleteCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiDeleteCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CloudCredentialRest**](CloudCredentialRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDynamicProtection**
> DeleteDynamicProtection(ctx, dynamicprotectionId)
Remove a dynamic protection entry. It requires backupdr.managementServers.deleteDynamicProtection IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dynamicprotectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverCloudVm**
> DiscoverCloudVm(ctx, cloudcredentialId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
 **optional** | ***DefaultApiDiscoverCloudVmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiDiscoverCloudVmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CloudVmDiscoveryRest**](CloudVmDiscoveryRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllZones**
> GetAllZones(ctx, cloudcredentialId, region)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
  **region** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudCredentialMetaInfo**
> GetCloudCredentialMetaInfo(ctx, type_)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudTypes**
> GetCloudTypes(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudVmMetaInfo**
> GetCloudVmMetaInfo(ctx, type_)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCredential**
> GetCredential(ctx, cloudcredentialId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDynamicProtection**
> DynamicProtectionRest GetDynamicProtection(ctx, dynamicprotectionId)
Get individual dynamic protection entry details. It requires backupdr.managementServers.getDynamicProtection IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dynamicprotectionId** | **string**|  | 

### Return type

[**DynamicProtectionRest**](DynamicProtectionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalGrammar**
> GetExternalGrammar(ctx, path)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGCPProjects**
> GetGCPProjects(ctx, cloudcredentialId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMountVmParams**
> GetMountVmParams(ctx, cloudcredentialId, region)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
  **region** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegions**
> GetRegions(ctx, type_, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
 **optional** | ***DefaultApiGetRegionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetRegionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudcredential** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWadl**
> GetWadl(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sun.wadl+xml, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCredentials**
> ListCredentials(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDiskTypes**
> ListDiskTypes(ctx, cloudcredentialId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
 **optional** | ***DefaultApiListDiskTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListDiskTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | **optional.String**|  | 
 **region** | **optional.String**|  | 
 **type_** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDynamicProtections**
> ListRestDynamicProtectionRest ListDynamicProtections(ctx, optional)
Get a list of dynamic protection entries. It requires backupdr.managementServers.listDynamicProtection IAM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiListDynamicProtectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListDynamicProtectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListRestDynamicProtectionRest**](ListRestDynamicProtectionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PromoteExistingCredential**
> PromoteExistingCredential(ctx, cloudcredentialId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
 **optional** | ***DefaultApiPromoteExistingCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiPromoteExistingCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CloudCredentialRest**](CloudCredentialRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunDynamicProtectionJob**
> InlineResponse200 RunDynamicProtectionJob(ctx, )
Run dynamic protection job. It requires backupdr.managementServers.assignBackupPlans IAM permission

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDynamicProtectionJobConfig**
> DynamicProtectionJobConfigRest SetDynamicProtectionJobConfig(ctx, optional)
Update dynamic protection job config. It requires backupdr.managementServers.assignBackupPlans IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiSetDynamicProtectionJobConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSetDynamicProtectionJobConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DynamicProtectionJobConfigRest**](DynamicProtectionJobConfigRest.md)|  | 

### Return type

[**DynamicProtectionJobConfigRest**](DynamicProtectionJobConfigRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestCredential**
> TestCredential(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiTestCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTestCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CloudCredentialRest**](CloudCredentialRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCloudPool**
> UpdateCloudPool(ctx, cloudcredentialId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
 **optional** | ***DefaultApiUpdateCloudPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUpdateCloudPoolOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DiskPoolRest**](DiskPoolRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCredential**
> UpdateCredential(ctx, cloudcredentialId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudcredentialId** | **string**|  | 
 **optional** | ***DefaultApiUpdateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUpdateCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CloudCredentialRest**](CloudCredentialRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

