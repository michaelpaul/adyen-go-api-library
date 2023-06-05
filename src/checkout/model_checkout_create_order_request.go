/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the CheckoutCreateOrderRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutCreateOrderRequest{}

// CheckoutCreateOrderRequest struct for CheckoutCreateOrderRequest
type CheckoutCreateOrderRequest struct {
	Amount Amount `json:"amount"`
	// The date that order expires; e.g. 2019-03-23T12:25:28Z. If not provided, the default expiry duration is 1 day.
	ExpiresAt *string `json:"expiresAt,omitempty"`
	// The merchant account identifier, with which you want to process the order.
	MerchantAccount string `json:"merchantAccount"`
	// A custom reference identifying the order.
	Reference string `json:"reference"`
}

// NewCheckoutCreateOrderRequest instantiates a new CheckoutCreateOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutCreateOrderRequest(amount Amount, merchantAccount string, reference string) *CheckoutCreateOrderRequest {
	this := CheckoutCreateOrderRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	this.Reference = reference
	return &this
}

// NewCheckoutCreateOrderRequestWithDefaults instantiates a new CheckoutCreateOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutCreateOrderRequestWithDefaults() *CheckoutCreateOrderRequest {
	this := CheckoutCreateOrderRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CheckoutCreateOrderRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CheckoutCreateOrderRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CheckoutCreateOrderRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CheckoutCreateOrderRequest) GetExpiresAt() string {
	if o == nil || common.IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutCreateOrderRequest) GetExpiresAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CheckoutCreateOrderRequest) HasExpiresAt() bool {
	if o != nil && !common.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CheckoutCreateOrderRequest) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *CheckoutCreateOrderRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *CheckoutCreateOrderRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *CheckoutCreateOrderRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetReference returns the Reference field value
func (o *CheckoutCreateOrderRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *CheckoutCreateOrderRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *CheckoutCreateOrderRequest) SetReference(v string) {
	o.Reference = v
}

func (o CheckoutCreateOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutCreateOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["reference"] = o.Reference
	return toSerialize, nil
}

type NullableCheckoutCreateOrderRequest struct {
	value *CheckoutCreateOrderRequest
	isSet bool
}

func (v NullableCheckoutCreateOrderRequest) Get() *CheckoutCreateOrderRequest {
	return v.value
}

func (v *NullableCheckoutCreateOrderRequest) Set(val *CheckoutCreateOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutCreateOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutCreateOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutCreateOrderRequest(val *CheckoutCreateOrderRequest) *NullableCheckoutCreateOrderRequest {
	return &NullableCheckoutCreateOrderRequest{value: val, isSet: true}
}

func (v NullableCheckoutCreateOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutCreateOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
