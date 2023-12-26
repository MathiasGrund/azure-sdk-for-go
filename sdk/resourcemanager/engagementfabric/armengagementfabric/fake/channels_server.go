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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric"
	"net/http"
	"net/url"
	"regexp"
)

// ChannelsServer is a fake server for instances of the armengagementfabric.ChannelsClient type.
type ChannelsServer struct {
	// CreateOrUpdate is the fake for method ChannelsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, accountName string, channelName string, channel armengagementfabric.Channel, options *armengagementfabric.ChannelsClientCreateOrUpdateOptions) (resp azfake.Responder[armengagementfabric.ChannelsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ChannelsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, channelName string, options *armengagementfabric.ChannelsClientDeleteOptions) (resp azfake.Responder[armengagementfabric.ChannelsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ChannelsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, channelName string, options *armengagementfabric.ChannelsClientGetOptions) (resp azfake.Responder[armengagementfabric.ChannelsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAccountPager is the fake for method ChannelsClient.NewListByAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAccountPager func(resourceGroupName string, accountName string, options *armengagementfabric.ChannelsClientListByAccountOptions) (resp azfake.PagerResponder[armengagementfabric.ChannelsClientListByAccountResponse])
}

// NewChannelsServerTransport creates a new instance of ChannelsServerTransport with the provided implementation.
// The returned ChannelsServerTransport instance is connected to an instance of armengagementfabric.ChannelsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewChannelsServerTransport(srv *ChannelsServer) *ChannelsServerTransport {
	return &ChannelsServerTransport{
		srv:                   srv,
		newListByAccountPager: newTracker[azfake.PagerResponder[armengagementfabric.ChannelsClientListByAccountResponse]](),
	}
}

// ChannelsServerTransport connects instances of armengagementfabric.ChannelsClient to instances of ChannelsServer.
// Don't use this type directly, use NewChannelsServerTransport instead.
type ChannelsServerTransport struct {
	srv                   *ChannelsServer
	newListByAccountPager *tracker[azfake.PagerResponder[armengagementfabric.ChannelsClientListByAccountResponse]]
}

// Do implements the policy.Transporter interface for ChannelsServerTransport.
func (c *ChannelsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ChannelsClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "ChannelsClient.Delete":
		resp, err = c.dispatchDelete(req)
	case "ChannelsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ChannelsClient.NewListByAccountPager":
		resp, err = c.dispatchNewListByAccountPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ChannelsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EngagementFabric/Accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Channels/(?P<channelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armengagementfabric.Channel](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	channelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("channelName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, accountNameParam, channelNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Channel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ChannelsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EngagementFabric/Accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Channels/(?P<channelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	channelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("channelName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, accountNameParam, channelNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ChannelsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EngagementFabric/Accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Channels/(?P<channelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	channelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("channelName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, channelNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Channel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ChannelsServerTransport) dispatchNewListByAccountPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAccountPager not implemented")}
	}
	newListByAccountPager := c.newListByAccountPager.get(req)
	if newListByAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EngagementFabric/Accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Channels`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByAccountPager(resourceGroupNameParam, accountNameParam, nil)
		newListByAccountPager = &resp
		c.newListByAccountPager.add(req, newListByAccountPager)
	}
	resp, err := server.PagerResponderNext(newListByAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAccountPager) {
		c.newListByAccountPager.remove(req)
	}
	return resp, nil
}
