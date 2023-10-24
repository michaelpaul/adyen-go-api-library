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

// checks if the UpdatePaymentLinkRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdatePaymentLinkRequest{}

// UpdatePaymentLinkRequest struct for UpdatePaymentLinkRequest
type UpdatePaymentLinkRequest struct {
	// Status of the payment link. Possible values: * **expired**
	Status string `json:"status"`
}

// NewUpdatePaymentLinkRequest instantiates a new UpdatePaymentLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePaymentLinkRequest(status string) *UpdatePaymentLinkRequest {
	this := UpdatePaymentLinkRequest{}
	this.Status = status
	return &this
}

// NewUpdatePaymentLinkRequestWithDefaults instantiates a new UpdatePaymentLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePaymentLinkRequestWithDefaults() *UpdatePaymentLinkRequest {
	this := UpdatePaymentLinkRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *UpdatePaymentLinkRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdatePaymentLinkRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdatePaymentLinkRequest) SetStatus(v string) {
	o.Status = v
}

func (o UpdatePaymentLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePaymentLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableUpdatePaymentLinkRequest struct {
	value *UpdatePaymentLinkRequest
	isSet bool
}

func (v NullableUpdatePaymentLinkRequest) Get() *UpdatePaymentLinkRequest {
	return v.value
}

func (v *NullableUpdatePaymentLinkRequest) Set(val *UpdatePaymentLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePaymentLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePaymentLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePaymentLinkRequest(val *UpdatePaymentLinkRequest) *NullableUpdatePaymentLinkRequest {
	return &NullableUpdatePaymentLinkRequest{value: val, isSet: true}
}

func (v NullableUpdatePaymentLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePaymentLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UpdatePaymentLinkRequest) isValidStatus() bool {
	var allowedEnumValues = []string{"expired"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
