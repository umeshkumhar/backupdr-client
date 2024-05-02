# {{classname}}

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVersion**](ConfigurationApi.md#GetVersion) | **Get** /config/version | Get version. No special IAM permission requirement as long as there is valid session id.
[**GetVersionDetail**](ConfigurationApi.md#GetVersionDetail) | **Get** /config/versiondetail | Get version details. No special IAM permission requirement as long as there is valid session id.
[**GetVmMetadataDetails**](ConfigurationApi.md#GetVmMetadataDetails) | **Get** /config/mgmtconsoledetails | 

# **GetVersion**
> VersionRest GetVersion(ctx, )
Get version. No special IAM permission requirement as long as there is valid session id.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VersionRest**](VersionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersionDetail**
> VersionDetailRest GetVersionDetail(ctx, )
Get version details. No special IAM permission requirement as long as there is valid session id.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VersionDetailRest**](VersionDetailRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVmMetadataDetails**
> GetVmMetadataDetails(ctx, )


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

