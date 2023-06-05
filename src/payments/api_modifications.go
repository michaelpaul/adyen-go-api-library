/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"context"
	_nethttp "net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// ModificationsApi ModificationsApi service
type ModificationsApi common.Service

type ModificationsApiAdjustAuthorisationConfig struct {
	ctx                        context.Context
	adjustAuthorisationRequest *AdjustAuthorisationRequest
}

func (r ModificationsApiAdjustAuthorisationConfig) AdjustAuthorisationRequest(adjustAuthorisationRequest AdjustAuthorisationRequest) ModificationsApiAdjustAuthorisationConfig {
	r.adjustAuthorisationRequest = &adjustAuthorisationRequest
	return r
}

/*
AdjustAuthorisation Change the authorised amount

Allows you to increase or decrease the authorised amount after the initial authorisation has taken place. This functionality enables for example tipping, improving the chances your authorisation will be valid, or charging the shopper when they have already left the merchant premises.

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce).
> If you have a [newer integration](https://docs.adyen.com/online-payments), and are doing:
> * [Asynchronous adjustments](https://docs.adyen.com/online-payments/adjust-authorisation#asynchronous-or-synchronous-adjustment), use the [`/payments/{paymentPspReference}/amountUpdates`](https://docs.adyen.com/api-explorer/#/CheckoutService/v67/post/payments/{paymentPspReference}/amountUpdates) endpoint on Checkout API.
> * [Synchronous adjustments](https://docs.adyen.com/online-payments/adjust-authorisation#asynchronous-or-synchronous-adjustment), use this endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ModificationsApiAdjustAuthorisationConfig
*/
func (a *ModificationsApi) AdjustAuthorisationConfig(ctx context.Context) ModificationsApiAdjustAuthorisationConfig {
	return ModificationsApiAdjustAuthorisationConfig{
		ctx: ctx,
	}
}

/*
Change the authorised amount
Allows you to increase or decrease the authorised amount after the initial authorisation has taken place. This functionality enables for example tipping, improving the chances your authorisation will be valid, or charging the shopper when they have already left the merchant premises.  &gt; This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). &gt; If you have a [newer integration](https://docs.adyen.com/online-payments), and are doing: &gt; * [Asynchronous adjustments](https://docs.adyen.com/online-payments/adjust-authorisation#asynchronous-or-synchronous-adjustment), use the [&#x60;/payments/{paymentPspReference}/amountUpdates&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/v67/post/payments/{paymentPspReference}/amountUpdates) endpoint on Checkout API. &gt; * [Synchronous adjustments](https://docs.adyen.com/online-payments/adjust-authorisation#asynchronous-or-synchronous-adjustment), use this endpoint.
 * @param req AdjustAuthorisationRequest - reference of AdjustAuthorisationRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ModificationResult
*/

func (a *ModificationsApi) AdjustAuthorisation(r ModificationsApiAdjustAuthorisationConfig) (ModificationResult, *_nethttp.Response, error) {
	res := &ModificationResult{}
	path := "/adjustAuthorisation"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.adjustAuthorisationRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type ModificationsApiCancelConfig struct {
	ctx           context.Context
	cancelRequest *CancelRequest
}

func (r ModificationsApiCancelConfig) CancelRequest(cancelRequest CancelRequest) ModificationsApiCancelConfig {
	r.cancelRequest = &cancelRequest
	return r
}

/*
Cancel Cancel an authorisation

Cancels the authorisation hold on a payment, returning a unique reference for this request. You can cancel payments after authorisation only for payment methods that support distinct authorisations and captures.

For more information, refer to [Cancel](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/cancel).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/{paymentPspReference}/cancels`](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/{paymentPspReference}/cancels) endpoint under Checkout API instead.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ModificationsApiCancelConfig
*/
func (a *ModificationsApi) CancelConfig(ctx context.Context) ModificationsApiCancelConfig {
	return ModificationsApiCancelConfig{
		ctx: ctx,
	}
}

/*
Cancel an authorisation
Cancels the authorisation hold on a payment, returning a unique reference for this request. You can cancel payments after authorisation only for payment methods that support distinct authorisations and captures.  For more information, refer to [Cancel](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/cancel).  &gt; This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [&#x60;/payments/{paymentPspReference}/cancels&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/{paymentPspReference}/cancels) endpoint under Checkout API instead.
 * @param req CancelRequest - reference of CancelRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ModificationResult
*/

func (a *ModificationsApi) Cancel(r ModificationsApiCancelConfig) (ModificationResult, *_nethttp.Response, error) {
	res := &ModificationResult{}
	path := "/cancel"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.cancelRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type ModificationsApiCancelOrRefundConfig struct {
	ctx                   context.Context
	cancelOrRefundRequest *CancelOrRefundRequest
}

func (r ModificationsApiCancelOrRefundConfig) CancelOrRefundRequest(cancelOrRefundRequest CancelOrRefundRequest) ModificationsApiCancelOrRefundConfig {
	r.cancelOrRefundRequest = &cancelOrRefundRequest
	return r
}

/*
CancelOrRefund Cancel or refund a payment

Cancels a payment if it has not been captured yet, or refunds it if it has already been captured. This is useful when it is not certain if the payment has been captured or not (for example, when using auto-capture).

Do not use this endpoint for payments that involve:
 * [Multiple partial captures](https://docs.adyen.com/online-payments/capture).
 * [Split data](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information) either at time of payment or capture for Adyen for Platforms.

 Instead, check if the payment has been captured and make a corresponding [`/refund`](https://docs.adyen.com/api-explorer/#/Payment/refund) or [`/cancel`](https://docs.adyen.com/api-explorer/#/Payment/cancel) call.

For more information, refer to [Cancel or refund](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/cancel-or-refund).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/{paymentPspReference}/reversals`](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/{paymentPspReference}/reversals) endpoint under Checkout API instead.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ModificationsApiCancelOrRefundConfig
*/
func (a *ModificationsApi) CancelOrRefundConfig(ctx context.Context) ModificationsApiCancelOrRefundConfig {
	return ModificationsApiCancelOrRefundConfig{
		ctx: ctx,
	}
}

/*
Cancel or refund a payment
Cancels a payment if it has not been captured yet, or refunds it if it has already been captured. This is useful when it is not certain if the payment has been captured or not (for example, when using auto-capture).  Do not use this endpoint for payments that involve:  * [Multiple partial captures](https://docs.adyen.com/online-payments/capture).  * [Split data](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information) either at time of payment or capture for Adyen for Platforms.   Instead, check if the payment has been captured and make a corresponding [&#x60;/refund&#x60;](https://docs.adyen.com/api-explorer/#/Payment/refund) or [&#x60;/cancel&#x60;](https://docs.adyen.com/api-explorer/#/Payment/cancel) call.  For more information, refer to [Cancel or refund](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/cancel-or-refund).  &gt; This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [&#x60;/payments/{paymentPspReference}/reversals&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/{paymentPspReference}/reversals) endpoint under Checkout API instead.
 * @param req CancelOrRefundRequest - reference of CancelOrRefundRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ModificationResult
*/

func (a *ModificationsApi) CancelOrRefund(r ModificationsApiCancelOrRefundConfig) (ModificationResult, *_nethttp.Response, error) {
	res := &ModificationResult{}
	path := "/cancelOrRefund"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.cancelOrRefundRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type ModificationsApiCaptureConfig struct {
	ctx            context.Context
	captureRequest *CaptureRequest
}

func (r ModificationsApiCaptureConfig) CaptureRequest(captureRequest CaptureRequest) ModificationsApiCaptureConfig {
	r.captureRequest = &captureRequest
	return r
}

/*
Capture Capture an authorisation

Captures the authorisation hold on a payment, returning a unique reference for this request. Usually the full authorisation amount is captured, however it's also possible to capture a smaller amount, which results in cancelling the remaining authorisation balance.

Payment methods that are captured automatically after authorisation don't need to be captured. However, submitting a capture request on these transactions will not result in double charges. If immediate or delayed auto-capture is enabled, calling the capture method is not neccessary.

For more information refer to [Capture](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/capture).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/{paymentPspReference}/captures`](https://docs.adyen.com/api-explorer/#/CheckoutService/v67/post/payments/{paymentPspReference}/captures) endpoint on Checkout API instead.



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ModificationsApiCaptureConfig
*/
func (a *ModificationsApi) CaptureConfig(ctx context.Context) ModificationsApiCaptureConfig {
	return ModificationsApiCaptureConfig{
		ctx: ctx,
	}
}

/*
Capture an authorisation
Captures the authorisation hold on a payment, returning a unique reference for this request. Usually the full authorisation amount is captured, however it&#39;s also possible to capture a smaller amount, which results in cancelling the remaining authorisation balance.  Payment methods that are captured automatically after authorisation don&#39;t need to be captured. However, submitting a capture request on these transactions will not result in double charges. If immediate or delayed auto-capture is enabled, calling the capture method is not neccessary.  For more information refer to [Capture](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/capture).  &gt; This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [&#x60;/payments/{paymentPspReference}/captures&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/v67/post/payments/{paymentPspReference}/captures) endpoint on Checkout API instead.
 * @param req CaptureRequest - reference of CaptureRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ModificationResult
*/

func (a *ModificationsApi) Capture(r ModificationsApiCaptureConfig) (ModificationResult, *_nethttp.Response, error) {
	res := &ModificationResult{}
	path := "/capture"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.captureRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type ModificationsApiDonateConfig struct {
	ctx             context.Context
	donationRequest *DonationRequest
}

func (r ModificationsApiDonateConfig) DonationRequest(donationRequest DonationRequest) ModificationsApiDonateConfig {
	r.donationRequest = &donationRequest
	return r
}

/*
Donate Create a donation

Schedules a new payment to be created (including a new authorisation request) for the specified donation using the payment details of the original payment.

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/donations`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/donations) endpoint under Checkout API instead.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ModificationsApiDonateConfig
*/
func (a *ModificationsApi) DonateConfig(ctx context.Context) ModificationsApiDonateConfig {
	return ModificationsApiDonateConfig{
		ctx: ctx,
	}
}

/*
Create a donation
Schedules a new payment to be created (including a new authorisation request) for the specified donation using the payment details of the original payment.  &gt; This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [&#x60;/donations&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/donations) endpoint under Checkout API instead.
 * @param req DonationRequest - reference of DonationRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ModificationResult
*/

func (a *ModificationsApi) Donate(r ModificationsApiDonateConfig) (ModificationResult, *_nethttp.Response, error) {
	res := &ModificationResult{}
	path := "/donate"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.donationRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type ModificationsApiRefundConfig struct {
	ctx           context.Context
	refundRequest *RefundRequest
}

func (r ModificationsApiRefundConfig) RefundRequest(refundRequest RefundRequest) ModificationsApiRefundConfig {
	r.refundRequest = &refundRequest
	return r
}

/*
Refund Refund a captured payment

Refunds a payment that has previously been captured, returning a unique reference for this request. Refunding can be done on the full captured amount or a partial amount. Multiple (partial) refunds will be accepted as long as their sum doesn't exceed the captured amount. Payments which have been authorised, but not captured, cannot be refunded, use the /cancel method instead.

Some payment methods/gateways do not support partial/multiple refunds.
A margin above the captured limit can be configured to cover shipping/handling costs.

For more information, refer to [Refund](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/refund).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/payments/{paymentPspReference}/refunds`](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/{paymentPspReference}/refunds) endpoint under Checkout API instead.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ModificationsApiRefundConfig
*/
func (a *ModificationsApi) RefundConfig(ctx context.Context) ModificationsApiRefundConfig {
	return ModificationsApiRefundConfig{
		ctx: ctx,
	}
}

/*
Refund a captured payment
Refunds a payment that has previously been captured, returning a unique reference for this request. Refunding can be done on the full captured amount or a partial amount. Multiple (partial) refunds will be accepted as long as their sum doesn&#39;t exceed the captured amount. Payments which have been authorised, but not captured, cannot be refunded, use the /cancel method instead.  Some payment methods/gateways do not support partial/multiple refunds. A margin above the captured limit can be configured to cover shipping/handling costs.  For more information, refer to [Refund](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/refund).  &gt; This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [&#x60;/payments/{paymentPspReference}/refunds&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/payments/{paymentPspReference}/refunds) endpoint under Checkout API instead.
 * @param req RefundRequest - reference of RefundRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ModificationResult
*/

func (a *ModificationsApi) Refund(r ModificationsApiRefundConfig) (ModificationResult, *_nethttp.Response, error) {
	res := &ModificationResult{}
	path := "/refund"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.refundRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type ModificationsApiTechnicalCancelConfig struct {
	ctx                    context.Context
	technicalCancelRequest *TechnicalCancelRequest
}

func (r ModificationsApiTechnicalCancelConfig) TechnicalCancelRequest(technicalCancelRequest TechnicalCancelRequest) ModificationsApiTechnicalCancelConfig {
	r.technicalCancelRequest = &technicalCancelRequest
	return r
}

/*
TechnicalCancel Cancel an authorisation using your reference

This endpoint allows you to cancel a payment if you do not have the PSP reference of the original payment request available.

In your call, refer to the original payment by using the `reference` that you specified in your payment request.

For more information, see [Technical cancel](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/cancel#technical-cancel).

> This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [`/cancels`](https://docs.adyen.com/api-explorer/#/CheckoutService/cancels) endpoint under Checkout API instead.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ModificationsApiTechnicalCancelConfig
*/
func (a *ModificationsApi) TechnicalCancelConfig(ctx context.Context) ModificationsApiTechnicalCancelConfig {
	return ModificationsApiTechnicalCancelConfig{
		ctx: ctx,
	}
}

/*
Cancel an authorisation using your reference
This endpoint allows you to cancel a payment if you do not have the PSP reference of the original payment request available.  In your call, refer to the original payment by using the &#x60;reference&#x60; that you specified in your payment request.  For more information, see [Technical cancel](https://docs.adyen.com/online-payments/classic-integrations/modify-payments/cancel#technical-cancel).   &gt; This endpoint is part of our [classic API integration](https://docs.adyen.com/online-payments/classic-integrations/api-integration-ecommerce). If using a [newer integration](https://docs.adyen.com/online-payments), use the [&#x60;/cancels&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/cancels) endpoint under Checkout API instead.
 * @param req TechnicalCancelRequest - reference of TechnicalCancelRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ModificationResult
*/

func (a *ModificationsApi) TechnicalCancel(r ModificationsApiTechnicalCancelConfig) (ModificationResult, *_nethttp.Response, error) {
	res := &ModificationResult{}
	path := "/technicalCancel"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.technicalCancelRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

type ModificationsApiVoidPendingRefundConfig struct {
	ctx                      context.Context
	voidPendingRefundRequest *VoidPendingRefundRequest
}

func (r ModificationsApiVoidPendingRefundConfig) VoidPendingRefundRequest(voidPendingRefundRequest VoidPendingRefundRequest) ModificationsApiVoidPendingRefundConfig {
	r.voidPendingRefundRequest = &voidPendingRefundRequest
	return r
}

/*
VoidPendingRefund Cancel an in-person refund

This endpoint allows you to cancel an unreferenced refund request before it has been completed.

In your call, you can refer to the original refund request either by using the `tenderReference`, or the `pspReference`. We recommend implementing based on the `tenderReference`, as this is generated for both offline and online transactions.

For more information, refer to [Cancel an unreferenced refund](https://docs.adyen.com/point-of-sale/refund-payment/cancel-unreferenced).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ModificationsApiVoidPendingRefundConfig
*/
func (a *ModificationsApi) VoidPendingRefundConfig(ctx context.Context) ModificationsApiVoidPendingRefundConfig {
	return ModificationsApiVoidPendingRefundConfig{
		ctx: ctx,
	}
}

/*
Cancel an in-person refund
This endpoint allows you to cancel an unreferenced refund request before it has been completed.  In your call, you can refer to the original refund request either by using the &#x60;tenderReference&#x60;, or the &#x60;pspReference&#x60;. We recommend implementing based on the &#x60;tenderReference&#x60;, as this is generated for both offline and online transactions.  For more information, refer to [Cancel an unreferenced refund](https://docs.adyen.com/point-of-sale/refund-payment/cancel-unreferenced).
 * @param req VoidPendingRefundRequest - reference of VoidPendingRefundRequest).
 * @param ctxs ...context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ModificationResult
*/

func (a *ModificationsApi) VoidPendingRefund(r ModificationsApiVoidPendingRefundConfig) (ModificationResult, *_nethttp.Response, error) {
	res := &ModificationResult{}
	path := "/voidPendingRefund"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		r.ctx,
		a.Client,
		r.voidPendingRefundRequest,
		res,
		_nethttp.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
