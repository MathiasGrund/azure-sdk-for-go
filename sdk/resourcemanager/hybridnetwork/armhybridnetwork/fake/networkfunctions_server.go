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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkFunctionsServer is a fake server for instances of the armhybridnetwork.NetworkFunctionsClient type.
type NetworkFunctionsServer struct {
	// BeginCreateOrUpdate is the fake for method NetworkFunctionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkFunctionName string, parameters armhybridnetwork.NetworkFunction, options *armhybridnetwork.NetworkFunctionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armhybridnetwork.NetworkFunctionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NetworkFunctionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkFunctionName string, options *armhybridnetwork.NetworkFunctionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armhybridnetwork.NetworkFunctionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginExecuteRequest is the fake for method NetworkFunctionsClient.BeginExecuteRequest
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExecuteRequest func(ctx context.Context, resourceGroupName string, networkFunctionName string, parameters armhybridnetwork.ExecuteRequestParameters, options *armhybridnetwork.NetworkFunctionsClientBeginExecuteRequestOptions) (resp azfake.PollerResponder[armhybridnetwork.NetworkFunctionsClientExecuteRequestResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NetworkFunctionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkFunctionName string, options *armhybridnetwork.NetworkFunctionsClientGetOptions) (resp azfake.Responder[armhybridnetwork.NetworkFunctionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method NetworkFunctionsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armhybridnetwork.NetworkFunctionsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armhybridnetwork.NetworkFunctionsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method NetworkFunctionsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armhybridnetwork.NetworkFunctionsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armhybridnetwork.NetworkFunctionsClientListBySubscriptionResponse])

	// UpdateTags is the fake for method NetworkFunctionsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, networkFunctionName string, parameters armhybridnetwork.TagsObject, options *armhybridnetwork.NetworkFunctionsClientUpdateTagsOptions) (resp azfake.Responder[armhybridnetwork.NetworkFunctionsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewNetworkFunctionsServerTransport creates a new instance of NetworkFunctionsServerTransport with the provided implementation.
// The returned NetworkFunctionsServerTransport instance is connected to an instance of armhybridnetwork.NetworkFunctionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkFunctionsServerTransport(srv *NetworkFunctionsServer) *NetworkFunctionsServerTransport {
	return &NetworkFunctionsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armhybridnetwork.NetworkFunctionsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armhybridnetwork.NetworkFunctionsClientDeleteResponse]](),
		beginExecuteRequest:         newTracker[azfake.PollerResponder[armhybridnetwork.NetworkFunctionsClientExecuteRequestResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armhybridnetwork.NetworkFunctionsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armhybridnetwork.NetworkFunctionsClientListBySubscriptionResponse]](),
	}
}

// NetworkFunctionsServerTransport connects instances of armhybridnetwork.NetworkFunctionsClient to instances of NetworkFunctionsServer.
// Don't use this type directly, use NewNetworkFunctionsServerTransport instead.
type NetworkFunctionsServerTransport struct {
	srv                         *NetworkFunctionsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armhybridnetwork.NetworkFunctionsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armhybridnetwork.NetworkFunctionsClientDeleteResponse]]
	beginExecuteRequest         *tracker[azfake.PollerResponder[armhybridnetwork.NetworkFunctionsClientExecuteRequestResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armhybridnetwork.NetworkFunctionsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armhybridnetwork.NetworkFunctionsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for NetworkFunctionsServerTransport.
func (n *NetworkFunctionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkFunctionsClient.BeginCreateOrUpdate":
		resp, err = n.dispatchBeginCreateOrUpdate(req)
	case "NetworkFunctionsClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NetworkFunctionsClient.BeginExecuteRequest":
		resp, err = n.dispatchBeginExecuteRequest(req)
	case "NetworkFunctionsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NetworkFunctionsClient.NewListByResourceGroupPager":
		resp, err = n.dispatchNewListByResourceGroupPager(req)
	case "NetworkFunctionsClient.NewListBySubscriptionPager":
		resp, err = n.dispatchNewListBySubscriptionPager(req)
	case "NetworkFunctionsClient.UpdateTags":
		resp, err = n.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkFunctionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/networkFunctions/(?P<networkFunctionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.NetworkFunction](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkFunctionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFunctionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, networkFunctionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NetworkFunctionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/networkFunctions/(?P<networkFunctionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkFunctionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFunctionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkFunctionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		n.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NetworkFunctionsServerTransport) dispatchBeginExecuteRequest(req *http.Request) (*http.Response, error) {
	if n.srv.BeginExecuteRequest == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExecuteRequest not implemented")}
	}
	beginExecuteRequest := n.beginExecuteRequest.get(req)
	if beginExecuteRequest == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/networkFunctions/(?P<networkFunctionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executeRequest`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.ExecuteRequestParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkFunctionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFunctionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginExecuteRequest(req.Context(), resourceGroupNameParam, networkFunctionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExecuteRequest = &respr
		n.beginExecuteRequest.add(req, beginExecuteRequest)
	}

	resp, err := server.PollerResponderNext(beginExecuteRequest, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginExecuteRequest.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExecuteRequest) {
		n.beginExecuteRequest.remove(req)
	}

	return resp, nil
}

func (n *NetworkFunctionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/networkFunctions/(?P<networkFunctionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkFunctionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFunctionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, networkFunctionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkFunction, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkFunctionsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := n.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/networkFunctions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		n.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armhybridnetwork.NetworkFunctionsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		n.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkFunctionsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := n.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/networkFunctions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := n.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		n.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armhybridnetwork.NetworkFunctionsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		n.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkFunctionsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if n.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/networkFunctions/(?P<networkFunctionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkFunctionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFunctionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.UpdateTags(req.Context(), resourceGroupNameParam, networkFunctionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkFunction, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
