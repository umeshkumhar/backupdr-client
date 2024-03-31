# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountApplianceUpdatesInstallationJobs**](ApplianceUpdateInstallationJobApi.md#CountApplianceUpdatesInstallationJobs) | **Head** /applianceupdateinstallationjob | Get the count of total appliance update installation jobs.
[**GetApplianceUpdateInstallationJob**](ApplianceUpdateInstallationJobApi.md#GetApplianceUpdateInstallationJob) | **Get** /applianceupdateinstallationjob/{update_id} | Get the appliance update job details.
[**GetReadMe**](ApplianceUpdateInstallationJobApi.md#GetReadMe) | **Get** /applianceupdateinstallationjob/readme | Get readme of the update id passed.
[**GetUpdateLogs**](ApplianceUpdateInstallationJobApi.md#GetUpdateLogs) | **Get** /applianceupdateinstallationjob/log | Get logs of the update id passed.
[**ListApplianceUpdatesInstallationjobs**](ApplianceUpdateInstallationJobApi.md#ListApplianceUpdatesInstallationjobs) | **Get** /applianceupdateinstallationjob | Get the list of update installation jobs.
[**OptionsForList**](ApplianceUpdateInstallationJobApi.md#OptionsForList) | **Options** /applianceupdateinstallationjob | Describes the fields available for filtering and sorting
[**UpdateNotification**](ApplianceUpdateInstallationJobApi.md#UpdateNotification) | **Put** /applianceupdateinstallationjob/notification/{id} | Update the acknowledge status of the notification to true

# **CountApplianceUpdatesInstallationJobs**
> CountApplianceUpdatesInstallationJobs(ctx, optional)
Get the count of total appliance update installation jobs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateInstallationJobApiCountApplianceUpdatesInstallationJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateInstallationJobApiCountApplianceUpdatesInstallationJobsOpts struct
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

# **GetApplianceUpdateInstallationJob**
> ApplianceUpdateInstallationJobRest GetApplianceUpdateInstallationJob(ctx, updateId)
Get the appliance update job details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateId** | **int64**|  | 

### Return type

[**ApplianceUpdateInstallationJobRest**](ApplianceUpdateInstallationJobRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReadMe**
> ApplianceUpdateReadmeRest GetReadMe(ctx, optional)
Get readme of the update id passed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateInstallationJobApiGetReadMeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateInstallationJobApiGetReadMeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**|  | 

### Return type

[**ApplianceUpdateReadmeRest**](ApplianceUpdateReadmeRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpdateLogs**
> ListApplianceUpdateLogsRest GetUpdateLogs(ctx, optional)
Get logs of the update id passed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateInstallationJobApiGetUpdateLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateInstallationJobApiGetUpdateLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**|  | 

### Return type

[**ListApplianceUpdateLogsRest**](ListApplianceUpdateLogsRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplianceUpdatesInstallationjobs**
> ListApplianceUpdateInstallationJobRest ListApplianceUpdatesInstallationjobs(ctx, optional)
Get the list of update installation jobs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateInstallationJobApiListApplianceUpdatesInstallationjobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateInstallationJobApiListApplianceUpdatesInstallationjobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListApplianceUpdateInstallationJobRest**](ListApplianceUpdateInstallationJobRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionsForList**
> OptionsRest OptionsForList(ctx, )
Describes the fields available for filtering and sorting

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

# **UpdateNotification**
> UpdateNotification(ctx, id)
Update the acknowledge status of the notification to true

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

