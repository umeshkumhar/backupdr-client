# \UserSessionAPI

All URIs are relative to */actifio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissions**](UserSessionAPI.md#GetPermissions) | **Get** /session/permissions | Get IAM permissions for the user
[**GetSessionInfo**](UserSessionAPI.md#GetSessionInfo) | **Get** /session/current | Get session information, including user preferences
[**Login**](UserSessionAPI.md#Login) | **Post** /session | Register an user session
[**Logout**](UserSessionAPI.md#Logout) | **Delete** /session/current | Deregister the current user session
[**PutUserpref**](UserSessionAPI.md#PutUserpref) | **Put** /session/current | Update user preferences



## GetPermissions

> GetPermissions(ctx).Execute()

Get IAM permissions for the user

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
	r, err := apiClient.UserSessionAPI.GetPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSessionAPI.GetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsRequest struct via the builder pattern


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


## GetSessionInfo

> SessionRest GetSessionInfo(ctx).Execute()

Get session information, including user preferences

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
	resp, r, err := apiClient.UserSessionAPI.GetSessionInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSessionAPI.GetSessionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionInfo`: SessionRest
	fmt.Fprintf(os.Stdout, "Response from `UserSessionAPI.GetSessionInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionInfoRequest struct via the builder pattern


### Return type

[**SessionRest**](SessionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> SessionRest Login(ctx).Execute()

Register an user session



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
	resp, r, err := apiClient.UserSessionAPI.Login(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSessionAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: SessionRest
	fmt.Fprintf(os.Stdout, "Response from `UserSessionAPI.Login`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


### Return type

[**SessionRest**](SessionRest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx).Execute()

Deregister the current user session



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
	r, err := apiClient.UserSessionAPI.Logout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSessionAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## PutUserpref

> PutUserpref(ctx).SessionRest(sessionRest).Execute()

Update user preferences

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
	sessionRest := *openapiclient.NewSessionRest() // SessionRest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserSessionAPI.PutUserpref(context.Background()).SessionRest(sessionRest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSessionAPI.PutUserpref``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUserprefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionRest** | [**SessionRest**](SessionRest.md) |  | 

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

