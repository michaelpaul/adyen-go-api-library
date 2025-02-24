/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// TransfersApi service
type TransfersApi common.Service

// All parameters accepted by TransfersApi.ReturnTransfer
type TransfersApiReturnTransferInput struct {
	id                    string
	returnTransferRequest *ReturnTransferRequest
}

func (r TransfersApiReturnTransferInput) ReturnTransferRequest(returnTransferRequest ReturnTransferRequest) TransfersApiReturnTransferInput {
	r.returnTransferRequest = &returnTransferRequest
	return r
}

/*
Prepare a request for ReturnTransfer
@param id The unique identifier of the transfer to be returned.
@return TransfersApiReturnTransferInput
*/
func (a *TransfersApi) ReturnTransferInput(id string) TransfersApiReturnTransferInput {
	return TransfersApiReturnTransferInput{
		id: id,
	}
}

/*
ReturnTransfer Return a transfer

Returns previously transferred funds without creating a new `transferId`.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransfersApiReturnTransferInput - Request parameters, see ReturnTransferInput
@return ReturnTransferResponse, *http.Response, error
*/
func (a *TransfersApi) ReturnTransfer(ctx context.Context, r TransfersApiReturnTransferInput) (ReturnTransferResponse, *http.Response, error) {
	res := &ReturnTransferResponse{}
	path := "/transfers/{id}/returns"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.returnTransferRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TransfersApi.TransferFunds
type TransfersApiTransferFundsInput struct {
	wWWAuthenticate *string
	transferInfo    *TransferInfo
}

// Header for authenticating through SCA
func (r TransfersApiTransferFundsInput) WWWAuthenticate(wWWAuthenticate string) TransfersApiTransferFundsInput {
	r.wWWAuthenticate = &wWWAuthenticate
	return r
}

func (r TransfersApiTransferFundsInput) TransferInfo(transferInfo TransferInfo) TransfersApiTransferFundsInput {
	r.transferInfo = &transferInfo
	return r
}

/*
Prepare a request for TransferFunds

@return TransfersApiTransferFundsInput
*/
func (a *TransfersApi) TransferFundsInput() TransfersApiTransferFundsInput {
	return TransfersApiTransferFundsInput{}
}

/*
TransferFunds Transfer funds

>Versions 1 and 2 of the Transfers API are deprecated. If you are just starting your implementation, use the latest version.

Starts a request to transfer funds to [balance accounts](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts), [transfer instruments](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments), or third-party bank accounts. Adyen sends the outcome of the transfer request through webhooks.

To use this endpoint, you need an additional role for your API credential and transfers must be enabled for the source balance account. Your Adyen contact will set these up for you.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransfersApiTransferFundsInput - Request parameters, see TransferFundsInput
@return Transfer, *http.Response, error
*/
func (a *TransfersApi) TransferFunds(ctx context.Context, r TransfersApiTransferFundsInput) (Transfer, *http.Response, error) {
	res := &Transfer{}
	path := "/transfers"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.wWWAuthenticate != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "WWW-Authenticate", r.wWWAuthenticate, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.transferInfo,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}
