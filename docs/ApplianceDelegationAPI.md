# \ApplianceDelegationAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DelegateGetCallDownloadLog**](ApplianceDelegationAPI.md#DelegateGetCallDownloadLog) | **Get** /appliancedelegation/{cluster_id}/config/download/log | Download logs from backup/recovery appliance
[**DownloadConnector**](ApplianceDelegationAPI.md#DownloadConnector) | **Get** /appliancedelegation/{cluster_id}/connectorbinary/{connectorname} | Download connectors from backup/recovery appliance
[**DownloadOssNotice**](ApplianceDelegationAPI.md#DownloadOssNotice) | **Get** /appliancedelegation/{cluster_id}/config/download/ossnotice | Download zip file containing licenses and notices for open-source components from backup/recovery appliance
[**UploadSoftwareUpgradeToAppliance**](ApplianceDelegationAPI.md#UploadSoftwareUpgradeToAppliance) | **Post** /appliancedelegation/{cluster_id}/cluster/uploadupdate | Upload software upgrade packages



## DelegateGetCallDownloadLog

> map[string]interface{} DelegateGetCallDownloadLog(ctx, clusterId).Execute()

Download logs from backup/recovery appliance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umeshkumhar/backupdr-client"
)

func main() {
	clusterId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceDelegationAPI.DelegateGetCallDownloadLog(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceDelegationAPI.DelegateGetCallDownloadLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelegateGetCallDownloadLog`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceDelegationAPI.DelegateGetCallDownloadLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelegateGetCallDownloadLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadConnector

> map[string]interface{} DownloadConnector(ctx, connectorname, clusterId).Execute()

Download connectors from backup/recovery appliance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umeshkumhar/backupdr-client"
)

func main() {
	connectorname := "connectorname_example" // string | 
	clusterId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceDelegationAPI.DownloadConnector(context.Background(), connectorname, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceDelegationAPI.DownloadConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadConnector`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceDelegationAPI.DownloadConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectorname** | **string** |  | 
**clusterId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadOssNotice

> map[string]interface{} DownloadOssNotice(ctx, clusterId).Execute()

Download zip file containing licenses and notices for open-source components from backup/recovery appliance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umeshkumhar/backupdr-client"
)

func main() {
	clusterId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceDelegationAPI.DownloadOssNotice(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceDelegationAPI.DownloadOssNotice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadOssNotice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceDelegationAPI.DownloadOssNotice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOssNoticeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSoftwareUpgradeToAppliance

> string UploadSoftwareUpgradeToAppliance(ctx, clusterId).File(file).Execute()

Upload software upgrade packages

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umeshkumhar/backupdr-client"
)

func main() {
	clusterId := int64(789) // int64 | 
	file := *openapiclient.NewFormDataContentDisposition() // FormDataContentDisposition |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceDelegationAPI.UploadSoftwareUpgradeToAppliance(context.Background(), clusterId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceDelegationAPI.UploadSoftwareUpgradeToAppliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadSoftwareUpgradeToAppliance`: string
	fmt.Fprintf(os.Stdout, "Response from `ApplianceDelegationAPI.UploadSoftwareUpgradeToAppliance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSoftwareUpgradeToApplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**FormDataContentDisposition**](FormDataContentDisposition.md) |  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

