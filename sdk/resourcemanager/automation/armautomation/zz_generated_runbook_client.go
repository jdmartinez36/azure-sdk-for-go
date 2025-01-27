//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RunbookClient contains the methods for the Runbook group.
// Don't use this type directly, use NewRunbookClient() instead.
type RunbookClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRunbookClient creates a new instance of RunbookClient with the specified values.
func NewRunbookClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RunbookClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RunbookClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create the runbook identified by runbook name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunbookClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters RunbookCreateOrUpdateParameters, options *RunbookCreateOrUpdateOptions) (RunbookCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, parameters, options)
	if err != nil {
		return RunbookCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RunbookCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RunbookCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RunbookClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters RunbookCreateOrUpdateParameters, options *RunbookCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RunbookClient) createOrUpdateHandleResponse(resp *http.Response) (RunbookCreateOrUpdateResponse, error) {
	result := RunbookCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Runbook); err != nil {
		return RunbookCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RunbookClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete the runbook by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunbookClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookDeleteOptions) (RunbookDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return RunbookDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RunbookDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RunbookDeleteResponse{}, client.deleteHandleError(resp)
	}
	return RunbookDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RunbookClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RunbookClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieve the runbook identified by runbook name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunbookClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookGetOptions) (RunbookGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return RunbookGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RunbookGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RunbookGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RunbookClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RunbookClient) getHandleResponse(resp *http.Response) (RunbookGetResponse, error) {
	result := RunbookGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Runbook); err != nil {
		return RunbookGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RunbookClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetContent - Retrieve the content of runbook identified by runbook name.
// If the operation fails it returns a generic error.
func (client *RunbookClient) GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookGetContentOptions) (RunbookGetContentResponse, error) {
	req, err := client.getContentCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return RunbookGetContentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RunbookGetContentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return RunbookGetContentResponse{}, client.getContentHandleError(resp)
	}
	return RunbookGetContentResponse{RawResponse: resp}, nil
}

// getContentCreateRequest creates the GetContent request.
func (client *RunbookClient) getContentCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookGetContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/content"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "text/powershell")
	return req, nil
}

// getContentHandleError handles the GetContent error response.
func (client *RunbookClient) getContentHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByAutomationAccount - Retrieve a list of runbooks.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunbookClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string, options *RunbookListByAutomationAccountOptions) *RunbookListByAutomationAccountPager {
	return &RunbookListByAutomationAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
		},
		advancer: func(ctx context.Context, resp RunbookListByAutomationAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RunbookListResult.NextLink)
		},
	}
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *RunbookClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *RunbookListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *RunbookClient) listByAutomationAccountHandleResponse(resp *http.Response) (RunbookListByAutomationAccountResponse, error) {
	result := RunbookListByAutomationAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunbookListResult); err != nil {
		return RunbookListByAutomationAccountResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByAutomationAccountHandleError handles the ListByAutomationAccount error response.
func (client *RunbookClient) listByAutomationAccountHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginPublish - Publish runbook draft.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunbookClient) BeginPublish(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookBeginPublishOptions) (RunbookPublishPollerResponse, error) {
	resp, err := client.publish(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return RunbookPublishPollerResponse{}, err
	}
	result := RunbookPublishPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RunbookClient.Publish", "", resp, client.pl, client.publishHandleError)
	if err != nil {
		return RunbookPublishPollerResponse{}, err
	}
	result.Poller = &RunbookPublishPoller{
		pt: pt,
	}
	return result, nil
}

// Publish - Publish runbook draft.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunbookClient) publish(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookBeginPublishOptions) (*http.Response, error) {
	req, err := client.publishCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.publishHandleError(resp)
	}
	return resp, nil
}

// publishCreateRequest creates the Publish request.
func (client *RunbookClient) publishCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *RunbookBeginPublishOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/publish"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// publishHandleError handles the Publish error response.
func (client *RunbookClient) publishHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update the runbook identified by runbook name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunbookClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters RunbookUpdateParameters, options *RunbookUpdateOptions) (RunbookUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, parameters, options)
	if err != nil {
		return RunbookUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RunbookUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RunbookUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *RunbookClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters RunbookUpdateParameters, options *RunbookUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *RunbookClient) updateHandleResponse(resp *http.Response) (RunbookUpdateResponse, error) {
	result := RunbookUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Runbook); err != nil {
		return RunbookUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *RunbookClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
