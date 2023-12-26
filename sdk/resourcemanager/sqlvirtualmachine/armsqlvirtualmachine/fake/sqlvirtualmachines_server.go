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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
	"net/http"
	"net/url"
	"regexp"
)

// SQLVirtualMachinesServer is a fake server for instances of the armsqlvirtualmachine.SQLVirtualMachinesClient type.
type SQLVirtualMachinesServer struct {
	// BeginCreateOrUpdate is the fake for method SQLVirtualMachinesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters armsqlvirtualmachine.SQLVirtualMachine, options *armsqlvirtualmachine.SQLVirtualMachinesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SQLVirtualMachinesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *armsqlvirtualmachine.SQLVirtualMachinesClientBeginDeleteOptions) (resp azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SQLVirtualMachinesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *armsqlvirtualmachine.SQLVirtualMachinesClientGetOptions) (resp azfake.Responder[armsqlvirtualmachine.SQLVirtualMachinesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SQLVirtualMachinesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsqlvirtualmachine.SQLVirtualMachinesClientListOptions) (resp azfake.PagerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientListResponse])

	// NewListByResourceGroupPager is the fake for method SQLVirtualMachinesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armsqlvirtualmachine.SQLVirtualMachinesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientListByResourceGroupResponse])

	// NewListBySQLVMGroupPager is the fake for method SQLVirtualMachinesClient.NewListBySQLVMGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySQLVMGroupPager func(resourceGroupName string, sqlVirtualMachineGroupName string, options *armsqlvirtualmachine.SQLVirtualMachinesClientListBySQLVMGroupOptions) (resp azfake.PagerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientListBySQLVMGroupResponse])

	// BeginRedeploy is the fake for method SQLVirtualMachinesClient.BeginRedeploy
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRedeploy func(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *armsqlvirtualmachine.SQLVirtualMachinesClientBeginRedeployOptions) (resp azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientRedeployResponse], errResp azfake.ErrorResponder)

	// BeginStartAssessment is the fake for method SQLVirtualMachinesClient.BeginStartAssessment
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStartAssessment func(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, options *armsqlvirtualmachine.SQLVirtualMachinesClientBeginStartAssessmentOptions) (resp azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientStartAssessmentResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method SQLVirtualMachinesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, sqlVirtualMachineName string, parameters armsqlvirtualmachine.Update, options *armsqlvirtualmachine.SQLVirtualMachinesClientBeginUpdateOptions) (resp azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewSQLVirtualMachinesServerTransport creates a new instance of SQLVirtualMachinesServerTransport with the provided implementation.
// The returned SQLVirtualMachinesServerTransport instance is connected to an instance of armsqlvirtualmachine.SQLVirtualMachinesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLVirtualMachinesServerTransport(srv *SQLVirtualMachinesServer) *SQLVirtualMachinesServerTransport {
	return &SQLVirtualMachinesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientListByResourceGroupResponse]](),
		newListBySQLVMGroupPager:    newTracker[azfake.PagerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientListBySQLVMGroupResponse]](),
		beginRedeploy:               newTracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientRedeployResponse]](),
		beginStartAssessment:        newTracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientStartAssessmentResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientUpdateResponse]](),
	}
}

// SQLVirtualMachinesServerTransport connects instances of armsqlvirtualmachine.SQLVirtualMachinesClient to instances of SQLVirtualMachinesServer.
// Don't use this type directly, use NewSQLVirtualMachinesServerTransport instead.
type SQLVirtualMachinesServerTransport struct {
	srv                         *SQLVirtualMachinesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientListByResourceGroupResponse]]
	newListBySQLVMGroupPager    *tracker[azfake.PagerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientListBySQLVMGroupResponse]]
	beginRedeploy               *tracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientRedeployResponse]]
	beginStartAssessment        *tracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientStartAssessmentResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armsqlvirtualmachine.SQLVirtualMachinesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for SQLVirtualMachinesServerTransport.
func (s *SQLVirtualMachinesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SQLVirtualMachinesClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "SQLVirtualMachinesClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SQLVirtualMachinesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SQLVirtualMachinesClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "SQLVirtualMachinesClient.NewListByResourceGroupPager":
		resp, err = s.dispatchNewListByResourceGroupPager(req)
	case "SQLVirtualMachinesClient.NewListBySQLVMGroupPager":
		resp, err = s.dispatchNewListBySQLVMGroupPager(req)
	case "SQLVirtualMachinesClient.BeginRedeploy":
		resp, err = s.dispatchBeginRedeploy(req)
	case "SQLVirtualMachinesClient.BeginStartAssessment":
		resp, err = s.dispatchBeginStartAssessment(req)
	case "SQLVirtualMachinesClient.BeginUpdate":
		resp, err = s.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SQLVirtualMachinesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SqlVirtualMachine/sqlVirtualMachines/(?P<sqlVirtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsqlvirtualmachine.SQLVirtualMachine](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlVirtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlVirtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, sqlVirtualMachineNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *SQLVirtualMachinesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SqlVirtualMachine/sqlVirtualMachines/(?P<sqlVirtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlVirtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlVirtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, sqlVirtualMachineNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SQLVirtualMachinesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SqlVirtualMachine/sqlVirtualMachines/(?P<sqlVirtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	sqlVirtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlVirtualMachineName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armsqlvirtualmachine.SQLVirtualMachinesClientGetOptions
	if expandParam != nil {
		options = &armsqlvirtualmachine.SQLVirtualMachinesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, sqlVirtualMachineNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SQLVirtualMachine, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQLVirtualMachinesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SqlVirtualMachine/sqlVirtualMachines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListPager(nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsqlvirtualmachine.SQLVirtualMachinesClientListResponse, createLink func() string) {
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

func (s *SQLVirtualMachinesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := s.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SqlVirtualMachine/sqlVirtualMachines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		s.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armsqlvirtualmachine.SQLVirtualMachinesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		s.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (s *SQLVirtualMachinesServerTransport) dispatchNewListBySQLVMGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySQLVMGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySQLVMGroupPager not implemented")}
	}
	newListBySQLVMGroupPager := s.newListBySQLVMGroupPager.get(req)
	if newListBySQLVMGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SqlVirtualMachine/sqlVirtualMachineGroups/(?P<sqlVirtualMachineGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlVirtualMachines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlVirtualMachineGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlVirtualMachineGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListBySQLVMGroupPager(resourceGroupNameParam, sqlVirtualMachineGroupNameParam, nil)
		newListBySQLVMGroupPager = &resp
		s.newListBySQLVMGroupPager.add(req, newListBySQLVMGroupPager)
		server.PagerResponderInjectNextLinks(newListBySQLVMGroupPager, req, func(page *armsqlvirtualmachine.SQLVirtualMachinesClientListBySQLVMGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySQLVMGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListBySQLVMGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySQLVMGroupPager) {
		s.newListBySQLVMGroupPager.remove(req)
	}
	return resp, nil
}

func (s *SQLVirtualMachinesServerTransport) dispatchBeginRedeploy(req *http.Request) (*http.Response, error) {
	if s.srv.BeginRedeploy == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRedeploy not implemented")}
	}
	beginRedeploy := s.beginRedeploy.get(req)
	if beginRedeploy == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SqlVirtualMachine/sqlVirtualMachines/(?P<sqlVirtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/redeploy`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlVirtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlVirtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginRedeploy(req.Context(), resourceGroupNameParam, sqlVirtualMachineNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRedeploy = &respr
		s.beginRedeploy.add(req, beginRedeploy)
	}

	resp, err := server.PollerResponderNext(beginRedeploy, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginRedeploy.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRedeploy) {
		s.beginRedeploy.remove(req)
	}

	return resp, nil
}

func (s *SQLVirtualMachinesServerTransport) dispatchBeginStartAssessment(req *http.Request) (*http.Response, error) {
	if s.srv.BeginStartAssessment == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStartAssessment not implemented")}
	}
	beginStartAssessment := s.beginStartAssessment.get(req)
	if beginStartAssessment == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SqlVirtualMachine/sqlVirtualMachines/(?P<sqlVirtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/startAssessment`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlVirtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlVirtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginStartAssessment(req.Context(), resourceGroupNameParam, sqlVirtualMachineNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStartAssessment = &respr
		s.beginStartAssessment.add(req, beginStartAssessment)
	}

	resp, err := server.PollerResponderNext(beginStartAssessment, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginStartAssessment.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStartAssessment) {
		s.beginStartAssessment.remove(req)
	}

	return resp, nil
}

func (s *SQLVirtualMachinesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := s.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SqlVirtualMachine/sqlVirtualMachines/(?P<sqlVirtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsqlvirtualmachine.Update](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlVirtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlVirtualMachineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUpdate(req.Context(), resourceGroupNameParam, sqlVirtualMachineNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		s.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		s.beginUpdate.remove(req)
	}

	return resp, nil
}
