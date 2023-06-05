/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the BankAccount type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BankAccount{}

// BankAccount struct for BankAccount
type BankAccount struct {
	// The bank account number (without separators).
	BankAccountNumber *string `json:"bankAccountNumber,omitempty"`
	// The bank city.
	BankCity *string `json:"bankCity,omitempty"`
	// The location id of the bank. The field value is `nil` in most cases.
	BankLocationId *string `json:"bankLocationId,omitempty"`
	// The name of the bank.
	BankName *string `json:"bankName,omitempty"`
	// The [Business Identifier Code](https://en.wikipedia.org/wiki/ISO_9362) (BIC) is the SWIFT address assigned to a bank. The field value is `nil` in most cases.
	Bic *string `json:"bic,omitempty"`
	// Country code where the bank is located.  A valid value is an ISO two-character country code (e.g. 'NL').
	CountryCode *string `json:"countryCode,omitempty"`
	// The [International Bank Account Number](https://en.wikipedia.org/wiki/International_Bank_Account_Number) (IBAN).
	Iban *string `json:"iban,omitempty"`
	// The name of the bank account holder. If you submit a name with non-Latin characters, we automatically replace some of them with corresponding Latin characters to meet the FATF recommendations. For example: * χ12 is converted to ch12. * üA is converted to euA. * Peter Møller is converted to Peter Mller, because banks don't accept 'ø'. After replacement, the ownerName must have at least three alphanumeric characters (A-Z, a-z, 0-9), and at least one of them must be a valid Latin character (A-Z, a-z). For example: * John17 - allowed. * J17 - allowed. * 171 - not allowed. * John-7 - allowed. > If provided details don't match the required format, the response returns the error message: 203 'Invalid bank account holder name'.
	OwnerName *string `json:"ownerName,omitempty"`
	// The bank account holder's tax ID.
	TaxId *string `json:"taxId,omitempty"`
}

// NewBankAccount instantiates a new BankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccount() *BankAccount {
	this := BankAccount{}
	return &this
}

// NewBankAccountWithDefaults instantiates a new BankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountWithDefaults() *BankAccount {
	this := BankAccount{}
	return &this
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *BankAccount) GetBankAccountNumber() string {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *BankAccount) HasBankAccountNumber() bool {
	if o != nil && !common.IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *BankAccount) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetBankCity returns the BankCity field value if set, zero value otherwise.
func (o *BankAccount) GetBankCity() string {
	if o == nil || common.IsNil(o.BankCity) {
		var ret string
		return ret
	}
	return *o.BankCity
}

// GetBankCityOk returns a tuple with the BankCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankCity) {
		return nil, false
	}
	return o.BankCity, true
}

// HasBankCity returns a boolean if a field has been set.
func (o *BankAccount) HasBankCity() bool {
	if o != nil && !common.IsNil(o.BankCity) {
		return true
	}

	return false
}

// SetBankCity gets a reference to the given string and assigns it to the BankCity field.
func (o *BankAccount) SetBankCity(v string) {
	o.BankCity = &v
}

// GetBankLocationId returns the BankLocationId field value if set, zero value otherwise.
func (o *BankAccount) GetBankLocationId() string {
	if o == nil || common.IsNil(o.BankLocationId) {
		var ret string
		return ret
	}
	return *o.BankLocationId
}

// GetBankLocationIdOk returns a tuple with the BankLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankLocationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankLocationId) {
		return nil, false
	}
	return o.BankLocationId, true
}

// HasBankLocationId returns a boolean if a field has been set.
func (o *BankAccount) HasBankLocationId() bool {
	if o != nil && !common.IsNil(o.BankLocationId) {
		return true
	}

	return false
}

// SetBankLocationId gets a reference to the given string and assigns it to the BankLocationId field.
func (o *BankAccount) SetBankLocationId(v string) {
	o.BankLocationId = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *BankAccount) GetBankName() string {
	if o == nil || common.IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *BankAccount) HasBankName() bool {
	if o != nil && !common.IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *BankAccount) SetBankName(v string) {
	o.BankName = &v
}

// GetBic returns the Bic field value if set, zero value otherwise.
func (o *BankAccount) GetBic() string {
	if o == nil || common.IsNil(o.Bic) {
		var ret string
		return ret
	}
	return *o.Bic
}

// GetBicOk returns a tuple with the Bic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBicOk() (*string, bool) {
	if o == nil || common.IsNil(o.Bic) {
		return nil, false
	}
	return o.Bic, true
}

// HasBic returns a boolean if a field has been set.
func (o *BankAccount) HasBic() bool {
	if o != nil && !common.IsNil(o.Bic) {
		return true
	}

	return false
}

// SetBic gets a reference to the given string and assigns it to the Bic field.
func (o *BankAccount) SetBic(v string) {
	o.Bic = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *BankAccount) GetCountryCode() string {
	if o == nil || common.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *BankAccount) HasCountryCode() bool {
	if o != nil && !common.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *BankAccount) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *BankAccount) GetIban() string {
	if o == nil || common.IsNil(o.Iban) {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetIbanOk() (*string, bool) {
	if o == nil || common.IsNil(o.Iban) {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *BankAccount) HasIban() bool {
	if o != nil && !common.IsNil(o.Iban) {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *BankAccount) SetIban(v string) {
	o.Iban = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *BankAccount) GetOwnerName() string {
	if o == nil || common.IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetOwnerNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *BankAccount) HasOwnerName() bool {
	if o != nil && !common.IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *BankAccount) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *BankAccount) GetTaxId() string {
	if o == nil || common.IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetTaxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *BankAccount) HasTaxId() bool {
	if o != nil && !common.IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *BankAccount) SetTaxId(v string) {
	o.TaxId = &v
}

func (o BankAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BankAccountNumber) {
		toSerialize["bankAccountNumber"] = o.BankAccountNumber
	}
	if !common.IsNil(o.BankCity) {
		toSerialize["bankCity"] = o.BankCity
	}
	if !common.IsNil(o.BankLocationId) {
		toSerialize["bankLocationId"] = o.BankLocationId
	}
	if !common.IsNil(o.BankName) {
		toSerialize["bankName"] = o.BankName
	}
	if !common.IsNil(o.Bic) {
		toSerialize["bic"] = o.Bic
	}
	if !common.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !common.IsNil(o.Iban) {
		toSerialize["iban"] = o.Iban
	}
	if !common.IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !common.IsNil(o.TaxId) {
		toSerialize["taxId"] = o.TaxId
	}
	return toSerialize, nil
}

type NullableBankAccount struct {
	value *BankAccount
	isSet bool
}

func (v NullableBankAccount) Get() *BankAccount {
	return v.value
}

func (v *NullableBankAccount) Set(val *BankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccount(val *BankAccount) *NullableBankAccount {
	return &NullableBankAccount{value: val, isSet: true}
}

func (v NullableBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
