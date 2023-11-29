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
)

// AccountHolderStatus struct for AccountHolderStatus
type AccountHolderStatus struct {
	// A list of events scheduled for the account holder.
	Events          *[]AccountEvent         `json:"events,omitempty"`
	PayoutState     *AccountPayoutState     `json:"payoutState,omitempty"`
	ProcessingState *AccountProcessingState `json:"processingState,omitempty"`
	// The status of the account holder. >Permitted values: `Active`, `Inactive`, `Suspended`, `Closed`.
	Status string `json:"status"`
	// The reason why the status was assigned to the account holder.
	StatusReason *string `json:"statusReason,omitempty"`
}

// NewAccountHolderStatus instantiates a new AccountHolderStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderStatus(status string) *AccountHolderStatus {
	this := AccountHolderStatus{}
	this.Status = status
	return &this
}

// NewAccountHolderStatusWithDefaults instantiates a new AccountHolderStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderStatusWithDefaults() *AccountHolderStatus {
	this := AccountHolderStatus{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AccountHolderStatus) GetEvents() []AccountEvent {
	if o == nil || o.Events == nil {
		var ret []AccountEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderStatus) GetEventsOk() (*[]AccountEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AccountHolderStatus) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []AccountEvent and assigns it to the Events field.
func (o *AccountHolderStatus) SetEvents(v []AccountEvent) {
	o.Events = &v
}

// GetPayoutState returns the PayoutState field value if set, zero value otherwise.
func (o *AccountHolderStatus) GetPayoutState() AccountPayoutState {
	if o == nil || o.PayoutState == nil {
		var ret AccountPayoutState
		return ret
	}
	return *o.PayoutState
}

// GetPayoutStateOk returns a tuple with the PayoutState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderStatus) GetPayoutStateOk() (*AccountPayoutState, bool) {
	if o == nil || o.PayoutState == nil {
		return nil, false
	}
	return o.PayoutState, true
}

// HasPayoutState returns a boolean if a field has been set.
func (o *AccountHolderStatus) HasPayoutState() bool {
	if o != nil && o.PayoutState != nil {
		return true
	}

	return false
}

// SetPayoutState gets a reference to the given AccountPayoutState and assigns it to the PayoutState field.
func (o *AccountHolderStatus) SetPayoutState(v AccountPayoutState) {
	o.PayoutState = &v
}

// GetProcessingState returns the ProcessingState field value if set, zero value otherwise.
func (o *AccountHolderStatus) GetProcessingState() AccountProcessingState {
	if o == nil || o.ProcessingState == nil {
		var ret AccountProcessingState
		return ret
	}
	return *o.ProcessingState
}

// GetProcessingStateOk returns a tuple with the ProcessingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderStatus) GetProcessingStateOk() (*AccountProcessingState, bool) {
	if o == nil || o.ProcessingState == nil {
		return nil, false
	}
	return o.ProcessingState, true
}

// HasProcessingState returns a boolean if a field has been set.
func (o *AccountHolderStatus) HasProcessingState() bool {
	if o != nil && o.ProcessingState != nil {
		return true
	}

	return false
}

// SetProcessingState gets a reference to the given AccountProcessingState and assigns it to the ProcessingState field.
func (o *AccountHolderStatus) SetProcessingState(v AccountProcessingState) {
	o.ProcessingState = &v
}

// GetStatus returns the Status field value
func (o *AccountHolderStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountHolderStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccountHolderStatus) SetStatus(v string) {
	o.Status = v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *AccountHolderStatus) GetStatusReason() string {
	if o == nil || o.StatusReason == nil {
		var ret string
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderStatus) GetStatusReasonOk() (*string, bool) {
	if o == nil || o.StatusReason == nil {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *AccountHolderStatus) HasStatusReason() bool {
	if o != nil && o.StatusReason != nil {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given string and assigns it to the StatusReason field.
func (o *AccountHolderStatus) SetStatusReason(v string) {
	o.StatusReason = &v
}

func (o AccountHolderStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.PayoutState != nil {
		toSerialize["payoutState"] = o.PayoutState
	}
	if o.ProcessingState != nil {
		toSerialize["processingState"] = o.ProcessingState
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.StatusReason != nil {
		toSerialize["statusReason"] = o.StatusReason
	}
	return json.Marshal(toSerialize)
}

type NullableAccountHolderStatus struct {
	value *AccountHolderStatus
	isSet bool
}

func (v NullableAccountHolderStatus) Get() *AccountHolderStatus {
	return v.value
}

func (v *NullableAccountHolderStatus) Set(val *AccountHolderStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderStatus(val *AccountHolderStatus) *NullableAccountHolderStatus {
	return &NullableAccountHolderStatus{value: val, isSet: true}
}

func (v NullableAccountHolderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
