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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
	"net/http"
	"net/url"
	"regexp"
)

// RequestsServer is a fake server for instances of the armcustomerlockbox.RequestsClient type.
type RequestsServer struct {
	// Get is the fake for method RequestsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, requestID string, subscriptionID string, options *armcustomerlockbox.RequestsClientGetOptions) (resp azfake.Responder[armcustomerlockbox.RequestsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RequestsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(subscriptionID string, options *armcustomerlockbox.RequestsClientListOptions) (resp azfake.PagerResponder[armcustomerlockbox.RequestsClientListResponse])

	// UpdateStatus is the fake for method RequestsClient.UpdateStatus
	// HTTP status codes to indicate success: http.StatusOK
	UpdateStatus func(ctx context.Context, subscriptionID string, requestID string, approval armcustomerlockbox.Approval, options *armcustomerlockbox.RequestsClientUpdateStatusOptions) (resp azfake.Responder[armcustomerlockbox.RequestsClientUpdateStatusResponse], errResp azfake.ErrorResponder)
}

// NewRequestsServerTransport creates a new instance of RequestsServerTransport with the provided implementation.
// The returned RequestsServerTransport instance is connected to an instance of armcustomerlockbox.RequestsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRequestsServerTransport(srv *RequestsServer) *RequestsServerTransport {
	return &RequestsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcustomerlockbox.RequestsClientListResponse]](),
	}
}

// RequestsServerTransport connects instances of armcustomerlockbox.RequestsClient to instances of RequestsServer.
// Don't use this type directly, use NewRequestsServerTransport instead.
type RequestsServerTransport struct {
	srv          *RequestsServer
	newListPager *tracker[azfake.PagerResponder[armcustomerlockbox.RequestsClientListResponse]]
}

// Do implements the policy.Transporter interface for RequestsServerTransport.
func (r *RequestsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RequestsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RequestsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	case "RequestsClient.UpdateStatus":
		resp, err = r.dispatchUpdateStatus(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RequestsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerLockbox/requests/(?P<requestId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	requestIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("requestId")])
	if err != nil {
		return nil, err
	}
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), requestIDParam, subscriptionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LockboxRequestResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RequestsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerLockbox/requests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		var options *armcustomerlockbox.RequestsClientListOptions
		if filterParam != nil {
			options = &armcustomerlockbox.RequestsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := r.srv.NewListPager(subscriptionIDParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcustomerlockbox.RequestsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

func (r *RequestsServerTransport) dispatchUpdateStatus(req *http.Request) (*http.Response, error) {
	if r.srv.UpdateStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerLockbox/requests/(?P<requestId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateApproval`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcustomerlockbox.Approval](req)
	if err != nil {
		return nil, err
	}
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	requestIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("requestId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.UpdateStatus(req.Context(), subscriptionIDParam, requestIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Approval, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
