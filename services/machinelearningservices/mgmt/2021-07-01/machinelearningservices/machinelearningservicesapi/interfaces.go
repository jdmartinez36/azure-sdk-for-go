package machinelearningservicesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/machinelearningservices/mgmt/2021-07-01/machinelearningservices"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result machinelearningservices.OperationListResult, err error)
}

var _ OperationsClientAPI = (*machinelearningservices.OperationsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters machinelearningservices.Workspace) (result machinelearningservices.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.WorkspacesDeleteFuture, err error)
	Diagnose(ctx context.Context, resourceGroupName string, workspaceName string, parameters *machinelearningservices.DiagnoseWorkspaceParameters) (result machinelearningservices.WorkspacesDiagnoseFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.Workspace, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skip string) (result machinelearningservices.WorkspaceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skip string) (result machinelearningservices.WorkspaceListResultIterator, err error)
	ListBySubscription(ctx context.Context, skip string) (result machinelearningservices.WorkspaceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, skip string) (result machinelearningservices.WorkspaceListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListWorkspaceKeysResult, err error)
	ListNotebookAccessToken(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.NotebookAccessTokenResult, err error)
	ListNotebookKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListNotebookKeysResult, err error)
	ListOutboundNetworkDependenciesEndpoints(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ExternalFQDNResponse, err error)
	ListStorageAccountKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListStorageAccountKeysResult, err error)
	PrepareNotebook(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.WorkspacesPrepareNotebookFuture, err error)
	ResyncKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.WorkspacesResyncKeysFuture, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters machinelearningservices.WorkspaceUpdateParameters) (result machinelearningservices.Workspace, err error)
}

var _ WorkspacesClientAPI = (*machinelearningservices.WorkspacesClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result machinelearningservices.ListUsagesResultPage, err error)
	ListComplete(ctx context.Context, location string) (result machinelearningservices.ListUsagesResultIterator, err error)
}

var _ UsagesClientAPI = (*machinelearningservices.UsagesClient)(nil)

// VirtualMachineSizesClientAPI contains the set of methods on the VirtualMachineSizesClient type.
type VirtualMachineSizesClientAPI interface {
	List(ctx context.Context, location string) (result machinelearningservices.VirtualMachineSizeListResult, err error)
}

var _ VirtualMachineSizesClientAPI = (*machinelearningservices.VirtualMachineSizesClient)(nil)

// QuotasClientAPI contains the set of methods on the QuotasClient type.
type QuotasClientAPI interface {
	List(ctx context.Context, location string) (result machinelearningservices.ListWorkspaceQuotasPage, err error)
	ListComplete(ctx context.Context, location string) (result machinelearningservices.ListWorkspaceQuotasIterator, err error)
	Update(ctx context.Context, location string, parameters machinelearningservices.QuotaUpdateParameters) (result machinelearningservices.UpdateWorkspaceQuotasResult, err error)
}

var _ QuotasClientAPI = (*machinelearningservices.QuotasClient)(nil)

// ComputeClientAPI contains the set of methods on the ComputeClient type.
type ComputeClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters machinelearningservices.ComputeResource) (result machinelearningservices.ComputeCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction machinelearningservices.UnderlyingResourceAction) (result machinelearningservices.ComputeDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeResource, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, skip string) (result machinelearningservices.PaginatedComputeResourcesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, skip string) (result machinelearningservices.PaginatedComputeResourcesListIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeSecretsModel, err error)
	ListNodes(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.AmlComputeNodesInformationPage, err error)
	ListNodesComplete(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.AmlComputeNodesInformationIterator, err error)
	Restart(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters machinelearningservices.ClusterUpdateParameters) (result machinelearningservices.ComputeUpdateFuture, err error)
}

var _ ComputeClientAPI = (*machinelearningservices.ComputeClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, properties machinelearningservices.PrivateEndpointConnection) (result machinelearningservices.PrivateEndpointConnection, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string) (result machinelearningservices.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.PrivateEndpointConnectionListResult, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*machinelearningservices.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*machinelearningservices.PrivateLinkResourcesClient)(nil)

// WorkspaceConnectionsClientAPI contains the set of methods on the WorkspaceConnectionsClient type.
type WorkspaceConnectionsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, parameters machinelearningservices.WorkspaceConnection) (result machinelearningservices.WorkspaceConnection, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string) (result machinelearningservices.WorkspaceConnection, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, target string, category string) (result machinelearningservices.PaginatedWorkspaceConnectionsList, err error)
}

var _ WorkspaceConnectionsClientAPI = (*machinelearningservices.WorkspaceConnectionsClient)(nil)

// WorkspaceFeaturesClientAPI contains the set of methods on the WorkspaceFeaturesClient type.
type WorkspaceFeaturesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListAmlUserFeatureResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListAmlUserFeatureResultIterator, err error)
}

var _ WorkspaceFeaturesClientAPI = (*machinelearningservices.WorkspaceFeaturesClient)(nil)

// WorkspaceSkusClientAPI contains the set of methods on the WorkspaceSkusClient type.
type WorkspaceSkusClientAPI interface {
	List(ctx context.Context) (result machinelearningservices.SkuListResultPage, err error)
	ListComplete(ctx context.Context) (result machinelearningservices.SkuListResultIterator, err error)
}

var _ WorkspaceSkusClientAPI = (*machinelearningservices.WorkspaceSkusClient)(nil)