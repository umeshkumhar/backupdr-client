# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneTemplates**](SLATemplateApi.md#CloneTemplates) | **Post** /slt/{slt_id}/clone | Clone a specific slt. It requires SLA Manage right.
[**CountSlts**](SLATemplateApi.md#CountSlts) | **Head** /slt | Get a count of total slts matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.
[**CreateOptionForPolicy**](SLATemplateApi.md#CreateOptionForPolicy) | **Post** /slt/{slt_id}/policy/{policy_id}/settableoption | Create a new settable option for the specific policy. It requires SLA Assign or SLA Manage rights.
[**CreatePolicy**](SLATemplateApi.md#CreatePolicy) | **Post** /slt/{slt_id}/policy | Create a new policy. It requires SLA Manage right.
[**CreateSlt**](SLATemplateApi.md#CreateSlt) | **Post** /slt | Create a new slt. It requires SLA Manage right.
[**DeleteOptionForPolicy**](SLATemplateApi.md#DeleteOptionForPolicy) | **Delete** /slt/{slt_id}/policy/{policy_id}/settableoption/{option_id} | Remove a settable option for the specific policy. It requires SLA Assign or SLA Manage rights.
[**DeletePolicy**](SLATemplateApi.md#DeletePolicy) | **Delete** /slt/{slt_id}/policy/{policy_id} | Remove a policy. It requires SLA Manage right.
[**DeleteSlt**](SLATemplateApi.md#DeleteSlt) | **Delete** /slt/{slt_id} | Remove a slt. It requires SLA Manage right.
[**GetOptionForPolicy**](SLATemplateApi.md#GetOptionForPolicy) | **Get** /slt/{slt_id}/policy/{policy_id}/settableoption/{option_id} | Get a specific settable option of the specific policy. It requires SLA View, SLA Assign or SLA Manage rights.
[**GetPolicy**](SLATemplateApi.md#GetPolicy) | **Get** /slt/{slt_id}/policy/{policy_id} | Get individual policy. It requires SLA View, SLA Assign or SLA Manage rights.
[**GetSlt**](SLATemplateApi.md#GetSlt) | **Get** /slt/{slt_id} | Get individual slt details. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListOptionForPolicy**](SLATemplateApi.md#ListOptionForPolicy) | **Get** /slt/{slt_id}/policy/{policy_id}/settableoption | List all existing settable options of the specific policy. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListPolicies**](SLATemplateApi.md#ListPolicies) | **Get** /slt/{slt_id}/policy | Get policy list from the specific slt. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListSlts**](SLATemplateApi.md#ListSlts) | **Get** /slt | Get a list of slts. It requires SLA View, SLA Assign or SLA Manage rights.
[**OptionsForList15**](SLATemplateApi.md#OptionsForList15) | **Options** /slt | Describes the fields available for filtering and sorting
[**SettableOptionMetadataForPolicy**](SLATemplateApi.md#SettableOptionMetadataForPolicy) | **Options** /slt/{slt_id}/policy/{policy_id}/settableoption | Get settable option metadata for the specific policy.
[**SettableOptionMetadataForPolicyType1**](SLATemplateApi.md#SettableOptionMetadataForPolicyType1) | **Options** /slt/settableoption/{policytype} | Get settable option metadata for the specific policy type.
[**UpdateOptionForPolicy**](SLATemplateApi.md#UpdateOptionForPolicy) | **Put** /slt/{slt_id}/policy/{policy_id}/settableoption/{option_id} | Update a settable option for the specific policy. It requires SLA Assgin or SLA Manage rights.
[**UpdatePolicy**](SLATemplateApi.md#UpdatePolicy) | **Put** /slt/{slt_id}/policy/{policy_id} | Update a policy. It requires SLA Manage right.
[**UpdateSlt**](SLATemplateApi.md#UpdateSlt) | **Put** /slt/{slt_id} | Update a slt. It requires SLA Manage right.

# **CloneTemplates**
> SltRest CloneTemplates(ctx, sltId, optional)
Clone a specific slt. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **int64**|  | 
 **optional** | ***SLATemplateApiCloneTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiCloneTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SltRest**](SltRest.md)|  | 

### Return type

[**SltRest**](SltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountSlts**
> CountSlts(ctx, optional)
Get a count of total slts matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLATemplateApiCountSltsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiCountSltsOpts struct
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

# **CreateOptionForPolicy**
> AdvancedOptionRest CreateOptionForPolicy(ctx, sltId, policyId, optional)
Create a new settable option for the specific policy. It requires SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
  **policyId** | **string**|  | 
 **optional** | ***SLATemplateApiCreateOptionForPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiCreateOptionForPolicyOpts struct
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

# **CreatePolicy**
> PolicyRest CreatePolicy(ctx, sltId, optional)
Create a new policy. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
 **optional** | ***SLATemplateApiCreatePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiCreatePolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PolicyRest**](PolicyRest.md)|  | 

### Return type

[**PolicyRest**](PolicyRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSlt**
> SltRest CreateSlt(ctx, optional)
Create a new slt. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLATemplateApiCreateSltOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiCreateSltOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SltRest**](SltRest.md)|  | 

### Return type

[**SltRest**](SltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOptionForPolicy**
> DeleteOptionForPolicy(ctx, sltId, policyId, optionId)
Remove a settable option for the specific policy. It requires SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
  **policyId** | **string**|  | 
  **optionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePolicy**
> DeletePolicy(ctx, sltId, policyId)
Remove a policy. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSlt**
> DeleteSlt(ctx, sltId)
Remove a slt. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOptionForPolicy**
> AdvancedOptionRest GetOptionForPolicy(ctx, sltId, policyId, optionId)
Get a specific settable option of the specific policy. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
  **policyId** | **string**|  | 
  **optionId** | **string**|  | 

### Return type

[**AdvancedOptionRest**](AdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicy**
> PolicyRest GetPolicy(ctx, sltId, policyId)
Get individual policy. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

[**PolicyRest**](PolicyRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSlt**
> SltRest GetSlt(ctx, sltId)
Get individual slt details. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 

### Return type

[**SltRest**](SltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOptionForPolicy**
> ListAdvancedOptionRest ListOptionForPolicy(ctx, sltId, policyId)
List all existing settable options of the specific policy. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

[**ListAdvancedOptionRest**](ListAdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPolicies**
> ListPolicyRest ListPolicies(ctx, sltId)
Get policy list from the specific slt. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **int64**|  | 

### Return type

[**ListPolicyRest**](ListPolicyRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSlts**
> ListSltRest ListSlts(ctx, optional)
Get a list of slts. It requires SLA View, SLA Assign or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SLATemplateApiListSltsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiListSltsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#x27;:asc&#x27; or &#x27;:desc&#x27; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **optional.String**| Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **optional.Int64**| Limit on the number of results to return | 
 **offset** | **optional.Int64**| Used with limit to support pagination | 

### Return type

[**ListSltRest**](ListSltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptionsForList15**
> OptionsRest OptionsForList15(ctx, )
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

# **SettableOptionMetadataForPolicy**
> SettableOptionMetadataForPolicy(ctx, sltId, policyId)
Get settable option metadata for the specific policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
  **policyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettableOptionMetadataForPolicyType1**
> SettableOptionMetadataForPolicyType1(ctx, policytype, optional)
Get settable option metadata for the specific policy type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policytype** | **string**|  | 
 **optional** | ***SLATemplateApiSettableOptionMetadataForPolicyType1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiSettableOptionMetadataForPolicyType1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apptype** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOptionForPolicy**
> AdvancedOptionRest UpdateOptionForPolicy(ctx, sltId, policyId, optionId, optional)
Update a settable option for the specific policy. It requires SLA Assgin or SLA Manage rights.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
  **policyId** | **string**|  | 
  **optionId** | **string**|  | 
 **optional** | ***SLATemplateApiUpdateOptionForPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiUpdateOptionForPolicyOpts struct
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

# **UpdatePolicy**
> PolicyRest UpdatePolicy(ctx, sltId, policyId, optional)
Update a policy. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
  **policyId** | **string**|  | 
 **optional** | ***SLATemplateApiUpdatePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiUpdatePolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PolicyRest**](PolicyRest.md)|  | 

### Return type

[**PolicyRest**](PolicyRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSlt**
> SltRest UpdateSlt(ctx, sltId, optional)
Update a slt. It requires SLA Manage right.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sltId** | **string**|  | 
 **optional** | ***SLATemplateApiUpdateSltOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SLATemplateApiUpdateSltOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SltRest**](SltRest.md)|  | 

### Return type

[**SltRest**](SltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

