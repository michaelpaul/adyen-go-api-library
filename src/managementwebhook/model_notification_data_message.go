/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the NotificationDataMessage type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotificationDataMessage{}

// NotificationDataMessage struct for NotificationDataMessage
type NotificationDataMessage struct {
	// Timestamp for when the webhook was created.
	CreatedAt time.Time                  `json:"createdAt"`
	Data      MidServiceNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of notification.
	Type string `json:"type"`
}

// NewNotificationDataMessage instantiates a new NotificationDataMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationDataMessage(createdAt time.Time, data MidServiceNotificationData, environment string, type_ string) *NotificationDataMessage {
	this := NotificationDataMessage{}
	this.CreatedAt = createdAt
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewNotificationDataMessageWithDefaults instantiates a new NotificationDataMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationDataMessageWithDefaults() *NotificationDataMessage {
	this := NotificationDataMessage{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *NotificationDataMessage) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NotificationDataMessage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NotificationDataMessage) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetData returns the Data field value
func (o *NotificationDataMessage) GetData() MidServiceNotificationData {
	if o == nil {
		var ret MidServiceNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NotificationDataMessage) GetDataOk() (*MidServiceNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NotificationDataMessage) SetData(v MidServiceNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *NotificationDataMessage) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *NotificationDataMessage) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *NotificationDataMessage) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *NotificationDataMessage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationDataMessage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationDataMessage) SetType(v string) {
	o.Type = v
}

func (o NotificationDataMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationDataMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableNotificationDataMessage struct {
	value *NotificationDataMessage
	isSet bool
}

func (v NullableNotificationDataMessage) Get() *NotificationDataMessage {
	return v.value
}

func (v *NullableNotificationDataMessage) Set(val *NotificationDataMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationDataMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationDataMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationDataMessage(val *NotificationDataMessage) *NullableNotificationDataMessage {
	return &NullableNotificationDataMessage{value: val, isSet: true}
}

func (v NullableNotificationDataMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationDataMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
