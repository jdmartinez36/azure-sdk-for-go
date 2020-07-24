// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GalleriesOperations contains the methods for the Galleries group.
type GalleriesOperations interface {
	// BeginCreateOrUpdate - Create or update a Shared Image Gallery.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery Gallery) (*GalleryPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (GalleryPoller, error)
	// BeginDelete - Delete a Shared Image Gallery.
	BeginDelete(ctx context.Context, resourceGroupName string, galleryName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves information about a Shared Image Gallery.
	Get(ctx context.Context, resourceGroupName string, galleryName string) (*GalleryResponse, error)
	// List - List galleries under a subscription.
	List() (GalleryListPager, error)
	// ListByResourceGroup - List galleries under a resource group.
	ListByResourceGroup(resourceGroupName string) (GalleryListPager, error)
	// BeginUpdate - Update a Shared Image Gallery.
	BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery GalleryUpdate) (*GalleryPollerResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (GalleryPoller, error)
}

// galleriesOperations implements the GalleriesOperations interface.
type galleriesOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Create or update a Shared Image Gallery.
func (client *galleriesOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery Gallery) (*GalleryPollerResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, galleryName, gallery)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("galleriesOperations.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &galleryPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*GalleryResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *galleriesOperations) ResumeCreateOrUpdate(token string) (GalleryPoller, error) {
	pt, err := resumePollingTracker("galleriesOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *galleriesOperations) createOrUpdateCreateRequest(resourceGroupName string, galleryName string, gallery Gallery) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(gallery)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *galleriesOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*GalleryPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return &GalleryPollerResponse{RawResponse: resp.Response}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *galleriesOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Delete a Shared Image Gallery.
func (client *galleriesOperations) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string) (*HTTPPollerResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, galleryName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("galleriesOperations.Delete", "", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *galleriesOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("galleriesOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *galleriesOperations) deleteCreateRequest(resourceGroupName string, galleryName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *galleriesOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *galleriesOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Retrieves information about a Shared Image Gallery.
func (client *galleriesOperations) Get(ctx context.Context, resourceGroupName string, galleryName string) (*GalleryResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, galleryName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *galleriesOperations) getCreateRequest(resourceGroupName string, galleryName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *galleriesOperations) getHandleResponse(resp *azcore.Response) (*GalleryResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := GalleryResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Gallery)
}

// getHandleError handles the Get error response.
func (client *galleriesOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - List galleries under a subscription.
func (client *galleriesOperations) List() (GalleryListPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &galleryListPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *GalleryListResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.GalleryList.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.GalleryList.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *galleriesOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/galleries"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *galleriesOperations) listHandleResponse(resp *azcore.Response) (*GalleryListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := GalleryListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryList)
}

// listHandleError handles the List error response.
func (client *galleriesOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - List galleries under a resource group.
func (client *galleriesOperations) ListByResourceGroup(resourceGroupName string) (GalleryListPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &galleryListPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *GalleryListResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.GalleryList.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.GalleryList.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *galleriesOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *galleriesOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*GalleryListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := GalleryListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryList)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *galleriesOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Update - Update a Shared Image Gallery.
func (client *galleriesOperations) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery GalleryUpdate) (*GalleryPollerResponse, error) {
	req, err := client.updateCreateRequest(resourceGroupName, galleryName, gallery)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("galleriesOperations.Update", "", resp, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &galleryPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*GalleryResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *galleriesOperations) ResumeUpdate(token string) (GalleryPoller, error) {
	pt, err := resumePollingTracker("galleriesOperations.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// updateCreateRequest creates the Update request.
func (client *galleriesOperations) updateCreateRequest(resourceGroupName string, galleryName string, gallery GalleryUpdate) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(gallery)
}

// updateHandleResponse handles the Update response.
func (client *galleriesOperations) updateHandleResponse(resp *azcore.Response) (*GalleryPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.updateHandleError(resp)
	}
	return &GalleryPollerResponse{RawResponse: resp.Response}, nil
}

// updateHandleError handles the Update error response.
func (client *galleriesOperations) updateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}