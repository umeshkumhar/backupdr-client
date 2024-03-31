# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissions**](UserSessionApi.md#GetPermissions) | **Get** /session/permissions | Get IAM permissions for the user
[**GetSessionInfo**](UserSessionApi.md#GetSessionInfo) | **Get** /session/current | Get session information, including user preferences
[**Login**](UserSessionApi.md#Login) | **Post** /session | Register an user session
[**Logout**](UserSessionApi.md#Logout) | **Delete** /session/current | Deregister the current user session
[**PutUserpref**](UserSessionApi.md#PutUserpref) | **Put** /session/current | Update user preferences

# **GetPermissions**
> GetPermissions(ctx, )
Get IAM permissions for the user

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

# **GetSessionInfo**
> SessionRest GetSessionInfo(ctx, )
Get session information, including user preferences

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

# **Login**
> SessionRest Login(ctx, )
Register an user session

This API will generate a session id that will be your API key for other calls.

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

# **Logout**
> Logout(ctx, )
Deregister the current user session

Deregister the current user session

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

# **PutUserpref**
> PutUserpref(ctx, optional)
Update user preferences

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserSessionApiPutUserprefOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserSessionApiPutUserprefOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SessionRest**](SessionRest.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

