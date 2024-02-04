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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
	"net/http"
	"net/url"
	"regexp"
)

// CloudHsmClusterPrivateEndpointConnectionsServer is a fake server for instances of the armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClient type.
type CloudHsmClusterPrivateEndpointConnectionsServer struct {
	// Create is the fake for method CloudHsmClusterPrivateEndpointConnectionsClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, properties armhardwaresecuritymodules.PrivateEndpointConnection, options *armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClientCreateOptions) (resp azfake.Responder[armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CloudHsmClusterPrivateEndpointConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, options *armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CloudHsmClusterPrivateEndpointConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, options *armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClientGetOptions) (resp azfake.Responder[armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewCloudHsmClusterPrivateEndpointConnectionsServerTransport creates a new instance of CloudHsmClusterPrivateEndpointConnectionsServerTransport with the provided implementation.
// The returned CloudHsmClusterPrivateEndpointConnectionsServerTransport instance is connected to an instance of armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCloudHsmClusterPrivateEndpointConnectionsServerTransport(srv *CloudHsmClusterPrivateEndpointConnectionsServer) *CloudHsmClusterPrivateEndpointConnectionsServerTransport {
	return &CloudHsmClusterPrivateEndpointConnectionsServerTransport{
		srv:         srv,
		beginDelete: newTracker[azfake.PollerResponder[armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClientDeleteResponse]](),
	}
}

// CloudHsmClusterPrivateEndpointConnectionsServerTransport connects instances of armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClient to instances of CloudHsmClusterPrivateEndpointConnectionsServer.
// Don't use this type directly, use NewCloudHsmClusterPrivateEndpointConnectionsServerTransport instead.
type CloudHsmClusterPrivateEndpointConnectionsServerTransport struct {
	srv         *CloudHsmClusterPrivateEndpointConnectionsServer
	beginDelete *tracker[azfake.PollerResponder[armhardwaresecuritymodules.CloudHsmClusterPrivateEndpointConnectionsClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for CloudHsmClusterPrivateEndpointConnectionsServerTransport.
func (c *CloudHsmClusterPrivateEndpointConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CloudHsmClusterPrivateEndpointConnectionsClient.Create":
		resp, err = c.dispatchCreate(req)
	case "CloudHsmClusterPrivateEndpointConnectionsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CloudHsmClusterPrivateEndpointConnectionsClient.Get":
		resp, err = c.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudHsmClusterPrivateEndpointConnectionsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if c.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HardwareSecurityModules/cloudHsmClusters/(?P<cloudHsmClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<peConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhardwaresecuritymodules.PrivateEndpointConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudHsmClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudHsmClusterName")])
	if err != nil {
		return nil, err
	}
	peConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Create(req.Context(), resourceGroupNameParam, cloudHsmClusterNameParam, peConnectionNameParam, body, nil)
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

func (c *CloudHsmClusterPrivateEndpointConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HardwareSecurityModules/cloudHsmClusters/(?P<cloudHsmClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<peConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudHsmClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudHsmClusterName")])
		if err != nil {
			return nil, err
		}
		peConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, cloudHsmClusterNameParam, peConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *CloudHsmClusterPrivateEndpointConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HardwareSecurityModules/cloudHsmClusters/(?P<cloudHsmClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<peConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudHsmClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudHsmClusterName")])
	if err != nil {
		return nil, err
	}
	peConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, cloudHsmClusterNameParam, peConnectionNameParam, nil)
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
