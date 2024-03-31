# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSlas**](SLAApi.md#CountSlas) | **Head** /sla | Get a count of total SLAs matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.
[**CreateOptionForSla**](SLAApi.md#CreateOptionForSla) | **Post** /sla/{sla_id}/settableoption | Create an overridden option on the specific SLA. It requires SLA Assign or SLA Manage rights.
[**CreateSla**](SLAApi.md#CreateSla) | **Post** /sla | Create a new SLA. It requires SLA Assign right.
[**DeleteOptionForSla**](SLAApi.md#DeleteOptionForSla) | **Delete** /sla/{sla_id}/settableoption/{option_id} | Delete a specific overridden option on the specific SLA. It requires SLA Assign or SLA Manage rights.
[**DeleteSla**](SLAApi.md#DeleteSla) | **Delete** /sla/{sla_id} | Remove the specific SLA. It requires SLA Assign right.
[**GetOptionForSla**](SLAApi.md#GetOptionForSla) | **Get** /sla/{sla_id}/settableoption/{option_id} | Get the details of a specific overridden option on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.
[**GetSla**](SLAApi.md#GetSla) | **Get** /sla/{sla_id} | Get individual SLA details. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListOptionForSla**](SLAApi.md#ListOptionForSla) | **Get** /sla/{sla_id}/settableoption | List all overridden options already set on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListSlas**](SLAApi.md#ListSlas) | **Get** /sla | Get a list of SLAs. It requires SLA View, SLA Assign or SLA Manage rights.
[**SettableOptionMetadataForSla**](SLAApi.md#SettableOptionMetadataForSla) | **Options** /sla/{sla_id}/settableoption | List all overridable option metadata for the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.
[**UpdateOptionForSla**](SLAApi.md#UpdateOptionForSla) | **Put** /sla/{sla_id}/settableoption/{option_id} | Update a specific overridden option on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.
[**UpdateSla**](SLAApi.md#UpdateSla) | **Put** /sla/{sla_id} | Update the specific slp. It requires SLA Assign right.

# **CountSlas**
> CountSlas(ctx, optional)
Get a count of total SLAs matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLAApiCountSlasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiCountSlasOpts struct
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

# **CreateOptionForSla**
> AdvancedOptionRest CreateOptionForSla(ctx, slaId, optional)
Create an overridden option on the specific SLA. It requires SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaId** | **string**|  | 
 **optional** | ***SLAApiCreateOptionForSlaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiCreateOptionForSlaOpts struct
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

# **CreateSla**
> SlaRest CreateSla(ctx, optional)
Create a new SLA. It requires SLA Assign right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLAApiCreateSlaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiCreateSlaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SlaRest**](SlaRest.md)|  | 

### Return type

[**SlaRest**](SlaRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOptionForSla**
> DeleteOptionForSla(ctx, slaId, optionId)
Delete a specific overridden option on the specific SLA. It requires SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaId** | **string**|  | 
  **optionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSla**
> DeleteSla(ctx, slaId)
Remove the specific SLA. It requires SLA Assign right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOptionForSla**
> AdvancedOptionRest GetOptionForSla(ctx, slaId, optionId)
Get the details of a specific overridden option on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaId** | **string**|  | 
  **optionId** | **string**|  | 

### Return type

[**AdvancedOptionRest**](AdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSla**
> SlaRest GetSla(ctx, slaId)
Get individual SLA details. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaId** | **string**|  | 

### Return type

[**SlaRest**](SlaRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOptionForSla**
> ListAdvancedOptionRest ListOptionForSla(ctx, slaId)
List all overridden options already set on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaId** | **string**|  | 

### Return type

[**ListAdvancedOptionRest**](ListAdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSlas**
> ListSlaRest ListSlas(ctx, optional)
Get a list of SLAs. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLAApiListSlasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiListSlasOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListSlaRest**](ListSlaRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettableOptionMetadataForSla**
> JsonArray SettableOptionMetadataForSla(ctx, slaId)
List all overridable option metadata for the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaId** | **string**|  | 

### Return type

[**JsonArray**](JSONArray.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOptionForSla**
> AdvancedOptionRest UpdateOptionForSla(ctx, slaId, optionId, optional)
Update a specific overridden option on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaId** | **string**|  | 
  **optionId** | **string**|  | 
 **optional** | ***SLAApiUpdateOptionForSlaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiUpdateOptionForSlaOpts struct
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

# **UpdateSla**
> SlaRest UpdateSla(ctx, slaId, optional)
Update the specific slp. It requires SLA Assign right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slaId** | **string**|  | 
 **optional** | ***SLAApiUpdateSlaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAApiUpdateSlaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SlaRest**](SlaRest.md)|  | 

### Return type

[**SlaRest**](SlaRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

