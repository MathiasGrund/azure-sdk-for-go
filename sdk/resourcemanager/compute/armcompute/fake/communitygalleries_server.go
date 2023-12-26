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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// CommunityGalleriesServer is a fake server for instances of the armcompute.CommunityGalleriesClient type.
type CommunityGalleriesServer struct {
	// Get is the fake for method CommunityGalleriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, publicGalleryName string, options *armcompute.CommunityGalleriesClientGetOptions) (resp azfake.Responder[armcompute.CommunityGalleriesClientGetResponse], errResp azfake.ErrorResponder)
}

// NewCommunityGalleriesServerTransport creates a new instance of CommunityGalleriesServerTransport with the provided implementation.
// The returned CommunityGalleriesServerTransport instance is connected to an instance of armcompute.CommunityGalleriesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCommunityGalleriesServerTransport(srv *CommunityGalleriesServer) *CommunityGalleriesServerTransport {
	return &CommunityGalleriesServerTransport{srv: srv}
}

// CommunityGalleriesServerTransport connects instances of armcompute.CommunityGalleriesClient to instances of CommunityGalleriesServer.
// Don't use this type directly, use NewCommunityGalleriesServerTransport instead.
type CommunityGalleriesServerTransport struct {
	srv *CommunityGalleriesServer
}

// Do implements the policy.Transporter interface for CommunityGalleriesServerTransport.
func (c *CommunityGalleriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CommunityGalleriesClient.Get":
		resp, err = c.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CommunityGalleriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/communityGalleries/(?P<publicGalleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publicGalleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publicGalleryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), locationParam, publicGalleryNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CommunityGallery, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
