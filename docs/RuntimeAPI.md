# \RuntimeAPI

All URIs are relative to *http://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRuntime**](RuntimeAPI.md#CreateRuntime) | **Post** /runtime | Create runtime
[**GetRuntime**](RuntimeAPI.md#GetRuntime) | **Get** /runtime/{id} | Get runtime
[**ListRuntimes**](RuntimeAPI.md#ListRuntimes) | **Get** /runtime | List runtimes



## CreateRuntime

> Runtime CreateRuntime(ctx).CreateRuntime(createRuntime).Execute()

Create runtime



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/onpremless/go-client"
)

func main() {
    createRuntime := *openapiclient.NewCreateRuntime("Dockerfile_example", "Name_example") // CreateRuntime | Create runtime body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuntimeAPI.CreateRuntime(context.Background()).CreateRuntime(createRuntime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeAPI.CreateRuntime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRuntime`: Runtime
    fmt.Fprintf(os.Stdout, "Response from `RuntimeAPI.CreateRuntime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuntimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRuntime** | [**CreateRuntime**](CreateRuntime.md) | Create runtime body | 

### Return type

[**Runtime**](Runtime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuntime

> Runtime GetRuntime(ctx, id).Execute()

Get runtime



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/onpremless/go-client"
)

func main() {
    id := "id_example" // string | Runtime id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuntimeAPI.GetRuntime(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeAPI.GetRuntime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuntime`: Runtime
    fmt.Fprintf(os.Stdout, "Response from `RuntimeAPI.GetRuntime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Runtime id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuntimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Runtime**](Runtime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRuntimes

> []Runtime ListRuntimes(ctx).Execute()

List runtimes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/onpremless/go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuntimeAPI.ListRuntimes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeAPI.ListRuntimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRuntimes`: []Runtime
    fmt.Fprintf(os.Stdout, "Response from `RuntimeAPI.ListRuntimes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRuntimesRequest struct via the builder pattern


### Return type

[**[]Runtime**](Runtime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

