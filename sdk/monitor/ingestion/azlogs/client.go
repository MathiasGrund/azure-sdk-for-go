//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azlogs

import (
	"bytes"
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the Client group.
// Don't use this type directly, use a constructor function instead.
type Client struct {
	internal *azcore.Client
	endpoint string
}

// Upload - Ingestion API used to directly ingest data using Data Collection Rules. Maximum size of of API call is 1 MB.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - ruleID - The immutable Id of the Data Collection Rule resource.
//   - streamName - The streamDeclaration name as defined in the Data Collection Rule.
//   - logs - An array of objects matching the schema defined by the provided stream.
//   - options - UploadOptions contains the optional parameters for the Client.Upload method.
func (client *Client) Upload(ctx context.Context, ruleID string, streamName string, logs []byte, options *UploadOptions) (UploadResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "Client.Upload", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.uploadCreateRequest(ctx, ruleID, streamName, logs, options)
	if err != nil {
		return UploadResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UploadResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UploadResponse{}, err
	}
	return UploadResponse{}, nil
}

// uploadCreateRequest creates the Upload request.
func (client *Client) uploadCreateRequest(ctx context.Context, ruleID string, streamName string, logs []byte, options *UploadOptions) (*policy.Request, error) {
	urlPath := "/dataCollectionRules/{ruleId}/streams/{stream}"
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	if streamName == "" {
		return nil, errors.New("parameter streamName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{stream}", url.PathEscape(streamName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ContentEncoding != nil {
		req.Raw().Header["Content-Encoding"] = []string{*options.ContentEncoding}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := req.SetBody(streaming.NopCloser(bytes.NewReader(logs)), "application/json"); err != nil {
		return nil, err
	}
	return req, nil
}
