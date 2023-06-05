/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the EntityReference type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EntityReference{}

// EntityReference struct for EntityReference
type EntityReference struct {
	// The unique identifier of the resource.
	Id *string `json:"id,omitempty"`
}

// NewEntityReference instantiates a new EntityReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityReference() *EntityReference {
	this := EntityReference{}
	return &this
}

// NewEntityReferenceWithDefaults instantiates a new EntityReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityReferenceWithDefaults() *EntityReference {
	this := EntityReference{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntityReference) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityReference) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntityReference) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntityReference) SetId(v string) {
	o.Id = &v
}

func (o EntityReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableEntityReference struct {
	value *EntityReference
	isSet bool
}

func (v NullableEntityReference) Get() *EntityReference {
	return v.value
}

func (v *NullableEntityReference) Set(val *EntityReference) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityReference) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityReference(val *EntityReference) *NullableEntityReference {
	return &NullableEntityReference{value: val, isSet: true}
}

func (v NullableEntityReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
