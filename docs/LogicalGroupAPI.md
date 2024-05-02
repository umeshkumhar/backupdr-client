# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountLogicalGroups**](LogicalGroupApi.md#CountLogicalGroups) | **Head** /logicalgroup | Get a count of total logical groups matching the filters. It requires backupdr.managementServers.access IAM permission
[**CreateLogicalGroup**](LogicalGroupApi.md#CreateLogicalGroup) | **Post** /logicalgroup | Create a new logical group. It requires backupdr.managementServers.manageApplications IAM permission
[**CreateLogicalGroupSla**](LogicalGroupApi.md#CreateLogicalGroupSla) | **Post** /logicalgroup/{group_id}/sla | Protect a logical group. It creates individual SLAs for all members. It requires backupdr.managementServers.assignBackupPlans IAM permission
[**DeleteLogicalGroup**](LogicalGroupApi.md#DeleteLogicalGroup) | **Delete** /logicalgroup/{group_id} | Remove a logical group. It requires backupdr.managementServers.manageApplications IAM permission
[**DeleteLogicalGroupSla**](LogicalGroupApi.md#DeleteLogicalGroupSla) | **Delete** /logicalgroup/{group_id}/sla | Unprotect a logical group. It removes SLAs for all members. It requires backupdr.managementServers.assignBackupPlans IAM permission
[**GetLogicalGroup**](LogicalGroupApi.md#GetLogicalGroup) | **Get** /logicalgroup/{group_id} | Get individual logical group details. It requires backupdr.managementServers.access IAM permission
[**ListLogicalGroup**](LogicalGroupApi.md#ListLogicalGroup) | **Get** /logicalgroup | Get a list of logical groups. It requires backupdr.managementServers.access IAM permission
[**ListLogicalGroupMembers**](LogicalGroupApi.md#ListLogicalGroupMembers) | **Get** /logicalgroup/{group_id}/member | Get logical group&#x27;s members. It requires backupdr.managementServers.access IAM permission
[**ModifyLogicalGroupMembers**](LogicalGroupApi.md#ModifyLogicalGroupMembers) | **Post** /logicalgroup/{group_id}/member | Incrementally add/delete logical group members. It requires backupdr.managementServers.manageApplications IAM permission
[**OptionsForListLogicalGroup**](LogicalGroupApi.md#OptionsForListLogicalGroup) | **Options** /logicalgroup | Describes the fields available for filtering and sorting. It requires backupdr.managementServers.access IAM permission
[**UpdateLogicalGroup**](LogicalGroupApi.md#UpdateLogicalGroup) | **Put** /logicalgroup/{group_id} | Update a logical group. It requires backupdr.managementServers.manageApplications IAM permission
[**UpdateLogicalGroupSla**](LogicalGroupApi.md#UpdateLogicalGroupSla) | **Put** /logicalgroup/{group_id}/sla | Update the current SLAs for a logical group. It updates individual SLAs for all members. It requires backupdr.managementServers.assignBackupPlans IAM permission

# **CountLogicalGroups**
> CountLogicalGroups(ctx, optional)
Get a count of total logical groups matching the filters. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogicalGroupApiCountLogicalGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalGroupApiCountLogicalGroupsOpts struct
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

# **CreateLogicalGroup**
> LogicalGroupRest CreateLogicalGroup(ctx, optional)
Create a new logical group. It requires backupdr.managementServers.manageApplications IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogicalGroupApiCreateLogicalGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalGroupApiCreateLogicalGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LogicalGroupRest**](LogicalGroupRest.md)|  | 

### Return type

[**LogicalGroupRest**](LogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLogicalGroupSla**
> CreateLogicalGroupSla(ctx, groupId, optional)
Protect a logical group. It creates individual SLAs for all members. It requires backupdr.managementServers.assignBackupPlans IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***LogicalGroupApiCreateLogicalGroupSlaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalGroupApiCreateLogicalGroupSlaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SlaRest**](SlaRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogicalGroup**
> DeleteLogicalGroup(ctx, groupId)
Remove a logical group. It requires backupdr.managementServers.manageApplications IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogicalGroupSla**
> DeleteLogicalGroupSla(ctx, groupId)
Unprotect a logical group. It removes SLAs for all members. It requires backupdr.managementServers.assignBackupPlans IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalGroup**
> LogicalGroupRest GetLogicalGroup(ctx, groupId)
Get individual logical group details. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

[**LogicalGroupRest**](LogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLogicalGroup**
> ListLogicalGroupRest ListLogicalGroup(ctx, optional)
Get a list of logical groups. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogicalGroupApiListLogicalGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalGroupApiListLogicalGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListLogicalGroupRest**](ListLogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLogicalGroupMembers**
> ListApplicationRest ListLogicalGroupMembers(ctx, groupId)
Get logical group's members. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

[**ListApplicationRest**](ListApplicationRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyLogicalGroupMembers**
> ModifyLogicalGroupMembers(ctx, groupId, optional)
Incrementally add/delete logical group members. It requires backupdr.managementServers.manageApplications IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***LogicalGroupApiModifyLogicalGroupMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalGroupApiModifyLogicalGroupMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []MembershipChangeRest**](MembershipChangeRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionsForListLogicalGroup**
> OptionsRest OptionsForListLogicalGroup(ctx, )
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

# **UpdateLogicalGroup**
> LogicalGroupRest UpdateLogicalGroup(ctx, groupId, optional)
Update a logical group. It requires backupdr.managementServers.manageApplications IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***LogicalGroupApiUpdateLogicalGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalGroupApiUpdateLogicalGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of LogicalGroupRest**](LogicalGroupRest.md)|  | 

### Return type

[**LogicalGroupRest**](LogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLogicalGroupSla**
> LogicalGroupRest UpdateLogicalGroupSla(ctx, groupId, optional)
Update the current SLAs for a logical group. It updates individual SLAs for all members. It requires backupdr.managementServers.assignBackupPlans IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***LogicalGroupApiUpdateLogicalGroupSlaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogicalGroupApiUpdateLogicalGroupSlaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SlaRest**](SlaRest.md)|  | 

### Return type

[**LogicalGroupRest**](LogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

