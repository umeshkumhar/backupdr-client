# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupNow**](ApplicationApi.md#BackupNow) | **Post** /application/{application_id}/backup | Run a backup job. It requires backupdr.managementServers.manageBackups IAM permission
[**CountApplications**](ApplicationApi.md#CountApplications) | **Head** /application | Get a count of total applications matching the filters. It requires backupdr.managementServers.access IAM permission
[**CreateOptionForApp**](ApplicationApi.md#CreateOptionForApp) | **Post** /application/{application_id}/settableoption | Create a settable option for the particular application. It requires backupdr.managementServers.assignBackupPlans IAM permission
[**CreateWorkflow**](ApplicationApi.md#CreateWorkflow) | **Post** /application/{application_id}/workflow | Create new workflow for the particular application. It requires backupdr.managementServers.manageWorkflows IAM permission
[**DeleteApplication**](ApplicationApi.md#DeleteApplication) | **Delete** /application/{application_id} | Delete application. It requires backupdr.managementServers.manageApplications IAM permission
[**DeleteOptionForApp**](ApplicationApi.md#DeleteOptionForApp) | **Delete** /application/{application_id}/settableoption/{option_id} | Delete the particular option from the particular application. It requires backupdr.managementServers.assignBackupPlans IAM permission
[**DeleteWorkflow**](ApplicationApi.md#DeleteWorkflow) | **Delete** /application/{application_id}/workflow/{workflow_id} | Delete existing workflow. It requires backupdr.managementServers.manageWorkflows IAM permission
[**ExpireBackups**](ApplicationApi.md#ExpireBackups) | **Post** /application/{application_id}/expirebackup | Expires backups of the particular application. It requires backupdr.managementServers.manageExpiration IAM permission
[**GetAppClass**](ApplicationApi.md#GetAppClass) | **Get** /application/{application_id}/appclass | Get the particular application&#x27;s appclass metadata. It requires backupdr.managementServers.access IAM permission
[**GetAppClassByAppclassName**](ApplicationApi.md#GetAppClassByAppclassName) | **Get** /application/appclass/{appclass_name} | Get appclass metadata for the particular appclass name. It requires backupdr.managementServers.access IAM permission
[**GetAppClasses**](ApplicationApi.md#GetAppClasses) | **Get** /application/appclass | Get all available app classes from given cluster. It requires backupdr.managementServers.access IAM permission
[**GetApplication**](ApplicationApi.md#GetApplication) | **Get** /application/{application_id} | Get details on the particular application. It requires backupdr.managementServers.access IAM permission
[**GetOptionForApp**](ApplicationApi.md#GetOptionForApp) | **Get** /application/{application_id}/settableoption/{option_id} | Get the particular option of the particular application. It requires backupdr.managementServers.viewBackupPlans IAM permission
[**GetWorkflow**](ApplicationApi.md#GetWorkflow) | **Get** /application/{application_id}/workflow/{workflow_id} | Get individual workflow. It requires backupdr.managementServers.viewWorkflows IAM permission
[**ListActiveImages**](ApplicationApi.md#ListActiveImages) | **Get** /application/{application_id}/activeimage | Get active images for an application. It requires backupdr.managementServers.access IAM permission
[**ListApplicationTypes**](ApplicationApi.md#ListApplicationTypes) | **Get** /application/types | Get list of application types that are currently in the system. It requires backupdr.managementServers.access IAM permission
[**ListApplications**](ApplicationApi.md#ListApplications) | **Get** /application | List applications. It requires backupdr.managementServers.access IAM permission
[**ListOptionForApp**](ApplicationApi.md#ListOptionForApp) | **Get** /application/{application_id}/settableoption | List all existing settable options of the application. It requires backupdr.managementServers.access IAM permission
[**ListWorkflowsPerApp**](ApplicationApi.md#ListWorkflowsPerApp) | **Get** /application/{application_id}/workflow | Get list of workflows for the particular application. It requires backupdr.managementServers.viewWorkflows IAM permission
[**OptionsForListApplication**](ApplicationApi.md#OptionsForListApplication) | **Options** /application | Describes the fields available for filtering and sorting. It requires backupdr.managementServers.access IAM permission
[**SettableOptionMetadataForApp**](ApplicationApi.md#SettableOptionMetadataForApp) | **Options** /application/{application_id}/settableoption | Get settable option metadata of the particular application. It requires backupdr.managementServers.access IAM permission
[**SettableOptionMetadataForAppType**](ApplicationApi.md#SettableOptionMetadataForAppType) | **Options** /application/settableoption/{apptype} | Settable option metadata for the particular application type. It requires backupdr.managementServers.access IAM permission
[**UpdateApplication**](ApplicationApi.md#UpdateApplication) | **Put** /application/{application_id} | Update application data. It requires backupdr.managementServers.manageApplications and backupdr.managementServers.manageSensitiveData (for sensitive app) IAM permissions
[**UpdateOptionForApp**](ApplicationApi.md#UpdateOptionForApp) | **Put** /application/{application_id}/settableoption/{option_id} | Update the particular option of the particular application. It requires backupdr.managementServers.assignBackupPlans IAM permission
[**UpdateWorkflow**](ApplicationApi.md#UpdateWorkflow) | **Put** /application/{application_id}/workflow/{workflow_id} | Update existing workflow. It requires backupdr.managementServers.manageWorkflows IAM permission

# **BackupNow**
> BackupNow(ctx, applicationId, optional)
Run a backup job. It requires backupdr.managementServers.manageBackups IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
 **optional** | ***ApplicationApiBackupNowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiBackupNowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BackupNowRest**](BackupNowRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountApplications**
> CountApplications(ctx, optional)
Get a count of total applications matching the filters. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiCountApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiCountApplicationsOpts struct
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

# **CreateOptionForApp**
> AdvancedOptionRest CreateOptionForApp(ctx, applicationId, optional)
Create a settable option for the particular application. It requires backupdr.managementServers.assignBackupPlans IAM permission

Available options can be retrieved from the OPTIONS API. Existing options can be retrieved from GET API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
 **optional** | ***ApplicationApiCreateOptionForAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiCreateOptionForAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AdvancedOptionRest**](AdvancedOptionRest.md)|  | 

### Return type

[**AdvancedOptionRest**](AdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWorkflow**
> WorkflowRest CreateWorkflow(ctx, applicationId, optional)
Create new workflow for the particular application. It requires backupdr.managementServers.manageWorkflows IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
 **optional** | ***ApplicationApiCreateWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiCreateWorkflowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of WorkflowRest**](WorkflowRest.md)|  | 

### Return type

[**WorkflowRest**](WorkflowRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplication**
> DeleteApplication(ctx, applicationId)
Delete application. It requires backupdr.managementServers.manageApplications IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOptionForApp**
> DeleteOptionForApp(ctx, applicationId, optionId)
Delete the particular option from the particular application. It requires backupdr.managementServers.assignBackupPlans IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
  **optionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWorkflow**
> DeleteWorkflow(ctx, applicationId, workflowId)
Delete existing workflow. It requires backupdr.managementServers.manageWorkflows IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
  **workflowId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpireBackups**
> ExpireBackups(ctx, applicationId, optional)
Expires backups of the particular application. It requires backupdr.managementServers.manageExpiration IAM permission

Default to all backups. Optionally the type of backups can be specified in the payload (ExpireBackupRest). Valid type includes snapshot, dedup, remote-dedup and vault.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
 **optional** | ***ApplicationApiExpireBackupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiExpireBackupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExpireBackupRest**](ExpireBackupRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppClass**
> AppClassRest GetAppClass(ctx, applicationId, optional)
Get the particular application's appclass metadata. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
 **optional** | ***ApplicationApiGetAppClassOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiGetAppClassOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostid** | **optional.String**|  | 
 **isasm** | **optional.String**|  | 
 **operation** | **optional.String**|  | 
 **backupid** | **optional.Int64**|  | 

### Return type

[**AppClassRest**](AppClassRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppClassByAppclassName**
> AppClassRest GetAppClassByAppclassName(ctx, appclassName, optional)
Get appclass metadata for the particular appclass name. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appclassName** | **string**|  | 
 **optional** | ***ApplicationApiGetAppClassByAppclassNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiGetAppClassByAppclassNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostid** | **optional.String**|  | 
 **clusterid** | **optional.String**|  | 
 **isasm** | **optional.String**|  | 
 **operation** | **optional.String**|  | 
 **backupid** | **optional.Int64**|  | 

### Return type

[**AppClassRest**](AppClassRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppClasses**
> AppClassRest GetAppClasses(ctx, optional)
Get all available app classes from given cluster. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiGetAppClassesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiGetAppClassesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterid** | **optional.Int64**|  | 

### Return type

[**AppClassRest**](AppClassRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplication**
> ApplicationRest GetApplication(ctx, applicationId)
Get details on the particular application. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 

### Return type

[**ApplicationRest**](ApplicationRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOptionForApp**
> AdvancedOptionRest GetOptionForApp(ctx, applicationId, optionId)
Get the particular option of the particular application. It requires backupdr.managementServers.viewBackupPlans IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
  **optionId** | **string**|  | 

### Return type

[**AdvancedOptionRest**](AdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflow**
> WorkflowRest GetWorkflow(ctx, applicationId, workflowId)
Get individual workflow. It requires backupdr.managementServers.viewWorkflows IAM permission

Note: Workflow_id is not unique on AGM (it's actually the id on the appliance). It's unique per application_id though

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
  **workflowId** | **string**|  | 

### Return type

[**WorkflowRest**](WorkflowRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListActiveImages**
> ListBackupRest ListActiveImages(ctx, applicationId)
Get active images for an application. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 

### Return type

[**ListBackupRest**](ListBackupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplicationTypes**
> ListStringRest ListApplicationTypes(ctx, )
Get list of application types that are currently in the system. It requires backupdr.managementServers.access IAM permission

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListStringRest**](ListStringRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplications**
> ListApplicationRest ListApplications(ctx, optional)
List applications. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiListApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiListApplicationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListApplicationRest**](ListApplicationRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOptionForApp**
> ListAdvancedOptionRest ListOptionForApp(ctx, applicationId)
List all existing settable options of the application. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 

### Return type

[**ListAdvancedOptionRest**](ListAdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWorkflowsPerApp**
> ListWorkflowRest ListWorkflowsPerApp(ctx, applicationId)
Get list of workflows for the particular application. It requires backupdr.managementServers.viewWorkflows IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 

### Return type

[**ListWorkflowRest**](ListWorkflowRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionsForListApplication**
> OptionsRest OptionsForListApplication(ctx, )
Describes the fields available for filtering and sorting. It requires backupdr.managementServers.access IAM permission

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OptionsRest**](OptionsRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettableOptionMetadataForApp**
> string SettableOptionMetadataForApp(ctx, applicationId)
Get settable option metadata of the particular application. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettableOptionMetadataForAppType**
> string SettableOptionMetadataForAppType(ctx, apptype)
Settable option metadata for the particular application type. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apptype** | **string**|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplication**
> UpdateApplication(ctx, applicationId, optional)
Update application data. It requires backupdr.managementServers.manageApplications and backupdr.managementServers.manageSensitiveData (for sensitive app) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
 **optional** | ***ApplicationApiUpdateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiUpdateApplicationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApplicationRest**](ApplicationRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOptionForApp**
> AdvancedOptionRest UpdateOptionForApp(ctx, applicationId, optionId, optional)
Update the particular option of the particular application. It requires backupdr.managementServers.assignBackupPlans IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
  **optionId** | **string**|  | 
 **optional** | ***ApplicationApiUpdateOptionForAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiUpdateOptionForAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AdvancedOptionRest**](AdvancedOptionRest.md)|  | 

### Return type

[**AdvancedOptionRest**](AdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWorkflow**
> WorkflowRest UpdateWorkflow(ctx, applicationId, workflowId, optional)
Update existing workflow. It requires backupdr.managementServers.manageWorkflows IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **string**|  | 
  **workflowId** | **string**|  | 
 **optional** | ***ApplicationApiUpdateWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiUpdateWorkflowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WorkflowRest**](WorkflowRest.md)|  | 

### Return type

[**WorkflowRest**](WorkflowRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

