# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDbDumpUploadUrls**](ManageACLApi.md#GetDbDumpUploadUrls) | **Post** /manageacl/migration/getdbdumpuploadurls | 
[**GetRMMigrationStatus**](ManageACLApi.md#GetRMMigrationStatus) | **Get** /manageacl/migration/rmmigrationstatus | 
[**IsInMigrationMode**](ManageACLApi.md#IsInMigrationMode) | **Get** /manageacl/migration/isinmigrationmode | 
[**PromoteUser**](ManageACLApi.md#PromoteUser) | **Put** /manageacl/promoteUser | To promote the IAP Manage Acl users to Administrator role
[**StoreSecretKeys**](ManageACLApi.md#StoreSecretKeys) | **Post** /manageacl/migration/savesecret | 

# **GetDbDumpUploadUrls**
> GetDbDumpUploadUrls(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManageACLApiGetDbDumpUploadUrlsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManageACLApiGetDbDumpUploadUrlsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rmExists** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRMMigrationStatus**
> GetRMMigrationStatus(ctx, )


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

# **IsInMigrationMode**
> IsInMigrationMode(ctx, )


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

# **PromoteUser**
> SessionRest PromoteUser(ctx, )
To promote the IAP Manage Acl users to Administrator role

If the caller identity has IAM ManageInternalACL permission, this call will grant the current user session temporary (up to 5 minutes) AGM Administrator role (which can be granted explicitly alternatively) and temporary full privileges with AGM access control management (such as assigning users to organizations/roles and modifying role/right definitions). If the caller identity doesn't have the IAM ManageInternalACL permission, this call fails with 403 not authorized

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SessionRest**](SessionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreSecretKeys**
> StoreSecretKeys(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManageACLApiStoreSecretKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManageACLApiStoreSecretKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ListNameValueRest**](ListNameValueRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

