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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"net/http"
	"net/url"
	"regexp"
)

// AccountsServer is a fake server for instances of the armbilling.AccountsClient type.
type AccountsServer struct {
	// Get is the fake for method AccountsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, billingAccountName string, options *armbilling.AccountsClientGetOptions) (resp azfake.Responder[armbilling.AccountsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AccountsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armbilling.AccountsClientListOptions) (resp azfake.PagerResponder[armbilling.AccountsClientListResponse])

	// NewListInvoiceSectionsByCreateSubscriptionPermissionPager is the fake for method AccountsClient.NewListInvoiceSectionsByCreateSubscriptionPermissionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListInvoiceSectionsByCreateSubscriptionPermissionPager func(billingAccountName string, options *armbilling.AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionOptions) (resp azfake.PagerResponder[armbilling.AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse])

	// BeginUpdate is the fake for method AccountsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, billingAccountName string, parameters armbilling.AccountUpdateRequest, options *armbilling.AccountsClientBeginUpdateOptions) (resp azfake.PollerResponder[armbilling.AccountsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAccountsServerTransport creates a new instance of AccountsServerTransport with the provided implementation.
// The returned AccountsServerTransport instance is connected to an instance of armbilling.AccountsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccountsServerTransport(srv *AccountsServer) *AccountsServerTransport {
	return &AccountsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armbilling.AccountsClientListResponse]](),
		newListInvoiceSectionsByCreateSubscriptionPermissionPager: newTracker[azfake.PagerResponder[armbilling.AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse]](),
		beginUpdate: newTracker[azfake.PollerResponder[armbilling.AccountsClientUpdateResponse]](),
	}
}

// AccountsServerTransport connects instances of armbilling.AccountsClient to instances of AccountsServer.
// Don't use this type directly, use NewAccountsServerTransport instead.
type AccountsServerTransport struct {
	srv                                                       *AccountsServer
	newListPager                                              *tracker[azfake.PagerResponder[armbilling.AccountsClientListResponse]]
	newListInvoiceSectionsByCreateSubscriptionPermissionPager *tracker[azfake.PagerResponder[armbilling.AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse]]
	beginUpdate                                               *tracker[azfake.PollerResponder[armbilling.AccountsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for AccountsServerTransport.
func (a *AccountsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AccountsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AccountsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AccountsClient.NewListInvoiceSectionsByCreateSubscriptionPermissionPager":
		resp, err = a.dispatchNewListInvoiceSectionsByCreateSubscriptionPermissionPager(req)
	case "AccountsClient.BeginUpdate":
		resp, err = a.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccountsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armbilling.AccountsClientGetOptions
	if expandParam != nil {
		options = &armbilling.AccountsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := a.srv.Get(req.Context(), billingAccountNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Account, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		qp := req.URL.Query()
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armbilling.AccountsClientListOptions
		if expandParam != nil {
			options = &armbilling.AccountsClientListOptions{
				Expand: expandParam,
			}
		}
		resp := a.srv.NewListPager(options)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armbilling.AccountsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchNewListInvoiceSectionsByCreateSubscriptionPermissionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListInvoiceSectionsByCreateSubscriptionPermissionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListInvoiceSectionsByCreateSubscriptionPermissionPager not implemented")}
	}
	newListInvoiceSectionsByCreateSubscriptionPermissionPager := a.newListInvoiceSectionsByCreateSubscriptionPermissionPager.get(req)
	if newListInvoiceSectionsByCreateSubscriptionPermissionPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listInvoiceSectionsWithCreateSubscriptionPermission`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListInvoiceSectionsByCreateSubscriptionPermissionPager(billingAccountNameParam, nil)
		newListInvoiceSectionsByCreateSubscriptionPermissionPager = &resp
		a.newListInvoiceSectionsByCreateSubscriptionPermissionPager.add(req, newListInvoiceSectionsByCreateSubscriptionPermissionPager)
		server.PagerResponderInjectNextLinks(newListInvoiceSectionsByCreateSubscriptionPermissionPager, req, func(page *armbilling.AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListInvoiceSectionsByCreateSubscriptionPermissionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListInvoiceSectionsByCreateSubscriptionPermissionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListInvoiceSectionsByCreateSubscriptionPermissionPager) {
		a.newListInvoiceSectionsByCreateSubscriptionPermissionPager.remove(req)
	}
	return resp, nil
}

func (a *AccountsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := a.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armbilling.AccountUpdateRequest](req)
		if err != nil {
			return nil, err
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdate(req.Context(), billingAccountNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		a.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		a.beginUpdate.remove(req)
	}

	return resp, nil
}
