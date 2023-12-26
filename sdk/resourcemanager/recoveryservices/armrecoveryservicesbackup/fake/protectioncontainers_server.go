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

// ProtectionContainersServer is a fake server for instances of the armrecoveryservicesbackup.ProtectionContainersClient type.
type ProtectionContainersServer struct {
	// Get is the fake for method ProtectionContainersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *armrecoveryservicesbackup.ProtectionContainersClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectionContainersClientGetResponse], errResp azfake.ErrorResponder)

	// Inquire is the fake for method ProtectionContainersClient.Inquire
	// HTTP status codes to indicate success: http.StatusAccepted
	Inquire func(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *armrecoveryservicesbackup.ProtectionContainersClientInquireOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectionContainersClientInquireResponse], errResp azfake.ErrorResponder)

	// Refresh is the fake for method ProtectionContainersClient.Refresh
	// HTTP status codes to indicate success: http.StatusAccepted
	Refresh func(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, options *armrecoveryservicesbackup.ProtectionContainersClientRefreshOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectionContainersClientRefreshResponse], errResp azfake.ErrorResponder)

	// Register is the fake for method ProtectionContainersClient.Register
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	Register func(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, parameters armrecoveryservicesbackup.ProtectionContainerResource, options *armrecoveryservicesbackup.ProtectionContainersClientRegisterOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectionContainersClientRegisterResponse], errResp azfake.ErrorResponder)

	// Unregister is the fake for method ProtectionContainersClient.Unregister
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	Unregister func(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *armrecoveryservicesbackup.ProtectionContainersClientUnregisterOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectionContainersClientUnregisterResponse], errResp azfake.ErrorResponder)
}

// NewProtectionContainersServerTransport creates a new instance of ProtectionContainersServerTransport with the provided implementation.
// The returned ProtectionContainersServerTransport instance is connected to an instance of armrecoveryservicesbackup.ProtectionContainersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProtectionContainersServerTransport(srv *ProtectionContainersServer) *ProtectionContainersServerTransport {
	return &ProtectionContainersServerTransport{srv: srv}
}

// ProtectionContainersServerTransport connects instances of armrecoveryservicesbackup.ProtectionContainersClient to instances of ProtectionContainersServer.
// Don't use this type directly, use NewProtectionContainersServerTransport instead.
type ProtectionContainersServerTransport struct {
	srv *ProtectionContainersServer
}

// Do implements the policy.Transporter interface for ProtectionContainersServerTransport.
func (p *ProtectionContainersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProtectionContainersClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProtectionContainersClient.Inquire":
		resp, err = p.dispatchInquire(req)
	case "ProtectionContainersClient.Refresh":
		resp, err = p.dispatchRefresh(req)
	case "ProtectionContainersClient.Register":
		resp, err = p.dispatchRegister(req)
	case "ProtectionContainersClient.Unregister":
		resp, err = p.dispatchUnregister(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProtectionContainersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectionContainers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, fabricNameParam, containerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProtectionContainerResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProtectionContainersServerTransport) dispatchInquire(req *http.Request) (*http.Response, error) {
	if p.srv.Inquire == nil {
		return nil, &nonRetriableError{errors.New("fake for method Inquire not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectionContainers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inquire`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	var options *armrecoveryservicesbackup.ProtectionContainersClientInquireOptions
	if filterParam != nil {
		options = &armrecoveryservicesbackup.ProtectionContainersClientInquireOptions{
			Filter: filterParam,
		}
	}
	respr, errRespr := p.srv.Inquire(req.Context(), vaultNameParam, resourceGroupNameParam, fabricNameParam, containerNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProtectionContainersServerTransport) dispatchRefresh(req *http.Request) (*http.Response, error) {
	if p.srv.Refresh == nil {
		return nil, &nonRetriableError{errors.New("fake for method Refresh not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/refreshContainers`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	var options *armrecoveryservicesbackup.ProtectionContainersClientRefreshOptions
	if filterParam != nil {
		options = &armrecoveryservicesbackup.ProtectionContainersClientRefreshOptions{
			Filter: filterParam,
		}
	}
	respr, errRespr := p.srv.Refresh(req.Context(), vaultNameParam, resourceGroupNameParam, fabricNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProtectionContainersServerTransport) dispatchRegister(req *http.Request) (*http.Response, error) {
	if p.srv.Register == nil {
		return nil, &nonRetriableError{errors.New("fake for method Register not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectionContainers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.ProtectionContainerResource](req)
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Register(req.Context(), vaultNameParam, resourceGroupNameParam, fabricNameParam, containerNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProtectionContainerResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProtectionContainersServerTransport) dispatchUnregister(req *http.Request) (*http.Response, error) {
	if p.srv.Unregister == nil {
		return nil, &nonRetriableError{errors.New("fake for method Unregister not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectionContainers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Unregister(req.Context(), vaultNameParam, resourceGroupNameParam, fabricNameParam, containerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
