# \SLATemplateAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneTemplates**](SLATemplateAPI.md#CloneTemplates) | **Post** /slt/{slt_id}/clone | Clone a specific slt. It requires SLA Manage right.
[**CountSlts**](SLATemplateAPI.md#CountSlts) | **Head** /slt | Get a count of total slts matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.
[**CreateOptionForPolicy**](SLATemplateAPI.md#CreateOptionForPolicy) | **Post** /slt/{slt_id}/policy/{policy_id}/settableoption | Create a new settable option for the specific policy. It requires SLA Assign or SLA Manage rights.
[**CreatePolicy**](SLATemplateAPI.md#CreatePolicy) | **Post** /slt/{slt_id}/policy | Create a new policy. It requires SLA Manage right.
[**CreateSlt**](SLATemplateAPI.md#CreateSlt) | **Post** /slt | Create a new slt. It requires SLA Manage right.
[**DeleteOptionForPolicy**](SLATemplateAPI.md#DeleteOptionForPolicy) | **Delete** /slt/{slt_id}/policy/{policy_id}/settableoption/{option_id} | Remove a settable option for the specific policy. It requires SLA Assign or SLA Manage rights.
[**DeletePolicy**](SLATemplateAPI.md#DeletePolicy) | **Delete** /slt/{slt_id}/policy/{policy_id} | Remove a policy. It requires SLA Manage right.
[**DeleteSlt**](SLATemplateAPI.md#DeleteSlt) | **Delete** /slt/{slt_id} | Remove a slt. It requires SLA Manage right.
[**GetOptionForPolicy**](SLATemplateAPI.md#GetOptionForPolicy) | **Get** /slt/{slt_id}/policy/{policy_id}/settableoption/{option_id} | Get a specific settable option of the specific policy. It requires SLA View, SLA Assign or SLA Manage rights.
[**GetPolicy**](SLATemplateAPI.md#GetPolicy) | **Get** /slt/{slt_id}/policy/{policy_id} | Get individual policy. It requires SLA View, SLA Assign or SLA Manage rights.
[**GetSlt**](SLATemplateAPI.md#GetSlt) | **Get** /slt/{slt_id} | Get individual slt details. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListOptionForPolicy**](SLATemplateAPI.md#ListOptionForPolicy) | **Get** /slt/{slt_id}/policy/{policy_id}/settableoption | List all existing settable options of the specific policy. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListPolicies**](SLATemplateAPI.md#ListPolicies) | **Get** /slt/{slt_id}/policy | Get policy list from the specific slt. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListSlts**](SLATemplateAPI.md#ListSlts) | **Get** /slt | Get a list of slts. It requires SLA View, SLA Assign or SLA Manage rights.
[**OptionsForList15**](SLATemplateAPI.md#OptionsForList15) | **Options** /slt | Describes the fields available for filtering and sorting
[**SettableOptionMetadataForPolicy**](SLATemplateAPI.md#SettableOptionMetadataForPolicy) | **Options** /slt/{slt_id}/policy/{policy_id}/settableoption | Get settable option metadata for the specific policy.
[**SettableOptionMetadataForPolicyType1**](SLATemplateAPI.md#SettableOptionMetadataForPolicyType1) | **Options** /slt/settableoption/{policytype} | Get settable option metadata for the specific policy type.
[**UpdateOptionForPolicy**](SLATemplateAPI.md#UpdateOptionForPolicy) | **Put** /slt/{slt_id}/policy/{policy_id}/settableoption/{option_id} | Update a settable option for the specific policy. It requires SLA Assgin or SLA Manage rights.
[**UpdatePolicy**](SLATemplateAPI.md#UpdatePolicy) | **Put** /slt/{slt_id}/policy/{policy_id} | Update a policy. It requires SLA Manage right.
[**UpdateSlt**](SLATemplateAPI.md#UpdateSlt) | **Put** /slt/{slt_id} | Update a slt. It requires SLA Manage right.



## CloneTemplates

> SltRest CloneTemplates(ctx, sltId).SltRest(sltRest).Execute()

Clone a specific slt. It requires SLA Manage right.

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
	sltId := int64(789) // int64 | 
	sltRest := *openapiclient.NewSltRest() // SltRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.CloneTemplates(context.Background(), sltId).SltRest(sltRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.CloneTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneTemplates`: SltRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.CloneTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sltRest** | [**SltRest**](SltRest.md) |  | 

### Return type

[**SltRest**](SltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountSlts

> CountSlts(ctx).Filter(filter).Execute()

Get a count of total slts matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.

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
	filter := "filter_example" // string | Filter field. Use OPTIONS method to get possible filter fields.<br>Then append an operator and value. Operators always begin with a colon and include:<br><table><tr><th>Operator</th><th>Meaning</th></tr><tr><td>:==</td><td>equals</td></tr><tr><td>:=|</td><td>contains (case-insensitive)</td></tr><tr><td>:>=</td><td>greater than or equal to</td></tr><tr><td>:<=</td><td>less than or equal to</td></tr><tr><td>:=b</td><td>bitwise and</td></tr></table> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SLATemplateAPI.CountSlts(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.CountSlts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountSltsRequest struct via the builder pattern


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


## CreateOptionForPolicy

> AdvancedOptionRest CreateOptionForPolicy(ctx, sltId, policyId).AdvancedOptionRest(advancedOptionRest).Execute()

Create a new settable option for the specific policy. It requires SLA Assign or SLA Manage rights.

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
	sltId := "sltId_example" // string | 
	policyId := "policyId_example" // string | 
	advancedOptionRest := *openapiclient.NewAdvancedOptionRest() // AdvancedOptionRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.CreateOptionForPolicy(context.Background(), sltId, policyId).AdvancedOptionRest(advancedOptionRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.CreateOptionForPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOptionForPolicy`: AdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.CreateOptionForPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOptionForPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **advancedOptionRest** | [**AdvancedOptionRest**](AdvancedOptionRest.md) |  | 

### Return type

[**AdvancedOptionRest**](AdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicy

> PolicyRest CreatePolicy(ctx, sltId).PolicyRest(policyRest).Execute()

Create a new policy. It requires SLA Manage right.

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
	sltId := "sltId_example" // string | 
	policyRest := *openapiclient.NewPolicyRest() // PolicyRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.CreatePolicy(context.Background(), sltId).PolicyRest(policyRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.CreatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicy`: PolicyRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyRest** | [**PolicyRest**](PolicyRest.md) |  | 

### Return type

[**PolicyRest**](PolicyRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSlt

> SltRest CreateSlt(ctx).SltRest(sltRest).Execute()

Create a new slt. It requires SLA Manage right.

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
	sltRest := *openapiclient.NewSltRest() // SltRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.CreateSlt(context.Background()).SltRest(sltRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.CreateSlt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSlt`: SltRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.CreateSlt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSltRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sltRest** | [**SltRest**](SltRest.md) |  | 

### Return type

[**SltRest**](SltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOptionForPolicy

> DeleteOptionForPolicy(ctx, sltId, policyId, optionId).Execute()

Remove a settable option for the specific policy. It requires SLA Assign or SLA Manage rights.

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
	sltId := "sltId_example" // string | 
	policyId := "policyId_example" // string | 
	optionId := "optionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SLATemplateAPI.DeleteOptionForPolicy(context.Background(), sltId, policyId, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.DeleteOptionForPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 
**policyId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOptionForPolicyRequest struct via the builder pattern


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


## DeletePolicy

> DeletePolicy(ctx, sltId, policyId).Execute()

Remove a policy. It requires SLA Manage right.

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
	sltId := "sltId_example" // string | 
	policyId := "policyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SLATemplateAPI.DeletePolicy(context.Background(), sltId, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.DeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


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


## DeleteSlt

> DeleteSlt(ctx, sltId).Execute()

Remove a slt. It requires SLA Manage right.

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
	sltId := "sltId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SLATemplateAPI.DeleteSlt(context.Background(), sltId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.DeleteSlt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSltRequest struct via the builder pattern


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


## GetOptionForPolicy

> AdvancedOptionRest GetOptionForPolicy(ctx, sltId, policyId, optionId).Execute()

Get a specific settable option of the specific policy. It requires SLA View, SLA Assign or SLA Manage rights.

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
	sltId := "sltId_example" // string | 
	policyId := "policyId_example" // string | 
	optionId := "optionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.GetOptionForPolicy(context.Background(), sltId, policyId, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.GetOptionForPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptionForPolicy`: AdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.GetOptionForPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 
**policyId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionForPolicyRequest struct via the builder pattern


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


## GetPolicy

> PolicyRest GetPolicy(ctx, sltId, policyId).Execute()

Get individual policy. It requires SLA View, SLA Assign or SLA Manage rights.

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
	sltId := "sltId_example" // string | 
	policyId := "policyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.GetPolicy(context.Background(), sltId, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.GetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicy`: PolicyRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PolicyRest**](PolicyRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlt

> SltRest GetSlt(ctx, sltId).Execute()

Get individual slt details. It requires SLA View, SLA Assign or SLA Manage rights.

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
	sltId := "sltId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.GetSlt(context.Background(), sltId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.GetSlt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlt`: SltRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.GetSlt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSltRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SltRest**](SltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptionForPolicy

> ListAdvancedOptionRest ListOptionForPolicy(ctx, sltId, policyId).Execute()

List all existing settable options of the specific policy. It requires SLA View, SLA Assign or SLA Manage rights.

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
	sltId := "sltId_example" // string | 
	policyId := "policyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.ListOptionForPolicy(context.Background(), sltId, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.ListOptionForPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOptionForPolicy`: ListAdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.ListOptionForPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOptionForPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListAdvancedOptionRest**](ListAdvancedOptionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> ListPolicyRest ListPolicies(ctx, sltId).Execute()

Get policy list from the specific slt. It requires SLA View, SLA Assign or SLA Manage rights.

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
	sltId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.ListPolicies(context.Background(), sltId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.ListPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicies`: ListPolicyRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.ListPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPolicyRest**](ListPolicyRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlts

> ListSltRest ListSlts(ctx).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()

Get a list of slts. It requires SLA View, SLA Assign or SLA Manage rights.

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
	sort := "sort_example" // string | Sort field. Use OPTIONS method to get possible sort fields.<br>Then append ':asc' or ':desc' for ascending or descending sort.<br>Sorting is case-sensitive. (optional)
	filter := "filter_example" // string | Filter field. Use OPTIONS method to get possible filter fields.<br>Then append an operator and value. Operators always begin with a colon and include:<br><table><tr><th>Operator</th><th>Meaning</th></tr><tr><td>:==</td><td>equals</td></tr><tr><td>:=|</td><td>contains (case-insensitive)</td></tr><tr><td>:>=</td><td>greater than or equal to</td></tr><tr><td>:<=</td><td>less than or equal to</td></tr><tr><td>:=b</td><td>bitwise and</td></tr></table> (optional)
	limit := int64(789) // int64 | Limit on the number of results to return (optional)
	offset := int64(789) // int64 | Used with limit to support pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.ListSlts(context.Background()).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.ListSlts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSlts`: ListSltRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.ListSlts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSltsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#39;:asc&#39; or &#39;:desc&#39; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **int64** | Limit on the number of results to return | 
 **offset** | **int64** | Used with limit to support pagination | 

### Return type

[**ListSltRest**](ListSltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsForList15

> OptionsRest OptionsForList15(ctx).Execute()

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
	resp, r, err := apiClient.SLATemplateAPI.OptionsForList15(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.OptionsForList15``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsForList15`: OptionsRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.OptionsForList15`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsForList15Request struct via the builder pattern


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


## SettableOptionMetadataForPolicy

> SettableOptionMetadataForPolicy(ctx, sltId, policyId).Execute()

Get settable option metadata for the specific policy.

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
	sltId := "sltId_example" // string | 
	policyId := "policyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SLATemplateAPI.SettableOptionMetadataForPolicy(context.Background(), sltId, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.SettableOptionMetadataForPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettableOptionMetadataForPolicyRequest struct via the builder pattern


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


## SettableOptionMetadataForPolicyType1

> SettableOptionMetadataForPolicyType1(ctx, policytype).Apptype(apptype).Execute()

Get settable option metadata for the specific policy type.

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
	policytype := "policytype_example" // string | 
	apptype := "apptype_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SLATemplateAPI.SettableOptionMetadataForPolicyType1(context.Background(), policytype).Apptype(apptype).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.SettableOptionMetadataForPolicyType1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policytype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettableOptionMetadataForPolicyType1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apptype** | **string** |  | 

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


## UpdateOptionForPolicy

> AdvancedOptionRest UpdateOptionForPolicy(ctx, sltId, policyId, optionId).AdvancedOptionRest(advancedOptionRest).Execute()

Update a settable option for the specific policy. It requires SLA Assgin or SLA Manage rights.

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
	sltId := "sltId_example" // string | 
	policyId := "policyId_example" // string | 
	optionId := "optionId_example" // string | 
	advancedOptionRest := *openapiclient.NewAdvancedOptionRest() // AdvancedOptionRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.UpdateOptionForPolicy(context.Background(), sltId, policyId, optionId).AdvancedOptionRest(advancedOptionRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.UpdateOptionForPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOptionForPolicy`: AdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.UpdateOptionForPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 
**policyId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOptionForPolicyRequest struct via the builder pattern


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


## UpdatePolicy

> PolicyRest UpdatePolicy(ctx, sltId, policyId).PolicyRest(policyRest).Execute()

Update a policy. It requires SLA Manage right.

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
	sltId := "sltId_example" // string | 
	policyId := "policyId_example" // string | 
	policyRest := *openapiclient.NewPolicyRest() // PolicyRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.UpdatePolicy(context.Background(), sltId, policyId).PolicyRest(policyRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.UpdatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicy`: PolicyRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyRest** | [**PolicyRest**](PolicyRest.md) |  | 

### Return type

[**PolicyRest**](PolicyRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlt

> SltRest UpdateSlt(ctx, sltId).SltRest(sltRest).Execute()

Update a slt. It requires SLA Manage right.

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
	sltId := "sltId_example" // string | 
	sltRest := *openapiclient.NewSltRest() // SltRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLATemplateAPI.UpdateSlt(context.Background(), sltId).SltRest(sltRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLATemplateAPI.UpdateSlt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSlt`: SltRest
	fmt.Fprintf(os.Stdout, "Response from `SLATemplateAPI.UpdateSlt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sltId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSltRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sltRest** | [**SltRest**](SltRest.md) |  | 

### Return type

[**SltRest**](SltRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

