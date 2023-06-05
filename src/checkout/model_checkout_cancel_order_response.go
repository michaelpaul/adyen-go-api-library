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

// checks if the CheckoutCancelOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutCancelOrderResponse{}

// CheckoutCancelOrderResponse struct for CheckoutCancelOrderResponse
type CheckoutCancelOrderResponse struct {
	// A unique reference of the cancellation request.
	PspReference string `json:"pspReference"`
	// The result of the cancellation request.  Possible values:  * **Received** – Indicates the cancellation has successfully been received by Adyen, and will be processed.
	ResultCode string `json:"resultCode"`
}

// NewCheckoutCancelOrderResponse instantiates a new CheckoutCancelOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutCancelOrderResponse(pspReference string, resultCode string) *CheckoutCancelOrderResponse {
	this := CheckoutCancelOrderResponse{}
	this.PspReference = pspReference
	this.ResultCode = resultCode
	return &this
}

// NewCheckoutCancelOrderResponseWithDefaults instantiates a new CheckoutCancelOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutCancelOrderResponseWithDefaults() *CheckoutCancelOrderResponse {
	this := CheckoutCancelOrderResponse{}
	return &this
}

// GetPspReference returns the PspReference field value
func (o *CheckoutCancelOrderResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *CheckoutCancelOrderResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *CheckoutCancelOrderResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetResultCode returns the ResultCode field value
func (o *CheckoutCancelOrderResponse) GetResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value
// and a boolean to check if the value has been set.
func (o *CheckoutCancelOrderResponse) GetResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCode, true
}

// SetResultCode sets field value
func (o *CheckoutCancelOrderResponse) SetResultCode(v string) {
	o.ResultCode = v
}

func (o CheckoutCancelOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutCancelOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pspReference"] = o.PspReference
	toSerialize["resultCode"] = o.ResultCode
	return toSerialize, nil
}

type NullableCheckoutCancelOrderResponse struct {
	value *CheckoutCancelOrderResponse
	isSet bool
}

func (v NullableCheckoutCancelOrderResponse) Get() *CheckoutCancelOrderResponse {
	return v.value
}

func (v *NullableCheckoutCancelOrderResponse) Set(val *CheckoutCancelOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutCancelOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutCancelOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutCancelOrderResponse(val *CheckoutCancelOrderResponse) *NullableCheckoutCancelOrderResponse {
	return &NullableCheckoutCancelOrderResponse{value: val, isSet: true}
}

func (v NullableCheckoutCancelOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutCancelOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutCancelOrderResponse) isValidResultCode() bool {
	var allowedEnumValues = []string{"Received"}
	for _, allowed := range allowedEnumValues {
		if o.GetResultCode() == allowed {
			return true
		}
	}
	return false
}
