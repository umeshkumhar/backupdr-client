# \SLAAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSlas**](SLAAPI.md#CountSlas) | **Head** /sla | Get a count of total SLAs matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.
[**CreateOptionForSla**](SLAAPI.md#CreateOptionForSla) | **Post** /sla/{sla_id}/settableoption | Create an overridden option on the specific SLA. It requires SLA Assign or SLA Manage rights.
[**CreateSla**](SLAAPI.md#CreateSla) | **Post** /sla | Create a new SLA. It requires SLA Assign right.
[**DeleteOptionForSla**](SLAAPI.md#DeleteOptionForSla) | **Delete** /sla/{sla_id}/settableoption/{option_id} | Delete a specific overridden option on the specific SLA. It requires SLA Assign or SLA Manage rights.
[**DeleteSla**](SLAAPI.md#DeleteSla) | **Delete** /sla/{sla_id} | Remove the specific SLA. It requires SLA Assign right.
[**GetOptionForSla**](SLAAPI.md#GetOptionForSla) | **Get** /sla/{sla_id}/settableoption/{option_id} | Get the details of a specific overridden option on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.
[**GetSla**](SLAAPI.md#GetSla) | **Get** /sla/{sla_id} | Get individual SLA details. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListOptionForSla**](SLAAPI.md#ListOptionForSla) | **Get** /sla/{sla_id}/settableoption | List all overridden options already set on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.
[**ListSlas**](SLAAPI.md#ListSlas) | **Get** /sla | Get a list of SLAs. It requires SLA View, SLA Assign or SLA Manage rights.
[**SettableOptionMetadataForSla**](SLAAPI.md#SettableOptionMetadataForSla) | **Options** /sla/{sla_id}/settableoption | List all overridable option metadata for the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.
[**UpdateOptionForSla**](SLAAPI.md#UpdateOptionForSla) | **Put** /sla/{sla_id}/settableoption/{option_id} | Update a specific overridden option on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.
[**UpdateSla**](SLAAPI.md#UpdateSla) | **Put** /sla/{sla_id} | Update the specific slp. It requires SLA Assign right.



## CountSlas

> CountSlas(ctx).Filter(filter).Execute()

Get a count of total SLAs matching the filters. It requires SLA View, SLA Assign or SLA Manage rights.

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
	r, err := apiClient.SLAAPI.CountSlas(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.CountSlas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountSlasRequest struct via the builder pattern


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


## CreateOptionForSla

> AdvancedOptionRest CreateOptionForSla(ctx, slaId).AdvancedOptionRest(advancedOptionRest).Execute()

Create an overridden option on the specific SLA. It requires SLA Assign or SLA Manage rights.

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
	slaId := "slaId_example" // string | 
	advancedOptionRest := *openapiclient.NewAdvancedOptionRest() // AdvancedOptionRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAAPI.CreateOptionForSla(context.Background(), slaId).AdvancedOptionRest(advancedOptionRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.CreateOptionForSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOptionForSla`: AdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `SLAAPI.CreateOptionForSla`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOptionForSlaRequest struct via the builder pattern


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


## CreateSla

> SlaRest CreateSla(ctx).SlaRest(slaRest).Execute()

Create a new SLA. It requires SLA Assign right.

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
	slaRest := *openapiclient.NewSlaRest() // SlaRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAAPI.CreateSla(context.Background()).SlaRest(slaRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.CreateSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSla`: SlaRest
	fmt.Fprintf(os.Stdout, "Response from `SLAAPI.CreateSla`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slaRest** | [**SlaRest**](SlaRest.md) |  | 

### Return type

[**SlaRest**](SlaRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOptionForSla

> DeleteOptionForSla(ctx, slaId, optionId).Execute()

Delete a specific overridden option on the specific SLA. It requires SLA Assign or SLA Manage rights.

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
	slaId := "slaId_example" // string | 
	optionId := "optionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SLAAPI.DeleteOptionForSla(context.Background(), slaId, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.DeleteOptionForSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOptionForSlaRequest struct via the builder pattern


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


## DeleteSla

> DeleteSla(ctx, slaId).Execute()

Remove the specific SLA. It requires SLA Assign right.

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
	slaId := "slaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SLAAPI.DeleteSla(context.Background(), slaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.DeleteSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSlaRequest struct via the builder pattern


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


## GetOptionForSla

> AdvancedOptionRest GetOptionForSla(ctx, slaId, optionId).Execute()

Get the details of a specific overridden option on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.

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
	slaId := "slaId_example" // string | 
	optionId := "optionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAAPI.GetOptionForSla(context.Background(), slaId, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.GetOptionForSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptionForSla`: AdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `SLAAPI.GetOptionForSla`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionForSlaRequest struct via the builder pattern


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


## GetSla

> SlaRest GetSla(ctx, slaId).Execute()

Get individual SLA details. It requires SLA View, SLA Assign or SLA Manage rights.

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
	slaId := "slaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAAPI.GetSla(context.Background(), slaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.GetSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSla`: SlaRest
	fmt.Fprintf(os.Stdout, "Response from `SLAAPI.GetSla`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlaRest**](SlaRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptionForSla

> ListAdvancedOptionRest ListOptionForSla(ctx, slaId).Execute()

List all overridden options already set on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.

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
	slaId := "slaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAAPI.ListOptionForSla(context.Background(), slaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.ListOptionForSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOptionForSla`: ListAdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `SLAAPI.ListOptionForSla`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOptionForSlaRequest struct via the builder pattern


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


## ListSlas

> ListSlaRest ListSlas(ctx).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()

Get a list of SLAs. It requires SLA View, SLA Assign or SLA Manage rights.

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
	resp, r, err := apiClient.SLAAPI.ListSlas(context.Background()).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.ListSlas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSlas`: ListSlaRest
	fmt.Fprintf(os.Stdout, "Response from `SLAAPI.ListSlas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSlasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#39;:asc&#39; or &#39;:desc&#39; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **int64** | Limit on the number of results to return | 
 **offset** | **int64** | Used with limit to support pagination | 

### Return type

[**ListSlaRest**](ListSlaRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettableOptionMetadataForSla

> JSONArray SettableOptionMetadataForSla(ctx, slaId).Execute()

List all overridable option metadata for the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.

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
	slaId := "slaId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAAPI.SettableOptionMetadataForSla(context.Background(), slaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.SettableOptionMetadataForSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettableOptionMetadataForSla`: JSONArray
	fmt.Fprintf(os.Stdout, "Response from `SLAAPI.SettableOptionMetadataForSla`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSettableOptionMetadataForSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSONArray**](JSONArray.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOptionForSla

> AdvancedOptionRest UpdateOptionForSla(ctx, slaId, optionId).AdvancedOptionRest(advancedOptionRest).Execute()

Update a specific overridden option on the specific SLA. It requires SLA View, SLA Assign or SLA Manage rights.

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
	slaId := "slaId_example" // string | 
	optionId := "optionId_example" // string | 
	advancedOptionRest := *openapiclient.NewAdvancedOptionRest() // AdvancedOptionRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAAPI.UpdateOptionForSla(context.Background(), slaId, optionId).AdvancedOptionRest(advancedOptionRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.UpdateOptionForSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOptionForSla`: AdvancedOptionRest
	fmt.Fprintf(os.Stdout, "Response from `SLAAPI.UpdateOptionForSla`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOptionForSlaRequest struct via the builder pattern


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


## UpdateSla

> SlaRest UpdateSla(ctx, slaId).SlaRest(slaRest).Execute()

Update the specific slp. It requires SLA Assign right.

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
	slaId := "slaId_example" // string | 
	slaRest := *openapiclient.NewSlaRest() // SlaRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAAPI.UpdateSla(context.Background(), slaId).SlaRest(slaRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAAPI.UpdateSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSla`: SlaRest
	fmt.Fprintf(os.Stdout, "Response from `SLAAPI.UpdateSla`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slaRest** | [**SlaRest**](SlaRest.md) |  | 

### Return type

[**SlaRest**](SlaRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

