//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/OutboundFirewallRuleGet.json
func ExampleOutboundFirewallRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOutboundFirewallRulesClient().Get(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", "server.database.windows.net", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OutboundFirewallRule = armsql.OutboundFirewallRule{
	// 	Name: to.Ptr("server.database.windows.net"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/outboundFirewallRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/outboundFirewallRules/server.datbase.windows.net"),
	// 	Properties: &armsql.OutboundFirewallRuleProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/OutboundFirewallRuleCreate.json
func ExampleOutboundFirewallRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOutboundFirewallRulesClient().BeginCreateOrUpdate(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", "server.database.windows.net", armsql.OutboundFirewallRule{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OutboundFirewallRule = armsql.OutboundFirewallRule{
	// 	Name: to.Ptr("server.database.windows.net"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/outboundFirewallRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/outboundFirewallRules/server.datbase.windows.net"),
	// 	Properties: &armsql.OutboundFirewallRuleProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/OutboundFirewallRuleDelete.json
func ExampleOutboundFirewallRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOutboundFirewallRulesClient().BeginDelete(ctx, "sqlcrudtest-7398", "sqlcrudtest-6661", "server.database.windows.net", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/OutboundFirewallRuleList.json
func ExampleOutboundFirewallRulesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOutboundFirewallRulesClient().NewListByServerPager("sqlcrudtest-7398", "sqlcrudtest-4645", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OutboundFirewallRuleListResult = armsql.OutboundFirewallRuleListResult{
		// 	Value: []*armsql.OutboundFirewallRule{
		// 		{
		// 			Name: to.Ptr("server.database.windows.net"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/outboundFirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/outboundFirewallRules/server.datbase.windows.net"),
		// 			Properties: &armsql.OutboundFirewallRuleProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("storage.blob.windows.net"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/outboundFirewallRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/outboundFirewallRules/storage.blob.windows.net"),
		// 			Properties: &armsql.OutboundFirewallRuleProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}
