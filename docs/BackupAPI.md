# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneBackup**](BackupApi.md#CloneBackup) | **Post** /backup/{backup_id}/clone | Clone a specific backup. It requires backupdr.managementServers.manageClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**CountBackups**](BackupApi.md#CountBackups) | **Head** /backup | Get a count of total backups matching the filters. It requires backupdr.managementServers.access IAM permission
[**DeleteBackup**](BackupApi.md#DeleteBackup) | **Delete** /backup/{backup_id} | Delete a specific backup. It requires backupdr.managementServers.manageExpiration and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**ExpireBackup**](BackupApi.md#ExpireBackup) | **Post** /backup/{backup_id}/expire | Expire a specific backup. It requires backupdr.managementServers.manageExpiration and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**FetchApplicationOptions**](BackupApi.md#FetchApplicationOptions) | **Get** /backup/{backupid}/applicationOptions | Gets the dynamic list of application options and corresponding default value (if any) for given backup image. It requires backupdr.managementServers.manageBackups and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**GetBackup**](BackupApi.md#GetBackup) | **Get** /backup/{backup_id} | Get individual backup details. It requires backupdr.managementServers.access and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**GetDiskMapping**](BackupApi.md#GetDiskMapping) | **Get** /backup/{backupid}/diskmapping | Gets disk mapping options for restore (source disks, target disks, RAC node list etc). It requires backupdr.managementServers.manageBackups and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**ListBackups**](BackupApi.md#ListBackups) | **Get** /backup | Get a list of backups. It requires backupdr.managementServers.access IAM permission
[**LiveCloneBackup**](BackupApi.md#LiveCloneBackup) | **Post** /backup/{backup_id}/liveclone | Create live-clone of a specific backup. It requires backupdr.managementServers.manageLiveClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**MigrateRestoreMountBackup**](BackupApi.md#MigrateRestoreMountBackup) | **Post** /backup/{backupid}/restoremigrate | Submit restore-migrate for Oracle restore-mounted image. It requires backupdr.managementServers.manageMigrations and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**MountBackup**](BackupApi.md#MountBackup) | **Post** /backup/{backup_id}/mount | Mount a specific backup. It require backupdr.managementServers.manageMounts IAM permission for mounting any backup, backupdr.managementServers.manageRestores IAM permission for restoring GCE backup, and backupdr.managementServers.accessSensitiveData IAM permission to access sensitive backup
[**MountMigratePreflight**](BackupApi.md#MountMigratePreflight) | **Post** /backup/{backupid}/mountmigratepreflight | Performs preflight check for Oracle restore-mount/restore-migrate operations. It requires backupdr.managementServers.manageMigrations and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**OptionsForListBackup**](BackupApi.md#OptionsForListBackup) | **Options** /backup | Describes the fields available for filtering and sorting. It requires backupdr.managementServers.access IAM permission
[**PrepMountLiveCloneNew**](BackupApi.md#PrepMountLiveCloneNew) | **Post** /backup/{backup_id}/prepmount | Prep-mount a specific backup. It requires backupdr.managementServers.manageLiveClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**PrepUnmountBackup**](BackupApi.md#PrepUnmountBackup) | **Post** /backup/{backup_id}/prepunmount | Prep-unmount a specific backup. It requires backupdr.managementServers.manageLiveClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**RefreshLiveClone**](BackupApi.md#RefreshLiveClone) | **Post** /backup/{backup_id}/refresh | Refresh a live-clone backup. It requires backupdr.managementServers.manageLiveClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**RestoreBackup**](BackupApi.md#RestoreBackup) | **Post** /backup/{backup_id}/restore | Restore an application from a specific backup. It requires backupdr.managementServers.manageRestores and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**UnmountBackup**](BackupApi.md#UnmountBackup) | **Post** /backup/{backup_id}/unmount | Unmount a specific backup. It requires backupdr.managementServers.manageMounts and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions
[**UpdateBackup**](BackupApi.md#UpdateBackup) | **Put** /backup/{backup_id} | Update the specific backup. It requires backupdr.managementServers.manageExpiration IAM permission to update expiration, backupdr.managementServers.manageBackups IAM permission to update rest of the fields, backupdr.managementServers.accessSensitiveData IAM permission to update sensitive backup, and backupdr.managementServers.manageSensitiveData IAM permission to update sensitivity flag

# **CloneBackup**
> CloneBackup(ctx, backupId, optional)
Clone a specific backup. It requires backupdr.managementServers.manageClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiCloneBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiCloneBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CloneRest**](CloneRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountBackups**
> CountBackups(ctx, optional)
Get a count of total backups matching the filters. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BackupApiCountBackupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiCountBackupsOpts struct
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

# **DeleteBackup**
> DeleteBackup(ctx, backupId)
Delete a specific backup. It requires backupdr.managementServers.manageExpiration and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpireBackup**
> ExpireBackup(ctx, backupId, optional)
Expire a specific backup. It requires backupdr.managementServers.manageExpiration and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiExpireBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiExpireBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchApplicationOptions**
> AppClassRest FetchApplicationOptions(ctx, backupid)
Gets the dynamic list of application options and corresponding default value (if any) for given backup image. It requires backupdr.managementServers.manageBackups and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupid** | **string**| Id of the backup image whose application options are required. | 

### Return type

[**AppClassRest**](AppClassRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackup**
> BackupRest GetBackup(ctx, backupId)
Get individual backup details. It requires backupdr.managementServers.access and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 

### Return type

[**BackupRest**](BackupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiskMapping**
> DiskMappingRest GetDiskMapping(ctx, backupid, optional)
Gets disk mapping options for restore (source disks, target disks, RAC node list etc). It requires backupdr.managementServers.manageBackups and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupid** | **string**| Id of the backup image to be restored | 
 **optional** | ***BackupApiGetDiskMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiGetDiskMappingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostid** | **optional.String**| Restore target host | 
 **targetstoragetype** | **optional.String**| Restore Target Storage Type (ASM or FS) | 

### Return type

[**DiskMappingRest**](DiskMappingRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBackups**
> ListBackupRest ListBackups(ctx, optional)
Get a list of backups. It requires backupdr.managementServers.access IAM permission

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BackupApiListBackupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiListBackupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListBackupRest**](ListBackupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LiveCloneBackup**
> LiveCloneBackup(ctx, backupId, optional)
Create live-clone of a specific backup. It requires backupdr.managementServers.manageLiveClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiLiveCloneBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiLiveCloneBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of LiveCloneRest**](LiveCloneRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateRestoreMountBackup**
> MigrateRestoreMountBackup(ctx, backupid, optional)
Submit restore-migrate for Oracle restore-mounted image. It requires backupdr.managementServers.manageMigrations and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupid** | **string**| Id of the restore-mounted image to be migrated | 
 **optional** | ***BackupApiMigrateRestoreMountBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiMigrateRestoreMountBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RestoreMigrateRest**](RestoreMigrateRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MountBackup**
> MountBackup(ctx, backupId, optional)
Mount a specific backup. It require backupdr.managementServers.manageMounts IAM permission for mounting any backup, backupdr.managementServers.manageRestores IAM permission for restoring GCE backup, and backupdr.managementServers.accessSensitiveData IAM permission to access sensitive backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiMountBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiMountBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MountRest**](MountRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MountMigratePreflight**
> RestorePreflightRest MountMigratePreflight(ctx, backupid, optional)
Performs preflight check for Oracle restore-mount/restore-migrate operations. It requires backupdr.managementServers.manageMigrations and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupid** | **string**| Id of the backup image which will be restore-mounted or restore-migrated. | 
 **optional** | ***BackupApiMountMigratePreflightOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiMountMigratePreflightOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PreflightRest**](PreflightRest.md)|  | 

### Return type

[**RestorePreflightRest**](RestorePreflightRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionsForListBackup**
> OptionsRest OptionsForListBackup(ctx, )
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

# **PrepMountLiveCloneNew**
> PrepMountLiveCloneNew(ctx, backupId, optional)
Prep-mount a specific backup. It requires backupdr.managementServers.manageLiveClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiPrepMountLiveCloneNewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiPrepMountLiveCloneNewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PrepmountRest**](PrepmountRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrepUnmountBackup**
> PrepUnmountBackup(ctx, backupId, optional)
Prep-unmount a specific backup. It requires backupdr.managementServers.manageLiveClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiPrepUnmountBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiPrepUnmountBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UnmountRest**](UnmountRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshLiveClone**
> RefreshLiveClone(ctx, backupId, optional)
Refresh a live-clone backup. It requires backupdr.managementServers.manageLiveClones and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiRefreshLiveCloneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiRefreshLiveCloneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RefreshRest**](RefreshRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreBackup**
> RestoreBackup(ctx, backupId, optional)
Restore an application from a specific backup. It requires backupdr.managementServers.manageRestores and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiRestoreBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiRestoreBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RestoreRest**](RestoreRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnmountBackup**
> UnmountBackup(ctx, backupId, optional)
Unmount a specific backup. It requires backupdr.managementServers.manageMounts and backupdr.managementServers.accessSensitiveData (for sensitive backup) IAM permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiUnmountBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiUnmountBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UnmountRest**](UnmountRest.md)|  | 
 **returnjob** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBackup**
> BackupRest UpdateBackup(ctx, backupId, optional)
Update the specific backup. It requires backupdr.managementServers.manageExpiration IAM permission to update expiration, backupdr.managementServers.manageBackups IAM permission to update rest of the fields, backupdr.managementServers.accessSensitiveData IAM permission to update sensitive backup, and backupdr.managementServers.manageSensitiveData IAM permission to update sensitivity flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **backupId** | **string**|  | 
 **optional** | ***BackupApiUpdateBackupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackupApiUpdateBackupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BackupRest**](BackupRest.md)|  | 

### Return type

[**BackupRest**](BackupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

