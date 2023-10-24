/*
Transfer webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the NZLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NZLocalAccountIdentification{}

// NZLocalAccountIdentification struct for NZLocalAccountIdentification
type NZLocalAccountIdentification struct {
	// The 7-digit bank account number, without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// The 2- to 3-digit account suffix, without separators or whitespace.
	AccountSuffix string `json:"accountSuffix"`
	// The 6-digit bank code including the 2-digit bank code and 4-digit branch code, without separators or whitespace.
	BankCode string `json:"bankCode"`
	// **nzLocal**
	Type string `json:"type"`
}

// NewNZLocalAccountIdentification instantiates a new NZLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNZLocalAccountIdentification(accountNumber string, accountSuffix string, bankCode string, type_ string) *NZLocalAccountIdentification {
	this := NZLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.AccountSuffix = accountSuffix
	this.BankCode = bankCode
	this.Type = type_
	return &this
}

// NewNZLocalAccountIdentificationWithDefaults instantiates a new NZLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNZLocalAccountIdentificationWithDefaults() *NZLocalAccountIdentification {
	this := NZLocalAccountIdentification{}
	var type_ string = "nzLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *NZLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *NZLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *NZLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAccountSuffix returns the AccountSuffix field value
func (o *NZLocalAccountIdentification) GetAccountSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountSuffix
}

// GetAccountSuffixOk returns a tuple with the AccountSuffix field value
// and a boolean to check if the value has been set.
func (o *NZLocalAccountIdentification) GetAccountSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountSuffix, true
}

// SetAccountSuffix sets field value
func (o *NZLocalAccountIdentification) SetAccountSuffix(v string) {
	o.AccountSuffix = v
}

// GetBankCode returns the BankCode field value
func (o *NZLocalAccountIdentification) GetBankCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value
// and a boolean to check if the value has been set.
func (o *NZLocalAccountIdentification) GetBankCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCode, true
}

// SetBankCode sets field value
func (o *NZLocalAccountIdentification) SetBankCode(v string) {
	o.BankCode = v
}

// GetType returns the Type field value
func (o *NZLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NZLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NZLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o NZLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NZLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["accountSuffix"] = o.AccountSuffix
	toSerialize["bankCode"] = o.BankCode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableNZLocalAccountIdentification struct {
	value *NZLocalAccountIdentification
	isSet bool
}

func (v NullableNZLocalAccountIdentification) Get() *NZLocalAccountIdentification {
	return v.value
}

func (v *NullableNZLocalAccountIdentification) Set(val *NZLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableNZLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableNZLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNZLocalAccountIdentification(val *NZLocalAccountIdentification) *NullableNZLocalAccountIdentification {
	return &NullableNZLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableNZLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNZLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *NZLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"nzLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
