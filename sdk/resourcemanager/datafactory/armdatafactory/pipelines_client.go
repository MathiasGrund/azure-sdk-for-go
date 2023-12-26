//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PipelinesClient contains the methods for the Pipelines group.
// Don't use this type directly, use NewPipelinesClient() instead.
type PipelinesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPipelinesClient creates a new instance of PipelinesClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPipelinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PipelinesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PipelinesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - pipelineName - The pipeline name.
//   - pipeline - Pipeline resource definition.
//   - options - PipelinesClientCreateOrUpdateOptions contains the optional parameters for the PipelinesClient.CreateOrUpdate
//     method.
func (client *PipelinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline PipelineResource, options *PipelinesClientCreateOrUpdateOptions) (PipelinesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "PipelinesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, pipeline, options)
	if err != nil {
		return PipelinesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelinesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelinesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PipelinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline PipelineResource, options *PipelinesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, pipeline); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PipelinesClient) createOrUpdateHandleResponse(resp *http.Response) (PipelinesClientCreateOrUpdateResponse, error) {
	result := PipelinesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineResource); err != nil {
		return PipelinesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateRun - Creates a run of a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - pipelineName - The pipeline name.
//   - options - PipelinesClientCreateRunOptions contains the optional parameters for the PipelinesClient.CreateRun method.
func (client *PipelinesClient) CreateRun(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientCreateRunOptions) (PipelinesClientCreateRunResponse, error) {
	var err error
	const operationName = "PipelinesClient.CreateRun"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createRunCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, options)
	if err != nil {
		return PipelinesClientCreateRunResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelinesClientCreateRunResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PipelinesClientCreateRunResponse{}, err
	}
	resp, err := client.createRunHandleResponse(httpResp)
	return resp, err
}

// createRunCreateRequest creates the CreateRun request.
func (client *PipelinesClient) createRunCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientCreateRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}/createRun"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	if options != nil && options.ReferencePipelineRunID != nil {
		reqQP.Set("referencePipelineRunId", *options.ReferencePipelineRunID)
	}
	if options != nil && options.IsRecovery != nil {
		reqQP.Set("isRecovery", strconv.FormatBool(*options.IsRecovery))
	}
	if options != nil && options.StartActivityName != nil {
		reqQP.Set("startActivityName", *options.StartActivityName)
	}
	if options != nil && options.StartFromFailure != nil {
		reqQP.Set("startFromFailure", strconv.FormatBool(*options.StartFromFailure))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createRunHandleResponse handles the CreateRun response.
func (client *PipelinesClient) createRunHandleResponse(resp *http.Response) (PipelinesClientCreateRunResponse, error) {
	result := PipelinesClientCreateRunResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreateRunResponse); err != nil {
		return PipelinesClientCreateRunResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - pipelineName - The pipeline name.
//   - options - PipelinesClientDeleteOptions contains the optional parameters for the PipelinesClient.Delete method.
func (client *PipelinesClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientDeleteOptions) (PipelinesClientDeleteResponse, error) {
	var err error
	const operationName = "PipelinesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, options)
	if err != nil {
		return PipelinesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelinesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PipelinesClientDeleteResponse{}, err
	}
	return PipelinesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PipelinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - pipelineName - The pipeline name.
//   - options - PipelinesClientGetOptions contains the optional parameters for the PipelinesClient.Get method.
func (client *PipelinesClient) Get(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientGetOptions) (PipelinesClientGetResponse, error) {
	var err error
	const operationName = "PipelinesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, options)
	if err != nil {
		return PipelinesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotModified) {
		err = runtime.NewResponseError(httpResp)
		return PipelinesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PipelinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelinesClient) getHandleResponse(resp *http.Response) (PipelinesClientGetResponse, error) {
	result := PipelinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineResource); err != nil {
		return PipelinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFactoryPager - Lists pipelines.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - options - PipelinesClientListByFactoryOptions contains the optional parameters for the PipelinesClient.NewListByFactoryPager
//     method.
func (client *PipelinesClient) NewListByFactoryPager(resourceGroupName string, factoryName string, options *PipelinesClientListByFactoryOptions) *runtime.Pager[PipelinesClientListByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelinesClientListByFactoryResponse]{
		More: func(page PipelinesClientListByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelinesClientListByFactoryResponse) (PipelinesClientListByFactoryResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PipelinesClient.NewListByFactoryPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
			}, nil)
			if err != nil {
				return PipelinesClientListByFactoryResponse{}, err
			}
			return client.listByFactoryHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *PipelinesClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *PipelinesClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *PipelinesClient) listByFactoryHandleResponse(resp *http.Response) (PipelinesClientListByFactoryResponse, error) {
	result := PipelinesClientListByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineListResponse); err != nil {
		return PipelinesClientListByFactoryResponse{}, err
	}
	return result, nil
}
