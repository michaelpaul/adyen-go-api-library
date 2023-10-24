/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the CustomNotification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CustomNotification{}

// CustomNotification struct for CustomNotification
type CustomNotification struct {
	Amount *Amount `json:"amount,omitempty"`
	// The event that caused the notification to be sent.Currently supported values: * **AUTHORISATION** * **CANCELLATION** * **REFUND** * **CAPTURE** * **REPORT_AVAILABLE** * **CHARGEBACK** * **REQUEST_FOR_INFORMATION** * **NOTIFICATION_OF_CHARGEBACK** * **NOTIFICATIONTEST** * **ORDER_OPENED** * **ORDER_CLOSED** * **CHARGEBACK_REVERSED** * **REFUNDED_REVERSED** * **REFUND_WITH_DATA**
	EventCode *string `json:"eventCode,omitempty"`
	// The time of the event. Format: [ISO 8601](http://www.w3.org/TR/NOTE-datetime), YYYY-MM-DDThh:mm:ssTZD.
	EventDate *time.Time `json:"eventDate,omitempty"`
	// Your reference for the custom test notification.
	MerchantReference *string `json:"merchantReference,omitempty"`
	// The payment method for the payment that the notification is about. Possible values: * **amex** * **visa** * **mc** * **maestro** * **bcmc** * **paypal**  * **sms**  * **bankTransfer_NL** * **bankTransfer_DE** * **bankTransfer_BE** * **ideal** * **elv** * **sepadirectdebit**
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	// A descripton of what caused the notification.
	Reason *string `json:"reason,omitempty"`
	// The outcome of the event which the notification is about. Set to either **true** or **false**.
	Success *bool `json:"success,omitempty"`
}

// NewCustomNotification instantiates a new CustomNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomNotification() *CustomNotification {
	this := CustomNotification{}
	return &this
}

// NewCustomNotificationWithDefaults instantiates a new CustomNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomNotificationWithDefaults() *CustomNotification {
	this := CustomNotification{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CustomNotification) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNotification) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CustomNotification) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *CustomNotification) SetAmount(v Amount) {
	o.Amount = &v
}

// GetEventCode returns the EventCode field value if set, zero value otherwise.
func (o *CustomNotification) GetEventCode() string {
	if o == nil || common.IsNil(o.EventCode) {
		var ret string
		return ret
	}
	return *o.EventCode
}

// GetEventCodeOk returns a tuple with the EventCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNotification) GetEventCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EventCode) {
		return nil, false
	}
	return o.EventCode, true
}

// HasEventCode returns a boolean if a field has been set.
func (o *CustomNotification) HasEventCode() bool {
	if o != nil && !common.IsNil(o.EventCode) {
		return true
	}

	return false
}

// SetEventCode gets a reference to the given string and assigns it to the EventCode field.
func (o *CustomNotification) SetEventCode(v string) {
	o.EventCode = &v
}

// GetEventDate returns the EventDate field value if set, zero value otherwise.
func (o *CustomNotification) GetEventDate() time.Time {
	if o == nil || common.IsNil(o.EventDate) {
		var ret time.Time
		return ret
	}
	return *o.EventDate
}

// GetEventDateOk returns a tuple with the EventDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNotification) GetEventDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.EventDate) {
		return nil, false
	}
	return o.EventDate, true
}

// HasEventDate returns a boolean if a field has been set.
func (o *CustomNotification) HasEventDate() bool {
	if o != nil && !common.IsNil(o.EventDate) {
		return true
	}

	return false
}

// SetEventDate gets a reference to the given time.Time and assigns it to the EventDate field.
func (o *CustomNotification) SetEventDate(v time.Time) {
	o.EventDate = &v
}

// GetMerchantReference returns the MerchantReference field value if set, zero value otherwise.
func (o *CustomNotification) GetMerchantReference() string {
	if o == nil || common.IsNil(o.MerchantReference) {
		var ret string
		return ret
	}
	return *o.MerchantReference
}

// GetMerchantReferenceOk returns a tuple with the MerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNotification) GetMerchantReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantReference) {
		return nil, false
	}
	return o.MerchantReference, true
}

// HasMerchantReference returns a boolean if a field has been set.
func (o *CustomNotification) HasMerchantReference() bool {
	if o != nil && !common.IsNil(o.MerchantReference) {
		return true
	}

	return false
}

// SetMerchantReference gets a reference to the given string and assigns it to the MerchantReference field.
func (o *CustomNotification) SetMerchantReference(v string) {
	o.MerchantReference = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *CustomNotification) GetPaymentMethod() string {
	if o == nil || common.IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNotification) GetPaymentMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *CustomNotification) HasPaymentMethod() bool {
	if o != nil && !common.IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *CustomNotification) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CustomNotification) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNotification) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CustomNotification) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CustomNotification) SetReason(v string) {
	o.Reason = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CustomNotification) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNotification) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CustomNotification) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CustomNotification) SetSuccess(v bool) {
	o.Success = &v
}

func (o CustomNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.EventCode) {
		toSerialize["eventCode"] = o.EventCode
	}
	if !common.IsNil(o.EventDate) {
		toSerialize["eventDate"] = o.EventDate
	}
	if !common.IsNil(o.MerchantReference) {
		toSerialize["merchantReference"] = o.MerchantReference
	}
	if !common.IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableCustomNotification struct {
	value *CustomNotification
	isSet bool
}

func (v NullableCustomNotification) Get() *CustomNotification {
	return v.value
}

func (v *NullableCustomNotification) Set(val *CustomNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomNotification(val *CustomNotification) *NullableCustomNotification {
	return &NullableCustomNotification{value: val, isSet: true}
}

func (v NullableCustomNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
