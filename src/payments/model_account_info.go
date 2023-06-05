/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the AccountInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountInfo{}

// AccountInfo struct for AccountInfo
type AccountInfo struct {
	// Indicator for the length of time since this shopper account was created in the merchant's environment. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	AccountAgeIndicator *string `json:"accountAgeIndicator,omitempty"`
	// Date when the shopper's account was last changed.
	AccountChangeDate *time.Time `json:"accountChangeDate,omitempty"`
	// Indicator for the length of time since the shopper's account was last updated. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	AccountChangeIndicator *string `json:"accountChangeIndicator,omitempty"`
	// Date when the shopper's account was created.
	AccountCreationDate *time.Time `json:"accountCreationDate,omitempty"`
	// Indicates the type of account. For example, for a multi-account card product. Allowed values: * notApplicable * credit * debit
	AccountType *string `json:"accountType,omitempty"`
	// Number of attempts the shopper tried to add a card to their account in the last day.
	AddCardAttemptsDay *int32 `json:"addCardAttemptsDay,omitempty"`
	// Date the selected delivery address was first used.
	DeliveryAddressUsageDate *time.Time `json:"deliveryAddressUsageDate,omitempty"`
	// Indicator for the length of time since this delivery address was first used. Allowed values: * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	DeliveryAddressUsageIndicator *string `json:"deliveryAddressUsageIndicator,omitempty"`
	// Shopper's home phone number (including the country code).
	// Deprecated
	HomePhone *string `json:"homePhone,omitempty"`
	// Shopper's mobile phone number (including the country code).
	// Deprecated
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// Date when the shopper last changed their password.
	PasswordChangeDate *time.Time `json:"passwordChangeDate,omitempty"`
	// Indicator when the shopper has changed their password. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	PasswordChangeIndicator *string `json:"passwordChangeIndicator,omitempty"`
	// Number of all transactions (successful and abandoned) from this shopper in the past 24 hours.
	PastTransactionsDay *int32 `json:"pastTransactionsDay,omitempty"`
	// Number of all transactions (successful and abandoned) from this shopper in the past year.
	PastTransactionsYear *int32 `json:"pastTransactionsYear,omitempty"`
	// Date this payment method was added to the shopper's account.
	PaymentAccountAge *time.Time `json:"paymentAccountAge,omitempty"`
	// Indicator for the length of time since this payment method was added to this shopper's account. Allowed values: * notApplicable * thisTransaction * lessThan30Days * from30To60Days * moreThan60Days
	PaymentAccountIndicator *string `json:"paymentAccountIndicator,omitempty"`
	// Number of successful purchases in the last six months.
	PurchasesLast6Months *int32 `json:"purchasesLast6Months,omitempty"`
	// Whether suspicious activity was recorded on this account.
	SuspiciousActivity *bool `json:"suspiciousActivity,omitempty"`
	// Shopper's work phone number (including the country code).
	// Deprecated
	WorkPhone *string `json:"workPhone,omitempty"`
}

// NewAccountInfo instantiates a new AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfo() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// NewAccountInfoWithDefaults instantiates a new AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoWithDefaults() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// GetAccountAgeIndicator returns the AccountAgeIndicator field value if set, zero value otherwise.
func (o *AccountInfo) GetAccountAgeIndicator() string {
	if o == nil || common.IsNil(o.AccountAgeIndicator) {
		var ret string
		return ret
	}
	return *o.AccountAgeIndicator
}

// GetAccountAgeIndicatorOk returns a tuple with the AccountAgeIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAccountAgeIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountAgeIndicator) {
		return nil, false
	}
	return o.AccountAgeIndicator, true
}

// HasAccountAgeIndicator returns a boolean if a field has been set.
func (o *AccountInfo) HasAccountAgeIndicator() bool {
	if o != nil && !common.IsNil(o.AccountAgeIndicator) {
		return true
	}

	return false
}

// SetAccountAgeIndicator gets a reference to the given string and assigns it to the AccountAgeIndicator field.
func (o *AccountInfo) SetAccountAgeIndicator(v string) {
	o.AccountAgeIndicator = &v
}

// GetAccountChangeDate returns the AccountChangeDate field value if set, zero value otherwise.
func (o *AccountInfo) GetAccountChangeDate() time.Time {
	if o == nil || common.IsNil(o.AccountChangeDate) {
		var ret time.Time
		return ret
	}
	return *o.AccountChangeDate
}

// GetAccountChangeDateOk returns a tuple with the AccountChangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAccountChangeDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.AccountChangeDate) {
		return nil, false
	}
	return o.AccountChangeDate, true
}

// HasAccountChangeDate returns a boolean if a field has been set.
func (o *AccountInfo) HasAccountChangeDate() bool {
	if o != nil && !common.IsNil(o.AccountChangeDate) {
		return true
	}

	return false
}

// SetAccountChangeDate gets a reference to the given time.Time and assigns it to the AccountChangeDate field.
func (o *AccountInfo) SetAccountChangeDate(v time.Time) {
	o.AccountChangeDate = &v
}

// GetAccountChangeIndicator returns the AccountChangeIndicator field value if set, zero value otherwise.
func (o *AccountInfo) GetAccountChangeIndicator() string {
	if o == nil || common.IsNil(o.AccountChangeIndicator) {
		var ret string
		return ret
	}
	return *o.AccountChangeIndicator
}

// GetAccountChangeIndicatorOk returns a tuple with the AccountChangeIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAccountChangeIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountChangeIndicator) {
		return nil, false
	}
	return o.AccountChangeIndicator, true
}

// HasAccountChangeIndicator returns a boolean if a field has been set.
func (o *AccountInfo) HasAccountChangeIndicator() bool {
	if o != nil && !common.IsNil(o.AccountChangeIndicator) {
		return true
	}

	return false
}

// SetAccountChangeIndicator gets a reference to the given string and assigns it to the AccountChangeIndicator field.
func (o *AccountInfo) SetAccountChangeIndicator(v string) {
	o.AccountChangeIndicator = &v
}

// GetAccountCreationDate returns the AccountCreationDate field value if set, zero value otherwise.
func (o *AccountInfo) GetAccountCreationDate() time.Time {
	if o == nil || common.IsNil(o.AccountCreationDate) {
		var ret time.Time
		return ret
	}
	return *o.AccountCreationDate
}

// GetAccountCreationDateOk returns a tuple with the AccountCreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAccountCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.AccountCreationDate) {
		return nil, false
	}
	return o.AccountCreationDate, true
}

// HasAccountCreationDate returns a boolean if a field has been set.
func (o *AccountInfo) HasAccountCreationDate() bool {
	if o != nil && !common.IsNil(o.AccountCreationDate) {
		return true
	}

	return false
}

// SetAccountCreationDate gets a reference to the given time.Time and assigns it to the AccountCreationDate field.
func (o *AccountInfo) SetAccountCreationDate(v time.Time) {
	o.AccountCreationDate = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *AccountInfo) GetAccountType() string {
	if o == nil || common.IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *AccountInfo) HasAccountType() bool {
	if o != nil && !common.IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *AccountInfo) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAddCardAttemptsDay returns the AddCardAttemptsDay field value if set, zero value otherwise.
func (o *AccountInfo) GetAddCardAttemptsDay() int32 {
	if o == nil || common.IsNil(o.AddCardAttemptsDay) {
		var ret int32
		return ret
	}
	return *o.AddCardAttemptsDay
}

// GetAddCardAttemptsDayOk returns a tuple with the AddCardAttemptsDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAddCardAttemptsDayOk() (*int32, bool) {
	if o == nil || common.IsNil(o.AddCardAttemptsDay) {
		return nil, false
	}
	return o.AddCardAttemptsDay, true
}

// HasAddCardAttemptsDay returns a boolean if a field has been set.
func (o *AccountInfo) HasAddCardAttemptsDay() bool {
	if o != nil && !common.IsNil(o.AddCardAttemptsDay) {
		return true
	}

	return false
}

// SetAddCardAttemptsDay gets a reference to the given int32 and assigns it to the AddCardAttemptsDay field.
func (o *AccountInfo) SetAddCardAttemptsDay(v int32) {
	o.AddCardAttemptsDay = &v
}

// GetDeliveryAddressUsageDate returns the DeliveryAddressUsageDate field value if set, zero value otherwise.
func (o *AccountInfo) GetDeliveryAddressUsageDate() time.Time {
	if o == nil || common.IsNil(o.DeliveryAddressUsageDate) {
		var ret time.Time
		return ret
	}
	return *o.DeliveryAddressUsageDate
}

// GetDeliveryAddressUsageDateOk returns a tuple with the DeliveryAddressUsageDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetDeliveryAddressUsageDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.DeliveryAddressUsageDate) {
		return nil, false
	}
	return o.DeliveryAddressUsageDate, true
}

// HasDeliveryAddressUsageDate returns a boolean if a field has been set.
func (o *AccountInfo) HasDeliveryAddressUsageDate() bool {
	if o != nil && !common.IsNil(o.DeliveryAddressUsageDate) {
		return true
	}

	return false
}

// SetDeliveryAddressUsageDate gets a reference to the given time.Time and assigns it to the DeliveryAddressUsageDate field.
func (o *AccountInfo) SetDeliveryAddressUsageDate(v time.Time) {
	o.DeliveryAddressUsageDate = &v
}

// GetDeliveryAddressUsageIndicator returns the DeliveryAddressUsageIndicator field value if set, zero value otherwise.
func (o *AccountInfo) GetDeliveryAddressUsageIndicator() string {
	if o == nil || common.IsNil(o.DeliveryAddressUsageIndicator) {
		var ret string
		return ret
	}
	return *o.DeliveryAddressUsageIndicator
}

// GetDeliveryAddressUsageIndicatorOk returns a tuple with the DeliveryAddressUsageIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetDeliveryAddressUsageIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliveryAddressUsageIndicator) {
		return nil, false
	}
	return o.DeliveryAddressUsageIndicator, true
}

// HasDeliveryAddressUsageIndicator returns a boolean if a field has been set.
func (o *AccountInfo) HasDeliveryAddressUsageIndicator() bool {
	if o != nil && !common.IsNil(o.DeliveryAddressUsageIndicator) {
		return true
	}

	return false
}

// SetDeliveryAddressUsageIndicator gets a reference to the given string and assigns it to the DeliveryAddressUsageIndicator field.
func (o *AccountInfo) SetDeliveryAddressUsageIndicator(v string) {
	o.DeliveryAddressUsageIndicator = &v
}

// GetHomePhone returns the HomePhone field value if set, zero value otherwise.
// Deprecated
func (o *AccountInfo) GetHomePhone() string {
	if o == nil || common.IsNil(o.HomePhone) {
		var ret string
		return ret
	}
	return *o.HomePhone
}

// GetHomePhoneOk returns a tuple with the HomePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccountInfo) GetHomePhoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.HomePhone) {
		return nil, false
	}
	return o.HomePhone, true
}

// HasHomePhone returns a boolean if a field has been set.
func (o *AccountInfo) HasHomePhone() bool {
	if o != nil && !common.IsNil(o.HomePhone) {
		return true
	}

	return false
}

// SetHomePhone gets a reference to the given string and assigns it to the HomePhone field.
// Deprecated
func (o *AccountInfo) SetHomePhone(v string) {
	o.HomePhone = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
// Deprecated
func (o *AccountInfo) GetMobilePhone() string {
	if o == nil || common.IsNil(o.MobilePhone) {
		var ret string
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccountInfo) GetMobilePhoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.MobilePhone) {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *AccountInfo) HasMobilePhone() bool {
	if o != nil && !common.IsNil(o.MobilePhone) {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.
// Deprecated
func (o *AccountInfo) SetMobilePhone(v string) {
	o.MobilePhone = &v
}

// GetPasswordChangeDate returns the PasswordChangeDate field value if set, zero value otherwise.
func (o *AccountInfo) GetPasswordChangeDate() time.Time {
	if o == nil || common.IsNil(o.PasswordChangeDate) {
		var ret time.Time
		return ret
	}
	return *o.PasswordChangeDate
}

// GetPasswordChangeDateOk returns a tuple with the PasswordChangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetPasswordChangeDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.PasswordChangeDate) {
		return nil, false
	}
	return o.PasswordChangeDate, true
}

// HasPasswordChangeDate returns a boolean if a field has been set.
func (o *AccountInfo) HasPasswordChangeDate() bool {
	if o != nil && !common.IsNil(o.PasswordChangeDate) {
		return true
	}

	return false
}

// SetPasswordChangeDate gets a reference to the given time.Time and assigns it to the PasswordChangeDate field.
func (o *AccountInfo) SetPasswordChangeDate(v time.Time) {
	o.PasswordChangeDate = &v
}

// GetPasswordChangeIndicator returns the PasswordChangeIndicator field value if set, zero value otherwise.
func (o *AccountInfo) GetPasswordChangeIndicator() string {
	if o == nil || common.IsNil(o.PasswordChangeIndicator) {
		var ret string
		return ret
	}
	return *o.PasswordChangeIndicator
}

// GetPasswordChangeIndicatorOk returns a tuple with the PasswordChangeIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetPasswordChangeIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.PasswordChangeIndicator) {
		return nil, false
	}
	return o.PasswordChangeIndicator, true
}

// HasPasswordChangeIndicator returns a boolean if a field has been set.
func (o *AccountInfo) HasPasswordChangeIndicator() bool {
	if o != nil && !common.IsNil(o.PasswordChangeIndicator) {
		return true
	}

	return false
}

// SetPasswordChangeIndicator gets a reference to the given string and assigns it to the PasswordChangeIndicator field.
func (o *AccountInfo) SetPasswordChangeIndicator(v string) {
	o.PasswordChangeIndicator = &v
}

// GetPastTransactionsDay returns the PastTransactionsDay field value if set, zero value otherwise.
func (o *AccountInfo) GetPastTransactionsDay() int32 {
	if o == nil || common.IsNil(o.PastTransactionsDay) {
		var ret int32
		return ret
	}
	return *o.PastTransactionsDay
}

// GetPastTransactionsDayOk returns a tuple with the PastTransactionsDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetPastTransactionsDayOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PastTransactionsDay) {
		return nil, false
	}
	return o.PastTransactionsDay, true
}

// HasPastTransactionsDay returns a boolean if a field has been set.
func (o *AccountInfo) HasPastTransactionsDay() bool {
	if o != nil && !common.IsNil(o.PastTransactionsDay) {
		return true
	}

	return false
}

// SetPastTransactionsDay gets a reference to the given int32 and assigns it to the PastTransactionsDay field.
func (o *AccountInfo) SetPastTransactionsDay(v int32) {
	o.PastTransactionsDay = &v
}

// GetPastTransactionsYear returns the PastTransactionsYear field value if set, zero value otherwise.
func (o *AccountInfo) GetPastTransactionsYear() int32 {
	if o == nil || common.IsNil(o.PastTransactionsYear) {
		var ret int32
		return ret
	}
	return *o.PastTransactionsYear
}

// GetPastTransactionsYearOk returns a tuple with the PastTransactionsYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetPastTransactionsYearOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PastTransactionsYear) {
		return nil, false
	}
	return o.PastTransactionsYear, true
}

// HasPastTransactionsYear returns a boolean if a field has been set.
func (o *AccountInfo) HasPastTransactionsYear() bool {
	if o != nil && !common.IsNil(o.PastTransactionsYear) {
		return true
	}

	return false
}

// SetPastTransactionsYear gets a reference to the given int32 and assigns it to the PastTransactionsYear field.
func (o *AccountInfo) SetPastTransactionsYear(v int32) {
	o.PastTransactionsYear = &v
}

// GetPaymentAccountAge returns the PaymentAccountAge field value if set, zero value otherwise.
func (o *AccountInfo) GetPaymentAccountAge() time.Time {
	if o == nil || common.IsNil(o.PaymentAccountAge) {
		var ret time.Time
		return ret
	}
	return *o.PaymentAccountAge
}

// GetPaymentAccountAgeOk returns a tuple with the PaymentAccountAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetPaymentAccountAgeOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.PaymentAccountAge) {
		return nil, false
	}
	return o.PaymentAccountAge, true
}

// HasPaymentAccountAge returns a boolean if a field has been set.
func (o *AccountInfo) HasPaymentAccountAge() bool {
	if o != nil && !common.IsNil(o.PaymentAccountAge) {
		return true
	}

	return false
}

// SetPaymentAccountAge gets a reference to the given time.Time and assigns it to the PaymentAccountAge field.
func (o *AccountInfo) SetPaymentAccountAge(v time.Time) {
	o.PaymentAccountAge = &v
}

// GetPaymentAccountIndicator returns the PaymentAccountIndicator field value if set, zero value otherwise.
func (o *AccountInfo) GetPaymentAccountIndicator() string {
	if o == nil || common.IsNil(o.PaymentAccountIndicator) {
		var ret string
		return ret
	}
	return *o.PaymentAccountIndicator
}

// GetPaymentAccountIndicatorOk returns a tuple with the PaymentAccountIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetPaymentAccountIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentAccountIndicator) {
		return nil, false
	}
	return o.PaymentAccountIndicator, true
}

// HasPaymentAccountIndicator returns a boolean if a field has been set.
func (o *AccountInfo) HasPaymentAccountIndicator() bool {
	if o != nil && !common.IsNil(o.PaymentAccountIndicator) {
		return true
	}

	return false
}

// SetPaymentAccountIndicator gets a reference to the given string and assigns it to the PaymentAccountIndicator field.
func (o *AccountInfo) SetPaymentAccountIndicator(v string) {
	o.PaymentAccountIndicator = &v
}

// GetPurchasesLast6Months returns the PurchasesLast6Months field value if set, zero value otherwise.
func (o *AccountInfo) GetPurchasesLast6Months() int32 {
	if o == nil || common.IsNil(o.PurchasesLast6Months) {
		var ret int32
		return ret
	}
	return *o.PurchasesLast6Months
}

// GetPurchasesLast6MonthsOk returns a tuple with the PurchasesLast6Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetPurchasesLast6MonthsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.PurchasesLast6Months) {
		return nil, false
	}
	return o.PurchasesLast6Months, true
}

// HasPurchasesLast6Months returns a boolean if a field has been set.
func (o *AccountInfo) HasPurchasesLast6Months() bool {
	if o != nil && !common.IsNil(o.PurchasesLast6Months) {
		return true
	}

	return false
}

// SetPurchasesLast6Months gets a reference to the given int32 and assigns it to the PurchasesLast6Months field.
func (o *AccountInfo) SetPurchasesLast6Months(v int32) {
	o.PurchasesLast6Months = &v
}

// GetSuspiciousActivity returns the SuspiciousActivity field value if set, zero value otherwise.
func (o *AccountInfo) GetSuspiciousActivity() bool {
	if o == nil || common.IsNil(o.SuspiciousActivity) {
		var ret bool
		return ret
	}
	return *o.SuspiciousActivity
}

// GetSuspiciousActivityOk returns a tuple with the SuspiciousActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetSuspiciousActivityOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SuspiciousActivity) {
		return nil, false
	}
	return o.SuspiciousActivity, true
}

// HasSuspiciousActivity returns a boolean if a field has been set.
func (o *AccountInfo) HasSuspiciousActivity() bool {
	if o != nil && !common.IsNil(o.SuspiciousActivity) {
		return true
	}

	return false
}

// SetSuspiciousActivity gets a reference to the given bool and assigns it to the SuspiciousActivity field.
func (o *AccountInfo) SetSuspiciousActivity(v bool) {
	o.SuspiciousActivity = &v
}

// GetWorkPhone returns the WorkPhone field value if set, zero value otherwise.
// Deprecated
func (o *AccountInfo) GetWorkPhone() string {
	if o == nil || common.IsNil(o.WorkPhone) {
		var ret string
		return ret
	}
	return *o.WorkPhone
}

// GetWorkPhoneOk returns a tuple with the WorkPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccountInfo) GetWorkPhoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkPhone) {
		return nil, false
	}
	return o.WorkPhone, true
}

// HasWorkPhone returns a boolean if a field has been set.
func (o *AccountInfo) HasWorkPhone() bool {
	if o != nil && !common.IsNil(o.WorkPhone) {
		return true
	}

	return false
}

// SetWorkPhone gets a reference to the given string and assigns it to the WorkPhone field.
// Deprecated
func (o *AccountInfo) SetWorkPhone(v string) {
	o.WorkPhone = &v
}

func (o AccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountAgeIndicator) {
		toSerialize["accountAgeIndicator"] = o.AccountAgeIndicator
	}
	if !common.IsNil(o.AccountChangeDate) {
		toSerialize["accountChangeDate"] = o.AccountChangeDate
	}
	if !common.IsNil(o.AccountChangeIndicator) {
		toSerialize["accountChangeIndicator"] = o.AccountChangeIndicator
	}
	if !common.IsNil(o.AccountCreationDate) {
		toSerialize["accountCreationDate"] = o.AccountCreationDate
	}
	if !common.IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !common.IsNil(o.AddCardAttemptsDay) {
		toSerialize["addCardAttemptsDay"] = o.AddCardAttemptsDay
	}
	if !common.IsNil(o.DeliveryAddressUsageDate) {
		toSerialize["deliveryAddressUsageDate"] = o.DeliveryAddressUsageDate
	}
	if !common.IsNil(o.DeliveryAddressUsageIndicator) {
		toSerialize["deliveryAddressUsageIndicator"] = o.DeliveryAddressUsageIndicator
	}
	if !common.IsNil(o.HomePhone) {
		toSerialize["homePhone"] = o.HomePhone
	}
	if !common.IsNil(o.MobilePhone) {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if !common.IsNil(o.PasswordChangeDate) {
		toSerialize["passwordChangeDate"] = o.PasswordChangeDate
	}
	if !common.IsNil(o.PasswordChangeIndicator) {
		toSerialize["passwordChangeIndicator"] = o.PasswordChangeIndicator
	}
	if !common.IsNil(o.PastTransactionsDay) {
		toSerialize["pastTransactionsDay"] = o.PastTransactionsDay
	}
	if !common.IsNil(o.PastTransactionsYear) {
		toSerialize["pastTransactionsYear"] = o.PastTransactionsYear
	}
	if !common.IsNil(o.PaymentAccountAge) {
		toSerialize["paymentAccountAge"] = o.PaymentAccountAge
	}
	if !common.IsNil(o.PaymentAccountIndicator) {
		toSerialize["paymentAccountIndicator"] = o.PaymentAccountIndicator
	}
	if !common.IsNil(o.PurchasesLast6Months) {
		toSerialize["purchasesLast6Months"] = o.PurchasesLast6Months
	}
	if !common.IsNil(o.SuspiciousActivity) {
		toSerialize["suspiciousActivity"] = o.SuspiciousActivity
	}
	if !common.IsNil(o.WorkPhone) {
		toSerialize["workPhone"] = o.WorkPhone
	}
	return toSerialize, nil
}

type NullableAccountInfo struct {
	value *AccountInfo
	isSet bool
}

func (v NullableAccountInfo) Get() *AccountInfo {
	return v.value
}

func (v *NullableAccountInfo) Set(val *AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfo(val *AccountInfo) *NullableAccountInfo {
	return &NullableAccountInfo{value: val, isSet: true}
}

func (v NullableAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AccountInfo) isValidAccountAgeIndicator() bool {
	var allowedEnumValues = []string{"notApplicable", "thisTransaction", "lessThan30Days", "from30To60Days", "moreThan60Days"}
	for _, allowed := range allowedEnumValues {
		if o.GetAccountAgeIndicator() == allowed {
			return true
		}
	}
	return false
}
func (o *AccountInfo) isValidAccountChangeIndicator() bool {
	var allowedEnumValues = []string{"thisTransaction", "lessThan30Days", "from30To60Days", "moreThan60Days"}
	for _, allowed := range allowedEnumValues {
		if o.GetAccountChangeIndicator() == allowed {
			return true
		}
	}
	return false
}
func (o *AccountInfo) isValidAccountType() bool {
	var allowedEnumValues = []string{"notApplicable", "credit", "debit"}
	for _, allowed := range allowedEnumValues {
		if o.GetAccountType() == allowed {
			return true
		}
	}
	return false
}
func (o *AccountInfo) isValidDeliveryAddressUsageIndicator() bool {
	var allowedEnumValues = []string{"thisTransaction", "lessThan30Days", "from30To60Days", "moreThan60Days"}
	for _, allowed := range allowedEnumValues {
		if o.GetDeliveryAddressUsageIndicator() == allowed {
			return true
		}
	}
	return false
}
func (o *AccountInfo) isValidPasswordChangeIndicator() bool {
	var allowedEnumValues = []string{"notApplicable", "thisTransaction", "lessThan30Days", "from30To60Days", "moreThan60Days"}
	for _, allowed := range allowedEnumValues {
		if o.GetPasswordChangeIndicator() == allowed {
			return true
		}
	}
	return false
}
func (o *AccountInfo) isValidPaymentAccountIndicator() bool {
	var allowedEnumValues = []string{"notApplicable", "thisTransaction", "lessThan30Days", "from30To60Days", "moreThan60Days"}
	for _, allowed := range allowedEnumValues {
		if o.GetPaymentAccountIndicator() == allowed {
			return true
		}
	}
	return false
}
