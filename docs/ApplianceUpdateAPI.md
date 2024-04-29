# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountApplianceUpdates**](ApplianceUpdateApi.md#CountApplianceUpdates) | **Head** /applianceupdate | Get a count of total applianceupdates matching the filters. It requires backupdr.managementServers.viewSystem IAM permission
[**DeleteSchedule**](ApplianceUpdateApi.md#DeleteSchedule) | **Delete** /applianceupdate/schedule | Deletes the schedule for update. It requires backupdr.managementServers.manageSystem IAM permission
[**DiscoverUpdates**](ApplianceUpdateApi.md#DiscoverUpdates) | **Post** /applianceupdate/discover | Discover new updates for the managed appliances. It requires backupdr.managementServers.manageSystem IAM permission
[**GetApplianceUpdate**](ApplianceUpdateApi.md#GetApplianceUpdate) | **Get** /applianceupdate/{update_id} | Get update details. It requires backupdr.managementServers.viewSystem IAM permission
[**GetNotifications**](ApplianceUpdateApi.md#GetNotifications) | **Get** /applianceupdate/notification | Fetch the notifications that needs to be displayed to the user. It requires backupdr.managementServers.access IAM permission
[**GetReadMeForApplianceUpdate**](ApplianceUpdateApi.md#GetReadMeForApplianceUpdate) | **Get** /applianceupdate/readme | Get readme with respect to the update id passed. It requires backupdr.managementServers.viewSystem IAM permission
[**InstallApplianceUpdateNow**](ApplianceUpdateApi.md#InstallApplianceUpdateNow) | **Post** /applianceupdate/installnow | Install the updates now. It requires backupdr.managementServers.manageSystem IAM permission
[**ListApplianceUpdates**](ApplianceUpdateApi.md#ListApplianceUpdates) | **Get** /applianceupdate | Get the list of actionable updates. It requires backupdr.managementServers.viewSystem IAM permission
[**OptionsForListApplianceUpdates**](ApplianceUpdateApi.md#OptionsForListApplianceUpdates) | **Options** /applianceupdate | Describes the fields available for filtering and sorting. It requires backupdr.managementServers.access IAM permission
[**UpdateNotificationForApplianceUpdate**](ApplianceUpdateApi.md#UpdateNotificationForApplianceUpdate) | **Put** /applianceupdate/notification | Update the acknowledge status of the notification to true
[**UpdateSchedule**](ApplianceUpdateApi.md#UpdateSchedule) | **Post** /applianceupdate/schedule | Creates the schedule for update. It requires backupdr.managementServers.manageSystem IAM permission

# **CountApplianceUpdates**
> CountApplianceUpdates(ctx, optional)
Get a count of total applianceupdates matching the filters. It requires backupdr.managementServers.viewSystem IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateApiCountApplianceUpdatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateApiCountApplianceUpdatesOpts struct
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

# **DeleteSchedule**
> DeleteSchedule(ctx, optional)
Deletes the schedule for update. It requires backupdr.managementServers.manageSystem IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateApiDeleteScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateApiDeleteScheduleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdateListRest**](UpdateListRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverUpdates**
> DiscoverUpdates(ctx, )
Discover new updates for the managed appliances. It requires backupdr.managementServers.manageSystem IAM permission

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

# **GetApplianceUpdate**
> ApplianceUpdateRest GetApplianceUpdate(ctx, updateId)
Get update details. It requires backupdr.managementServers.viewSystem IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateId** | **int64**|  | 

### Return type

[**ApplianceUpdateRest**](ApplianceUpdateRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotifications**
> ListRest GetNotifications(ctx, )
Fetch the notifications that needs to be displayed to the user. It requires backupdr.managementServers.access IAM permission

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListRest**](ListRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReadMeForApplianceUpdate**
> ApplianceUpdateReadmeRest GetReadMeForApplianceUpdate(ctx, optional)
Get readme with respect to the update id passed. It requires backupdr.managementServers.viewSystem IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateApiGetReadMeForApplianceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateApiGetReadMeForApplianceUpdateOpts struct
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

# **InstallApplianceUpdateNow**
> InstallApplianceUpdateNow(ctx, optional)
Install the updates now. It requires backupdr.managementServers.manageSystem IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateApiInstallApplianceUpdateNowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateApiInstallApplianceUpdateNowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdateListRest**](UpdateListRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplianceUpdates**
> ListConsolidatedApplianceUpdatesRest ListApplianceUpdates(ctx, optional)
Get the list of actionable updates. It requires backupdr.managementServers.viewSystem IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateApiListApplianceUpdatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateApiListApplianceUpdatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListConsolidatedApplianceUpdatesRest**](ListConsolidatedApplianceUpdatesRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionsForListApplianceUpdates**
> OptionsRest OptionsForListApplianceUpdates(ctx, )
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

# **UpdateNotificationForApplianceUpdate**
> UpdateNotificationForApplianceUpdate(ctx, optional)
Update the acknowledge status of the notification to true

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateApiUpdateNotificationForApplianceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateApiUpdateNotificationForApplianceUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApplianceUpdateNotificationRest**](ApplianceUpdateNotificationRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchedule**
> UpdateSchedule(ctx, optional)
Creates the schedule for update. It requires backupdr.managementServers.manageSystem IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplianceUpdateApiUpdateScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplianceUpdateApiUpdateScheduleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdateListRest**](UpdateListRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

