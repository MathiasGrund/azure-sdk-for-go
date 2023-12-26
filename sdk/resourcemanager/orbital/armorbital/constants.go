//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital"
	moduleVersion = "v2.1.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AutoTrackingConfiguration - Auto-tracking configuration.
type AutoTrackingConfiguration string

const (
	AutoTrackingConfigurationDisabled AutoTrackingConfiguration = "disabled"
	AutoTrackingConfigurationSBand    AutoTrackingConfiguration = "sBand"
	AutoTrackingConfigurationXBand    AutoTrackingConfiguration = "xBand"
)

// PossibleAutoTrackingConfigurationValues returns the possible values for the AutoTrackingConfiguration const type.
func PossibleAutoTrackingConfigurationValues() []AutoTrackingConfiguration {
	return []AutoTrackingConfiguration{
		AutoTrackingConfigurationDisabled,
		AutoTrackingConfigurationSBand,
		AutoTrackingConfigurationXBand,
	}
}

// Capability - Capability of the Ground Station.
type Capability string

const (
	CapabilityCommunication    Capability = "Communication"
	CapabilityEarthObservation Capability = "EarthObservation"
)

// PossibleCapabilityValues returns the possible values for the Capability const type.
func PossibleCapabilityValues() []Capability {
	return []Capability{
		CapabilityCommunication,
		CapabilityEarthObservation,
	}
}

type CapabilityParameter string

const (
	CapabilityParameterCommunication    CapabilityParameter = "Communication"
	CapabilityParameterEarthObservation CapabilityParameter = "EarthObservation"
)

// PossibleCapabilityParameterValues returns the possible values for the CapabilityParameter const type.
func PossibleCapabilityParameterValues() []CapabilityParameter {
	return []CapabilityParameter{
		CapabilityParameterCommunication,
		CapabilityParameterEarthObservation,
	}
}

// ContactProfilesPropertiesProvisioningState - The current state of the resource's creation, deletion, or modification.
type ContactProfilesPropertiesProvisioningState string

const (
	ContactProfilesPropertiesProvisioningStateCanceled  ContactProfilesPropertiesProvisioningState = "canceled"
	ContactProfilesPropertiesProvisioningStateCreating  ContactProfilesPropertiesProvisioningState = "creating"
	ContactProfilesPropertiesProvisioningStateDeleting  ContactProfilesPropertiesProvisioningState = "deleting"
	ContactProfilesPropertiesProvisioningStateFailed    ContactProfilesPropertiesProvisioningState = "failed"
	ContactProfilesPropertiesProvisioningStateSucceeded ContactProfilesPropertiesProvisioningState = "succeeded"
	ContactProfilesPropertiesProvisioningStateUpdating  ContactProfilesPropertiesProvisioningState = "updating"
)

// PossibleContactProfilesPropertiesProvisioningStateValues returns the possible values for the ContactProfilesPropertiesProvisioningState const type.
func PossibleContactProfilesPropertiesProvisioningStateValues() []ContactProfilesPropertiesProvisioningState {
	return []ContactProfilesPropertiesProvisioningState{
		ContactProfilesPropertiesProvisioningStateCanceled,
		ContactProfilesPropertiesProvisioningStateCreating,
		ContactProfilesPropertiesProvisioningStateDeleting,
		ContactProfilesPropertiesProvisioningStateFailed,
		ContactProfilesPropertiesProvisioningStateSucceeded,
		ContactProfilesPropertiesProvisioningStateUpdating,
	}
}

// ContactsPropertiesProvisioningState - The current state of the resource's creation, deletion, or modification.
type ContactsPropertiesProvisioningState string

const (
	ContactsPropertiesProvisioningStateCanceled  ContactsPropertiesProvisioningState = "canceled"
	ContactsPropertiesProvisioningStateCreating  ContactsPropertiesProvisioningState = "creating"
	ContactsPropertiesProvisioningStateDeleting  ContactsPropertiesProvisioningState = "deleting"
	ContactsPropertiesProvisioningStateFailed    ContactsPropertiesProvisioningState = "failed"
	ContactsPropertiesProvisioningStateSucceeded ContactsPropertiesProvisioningState = "succeeded"
	ContactsPropertiesProvisioningStateUpdating  ContactsPropertiesProvisioningState = "updating"
)

// PossibleContactsPropertiesProvisioningStateValues returns the possible values for the ContactsPropertiesProvisioningState const type.
func PossibleContactsPropertiesProvisioningStateValues() []ContactsPropertiesProvisioningState {
	return []ContactsPropertiesProvisioningState{
		ContactsPropertiesProvisioningStateCanceled,
		ContactsPropertiesProvisioningStateCreating,
		ContactsPropertiesProvisioningStateDeleting,
		ContactsPropertiesProvisioningStateFailed,
		ContactsPropertiesProvisioningStateSucceeded,
		ContactsPropertiesProvisioningStateUpdating,
	}
}

// ContactsStatus - Status of a contact.
type ContactsStatus string

const (
	ContactsStatusCancelled         ContactsStatus = "cancelled"
	ContactsStatusFailed            ContactsStatus = "failed"
	ContactsStatusProviderCancelled ContactsStatus = "providerCancelled"
	ContactsStatusScheduled         ContactsStatus = "scheduled"
	ContactsStatusSucceeded         ContactsStatus = "succeeded"
)

// PossibleContactsStatusValues returns the possible values for the ContactsStatus const type.
func PossibleContactsStatusValues() []ContactsStatus {
	return []ContactsStatus{
		ContactsStatusCancelled,
		ContactsStatusFailed,
		ContactsStatusProviderCancelled,
		ContactsStatusScheduled,
		ContactsStatusSucceeded,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// Direction - Direction (Uplink or Downlink).
type Direction string

const (
	DirectionDownlink Direction = "Downlink"
	DirectionUplink   Direction = "Uplink"
)

// PossibleDirectionValues returns the possible values for the Direction const type.
func PossibleDirectionValues() []Direction {
	return []Direction{
		DirectionDownlink,
		DirectionUplink,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// Polarization - Polarization. e.g. (RHCP, LHCP).
type Polarization string

const (
	PolarizationLHCP             Polarization = "LHCP"
	PolarizationLinearHorizontal Polarization = "linearHorizontal"
	PolarizationLinearVertical   Polarization = "linearVertical"
	PolarizationRHCP             Polarization = "RHCP"
)

// PossiblePolarizationValues returns the possible values for the Polarization const type.
func PossiblePolarizationValues() []Polarization {
	return []Polarization{
		PolarizationLHCP,
		PolarizationLinearHorizontal,
		PolarizationLinearVertical,
		PolarizationRHCP,
	}
}

// Protocol - Protocol either UDP or TCP.
type Protocol string

const (
	ProtocolTCP Protocol = "TCP"
	ProtocolUDP Protocol = "UDP"
)

// PossibleProtocolValues returns the possible values for the Protocol const type.
func PossibleProtocolValues() []Protocol {
	return []Protocol{
		ProtocolTCP,
		ProtocolUDP,
	}
}

// ProvisioningState - The current state of the resource's creation, deletion, or modification.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "canceled"
	ProvisioningStateCreating  ProvisioningState = "creating"
	ProvisioningStateDeleting  ProvisioningState = "deleting"
	ProvisioningStateFailed    ProvisioningState = "failed"
	ProvisioningStateSucceeded ProvisioningState = "succeeded"
	ProvisioningStateUpdating  ProvisioningState = "updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ReleaseMode - Release Status of a ground station.
type ReleaseMode string

const (
	ReleaseModeGA      ReleaseMode = "GA"
	ReleaseModePreview ReleaseMode = "Preview"
)

// PossibleReleaseModeValues returns the possible values for the ReleaseMode const type.
func PossibleReleaseModeValues() []ReleaseMode {
	return []ReleaseMode{
		ReleaseModeGA,
		ReleaseModePreview,
	}
}

// SpacecraftsPropertiesProvisioningState - The current state of the resource's creation, deletion, or modification.
type SpacecraftsPropertiesProvisioningState string

const (
	SpacecraftsPropertiesProvisioningStateCanceled  SpacecraftsPropertiesProvisioningState = "canceled"
	SpacecraftsPropertiesProvisioningStateCreating  SpacecraftsPropertiesProvisioningState = "creating"
	SpacecraftsPropertiesProvisioningStateDeleting  SpacecraftsPropertiesProvisioningState = "deleting"
	SpacecraftsPropertiesProvisioningStateFailed    SpacecraftsPropertiesProvisioningState = "failed"
	SpacecraftsPropertiesProvisioningStateSucceeded SpacecraftsPropertiesProvisioningState = "succeeded"
	SpacecraftsPropertiesProvisioningStateUpdating  SpacecraftsPropertiesProvisioningState = "updating"
)

// PossibleSpacecraftsPropertiesProvisioningStateValues returns the possible values for the SpacecraftsPropertiesProvisioningState const type.
func PossibleSpacecraftsPropertiesProvisioningStateValues() []SpacecraftsPropertiesProvisioningState {
	return []SpacecraftsPropertiesProvisioningState{
		SpacecraftsPropertiesProvisioningStateCanceled,
		SpacecraftsPropertiesProvisioningStateCreating,
		SpacecraftsPropertiesProvisioningStateDeleting,
		SpacecraftsPropertiesProvisioningStateFailed,
		SpacecraftsPropertiesProvisioningStateSucceeded,
		SpacecraftsPropertiesProvisioningStateUpdating,
	}
}

// Status - The status of operation.
type Status string

const (
	StatusCanceled  Status = "Canceled"
	StatusFailed    Status = "Failed"
	StatusRunning   Status = "Running"
	StatusSucceeded Status = "Succeeded"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusCanceled,
		StatusFailed,
		StatusRunning,
		StatusSucceeded,
	}
}
