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

// checks if the InstallmentsNumber type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &InstallmentsNumber{}

// InstallmentsNumber struct for InstallmentsNumber
type InstallmentsNumber struct {
	// Maximum number of installments
	MaxNumberOfInstallments int32 `json:"maxNumberOfInstallments"`
}

// NewInstallmentsNumber instantiates a new InstallmentsNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallmentsNumber(maxNumberOfInstallments int32) *InstallmentsNumber {
	this := InstallmentsNumber{}
	this.MaxNumberOfInstallments = maxNumberOfInstallments
	return &this
}

// NewInstallmentsNumberWithDefaults instantiates a new InstallmentsNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallmentsNumberWithDefaults() *InstallmentsNumber {
	this := InstallmentsNumber{}
	return &this
}

// GetMaxNumberOfInstallments returns the MaxNumberOfInstallments field value
func (o *InstallmentsNumber) GetMaxNumberOfInstallments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxNumberOfInstallments
}

// GetMaxNumberOfInstallmentsOk returns a tuple with the MaxNumberOfInstallments field value
// and a boolean to check if the value has been set.
func (o *InstallmentsNumber) GetMaxNumberOfInstallmentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxNumberOfInstallments, true
}

// SetMaxNumberOfInstallments sets field value
func (o *InstallmentsNumber) SetMaxNumberOfInstallments(v int32) {
	o.MaxNumberOfInstallments = v
}

func (o InstallmentsNumber) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallmentsNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxNumberOfInstallments"] = o.MaxNumberOfInstallments
	return toSerialize, nil
}

type NullableInstallmentsNumber struct {
	value *InstallmentsNumber
	isSet bool
}

func (v NullableInstallmentsNumber) Get() *InstallmentsNumber {
	return v.value
}

func (v *NullableInstallmentsNumber) Set(val *InstallmentsNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallmentsNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallmentsNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallmentsNumber(val *InstallmentsNumber) *NullableInstallmentsNumber {
	return &NullableInstallmentsNumber{value: val, isSet: true}
}

func (v NullableInstallmentsNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallmentsNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
