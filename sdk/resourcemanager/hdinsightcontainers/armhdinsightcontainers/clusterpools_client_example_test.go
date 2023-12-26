//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7f70e351393addbc31d790a908c994c7c8644d9c/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/GetClusterPool.json
func ExampleClusterPoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClusterPoolsClient().Get(ctx, "hiloResourcegroup", "clusterpool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterPool = armhdinsightcontainers.ClusterPool{
	// 	Name: to.Ptr("clusterpool1"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusterPools"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1"),
	// 	SystemData: &armhdinsightcontainers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("West US 2"),
	// 	Properties: &armhdinsightcontainers.ClusterPoolResourceProperties{
	// 		AksClusterProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesAksClusterProfile{
	// 			AksClusterAgentPoolIdentityProfile: &armhdinsightcontainers.AksClusterProfileAksClusterAgentPoolIdentityProfile{
	// 				MsiClientID: to.Ptr("a89fb478-2a84-4d9b-8f18-3e8c4d1db3eb"),
	// 				MsiObjectID: to.Ptr("dc7ef861-8b55-4ffb-9003-20885cd895a9"),
	// 				MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ManagedIdentity/userAssignedIdentities/clusterpool1-agentpool"),
	// 			},
	// 			AksClusterResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ContainerService/managedClusters/clusterpool1"),
	// 			AksVersion: to.Ptr("1.24"),
	// 		},
	// 		ClusterPoolProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesClusterPoolProfile{
	// 			ClusterPoolVersion: to.Ptr("1.2"),
	// 		},
	// 		ComputeProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesComputeProfile{
	// 			Count: to.Ptr[int32](3),
	// 			VMSize: to.Ptr("Standard_D3_v2"),
	// 		},
	// 		DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
	// 		ManagedResourceGroupName: to.Ptr("hdi-45cd32aead6e4a91b079a0cdbfac8c36"),
	// 		ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7f70e351393addbc31d790a908c994c7c8644d9c/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/CreateClusterPool.json
func ExampleClusterPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterPoolsClient().BeginCreateOrUpdate(ctx, "hiloResourcegroup", "clusterpool1", armhdinsightcontainers.ClusterPool{
		Location: to.Ptr("West US 2"),
		Properties: &armhdinsightcontainers.ClusterPoolResourceProperties{
			ClusterPoolProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesClusterPoolProfile{
				ClusterPoolVersion: to.Ptr("1.2"),
			},
			ComputeProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesComputeProfile{
				VMSize: to.Ptr("Standard_D3_v2"),
			},
		},
	}, nil)
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
	// res.ClusterPool = armhdinsightcontainers.ClusterPool{
	// 	Name: to.Ptr("clusterpool1"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusterPools"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1"),
	// 	SystemData: &armhdinsightcontainers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("West US 2"),
	// 	Properties: &armhdinsightcontainers.ClusterPoolResourceProperties{
	// 		AksClusterProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesAksClusterProfile{
	// 			AksClusterAgentPoolIdentityProfile: &armhdinsightcontainers.AksClusterProfileAksClusterAgentPoolIdentityProfile{
	// 				MsiClientID: to.Ptr("a89fb478-2a84-4d9b-8f18-3e8c4d1db3eb"),
	// 				MsiObjectID: to.Ptr("dc7ef861-8b55-4ffb-9003-20885cd895a9"),
	// 				MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ManagedIdentity/userAssignedIdentities/clusterpool1-agentpool"),
	// 			},
	// 			AksClusterResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ContainerService/managedClusters/clusterpool1"),
	// 			AksVersion: to.Ptr("1.24"),
	// 		},
	// 		ClusterPoolProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesClusterPoolProfile{
	// 			ClusterPoolVersion: to.Ptr("1.2"),
	// 		},
	// 		ComputeProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesComputeProfile{
	// 			Count: to.Ptr[int32](3),
	// 			VMSize: to.Ptr("Standard_D3_v2"),
	// 		},
	// 		DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
	// 		ManagedResourceGroupName: to.Ptr("hdi-45cd32aead6e4a91b079a0cdbfac8c36"),
	// 		ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7f70e351393addbc31d790a908c994c7c8644d9c/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/PatchClusterPool.json
func ExampleClusterPoolsClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterPoolsClient().BeginUpdateTags(ctx, "hiloResourcegroup", "clusterpool1", armhdinsightcontainers.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
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
	// res.ClusterPool = armhdinsightcontainers.ClusterPool{
	// 	Name: to.Ptr("clusterpool1"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusterPools"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1"),
	// 	SystemData: &armhdinsightcontainers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("West US 2"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armhdinsightcontainers.ClusterPoolResourceProperties{
	// 		AksClusterProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesAksClusterProfile{
	// 			AksClusterAgentPoolIdentityProfile: &armhdinsightcontainers.AksClusterProfileAksClusterAgentPoolIdentityProfile{
	// 				MsiClientID: to.Ptr("a89fb478-2a84-4d9b-8f18-3e8c4d1db3eb"),
	// 				MsiObjectID: to.Ptr("dc7ef861-8b55-4ffb-9003-20885cd895a9"),
	// 				MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ManagedIdentity/userAssignedIdentities/clusterpool1-agentpool"),
	// 			},
	// 			AksClusterResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ContainerService/managedClusters/clusterpool1"),
	// 		},
	// 		ComputeProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesComputeProfile{
	// 			Count: to.Ptr[int32](3),
	// 			VMSize: to.Ptr("Standard_D3_v2"),
	// 		},
	// 		DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
	// 		ManagedResourceGroupName: to.Ptr("hdi-45cd32aead6e4a91b079a0cdbfac8c36"),
	// 		ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7f70e351393addbc31d790a908c994c7c8644d9c/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/DeleteClusterPool.json
func ExampleClusterPoolsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterPoolsClient().BeginDelete(ctx, "rg1", "clusterpool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7f70e351393addbc31d790a908c994c7c8644d9c/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/ListClusterPoolsSubscription.json
func ExampleClusterPoolsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterPoolsClient().NewListBySubscriptionPager(nil)
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
		// page.ClusterPoolListResult = armhdinsightcontainers.ClusterPoolListResult{
		// 	Value: []*armhdinsightcontainers.ClusterPool{
		// 		{
		// 			Name: to.Ptr("clusterpool1"),
		// 			Type: to.Ptr("Microsoft.HDInsight/clusterPools"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1"),
		// 			Location: to.Ptr("West US 2"),
		// 			Tags: map[string]*string{
		// 				"company": to.Ptr("Contoso"),
		// 				"department": to.Ptr("MightyMight"),
		// 			},
		// 			Properties: &armhdinsightcontainers.ClusterPoolResourceProperties{
		// 				AksClusterProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesAksClusterProfile{
		// 					AksClusterAgentPoolIdentityProfile: &armhdinsightcontainers.AksClusterProfileAksClusterAgentPoolIdentityProfile{
		// 						MsiClientID: to.Ptr("a89fb478-2a84-4d9b-8f18-3e8c4d1db3eb"),
		// 						MsiObjectID: to.Ptr("dc7ef861-8b55-4ffb-9003-20885cd895a9"),
		// 						MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ManagedIdentity/userAssignedIdentities/clusterpool1-agentpool"),
		// 					},
		// 					AksClusterResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ContainerService/managedClusters/clusterpool1"),
		// 					AksVersion: to.Ptr("1.24"),
		// 				},
		// 				ClusterPoolProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesClusterPoolProfile{
		// 					ClusterPoolVersion: to.Ptr("1.2"),
		// 				},
		// 				ComputeProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesComputeProfile{
		// 					Count: to.Ptr[int32](3),
		// 					VMSize: to.Ptr("Standard_D3_v2"),
		// 				},
		// 				DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
		// 				ManagedResourceGroupName: to.Ptr("hdi-45cd32aead6e4a91b079a0cdbfac8c36"),
		// 				ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7f70e351393addbc31d790a908c994c7c8644d9c/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/ListClusterPools.json
func ExampleClusterPoolsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterPoolsClient().NewListByResourceGroupPager("hiloResourcegroup", nil)
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
		// page.ClusterPoolListResult = armhdinsightcontainers.ClusterPoolListResult{
		// 	Value: []*armhdinsightcontainers.ClusterPool{
		// 		{
		// 			Name: to.Ptr("clusterpool1"),
		// 			Type: to.Ptr("Microsoft.HDInsight/clusterPools"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1"),
		// 			Location: to.Ptr("West US 2"),
		// 			Tags: map[string]*string{
		// 				"company": to.Ptr("Contoso"),
		// 				"department": to.Ptr("MightyMight"),
		// 			},
		// 			Properties: &armhdinsightcontainers.ClusterPoolResourceProperties{
		// 				AksClusterProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesAksClusterProfile{
		// 					AksClusterAgentPoolIdentityProfile: &armhdinsightcontainers.AksClusterProfileAksClusterAgentPoolIdentityProfile{
		// 						MsiClientID: to.Ptr("a89fb478-2a84-4d9b-8f18-3e8c4d1db3eb"),
		// 						MsiObjectID: to.Ptr("dc7ef861-8b55-4ffb-9003-20885cd895a9"),
		// 						MsiResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ManagedIdentity/userAssignedIdentities/clusterpool1-agentpool"),
		// 					},
		// 					AksClusterResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hdi-45cd32aead6e4a91b079a0cdbfac8c36/providers/Microsoft.ContainerService/managedClusters/clusterpool1"),
		// 					AksVersion: to.Ptr("1.24"),
		// 				},
		// 				ClusterPoolProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesClusterPoolProfile{
		// 					ClusterPoolVersion: to.Ptr("1.2"),
		// 				},
		// 				ComputeProfile: &armhdinsightcontainers.ClusterPoolResourcePropertiesComputeProfile{
		// 					Count: to.Ptr[int32](3),
		// 					VMSize: to.Ptr("Standard_D3_v2"),
		// 				},
		// 				DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
		// 				ManagedResourceGroupName: to.Ptr("hdi-45cd32aead6e4a91b079a0cdbfac8c36"),
		// 				ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}
