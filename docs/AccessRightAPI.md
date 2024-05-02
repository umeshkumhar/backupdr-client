# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRight**](AccessRightApi.md#GetRight) | **Get** /right/{right_name} | Get details for a specific right. It requires System View right.
[**ListRights**](AccessRightApi.md#ListRights) | **Get** /right | Get a list of rights in the system. Rights are predefined and cannot be changed. It requires System View right.
[**OptionsForList1**](AccessRightApi.md#OptionsForList1) | **Options** /right | Describes the fields available for filtering and sorting

# **GetRight**
> RightRest GetRight(ctx, rightName)
Get details for a specific right. It requires System View right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rightName** | **string**|  | 

### Return type

[**RightRest**](RightRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRights**
> ListRightRest ListRights(ctx, )
Get a list of rights in the system. Rights are predefined and cannot be changed. It requires System View right.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListRightRest**](ListRightRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionsForList1**
> OptionsRest OptionsForList1(ctx, )
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

