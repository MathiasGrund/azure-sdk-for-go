//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armstorage_test

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/fake"
)

func ExampleAccountsServer() {
	// first, create an instance of the fake server for the client you wish to test.
	// the type name of the server will be similar to the corresponding client, with
	// the suffix "Server" instead of "Client".
	fakeAccountsServer := fake.AccountsServer{

		// next, provide implementations for the APIs you wish to fake.
		// this fake corresponds to the AccountsClient.GetProperties() API.
		GetProperties: func(ctx context.Context, resourceGroupName string, accountName string, options *armstorage.AccountsClientGetPropertiesOptions) (resp azfake.Responder[armstorage.AccountsClientGetPropertiesResponse], errResp azfake.ErrorResponder) {
			// the values of ctx, resourceGroupName, accountName, and options come from the API call.

			// the named return values resp and errResp are used to construct the response
			// and are meant to be mutually exclusive. if both responses have been constructed,
			// the error response is selected.

			// construct the response type, populating fields as required
			accountResp := armstorage.AccountsClientGetPropertiesResponse{}
			accountResp.ID = to.Ptr("/fake/resource/id")

			// use resp to set the desired response
			resp.SetResponse(http.StatusOK, accountResp, nil)

			// to simulate the failure case, use errResp
			//errResp.SetResponseError(http.StatusBadRequest, "ThisIsASimulatedError")

			return
		},
	}

	// now create the corresponding client, connecting the fake server via the client options
	client, err := armstorage.NewAccountsClient("subscriptionID", &azfake.TokenCredential{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewAccountsServerTransport(&fakeAccountsServer),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// call the API. the provided values will be passed to the fake's implementation.
	// the response or error values returned by the API call are from the fake.
	resp, err := client.GetProperties(context.TODO(), "fakeResourceGroup", "fakeVM", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*resp.ID)

	// APIs that haven't been faked will return an error
	_, err = client.CheckNameAvailability(context.TODO(), armstorage.AccountCheckNameAvailabilityParameters{}, nil)

	fmt.Println(err.Error())

	// Output:
	// /fake/resource/id
	// fake for method CheckNameAvailability not implemented
}
