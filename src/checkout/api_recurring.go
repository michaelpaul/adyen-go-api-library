/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"context"
	_nethttp "net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// RecurringApi RecurringApi service
type RecurringApi common.Service

type RecurringApiDeleteTokenForStoredPaymentDetailsConfig struct {
	ctx              context.Context
	recurringId      string
	shopperReference *string
	merchantAccount  *string
}

// Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address.
func (r RecurringApiDeleteTokenForStoredPaymentDetailsConfig) ShopperReference(shopperReference string) RecurringApiDeleteTokenForStoredPaymentDetailsConfig {
	r.shopperReference = &shopperReference
	return r
}

// Your merchant account.
func (r RecurringApiDeleteTokenForStoredPaymentDetailsConfig) MerchantAccount(merchantAccount string) RecurringApiDeleteTokenForStoredPaymentDetailsConfig {
	r.merchantAccount = &merchantAccount
	return r
}

/*
DeleteTokenForStoredPaymentDetails Delete a token for stored payment details

Deletes the token identified in the path. The token can no longer be used with payment requests.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param recurringId The unique identifier of the token.
 @return RecurringApiDeleteTokenForStoredPaymentDetailsConfig
*/
func (a *RecurringApi) DeleteTokenForStoredPaymentDetailsConfig(ctx context.Context, recurringId string) RecurringApiDeleteTokenForStoredPaymentDetailsConfig {
	return RecurringApiDeleteTokenForStoredPaymentDetailsConfig{
		ctx:         ctx,
		recurringId: recurringId,
	}
}

/*
Delete a token for stored payment details
Deletes the token identified in the path. The token can no longer be used with payment requests.
 * @param recurringId The unique identifier of the token.
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return StoredPaymentMethodResource
*/

func (a *RecurringApi) DeleteTokenForStoredPaymentDetails(r RecurringApiDeleteTokenForStoredPaymentDetailsConfig) (StoredPaymentMethodResource, *_nethttp.Response, error) {
	res := &StoredPaymentMethodResource{}
	path := "/storedPaymentMethods/{recurringId}"
	path = strings.Replace(path, "{"+"recurringId"+"}", url.PathEscape(common.ParameterValueToString(r.recurringId, "recurringId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.shopperReference != nil {
		common.ParameterAddToQuery(queryParams, "shopperReference", r.shopperReference, "")
	}
	if r.merchantAccount != nil {
		common.ParameterAddToQuery(queryParams, "merchantAccount", r.merchantAccount, "")
	}
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		nil,
		res,
		_nethttp.MethodDelete,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type RecurringApiGetTokensForStoredPaymentDetailsConfig struct {
	ctx              context.Context
	shopperReference *string
	merchantAccount  *string
}

// Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address.
func (r RecurringApiGetTokensForStoredPaymentDetailsConfig) ShopperReference(shopperReference string) RecurringApiGetTokensForStoredPaymentDetailsConfig {
	r.shopperReference = &shopperReference
	return r
}

// Your merchant account.
func (r RecurringApiGetTokensForStoredPaymentDetailsConfig) MerchantAccount(merchantAccount string) RecurringApiGetTokensForStoredPaymentDetailsConfig {
	r.merchantAccount = &merchantAccount
	return r
}

/*
GetTokensForStoredPaymentDetails Get tokens for stored payment details

Lists the tokens for stored payment details for the shopper identified in the path, if there are any available. The token ID can be used with payment requests for the shopper's payment. A summary of the stored details is included.



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RecurringApiGetTokensForStoredPaymentDetailsConfig
*/
func (a *RecurringApi) GetTokensForStoredPaymentDetailsConfig(ctx context.Context) RecurringApiGetTokensForStoredPaymentDetailsConfig {
	return RecurringApiGetTokensForStoredPaymentDetailsConfig{
		ctx: ctx,
	}
}

/*
Get tokens for stored payment details
Lists the tokens for stored payment details for the shopper identified in the path, if there are any available. The token ID can be used with payment requests for the shopper&#39;s payment. A summary of the stored details is included.
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ListStoredPaymentMethodsResponse
*/

func (a *RecurringApi) GetTokensForStoredPaymentDetails(r RecurringApiGetTokensForStoredPaymentDetailsConfig) (ListStoredPaymentMethodsResponse, *_nethttp.Response, error) {
	res := &ListStoredPaymentMethodsResponse{}
	path := "/storedPaymentMethods"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.shopperReference != nil {
		common.ParameterAddToQuery(queryParams, "shopperReference", r.shopperReference, "")
	}
	if r.merchantAccount != nil {
		common.ParameterAddToQuery(queryParams, "merchantAccount", r.merchantAccount, "")
	}
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		nil,
		res,
		_nethttp.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
