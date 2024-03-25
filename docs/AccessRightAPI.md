# \AccessRightAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRight**](AccessRightAPI.md#GetRight) | **Get** /right/{right_name} | Get details for a specific right. It requires System View right.
[**ListRights**](AccessRightAPI.md#ListRights) | **Get** /right | Get a list of rights in the system. Rights are predefined and cannot be changed. It requires System View right.
[**OptionsForList13**](AccessRightAPI.md#OptionsForList13) | **Options** /right | Describes the fields available for filtering and sorting



## GetRight

> RightRest GetRight(ctx, rightName).Execute()

Get details for a specific right. It requires System View right.

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
	rightName := "rightName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRightAPI.GetRight(context.Background(), rightName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRightAPI.GetRight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRight`: RightRest
	fmt.Fprintf(os.Stdout, "Response from `AccessRightAPI.GetRight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rightName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RightRest**](RightRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRights

> ListRightRest ListRights(ctx).Execute()

Get a list of rights in the system. Rights are predefined and cannot be changed. It requires System View right.

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
	resp, r, err := apiClient.AccessRightAPI.ListRights(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRightAPI.ListRights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRights`: ListRightRest
	fmt.Fprintf(os.Stdout, "Response from `AccessRightAPI.ListRights`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRightsRequest struct via the builder pattern


### Return type

[**ListRightRest**](ListRightRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsForList13

> OptionsRest OptionsForList13(ctx).Execute()

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
	resp, r, err := apiClient.AccessRightAPI.OptionsForList13(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRightAPI.OptionsForList13``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsForList13`: OptionsRest
	fmt.Fprintf(os.Stdout, "Response from `AccessRightAPI.OptionsForList13`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsForList13Request struct via the builder pattern


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

