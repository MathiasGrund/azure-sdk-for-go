//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// ScopeAccessReviewInstanceClient contains the methods for the ScopeAccessReviewInstance group.
// Don't use this type directly, use NewScopeAccessReviewInstanceClient() instead.
type ScopeAccessReviewInstanceClient struct {
	internal *arm.Client
}

// NewScopeAccessReviewInstanceClient creates a new instance of ScopeAccessReviewInstanceClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScopeAccessReviewInstanceClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ScopeAccessReviewInstanceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScopeAccessReviewInstanceClient{
		internal: cl,
	}
	return client, nil
}

// ApplyDecisions - An action to apply all decisions for an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - ScopeAccessReviewInstanceClientApplyDecisionsOptions contains the optional parameters for the ScopeAccessReviewInstanceClient.ApplyDecisions
//     method.
func (client *ScopeAccessReviewInstanceClient) ApplyDecisions(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceClientApplyDecisionsOptions) (ScopeAccessReviewInstanceClientApplyDecisionsResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewInstanceClient.ApplyDecisions"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.applyDecisionsCreateRequest(ctx, scope, scheduleDefinitionID, id, options)
	if err != nil {
		return ScopeAccessReviewInstanceClientApplyDecisionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewInstanceClientApplyDecisionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewInstanceClientApplyDecisionsResponse{}, err
	}
	return ScopeAccessReviewInstanceClientApplyDecisionsResponse{}, nil
}

// applyDecisionsCreateRequest creates the ApplyDecisions request.
func (client *ScopeAccessReviewInstanceClient) applyDecisionsCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceClientApplyDecisionsOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/applyDecisions"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// RecordAllDecisions - An action to approve/deny all decisions for a review with certain filters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - properties - Record all decisions payload.
//   - options - ScopeAccessReviewInstanceClientRecordAllDecisionsOptions contains the optional parameters for the ScopeAccessReviewInstanceClient.RecordAllDecisions
//     method.
func (client *ScopeAccessReviewInstanceClient) RecordAllDecisions(ctx context.Context, scope string, scheduleDefinitionID string, id string, properties RecordAllDecisionsProperties, options *ScopeAccessReviewInstanceClientRecordAllDecisionsOptions) (ScopeAccessReviewInstanceClientRecordAllDecisionsResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewInstanceClient.RecordAllDecisions"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.recordAllDecisionsCreateRequest(ctx, scope, scheduleDefinitionID, id, properties, options)
	if err != nil {
		return ScopeAccessReviewInstanceClientRecordAllDecisionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewInstanceClientRecordAllDecisionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewInstanceClientRecordAllDecisionsResponse{}, err
	}
	return ScopeAccessReviewInstanceClientRecordAllDecisionsResponse{}, nil
}

// recordAllDecisionsCreateRequest creates the RecordAllDecisions request.
func (client *ScopeAccessReviewInstanceClient) recordAllDecisionsCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, id string, properties RecordAllDecisionsProperties, options *ScopeAccessReviewInstanceClientRecordAllDecisionsOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/recordAllDecisions"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// ResetDecisions - An action to reset all decisions for an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - ScopeAccessReviewInstanceClientResetDecisionsOptions contains the optional parameters for the ScopeAccessReviewInstanceClient.ResetDecisions
//     method.
func (client *ScopeAccessReviewInstanceClient) ResetDecisions(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceClientResetDecisionsOptions) (ScopeAccessReviewInstanceClientResetDecisionsResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewInstanceClient.ResetDecisions"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resetDecisionsCreateRequest(ctx, scope, scheduleDefinitionID, id, options)
	if err != nil {
		return ScopeAccessReviewInstanceClientResetDecisionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewInstanceClientResetDecisionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewInstanceClientResetDecisionsResponse{}, err
	}
	return ScopeAccessReviewInstanceClientResetDecisionsResponse{}, nil
}

// resetDecisionsCreateRequest creates the ResetDecisions request.
func (client *ScopeAccessReviewInstanceClient) resetDecisionsCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceClientResetDecisionsOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/resetDecisions"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// SendReminders - An action to send reminders for an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - ScopeAccessReviewInstanceClientSendRemindersOptions contains the optional parameters for the ScopeAccessReviewInstanceClient.SendReminders
//     method.
func (client *ScopeAccessReviewInstanceClient) SendReminders(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceClientSendRemindersOptions) (ScopeAccessReviewInstanceClientSendRemindersResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewInstanceClient.SendReminders"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sendRemindersCreateRequest(ctx, scope, scheduleDefinitionID, id, options)
	if err != nil {
		return ScopeAccessReviewInstanceClientSendRemindersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewInstanceClientSendRemindersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewInstanceClientSendRemindersResponse{}, err
	}
	return ScopeAccessReviewInstanceClientSendRemindersResponse{}, nil
}

// sendRemindersCreateRequest creates the SendReminders request.
func (client *ScopeAccessReviewInstanceClient) sendRemindersCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceClientSendRemindersOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/sendReminders"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Stop - An action to stop an access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - ScopeAccessReviewInstanceClientStopOptions contains the optional parameters for the ScopeAccessReviewInstanceClient.Stop
//     method.
func (client *ScopeAccessReviewInstanceClient) Stop(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceClientStopOptions) (ScopeAccessReviewInstanceClientStopResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewInstanceClient.Stop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, scope, scheduleDefinitionID, id, options)
	if err != nil {
		return ScopeAccessReviewInstanceClientStopResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewInstanceClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewInstanceClientStopResponse{}, err
	}
	return ScopeAccessReviewInstanceClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *ScopeAccessReviewInstanceClient) stopCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceClientStopOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/stop"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
