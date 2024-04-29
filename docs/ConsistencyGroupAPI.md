# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountConsistencyGroups**](ConsistencyGroupApi.md#CountConsistencyGroups) | **Head** /consistencygroup | Get a count of total consistency groups matching the filters. It requires backupdr.managementServers.access IAM permission
[**CreateConsistencyGroup**](ConsistencyGroupApi.md#CreateConsistencyGroup) | **Post** /consistencygroup | Create a new consistency group. It requires backupdr.managementServers.manageApplications IAM permission
[**DeleteConsistencyGroup**](ConsistencyGroupApi.md#DeleteConsistencyGroup) | **Delete** /consistencygroup/{group_id} | Remove a consistency group. It requires backupdr.managementServers.manageApplications IAM permission
[**GetConsistencyGroup**](ConsistencyGroupApi.md#GetConsistencyGroup) | **Get** /consistencygroup/{group_id} | Get individual consistency group details. It requires backupdr.managementServers.access IAM permission
[**GetConsistencyGroupMember**](ConsistencyGroupApi.md#GetConsistencyGroupMember) | **Get** /consistencygroup/{group_id}/member | Get consistency group&#x27;s members. It requires backupdr.managementServers.access IAM permission
[**ListConsistencyGroups**](ConsistencyGroupApi.md#ListConsistencyGroups) | **Get** /consistencygroup | Get a list of consistency groups. It requires backupdr.managementServers.access IAM permission
[**ModifyConsistencyGroupMember**](ConsistencyGroupApi.md#ModifyConsistencyGroupMember) | **Post** /consistencygroup/{group_id}/member | Incrementally add/delete consistency group members. It requires backupdr.managementServers.manageApplications IAM permission
[**OptionsForListConsistencyGroup**](ConsistencyGroupApi.md#OptionsForListConsistencyGroup) | **Options** /consistencygroup | Describes the fields available for filtering and sorting. It requires backupdr.managementServers.access IAM permission
[**UpdateConsistencyGroup**](ConsistencyGroupApi.md#UpdateConsistencyGroup) | **Put** /consistencygroup/{group_id} | Update a consistency group. It requires backupdr.managementServers.manageApplications IAM permission

# **CountConsistencyGroups**
> CountConsistencyGroups(ctx, optional)
Get a count of total consistency groups matching the filters. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConsistencyGroupApiCountConsistencyGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConsistencyGroupApiCountConsistencyGroupsOpts struct
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

# **CreateConsistencyGroup**
> ConsistencyGroupRest CreateConsistencyGroup(ctx, optional)
Create a new consistency group. It requires backupdr.managementServers.manageApplications IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConsistencyGroupApiCreateConsistencyGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConsistencyGroupApiCreateConsistencyGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ConsistencyGroupRest**](ConsistencyGroupRest.md)|  | 

### Return type

[**ConsistencyGroupRest**](ConsistencyGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConsistencyGroup**
> DeleteConsistencyGroup(ctx, groupId)
Remove a consistency group. It requires backupdr.managementServers.manageApplications IAM permission

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

# **GetConsistencyGroup**
> ConsistencyGroupRest GetConsistencyGroup(ctx, groupId)
Get individual consistency group details. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

[**ConsistencyGroupRest**](ConsistencyGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsistencyGroupMember**
> ListApplicationRest GetConsistencyGroupMember(ctx, groupId, optional)
Get consistency group's members. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***ConsistencyGroupApiGetConsistencyGroupMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConsistencyGroupApiGetConsistencyGroupMemberOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | 

### Return type

[**ListApplicationRest**](ListApplicationRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConsistencyGroups**
> ListConsistencyGroupRest ListConsistencyGroups(ctx, optional)
Get a list of consistency groups. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConsistencyGroupApiListConsistencyGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConsistencyGroupApiListConsistencyGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListConsistencyGroupRest**](ListConsistencyGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyConsistencyGroupMember**
> ModifyConsistencyGroupMember(ctx, groupId, optional)
Incrementally add/delete consistency group members. It requires backupdr.managementServers.manageApplications IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***ConsistencyGroupApiModifyConsistencyGroupMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConsistencyGroupApiModifyConsistencyGroupMemberOpts struct
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

# **OptionsForListConsistencyGroup**
> OptionsRest OptionsForListConsistencyGroup(ctx, )
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

# **UpdateConsistencyGroup**
> ConsistencyGroupRest UpdateConsistencyGroup(ctx, groupId, optional)
Update a consistency group. It requires backupdr.managementServers.manageApplications IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***ConsistencyGroupApiUpdateConsistencyGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConsistencyGroupApiUpdateConsistencyGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ConsistencyGroupRest**](ConsistencyGroupRest.md)|  | 

### Return type

[**ConsistencyGroupRest**](ConsistencyGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

