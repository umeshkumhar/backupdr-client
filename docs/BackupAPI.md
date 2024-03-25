# \BackupAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneBackup**](BackupAPI.md#CloneBackup) | **Post** /backup/{backup_id}/clone | Clone a specific backup. It requires Application Manage, Host Manage or Clone Manage rights.
[**CountBackups**](BackupAPI.md#CountBackups) | **Head** /backup | Get a count of total backups matching the filters.
[**DeleteBackup**](BackupAPI.md#DeleteBackup) | **Delete** /backup/{backup_id} | Delete a specific backup. It requires Application Manage, Host Manage or Backup Manage rights.
[**ExpireBackup**](BackupAPI.md#ExpireBackup) | **Post** /backup/{backup_id}/expire | Expire a specific backup. It requires the manageExpiration permission
[**FetchApplicationOptions**](BackupAPI.md#FetchApplicationOptions) | **Get** /backup/{backupid}/applicationOptions | Gets the dynamic list of application options and corresponding default value (if any) for given backup image.
[**GetBackup**](BackupAPI.md#GetBackup) | **Get** /backup/{backup_id} | Get individual backup details.
[**GetDiskMapping**](BackupAPI.md#GetDiskMapping) | **Get** /backup/{backupid}/diskmapping | Gets disk mapping options for restore (source disks, target disks, RAC node list etc)
[**ListBackups**](BackupAPI.md#ListBackups) | **Get** /backup | Get a list of backups.
[**LiveCloneBackup**](BackupAPI.md#LiveCloneBackup) | **Post** /backup/{backup_id}/liveclone | Create live-clone of a specific backup. It requires Application Manage, Host Manage or Liveclone Manage rights.
[**MigrateRestoreMountBackup**](BackupAPI.md#MigrateRestoreMountBackup) | **Post** /backup/{backupid}/restoremigrate | Submit restore-migrate for Oracle restore-mounted image
[**MountBackup**](BackupAPI.md#MountBackup) | **Post** /backup/{backup_id}/mount | Mount a specific backup. It requires Application Manage, Host Manage or Mount Manage rights.
[**MountMigratePreflight**](BackupAPI.md#MountMigratePreflight) | **Post** /backup/{backupid}/mountmigratepreflight | Performs preflight check for Oracle restore-mount/restore-migrate operations
[**OptionsForList4**](BackupAPI.md#OptionsForList4) | **Options** /backup | Describes the fields available for filtering and sorting
[**PrepMountLiveCloneNew**](BackupAPI.md#PrepMountLiveCloneNew) | **Post** /backup/{backup_id}/prepmount | Prep-mount a specific backup. It requires Application Manage, Host Manage, Mount Manage or Liveclone Manage rights.
[**PrepUnmountBackup**](BackupAPI.md#PrepUnmountBackup) | **Post** /backup/{backup_id}/prepunmount | Prep-unmount a specific backup. It requires Application Manage, Host Manage, Mount Manage or Liveclone Manage rights.
[**RefreshLiveClone**](BackupAPI.md#RefreshLiveClone) | **Post** /backup/{backup_id}/refresh | Refresh a live-clone backup. It requires Application Manage, Host Manage or Liveclone Manage rights.
[**RestoreBackup**](BackupAPI.md#RestoreBackup) | **Post** /backup/{backup_id}/restore | Restore an application from a specific backup. It requires Application Manage, Host Manage, Restore Manage or Mirroring Manage rights.
[**UnmountBackup**](BackupAPI.md#UnmountBackup) | **Post** /backup/{backup_id}/unmount | Unmount a specific backup. It requires Application Manage, Host Manage or Mount Manage rights.
[**UpdateBackup**](BackupAPI.md#UpdateBackup) | **Put** /backup/{backup_id} | Update the specific backup. It requires manageExpiration permission to update expiration, and manageBackups permission to update rest of the fields.



## CloneBackup

> CloneBackup(ctx, backupId).CloneRest(cloneRest).Execute()

Clone a specific backup. It requires Application Manage, Host Manage or Clone Manage rights.

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
	backupId := "backupId_example" // string | 
	cloneRest := *openapiclient.NewCloneRest() // CloneRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.CloneBackup(context.Background(), backupId).CloneRest(cloneRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.CloneBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneRest** | [**CloneRest**](CloneRest.md) |  | 

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


## CountBackups

> CountBackups(ctx).Filter(filter).Execute()

Get a count of total backups matching the filters.

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
	r, err := apiClient.BackupAPI.CountBackups(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.CountBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountBackupsRequest struct via the builder pattern


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


## DeleteBackup

> DeleteBackup(ctx, backupId).Execute()

Delete a specific backup. It requires Application Manage, Host Manage or Backup Manage rights.

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
	backupId := "backupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.DeleteBackup(context.Background(), backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.DeleteBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupRequest struct via the builder pattern


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


## ExpireBackup

> ExpireBackup(ctx, backupId).Force(force).Execute()

Expire a specific backup. It requires the manageExpiration permission

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
	backupId := "backupId_example" // string | 
	force := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.ExpireBackup(context.Background(), backupId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.ExpireBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | [default to false]

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


## FetchApplicationOptions

> AppClassRest FetchApplicationOptions(ctx, backupid, backupid2).Execute()

Gets the dynamic list of application options and corresponding default value (if any) for given backup image.

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
	backupid := "backupid_example" // string | Id of the backup image whose application options are required.
	backupid2 := "backupid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.FetchApplicationOptions(context.Background(), backupid, backupid2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.FetchApplicationOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchApplicationOptions`: AppClassRest
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.FetchApplicationOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupid** | **string** | Id of the backup image whose application options are required. | 
**backupid2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchApplicationOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetBackup

> BackupRest GetBackup(ctx, backupId).Execute()

Get individual backup details.

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
	backupId := "backupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetBackup(context.Background(), backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackup`: BackupRest
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupRest**](BackupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiskMapping

> DiskMappingRest GetDiskMapping(ctx, backupid, backupid2).Hostid(hostid).Targetstoragetype(targetstoragetype).Hostid2(hostid2).Targetstoragetype2(targetstoragetype2).Execute()

Gets disk mapping options for restore (source disks, target disks, RAC node list etc)

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
	backupid := "backupid_example" // string | Id of the backup image to be restored
	backupid2 := "backupid_example" // string | 
	hostid := "hostid_example" // string | Restore target host (optional)
	targetstoragetype := "targetstoragetype_example" // string | Restore Target Storage Type (ASM or FS) (optional)
	hostid2 := "hostid_example" // string |  (optional)
	targetstoragetype2 := "targetstoragetype_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetDiskMapping(context.Background(), backupid, backupid2).Hostid(hostid).Targetstoragetype(targetstoragetype).Hostid2(hostid2).Targetstoragetype2(targetstoragetype2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetDiskMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiskMapping`: DiskMappingRest
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetDiskMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupid** | **string** | Id of the backup image to be restored | 
**backupid2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiskMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hostid** | **string** | Restore target host | 
 **targetstoragetype** | **string** | Restore Target Storage Type (ASM or FS) | 
 **hostid2** | **string** |  | 
 **targetstoragetype2** | **string** |  | 

### Return type

[**DiskMappingRest**](DiskMappingRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackups

> ListBackupRest ListBackups(ctx).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()

Get a list of backups.

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
	resp, r, err := apiClient.BackupAPI.ListBackups(context.Background()).Sort(sort).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.ListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackups`: ListBackupRest
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.ListBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort field. Use OPTIONS method to get possible sort fields.&lt;br&gt;Then append &#39;:asc&#39; or &#39;:desc&#39; for ascending or descending sort.&lt;br&gt;Sorting is case-sensitive. | 
 **filter** | **string** | Filter field. Use OPTIONS method to get possible filter fields.&lt;br&gt;Then append an operator and value. Operators always begin with a colon and include:&lt;br&gt;&lt;table&gt;&lt;tr&gt;&lt;th&gt;Operator&lt;/th&gt;&lt;th&gt;Meaning&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;&#x3D;&lt;/td&gt;&lt;td&gt;equals&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;|&lt;/td&gt;&lt;td&gt;contains (case-insensitive)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&gt;&#x3D;&lt;/td&gt;&lt;td&gt;greater than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&lt;&#x3D;&lt;/td&gt;&lt;td&gt;less than or equal to&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td&gt;:&#x3D;b&lt;/td&gt;&lt;td&gt;bitwise and&lt;/td&gt;&lt;/tr&gt;&lt;/table&gt; | 
 **limit** | **int64** | Limit on the number of results to return | 
 **offset** | **int64** | Used with limit to support pagination | 

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


## LiveCloneBackup

> LiveCloneBackup(ctx, backupId).LiveCloneRest(liveCloneRest).Execute()

Create live-clone of a specific backup. It requires Application Manage, Host Manage or Liveclone Manage rights.

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
	backupId := "backupId_example" // string | 
	liveCloneRest := *openapiclient.NewLiveCloneRest() // LiveCloneRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.LiveCloneBackup(context.Background(), backupId).LiveCloneRest(liveCloneRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.LiveCloneBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveCloneBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveCloneRest** | [**LiveCloneRest**](LiveCloneRest.md) |  | 

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


## MigrateRestoreMountBackup

> MigrateRestoreMountBackup(ctx, backupid, backupid2).RestoreMigrateRest(restoreMigrateRest).Execute()

Submit restore-migrate for Oracle restore-mounted image

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
	backupid := "backupid_example" // string | Id of the restore-mounted image to be migrated
	backupid2 := "backupid_example" // string | 
	restoreMigrateRest := *openapiclient.NewRestoreMigrateRest() // RestoreMigrateRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.MigrateRestoreMountBackup(context.Background(), backupid, backupid2).RestoreMigrateRest(restoreMigrateRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.MigrateRestoreMountBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupid** | **string** | Id of the restore-mounted image to be migrated | 
**backupid2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateRestoreMountBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restoreMigrateRest** | [**RestoreMigrateRest**](RestoreMigrateRest.md) |  | 

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


## MountBackup

> MountBackup(ctx, backupId).MountRest(mountRest).Execute()

Mount a specific backup. It requires Application Manage, Host Manage or Mount Manage rights.

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
	backupId := "backupId_example" // string | 
	mountRest := *openapiclient.NewMountRest() // MountRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.MountBackup(context.Background(), backupId).MountRest(mountRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.MountBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMountBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mountRest** | [**MountRest**](MountRest.md) |  | 

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


## MountMigratePreflight

> RestorePreflightRest MountMigratePreflight(ctx, backupid, backupid2).PreflightRest(preflightRest).Execute()

Performs preflight check for Oracle restore-mount/restore-migrate operations

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
	backupid := "backupid_example" // string | Id of the backup image which will be restore-mounted or restore-migrated.
	backupid2 := "backupid_example" // string | 
	preflightRest := *openapiclient.NewPreflightRest() // PreflightRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.MountMigratePreflight(context.Background(), backupid, backupid2).PreflightRest(preflightRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.MountMigratePreflight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MountMigratePreflight`: RestorePreflightRest
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.MountMigratePreflight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupid** | **string** | Id of the backup image which will be restore-mounted or restore-migrated. | 
**backupid2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMountMigratePreflightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **preflightRest** | [**PreflightRest**](PreflightRest.md) |  | 

### Return type

[**RestorePreflightRest**](RestorePreflightRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsForList4

> OptionsRest OptionsForList4(ctx).Execute()

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
	resp, r, err := apiClient.BackupAPI.OptionsForList4(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.OptionsForList4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsForList4`: OptionsRest
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.OptionsForList4`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsForList4Request struct via the builder pattern


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


## PrepMountLiveCloneNew

> PrepMountLiveCloneNew(ctx, backupId).PrepmountRest(prepmountRest).Execute()

Prep-mount a specific backup. It requires Application Manage, Host Manage, Mount Manage or Liveclone Manage rights.

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
	backupId := "backupId_example" // string | 
	prepmountRest := *openapiclient.NewPrepmountRest() // PrepmountRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.PrepMountLiveCloneNew(context.Background(), backupId).PrepmountRest(prepmountRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.PrepMountLiveCloneNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrepMountLiveCloneNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prepmountRest** | [**PrepmountRest**](PrepmountRest.md) |  | 

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


## PrepUnmountBackup

> PrepUnmountBackup(ctx, backupId).UnmountRest(unmountRest).Execute()

Prep-unmount a specific backup. It requires Application Manage, Host Manage, Mount Manage or Liveclone Manage rights.

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
	backupId := "backupId_example" // string | 
	unmountRest := *openapiclient.NewUnmountRest() // UnmountRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.PrepUnmountBackup(context.Background(), backupId).UnmountRest(unmountRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.PrepUnmountBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrepUnmountBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unmountRest** | [**UnmountRest**](UnmountRest.md) |  | 

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


## RefreshLiveClone

> RefreshLiveClone(ctx, backupId).RefreshRest(refreshRest).Execute()

Refresh a live-clone backup. It requires Application Manage, Host Manage or Liveclone Manage rights.

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
	backupId := "backupId_example" // string | 
	refreshRest := *openapiclient.NewRefreshRest() // RefreshRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.RefreshLiveClone(context.Background(), backupId).RefreshRest(refreshRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.RefreshLiveClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshLiveCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshRest** | [**RefreshRest**](RefreshRest.md) |  | 

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


## RestoreBackup

> RestoreBackup(ctx, backupId).RestoreRest(restoreRest).Execute()

Restore an application from a specific backup. It requires Application Manage, Host Manage, Restore Manage or Mirroring Manage rights.

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
	backupId := "backupId_example" // string | 
	restoreRest := *openapiclient.NewRestoreRest() // RestoreRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.RestoreBackup(context.Background(), backupId).RestoreRest(restoreRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.RestoreBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreRest** | [**RestoreRest**](RestoreRest.md) |  | 

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


## UnmountBackup

> UnmountBackup(ctx, backupId).Returnjob(returnjob).UnmountRest(unmountRest).Execute()

Unmount a specific backup. It requires Application Manage, Host Manage or Mount Manage rights.

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
	backupId := "backupId_example" // string | 
	returnjob := true // bool |  (optional)
	unmountRest := *openapiclient.NewUnmountRest() // UnmountRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.UnmountBackup(context.Background(), backupId).Returnjob(returnjob).UnmountRest(unmountRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.UnmountBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmountBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnjob** | **bool** |  | 
 **unmountRest** | [**UnmountRest**](UnmountRest.md) |  | 

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


## UpdateBackup

> BackupRest UpdateBackup(ctx, backupId).BackupRest(backupRest).Execute()

Update the specific backup. It requires manageExpiration permission to update expiration, and manageBackups permission to update rest of the fields.

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
	backupId := "backupId_example" // string | 
	backupRest := *openapiclient.NewBackupRest() // BackupRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.UpdateBackup(context.Background(), backupId).BackupRest(backupRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.UpdateBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBackup`: BackupRest
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.UpdateBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupRest** | [**BackupRest**](BackupRest.md) |  | 

### Return type

[**BackupRest**](BackupRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

