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

// TerminalSettingsTerminalLevelApi TerminalSettingsTerminalLevelApi service
type TerminalSettingsTerminalLevelApi common.Service

type TerminalSettingsTerminalLevelApiGetTerminalLogoConfig struct {
	ctx        context.Context
	terminalId string
}

/*
GetTerminalLogo Get the terminal logo

Returns the logo that is configured for the payment terminal identified in the path.
The logo is returned as a Base64-encoded string. You need to Base64-decode the string to get the actual image file.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param terminalId The unique identifier of the payment terminal.
 @return TerminalSettingsTerminalLevelApiGetTerminalLogoConfig
*/
func (a *TerminalSettingsTerminalLevelApi) GetTerminalLogoConfig(ctx context.Context, terminalId string) TerminalSettingsTerminalLevelApiGetTerminalLogoConfig {
	return TerminalSettingsTerminalLevelApiGetTerminalLogoConfig{
		ctx:        ctx,
		terminalId: terminalId,
	}
}

/*
Get the terminal logo
Returns the logo that is configured for the payment terminal identified in the path. The logo is returned as a Base64-encoded string. You need to Base64-decode the string to get the actual image file.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read * Management API—Terminal settings read and write
 * @param terminalId The unique identifier of the payment terminal.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Logo
*/

func (a *TerminalSettingsTerminalLevelApi) GetTerminalLogo(r TerminalSettingsTerminalLevelApiGetTerminalLogoConfig) (Logo, *_nethttp.Response, error) {
	res := &Logo{}
	path := "/terminals/{terminalId}/terminalLogos"
	path = strings.Replace(path, "{"+"terminalId"+"}", url.PathEscape(common.ParameterValueToString(r.terminalId, "terminalId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsTerminalLevelApiGetTerminalSettingsConfig struct {
	ctx        context.Context
	terminalId string
}

/*
GetTerminalSettings Get terminal settings

Returns the settings that are configured for the payment terminal identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param terminalId The unique identifier of the payment terminal.
 @return TerminalSettingsTerminalLevelApiGetTerminalSettingsConfig
*/
func (a *TerminalSettingsTerminalLevelApi) GetTerminalSettingsConfig(ctx context.Context, terminalId string) TerminalSettingsTerminalLevelApiGetTerminalSettingsConfig {
	return TerminalSettingsTerminalLevelApiGetTerminalSettingsConfig{
		ctx:        ctx,
		terminalId: terminalId,
	}
}

/*
Get terminal settings
Returns the settings that are configured for the payment terminal identified in the path.  To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read * Management API—Terminal settings read and write  For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role: * Management API—Terminal settings Advanced read and write
 * @param terminalId The unique identifier of the payment terminal.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalSettings
*/

func (a *TerminalSettingsTerminalLevelApi) GetTerminalSettings(r TerminalSettingsTerminalLevelApiGetTerminalSettingsConfig) (TerminalSettings, *_nethttp.Response, error) {
	res := &TerminalSettings{}
	path := "/terminals/{terminalId}/terminalSettings"
	path = strings.Replace(path, "{"+"terminalId"+"}", url.PathEscape(common.ParameterValueToString(r.terminalId, "terminalId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodGet, nil, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsTerminalLevelApiUpdateLogoConfig struct {
	ctx        context.Context
	terminalId string
	logo       *Logo
}

func (r TerminalSettingsTerminalLevelApiUpdateLogoConfig) Logo(logo Logo) TerminalSettingsTerminalLevelApiUpdateLogoConfig {
	r.logo = &logo
	return r
}

/*
UpdateLogo Update the logo

Updates the logo for the payment terminal identified in the path.

* To change the logo, specify the image file as a Base64-encoded string.
* To restore the logo inherited from a higher level (store, merchant account, or company account), specify an empty logo value.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param terminalId The unique identifier of the payment terminal.
 @return TerminalSettingsTerminalLevelApiUpdateLogoConfig
*/
func (a *TerminalSettingsTerminalLevelApi) UpdateLogoConfig(ctx context.Context, terminalId string) TerminalSettingsTerminalLevelApiUpdateLogoConfig {
	return TerminalSettingsTerminalLevelApiUpdateLogoConfig{
		ctx:        ctx,
		terminalId: terminalId,
	}
}

/*
Update the logo
Updates the logo for the payment terminal identified in the path.  * To change the logo, specify the image file as a Base64-encoded string. * To restore the logo inherited from a higher level (store, merchant account, or company account), specify an empty logo value.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read and write
 * @param terminalId The unique identifier of the payment terminal.
 * @param req Logo - reference of Logo).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return Logo
*/

func (a *TerminalSettingsTerminalLevelApi) UpdateLogo(r TerminalSettingsTerminalLevelApiUpdateLogoConfig) (Logo, *_nethttp.Response, error) {
	res := &Logo{}
	path := "/terminals/{terminalId}/terminalLogos"
	path = strings.Replace(path, "{"+"terminalId"+"}", url.PathEscape(common.ParameterValueToString(r.terminalId, "terminalId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.logo, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}

type TerminalSettingsTerminalLevelApiUpdateTerminalSettingsConfig struct {
	ctx              context.Context
	terminalId       string
	terminalSettings *TerminalSettings
}

func (r TerminalSettingsTerminalLevelApiUpdateTerminalSettingsConfig) TerminalSettings(terminalSettings TerminalSettings) TerminalSettingsTerminalLevelApiUpdateTerminalSettingsConfig {
	r.terminalSettings = &terminalSettings
	return r
}

/*
UpdateTerminalSettings Update terminal settings

Updates the settings that are configured for the payment terminal identified in the path.

* To change a parameter value, include the full object that contains the parameter, even if you don't want to change all parameters in the object.
* To restore a parameter value inherited from a higher level, include the full object that contains the parameter, and specify an empty value for the parameter or omit the parameter.
* Objects that are not included in the request are not updated.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param terminalId The unique identifier of the payment terminal.
 @return TerminalSettingsTerminalLevelApiUpdateTerminalSettingsConfig
*/
func (a *TerminalSettingsTerminalLevelApi) UpdateTerminalSettingsConfig(ctx context.Context, terminalId string) TerminalSettingsTerminalLevelApiUpdateTerminalSettingsConfig {
	return TerminalSettingsTerminalLevelApiUpdateTerminalSettingsConfig{
		ctx:        ctx,
		terminalId: terminalId,
	}
}

/*
Update terminal settings
Updates the settings that are configured for the payment terminal identified in the path.  * To change a parameter value, include the full object that contains the parameter, even if you don&#39;t want to change all parameters in the object. * To restore a parameter value inherited from a higher level, include the full object that contains the parameter, and specify an empty value for the parameter or omit the parameter. * Objects that are not included in the request are not updated.  To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions): * Management API—Terminal settings read and write  For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role: * Management API—Terminal settings Advanced read and write
 * @param terminalId The unique identifier of the payment terminal.
 * @param req TerminalSettings - reference of TerminalSettings).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return TerminalSettings
*/

func (a *TerminalSettingsTerminalLevelApi) UpdateTerminalSettings(r TerminalSettingsTerminalLevelApiUpdateTerminalSettingsConfig) (TerminalSettings, *_nethttp.Response, error) {
	res := &TerminalSettings{}
	path := "/terminals/{terminalId}/terminalSettings"
	path = strings.Replace(path, "{"+"terminalId"+"}", url.PathEscape(common.ParameterValueToString(r.terminalId, "terminalId")), -1)
	httpRes, err := common.CreateHTTPRequest(a.Client, _nethttp.MethodPatch, r.terminalSettings, res, a.BasePath()+path, []_context.Context{r.ctx})
	return *res, httpRes, err
}
