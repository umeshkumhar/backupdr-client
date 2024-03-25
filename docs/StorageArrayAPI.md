# \StorageArrayAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountArrays**](StorageArrayAPI.md#CountArrays) | **Head** /array | Get a count of total storage arrays matching the filters.
[**CreateArray**](StorageArrayAPI.md#CreateArray) | **Post** /array | Create a storage array. It requires Storage Manage right.
[**CredentialTest**](StorageArrayAPI.md#CredentialTest) | **Post** /array/pretest | Credential test before creating a storage array. It requires Storage Manage right.
[**DeleteArray**](StorageArrayAPI.md#DeleteArray) | **Delete** /array/{array_id} | Remove the specific storage array from all appliances. It requires Storage Manage right.
[**GetArray**](StorageArrayAPI.md#GetArray) | **Get** /array/{array_id} | Get individual storage array details.
[**GetArrayTypes**](StorageArrayAPI.md#GetArrayTypes) | **Get** /array/type | Get all supported storage array types from appliances. It requires Storage View right.
[**ListArrays**](StorageArrayAPI.md#ListArrays) | **Get** /array | Get a list of storage arrays
[**OptionsForList3**](StorageArrayAPI.md#OptionsForList3) | **Options** /array | Describes the fields available for filtering and sorting
[**TestArray**](StorageArrayAPI.md#TestArray) | **Post** /array/{array_id} | Connectivity test on the specific storage array. It requires Storage Manage right.
[**UpdateArray**](StorageArrayAPI.md#UpdateArray) | **Put** /array/{array_id} | Update a specific storage array. It requires Storage Manage right.



## CountArrays

> CountArrays(ctx).Filter(filter).Execute()

Get a count of total storage arrays matching the filters.

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
	r, err := apiClient.StorageArrayAPI.CountArrays(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.CountArrays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountArraysRequest struct via the builder pattern


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


## CreateArray

> ArrayRest CreateArray(ctx).ArrayRest(arrayRest).Execute()

Create a storage array. It requires Storage Manage right.

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
	arrayRest := *openapiclient.NewArrayRest() // ArrayRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageArrayAPI.CreateArray(context.Background()).ArrayRest(arrayRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.CreateArray``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateArray`: ArrayRest
	fmt.Fprintf(os.Stdout, "Response from `StorageArrayAPI.CreateArray`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arrayRest** | [**ArrayRest**](ArrayRest.md) |  | 

### Return type

[**ArrayRest**](ArrayRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CredentialTest

> CredentialTest(ctx).ArrayRest(arrayRest).Execute()

Credential test before creating a storage array. It requires Storage Manage right.

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
	arrayRest := *openapiclient.NewArrayRest() // ArrayRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageArrayAPI.CredentialTest(context.Background()).ArrayRest(arrayRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.CredentialTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arrayRest** | [**ArrayRest**](ArrayRest.md) |  | 

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


## DeleteArray

> DeleteArray(ctx, arrayId).Execute()

Remove the specific storage array from all appliances. It requires Storage Manage right.

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
	arrayId := "arrayId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageArrayAPI.DeleteArray(context.Background(), arrayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.DeleteArray``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**arrayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArrayRequest struct via the builder pattern


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


## GetArray

> ArrayRest GetArray(ctx, arrayId).Execute()

Get individual storage array details.

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
	arrayId := "arrayId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageArrayAPI.GetArray(context.Background(), arrayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.GetArray``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArray`: ArrayRest
	fmt.Fprintf(os.Stdout, "Response from `StorageArrayAPI.GetArray`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**arrayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArrayRest**](ArrayRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArrayTypes

> ListArrayTypeRest GetArrayTypes(ctx).Execute()

Get all supported storage array types from appliances. It requires Storage View right.

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
	resp, r, err := apiClient.StorageArrayAPI.GetArrayTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.GetArrayTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArrayTypes`: ListArrayTypeRest
	fmt.Fprintf(os.Stdout, "Response from `StorageArrayAPI.GetArrayTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetArrayTypesRequest struct via the builder pattern


### Return type

[**ListArrayTypeRest**](ListArrayTypeRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArrays

> ListArrayRest ListArrays(ctx).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()

Get a list of storage arrays

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
	resp, r, err := apiClient.StorageArrayAPI.ListArrays(context.Background()).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.ListArrays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListArrays`: ListArrayRest
	fmt.Fprintf(os.Stdout, "Response from `StorageArrayAPI.ListArrays`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListArraysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#39;:asc&#39; or &#39;:desc&#39; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **int64** | Limit on the number of results to return | 
 **offset** | **int64** | Used with limit to support pagination | 

### Return type

[**ListArrayRest**](ListArrayRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsForList3

> OptionsRest OptionsForList3(ctx).Execute()

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
	resp, r, err := apiClient.StorageArrayAPI.OptionsForList3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.OptionsForList3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsForList3`: OptionsRest
	fmt.Fprintf(os.Stdout, "Response from `StorageArrayAPI.OptionsForList3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsForList3Request struct via the builder pattern


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


## TestArray

> ListArrayTestResultRest TestArray(ctx, arrayId).Execute()

Connectivity test on the specific storage array. It requires Storage Manage right.

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
	arrayId := "arrayId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageArrayAPI.TestArray(context.Background(), arrayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.TestArray``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestArray`: ListArrayTestResultRest
	fmt.Fprintf(os.Stdout, "Response from `StorageArrayAPI.TestArray`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**arrayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListArrayTestResultRest**](ListArrayTestResultRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArray

> ArrayRest UpdateArray(ctx, arrayId).ArrayRest(arrayRest).Execute()

Update a specific storage array. It requires Storage Manage right.

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
	arrayId := "arrayId_example" // string | 
	arrayRest := *openapiclient.NewArrayRest() // ArrayRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageArrayAPI.UpdateArray(context.Background(), arrayId).ArrayRest(arrayRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageArrayAPI.UpdateArray``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateArray`: ArrayRest
	fmt.Fprintf(os.Stdout, "Response from `StorageArrayAPI.UpdateArray`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**arrayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArrayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **arrayRest** | [**ArrayRest**](ArrayRest.md) |  | 

### Return type

[**ArrayRest**](ArrayRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

