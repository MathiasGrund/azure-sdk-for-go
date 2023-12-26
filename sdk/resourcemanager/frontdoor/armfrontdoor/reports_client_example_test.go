//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetLatencyScorecard.json
func ExampleReportsClient_GetLatencyScorecards() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReportsClient().GetLatencyScorecards(ctx, "MyResourceGroup", "MyProfile", "MyExperiment", armfrontdoor.LatencyScorecardAggregationIntervalDaily, &armfrontdoor.ReportsClientGetLatencyScorecardsOptions{EndDateTimeUTC: nil,
		Country: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LatencyScorecard = armfrontdoor.LatencyScorecard{
	// 	Name: to.Ptr("DailyLatencyScorecard"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/MyResourceGroup/providers/Microsoft.Network/NetworkExperimentProfiles/MyProfile/Experiments/MyExperiment/LatencyScorecard"),
	// 	Properties: &armfrontdoor.LatencyScorecardProperties{
	// 		Description: to.Ptr("This scorecard is the latency scorecard, aggregated over a day"),
	// 		Country: to.Ptr("USA"),
	// 		EndDateTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-21T17:32:28.000Z"); return t}()),
	// 		EndpointA: to.Ptr("https://endpointA.com"),
	// 		EndpointB: to.Ptr("https://endpoingB.com"),
	// 		StartDateTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-21T17:32:28.000Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetTimeseries.json
func ExampleReportsClient_GetTimeseries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReportsClient().GetTimeseries(ctx, "MyResourceGroup", "MyProfile", "MyExperiment", func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-21T17:32:28.000Z"); return t }(), func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-21T17:32:28.000Z"); return t }(), armfrontdoor.TimeseriesAggregationIntervalHourly, armfrontdoor.TimeseriesTypeMeasurementCounts, &armfrontdoor.ReportsClientGetTimeseriesOptions{Endpoint: nil,
		Country: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Timeseries = armfrontdoor.Timeseries{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/MyResourceGroup/providers/Microsoft.Network/NetworkExperimentProfiles/MyProfile/Experiments/MyExperiment/Timeseries"),
	// 	Properties: &armfrontdoor.TimeseriesProperties{
	// 		AggregationInterval: to.Ptr(armfrontdoor.AggregationIntervalHourly),
	// 		Country: to.Ptr("United States"),
	// 		EndDateTimeUTC: to.Ptr("2019-08-02"),
	// 		Endpoint: to.Ptr("https://endpointA.com"),
	// 		StartDateTimeUTC: to.Ptr("2019-07-29"),
	// 		TimeseriesData: []*armfrontdoor.TimeseriesDataPoint{
	// 			{
	// 				DateTimeUTC: to.Ptr("2019-07-22T17:32:28Z"),
	// 				Value: to.Ptr[float32](79),
	// 		}},
	// 		TimeseriesType: to.Ptr(armfrontdoor.TimeseriesTypeMeasurementCounts),
	// 	},
	// }
}
