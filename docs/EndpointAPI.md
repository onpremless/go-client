# \EndpointAPI

All URIs are relative to *http://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpoint**](EndpointAPI.md#CreateEndpoint) | **Post** /endpoint | Create endpoint
[**DeleteEndpoint**](EndpointAPI.md#DeleteEndpoint) | **Delete** /endpoint/{id} | Delete endpoint
[**GetEndpoint**](EndpointAPI.md#GetEndpoint) | **Get** /endpoint/{id} | Get endpoint
[**ListEndpoints**](EndpointAPI.md#ListEndpoints) | **Get** /endpoint | List endpoints



## CreateEndpoint

> Endpoint CreateEndpoint(ctx).CreateEndpoint(createEndpoint).Execute()

Create endpoint



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
    createEndpoint := *openapiclient.NewCreateEndpoint("Name_example", "Path_example", "Lambda_example") // CreateEndpoint | Create endpoint body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndpointAPI.CreateEndpoint(context.Background()).CreateEndpoint(createEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.CreateEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEndpoint`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.CreateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEndpoint** | [**CreateEndpoint**](CreateEndpoint.md) | Create endpoint body | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpoint

> DeleteEndpoint(ctx, id).Execute()

Delete endpoint



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
    id := "id_example" // string | endpoint id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EndpointAPI.DeleteEndpoint(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.DeleteEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | endpoint id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpoint

> Endpoint GetEndpoint(ctx, id).Execute()

Get endpoint



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
    id := "id_example" // string | endpoint id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndpointAPI.GetEndpoint(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.GetEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEndpoint`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.GetEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | endpoint id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpoints

> []Endpoint ListEndpoints(ctx).Execute()

List endpoints



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
    resp, r, err := apiClient.EndpointAPI.ListEndpoints(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndpointAPI.ListEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpoints`: []Endpoint
    fmt.Fprintf(os.Stdout, "Response from `EndpointAPI.ListEndpoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsRequest struct via the builder pattern


### Return type

[**[]Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

