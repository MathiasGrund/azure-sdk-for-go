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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// CatalogsServer is a fake server for instances of the armsphere.CatalogsClient type.
type CatalogsServer struct {
	// CountDevices is the fake for method CatalogsClient.CountDevices
	// HTTP status codes to indicate success: http.StatusOK
	CountDevices func(ctx context.Context, resourceGroupName string, catalogName string, options *armsphere.CatalogsClientCountDevicesOptions) (resp azfake.Responder[armsphere.CatalogsClientCountDevicesResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method CatalogsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, catalogName string, resource armsphere.Catalog, options *armsphere.CatalogsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsphere.CatalogsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CatalogsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, catalogName string, options *armsphere.CatalogsClientBeginDeleteOptions) (resp azfake.PollerResponder[armsphere.CatalogsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CatalogsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, catalogName string, options *armsphere.CatalogsClientGetOptions) (resp azfake.Responder[armsphere.CatalogsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method CatalogsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armsphere.CatalogsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armsphere.CatalogsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method CatalogsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armsphere.CatalogsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armsphere.CatalogsClientListBySubscriptionResponse])

	// NewListDeploymentsPager is the fake for method CatalogsClient.NewListDeploymentsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDeploymentsPager func(resourceGroupName string, catalogName string, options *armsphere.CatalogsClientListDeploymentsOptions) (resp azfake.PagerResponder[armsphere.CatalogsClientListDeploymentsResponse])

	// NewListDeviceGroupsPager is the fake for method CatalogsClient.NewListDeviceGroupsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDeviceGroupsPager func(resourceGroupName string, catalogName string, listDeviceGroupsRequest armsphere.ListDeviceGroupsRequest, options *armsphere.CatalogsClientListDeviceGroupsOptions) (resp azfake.PagerResponder[armsphere.CatalogsClientListDeviceGroupsResponse])

	// NewListDeviceInsightsPager is the fake for method CatalogsClient.NewListDeviceInsightsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDeviceInsightsPager func(resourceGroupName string, catalogName string, options *armsphere.CatalogsClientListDeviceInsightsOptions) (resp azfake.PagerResponder[armsphere.CatalogsClientListDeviceInsightsResponse])

	// NewListDevicesPager is the fake for method CatalogsClient.NewListDevicesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDevicesPager func(resourceGroupName string, catalogName string, options *armsphere.CatalogsClientListDevicesOptions) (resp azfake.PagerResponder[armsphere.CatalogsClientListDevicesResponse])

	// Update is the fake for method CatalogsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, catalogName string, properties armsphere.CatalogUpdate, options *armsphere.CatalogsClientUpdateOptions) (resp azfake.Responder[armsphere.CatalogsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCatalogsServerTransport creates a new instance of CatalogsServerTransport with the provided implementation.
// The returned CatalogsServerTransport instance is connected to an instance of armsphere.CatalogsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCatalogsServerTransport(srv *CatalogsServer) *CatalogsServerTransport {
	return &CatalogsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armsphere.CatalogsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armsphere.CatalogsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armsphere.CatalogsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armsphere.CatalogsClientListBySubscriptionResponse]](),
		newListDeploymentsPager:     newTracker[azfake.PagerResponder[armsphere.CatalogsClientListDeploymentsResponse]](),
		newListDeviceGroupsPager:    newTracker[azfake.PagerResponder[armsphere.CatalogsClientListDeviceGroupsResponse]](),
		newListDeviceInsightsPager:  newTracker[azfake.PagerResponder[armsphere.CatalogsClientListDeviceInsightsResponse]](),
		newListDevicesPager:         newTracker[azfake.PagerResponder[armsphere.CatalogsClientListDevicesResponse]](),
	}
}

// CatalogsServerTransport connects instances of armsphere.CatalogsClient to instances of CatalogsServer.
// Don't use this type directly, use NewCatalogsServerTransport instead.
type CatalogsServerTransport struct {
	srv                         *CatalogsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armsphere.CatalogsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armsphere.CatalogsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armsphere.CatalogsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armsphere.CatalogsClientListBySubscriptionResponse]]
	newListDeploymentsPager     *tracker[azfake.PagerResponder[armsphere.CatalogsClientListDeploymentsResponse]]
	newListDeviceGroupsPager    *tracker[azfake.PagerResponder[armsphere.CatalogsClientListDeviceGroupsResponse]]
	newListDeviceInsightsPager  *tracker[azfake.PagerResponder[armsphere.CatalogsClientListDeviceInsightsResponse]]
	newListDevicesPager         *tracker[azfake.PagerResponder[armsphere.CatalogsClientListDevicesResponse]]
}

// Do implements the policy.Transporter interface for CatalogsServerTransport.
func (c *CatalogsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CatalogsClient.CountDevices":
		resp, err = c.dispatchCountDevices(req)
	case "CatalogsClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CatalogsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CatalogsClient.Get":
		resp, err = c.dispatchGet(req)
	case "CatalogsClient.NewListByResourceGroupPager":
		resp, err = c.dispatchNewListByResourceGroupPager(req)
	case "CatalogsClient.NewListBySubscriptionPager":
		resp, err = c.dispatchNewListBySubscriptionPager(req)
	case "CatalogsClient.NewListDeploymentsPager":
		resp, err = c.dispatchNewListDeploymentsPager(req)
	case "CatalogsClient.NewListDeviceGroupsPager":
		resp, err = c.dispatchNewListDeviceGroupsPager(req)
	case "CatalogsClient.NewListDeviceInsightsPager":
		resp, err = c.dispatchNewListDeviceInsightsPager(req)
	case "CatalogsClient.NewListDevicesPager":
		resp, err = c.dispatchNewListDevicesPager(req)
	case "CatalogsClient.Update":
		resp, err = c.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CatalogsServerTransport) dispatchCountDevices(req *http.Request) (*http.Response, error) {
	if c.srv.CountDevices == nil {
		return nil, &nonRetriableError{errors.New("fake for method CountDevices not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/countDevices`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CountDevices(req.Context(), resourceGroupNameParam, catalogNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CountDeviceResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsphere.Catalog](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, catalogNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *CatalogsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, catalogNameParam, nil)
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

func (c *CatalogsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, catalogNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Catalog, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := c.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		c.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armsphere.CatalogsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		c.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := c.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		c.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armsphere.CatalogsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		c.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchNewListDeploymentsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListDeploymentsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDeploymentsPager not implemented")}
	}
	newListDeploymentsPager := c.newListDeploymentsPager.get(req)
	if newListDeploymentsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listDeployments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("$maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsphere.CatalogsClientListDeploymentsOptions
		if filterParam != nil || topParam != nil || skipParam != nil || maxpagesizeParam != nil {
			options = &armsphere.CatalogsClientListDeploymentsOptions{
				Filter:      filterParam,
				Top:         topParam,
				Skip:        skipParam,
				Maxpagesize: maxpagesizeParam,
			}
		}
		resp := c.srv.NewListDeploymentsPager(resourceGroupNameParam, catalogNameParam, options)
		newListDeploymentsPager = &resp
		c.newListDeploymentsPager.add(req, newListDeploymentsPager)
		server.PagerResponderInjectNextLinks(newListDeploymentsPager, req, func(page *armsphere.CatalogsClientListDeploymentsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDeploymentsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListDeploymentsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDeploymentsPager) {
		c.newListDeploymentsPager.remove(req)
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchNewListDeviceGroupsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListDeviceGroupsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDeviceGroupsPager not implemented")}
	}
	newListDeviceGroupsPager := c.newListDeviceGroupsPager.get(req)
	if newListDeviceGroupsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listDeviceGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armsphere.ListDeviceGroupsRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("$maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsphere.CatalogsClientListDeviceGroupsOptions
		if filterParam != nil || topParam != nil || skipParam != nil || maxpagesizeParam != nil {
			options = &armsphere.CatalogsClientListDeviceGroupsOptions{
				Filter:      filterParam,
				Top:         topParam,
				Skip:        skipParam,
				Maxpagesize: maxpagesizeParam,
			}
		}
		resp := c.srv.NewListDeviceGroupsPager(resourceGroupNameParam, catalogNameParam, body, options)
		newListDeviceGroupsPager = &resp
		c.newListDeviceGroupsPager.add(req, newListDeviceGroupsPager)
		server.PagerResponderInjectNextLinks(newListDeviceGroupsPager, req, func(page *armsphere.CatalogsClientListDeviceGroupsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDeviceGroupsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListDeviceGroupsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDeviceGroupsPager) {
		c.newListDeviceGroupsPager.remove(req)
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchNewListDeviceInsightsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListDeviceInsightsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDeviceInsightsPager not implemented")}
	}
	newListDeviceInsightsPager := c.newListDeviceInsightsPager.get(req)
	if newListDeviceInsightsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listDeviceInsights`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("$maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsphere.CatalogsClientListDeviceInsightsOptions
		if filterParam != nil || topParam != nil || skipParam != nil || maxpagesizeParam != nil {
			options = &armsphere.CatalogsClientListDeviceInsightsOptions{
				Filter:      filterParam,
				Top:         topParam,
				Skip:        skipParam,
				Maxpagesize: maxpagesizeParam,
			}
		}
		resp := c.srv.NewListDeviceInsightsPager(resourceGroupNameParam, catalogNameParam, options)
		newListDeviceInsightsPager = &resp
		c.newListDeviceInsightsPager.add(req, newListDeviceInsightsPager)
		server.PagerResponderInjectNextLinks(newListDeviceInsightsPager, req, func(page *armsphere.CatalogsClientListDeviceInsightsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDeviceInsightsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListDeviceInsightsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDeviceInsightsPager) {
		c.newListDeviceInsightsPager.remove(req)
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchNewListDevicesPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListDevicesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDevicesPager not implemented")}
	}
	newListDevicesPager := c.newListDevicesPager.get(req)
	if newListDevicesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listDevices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("$maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsphere.CatalogsClientListDevicesOptions
		if filterParam != nil || topParam != nil || skipParam != nil || maxpagesizeParam != nil {
			options = &armsphere.CatalogsClientListDevicesOptions{
				Filter:      filterParam,
				Top:         topParam,
				Skip:        skipParam,
				Maxpagesize: maxpagesizeParam,
			}
		}
		resp := c.srv.NewListDevicesPager(resourceGroupNameParam, catalogNameParam, options)
		newListDevicesPager = &resp
		c.newListDevicesPager.add(req, newListDevicesPager)
		server.PagerResponderInjectNextLinks(newListDevicesPager, req, func(page *armsphere.CatalogsClientListDevicesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDevicesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListDevicesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDevicesPager) {
		c.newListDevicesPager.remove(req)
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureSphere/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsphere.CatalogUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, catalogNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Catalog, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
