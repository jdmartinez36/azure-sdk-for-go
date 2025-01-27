//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AccountsCheckNameAvailabilityResponse contains the response from method Accounts.CheckNameAvailability.
type AccountsCheckNameAvailabilityResponse struct {
	AccountsCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsCheckNameAvailabilityResult contains the result from method Accounts.CheckNameAvailability.
type AccountsCheckNameAvailabilityResult struct {
	NameAvailabilityInformation
}

// AccountsCreatePollerResponse contains the response from method Accounts.Create.
type AccountsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsCreateResponse, error) {
	respType := AccountsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DataLakeAnalyticsAccount)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsCreatePollerResponse from the provided client and resume token.
func (l *AccountsCreatePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &AccountsCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsCreateResponse contains the response from method Accounts.Create.
type AccountsCreateResponse struct {
	AccountsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsCreateResult contains the result from method Accounts.Create.
type AccountsCreateResult struct {
	DataLakeAnalyticsAccount
}

// AccountsDeletePollerResponse contains the response from method Accounts.Delete.
type AccountsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsDeleteResponse, error) {
	respType := AccountsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsDeletePollerResponse from the provided client and resume token.
func (l *AccountsDeletePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &AccountsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsDeleteResponse contains the response from method Accounts.Delete.
type AccountsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsGetResponse contains the response from method Accounts.Get.
type AccountsGetResponse struct {
	AccountsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsGetResult contains the result from method Accounts.Get.
type AccountsGetResult struct {
	DataLakeAnalyticsAccount
}

// AccountsListByResourceGroupResponse contains the response from method Accounts.ListByResourceGroup.
type AccountsListByResourceGroupResponse struct {
	AccountsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsListByResourceGroupResult contains the result from method Accounts.ListByResourceGroup.
type AccountsListByResourceGroupResult struct {
	DataLakeAnalyticsAccountListResult
}

// AccountsListResponse contains the response from method Accounts.List.
type AccountsListResponse struct {
	AccountsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsListResult contains the result from method Accounts.List.
type AccountsListResult struct {
	DataLakeAnalyticsAccountListResult
}

// AccountsUpdatePollerResponse contains the response from method Accounts.Update.
type AccountsUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AccountsUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AccountsUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AccountsUpdateResponse, error) {
	respType := AccountsUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DataLakeAnalyticsAccount)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AccountsUpdatePollerResponse from the provided client and resume token.
func (l *AccountsUpdatePollerResponse) Resume(ctx context.Context, client *AccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AccountsClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &AccountsUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AccountsUpdateResponse contains the response from method Accounts.Update.
type AccountsUpdateResponse struct {
	AccountsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AccountsUpdateResult contains the result from method Accounts.Update.
type AccountsUpdateResult struct {
	DataLakeAnalyticsAccount
}

// ComputePoliciesCreateOrUpdateResponse contains the response from method ComputePolicies.CreateOrUpdate.
type ComputePoliciesCreateOrUpdateResponse struct {
	ComputePoliciesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputePoliciesCreateOrUpdateResult contains the result from method ComputePolicies.CreateOrUpdate.
type ComputePoliciesCreateOrUpdateResult struct {
	ComputePolicy
}

// ComputePoliciesDeleteResponse contains the response from method ComputePolicies.Delete.
type ComputePoliciesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputePoliciesGetResponse contains the response from method ComputePolicies.Get.
type ComputePoliciesGetResponse struct {
	ComputePoliciesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputePoliciesGetResult contains the result from method ComputePolicies.Get.
type ComputePoliciesGetResult struct {
	ComputePolicy
}

// ComputePoliciesListByAccountResponse contains the response from method ComputePolicies.ListByAccount.
type ComputePoliciesListByAccountResponse struct {
	ComputePoliciesListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputePoliciesListByAccountResult contains the result from method ComputePolicies.ListByAccount.
type ComputePoliciesListByAccountResult struct {
	ComputePolicyListResult
}

// ComputePoliciesUpdateResponse contains the response from method ComputePolicies.Update.
type ComputePoliciesUpdateResponse struct {
	ComputePoliciesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ComputePoliciesUpdateResult contains the result from method ComputePolicies.Update.
type ComputePoliciesUpdateResult struct {
	ComputePolicy
}

// DataLakeStoreAccountsAddResponse contains the response from method DataLakeStoreAccounts.Add.
type DataLakeStoreAccountsAddResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataLakeStoreAccountsDeleteResponse contains the response from method DataLakeStoreAccounts.Delete.
type DataLakeStoreAccountsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataLakeStoreAccountsGetResponse contains the response from method DataLakeStoreAccounts.Get.
type DataLakeStoreAccountsGetResponse struct {
	DataLakeStoreAccountsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataLakeStoreAccountsGetResult contains the result from method DataLakeStoreAccounts.Get.
type DataLakeStoreAccountsGetResult struct {
	DataLakeStoreAccountInformation
}

// DataLakeStoreAccountsListByAccountResponse contains the response from method DataLakeStoreAccounts.ListByAccount.
type DataLakeStoreAccountsListByAccountResponse struct {
	DataLakeStoreAccountsListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataLakeStoreAccountsListByAccountResult contains the result from method DataLakeStoreAccounts.ListByAccount.
type DataLakeStoreAccountsListByAccountResult struct {
	DataLakeStoreAccountInformationListResult
}

// FirewallRulesCreateOrUpdateResponse contains the response from method FirewallRules.CreateOrUpdate.
type FirewallRulesCreateOrUpdateResponse struct {
	FirewallRulesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesCreateOrUpdateResult contains the result from method FirewallRules.CreateOrUpdate.
type FirewallRulesCreateOrUpdateResult struct {
	FirewallRule
}

// FirewallRulesDeleteResponse contains the response from method FirewallRules.Delete.
type FirewallRulesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesGetResponse contains the response from method FirewallRules.Get.
type FirewallRulesGetResponse struct {
	FirewallRulesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesGetResult contains the result from method FirewallRules.Get.
type FirewallRulesGetResult struct {
	FirewallRule
}

// FirewallRulesListByAccountResponse contains the response from method FirewallRules.ListByAccount.
type FirewallRulesListByAccountResponse struct {
	FirewallRulesListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesListByAccountResult contains the result from method FirewallRules.ListByAccount.
type FirewallRulesListByAccountResult struct {
	FirewallRuleListResult
}

// FirewallRulesUpdateResponse contains the response from method FirewallRules.Update.
type FirewallRulesUpdateResponse struct {
	FirewallRulesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FirewallRulesUpdateResult contains the result from method FirewallRules.Update.
type FirewallRulesUpdateResult struct {
	FirewallRule
}

// LocationsGetCapabilityResponse contains the response from method Locations.GetCapability.
type LocationsGetCapabilityResponse struct {
	LocationsGetCapabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LocationsGetCapabilityResult contains the result from method Locations.GetCapability.
type LocationsGetCapabilityResult struct {
	CapabilityInformation
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// StorageAccountsAddResponse contains the response from method StorageAccounts.Add.
type StorageAccountsAddResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageAccountsDeleteResponse contains the response from method StorageAccounts.Delete.
type StorageAccountsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageAccountsGetResponse contains the response from method StorageAccounts.Get.
type StorageAccountsGetResponse struct {
	StorageAccountsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageAccountsGetResult contains the result from method StorageAccounts.Get.
type StorageAccountsGetResult struct {
	StorageAccountInformation
}

// StorageAccountsGetStorageContainerResponse contains the response from method StorageAccounts.GetStorageContainer.
type StorageAccountsGetStorageContainerResponse struct {
	StorageAccountsGetStorageContainerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageAccountsGetStorageContainerResult contains the result from method StorageAccounts.GetStorageContainer.
type StorageAccountsGetStorageContainerResult struct {
	StorageContainer
}

// StorageAccountsListByAccountResponse contains the response from method StorageAccounts.ListByAccount.
type StorageAccountsListByAccountResponse struct {
	StorageAccountsListByAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageAccountsListByAccountResult contains the result from method StorageAccounts.ListByAccount.
type StorageAccountsListByAccountResult struct {
	StorageAccountInformationListResult
}

// StorageAccountsListSasTokensResponse contains the response from method StorageAccounts.ListSasTokens.
type StorageAccountsListSasTokensResponse struct {
	StorageAccountsListSasTokensResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageAccountsListSasTokensResult contains the result from method StorageAccounts.ListSasTokens.
type StorageAccountsListSasTokensResult struct {
	SasTokenInformationListResult
}

// StorageAccountsListStorageContainersResponse contains the response from method StorageAccounts.ListStorageContainers.
type StorageAccountsListStorageContainersResponse struct {
	StorageAccountsListStorageContainersResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageAccountsListStorageContainersResult contains the result from method StorageAccounts.ListStorageContainers.
type StorageAccountsListStorageContainersResult struct {
	StorageContainerListResult
}

// StorageAccountsUpdateResponse contains the response from method StorageAccounts.Update.
type StorageAccountsUpdateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
