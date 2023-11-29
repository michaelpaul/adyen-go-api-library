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

// BankAccountDetail struct for BankAccountDetail
type BankAccountDetail struct {
	// The bank account number (without separators). >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	AccountNumber string `json:"accountNumber,omitempty"`
	// The type of bank account. Only applicable to bank accounts held in the USA. The permitted values are: `checking`, `savings`.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	AccountType string `json:"accountType,omitempty"`
	// The name of the bank account.
	BankAccountName string `json:"bankAccountName,omitempty"`
	// Merchant reference to the bank account.
	BankAccountReference string `json:"bankAccountReference,omitempty"`
	// The unique identifier (UUID) of the Bank Account. >If, during an account holder create or update request, this field is left blank (but other fields provided), a new Bank Account will be created with a procedurally-generated UUID.  >If, during an account holder create request, a UUID is provided, the creation of the Bank Account will fail while the creation of the account holder will continue.  >If, during an account holder update request, a UUID that is not correlated with an existing Bank Account is provided, the update of the account holder will fail.  >If, during an account holder update request, a UUID that is correlated with an existing Bank Account is provided, the existing Bank Account will be updated.
	BankAccountUUID string `json:"bankAccountUUID,omitempty"`
	// The bank identifier code. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BankBicSwift string `json:"bankBicSwift,omitempty"`
	// The city in which the bank branch is located.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BankCity string `json:"bankCity,omitempty"`
	// The bank code of the banking institution with which the bank account is registered.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BankCode string `json:"bankCode,omitempty"`
	// The name of the banking institution with which the bank account is held.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BankName string `json:"bankName,omitempty"`
	// The branch code of the branch under which the bank account is registered. The value to be specified in this parameter depends on the country of the bank account: * United States - Routing number * United Kingdom - Sort code * Germany - Bankleitzahl >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	BranchCode string `json:"branchCode,omitempty"`
	// The check code of the bank account.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	CheckCode string `json:"checkCode,omitempty"`
	// The two-letter country code in which the bank account is registered. >The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL').  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	CountryCode string `json:"countryCode,omitempty"`
	// The currency in which the bank account deals. >The permitted currency codes are defined in ISO-4217 (e.g. 'EUR').  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	CurrencyCode string `json:"currencyCode,omitempty"`
	// The international bank account number. >The IBAN standard is defined in ISO-13616.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	Iban string `json:"iban,omitempty"`
	// The city of residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerCity string `json:"ownerCity,omitempty"`
	// The country code of the country of residence of the bank account owner. >The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL').  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerCountryCode string `json:"ownerCountryCode,omitempty"`
	// The date of birth of the bank account owner.
	OwnerDateOfBirth string `json:"ownerDateOfBirth,omitempty"`
	// The house name or number of the residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerHouseNumberOrName string `json:"ownerHouseNumberOrName,omitempty"`
	// The name of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerName string `json:"ownerName,omitempty"`
	// The country code of the country of nationality of the bank account owner. >The permitted country codes are defined in ISO-3166-1 alpha-2 (e.g. 'NL').  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerNationality string `json:"ownerNationality,omitempty"`
	// The postal code of the residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerPostalCode string `json:"ownerPostalCode,omitempty"`
	// The state of residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerState string `json:"ownerState,omitempty"`
	// The street name of the residence of the bank account owner. >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	OwnerStreet string `json:"ownerStreet,omitempty"`
	// If set to true, the bank account is a primary account.
	PrimaryAccount bool `json:"primaryAccount,omitempty"`
	// The tax ID number.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	TaxId string `json:"taxId,omitempty"`
	// The URL to be used for bank account verification. This may be generated on bank account creation.  >Refer to the [Onboarding and verification](https://docs.adyen.com/platforms/onboarding-and-verification) section for details on field requirements.
	UrlForVerification string `json:"urlForVerification,omitempty"`
}
