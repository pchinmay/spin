# \NetworkControllerApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllByCloudProviderUsingGET**](NetworkControllerApi.md#AllByCloudProviderUsingGET) | **Get** /networks/{cloudProvider} | Retrieve a list of networks for a given cloud provider
[**AllUsingGET1**](NetworkControllerApi.md#AllUsingGET1) | **Get** /networks | Retrieve a list of networks, grouped by cloud provider


# **AllByCloudProviderUsingGET**
> []HashMap AllByCloudProviderUsingGET(ctx, cloudProvider, optional)
Retrieve a list of networks for a given cloud provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProvider** | **string**| cloudProvider | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvider** | **string**| cloudProvider | 
 **xRateLimitApp** | **string**| X-RateLimit-App | 

### Return type

[**[]HashMap**](HashMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AllUsingGET1**
> map[string]interface{} AllUsingGET1(ctx, optional)
Retrieve a list of networks, grouped by cloud provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRateLimitApp** | **string**| X-RateLimit-App | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

