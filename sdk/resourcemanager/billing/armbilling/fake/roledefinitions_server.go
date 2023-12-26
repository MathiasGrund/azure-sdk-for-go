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

// RoleDefinitionsServer is a fake server for instances of the armbilling.RoleDefinitionsClient type.
type RoleDefinitionsServer struct {
	// GetByBillingAccount is the fake for method RoleDefinitionsClient.GetByBillingAccount
	// HTTP status codes to indicate success: http.StatusOK
	GetByBillingAccount func(ctx context.Context, billingAccountName string, billingRoleDefinitionName string, options *armbilling.RoleDefinitionsClientGetByBillingAccountOptions) (resp azfake.Responder[armbilling.RoleDefinitionsClientGetByBillingAccountResponse], errResp azfake.ErrorResponder)

	// GetByBillingProfile is the fake for method RoleDefinitionsClient.GetByBillingProfile
	// HTTP status codes to indicate success: http.StatusOK
	GetByBillingProfile func(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string, options *armbilling.RoleDefinitionsClientGetByBillingProfileOptions) (resp azfake.Responder[armbilling.RoleDefinitionsClientGetByBillingProfileResponse], errResp azfake.ErrorResponder)

	// GetByInvoiceSection is the fake for method RoleDefinitionsClient.GetByInvoiceSection
	// HTTP status codes to indicate success: http.StatusOK
	GetByInvoiceSection func(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleDefinitionName string, options *armbilling.RoleDefinitionsClientGetByInvoiceSectionOptions) (resp azfake.Responder[armbilling.RoleDefinitionsClientGetByInvoiceSectionResponse], errResp azfake.ErrorResponder)

	// NewListByBillingAccountPager is the fake for method RoleDefinitionsClient.NewListByBillingAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingAccountPager func(billingAccountName string, options *armbilling.RoleDefinitionsClientListByBillingAccountOptions) (resp azfake.PagerResponder[armbilling.RoleDefinitionsClientListByBillingAccountResponse])

	// NewListByBillingProfilePager is the fake for method RoleDefinitionsClient.NewListByBillingProfilePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBillingProfilePager func(billingAccountName string, billingProfileName string, options *armbilling.RoleDefinitionsClientListByBillingProfileOptions) (resp azfake.PagerResponder[armbilling.RoleDefinitionsClientListByBillingProfileResponse])

	// NewListByInvoiceSectionPager is the fake for method RoleDefinitionsClient.NewListByInvoiceSectionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByInvoiceSectionPager func(billingAccountName string, billingProfileName string, invoiceSectionName string, options *armbilling.RoleDefinitionsClientListByInvoiceSectionOptions) (resp azfake.PagerResponder[armbilling.RoleDefinitionsClientListByInvoiceSectionResponse])
}

// NewRoleDefinitionsServerTransport creates a new instance of RoleDefinitionsServerTransport with the provided implementation.
// The returned RoleDefinitionsServerTransport instance is connected to an instance of armbilling.RoleDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRoleDefinitionsServerTransport(srv *RoleDefinitionsServer) *RoleDefinitionsServerTransport {
	return &RoleDefinitionsServerTransport{
		srv:                          srv,
		newListByBillingAccountPager: newTracker[azfake.PagerResponder[armbilling.RoleDefinitionsClientListByBillingAccountResponse]](),
		newListByBillingProfilePager: newTracker[azfake.PagerResponder[armbilling.RoleDefinitionsClientListByBillingProfileResponse]](),
		newListByInvoiceSectionPager: newTracker[azfake.PagerResponder[armbilling.RoleDefinitionsClientListByInvoiceSectionResponse]](),
	}
}

// RoleDefinitionsServerTransport connects instances of armbilling.RoleDefinitionsClient to instances of RoleDefinitionsServer.
// Don't use this type directly, use NewRoleDefinitionsServerTransport instead.
type RoleDefinitionsServerTransport struct {
	srv                          *RoleDefinitionsServer
	newListByBillingAccountPager *tracker[azfake.PagerResponder[armbilling.RoleDefinitionsClientListByBillingAccountResponse]]
	newListByBillingProfilePager *tracker[azfake.PagerResponder[armbilling.RoleDefinitionsClientListByBillingProfileResponse]]
	newListByInvoiceSectionPager *tracker[azfake.PagerResponder[armbilling.RoleDefinitionsClientListByInvoiceSectionResponse]]
}

// Do implements the policy.Transporter interface for RoleDefinitionsServerTransport.
func (r *RoleDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RoleDefinitionsClient.GetByBillingAccount":
		resp, err = r.dispatchGetByBillingAccount(req)
	case "RoleDefinitionsClient.GetByBillingProfile":
		resp, err = r.dispatchGetByBillingProfile(req)
	case "RoleDefinitionsClient.GetByInvoiceSection":
		resp, err = r.dispatchGetByInvoiceSection(req)
	case "RoleDefinitionsClient.NewListByBillingAccountPager":
		resp, err = r.dispatchNewListByBillingAccountPager(req)
	case "RoleDefinitionsClient.NewListByBillingProfilePager":
		resp, err = r.dispatchNewListByBillingProfilePager(req)
	case "RoleDefinitionsClient.NewListByInvoiceSectionPager":
		resp, err = r.dispatchNewListByInvoiceSectionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RoleDefinitionsServerTransport) dispatchGetByBillingAccount(req *http.Request) (*http.Response, error) {
	if r.srv.GetByBillingAccount == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByBillingAccount not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingRoleDefinitions/(?P<billingRoleDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	billingRoleDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingRoleDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.GetByBillingAccount(req.Context(), billingAccountNameParam, billingRoleDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleDefinitionsServerTransport) dispatchGetByBillingProfile(req *http.Request) (*http.Response, error) {
	if r.srv.GetByBillingProfile == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByBillingProfile not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingRoleDefinitions/(?P<billingRoleDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
	if err != nil {
		return nil, err
	}
	billingRoleDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingRoleDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.GetByBillingProfile(req.Context(), billingAccountNameParam, billingProfileNameParam, billingRoleDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleDefinitionsServerTransport) dispatchGetByInvoiceSection(req *http.Request) (*http.Response, error) {
	if r.srv.GetByInvoiceSection == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByInvoiceSection not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/invoiceSections/(?P<invoiceSectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingRoleDefinitions/(?P<billingRoleDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
	if err != nil {
		return nil, err
	}
	invoiceSectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("invoiceSectionName")])
	if err != nil {
		return nil, err
	}
	billingRoleDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingRoleDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.GetByInvoiceSection(req.Context(), billingAccountNameParam, billingProfileNameParam, invoiceSectionNameParam, billingRoleDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RoleDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RoleDefinitionsServerTransport) dispatchNewListByBillingAccountPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByBillingAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingAccountPager not implemented")}
	}
	newListByBillingAccountPager := r.newListByBillingAccountPager.get(req)
	if newListByBillingAccountPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingRoleDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByBillingAccountPager(billingAccountNameParam, nil)
		newListByBillingAccountPager = &resp
		r.newListByBillingAccountPager.add(req, newListByBillingAccountPager)
		server.PagerResponderInjectNextLinks(newListByBillingAccountPager, req, func(page *armbilling.RoleDefinitionsClientListByBillingAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByBillingAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingAccountPager) {
		r.newListByBillingAccountPager.remove(req)
	}
	return resp, nil
}

func (r *RoleDefinitionsServerTransport) dispatchNewListByBillingProfilePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByBillingProfilePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBillingProfilePager not implemented")}
	}
	newListByBillingProfilePager := r.newListByBillingProfilePager.get(req)
	if newListByBillingProfilePager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingRoleDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByBillingProfilePager(billingAccountNameParam, billingProfileNameParam, nil)
		newListByBillingProfilePager = &resp
		r.newListByBillingProfilePager.add(req, newListByBillingProfilePager)
		server.PagerResponderInjectNextLinks(newListByBillingProfilePager, req, func(page *armbilling.RoleDefinitionsClientListByBillingProfileResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBillingProfilePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByBillingProfilePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBillingProfilePager) {
		r.newListByBillingProfilePager.remove(req)
	}
	return resp, nil
}

func (r *RoleDefinitionsServerTransport) dispatchNewListByInvoiceSectionPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByInvoiceSectionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByInvoiceSectionPager not implemented")}
	}
	newListByInvoiceSectionPager := r.newListByInvoiceSectionPager.get(req)
	if newListByInvoiceSectionPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/invoiceSections/(?P<invoiceSectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingRoleDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		invoiceSectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("invoiceSectionName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByInvoiceSectionPager(billingAccountNameParam, billingProfileNameParam, invoiceSectionNameParam, nil)
		newListByInvoiceSectionPager = &resp
		r.newListByInvoiceSectionPager.add(req, newListByInvoiceSectionPager)
		server.PagerResponderInjectNextLinks(newListByInvoiceSectionPager, req, func(page *armbilling.RoleDefinitionsClientListByInvoiceSectionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByInvoiceSectionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByInvoiceSectionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByInvoiceSectionPager) {
		r.newListByInvoiceSectionPager.remove(req)
	}
	return resp, nil
}
