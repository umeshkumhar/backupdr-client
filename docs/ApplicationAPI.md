# \ApplicationAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupNow**](ApplicationAPI.md#BackupNow) | **Post** /application/{application_id}/backup | Run a backup job. The user must have Application Manage, Host Manage, or Backup Manage right.
[**CountApplications**](ApplicationAPI.md#CountApplications) | **Head** /application | Get a count of total applications matching the filters.
[**CreateWorkflow**](ApplicationAPI.md#CreateWorkflow) | **Post** /application/{application_id}/workflow | Create new workflow for the particular application. It requires Workflow Manage right.
[**DeleteApplication**](ApplicationAPI.md#DeleteApplication) | **Delete** /application/{application_id} | Delete application. It requires Application Manage or Host Manage rights.
[**DeleteOptionForApp**](ApplicationAPI.md#DeleteOptionForApp) | **Delete** /application/{application_id}/settableoption/{option_id} | Delete the particular option from the particular application. It requires SLA Manage or SLA Assign rights.
[**DeleteWorkflow**](ApplicationAPI.md#DeleteWorkflow) | **Delete** /application/{application_id}/workflow/{workflow_id} | Delete existing workflow. It requires Workflow Manage right.
[**ExpireBackups**](ApplicationAPI.md#ExpireBackups) | **Post** /application/{application_id}/expirebackup | Expires backups of the particular application. It requires the manageExpiration permission
[**GetAppClass**](ApplicationAPI.md#GetAppClass) | **Get** /application/{application_id}/appclass | Get the particular application&#39;s appclass metadata
[**GetAppClassByAppclassName**](ApplicationAPI.md#GetAppClassByAppclassName) | **Get** /application/appclass/{appclass_name} | Get appclass metadata for the particular appclass name
[**GetAppClasses**](ApplicationAPI.md#GetAppClasses) | **Get** /application/appclass | Get all available app classes from given cluster
[**GetApplication**](ApplicationAPI.md#GetApplication) | **Get** /application/{application_id} | Get details on the particular application
[**GetOptionForApp**](ApplicationAPI.md#GetOptionForApp) | **Get** /application/{application_id}/settableoption/{option_id} | Get the particular option of the particular application. It requires SLA View, SLA Manage, or SLA Assign rights.
[**GetWorkflow**](ApplicationAPI.md#GetWorkflow) | **Get** /application/{application_id}/workflow/{workflow_id} | Get individual workflow. It requires Workflow View right.
[**ListActiveImages**](ApplicationAPI.md#ListActiveImages) | **Get** /application/{application_id}/activeimage | Get active images for an application
[**ListApplicationTypes**](ApplicationAPI.md#ListApplicationTypes) | **Get** /application/types | Get list of application types that are currently in the system.
[**ListApplications**](ApplicationAPI.md#ListApplications) | **Get** /application | List applications.
[**ListWorkflows**](ApplicationAPI.md#ListWorkflows) | **Get** /application/{application_id}/workflow | Get list of workflows for the particular application. It requires Workflow View right.
[**OptionsForList2**](ApplicationAPI.md#OptionsForList2) | **Options** /application | Describes the fields available for filtering and sorting
[**SettableOptionMetadataForPolicyType**](ApplicationAPI.md#SettableOptionMetadataForPolicyType) | **Options** /application/settableoption/{apptype} | Settable option metadata for the particular application type
[**UpdateApplication**](ApplicationAPI.md#UpdateApplication) | **Put** /application/{application_id} | Update application data. It requires Application Manage or Host Manage rights.
[**UpdateOptionForApp**](ApplicationAPI.md#UpdateOptionForApp) | **Put** /application/{application_id}/settableoption/{option_id} | Update the particular option of the particular application. It requires SLA Manage or SLA Assign rights.
[**UpdateWorkflow**](ApplicationAPI.md#UpdateWorkflow) | **Put** /application/{application_id}/workflow/{workflow_id} | Update existing workflow. It requires Workflow Manage right.



## BackupNow

> BackupNow(ctx, applicationId).BackupNowRest(backupNowRest).Execute()

Run a backup job. The user must have Application Manage, Host Manage, or Backup Manage right.

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
	applicationId := "applicationId_example" // string | 
	backupNowRest := *openapiclient.NewBackupNowRest() // BackupNowRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationAPI.BackupNow(context.Background(), applicationId).BackupNowRest(backupNowRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.BackupNow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupNowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupNowRest** | [**BackupNowRest**](BackupNowRest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountApplications

> CountApplications(ctx).Filter(filter).Execute()

Get a count of total applications matching the filters.

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
	filter := "hostname:=|myhost" // string | Filter field. Use OPTIONS method to get possible filter fields.<br>Then append an operator and value. Operators always begin with a colon and include:<br><table><tr><th>Operator</th><th>Meaning</th></tr><tr><td>:==</td><td>equals</td></tr><tr><td>:=|</td><td>contains (case-insensitive)</td></tr><tr><td>:>=</td><td>greater than or equal to</td></tr><tr><td>:<=</td><td>less than or equal to</td></tr><tr><td>:=b</td><td>bitwise and</td></tr></table> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationAPI.CountApplications(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.CountApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflow

> WorkflowRest CreateWorkflow(ctx, applicationId).WorkflowRest(workflowRest).Execute()

Create new workflow for the particular application. It requires Workflow Manage right.

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
	applicationId := "applicationId_example" // string | 
	workflowRest := *openapiclient.NewWorkflowRest() // WorkflowRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.CreateWorkflow(context.Background(), applicationId).WorkflowRest(workflowRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.CreateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflow`: WorkflowRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.CreateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowRest** | [**WorkflowRest**](WorkflowRest.md) |  | 

### Return type

[**WorkflowRest**](WorkflowRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, applicationId).Execute()

Delete application. It requires Application Manage or Host Manage rights.

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
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationAPI.DeleteApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.DeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOptionForApp

> DeleteOptionForApp(ctx, applicationId, optionId).Execute()

Delete the particular option from the particular application. It requires SLA Manage or SLA Assign rights.

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
	applicationId := "applicationId_example" // string | 
	optionId := "optionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationAPI.DeleteOptionForApp(context.Background(), applicationId, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.DeleteOptionForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOptionForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflow

> DeleteWorkflow(ctx, applicationId, workflowId).Execute()

Delete existing workflow. It requires Workflow Manage right.

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
	applicationId := "applicationId_example" // string | 
	workflowId := "workflowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationAPI.DeleteWorkflow(context.Background(), applicationId, workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.DeleteWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireBackups

> ExpireBackups(ctx, applicationId).ExpireBackupRest(expireBackupRest).Execute()

Expires backups of the particular application. It requires the manageExpiration permission



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
	applicationId := "applicationId_example" // string | 
	expireBackupRest := *openapiclient.NewExpireBackupRest() // ExpireBackupRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationAPI.ExpireBackups(context.Background(), applicationId).ExpireBackupRest(expireBackupRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.ExpireBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expireBackupRest** | [**ExpireBackupRest**](ExpireBackupRest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppClass

> AppClassRest GetAppClass(ctx, applicationId).Hostid(hostid).Isasm(isasm).Operation(operation).Backupid(backupid).Execute()

Get the particular application's appclass metadata

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
	applicationId := "applicationId_example" // string | 
	hostid := "hostid_example" // string |  (optional)
	isasm := "isasm_example" // string |  (optional)
	operation := "operation_example" // string |  (optional)
	backupid := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetAppClass(context.Background(), applicationId).Hostid(hostid).Isasm(isasm).Operation(operation).Backupid(backupid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetAppClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppClass`: AppClassRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetAppClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostid** | **string** |  | 
 **isasm** | **string** |  | 
 **operation** | **string** |  | 
 **backupid** | **int64** |  | 

### Return type

[**AppClassRest**](AppClassRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppClassByAppclassName

> AppClassRest GetAppClassByAppclassName(ctx, appclassName).Hostid(hostid).Clusterid(clusterid).Isasm(isasm).Operation(operation).Backupid(backupid).Execute()

Get appclass metadata for the particular appclass name

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
	appclassName := "appclassName_example" // string | 
	hostid := "hostid_example" // string |  (optional)
	clusterid := "clusterid_example" // string |  (optional)
	isasm := "isasm_example" // string |  (optional)
	operation := "operation_example" // string |  (optional)
	backupid := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetAppClassByAppclassName(context.Background(), appclassName).Hostid(hostid).Clusterid(clusterid).Isasm(isasm).Operation(operation).Backupid(backupid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetAppClassByAppclassName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppClassByAppclassName`: AppClassRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetAppClassByAppclassName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appclassName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppClassByAppclassNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostid** | **string** |  | 
 **clusterid** | **string** |  | 
 **isasm** | **string** |  | 
 **operation** | **string** |  | 
 **backupid** | **int64** |  | 

### Return type

[**AppClassRest**](AppClassRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppClasses

> AppClassRest GetAppClasses(ctx).Clusterid(clusterid).Execute()

Get all available app classes from given cluster

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
	clusterid := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetAppClasses(context.Background()).Clusterid(clusterid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetAppClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppClasses`: AppClassRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetAppClasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterid** | **int64** |  | 

### Return type

[**AppClassRest**](AppClassRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> ApplicationRest GetApplication(ctx, applicationId).Execute()

Get details on the particular application

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
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: ApplicationRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationRest**](ApplicationRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOptionForApp

> AdvancedOptionRest GetOptionForApp(ctx, applicationId, optionId).Execute()

Get the particular option of the particular application. It requires SLA View, SLA Manage, or SLA Assign rights.

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
	applicationId := "applicationId_example" // string | 
	optionId := "optionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetOptionForApp(context.Background(), applicationId, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetOptionForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptionForApp`: AdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetOptionForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AdvancedOptionRest**](AdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflow

> WorkflowRest GetWorkflow(ctx, applicationId, workflowId).Execute()

Get individual workflow. It requires Workflow View right.



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
	applicationId := "applicationId_example" // string | 
	workflowId := "workflowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetWorkflow(context.Background(), applicationId, workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflow`: WorkflowRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkflowRest**](WorkflowRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActiveImages

> ListBackupRest ListActiveImages(ctx, applicationId).Execute()

Get active images for an application

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
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.ListActiveImages(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.ListActiveImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActiveImages`: ListBackupRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.ListActiveImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActiveImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListBackupRest**](ListBackupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationTypes

> ListStringRest ListApplicationTypes(ctx).Execute()

Get list of application types that are currently in the system.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.ListApplicationTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.ListApplicationTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationTypes`: ListStringRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.ListApplicationTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationTypesRequest struct via the builder pattern


### Return type

[**ListStringRest**](ListStringRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> ListApplicationRest ListApplications(ctx).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()

List applications.

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
	sort := "appname:asc" // string | Sort field. Use OPTIONS method to get possible sort fields.<br>Then append ':asc' or ':desc' for ascending or descending sort.<br>Sorting is case-sensitive. (optional)
	filter := "hostname:=|myhost" // string | Filter field. Use OPTIONS method to get possible filter fields.<br>Then append an operator and value. Operators always begin with a colon and include:<br><table><tr><th>Operator</th><th>Meaning</th></tr><tr><td>:==</td><td>equals</td></tr><tr><td>:=|</td><td>contains (case-insensitive)</td></tr><tr><td>:>=</td><td>greater than or equal to</td></tr><tr><td>:<=</td><td>less than or equal to</td></tr><tr><td>:=b</td><td>bitwise and</td></tr></table> (optional)
	limit := int64(50) // int64 | Limit on the number of results to return (optional)
	offset := int64(0) // int64 | Used with limit to support pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.ListApplications(context.Background()).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.ListApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplications`: ListApplicationRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.ListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#39;:asc&#39; or &#39;:desc&#39; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **int64** | Limit on the number of results to return | 
 **offset** | **int64** | Used with limit to support pagination | 

### Return type

[**ListApplicationRest**](ListApplicationRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflows

> ListWorkflowRest ListWorkflows(ctx, applicationId).Execute()

Get list of workflows for the particular application. It requires Workflow View right.

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
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.ListWorkflows(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.ListWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkflows`: ListWorkflowRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.ListWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListWorkflowRest**](ListWorkflowRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsForList2

> OptionsRest OptionsForList2(ctx).Execute()

Describes the fields available for filtering and sorting

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.OptionsForList2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.OptionsForList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsForList2`: OptionsRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.OptionsForList2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsForList2Request struct via the builder pattern


### Return type

[**OptionsRest**](OptionsRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettableOptionMetadataForPolicyType

> string SettableOptionMetadataForPolicyType(ctx, apptype).Execute()

Settable option metadata for the particular application type

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
	apptype := "apptype_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.SettableOptionMetadataForPolicyType(context.Background(), apptype).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.SettableOptionMetadataForPolicyType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettableOptionMetadataForPolicyType`: string
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.SettableOptionMetadataForPolicyType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apptype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettableOptionMetadataForPolicyTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> UpdateApplication(ctx, applicationId).ApplicationRest(applicationRest).Execute()

Update application data. It requires Application Manage or Host Manage rights.

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
	applicationId := "applicationId_example" // string | 
	applicationRest := *openapiclient.NewApplicationRest() // ApplicationRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationAPI.UpdateApplication(context.Background(), applicationId).ApplicationRest(applicationRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.UpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationRest** | [**ApplicationRest**](ApplicationRest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOptionForApp

> AdvancedOptionRest UpdateOptionForApp(ctx, applicationId, optionId).AdvancedOptionRest(advancedOptionRest).Execute()

Update the particular option of the particular application. It requires SLA Manage or SLA Assign rights.

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
	applicationId := "applicationId_example" // string | 
	optionId := "optionId_example" // string | 
	advancedOptionRest := *openapiclient.NewAdvancedOptionRest() // AdvancedOptionRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.UpdateOptionForApp(context.Background(), applicationId, optionId).AdvancedOptionRest(advancedOptionRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.UpdateOptionForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOptionForApp`: AdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.UpdateOptionForApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOptionForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **advancedOptionRest** | [**AdvancedOptionRest**](AdvancedOptionRest.md) |  | 

### Return type

[**AdvancedOptionRest**](AdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflow

> WorkflowRest UpdateWorkflow(ctx, applicationId, workflowId).WorkflowRest(workflowRest).Execute()

Update existing workflow. It requires Workflow Manage right.

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
	applicationId := "applicationId_example" // string | 
	workflowId := "workflowId_example" // string | 
	workflowRest := *openapiclient.NewWorkflowRest() // WorkflowRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.UpdateWorkflow(context.Background(), applicationId, workflowId).WorkflowRest(workflowRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.UpdateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflow`: WorkflowRest
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.UpdateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workflowRest** | [**WorkflowRest**](WorkflowRest.md) |  | 

### Return type

[**WorkflowRest**](WorkflowRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

