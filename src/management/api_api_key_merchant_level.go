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

// APIKeyMerchantLevelApi APIKeyMerchantLevelApi service
type APIKeyMerchantLevelApi common.Service

type APIKeyMerchantLevelApiGenerateNewApiKeyConfig struct {
	ctx             context.Context
	merchantId      string
	apiCredentialId string
}

/*
GenerateNewApiKey Generate new API key

Returns a new API key for the API credential. You can use the new API key a few minutes after generating it. The old API key stops working 24 hours after generating a new one.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param merchantId The unique identifier of the merchant account.
 @param apiCredentialId Unique identifier of the API credential.
 @return APIKeyMerchantLevelApiGenerateNewApiKeyConfig
*/
func (a *APIKeyMerchantLevelApi) GenerateNewApiKeyConfig(ctx context.Context, merchantId string, apiCredentialId string) APIKeyMerchantLevelApiGenerateNewApiKeyConfig {
	return APIKeyMerchantLevelApiGenerateNewApiKeyConfig{
		ctx:             ctx,
		merchantId:      merchantId,
		apiCredentialId: apiCredentialId,
	}
}

/*
Generate new API key
Returns a new API key for the API credential. You can use the new API key a few minutes after generating it. The old API key stops working 24 hours after generating a new one.  To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—API credentials read and write
 * @param merchantId The unique identifier of the merchant account.
 * @param apiCredentialId Unique identifier of the API credential.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return GenerateApiKeyResponse
*/

func (a *APIKeyMerchantLevelApi) GenerateNewApiKey(r APIKeyMerchantLevelApiGenerateNewApiKeyConfig) (GenerateApiKeyResponse, *_nethttp.Response, error) {
	res := &GenerateApiKeyResponse{}
	path := "/merchants/{merchantId}/apiCredentials/{apiCredentialId}/generateApiKey"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPost, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}
