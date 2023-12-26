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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
	"net/http"
	"net/url"
	"regexp"
)

// KubernetesVersionsServer is a fake server for instances of the armhybridcontainerservice.KubernetesVersionsClient type.
type KubernetesVersionsServer struct {
	// NewListPager is the fake for method KubernetesVersionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(customLocationResourceURI string, options *armhybridcontainerservice.KubernetesVersionsClientListOptions) (resp azfake.PagerResponder[armhybridcontainerservice.KubernetesVersionsClientListResponse])
}

// NewKubernetesVersionsServerTransport creates a new instance of KubernetesVersionsServerTransport with the provided implementation.
// The returned KubernetesVersionsServerTransport instance is connected to an instance of armhybridcontainerservice.KubernetesVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewKubernetesVersionsServerTransport(srv *KubernetesVersionsServer) *KubernetesVersionsServerTransport {
	return &KubernetesVersionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armhybridcontainerservice.KubernetesVersionsClientListResponse]](),
	}
}

// KubernetesVersionsServerTransport connects instances of armhybridcontainerservice.KubernetesVersionsClient to instances of KubernetesVersionsServer.
// Don't use this type directly, use NewKubernetesVersionsServerTransport instead.
type KubernetesVersionsServerTransport struct {
	srv          *KubernetesVersionsServer
	newListPager *tracker[azfake.PagerResponder[armhybridcontainerservice.KubernetesVersionsClientListResponse]]
}

// Do implements the policy.Transporter interface for KubernetesVersionsServerTransport.
func (k *KubernetesVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "KubernetesVersionsClient.NewListPager":
		resp, err = k.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (k *KubernetesVersionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if k.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := k.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<customLocationResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/kubernetesVersions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		customLocationResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("customLocationResourceUri")])
		if err != nil {
			return nil, err
		}
		resp := k.srv.NewListPager(customLocationResourceURIParam, nil)
		newListPager = &resp
		k.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armhybridcontainerservice.KubernetesVersionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		k.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		k.newListPager.remove(req)
	}
	return resp, nil
}
