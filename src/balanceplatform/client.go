/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
    "github.com/adyen/adyen-go-api-library/v7/src/common"
)

// APIClient manages communication with the Configuration API API v2
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	common common.Service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccountHoldersApi *AccountHoldersApi

	BalanceAccountsApi *BalanceAccountsApi

	BankAccountValidationApi *BankAccountValidationApi

	GrantAccountsApi *GrantAccountsApi

	GrantOffersApi *GrantOffersApi

	NetworkTokensApi *NetworkTokensApi

	PaymentInstrumentGroupsApi *PaymentInstrumentGroupsApi

	PaymentInstrumentsApi *PaymentInstrumentsApi

	PlatformApi *PlatformApi

	TransactionRulesApi *TransactionRulesApi
}

// NewAPIClient creates a new API client.
func NewAPIClient(client *common.Client) *APIClient {
	c := &APIClient{}
    c.common.Client = client
    c.common.BasePath = func() string {
        return client.Cfg.BalancePlatformEndpoint
    }

	// API Services
	c.AccountHoldersApi = (*AccountHoldersApi)(&c.common)
	c.BalanceAccountsApi = (*BalanceAccountsApi)(&c.common)
	c.BankAccountValidationApi = (*BankAccountValidationApi)(&c.common)
	c.GrantAccountsApi = (*GrantAccountsApi)(&c.common)
	c.GrantOffersApi = (*GrantOffersApi)(&c.common)
	c.NetworkTokensApi = (*NetworkTokensApi)(&c.common)
	c.PaymentInstrumentGroupsApi = (*PaymentInstrumentGroupsApi)(&c.common)
	c.PaymentInstrumentsApi = (*PaymentInstrumentsApi)(&c.common)
	c.PlatformApi = (*PlatformApi)(&c.common)
	c.TransactionRulesApi = (*TransactionRulesApi)(&c.common)

	return c
}