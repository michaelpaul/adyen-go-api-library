/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the GrantLimit type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GrantLimit{}

// GrantLimit struct for GrantLimit
type GrantLimit struct {
	Amount *Amount `json:"amount,omitempty"`
}

// NewGrantLimit instantiates a new GrantLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantLimit() *GrantLimit {
	this := GrantLimit{}
	return &this
}

// NewGrantLimitWithDefaults instantiates a new GrantLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantLimitWithDefaults() *GrantLimit {
	this := GrantLimit{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GrantLimit) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantLimit) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GrantLimit) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *GrantLimit) SetAmount(v Amount) {
	o.Amount = &v
}

func (o GrantLimit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableGrantLimit struct {
	value *GrantLimit
	isSet bool
}

func (v NullableGrantLimit) Get() *GrantLimit {
	return v.value
}

func (v *NullableGrantLimit) Set(val *GrantLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantLimit(val *GrantLimit) *NullableGrantLimit {
	return &NullableGrantLimit{value: val, isSet: true}
}

func (v NullableGrantLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
