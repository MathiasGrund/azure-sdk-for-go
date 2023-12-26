//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbaremetalinfrastructure

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AzureBareMetalInstancesClient contains the methods for the AzureBareMetalInstances group.
// Don't use this type directly, use NewAzureBareMetalInstancesClient() instead.
type AzureBareMetalInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureBareMetalInstancesClient creates a new instance of AzureBareMetalInstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureBareMetalInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureBareMetalInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureBareMetalInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an Azure Bare Metal Instance for the specified subscription, resource group, and instance name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureBareMetalInstanceName - Name of the Azure Bare Metal Instance, also known as the ResourceName.
//   - options - AzureBareMetalInstancesClientGetOptions contains the optional parameters for the AzureBareMetalInstancesClient.Get
//     method.
func (client *AzureBareMetalInstancesClient) Get(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientGetOptions) (AzureBareMetalInstancesClientGetResponse, error) {
	var err error
	const operationName = "AzureBareMetalInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, azureBareMetalInstanceName, options)
	if err != nil {
		return AzureBareMetalInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureBareMetalInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureBareMetalInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AzureBareMetalInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalInstanceName}", url.PathEscape(azureBareMetalInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureBareMetalInstancesClient) getHandleResponse(resp *http.Response) (AzureBareMetalInstancesClientGetResponse, error) {
	result := AzureBareMetalInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalInstance); err != nil {
		return AzureBareMetalInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of Azure Bare Metal Instances in the specified subscription and resource group.
// The operations returns various properties of each Azure Bare Metal Instance.
//
// Generated from API version 2023-08-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - AzureBareMetalInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureBareMetalInstancesClient.NewListByResourceGroupPager
//     method.
func (client *AzureBareMetalInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *AzureBareMetalInstancesClientListByResourceGroupOptions) *runtime.Pager[AzureBareMetalInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureBareMetalInstancesClientListByResourceGroupResponse]{
		More: func(page AzureBareMetalInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureBareMetalInstancesClientListByResourceGroupResponse) (AzureBareMetalInstancesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureBareMetalInstancesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AzureBareMetalInstancesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AzureBareMetalInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AzureBareMetalInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AzureBareMetalInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (AzureBareMetalInstancesClientListByResourceGroupResponse, error) {
	result := AzureBareMetalInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalInstancesListResult); err != nil {
		return AzureBareMetalInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Returns a list of Azure Bare Metal Instances in the specified subscription. The operations
// returns various properties of each Azure Bare Metal Instance.
//
// Generated from API version 2023-08-04-preview
//   - options - AzureBareMetalInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureBareMetalInstancesClient.NewListBySubscriptionPager
//     method.
func (client *AzureBareMetalInstancesClient) NewListBySubscriptionPager(options *AzureBareMetalInstancesClientListBySubscriptionOptions) *runtime.Pager[AzureBareMetalInstancesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureBareMetalInstancesClientListBySubscriptionResponse]{
		More: func(page AzureBareMetalInstancesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureBareMetalInstancesClientListBySubscriptionResponse) (AzureBareMetalInstancesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureBareMetalInstancesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AzureBareMetalInstancesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AzureBareMetalInstancesClient) listBySubscriptionCreateRequest(ctx context.Context, options *AzureBareMetalInstancesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AzureBareMetalInstancesClient) listBySubscriptionHandleResponse(resp *http.Response) (AzureBareMetalInstancesClientListBySubscriptionResponse, error) {
	result := AzureBareMetalInstancesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalInstancesListResult); err != nil {
		return AzureBareMetalInstancesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginRestart - The operation to restart an Azure Bare Metal Instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureBareMetalInstanceName - Name of the Azure Bare Metal Instance, also known as the ResourceName.
//   - options - AzureBareMetalInstancesClientBeginRestartOptions contains the optional parameters for the AzureBareMetalInstancesClient.BeginRestart
//     method.
func (client *AzureBareMetalInstancesClient) BeginRestart(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientBeginRestartOptions) (*runtime.Poller[AzureBareMetalInstancesClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceGroupName, azureBareMetalInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureBareMetalInstancesClientRestartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureBareMetalInstancesClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restart - The operation to restart an Azure Bare Metal Instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-04-preview
func (client *AzureBareMetalInstancesClient) restart(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientBeginRestartOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureBareMetalInstancesClient.BeginRestart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartCreateRequest(ctx, resourceGroupName, azureBareMetalInstanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// restartCreateRequest creates the Restart request.
func (client *AzureBareMetalInstancesClient) restartCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalInstanceName}", url.PathEscape(azureBareMetalInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ForceParameter != nil {
		if err := runtime.MarshalAsJSON(req, *options.ForceParameter); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// BeginShutdown - The operation to shutdown an Azure Bare Metal Instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureBareMetalInstanceName - Name of the Azure Bare Metal Instance, also known as the ResourceName.
//   - options - AzureBareMetalInstancesClientBeginShutdownOptions contains the optional parameters for the AzureBareMetalInstancesClient.BeginShutdown
//     method.
func (client *AzureBareMetalInstancesClient) BeginShutdown(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientBeginShutdownOptions) (*runtime.Poller[AzureBareMetalInstancesClientShutdownResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.shutdown(ctx, resourceGroupName, azureBareMetalInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureBareMetalInstancesClientShutdownResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureBareMetalInstancesClientShutdownResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Shutdown - The operation to shutdown an Azure Bare Metal Instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-04-preview
func (client *AzureBareMetalInstancesClient) shutdown(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientBeginShutdownOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureBareMetalInstancesClient.BeginShutdown"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.shutdownCreateRequest(ctx, resourceGroupName, azureBareMetalInstanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// shutdownCreateRequest creates the Shutdown request.
func (client *AzureBareMetalInstancesClient) shutdownCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientBeginShutdownOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}/shutdown"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalInstanceName}", url.PathEscape(azureBareMetalInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStart - The operation to start an Azure Bare Metal instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureBareMetalInstanceName - Name of the Azure Bare Metal Instance, also known as the ResourceName.
//   - options - AzureBareMetalInstancesClientBeginStartOptions contains the optional parameters for the AzureBareMetalInstancesClient.BeginStart
//     method.
func (client *AzureBareMetalInstancesClient) BeginStart(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientBeginStartOptions) (*runtime.Poller[AzureBareMetalInstancesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, azureBareMetalInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureBareMetalInstancesClientStartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureBareMetalInstancesClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - The operation to start an Azure Bare Metal instance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-04-preview
func (client *AzureBareMetalInstancesClient) start(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureBareMetalInstancesClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, azureBareMetalInstanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startCreateRequest creates the Start request.
func (client *AzureBareMetalInstancesClient) startCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalInstanceName}", url.PathEscape(azureBareMetalInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Patches the Tags field of a Azure Bare Metal Instance for the specified subscription, resource group, and instance
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureBareMetalInstanceName - Name of the Azure Bare Metal Instance, also known as the ResourceName.
//   - tagsParameter - Request body that only contains the new Tags field
//   - options - AzureBareMetalInstancesClientUpdateOptions contains the optional parameters for the AzureBareMetalInstancesClient.Update
//     method.
func (client *AzureBareMetalInstancesClient) Update(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, tagsParameter Tags, options *AzureBareMetalInstancesClientUpdateOptions) (AzureBareMetalInstancesClientUpdateResponse, error) {
	var err error
	const operationName = "AzureBareMetalInstancesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, azureBareMetalInstanceName, tagsParameter, options)
	if err != nil {
		return AzureBareMetalInstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureBareMetalInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureBareMetalInstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AzureBareMetalInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, tagsParameter Tags, options *AzureBareMetalInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalInstanceName}", url.PathEscape(azureBareMetalInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, tagsParameter); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AzureBareMetalInstancesClient) updateHandleResponse(resp *http.Response) (AzureBareMetalInstancesClientUpdateResponse, error) {
	result := AzureBareMetalInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalInstance); err != nil {
		return AzureBareMetalInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
