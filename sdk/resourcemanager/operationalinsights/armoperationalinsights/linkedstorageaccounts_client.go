//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// LinkedStorageAccountsClient contains the methods for the LinkedStorageAccounts group.
// Don't use this type directly, use NewLinkedStorageAccountsClient() instead.
type LinkedStorageAccountsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLinkedStorageAccountsClient creates a new instance of LinkedStorageAccountsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLinkedStorageAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LinkedStorageAccountsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LinkedStorageAccountsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update a link relation between current workspace and a group of storage accounts of a specific
// data source type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataSourceType - Linked storage accounts type.
//   - parameters - The parameters required to create or update linked storage accounts.
//   - options - LinkedStorageAccountsClientCreateOrUpdateOptions contains the optional parameters for the LinkedStorageAccountsClient.CreateOrUpdate
//     method.
func (client *LinkedStorageAccountsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, parameters LinkedStorageAccountsResource, options *LinkedStorageAccountsClientCreateOrUpdateOptions) (LinkedStorageAccountsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "LinkedStorageAccountsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceType, parameters, options)
	if err != nil {
		return LinkedStorageAccountsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkedStorageAccountsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LinkedStorageAccountsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LinkedStorageAccountsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, parameters LinkedStorageAccountsResource, options *LinkedStorageAccountsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts/{dataSourceType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceType == "" {
		return nil, errors.New("parameter dataSourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceType}", url.PathEscape(string(dataSourceType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LinkedStorageAccountsClient) createOrUpdateHandleResponse(resp *http.Response) (LinkedStorageAccountsClientCreateOrUpdateResponse, error) {
	result := LinkedStorageAccountsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedStorageAccountsResource); err != nil {
		return LinkedStorageAccountsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes all linked storage accounts of a specific data source type associated with the specified workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataSourceType - Linked storage accounts type.
//   - options - LinkedStorageAccountsClientDeleteOptions contains the optional parameters for the LinkedStorageAccountsClient.Delete
//     method.
func (client *LinkedStorageAccountsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsClientDeleteOptions) (LinkedStorageAccountsClientDeleteResponse, error) {
	var err error
	const operationName = "LinkedStorageAccountsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceType, options)
	if err != nil {
		return LinkedStorageAccountsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkedStorageAccountsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LinkedStorageAccountsClientDeleteResponse{}, err
	}
	return LinkedStorageAccountsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkedStorageAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts/{dataSourceType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceType == "" {
		return nil, errors.New("parameter dataSourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceType}", url.PathEscape(string(dataSourceType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets all linked storage account of a specific data source type associated with the specified workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataSourceType - Linked storage accounts type.
//   - options - LinkedStorageAccountsClientGetOptions contains the optional parameters for the LinkedStorageAccountsClient.Get
//     method.
func (client *LinkedStorageAccountsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsClientGetOptions) (LinkedStorageAccountsClientGetResponse, error) {
	var err error
	const operationName = "LinkedStorageAccountsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceType, options)
	if err != nil {
		return LinkedStorageAccountsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkedStorageAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LinkedStorageAccountsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LinkedStorageAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceType DataSourceType, options *LinkedStorageAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts/{dataSourceType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceType == "" {
		return nil, errors.New("parameter dataSourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceType}", url.PathEscape(string(dataSourceType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedStorageAccountsClient) getHandleResponse(resp *http.Response) (LinkedStorageAccountsClientGetResponse, error) {
	result := LinkedStorageAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedStorageAccountsResource); err != nil {
		return LinkedStorageAccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Gets all linked storage accounts associated with the specified workspace, storage accounts will
// be sorted by their data source type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - LinkedStorageAccountsClientListByWorkspaceOptions contains the optional parameters for the LinkedStorageAccountsClient.NewListByWorkspacePager
//     method.
func (client *LinkedStorageAccountsClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *LinkedStorageAccountsClientListByWorkspaceOptions) *runtime.Pager[LinkedStorageAccountsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkedStorageAccountsClientListByWorkspaceResponse]{
		More: func(page LinkedStorageAccountsClientListByWorkspaceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *LinkedStorageAccountsClientListByWorkspaceResponse) (LinkedStorageAccountsClientListByWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LinkedStorageAccountsClient.NewListByWorkspacePager")
			req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			if err != nil {
				return LinkedStorageAccountsClientListByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LinkedStorageAccountsClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LinkedStorageAccountsClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *LinkedStorageAccountsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *LinkedStorageAccountsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedStorageAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *LinkedStorageAccountsClient) listByWorkspaceHandleResponse(resp *http.Response) (LinkedStorageAccountsClientListByWorkspaceResponse, error) {
	result := LinkedStorageAccountsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedStorageAccountsListResult); err != nil {
		return LinkedStorageAccountsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
