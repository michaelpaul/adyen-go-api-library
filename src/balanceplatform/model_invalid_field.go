/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the InvalidField type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &InvalidField{}

// InvalidField struct for InvalidField
type InvalidField struct {
	// Description of the validation error.
	Message string `json:"message"`
	// The field that has an invalid value.
	Name string `json:"name"`
	// The invalid value.
	Value string `json:"value"`
}

// NewInvalidField instantiates a new InvalidField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidField(message string, name string, value string) *InvalidField {
	this := InvalidField{}
	this.Message = message
	this.Name = name
	this.Value = value
	return &this
}

// NewInvalidFieldWithDefaults instantiates a new InvalidField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidFieldWithDefaults() *InvalidField {
	this := InvalidField{}
	return &this
}

// GetMessage returns the Message field value
func (o *InvalidField) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InvalidField) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InvalidField) SetMessage(v string) {
	o.Message = v
}

// GetName returns the Name field value
func (o *InvalidField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvalidField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvalidField) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *InvalidField) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InvalidField) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InvalidField) SetValue(v string) {
	o.Value = v
}

func (o InvalidField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvalidField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableInvalidField struct {
	value *InvalidField
	isSet bool
}

func (v NullableInvalidField) Get() *InvalidField {
	return v.value
}

func (v *NullableInvalidField) Set(val *InvalidField) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidField) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidField(val *InvalidField) *NullableInvalidField {
	return &NullableInvalidField{value: val, isSet: true}
}

func (v NullableInvalidField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
