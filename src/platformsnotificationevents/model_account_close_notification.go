/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
	"time"
)

// AccountCloseNotification struct for AccountCloseNotification
type AccountCloseNotification struct {
	Content CloseAccountResponse        `json:"content"`
	Error   *NotificationErrorContainer `json:"error,omitempty"`
	// The date and time when an event has been completed.
	EventDate time.Time `json:"eventDate"`
	// The event type of the notification.
	EventType string `json:"eventType"`
	// The user or process that has triggered the notification.
	ExecutingUserKey string `json:"executingUserKey"`
	// Indicates whether the notification originated from the live environment or the test environment. If true, the notification originated from the live environment. If false, the notification originated from the test environment.
	Live bool `json:"live"`
	// The PSP reference of the request from which the notification originates.
	PspReference string `json:"pspReference"`
}

// NewAccountCloseNotification instantiates a new AccountCloseNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCloseNotification(content CloseAccountResponse, eventDate time.Time, eventType string, executingUserKey string, live bool, pspReference string) *AccountCloseNotification {
	this := AccountCloseNotification{}
	this.Content = content
	this.EventDate = eventDate
	this.EventType = eventType
	this.ExecutingUserKey = executingUserKey
	this.Live = live
	this.PspReference = pspReference
	return &this
}

// NewAccountCloseNotificationWithDefaults instantiates a new AccountCloseNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCloseNotificationWithDefaults() *AccountCloseNotification {
	this := AccountCloseNotification{}
	return &this
}

// GetContent returns the Content field value
func (o *AccountCloseNotification) GetContent() CloseAccountResponse {
	if o == nil {
		var ret CloseAccountResponse
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AccountCloseNotification) GetContentOk() (*CloseAccountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *AccountCloseNotification) SetContent(v CloseAccountResponse) {
	o.Content = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AccountCloseNotification) GetError() NotificationErrorContainer {
	if o == nil || o.Error == nil {
		var ret NotificationErrorContainer
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCloseNotification) GetErrorOk() (*NotificationErrorContainer, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AccountCloseNotification) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given NotificationErrorContainer and assigns it to the Error field.
func (o *AccountCloseNotification) SetError(v NotificationErrorContainer) {
	o.Error = &v
}

// GetEventDate returns the EventDate field value
func (o *AccountCloseNotification) GetEventDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventDate
}

// GetEventDateOk returns a tuple with the EventDate field value
// and a boolean to check if the value has been set.
func (o *AccountCloseNotification) GetEventDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDate, true
}

// SetEventDate sets field value
func (o *AccountCloseNotification) SetEventDate(v time.Time) {
	o.EventDate = v
}

// GetEventType returns the EventType field value
func (o *AccountCloseNotification) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *AccountCloseNotification) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *AccountCloseNotification) SetEventType(v string) {
	o.EventType = v
}

// GetExecutingUserKey returns the ExecutingUserKey field value
func (o *AccountCloseNotification) GetExecutingUserKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecutingUserKey
}

// GetExecutingUserKeyOk returns a tuple with the ExecutingUserKey field value
// and a boolean to check if the value has been set.
func (o *AccountCloseNotification) GetExecutingUserKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutingUserKey, true
}

// SetExecutingUserKey sets field value
func (o *AccountCloseNotification) SetExecutingUserKey(v string) {
	o.ExecutingUserKey = v
}

// GetLive returns the Live field value
func (o *AccountCloseNotification) GetLive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Live
}

// GetLiveOk returns a tuple with the Live field value
// and a boolean to check if the value has been set.
func (o *AccountCloseNotification) GetLiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Live, true
}

// SetLive sets field value
func (o *AccountCloseNotification) SetLive(v bool) {
	o.Live = v
}

// GetPspReference returns the PspReference field value
func (o *AccountCloseNotification) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *AccountCloseNotification) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *AccountCloseNotification) SetPspReference(v string) {
	o.PspReference = v
}

func (o AccountCloseNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["eventDate"] = o.EventDate
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["executingUserKey"] = o.ExecutingUserKey
	}
	if true {
		toSerialize["live"] = o.Live
	}
	if true {
		toSerialize["pspReference"] = o.PspReference
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCloseNotification struct {
	value *AccountCloseNotification
	isSet bool
}

func (v NullableAccountCloseNotification) Get() *AccountCloseNotification {
	return v.value
}

func (v *NullableAccountCloseNotification) Set(val *AccountCloseNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCloseNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCloseNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCloseNotification(val *AccountCloseNotification) *NullableAccountCloseNotification {
	return &NullableAccountCloseNotification{value: val, isSet: true}
}

func (v NullableAccountCloseNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCloseNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
