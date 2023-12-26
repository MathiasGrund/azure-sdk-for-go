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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
	"net/http"
	"net/url"
	"regexp"
)

// TenantConfigurationServer is a fake server for instances of the armapimanagement.TenantConfigurationClient type.
type TenantConfigurationServer struct {
	// BeginDeploy is the fake for method TenantConfigurationClient.BeginDeploy
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDeploy func(ctx context.Context, resourceGroupName string, serviceName string, configurationName armapimanagement.ConfigurationIDName, parameters armapimanagement.DeployConfigurationParameters, options *armapimanagement.TenantConfigurationClientBeginDeployOptions) (resp azfake.PollerResponder[armapimanagement.TenantConfigurationClientDeployResponse], errResp azfake.ErrorResponder)

	// GetSyncState is the fake for method TenantConfigurationClient.GetSyncState
	// HTTP status codes to indicate success: http.StatusOK
	GetSyncState func(ctx context.Context, resourceGroupName string, serviceName string, configurationName armapimanagement.ConfigurationIDName, options *armapimanagement.TenantConfigurationClientGetSyncStateOptions) (resp azfake.Responder[armapimanagement.TenantConfigurationClientGetSyncStateResponse], errResp azfake.ErrorResponder)

	// BeginSave is the fake for method TenantConfigurationClient.BeginSave
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginSave func(ctx context.Context, resourceGroupName string, serviceName string, configurationName armapimanagement.ConfigurationIDName, parameters armapimanagement.SaveConfigurationParameter, options *armapimanagement.TenantConfigurationClientBeginSaveOptions) (resp azfake.PollerResponder[armapimanagement.TenantConfigurationClientSaveResponse], errResp azfake.ErrorResponder)

	// BeginValidate is the fake for method TenantConfigurationClient.BeginValidate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginValidate func(ctx context.Context, resourceGroupName string, serviceName string, configurationName armapimanagement.ConfigurationIDName, parameters armapimanagement.DeployConfigurationParameters, options *armapimanagement.TenantConfigurationClientBeginValidateOptions) (resp azfake.PollerResponder[armapimanagement.TenantConfigurationClientValidateResponse], errResp azfake.ErrorResponder)
}

// NewTenantConfigurationServerTransport creates a new instance of TenantConfigurationServerTransport with the provided implementation.
// The returned TenantConfigurationServerTransport instance is connected to an instance of armapimanagement.TenantConfigurationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTenantConfigurationServerTransport(srv *TenantConfigurationServer) *TenantConfigurationServerTransport {
	return &TenantConfigurationServerTransport{
		srv:           srv,
		beginDeploy:   newTracker[azfake.PollerResponder[armapimanagement.TenantConfigurationClientDeployResponse]](),
		beginSave:     newTracker[azfake.PollerResponder[armapimanagement.TenantConfigurationClientSaveResponse]](),
		beginValidate: newTracker[azfake.PollerResponder[armapimanagement.TenantConfigurationClientValidateResponse]](),
	}
}

// TenantConfigurationServerTransport connects instances of armapimanagement.TenantConfigurationClient to instances of TenantConfigurationServer.
// Don't use this type directly, use NewTenantConfigurationServerTransport instead.
type TenantConfigurationServerTransport struct {
	srv           *TenantConfigurationServer
	beginDeploy   *tracker[azfake.PollerResponder[armapimanagement.TenantConfigurationClientDeployResponse]]
	beginSave     *tracker[azfake.PollerResponder[armapimanagement.TenantConfigurationClientSaveResponse]]
	beginValidate *tracker[azfake.PollerResponder[armapimanagement.TenantConfigurationClientValidateResponse]]
}

// Do implements the policy.Transporter interface for TenantConfigurationServerTransport.
func (t *TenantConfigurationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TenantConfigurationClient.BeginDeploy":
		resp, err = t.dispatchBeginDeploy(req)
	case "TenantConfigurationClient.GetSyncState":
		resp, err = t.dispatchGetSyncState(req)
	case "TenantConfigurationClient.BeginSave":
		resp, err = t.dispatchBeginSave(req)
	case "TenantConfigurationClient.BeginValidate":
		resp, err = t.dispatchBeginValidate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TenantConfigurationServerTransport) dispatchBeginDeploy(req *http.Request) (*http.Response, error) {
	if t.srv.BeginDeploy == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeploy not implemented")}
	}
	beginDeploy := t.beginDeploy.get(req)
	if beginDeploy == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tenant/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deploy`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.DeployConfigurationParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		configurationNameParam, err := parseWithCast(matches[regex.SubexpIndex("configurationName")], func(v string) (armapimanagement.ConfigurationIDName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armapimanagement.ConfigurationIDName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginDeploy(req.Context(), resourceGroupNameParam, serviceNameParam, configurationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeploy = &respr
		t.beginDeploy.add(req, beginDeploy)
	}

	resp, err := server.PollerResponderNext(beginDeploy, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		t.beginDeploy.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeploy) {
		t.beginDeploy.remove(req)
	}

	return resp, nil
}

func (t *TenantConfigurationServerTransport) dispatchGetSyncState(req *http.Request) (*http.Response, error) {
	if t.srv.GetSyncState == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSyncState not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tenant/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncState`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	configurationNameParam, err := parseWithCast(matches[regex.SubexpIndex("configurationName")], func(v string) (armapimanagement.ConfigurationIDName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.ConfigurationIDName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.GetSyncState(req.Context(), resourceGroupNameParam, serviceNameParam, configurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TenantConfigurationSyncStateContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TenantConfigurationServerTransport) dispatchBeginSave(req *http.Request) (*http.Response, error) {
	if t.srv.BeginSave == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSave not implemented")}
	}
	beginSave := t.beginSave.get(req)
	if beginSave == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tenant/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/save`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.SaveConfigurationParameter](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		configurationNameParam, err := parseWithCast(matches[regex.SubexpIndex("configurationName")], func(v string) (armapimanagement.ConfigurationIDName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armapimanagement.ConfigurationIDName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginSave(req.Context(), resourceGroupNameParam, serviceNameParam, configurationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSave = &respr
		t.beginSave.add(req, beginSave)
	}

	resp, err := server.PollerResponderNext(beginSave, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		t.beginSave.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSave) {
		t.beginSave.remove(req)
	}

	return resp, nil
}

func (t *TenantConfigurationServerTransport) dispatchBeginValidate(req *http.Request) (*http.Response, error) {
	if t.srv.BeginValidate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginValidate not implemented")}
	}
	beginValidate := t.beginValidate.get(req)
	if beginValidate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tenant/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validate`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.DeployConfigurationParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		configurationNameParam, err := parseWithCast(matches[regex.SubexpIndex("configurationName")], func(v string) (armapimanagement.ConfigurationIDName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armapimanagement.ConfigurationIDName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginValidate(req.Context(), resourceGroupNameParam, serviceNameParam, configurationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginValidate = &respr
		t.beginValidate.add(req, beginValidate)
	}

	resp, err := server.PollerResponderNext(beginValidate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		t.beginValidate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginValidate) {
		t.beginValidate.remove(req)
	}

	return resp, nil
}
