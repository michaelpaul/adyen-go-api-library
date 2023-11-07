/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the PaymentMethodsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodsResponse{}

// PaymentMethodsResponse struct for PaymentMethodsResponse
type PaymentMethodsResponse struct {
	// Detailed list of payment methods required to generate payment forms.
	PaymentMethods []PaymentMethod `json:"paymentMethods,omitempty"`
	// List of all stored payment methods.
	StoredPaymentMethods []StoredPaymentMethod `json:"storedPaymentMethods,omitempty"`
}

// NewPaymentMethodsResponse instantiates a new PaymentMethodsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodsResponse() *PaymentMethodsResponse {
	this := PaymentMethodsResponse{}
	return &this
}

// NewPaymentMethodsResponseWithDefaults instantiates a new PaymentMethodsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodsResponseWithDefaults() *PaymentMethodsResponse {
	this := PaymentMethodsResponse{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *PaymentMethodsResponse) GetPaymentMethods() []PaymentMethod {
	if o == nil || common.IsNil(o.PaymentMethods) {
		var ret []PaymentMethod
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsResponse) GetPaymentMethodsOk() ([]PaymentMethod, bool) {
	if o == nil || common.IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *PaymentMethodsResponse) HasPaymentMethods() bool {
	if o != nil && !common.IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []PaymentMethod and assigns it to the PaymentMethods field.
func (o *PaymentMethodsResponse) SetPaymentMethods(v []PaymentMethod) {
	o.PaymentMethods = v
}

// GetStoredPaymentMethods returns the StoredPaymentMethods field value if set, zero value otherwise.
func (o *PaymentMethodsResponse) GetStoredPaymentMethods() []StoredPaymentMethod {
	if o == nil || common.IsNil(o.StoredPaymentMethods) {
		var ret []StoredPaymentMethod
		return ret
	}
	return o.StoredPaymentMethods
}

// GetStoredPaymentMethodsOk returns a tuple with the StoredPaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsResponse) GetStoredPaymentMethodsOk() ([]StoredPaymentMethod, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethods) {
		return nil, false
	}
	return o.StoredPaymentMethods, true
}

// HasStoredPaymentMethods returns a boolean if a field has been set.
func (o *PaymentMethodsResponse) HasStoredPaymentMethods() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethods) {
		return true
	}

	return false
}

// SetStoredPaymentMethods gets a reference to the given []StoredPaymentMethod and assigns it to the StoredPaymentMethods field.
func (o *PaymentMethodsResponse) SetStoredPaymentMethods(v []StoredPaymentMethod) {
	o.StoredPaymentMethods = v
}

func (o PaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PaymentMethods) {
		toSerialize["paymentMethods"] = o.PaymentMethods
	}
	if !common.IsNil(o.StoredPaymentMethods) {
		toSerialize["storedPaymentMethods"] = o.StoredPaymentMethods
	}
	return toSerialize, nil
}

type NullablePaymentMethodsResponse struct {
	value *PaymentMethodsResponse
	isSet bool
}

func (v NullablePaymentMethodsResponse) Get() *PaymentMethodsResponse {
	return v.value
}

func (v *NullablePaymentMethodsResponse) Set(val *PaymentMethodsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodsResponse(val *PaymentMethodsResponse) *NullablePaymentMethodsResponse {
	return &NullablePaymentMethodsResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
