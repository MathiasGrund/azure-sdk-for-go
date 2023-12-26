//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ProtectionPolicyOperationStatusesServer is a fake server for instances of the armrecoveryservicesbackup.ProtectionPolicyOperationStatusesClient type.
type ProtectionPolicyOperationStatusesServer struct {
	// Get is the fake for method ProtectionPolicyOperationStatusesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string, options *armrecoveryservicesbackup.ProtectionPolicyOperationStatusesClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectionPolicyOperationStatusesClientGetResponse], errResp azfake.ErrorResponder)
}

// NewProtectionPolicyOperationStatusesServerTransport creates a new instance of ProtectionPolicyOperationStatusesServerTransport with the provided implementation.
// The returned ProtectionPolicyOperationStatusesServerTransport instance is connected to an instance of armrecoveryservicesbackup.ProtectionPolicyOperationStatusesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProtectionPolicyOperationStatusesServerTransport(srv *ProtectionPolicyOperationStatusesServer) *ProtectionPolicyOperationStatusesServerTransport {
	return &ProtectionPolicyOperationStatusesServerTransport{srv: srv}
}

// ProtectionPolicyOperationStatusesServerTransport connects instances of armrecoveryservicesbackup.ProtectionPolicyOperationStatusesClient to instances of ProtectionPolicyOperationStatusesServer.
// Don't use this type directly, use NewProtectionPolicyOperationStatusesServerTransport instead.
type ProtectionPolicyOperationStatusesServerTransport struct {
	srv *ProtectionPolicyOperationStatusesServer
}

// Do implements the policy.Transporter interface for ProtectionPolicyOperationStatusesServerTransport.
func (p *ProtectionPolicyOperationStatusesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProtectionPolicyOperationStatusesClient.Get":
		resp, err = p.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProtectionPolicyOperationStatusesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	policyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, policyNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
