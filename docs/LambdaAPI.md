# \LambdaAPI

All URIs are relative to *http://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLambda**](LambdaAPI.md#CreateLambda) | **Post** /lambda | Create lambda
[**DestroyLambda**](LambdaAPI.md#DestroyLambda) | **Post** /lambda/{id}/destroy | Destroy lambda
[**GetLambda**](LambdaAPI.md#GetLambda) | **Get** /lambda/{id} | Get lambda
[**ListLambdas**](LambdaAPI.md#ListLambdas) | **Get** /lambda | List lambdas
[**StartLambda**](LambdaAPI.md#StartLambda) | **Post** /lambda/{id}/start | Start lambda



## CreateLambda

> Lambda CreateLambda(ctx).CreateLambda(createLambda).Execute()

Create lambda



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
    createLambda := *openapiclient.NewCreateLambda("Archive_example", "Name_example", "Runtime_example", "LambdaType_example") // CreateLambda | Create lambda body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LambdaAPI.CreateLambda(context.Background()).CreateLambda(createLambda).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LambdaAPI.CreateLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLambda`: Lambda
    fmt.Fprintf(os.Stdout, "Response from `LambdaAPI.CreateLambda`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLambda** | [**CreateLambda**](CreateLambda.md) | Create lambda body | 

### Return type

[**Lambda**](Lambda.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyLambda

> TaskResponse DestroyLambda(ctx, id).Execute()

Destroy lambda



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
    id := "id_example" // string | lambda id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LambdaAPI.DestroyLambda(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LambdaAPI.DestroyLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DestroyLambda`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `LambdaAPI.DestroyLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | lambda id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLambda

> Lambda GetLambda(ctx, id).Execute()

Get lambda



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
    id := "id_example" // string | lambda id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LambdaAPI.GetLambda(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LambdaAPI.GetLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLambda`: Lambda
    fmt.Fprintf(os.Stdout, "Response from `LambdaAPI.GetLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | lambda id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Lambda**](Lambda.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLambdas

> []Lambda ListLambdas(ctx).Execute()

List lambdas



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
    resp, r, err := apiClient.LambdaAPI.ListLambdas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LambdaAPI.ListLambdas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLambdas`: []Lambda
    fmt.Fprintf(os.Stdout, "Response from `LambdaAPI.ListLambdas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLambdasRequest struct via the builder pattern


### Return type

[**[]Lambda**](Lambda.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartLambda

> TaskResponse StartLambda(ctx, id).Execute()

Start lambda



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
    id := "id_example" // string | lambda id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LambdaAPI.StartLambda(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LambdaAPI.StartLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartLambda`: TaskResponse
    fmt.Fprintf(os.Stdout, "Response from `LambdaAPI.StartLambda`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | lambda id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskResponse**](TaskResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

