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

// checks if the ApplePayDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ApplePayDetails{}

// ApplePayDetails struct for ApplePayDetails
type ApplePayDetails struct {
	// The stringified and base64 encoded `paymentData` you retrieved from the Apple framework.
	ApplePayToken string `json:"applePayToken"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**.
	FundingSource *string `json:"fundingSource,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// **applepay**
	Type *string `json:"type,omitempty"`
}

// NewApplePayDetails instantiates a new ApplePayDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplePayDetails(applePayToken string) *ApplePayDetails {
	this := ApplePayDetails{}
	this.ApplePayToken = applePayToken
	var type_ string = "applepay"
	this.Type = &type_
	return &this
}

// NewApplePayDetailsWithDefaults instantiates a new ApplePayDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplePayDetailsWithDefaults() *ApplePayDetails {
	this := ApplePayDetails{}
	var type_ string = "applepay"
	this.Type = &type_
	return &this
}

// GetApplePayToken returns the ApplePayToken field value
func (o *ApplePayDetails) GetApplePayToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplePayToken
}

// GetApplePayTokenOk returns a tuple with the ApplePayToken field value
// and a boolean to check if the value has been set.
func (o *ApplePayDetails) GetApplePayTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplePayToken, true
}

// SetApplePayToken sets field value
func (o *ApplePayDetails) SetApplePayToken(v string) {
	o.ApplePayToken = v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *ApplePayDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplePayDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *ApplePayDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *ApplePayDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *ApplePayDetails) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplePayDetails) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *ApplePayDetails) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *ApplePayDetails) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated
func (o *ApplePayDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplePayDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *ApplePayDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated
func (o *ApplePayDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *ApplePayDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplePayDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *ApplePayDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *ApplePayDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApplePayDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplePayDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplePayDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApplePayDetails) SetType(v string) {
	o.Type = &v
}

func (o ApplePayDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplePayDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applePayToken"] = o.ApplePayToken
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableApplePayDetails struct {
	value *ApplePayDetails
	isSet bool
}

func (v NullableApplePayDetails) Get() *ApplePayDetails {
	return v.value
}

func (v *NullableApplePayDetails) Set(val *ApplePayDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableApplePayDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableApplePayDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplePayDetails(val *ApplePayDetails) *NullableApplePayDetails {
	return &NullableApplePayDetails{value: val, isSet: true}
}

func (v NullableApplePayDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplePayDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ApplePayDetails) isValidFundingSource() bool {
	var allowedEnumValues = []string{"debit"}
	for _, allowed := range allowedEnumValues {
		if o.GetFundingSource() == allowed {
			return true
		}
	}
	return false
}
func (o *ApplePayDetails) isValidType() bool {
	var allowedEnumValues = []string{"applepay"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
