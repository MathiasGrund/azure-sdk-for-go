//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBCollectionPartitionGetMetrics.json
func ExampleCollectionPartitionClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCollectionPartitionClient().NewListMetricsPager("rg1", "ddb1", "databaseRid", "collectionRid", "$filter=(name.value eq 'Max RUs Per Second') and timeGrain eq duration'PT1M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T23:58:55.2780000Z", nil)
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
		// page.PartitionMetricListResult = armcosmos.PartitionMetricListResult{
		// 	Value: []*armcosmos.PartitionMetric{
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Max RUs Per Second"),
		// 				Value: to.Ptr("Max RUs Per Second"),
		// 			},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-20T23:58:55.278Z"); return t}()),
		// 			MetricValues: []*armcosmos.MetricValue{
		// 				{
		// 					Maximum: to.Ptr[float64](5),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.278Z"); return t}()),
		// 				},
		// 				{
		// 					Maximum: to.Ptr[float64](5),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:54:55.278Z"); return t}()),
		// 				},
		// 				{
		// 					Maximum: to.Ptr[float64](5),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:55:55.278Z"); return t}()),
		// 				},
		// 				{
		// 					Maximum: to.Ptr[float64](5),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:56:55.278Z"); return t}()),
		// 				},
		// 				{
		// 					Maximum: to.Ptr[float64](5),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:57:55.278Z"); return t}()),
		// 				},
		// 				{
		// 					Maximum: to.Ptr[float64](5),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:58:55.278Z"); return t}()),
		// 			}},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.278Z"); return t}()),
		// 			TimeGrain: to.Ptr("PT1M"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeCount),
		// 			PartitionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			PartitionKeyRangeID: to.Ptr("0"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBCollectionPartitionGetUsages.json
func ExampleCollectionPartitionClient_NewListUsagesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCollectionPartitionClient().NewListUsagesPager("rg1", "ddb1", "databaseRid", "collectionRid", &armcosmos.CollectionPartitionClientListUsagesOptions{Filter: to.Ptr("$filter=name.value eq 'Partition Storage'")})
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
		// page.PartitionUsagesResult = armcosmos.PartitionUsagesResult{
		// 	Value: []*armcosmos.PartitionUsage{
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Storage"),
		// 				Value: to.Ptr("Storage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](10737418240),
		// 			QuotaPeriod: to.Ptr("P1D"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeBytes),
		// 			PartitionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			PartitionKeyRangeID: to.Ptr("0"),
		// 	}},
		// }
	}
}
