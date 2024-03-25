# \HostAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplication**](HostAPI.md#AddApplication) | **Post** /host/{host_id}/addapplication | Create a new application on the specific host. It requires Application Manage or Host Manage rights.
[**AppDiscovery**](HostAPI.md#AppDiscovery) | **Post** /host/{host_id}/appdiscovery | Discover applications on the specific host. It requires Application Manage or Host Manage rights.
[**CountHosts**](HostAPI.md#CountHosts) | **Head** /host | Get a count of total hosts matching the filters.
[**CreateHost**](HostAPI.md#CreateHost) | **Post** /host | Create a new host. It requires Application Manage or Host Manage rights.
[**DeleteHost**](HostAPI.md#DeleteHost) | **Delete** /host/{host_id} | Delete a host completely or remove the host from selective appliances. It requires Application Manage or Host Manage rights.
[**EnableConnectorUpgrade**](HostAPI.md#EnableConnectorUpgrade) | **Post** /host/enableconnectorupgrade | 
[**GetHost**](HostAPI.md#GetHost) | **Get** /host/{host_id} | Get individual host details.
[**ListHosts**](HostAPI.md#ListHosts) | **Get** /host | Get a list of host.
[**OptionsForList9**](HostAPI.md#OptionsForList9) | **Options** /host | Describes the fields available for filtering and sorting
[**RevokeCertificate**](HostAPI.md#RevokeCertificate) | **Put** /host/revokeCertificate | Revokes existing certificates of passed in host and blocks them from creating any new certificates. It requires Host Manage rights.
[**UpdateHost**](HostAPI.md#UpdateHost) | **Put** /host/{host_id} | Update a host. It requires Application Manage or Host Manage rights.
[**VmAddNew**](HostAPI.md#VmAddNew) | **Post** /host/{host_id}/host/{cluster_name}/addvms | Add discovered VMs to appliances asynchronously. It requires Hosts Manage right.
[**VmDiscoveryWithoutCluster**](HostAPI.md#VmDiscoveryWithoutCluster) | **Get** /host/{host_id}/discovervm | Discover VMS on the specific host. It requires Host Manage right.



## AddApplication

> ListApplicationRest AddApplication(ctx, hostId).AppCreationRest(appCreationRest).Execute()

Create a new application on the specific host. It requires Application Manage or Host Manage rights.

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
	hostId := "hostId_example" // string | 
	appCreationRest := *openapiclient.NewAppCreationRest() // AppCreationRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.AddApplication(context.Background(), hostId).AppCreationRest(appCreationRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.AddApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddApplication`: ListApplicationRest
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.AddApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appCreationRest** | [**AppCreationRest**](AppCreationRest.md) |  | 

### Return type

[**ListApplicationRest**](ListApplicationRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDiscovery

> ListApplicationRest AppDiscovery(ctx, hostId).AppDiscoveryRest(appDiscoveryRest).Execute()

Discover applications on the specific host. It requires Application Manage or Host Manage rights.

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
	hostId := "hostId_example" // string | 
	appDiscoveryRest := *openapiclient.NewAppDiscoveryRest() // AppDiscoveryRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.AppDiscovery(context.Background(), hostId).AppDiscoveryRest(appDiscoveryRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.AppDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppDiscovery`: ListApplicationRest
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.AppDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appDiscoveryRest** | [**AppDiscoveryRest**](AppDiscoveryRest.md) |  | 

### Return type

[**ListApplicationRest**](ListApplicationRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountHosts

> CountHosts(ctx).Filter(filter).Execute()

Get a count of total hosts matching the filters.

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
	r, err := apiClient.HostAPI.CountHosts(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.CountHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountHostsRequest struct via the builder pattern


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


## CreateHost

> HostRest CreateHost(ctx).HostRest(hostRest).Execute()

Create a new host. It requires Application Manage or Host Manage rights.

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
	hostRest := *openapiclient.NewHostRest() // HostRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.CreateHost(context.Background()).HostRest(hostRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.CreateHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHost`: HostRest
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.CreateHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostRest** | [**HostRest**](HostRest.md) |  | 

### Return type

[**HostRest**](HostRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHost

> DeleteHost(ctx, hostId).Execute()

Delete a host completely or remove the host from selective appliances. It requires Application Manage or Host Manage rights.

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
	hostId := "hostId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.DeleteHost(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.DeleteHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostRequest struct via the builder pattern


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


## EnableConnectorUpgrade

> EnableConnectorUpgrade(ctx).HostRest(hostRest).Execute()



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
	hostRest := []openapiclient.HostRest{*openapiclient.NewHostRest()} // []HostRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.EnableConnectorUpgrade(context.Background()).HostRest(hostRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.EnableConnectorUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableConnectorUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostRest** | [**[]HostRest**](HostRest.md) |  | 

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


## GetHost

> HostRest GetHost(ctx, hostId).FetchExtraInfo(fetchExtraInfo).Execute()

Get individual host details.

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
	hostId := "hostId_example" // string | 
	fetchExtraInfo := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.GetHost(context.Background(), hostId).FetchExtraInfo(fetchExtraInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.GetHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHost`: HostRest
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.GetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchExtraInfo** | **bool** |  | 

### Return type

[**HostRest**](HostRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> ListHostRest ListHosts(ctx).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()

Get a list of host.

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
	resp, r, err := apiClient.HostAPI.ListHosts(context.Background()).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ListHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosts`: ListHostRest
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.ListHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#39;:asc&#39; or &#39;:desc&#39; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **int64** | Limit on the number of results to return | 
 **offset** | **int64** | Used with limit to support pagination | 

### Return type

[**ListHostRest**](ListHostRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsForList9

> OptionsRest OptionsForList9(ctx).Execute()

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
	resp, r, err := apiClient.HostAPI.OptionsForList9(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.OptionsForList9``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsForList9`: OptionsRest
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.OptionsForList9`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsForList9Request struct via the builder pattern


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


## RevokeCertificate

> CertificateRevocationRest RevokeCertificate(ctx).RequestBody(requestBody).Execute()

Revokes existing certificates of passed in host and blocks them from creating any new certificates. It requires Host Manage rights.

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
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.RevokeCertificate(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.RevokeCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeCertificate`: CertificateRevocationRest
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.RevokeCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**CertificateRevocationRest**](CertificateRevocationRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHost

> HostRest UpdateHost(ctx, hostId).Clone(clone).HostRest(hostRest).Execute()

Update a host. It requires Application Manage or Host Manage rights.

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
	hostId := "hostId_example" // string | 
	clone := true // bool |  (optional)
	hostRest := *openapiclient.NewHostRest() // HostRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.UpdateHost(context.Background(), hostId).Clone(clone).HostRest(hostRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.UpdateHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHost`: HostRest
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.UpdateHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clone** | **bool** |  | 
 **hostRest** | [**HostRest**](HostRest.md) |  | 

### Return type

[**HostRest**](HostRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmAddNew

> VmAddNew(ctx, hostId, clusterName).VmDiscoveryRest(vmDiscoveryRest).Execute()

Add discovered VMs to appliances asynchronously. It requires Hosts Manage right.

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
	hostId := "hostId_example" // string | 
	clusterName := "clusterName_example" // string | 
	vmDiscoveryRest := *openapiclient.NewVmDiscoveryRest() // VmDiscoveryRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.VmAddNew(context.Background(), hostId, clusterName).VmDiscoveryRest(vmDiscoveryRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.VmAddNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 
**clusterName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmAddNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmDiscoveryRest** | [**VmDiscoveryRest**](VmDiscoveryRest.md) |  | 

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


## VmDiscoveryWithoutCluster

> VmDiscoveryWithoutCluster(ctx, hostId).Execute()

Discover VMS on the specific host. It requires Host Manage right.

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
	hostId := "hostId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.VmDiscoveryWithoutCluster(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.VmDiscoveryWithoutCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmDiscoveryWithoutClusterRequest struct via the builder pattern


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

