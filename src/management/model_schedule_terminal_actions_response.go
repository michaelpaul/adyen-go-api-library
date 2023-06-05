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

// checks if the ScheduleTerminalActionsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ScheduleTerminalActionsResponse{}

// ScheduleTerminalActionsResponse struct for ScheduleTerminalActionsResponse
type ScheduleTerminalActionsResponse struct {
	ActionDetails *ScheduleTerminalActionsRequestActionDetails `json:"actionDetails,omitempty"`
	// A list containing a terminal ID and an action ID for each terminal that the action was scheduled for.
	Items []TerminalActionScheduleDetail `json:"items,omitempty"`
	// The date and time when the action should happen.  Format: [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339), but without the **Z** before the time offset. For example, **2021-11-15T12:16:21+01:00**  The action is sent with the first [maintenance call](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api#when-actions-take-effect) after the specified date and time in the time zone of the terminal.  An empty value causes the action to be sent as soon as possible: at the next maintenance call.
	ScheduledAt *string `json:"scheduledAt,omitempty"`
	// The unique ID of the [store](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/stores). If present, all terminals in the `terminalIds` list must be assigned to this store.
	StoreId *string `json:"storeId,omitempty"`
	// A list of unique IDs of the terminals that the action applies to.
	// Deprecated
	TerminalIds []string `json:"terminalIds,omitempty"`
	// The validation errors that occurred in the list of terminals, and for each error the IDs of the terminals that the error applies to.
	TerminalsWithErrors *map[string][]string `json:"terminalsWithErrors,omitempty"`
	// The number of terminals for which scheduling the action failed.
	TotalErrors *int32 `json:"totalErrors,omitempty"`
	// The number of terminals for which the action was successfully scheduled. This doesn't mean the action has happened yet.
	TotalScheduled *int32 `json:"totalScheduled,omitempty"`
}

// NewScheduleTerminalActionsResponse instantiates a new ScheduleTerminalActionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleTerminalActionsResponse() *ScheduleTerminalActionsResponse {
	this := ScheduleTerminalActionsResponse{}
	return &this
}

// NewScheduleTerminalActionsResponseWithDefaults instantiates a new ScheduleTerminalActionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleTerminalActionsResponseWithDefaults() *ScheduleTerminalActionsResponse {
	this := ScheduleTerminalActionsResponse{}
	return &this
}

// GetActionDetails returns the ActionDetails field value if set, zero value otherwise.
func (o *ScheduleTerminalActionsResponse) GetActionDetails() ScheduleTerminalActionsRequestActionDetails {
	if o == nil || common.IsNil(o.ActionDetails) {
		var ret ScheduleTerminalActionsRequestActionDetails
		return ret
	}
	return *o.ActionDetails
}

// GetActionDetailsOk returns a tuple with the ActionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleTerminalActionsResponse) GetActionDetailsOk() (*ScheduleTerminalActionsRequestActionDetails, bool) {
	if o == nil || common.IsNil(o.ActionDetails) {
		return nil, false
	}
	return o.ActionDetails, true
}

// HasActionDetails returns a boolean if a field has been set.
func (o *ScheduleTerminalActionsResponse) HasActionDetails() bool {
	if o != nil && !common.IsNil(o.ActionDetails) {
		return true
	}

	return false
}

// SetActionDetails gets a reference to the given ScheduleTerminalActionsRequestActionDetails and assigns it to the ActionDetails field.
func (o *ScheduleTerminalActionsResponse) SetActionDetails(v ScheduleTerminalActionsRequestActionDetails) {
	o.ActionDetails = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ScheduleTerminalActionsResponse) GetItems() []TerminalActionScheduleDetail {
	if o == nil || common.IsNil(o.Items) {
		var ret []TerminalActionScheduleDetail
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleTerminalActionsResponse) GetItemsOk() ([]TerminalActionScheduleDetail, bool) {
	if o == nil || common.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ScheduleTerminalActionsResponse) HasItems() bool {
	if o != nil && !common.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TerminalActionScheduleDetail and assigns it to the Items field.
func (o *ScheduleTerminalActionsResponse) SetItems(v []TerminalActionScheduleDetail) {
	o.Items = v
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise.
func (o *ScheduleTerminalActionsResponse) GetScheduledAt() string {
	if o == nil || common.IsNil(o.ScheduledAt) {
		var ret string
		return ret
	}
	return *o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleTerminalActionsResponse) GetScheduledAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ScheduledAt) {
		return nil, false
	}
	return o.ScheduledAt, true
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *ScheduleTerminalActionsResponse) HasScheduledAt() bool {
	if o != nil && !common.IsNil(o.ScheduledAt) {
		return true
	}

	return false
}

// SetScheduledAt gets a reference to the given string and assigns it to the ScheduledAt field.
func (o *ScheduleTerminalActionsResponse) SetScheduledAt(v string) {
	o.ScheduledAt = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *ScheduleTerminalActionsResponse) GetStoreId() string {
	if o == nil || common.IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleTerminalActionsResponse) GetStoreIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *ScheduleTerminalActionsResponse) HasStoreId() bool {
	if o != nil && !common.IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *ScheduleTerminalActionsResponse) SetStoreId(v string) {
	o.StoreId = &v
}

// GetTerminalIds returns the TerminalIds field value if set, zero value otherwise.
// Deprecated
func (o *ScheduleTerminalActionsResponse) GetTerminalIds() []string {
	if o == nil || common.IsNil(o.TerminalIds) {
		var ret []string
		return ret
	}
	return o.TerminalIds
}

// GetTerminalIdsOk returns a tuple with the TerminalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ScheduleTerminalActionsResponse) GetTerminalIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.TerminalIds) {
		return nil, false
	}
	return o.TerminalIds, true
}

// HasTerminalIds returns a boolean if a field has been set.
func (o *ScheduleTerminalActionsResponse) HasTerminalIds() bool {
	if o != nil && !common.IsNil(o.TerminalIds) {
		return true
	}

	return false
}

// SetTerminalIds gets a reference to the given []string and assigns it to the TerminalIds field.
// Deprecated
func (o *ScheduleTerminalActionsResponse) SetTerminalIds(v []string) {
	o.TerminalIds = v
}

// GetTerminalsWithErrors returns the TerminalsWithErrors field value if set, zero value otherwise.
func (o *ScheduleTerminalActionsResponse) GetTerminalsWithErrors() map[string][]string {
	if o == nil || common.IsNil(o.TerminalsWithErrors) {
		var ret map[string][]string
		return ret
	}
	return *o.TerminalsWithErrors
}

// GetTerminalsWithErrorsOk returns a tuple with the TerminalsWithErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleTerminalActionsResponse) GetTerminalsWithErrorsOk() (*map[string][]string, bool) {
	if o == nil || common.IsNil(o.TerminalsWithErrors) {
		return nil, false
	}
	return o.TerminalsWithErrors, true
}

// HasTerminalsWithErrors returns a boolean if a field has been set.
func (o *ScheduleTerminalActionsResponse) HasTerminalsWithErrors() bool {
	if o != nil && !common.IsNil(o.TerminalsWithErrors) {
		return true
	}

	return false
}

// SetTerminalsWithErrors gets a reference to the given map[string][]string and assigns it to the TerminalsWithErrors field.
func (o *ScheduleTerminalActionsResponse) SetTerminalsWithErrors(v map[string][]string) {
	o.TerminalsWithErrors = &v
}

// GetTotalErrors returns the TotalErrors field value if set, zero value otherwise.
func (o *ScheduleTerminalActionsResponse) GetTotalErrors() int32 {
	if o == nil || common.IsNil(o.TotalErrors) {
		var ret int32
		return ret
	}
	return *o.TotalErrors
}

// GetTotalErrorsOk returns a tuple with the TotalErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleTerminalActionsResponse) GetTotalErrorsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.TotalErrors) {
		return nil, false
	}
	return o.TotalErrors, true
}

// HasTotalErrors returns a boolean if a field has been set.
func (o *ScheduleTerminalActionsResponse) HasTotalErrors() bool {
	if o != nil && !common.IsNil(o.TotalErrors) {
		return true
	}

	return false
}

// SetTotalErrors gets a reference to the given int32 and assigns it to the TotalErrors field.
func (o *ScheduleTerminalActionsResponse) SetTotalErrors(v int32) {
	o.TotalErrors = &v
}

// GetTotalScheduled returns the TotalScheduled field value if set, zero value otherwise.
func (o *ScheduleTerminalActionsResponse) GetTotalScheduled() int32 {
	if o == nil || common.IsNil(o.TotalScheduled) {
		var ret int32
		return ret
	}
	return *o.TotalScheduled
}

// GetTotalScheduledOk returns a tuple with the TotalScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleTerminalActionsResponse) GetTotalScheduledOk() (*int32, bool) {
	if o == nil || common.IsNil(o.TotalScheduled) {
		return nil, false
	}
	return o.TotalScheduled, true
}

// HasTotalScheduled returns a boolean if a field has been set.
func (o *ScheduleTerminalActionsResponse) HasTotalScheduled() bool {
	if o != nil && !common.IsNil(o.TotalScheduled) {
		return true
	}

	return false
}

// SetTotalScheduled gets a reference to the given int32 and assigns it to the TotalScheduled field.
func (o *ScheduleTerminalActionsResponse) SetTotalScheduled(v int32) {
	o.TotalScheduled = &v
}

func (o ScheduleTerminalActionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleTerminalActionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ActionDetails) {
		toSerialize["actionDetails"] = o.ActionDetails
	}
	if !common.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !common.IsNil(o.ScheduledAt) {
		toSerialize["scheduledAt"] = o.ScheduledAt
	}
	if !common.IsNil(o.StoreId) {
		toSerialize["storeId"] = o.StoreId
	}
	if !common.IsNil(o.TerminalIds) {
		toSerialize["terminalIds"] = o.TerminalIds
	}
	if !common.IsNil(o.TerminalsWithErrors) {
		toSerialize["terminalsWithErrors"] = o.TerminalsWithErrors
	}
	if !common.IsNil(o.TotalErrors) {
		toSerialize["totalErrors"] = o.TotalErrors
	}
	if !common.IsNil(o.TotalScheduled) {
		toSerialize["totalScheduled"] = o.TotalScheduled
	}
	return toSerialize, nil
}

type NullableScheduleTerminalActionsResponse struct {
	value *ScheduleTerminalActionsResponse
	isSet bool
}

func (v NullableScheduleTerminalActionsResponse) Get() *ScheduleTerminalActionsResponse {
	return v.value
}

func (v *NullableScheduleTerminalActionsResponse) Set(val *ScheduleTerminalActionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleTerminalActionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleTerminalActionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleTerminalActionsResponse(val *ScheduleTerminalActionsResponse) *NullableScheduleTerminalActionsResponse {
	return &NullableScheduleTerminalActionsResponse{value: val, isSet: true}
}

func (v NullableScheduleTerminalActionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleTerminalActionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
