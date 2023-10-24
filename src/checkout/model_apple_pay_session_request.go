/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the ApplePaySessionRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ApplePaySessionRequest{}

// ApplePaySessionRequest struct for ApplePaySessionRequest
type ApplePaySessionRequest struct {
	// This is the name that your shoppers will see in the Apple Pay interface.  The value returned as `configuration.merchantName` field from the [`/paymentMethods`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/paymentMethods) response.
	DisplayName string `json:"displayName"`
	// The domain name you provided when you added Apple Pay in your Customer Area.  This must match the `window.location.hostname` of the web shop.
	DomainName string `json:"domainName"`
	// Your merchant identifier registered with Apple Pay.  Use the value of the `configuration.merchantId` field from the [`/paymentMethods`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/paymentMethods) response.
	MerchantIdentifier string `json:"merchantIdentifier"`
}

// NewApplePaySessionRequest instantiates a new ApplePaySessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplePaySessionRequest(displayName string, domainName string, merchantIdentifier string) *ApplePaySessionRequest {
	this := ApplePaySessionRequest{}
	this.DisplayName = displayName
	this.DomainName = domainName
	this.MerchantIdentifier = merchantIdentifier
	return &this
}

// NewApplePaySessionRequestWithDefaults instantiates a new ApplePaySessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplePaySessionRequestWithDefaults() *ApplePaySessionRequest {
	this := ApplePaySessionRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *ApplePaySessionRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ApplePaySessionRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ApplePaySessionRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDomainName returns the DomainName field value
func (o *ApplePaySessionRequest) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *ApplePaySessionRequest) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *ApplePaySessionRequest) SetDomainName(v string) {
	o.DomainName = v
}

// GetMerchantIdentifier returns the MerchantIdentifier field value
func (o *ApplePaySessionRequest) GetMerchantIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantIdentifier
}

// GetMerchantIdentifierOk returns a tuple with the MerchantIdentifier field value
// and a boolean to check if the value has been set.
func (o *ApplePaySessionRequest) GetMerchantIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantIdentifier, true
}

// SetMerchantIdentifier sets field value
func (o *ApplePaySessionRequest) SetMerchantIdentifier(v string) {
	o.MerchantIdentifier = v
}

func (o ApplePaySessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplePaySessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["domainName"] = o.DomainName
	toSerialize["merchantIdentifier"] = o.MerchantIdentifier
	return toSerialize, nil
}

type NullableApplePaySessionRequest struct {
	value *ApplePaySessionRequest
	isSet bool
}

func (v NullableApplePaySessionRequest) Get() *ApplePaySessionRequest {
	return v.value
}

func (v *NullableApplePaySessionRequest) Set(val *ApplePaySessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplePaySessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplePaySessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplePaySessionRequest(val *ApplePaySessionRequest) *NullableApplePaySessionRequest {
	return &NullableApplePaySessionRequest{value: val, isSet: true}
}

func (v NullableApplePaySessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplePaySessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
