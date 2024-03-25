# \LogicalGroupAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountLogicalGroups**](LogicalGroupAPI.md#CountLogicalGroups) | **Head** /logicalgroup | Get a count of total logical groups matching the filters.
[**CreateLogicalGroup**](LogicalGroupAPI.md#CreateLogicalGroup) | **Post** /logicalgroup | Create a new logical group. It requires Application Manage or System Manage rights.
[**CreateLogicalGroupSla**](LogicalGroupAPI.md#CreateLogicalGroupSla) | **Post** /logicalgroup/{group_id}/sla | Protect a logical group. It creates individual SLAs for all members. It requires SLA Assign right.
[**DeleteLogicalGroup**](LogicalGroupAPI.md#DeleteLogicalGroup) | **Delete** /logicalgroup/{group_id} | Remove a logical group. It requires Application Manage or System Manage rights.
[**DeleteLogicalGroupSla**](LogicalGroupAPI.md#DeleteLogicalGroupSla) | **Delete** /logicalgroup/{group_id}/sla | Unprotect a logical group. It removes SLAs for all members. It requires SLA Assign right.
[**GetLogicalGroup**](LogicalGroupAPI.md#GetLogicalGroup) | **Get** /logicalgroup/{group_id} | Get individual logical group details. It requires System View right.
[**ListLogicalGroup**](LogicalGroupAPI.md#ListLogicalGroup) | **Get** /logicalgroup | Get a list of logical groups. It requires System View right.
[**ListLogicalGroupMembers**](LogicalGroupAPI.md#ListLogicalGroupMembers) | **Get** /logicalgroup/{group_id}/member | Get logical group&#39;s members. It requires SLA Assign or System View rights.
[**ModifyLogicalGroupMembers**](LogicalGroupAPI.md#ModifyLogicalGroupMembers) | **Post** /logicalgroup/{group_id}/member | Incrementally add/delete logical group members. It requires System Manage or SLA Assign rights.
[**OptionsForList11**](LogicalGroupAPI.md#OptionsForList11) | **Options** /logicalgroup | Describes the fields available for filtering and sorting
[**UpdateLogicalGroup**](LogicalGroupAPI.md#UpdateLogicalGroup) | **Put** /logicalgroup/{group_id} | Update a logical group. It requires Application Manage or System Manage rights.
[**UpdateLogicalGroupSla**](LogicalGroupAPI.md#UpdateLogicalGroupSla) | **Put** /logicalgroup/{group_id}/sla | Update the current SLAs for a logical group. It updates individual SLAs for all members. It requires SLA Manage right.



## CountLogicalGroups

> CountLogicalGroups(ctx).Filter(filter).Execute()

Get a count of total logical groups matching the filters.

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
	r, err := apiClient.LogicalGroupAPI.CountLogicalGroups(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.CountLogicalGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountLogicalGroupsRequest struct via the builder pattern


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


## CreateLogicalGroup

> LogicalGroupRest CreateLogicalGroup(ctx).LogicalGroupRest(logicalGroupRest).Execute()

Create a new logical group. It requires Application Manage or System Manage rights.

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
	logicalGroupRest := *openapiclient.NewLogicalGroupRest() // LogicalGroupRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalGroupAPI.CreateLogicalGroup(context.Background()).LogicalGroupRest(logicalGroupRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.CreateLogicalGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLogicalGroup`: LogicalGroupRest
	fmt.Fprintf(os.Stdout, "Response from `LogicalGroupAPI.CreateLogicalGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logicalGroupRest** | [**LogicalGroupRest**](LogicalGroupRest.md) |  | 

### Return type

[**LogicalGroupRest**](LogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogicalGroupSla

> CreateLogicalGroupSla(ctx, groupId).SlaRest(slaRest).Execute()

Protect a logical group. It creates individual SLAs for all members. It requires SLA Assign right.

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
	groupId := "groupId_example" // string | 
	slaRest := *openapiclient.NewSlaRest() // SlaRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalGroupAPI.CreateLogicalGroupSla(context.Background(), groupId).SlaRest(slaRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.CreateLogicalGroupSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogicalGroupSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slaRest** | [**SlaRest**](SlaRest.md) |  | 

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


## DeleteLogicalGroup

> DeleteLogicalGroup(ctx, groupId).Execute()

Remove a logical group. It requires Application Manage or System Manage rights.

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalGroupAPI.DeleteLogicalGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.DeleteLogicalGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalGroupRequest struct via the builder pattern


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


## DeleteLogicalGroupSla

> DeleteLogicalGroupSla(ctx, groupId).Execute()

Unprotect a logical group. It removes SLAs for all members. It requires SLA Assign right.

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalGroupAPI.DeleteLogicalGroupSla(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.DeleteLogicalGroupSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogicalGroupSlaRequest struct via the builder pattern


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


## GetLogicalGroup

> LogicalGroupRest GetLogicalGroup(ctx, groupId).Execute()

Get individual logical group details. It requires System View right.

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalGroupAPI.GetLogicalGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.GetLogicalGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogicalGroup`: LogicalGroupRest
	fmt.Fprintf(os.Stdout, "Response from `LogicalGroupAPI.GetLogicalGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogicalGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogicalGroupRest**](LogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogicalGroup

> ListLogicalGroupRest ListLogicalGroup(ctx).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()

Get a list of logical groups. It requires System View right.

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
	resp, r, err := apiClient.LogicalGroupAPI.ListLogicalGroup(context.Background()).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.ListLogicalGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogicalGroup`: ListLogicalGroupRest
	fmt.Fprintf(os.Stdout, "Response from `LogicalGroupAPI.ListLogicalGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogicalGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#39;:asc&#39; or &#39;:desc&#39; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **int64** | Limit on the number of results to return | 
 **offset** | **int64** | Used with limit to support pagination | 

### Return type

[**ListLogicalGroupRest**](ListLogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogicalGroupMembers

> ListApplicationRest ListLogicalGroupMembers(ctx, groupId).Execute()

Get logical group's members. It requires SLA Assign or System View rights.

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalGroupAPI.ListLogicalGroupMembers(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.ListLogicalGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogicalGroupMembers`: ListApplicationRest
	fmt.Fprintf(os.Stdout, "Response from `LogicalGroupAPI.ListLogicalGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLogicalGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ModifyLogicalGroupMembers

> ModifyLogicalGroupMembers(ctx, groupId).MembershipChangeRest(membershipChangeRest).Execute()

Incrementally add/delete logical group members. It requires System Manage or SLA Assign rights.

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
	groupId := "groupId_example" // string | 
	membershipChangeRest := []openapiclient.MembershipChangeRest{*openapiclient.NewMembershipChangeRest()} // []MembershipChangeRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogicalGroupAPI.ModifyLogicalGroupMembers(context.Background(), groupId).MembershipChangeRest(membershipChangeRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.ModifyLogicalGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyLogicalGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membershipChangeRest** | [**[]MembershipChangeRest**](MembershipChangeRest.md) |  | 

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


## OptionsForList11

> OptionsRest OptionsForList11(ctx).Execute()

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
	resp, r, err := apiClient.LogicalGroupAPI.OptionsForList11(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.OptionsForList11``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsForList11`: OptionsRest
	fmt.Fprintf(os.Stdout, "Response from `LogicalGroupAPI.OptionsForList11`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsForList11Request struct via the builder pattern


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


## UpdateLogicalGroup

> LogicalGroupRest UpdateLogicalGroup(ctx, groupId).LogicalGroupRest(logicalGroupRest).Execute()

Update a logical group. It requires Application Manage or System Manage rights.

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
	groupId := "groupId_example" // string | 
	logicalGroupRest := *openapiclient.NewLogicalGroupRest() // LogicalGroupRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalGroupAPI.UpdateLogicalGroup(context.Background(), groupId).LogicalGroupRest(logicalGroupRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.UpdateLogicalGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogicalGroup`: LogicalGroupRest
	fmt.Fprintf(os.Stdout, "Response from `LogicalGroupAPI.UpdateLogicalGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogicalGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logicalGroupRest** | [**LogicalGroupRest**](LogicalGroupRest.md) |  | 

### Return type

[**LogicalGroupRest**](LogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogicalGroupSla

> LogicalGroupRest UpdateLogicalGroupSla(ctx, groupId).SlaRest(slaRest).Execute()

Update the current SLAs for a logical group. It updates individual SLAs for all members. It requires SLA Manage right.

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
	groupId := "groupId_example" // string | 
	slaRest := *openapiclient.NewSlaRest() // SlaRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogicalGroupAPI.UpdateLogicalGroupSla(context.Background(), groupId).SlaRest(slaRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogicalGroupAPI.UpdateLogicalGroupSla``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLogicalGroupSla`: LogicalGroupRest
	fmt.Fprintf(os.Stdout, "Response from `LogicalGroupAPI.UpdateLogicalGroupSla`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogicalGroupSlaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slaRest** | [**SlaRest**](SlaRest.md) |  | 

### Return type

[**LogicalGroupRest**](LogicalGroupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

