/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the AssignTerminalsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssignTerminalsResponse{}

// AssignTerminalsResponse struct for AssignTerminalsResponse
type AssignTerminalsResponse struct {
	// Array that returns a list of the terminals, and for each terminal the result of assigning it to an account or store.  The results can be:    - `Done`: The terminal has been assigned.   - `AssignmentScheduled`: The terminal will be assigned asynschronously.   - `RemoveConfigScheduled`: The terminal was previously assigned and boarded. Wait for the terminal to synchronize with the Adyen platform. For more information, refer to [Reassigning boarded terminals](https://docs.adyen.com/point-of-sale/managing-terminals/assign-terminals#reassign-boarded-terminals).   - `Error`: There was an error when assigning the terminal.
	Results map[string]string `json:"results"`
}

// NewAssignTerminalsResponse instantiates a new AssignTerminalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignTerminalsResponse(results map[string]string) *AssignTerminalsResponse {
	this := AssignTerminalsResponse{}
	this.Results = results
	return &this
}

// NewAssignTerminalsResponseWithDefaults instantiates a new AssignTerminalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignTerminalsResponseWithDefaults() *AssignTerminalsResponse {
	this := AssignTerminalsResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *AssignTerminalsResponse) GetResults() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AssignTerminalsResponse) GetResultsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *AssignTerminalsResponse) SetResults(v map[string]string) {
	o.Results = v
}

func (o AssignTerminalsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignTerminalsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableAssignTerminalsResponse struct {
	value *AssignTerminalsResponse
	isSet bool
}

func (v NullableAssignTerminalsResponse) Get() *AssignTerminalsResponse {
	return v.value
}

func (v *NullableAssignTerminalsResponse) Set(val *AssignTerminalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignTerminalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignTerminalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignTerminalsResponse(val *AssignTerminalsResponse) *NullableAssignTerminalsResponse {
	return &NullableAssignTerminalsResponse{value: val, isSet: true}
}

func (v NullableAssignTerminalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignTerminalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
