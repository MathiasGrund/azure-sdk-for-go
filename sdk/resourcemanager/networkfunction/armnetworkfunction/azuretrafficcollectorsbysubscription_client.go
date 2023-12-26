//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkfunction

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AzureTrafficCollectorsBySubscriptionClient contains the methods for the AzureTrafficCollectorsBySubscription group.
// Don't use this type directly, use NewAzureTrafficCollectorsBySubscriptionClient() instead.
type AzureTrafficCollectorsBySubscriptionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureTrafficCollectorsBySubscriptionClient creates a new instance of AzureTrafficCollectorsBySubscriptionClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureTrafficCollectorsBySubscriptionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureTrafficCollectorsBySubscriptionClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureTrafficCollectorsBySubscriptionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Return list of Azure Traffic Collectors in a subscription
//
// Generated from API version 2022-11-01
//   - options - AzureTrafficCollectorsBySubscriptionClientListOptions contains the optional parameters for the AzureTrafficCollectorsBySubscriptionClient.NewListPager
//     method.
func (client *AzureTrafficCollectorsBySubscriptionClient) NewListPager(options *AzureTrafficCollectorsBySubscriptionClientListOptions) *runtime.Pager[AzureTrafficCollectorsBySubscriptionClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureTrafficCollectorsBySubscriptionClientListResponse]{
		More: func(page AzureTrafficCollectorsBySubscriptionClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureTrafficCollectorsBySubscriptionClientListResponse) (AzureTrafficCollectorsBySubscriptionClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureTrafficCollectorsBySubscriptionClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AzureTrafficCollectorsBySubscriptionClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AzureTrafficCollectorsBySubscriptionClient) listCreateRequest(ctx context.Context, options *AzureTrafficCollectorsBySubscriptionClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkFunction/azureTrafficCollectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AzureTrafficCollectorsBySubscriptionClient) listHandleResponse(resp *http.Response) (AzureTrafficCollectorsBySubscriptionClientListResponse, error) {
	result := AzureTrafficCollectorsBySubscriptionClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureTrafficCollectorListResult); err != nil {
		return AzureTrafficCollectorsBySubscriptionClientListResponse{}, err
	}
	return result, nil
}
