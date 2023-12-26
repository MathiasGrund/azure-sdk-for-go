//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
	"net/http"
	"net/url"
	"regexp"
)

// SKUsAvailabilityServer is a fake server for instances of the armvmwarecloudsimple.SKUsAvailabilityClient type.
type SKUsAvailabilityServer struct {
	// NewListPager is the fake for method SKUsAvailabilityClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(regionID string, options *armvmwarecloudsimple.SKUsAvailabilityClientListOptions) (resp azfake.PagerResponder[armvmwarecloudsimple.SKUsAvailabilityClientListResponse])
}

// NewSKUsAvailabilityServerTransport creates a new instance of SKUsAvailabilityServerTransport with the provided implementation.
// The returned SKUsAvailabilityServerTransport instance is connected to an instance of armvmwarecloudsimple.SKUsAvailabilityClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSKUsAvailabilityServerTransport(srv *SKUsAvailabilityServer) *SKUsAvailabilityServerTransport {
	return &SKUsAvailabilityServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armvmwarecloudsimple.SKUsAvailabilityClientListResponse]](),
	}
}

// SKUsAvailabilityServerTransport connects instances of armvmwarecloudsimple.SKUsAvailabilityClient to instances of SKUsAvailabilityServer.
// Don't use this type directly, use NewSKUsAvailabilityServerTransport instead.
type SKUsAvailabilityServerTransport struct {
	srv          *SKUsAvailabilityServer
	newListPager *tracker[azfake.PagerResponder[armvmwarecloudsimple.SKUsAvailabilityClientListResponse]]
}

// Do implements the policy.Transporter interface for SKUsAvailabilityServerTransport.
func (s *SKUsAvailabilityServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SKUsAvailabilityClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SKUsAvailabilityServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VMwareCloudSimple/locations/(?P<regionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/availabilities`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		regionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("regionId")])
		if err != nil {
			return nil, err
		}
		sKUIDUnescaped, err := url.QueryUnescape(qp.Get("skuId"))
		if err != nil {
			return nil, err
		}
		sKUIDParam := getOptional(sKUIDUnescaped)
		var options *armvmwarecloudsimple.SKUsAvailabilityClientListOptions
		if sKUIDParam != nil {
			options = &armvmwarecloudsimple.SKUsAvailabilityClientListOptions{
				SKUID: sKUIDParam,
			}
		}
		resp := s.srv.NewListPager(regionIDParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armvmwarecloudsimple.SKUsAvailabilityClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}
