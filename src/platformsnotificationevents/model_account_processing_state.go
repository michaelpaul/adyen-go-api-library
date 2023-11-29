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

// AccountProcessingState struct for AccountProcessingState
type AccountProcessingState struct {
	// The reason why processing has been disabled.
	DisableReason *string `json:"disableReason,omitempty"`
	// Indicates whether the processing of payments is allowed.
	Disabled      *bool   `json:"disabled,omitempty"`
	ProcessedFrom *Amount `json:"processedFrom,omitempty"`
	ProcessedTo   *Amount `json:"processedTo,omitempty"`
	// The processing tier that the account holder occupies.
	TierNumber *int32 `json:"tierNumber,omitempty"`
}

// NewAccountProcessingState instantiates a new AccountProcessingState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountProcessingState() *AccountProcessingState {
	this := AccountProcessingState{}
	return &this
}

// NewAccountProcessingStateWithDefaults instantiates a new AccountProcessingState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountProcessingStateWithDefaults() *AccountProcessingState {
	this := AccountProcessingState{}
	return &this
}

// GetDisableReason returns the DisableReason field value if set, zero value otherwise.
func (o *AccountProcessingState) GetDisableReason() string {
	if o == nil || o.DisableReason == nil {
		var ret string
		return ret
	}
	return *o.DisableReason
}

// GetDisableReasonOk returns a tuple with the DisableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountProcessingState) GetDisableReasonOk() (*string, bool) {
	if o == nil || o.DisableReason == nil {
		return nil, false
	}
	return o.DisableReason, true
}

// HasDisableReason returns a boolean if a field has been set.
func (o *AccountProcessingState) HasDisableReason() bool {
	if o != nil && o.DisableReason != nil {
		return true
	}

	return false
}

// SetDisableReason gets a reference to the given string and assigns it to the DisableReason field.
func (o *AccountProcessingState) SetDisableReason(v string) {
	o.DisableReason = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AccountProcessingState) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountProcessingState) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AccountProcessingState) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AccountProcessingState) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetProcessedFrom returns the ProcessedFrom field value if set, zero value otherwise.
func (o *AccountProcessingState) GetProcessedFrom() Amount {
	if o == nil || o.ProcessedFrom == nil {
		var ret Amount
		return ret
	}
	return *o.ProcessedFrom
}

// GetProcessedFromOk returns a tuple with the ProcessedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountProcessingState) GetProcessedFromOk() (*Amount, bool) {
	if o == nil || o.ProcessedFrom == nil {
		return nil, false
	}
	return o.ProcessedFrom, true
}

// HasProcessedFrom returns a boolean if a field has been set.
func (o *AccountProcessingState) HasProcessedFrom() bool {
	if o != nil && o.ProcessedFrom != nil {
		return true
	}

	return false
}

// SetProcessedFrom gets a reference to the given Amount and assigns it to the ProcessedFrom field.
func (o *AccountProcessingState) SetProcessedFrom(v Amount) {
	o.ProcessedFrom = &v
}

// GetProcessedTo returns the ProcessedTo field value if set, zero value otherwise.
func (o *AccountProcessingState) GetProcessedTo() Amount {
	if o == nil || o.ProcessedTo == nil {
		var ret Amount
		return ret
	}
	return *o.ProcessedTo
}

// GetProcessedToOk returns a tuple with the ProcessedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountProcessingState) GetProcessedToOk() (*Amount, bool) {
	if o == nil || o.ProcessedTo == nil {
		return nil, false
	}
	return o.ProcessedTo, true
}

// HasProcessedTo returns a boolean if a field has been set.
func (o *AccountProcessingState) HasProcessedTo() bool {
	if o != nil && o.ProcessedTo != nil {
		return true
	}

	return false
}

// SetProcessedTo gets a reference to the given Amount and assigns it to the ProcessedTo field.
func (o *AccountProcessingState) SetProcessedTo(v Amount) {
	o.ProcessedTo = &v
}

// GetTierNumber returns the TierNumber field value if set, zero value otherwise.
func (o *AccountProcessingState) GetTierNumber() int32 {
	if o == nil || o.TierNumber == nil {
		var ret int32
		return ret
	}
	return *o.TierNumber
}

// GetTierNumberOk returns a tuple with the TierNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountProcessingState) GetTierNumberOk() (*int32, bool) {
	if o == nil || o.TierNumber == nil {
		return nil, false
	}
	return o.TierNumber, true
}

// HasTierNumber returns a boolean if a field has been set.
func (o *AccountProcessingState) HasTierNumber() bool {
	if o != nil && o.TierNumber != nil {
		return true
	}

	return false
}

// SetTierNumber gets a reference to the given int32 and assigns it to the TierNumber field.
func (o *AccountProcessingState) SetTierNumber(v int32) {
	o.TierNumber = &v
}

func (o AccountProcessingState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisableReason != nil {
		toSerialize["disableReason"] = o.DisableReason
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.ProcessedFrom != nil {
		toSerialize["processedFrom"] = o.ProcessedFrom
	}
	if o.ProcessedTo != nil {
		toSerialize["processedTo"] = o.ProcessedTo
	}
	if o.TierNumber != nil {
		toSerialize["tierNumber"] = o.TierNumber
	}
	return json.Marshal(toSerialize)
}

type NullableAccountProcessingState struct {
	value *AccountProcessingState
	isSet bool
}

func (v NullableAccountProcessingState) Get() *AccountProcessingState {
	return v.value
}

func (v *NullableAccountProcessingState) Set(val *AccountProcessingState) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountProcessingState) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountProcessingState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountProcessingState(val *AccountProcessingState) *NullableAccountProcessingState {
	return &NullableAccountProcessingState{value: val, isSet: true}
}

func (v NullableAccountProcessingState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountProcessingState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
