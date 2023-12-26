//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armserialconsole

// MicrosoftSerialConsoleClientDisableConsoleResponse contains the response from method MicrosoftSerialConsoleClient.DisableConsole.
type MicrosoftSerialConsoleClientDisableConsoleResponse struct {
	// Possible types are DisableSerialConsoleResult, GetSerialConsoleSubscriptionNotFound
	Value any
}

// MicrosoftSerialConsoleClientEnableConsoleResponse contains the response from method MicrosoftSerialConsoleClient.EnableConsole.
type MicrosoftSerialConsoleClientEnableConsoleResponse struct {
	// Possible types are EnableSerialConsoleResult, GetSerialConsoleSubscriptionNotFound
	Value any
}

// MicrosoftSerialConsoleClientGetConsoleStatusResponse contains the response from method MicrosoftSerialConsoleClient.GetConsoleStatus.
type MicrosoftSerialConsoleClientGetConsoleStatusResponse struct {
	// Possible types are Status, GetSerialConsoleSubscriptionNotFound
	Value any
}

// MicrosoftSerialConsoleClientListOperationsResponse contains the response from method MicrosoftSerialConsoleClient.ListOperations.
type MicrosoftSerialConsoleClientListOperationsResponse struct {
	// Serial Console operations
	Operations
}

// SerialPortsClientConnectResponse contains the response from method SerialPortsClient.Connect.
type SerialPortsClientConnectResponse struct {
	// Returns a connection string to the serial port of the resource.
	SerialPortConnectResult
}

// SerialPortsClientCreateResponse contains the response from method SerialPortsClient.Create.
type SerialPortsClientCreateResponse struct {
	// Represents the serial port of the parent resource.
	SerialPort
}

// SerialPortsClientDeleteResponse contains the response from method SerialPortsClient.Delete.
type SerialPortsClientDeleteResponse struct {
	// placeholder for future response values
}

// SerialPortsClientGetResponse contains the response from method SerialPortsClient.Get.
type SerialPortsClientGetResponse struct {
	// Represents the serial port of the parent resource.
	SerialPort
}

// SerialPortsClientListBySubscriptionsResponse contains the response from method SerialPortsClient.ListBySubscriptions.
type SerialPortsClientListBySubscriptionsResponse struct {
	// The list serial ports operation response.
	SerialPortListResult
}

// SerialPortsClientListResponse contains the response from method SerialPortsClient.List.
type SerialPortsClientListResponse struct {
	// The list serial ports operation response.
	SerialPortListResult
}
