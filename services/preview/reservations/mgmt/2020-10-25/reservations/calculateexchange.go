package reservations

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CalculateExchangeClient is the client for the CalculateExchange methods of the Reservations service.
type CalculateExchangeClient struct {
	BaseClient
}

// NewCalculateExchangeClient creates an instance of the CalculateExchangeClient client.
func NewCalculateExchangeClient() CalculateExchangeClient {
	return NewCalculateExchangeClientWithBaseURI(DefaultBaseURI)
}

// NewCalculateExchangeClientWithBaseURI creates an instance of the CalculateExchangeClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewCalculateExchangeClientWithBaseURI(baseURI string) CalculateExchangeClient {
	return CalculateExchangeClient{NewWithBaseURI(baseURI)}
}

// Post calculates price for exchanging `Reservations` if there are no policy errors.
// Parameters:
// body - request containing purchases and refunds that need to be executed.
func (client CalculateExchangeClient) Post(ctx context.Context, body CalculateExchangeRequest) (result CalculateExchangePostFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CalculateExchangeClient.Post")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PostPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.CalculateExchangeClient", "Post", nil, "Failure preparing request")
		return
	}

	result, err = client.PostSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.CalculateExchangeClient", "Post", nil, "Failure sending request")
		return
	}

	return
}

// PostPreparer prepares the Post request.
func (client CalculateExchangeClient) PostPreparer(ctx context.Context, body CalculateExchangeRequest) (*http.Request, error) {
	const APIVersion = "2020-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Capacity/calculateExchange"),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostSender sends the Post request. The method will close the
// http.Response Body if it receives an error.
func (client CalculateExchangeClient) PostSender(req *http.Request) (future CalculateExchangePostFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client CalculateExchangeClient) (ceorr CalculateExchangeOperationResultResponse, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "reservations.CalculateExchangePostFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("reservations.CalculateExchangePostFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		if ceorr.Response.Response, err = future.GetResult(sender); err == nil && ceorr.Response.Response.StatusCode != http.StatusNoContent {
			ceorr, err = client.PostResponder(ceorr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "reservations.CalculateExchangePostFuture", "Result", ceorr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// PostResponder handles the response to the Post request. The method always
// closes the http.Response Body.
func (client CalculateExchangeClient) PostResponder(resp *http.Response) (result CalculateExchangeOperationResultResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}