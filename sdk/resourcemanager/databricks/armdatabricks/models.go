//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabricks

import "time"

// AccessConnector - Information about azure databricks accessConnector.
type AccessConnector struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity

	// Azure Databricks accessConnector properties
	Properties *AccessConnectorProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system metadata relating to this resource
	SystemData *SystemData

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// AccessConnectorListResult - List of azure databricks accessConnector.
type AccessConnectorListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// The array of azure databricks accessConnector.
	Value []*AccessConnector
}

type AccessConnectorProperties struct {
	// READ-ONLY; Provisioning status of the accessConnector.
	ProvisioningState *ProvisioningState
}

// AccessConnectorUpdate - An update to an azure databricks accessConnector.
type AccessConnectorUpdate struct {
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity

	// Resource tags.
	Tags map[string]*string
}

// AddressSpace contains an array of IP address ranges that can be used by subnets of the virtual network.
type AddressSpace struct {
	// A list of address blocks reserved for this virtual network in CIDR notation.
	AddressPrefixes []*string
}

// CreatedBy - Provides details of the entity that created/updated the workspace.
type CreatedBy struct {
	// READ-ONLY; The application ID of the application that initiated the creation of the workspace. For example, Azure Portal.
	ApplicationID *string

	// READ-ONLY; The Object ID that created the workspace.
	Oid *string

	// READ-ONLY; The Personal Object ID corresponding to the object ID above
	Puid *string
}

// Encryption - The object that contains details of encryption used on the workspace.
type Encryption struct {
	// The name of KeyVault key.
	KeyName *string

	// The encryption keySource (provider). Possible values (case-insensitive): Default, Microsoft.Keyvault
	KeySource *KeySource

	// The Uri of KeyVault.
	KeyVaultURI *string

	// The version of KeyVault key.
	KeyVersion *string
}

// EncryptionEntitiesDefinition - Encryption entities for databricks workspace resource.
type EncryptionEntitiesDefinition struct {
	// Encryption properties for the databricks managed disks.
	ManagedDisk *ManagedDiskEncryption

	// Encryption properties for the databricks managed services.
	ManagedServices *EncryptionV2
}

// EncryptionV2 - The object that contains details of encryption used on the workspace.
type EncryptionV2 struct {
	// REQUIRED; The encryption keySource (provider). Possible values (case-insensitive): Microsoft.Keyvault
	KeySource *EncryptionKeySource

	// Key Vault input properties for encryption.
	KeyVaultProperties *EncryptionV2KeyVaultProperties
}

// EncryptionV2KeyVaultProperties - Key Vault input properties for encryption.
type EncryptionV2KeyVaultProperties struct {
	// REQUIRED; The name of KeyVault key.
	KeyName *string

	// REQUIRED; The Uri of KeyVault.
	KeyVaultURI *string

	// REQUIRED; The version of KeyVault key.
	KeyVersion *string
}

// EndpointDependency - A domain name or IP address the Workspace is reaching at.
type EndpointDependency struct {
	// The domain name of the dependency.
	DomainName *string

	// The Ports used when connecting to domainName.
	EndpointDetails []*EndpointDetail
}

// EndpointDetail - Connect information from the Workspace to a single endpoint.
type EndpointDetail struct {
	// An IP Address that Domain Name currently resolves to.
	IPAddress *string

	// Whether it is possible to create a connection from the Workspace to this IpAddress at this Port.
	IsAccessible *bool

	// The time in milliseconds it takes for the connection to be created from the Workspace to this IpAddress at this Port.
	Latency *float64

	// The port an endpoint is connected to.
	Port *int32
}

// GroupIDInformation - The group information for creating a private endpoint on a workspace
type GroupIDInformation struct {
	// REQUIRED; The group id properties.
	Properties *GroupIDInformationProperties

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// GroupIDInformationProperties - The properties for a group information object
type GroupIDInformationProperties struct {
	// The group id
	GroupID *string

	// The required members for a specific group id
	RequiredMembers []*string

	// The required DNS zones for a specific group id
	RequiredZoneNames []*string
}

// ManagedDiskEncryption - The object that contains details of encryption used on the workspace.
type ManagedDiskEncryption struct {
	// REQUIRED; The encryption keySource (provider). Possible values (case-insensitive): Microsoft.Keyvault
	KeySource *EncryptionKeySource

	// REQUIRED; Key Vault input properties for encryption.
	KeyVaultProperties *ManagedDiskEncryptionKeyVaultProperties

	// Indicate whether the latest key version should be automatically used for Managed Disk Encryption.
	RotationToLatestKeyVersionEnabled *bool
}

// ManagedDiskEncryptionKeyVaultProperties - Key Vault input properties for encryption.
type ManagedDiskEncryptionKeyVaultProperties struct {
	// REQUIRED; The name of KeyVault key.
	KeyName *string

	// REQUIRED; The URI of KeyVault.
	KeyVaultURI *string

	// REQUIRED; The version of KeyVault key.
	KeyVersion *string
}

// ManagedIdentityConfiguration - The Managed Identity details for storage account.
type ManagedIdentityConfiguration struct {
	// READ-ONLY; The objectId of the Managed Identity that is linked to the Managed Storage account.
	PrincipalID *string

	// READ-ONLY; The tenant Id where the Managed Identity is created.
	TenantID *string

	// READ-ONLY; The type of Identity created. It can be either SystemAssigned or UserAssigned.
	Type *string
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// Operation - REST API operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Operation name: {provider}/{resource}/{operation}
	Name *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description for the resource operation.
	Description *string

	// Operation type: Read, write, delete, etc.
	Operation *string

	// Service provider: ex Microsoft.Databricks
	Provider *string

	// Resource on which the operation is performed.
	Resource *string
}

// OperationListResult - Result of the request to list Resource Provider operations. It contains a list of operations and
// a URL link to get the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of Resource Provider operations supported by the Resource Provider resource provider.
	Value []*Operation
}

// OutboundEnvironmentEndpoint - Egress endpoints which Workspace connects to for common purposes.
type OutboundEnvironmentEndpoint struct {
	// The category of endpoints accessed by the Workspace, e.g. azure-storage, azure-mysql, etc.
	Category *string

	// The endpoints that Workspace connect to
	Endpoints []*EndpointDependency
}

// PrivateEndpoint - The private endpoint property of a private endpoint connection
type PrivateEndpoint struct {
	// READ-ONLY; The resource identifier.
	ID *string
}

// PrivateEndpointConnection - The private endpoint connection of a workspace
type PrivateEndpointConnection struct {
	// REQUIRED; The private endpoint connection properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// PrivateEndpointConnectionProperties - The properties of a private endpoint connection
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; Private endpoint connection state
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// GroupIds from the private link service resource.
	GroupIDs []*string

	// Private endpoint
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; Provisioning state of the private endpoint connection.
	ProvisioningState *PrivateEndpointConnectionProvisioningState
}

// PrivateEndpointConnectionsList - List of private link connections.
type PrivateEndpointConnectionsList struct {
	// The URL to get the next set of endpoint connections.
	NextLink *string

	// The list of returned private endpoint connection.
	Value []*PrivateEndpointConnection
}

// PrivateLinkResourcesList - The available private link resources for a workspace
type PrivateLinkResourcesList struct {
	// The URL to get the next set of private link resources.
	NextLink *string

	// The list of available private link resources for a workspace
	Value []*GroupIDInformation
}

// PrivateLinkServiceConnectionState - The current state of a private endpoint connection
type PrivateLinkServiceConnectionState struct {
	// REQUIRED; The status of a private endpoint connection
	Status *PrivateLinkServiceConnectionStatus

	// Actions required for a private endpoint connection
	ActionsRequired *string

	// The description for the current state of a private endpoint connection
	Description *string
}

// SKU for the resource.
type SKU struct {
	// REQUIRED; The SKU name.
	Name *string

	// The SKU tier.
	Tier *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// VirtualNetworkPeering - Peerings in a VirtualNetwork resource
type VirtualNetworkPeering struct {
	// REQUIRED; List of properties for vNet Peering
	Properties *VirtualNetworkPeeringPropertiesFormat

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Name of the virtual network peering resource
	Name *string

	// READ-ONLY; type of the virtual network peering resource
	Type *string
}

// VirtualNetworkPeeringList - Gets all virtual network peerings under a workspace.
type VirtualNetworkPeeringList struct {
	// URL to get the next set of virtual network peering list results if there are any.
	NextLink *string

	// List of virtual network peerings on workspace.
	Value []*VirtualNetworkPeering
}

// VirtualNetworkPeeringPropertiesFormat - Properties of the virtual network peering.
type VirtualNetworkPeeringPropertiesFormat struct {
	// REQUIRED; The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	RemoteVirtualNetwork *VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork

	// Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
	AllowForwardedTraffic *bool

	// If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit *bool

	// Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
	AllowVirtualNetworkAccess *bool

	// The reference to the databricks virtual network address space.
	DatabricksAddressSpace *AddressSpace

	// The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	DatabricksVirtualNetwork *VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork

	// The reference to the remote virtual network address space.
	RemoteAddressSpace *AddressSpace

	// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering
	// is also true, virtual network will use gateways of remote virtual network
	// for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
	// gateway.
	UseRemoteGateways *bool

	// READ-ONLY; The status of the virtual network peering.
	PeeringState *PeeringState

	// READ-ONLY; The provisioning state of the virtual network peering resource.
	ProvisioningState *PeeringProvisioningState
}

// VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork - The remote virtual network should be in the same region.
// See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
type VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork struct {
	// The Id of the databricks virtual network.
	ID *string
}

// VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork - The remote virtual network should be in the same region. See
// here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
type VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork struct {
	// The Id of the remote virtual network.
	ID *string
}

// Workspace - Information about workspace.
type Workspace struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The workspace properties.
	Properties *WorkspaceProperties

	// The SKU of the resource.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system metadata relating to this resource
	SystemData *SystemData

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// WorkspaceCustomBooleanParameter - The value which should be used for this field.
type WorkspaceCustomBooleanParameter struct {
	// REQUIRED; The value which should be used for this field.
	Value *bool

	// READ-ONLY; The type of variable that this is
	Type *CustomParameterType
}

// WorkspaceCustomObjectParameter - The value which should be used for this field.
type WorkspaceCustomObjectParameter struct {
	// REQUIRED; The value which should be used for this field.
	Value any

	// READ-ONLY; The type of variable that this is
	Type *CustomParameterType
}

// WorkspaceCustomParameters - Custom Parameters used for Cluster Creation.
type WorkspaceCustomParameters struct {
	// The ID of a Azure Machine Learning workspace to link with Databricks workspace
	AmlWorkspaceID *WorkspaceCustomStringParameter

	// The name of the Private Subnet within the Virtual Network
	CustomPrivateSubnetName *WorkspaceCustomStringParameter

	// The name of a Public Subnet within the Virtual Network
	CustomPublicSubnetName *WorkspaceCustomStringParameter

	// The ID of a Virtual Network where this Databricks Cluster should be created
	CustomVirtualNetworkID *WorkspaceCustomStringParameter

	// Should the Public IP be Disabled?
	EnableNoPublicIP *WorkspaceCustomBooleanParameter

	// Contains the encryption details for Customer-Managed Key (CMK) enabled workspace.
	Encryption *WorkspaceEncryptionParameter

	// Name of the outbound Load Balancer Backend Pool for Secure Cluster Connectivity (No Public IP).
	LoadBalancerBackendPoolName *WorkspaceCustomStringParameter

	// Resource URI of Outbound Load balancer for Secure Cluster Connectivity (No Public IP) workspace.
	LoadBalancerID *WorkspaceCustomStringParameter

	// Name of the NAT gateway for Secure Cluster Connectivity (No Public IP) workspace subnets.
	NatGatewayName *WorkspaceCustomStringParameter

	// Prepare the workspace for encryption. Enables the Managed Identity for managed storage account.
	PrepareEncryption *WorkspaceCustomBooleanParameter

	// Name of the Public IP for No Public IP workspace with managed vNet.
	PublicIPName *WorkspaceCustomStringParameter

	// A boolean indicating whether or not the DBFS root file system will be enabled with secondary layer of encryption with platform
	// managed keys for data at rest.
	RequireInfrastructureEncryption *WorkspaceCustomBooleanParameter

	// Default DBFS storage account name.
	StorageAccountName *WorkspaceCustomStringParameter

	// Storage account SKU name, ex: StandardGRS, StandardLRS. Refer https://aka.ms/storageskus for valid inputs.
	StorageAccountSKUName *WorkspaceCustomStringParameter

	// Address prefix for Managed virtual network. Default value for this input is 10.139.
	VnetAddressPrefix *WorkspaceCustomStringParameter

	// READ-ONLY; Tags applied to resources under Managed resource group. These can be updated by updating tags at workspace level.
	ResourceTags *WorkspaceCustomObjectParameter
}

// WorkspaceCustomStringParameter - The Value.
type WorkspaceCustomStringParameter struct {
	// REQUIRED; The value which should be used for this field.
	Value *string

	// READ-ONLY; The type of variable that this is
	Type *CustomParameterType
}

// WorkspaceEncryptionParameter - The object that contains details of encryption used on the workspace.
type WorkspaceEncryptionParameter struct {
	// The value which should be used for this field.
	Value *Encryption

	// READ-ONLY; The type of variable that this is
	Type *CustomParameterType
}

// WorkspaceListResult - List of workspaces.
type WorkspaceListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// The array of workspaces.
	Value []*Workspace
}

// WorkspaceProperties - The workspace properties.
type WorkspaceProperties struct {
	// REQUIRED; The managed resource group Id.
	ManagedResourceGroupID *string

	// The workspace provider authorizations.
	Authorizations []*WorkspaceProviderAuthorization

	// Indicates the Object ID, PUID and Application ID of entity that created the workspace.
	CreatedBy *CreatedBy

	// Encryption properties for databricks workspace
	Encryption *WorkspacePropertiesEncryption

	// The details of Managed Identity of Disk Encryption Set used for Managed Disk Encryption
	ManagedDiskIdentity *ManagedIdentityConfiguration

	// The workspace's custom parameters.
	Parameters *WorkspaceCustomParameters

	// The network access type for accessing workspace. Set value to disabled to access workspace only via private link.
	PublicNetworkAccess *PublicNetworkAccess

	// Gets or sets a value indicating whether data plane (clusters) to control plane communication happen over private endpoint.
	// Supported values are 'AllRules' and 'NoAzureDatabricksRules'.
	// 'NoAzureServiceRules' value is for internal use only.
	RequiredNsgRules *RequiredNsgRules

	// The details of Managed Identity of Storage Account
	StorageAccountIdentity *ManagedIdentityConfiguration

	// The blob URI where the UI definition file is located.
	UIDefinitionURI *string

	// Indicates the Object ID, PUID and Application ID of entity that last updated the workspace.
	UpdatedBy *CreatedBy

	// READ-ONLY; Specifies the date and time when the workspace is created.
	CreatedDateTime *time.Time

	// READ-ONLY; The resource Id of the managed disk encryption set.
	DiskEncryptionSetID *string

	// READ-ONLY; Private endpoint connections created on the workspace
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; The workspace provisioning state.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The unique identifier of the databricks workspace in databricks control plane.
	WorkspaceID *string

	// READ-ONLY; The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
	WorkspaceURL *string
}

// WorkspacePropertiesEncryption - Encryption properties for databricks workspace
type WorkspacePropertiesEncryption struct {
	// REQUIRED; Encryption entities definition for the workspace.
	Entities *EncryptionEntitiesDefinition
}

// WorkspaceProviderAuthorization - The workspace provider authorization.
type WorkspaceProviderAuthorization struct {
	// REQUIRED; The provider's principal identifier. This is the identity that the provider will use to call ARM to manage the
	// workspace resources.
	PrincipalID *string

	// REQUIRED; The provider's role definition identifier. This role will define all the permissions that the provider must have
	// on the workspace's container resource group. This role definition cannot have
	// permission to delete the resource group.
	RoleDefinitionID *string
}

// WorkspaceUpdate - An update to a workspace.
type WorkspaceUpdate struct {
	// Resource tags.
	Tags map[string]*string
}
