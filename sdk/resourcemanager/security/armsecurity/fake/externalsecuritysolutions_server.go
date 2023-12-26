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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"net/http"
	"net/url"
	"regexp"
)

// ExternalSecuritySolutionsServer is a fake server for instances of the armsecurity.ExternalSecuritySolutionsClient type.
type ExternalSecuritySolutionsServer struct {
	// Get is the fake for method ExternalSecuritySolutionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, ascLocation string, externalSecuritySolutionsName string, options *armsecurity.ExternalSecuritySolutionsClientGetOptions) (resp azfake.Responder[armsecurity.ExternalSecuritySolutionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExternalSecuritySolutionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsecurity.ExternalSecuritySolutionsClientListOptions) (resp azfake.PagerResponder[armsecurity.ExternalSecuritySolutionsClientListResponse])

	// NewListByHomeRegionPager is the fake for method ExternalSecuritySolutionsClient.NewListByHomeRegionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHomeRegionPager func(ascLocation string, options *armsecurity.ExternalSecuritySolutionsClientListByHomeRegionOptions) (resp azfake.PagerResponder[armsecurity.ExternalSecuritySolutionsClientListByHomeRegionResponse])
}

// NewExternalSecuritySolutionsServerTransport creates a new instance of ExternalSecuritySolutionsServerTransport with the provided implementation.
// The returned ExternalSecuritySolutionsServerTransport instance is connected to an instance of armsecurity.ExternalSecuritySolutionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExternalSecuritySolutionsServerTransport(srv *ExternalSecuritySolutionsServer) *ExternalSecuritySolutionsServerTransport {
	return &ExternalSecuritySolutionsServerTransport{
		srv:                      srv,
		newListPager:             newTracker[azfake.PagerResponder[armsecurity.ExternalSecuritySolutionsClientListResponse]](),
		newListByHomeRegionPager: newTracker[azfake.PagerResponder[armsecurity.ExternalSecuritySolutionsClientListByHomeRegionResponse]](),
	}
}

// ExternalSecuritySolutionsServerTransport connects instances of armsecurity.ExternalSecuritySolutionsClient to instances of ExternalSecuritySolutionsServer.
// Don't use this type directly, use NewExternalSecuritySolutionsServerTransport instead.
type ExternalSecuritySolutionsServerTransport struct {
	srv                      *ExternalSecuritySolutionsServer
	newListPager             *tracker[azfake.PagerResponder[armsecurity.ExternalSecuritySolutionsClientListResponse]]
	newListByHomeRegionPager *tracker[azfake.PagerResponder[armsecurity.ExternalSecuritySolutionsClientListByHomeRegionResponse]]
}

// Do implements the policy.Transporter interface for ExternalSecuritySolutionsServerTransport.
func (e *ExternalSecuritySolutionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExternalSecuritySolutionsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExternalSecuritySolutionsClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	case "ExternalSecuritySolutionsClient.NewListByHomeRegionPager":
		resp, err = e.dispatchNewListByHomeRegionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExternalSecuritySolutionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/locations/(?P<ascLocation>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ExternalSecuritySolutions/(?P<externalSecuritySolutionsName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	ascLocationParam, err := url.PathUnescape(matches[regex.SubexpIndex("ascLocation")])
	if err != nil {
		return nil, err
	}
	externalSecuritySolutionsNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("externalSecuritySolutionsName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, ascLocationParam, externalSecuritySolutionsNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExternalSecuritySolutionClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExternalSecuritySolutionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/externalSecuritySolutions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := e.srv.NewListPager(nil)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurity.ExternalSecuritySolutionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		e.newListPager.remove(req)
	}
	return resp, nil
}

func (e *ExternalSecuritySolutionsServerTransport) dispatchNewListByHomeRegionPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByHomeRegionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHomeRegionPager not implemented")}
	}
	newListByHomeRegionPager := e.newListByHomeRegionPager.get(req)
	if newListByHomeRegionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/locations/(?P<ascLocation>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ExternalSecuritySolutions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		ascLocationParam, err := url.PathUnescape(matches[regex.SubexpIndex("ascLocation")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByHomeRegionPager(ascLocationParam, nil)
		newListByHomeRegionPager = &resp
		e.newListByHomeRegionPager.add(req, newListByHomeRegionPager)
		server.PagerResponderInjectNextLinks(newListByHomeRegionPager, req, func(page *armsecurity.ExternalSecuritySolutionsClientListByHomeRegionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHomeRegionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByHomeRegionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHomeRegionPager) {
		e.newListByHomeRegionPager.remove(req)
	}
	return resp, nil
}
