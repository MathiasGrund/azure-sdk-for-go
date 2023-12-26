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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateEndpointConnectionsForSCCPowershellServer is a fake server for instances of the armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClient type.
type PrivateEndpointConnectionsForSCCPowershellServer struct {
	// BeginCreateOrUpdate is the fake for method PrivateEndpointConnectionsForSCCPowershellClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, properties armm365securityandcompliance.PrivateEndpointConnection, options *armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateEndpointConnectionsForSCCPowershellClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientBeginDeleteOptions) (resp azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionsForSCCPowershellClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientGetOptions) (resp azfake.Responder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method PrivateEndpointConnectionsForSCCPowershellClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, resourceName string, options *armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientListByServiceOptions) (resp azfake.PagerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientListByServiceResponse])
}

// NewPrivateEndpointConnectionsForSCCPowershellServerTransport creates a new instance of PrivateEndpointConnectionsForSCCPowershellServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionsForSCCPowershellServerTransport instance is connected to an instance of armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionsForSCCPowershellServerTransport(srv *PrivateEndpointConnectionsForSCCPowershellServer) *PrivateEndpointConnectionsForSCCPowershellServerTransport {
	return &PrivateEndpointConnectionsForSCCPowershellServerTransport{
		srv:                   srv,
		beginCreateOrUpdate:   newTracker[azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientCreateOrUpdateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientDeleteResponse]](),
		newListByServicePager: newTracker[azfake.PagerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientListByServiceResponse]](),
	}
}

// PrivateEndpointConnectionsForSCCPowershellServerTransport connects instances of armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClient to instances of PrivateEndpointConnectionsForSCCPowershellServer.
// Don't use this type directly, use NewPrivateEndpointConnectionsForSCCPowershellServerTransport instead.
type PrivateEndpointConnectionsForSCCPowershellServerTransport struct {
	srv                   *PrivateEndpointConnectionsForSCCPowershellServer
	beginCreateOrUpdate   *tracker[azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientCreateOrUpdateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientDeleteResponse]]
	newListByServicePager *tracker[azfake.PagerResponder[armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionsForSCCPowershellServerTransport.
func (p *PrivateEndpointConnectionsForSCCPowershellServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateEndpointConnectionsForSCCPowershellClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PrivateEndpointConnectionsForSCCPowershellClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PrivateEndpointConnectionsForSCCPowershellClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateEndpointConnectionsForSCCPowershellClient.NewListByServicePager":
		resp, err = p.dispatchNewListByServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsForSCCPowershellServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armm365securityandcompliance.PrivateEndpointConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, privateEndpointConnectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsForSCCPowershellServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, privateEndpointConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsForSCCPowershellServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsForSCCPowershellServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := p.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByServicePager(resourceGroupNameParam, resourceNameParam, nil)
		newListByServicePager = &resp
		p.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armm365securityandcompliance.PrivateEndpointConnectionsForSCCPowershellClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		p.newListByServicePager.remove(req)
	}
	return resp, nil
}
