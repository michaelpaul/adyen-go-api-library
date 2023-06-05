/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the BcmcInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BcmcInfo{}

// BcmcInfo struct for BcmcInfo
type BcmcInfo struct {
	// Indicates if [Bancontact mobile](https://docs.adyen.com/payment-methods/bancontact/bancontact-mobile) is enabled.
	EnableBcmcMobile *bool `json:"enableBcmcMobile,omitempty"`
}

// NewBcmcInfo instantiates a new BcmcInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBcmcInfo() *BcmcInfo {
	this := BcmcInfo{}
	return &this
}

// NewBcmcInfoWithDefaults instantiates a new BcmcInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBcmcInfoWithDefaults() *BcmcInfo {
	this := BcmcInfo{}
	return &this
}

// GetEnableBcmcMobile returns the EnableBcmcMobile field value if set, zero value otherwise.
func (o *BcmcInfo) GetEnableBcmcMobile() bool {
	if o == nil || common.IsNil(o.EnableBcmcMobile) {
		var ret bool
		return ret
	}
	return *o.EnableBcmcMobile
}

// GetEnableBcmcMobileOk returns a tuple with the EnableBcmcMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BcmcInfo) GetEnableBcmcMobileOk() (*bool, bool) {
	if o == nil || common.IsNil(o.EnableBcmcMobile) {
		return nil, false
	}
	return o.EnableBcmcMobile, true
}

// HasEnableBcmcMobile returns a boolean if a field has been set.
func (o *BcmcInfo) HasEnableBcmcMobile() bool {
	if o != nil && !common.IsNil(o.EnableBcmcMobile) {
		return true
	}

	return false
}

// SetEnableBcmcMobile gets a reference to the given bool and assigns it to the EnableBcmcMobile field.
func (o *BcmcInfo) SetEnableBcmcMobile(v bool) {
	o.EnableBcmcMobile = &v
}

func (o BcmcInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BcmcInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EnableBcmcMobile) {
		toSerialize["enableBcmcMobile"] = o.EnableBcmcMobile
	}
	return toSerialize, nil
}

type NullableBcmcInfo struct {
	value *BcmcInfo
	isSet bool
}

func (v NullableBcmcInfo) Get() *BcmcInfo {
	return v.value
}

func (v *NullableBcmcInfo) Set(val *BcmcInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBcmcInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBcmcInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBcmcInfo(val *BcmcInfo) *NullableBcmcInfo {
	return &NullableBcmcInfo{value: val, isSet: true}
}

func (v NullableBcmcInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBcmcInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
