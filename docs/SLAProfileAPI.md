# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSlps**](SLAProfileApi.md#CountSlps) | **Head** /slp | Get a count of total slps matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.
[**CreateSlp**](SLAProfileApi.md#CreateSlp) | **Post** /slp | Create a new slp. It requires SLA Manage right.
[**DeleteSlp**](SLAProfileApi.md#DeleteSlp) | **Delete** /slp/{slp_id} | Remove the specific slp. It requires SLA Manage right.
[**GetSlp**](SLAProfileApi.md#GetSlp) | **Get** /slp/{slp_id} | Get individual slp details. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListSlps**](SLAProfileApi.md#ListSlps) | **Get** /slp | Get a list of slps. It requires SLA View, SLA Assign or SLA Manage rights.
[**UpdateSlp**](SLAProfileApi.md#UpdateSlp) | **Put** /slp/{slp_id} | Update the specific slp. It requires SLA Manage right.

# **CountSlps**
> CountSlps(ctx, optional)
Get a count of total slps matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLAProfileApiCountSlpsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAProfileApiCountSlpsOpts struct
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

# **CreateSlp**
> SlpRest CreateSlp(ctx, optional)
Create a new slp. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLAProfileApiCreateSlpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAProfileApiCreateSlpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SlpRest**](SlpRest.md)|  | 

### Return type

[**SlpRest**](SlpRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSlp**
> DeleteSlp(ctx, slpId)
Remove the specific slp. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slpId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSlp**
> SlpRest GetSlp(ctx, slpId)
Get individual slp details. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slpId** | **string**|  | 

### Return type

[**SlpRest**](SlpRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSlps**
> ListSlpRest ListSlps(ctx, optional)
Get a list of slps. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLAProfileApiListSlpsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAProfileApiListSlpsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListSlpRest**](ListSlpRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSlp**
> SlpRest UpdateSlp(ctx, slpId, optional)
Update the specific slp. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **slpId** | **string**|  | 
 **optional** | ***SLAProfileApiUpdateSlpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLAProfileApiUpdateSlpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SlpRest**](SlpRest.md)|  | 

### Return type

[**SlpRest**](SlpRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

