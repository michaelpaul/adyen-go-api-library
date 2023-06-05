/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the FraudResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FraudResult{}

// FraudResult struct for FraudResult
type FraudResult struct {
	// The total fraud score generated by the risk checks.
	AccountScore int32 `json:"accountScore"`
	// The result of the individual risk checks.
	Results []FraudCheckResultWrapper `json:"results,omitempty"`
}

// NewFraudResult instantiates a new FraudResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFraudResult(accountScore int32) *FraudResult {
	this := FraudResult{}
	this.AccountScore = accountScore
	return &this
}

// NewFraudResultWithDefaults instantiates a new FraudResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFraudResultWithDefaults() *FraudResult {
	this := FraudResult{}
	return &this
}

// GetAccountScore returns the AccountScore field value
func (o *FraudResult) GetAccountScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountScore
}

// GetAccountScoreOk returns a tuple with the AccountScore field value
// and a boolean to check if the value has been set.
func (o *FraudResult) GetAccountScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountScore, true
}

// SetAccountScore sets field value
func (o *FraudResult) SetAccountScore(v int32) {
	o.AccountScore = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *FraudResult) GetResults() []FraudCheckResultWrapper {
	if o == nil || common.IsNil(o.Results) {
		var ret []FraudCheckResultWrapper
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FraudResult) GetResultsOk() ([]FraudCheckResultWrapper, bool) {
	if o == nil || common.IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *FraudResult) HasResults() bool {
	if o != nil && !common.IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []FraudCheckResultWrapper and assigns it to the Results field.
func (o *FraudResult) SetResults(v []FraudCheckResultWrapper) {
	o.Results = v
}

func (o FraudResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FraudResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountScore"] = o.AccountScore
	if !common.IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableFraudResult struct {
	value *FraudResult
	isSet bool
}

func (v NullableFraudResult) Get() *FraudResult {
	return v.value
}

func (v *NullableFraudResult) Set(val *FraudResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFraudResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFraudResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFraudResult(val *FraudResult) *NullableFraudResult {
	return &NullableFraudResult{value: val, isSet: true}
}

func (v NullableFraudResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFraudResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
