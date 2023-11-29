/*
 * Adyen for Platforms: Fund API
 *
 * The Fund API provides endpoints for managing the funds in the accounts on your platform. These management operations include actions such as the transfer of funds from one account to another, the payout of funds to an account holder, and the retrieval of balances in an account.  For more information, refer to our [documentation](https://docs.adyen.com/platforms). ## Authentication To connect to the Fund API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Fund API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Fund/v6/accountHolderBalance ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsfund

// RefundFundsTransferRequest struct for RefundFundsTransferRequest
type RefundFundsTransferRequest struct {
	Amount Amount `json:"amount"`
	// A value that can be supplied at the discretion of the executing user in order to link multiple transactions to one another.
	MerchantReference string `json:"merchantReference,omitempty"`
	// A PSP reference of the original fund transfer.
	OriginalReference string `json:"originalReference"`
}
