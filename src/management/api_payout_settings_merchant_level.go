/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	_context "context"
	_nethttp "net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// PayoutSettingsMerchantLevelApi PayoutSettingsMerchantLevelApi service
type PayoutSettingsMerchantLevelApi common.Service

type PayoutSettingsMerchantLevelApiAddPayoutSettingConfig struct {
	ctx                   context.Context
	merchantId            string
	payoutSettingsRequest *PayoutSettingsRequest
}

func (r PayoutSettingsMerchantLevelApiAddPayoutSettingConfig) PayoutSettingsRequest(payoutSettingsRequest PayoutSettingsRequest) PayoutSettingsMerchantLevelApiAddPayoutSettingConfig {
	r.payoutSettingsRequest = &payoutSettingsRequest
	return r
}

/*
AddPayoutSetting Add a payout setting

Sends a request to add a payout setting for the merchant account specified in the path. A payout setting links the merchant account to the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments) that contains the details of the payout bank account. Adyen verifies the bank account before allowing and enabling the payout setting.

If you're accepting payments in multiple currencies, you may add multiple payout settings for the merchant account.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):

* Management API—Payout account settings read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @return PayoutSettingsMerchantLevelApiAddPayoutSettingConfig
*/
func (a *PayoutSettingsMerchantLevelApi) AddPayoutSettingConfig(ctx context.Context, merchantId string) PayoutSettingsMerchantLevelApiAddPayoutSettingConfig {
	return PayoutSettingsMerchantLevelApiAddPayoutSettingConfig{
		ctx:        ctx,
		merchantId: merchantId,
	}
}

/*
Add a payout setting
Sends a request to add a payout setting for the merchant account specified in the path. A payout setting links the merchant account to the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments) that contains the details of the payout bank account. Adyen verifies the bank account before allowing and enabling the payout setting.  If you&#39;re accepting payments in multiple currencies, you may add multiple payout settings for the merchant account.  Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):  * Management API—Payout account settings read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param req PayoutSettingsRequest - reference of PayoutSettingsRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PayoutSettings
*/

func (a *PayoutSettingsMerchantLevelApi) AddPayoutSetting(r PayoutSettingsMerchantLevelApiAddPayoutSettingConfig) (PayoutSettings, *_nethttp.Response, error) {
	res := &PayoutSettings{}
	path := "/merchants/{merchantId}/payoutSettings"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, r.payoutSettingsRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type PayoutSettingsMerchantLevelApiDeletePayoutSettingConfig struct {
	ctx              context.Context
	merchantId       string
	payoutSettingsId string
}

/*
DeletePayoutSetting Delete a payout setting

Deletes the payout setting identified in the path.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):

* Management API—Payout account settings read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param payoutSettingsId The unique identifier of the payout setting.
 @return PayoutSettingsMerchantLevelApiDeletePayoutSettingConfig
*/
func (a *PayoutSettingsMerchantLevelApi) DeletePayoutSettingConfig(ctx context.Context, merchantId string, payoutSettingsId string) PayoutSettingsMerchantLevelApiDeletePayoutSettingConfig {
	return PayoutSettingsMerchantLevelApiDeletePayoutSettingConfig{
		ctx:              ctx,
		merchantId:       merchantId,
		payoutSettingsId: payoutSettingsId,
	}
}

/*
Delete a payout setting
Deletes the payout setting identified in the path.  Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):  * Management API—Payout account settings read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param payoutSettingsId The unique identifier of the payout setting.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
*/

func (a *PayoutSettingsMerchantLevelApi) DeletePayoutSetting(r PayoutSettingsMerchantLevelApiDeletePayoutSettingConfig) (*_nethttp.Response, error) {
	var res interface{}
	path := "/merchants/{merchantId}/payoutSettings/{payoutSettingsId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"payoutSettingsId"+"}", url.PathEscape(common.ParameterValueToString(r.payoutSettingsId, "payoutSettingsId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodDelete, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return httpRes, err
}

type PayoutSettingsMerchantLevelApiGetPayoutSettingConfig struct {
	ctx              context.Context
	merchantId       string
	payoutSettingsId string
}

/*
GetPayoutSetting Get a payout setting

Returns the payout setting identified in the path.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payout account settings read

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param payoutSettingsId The unique identifier of the payout setting.
 @return PayoutSettingsMerchantLevelApiGetPayoutSettingConfig
*/
func (a *PayoutSettingsMerchantLevelApi) GetPayoutSettingConfig(ctx context.Context, merchantId string, payoutSettingsId string) PayoutSettingsMerchantLevelApiGetPayoutSettingConfig {
	return PayoutSettingsMerchantLevelApiGetPayoutSettingConfig{
		ctx:              ctx,
		merchantId:       merchantId,
		payoutSettingsId: payoutSettingsId,
	}
}

/*
Get a payout setting
Returns the payout setting identified in the path.  Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Payout account settings read
 * @param merchantId The unique identifier of the merchant account.
 * @param payoutSettingsId The unique identifier of the payout setting.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PayoutSettings
*/

func (a *PayoutSettingsMerchantLevelApi) GetPayoutSetting(r PayoutSettingsMerchantLevelApiGetPayoutSettingConfig) (PayoutSettings, *_nethttp.Response, error) {
	res := &PayoutSettings{}
	path := "/merchants/{merchantId}/payoutSettings/{payoutSettingsId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"payoutSettingsId"+"}", url.PathEscape(common.ParameterValueToString(r.payoutSettingsId, "payoutSettingsId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type PayoutSettingsMerchantLevelApiListPayoutSettingsConfig struct {
	ctx        context.Context
	merchantId string
}

/*
ListPayoutSettings Get a list of payout settings

Returns the list of payout settings for the merchant account identified in the path.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Payout account settings read

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @return PayoutSettingsMerchantLevelApiListPayoutSettingsConfig
*/
func (a *PayoutSettingsMerchantLevelApi) ListPayoutSettingsConfig(ctx context.Context, merchantId string) PayoutSettingsMerchantLevelApiListPayoutSettingsConfig {
	return PayoutSettingsMerchantLevelApiListPayoutSettingsConfig{
		ctx:        ctx,
		merchantId: merchantId,
	}
}

/*
Get a list of payout settings
Returns the list of payout settings for the merchant account identified in the path.  Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Payout account settings read
 * @param merchantId The unique identifier of the merchant account.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PayoutSettingsResponse
*/

func (a *PayoutSettingsMerchantLevelApi) ListPayoutSettings(r PayoutSettingsMerchantLevelApiListPayoutSettingsConfig) (PayoutSettingsResponse, *_nethttp.Response, error) {
	res := &PayoutSettingsResponse{}
	path := "/merchants/{merchantId}/payoutSettings"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type PayoutSettingsMerchantLevelApiUpdatePayoutSettingConfig struct {
	ctx                         context.Context
	merchantId                  string
	payoutSettingsId            string
	updatePayoutSettingsRequest *UpdatePayoutSettingsRequest
}

func (r PayoutSettingsMerchantLevelApiUpdatePayoutSettingConfig) UpdatePayoutSettingsRequest(updatePayoutSettingsRequest UpdatePayoutSettingsRequest) PayoutSettingsMerchantLevelApiUpdatePayoutSettingConfig {
	r.updatePayoutSettingsRequest = &updatePayoutSettingsRequest
	return r
}

/*
UpdatePayoutSetting Update a payout setting

Updates the payout setting identified in the path. You can enable or disable the payout setting.

Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):

* Management API—Payout account settings read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param payoutSettingsId The unique identifier of the payout setting.
 @return PayoutSettingsMerchantLevelApiUpdatePayoutSettingConfig
*/
func (a *PayoutSettingsMerchantLevelApi) UpdatePayoutSettingConfig(ctx context.Context, merchantId string, payoutSettingsId string) PayoutSettingsMerchantLevelApiUpdatePayoutSettingConfig {
	return PayoutSettingsMerchantLevelApiUpdatePayoutSettingConfig{
		ctx:              ctx,
		merchantId:       merchantId,
		payoutSettingsId: payoutSettingsId,
	}
}

/*
Update a payout setting
Updates the payout setting identified in the path. You can enable or disable the payout setting.  Use this endpoint if your integration requires it, such as Adyen for Platforms Manage. Your Adyen contact will set up your access.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):  * Management API—Payout account settings read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param payoutSettingsId The unique identifier of the payout setting.
 * @param req UpdatePayoutSettingsRequest - reference of UpdatePayoutSettingsRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PayoutSettings
*/

func (a *PayoutSettingsMerchantLevelApi) UpdatePayoutSetting(r PayoutSettingsMerchantLevelApiUpdatePayoutSettingConfig) (PayoutSettings, *_nethttp.Response, error) {
	res := &PayoutSettings{}
	path := "/merchants/{merchantId}/payoutSettings/{payoutSettingsId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"payoutSettingsId"+"}", url.PathEscape(common.ParameterValueToString(r.payoutSettingsId, "payoutSettingsId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.updatePayoutSettingsRequest, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}
