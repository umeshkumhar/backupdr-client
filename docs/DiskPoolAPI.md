# \DiskPoolAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountDiskPools**](DiskPoolAPI.md#CountDiskPools) | **Head** /diskpool | Get a count of total diskpools matching the filters.
[**CreateDiskPool**](DiskPoolAPI.md#CreateDiskPool) | **Post** /diskpool | Create a new diskpool. It requires Storage Manage right.
[**DeleteDiskPool**](DiskPoolAPI.md#DeleteDiskPool) | **Delete** /diskpool/{diskpool_id} | Remove the specific diskpool. It requires Storage Manage right.
[**GetDiskPool**](DiskPoolAPI.md#GetDiskPool) | **Get** /diskpool/{diskpool_id} | Get individual diskpool details.
[**GetDynamicJsonForPoolManage**](DiskPoolAPI.md#GetDynamicJsonForPoolManage) | **Get** /diskpool/vault/listdynamicfields/{pool_type} | List dynamic UI properties based on type of vault pool. Requires Storage Manage right.
[**GetVaultList**](DiskPoolAPI.md#GetVaultList) | **Get** /diskpool/vault/listtype | List types of vault pool. Requires Storage Manage right.
[**ListDiskPools**](DiskPoolAPI.md#ListDiskPools) | **Get** /diskpool | Get a list of diskpools
[**ListGcpVaultPool**](DiskPoolAPI.md#ListGcpVaultPool) | **Get** /diskpool/listgcpvaultpool | Get a list of GCP vault pools. This operation is network costly and UI needs to cache the result
[**OptionsForList8**](DiskPoolAPI.md#OptionsForList8) | **Options** /diskpool | Describes the fields available for filtering and sorting
[**UpdateDiskPool**](DiskPoolAPI.md#UpdateDiskPool) | **Put** /diskpool/{diskpool_id} | Update the specific diskpool. It requires Storage Manage right.



## CountDiskPools

> CountDiskPools(ctx).Filter(filter).Execute()

Get a count of total diskpools matching the filters.

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
	r, err := apiClient.DiskPoolAPI.CountDiskPools(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.CountDiskPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountDiskPoolsRequest struct via the builder pattern


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


## CreateDiskPool

> DiskPoolRest CreateDiskPool(ctx).DiskPoolRest(diskPoolRest).Execute()

Create a new diskpool. It requires Storage Manage right.

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
	diskPoolRest := *openapiclient.NewDiskPoolRest() // DiskPoolRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiskPoolAPI.CreateDiskPool(context.Background()).DiskPoolRest(diskPoolRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.CreateDiskPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDiskPool`: DiskPoolRest
	fmt.Fprintf(os.Stdout, "Response from `DiskPoolAPI.CreateDiskPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiskPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskPoolRest** | [**DiskPoolRest**](DiskPoolRest.md) |  | 

### Return type

[**DiskPoolRest**](DiskPoolRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDiskPool

> DeleteDiskPool(ctx, diskpoolId).Execute()

Remove the specific diskpool. It requires Storage Manage right.

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
	diskpoolId := "diskpoolId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiskPoolAPI.DeleteDiskPool(context.Background(), diskpoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.DeleteDiskPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskpoolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiskPoolRequest struct via the builder pattern


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


## GetDiskPool

> DiskPoolRest GetDiskPool(ctx, diskpoolId).Execute()

Get individual diskpool details.

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
	diskpoolId := "diskpoolId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiskPoolAPI.GetDiskPool(context.Background(), diskpoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.GetDiskPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiskPool`: DiskPoolRest
	fmt.Fprintf(os.Stdout, "Response from `DiskPoolAPI.GetDiskPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskpoolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiskPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiskPoolRest**](DiskPoolRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicJsonForPoolManage

> GetDynamicJsonForPoolManage(ctx, poolType).Execute()

List dynamic UI properties based on type of vault pool. Requires Storage Manage right.

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
	poolType := "poolType_example" // string | The vault type for which dynamic json will be rendered

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiskPoolAPI.GetDynamicJsonForPoolManage(context.Background(), poolType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.GetDynamicJsonForPoolManage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolType** | **string** | The vault type for which dynamic json will be rendered | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicJsonForPoolManageRequest struct via the builder pattern


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


## GetVaultList

> GetVaultList(ctx).Execute()

List types of vault pool. Requires Storage Manage right.

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
	r, err := apiClient.DiskPoolAPI.GetVaultList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.GetVaultList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultListRequest struct via the builder pattern


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


## ListDiskPools

> ListDiskPoolRest ListDiskPools(ctx).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()

Get a list of diskpools

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
	resp, r, err := apiClient.DiskPoolAPI.ListDiskPools(context.Background()).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.ListDiskPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDiskPools`: ListDiskPoolRest
	fmt.Fprintf(os.Stdout, "Response from `DiskPoolAPI.ListDiskPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDiskPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#39;:asc&#39; or &#39;:desc&#39; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **int64** | Limit on the number of results to return | 
 **offset** | **int64** | Used with limit to support pagination | 

### Return type

[**ListDiskPoolRest**](ListDiskPoolRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGcpVaultPool

> ListDiskPoolRest ListGcpVaultPool(ctx).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()

Get a list of GCP vault pools. This operation is network costly and UI needs to cache the result

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
	resp, r, err := apiClient.DiskPoolAPI.ListGcpVaultPool(context.Background()).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.ListGcpVaultPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGcpVaultPool`: ListDiskPoolRest
	fmt.Fprintf(os.Stdout, "Response from `DiskPoolAPI.ListGcpVaultPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGcpVaultPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#39;:asc&#39; or &#39;:desc&#39; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **int64** | Limit on the number of results to return | 
 **offset** | **int64** | Used with limit to support pagination | 

### Return type

[**ListDiskPoolRest**](ListDiskPoolRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsForList8

> OptionsRest OptionsForList8(ctx).Execute()

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
	resp, r, err := apiClient.DiskPoolAPI.OptionsForList8(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.OptionsForList8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsForList8`: OptionsRest
	fmt.Fprintf(os.Stdout, "Response from `DiskPoolAPI.OptionsForList8`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsForList8Request struct via the builder pattern


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


## UpdateDiskPool

> DiskPoolRest UpdateDiskPool(ctx, diskpoolId).DiskPoolRest(diskPoolRest).Execute()

Update the specific diskpool. It requires Storage Manage right.

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
	diskpoolId := "diskpoolId_example" // string | 
	diskPoolRest := *openapiclient.NewDiskPoolRest() // DiskPoolRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiskPoolAPI.UpdateDiskPool(context.Background(), diskpoolId).DiskPoolRest(diskPoolRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiskPoolAPI.UpdateDiskPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDiskPool`: DiskPoolRest
	fmt.Fprintf(os.Stdout, "Response from `DiskPoolAPI.UpdateDiskPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diskpoolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiskPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diskPoolRest** | [**DiskPoolRest**](DiskPoolRest.md) |  | 

### Return type

[**DiskPoolRest**](DiskPoolRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

