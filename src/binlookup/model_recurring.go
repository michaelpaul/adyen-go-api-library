/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the Recurring type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Recurring{}

// Recurring struct for Recurring
type Recurring struct {
	// The type of recurring contract to be used. Possible values: * `ONECLICK` – Payment details can be used to initiate a one-click payment, where the shopper enters the [card security code (CVC/CVV)](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-security-code-cvc-cvv-cid). * `RECURRING` – Payment details can be used without the card security code to initiate [card-not-present transactions](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-not-present-cnp). * `ONECLICK,RECURRING` – Payment details can be used regardless of whether the shopper is on your site or not. * `PAYOUT` – Payment details can be used to [make a payout](https://docs.adyen.com/online-payments/online-payouts).
	Contract *string `json:"contract,omitempty"`
	// A descriptive name for this detail.
	RecurringDetailName *string `json:"recurringDetailName,omitempty"`
	// Date after which no further authorisations shall be performed. Only for 3D Secure 2.
	RecurringExpiry *time.Time `json:"recurringExpiry,omitempty"`
	// Minimum number of days between authorisations. Only for 3D Secure 2.
	RecurringFrequency *string `json:"recurringFrequency,omitempty"`
	// The name of the token service.
	TokenService *string `json:"tokenService,omitempty"`
}

// NewRecurring instantiates a new Recurring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurring() *Recurring {
	this := Recurring{}
	return &this
}

// NewRecurringWithDefaults instantiates a new Recurring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringWithDefaults() *Recurring {
	this := Recurring{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *Recurring) GetContract() string {
	if o == nil || common.IsNil(o.Contract) {
		var ret string
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recurring) GetContractOk() (*string, bool) {
	if o == nil || common.IsNil(o.Contract) {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *Recurring) HasContract() bool {
	if o != nil && !common.IsNil(o.Contract) {
		return true
	}

	return false
}

// SetContract gets a reference to the given string and assigns it to the Contract field.
func (o *Recurring) SetContract(v string) {
	o.Contract = &v
}

// GetRecurringDetailName returns the RecurringDetailName field value if set, zero value otherwise.
func (o *Recurring) GetRecurringDetailName() string {
	if o == nil || common.IsNil(o.RecurringDetailName) {
		var ret string
		return ret
	}
	return *o.RecurringDetailName
}

// GetRecurringDetailNameOk returns a tuple with the RecurringDetailName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recurring) GetRecurringDetailNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailName) {
		return nil, false
	}
	return o.RecurringDetailName, true
}

// HasRecurringDetailName returns a boolean if a field has been set.
func (o *Recurring) HasRecurringDetailName() bool {
	if o != nil && !common.IsNil(o.RecurringDetailName) {
		return true
	}

	return false
}

// SetRecurringDetailName gets a reference to the given string and assigns it to the RecurringDetailName field.
func (o *Recurring) SetRecurringDetailName(v string) {
	o.RecurringDetailName = &v
}

// GetRecurringExpiry returns the RecurringExpiry field value if set, zero value otherwise.
func (o *Recurring) GetRecurringExpiry() time.Time {
	if o == nil || common.IsNil(o.RecurringExpiry) {
		var ret time.Time
		return ret
	}
	return *o.RecurringExpiry
}

// GetRecurringExpiryOk returns a tuple with the RecurringExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recurring) GetRecurringExpiryOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.RecurringExpiry) {
		return nil, false
	}
	return o.RecurringExpiry, true
}

// HasRecurringExpiry returns a boolean if a field has been set.
func (o *Recurring) HasRecurringExpiry() bool {
	if o != nil && !common.IsNil(o.RecurringExpiry) {
		return true
	}

	return false
}

// SetRecurringExpiry gets a reference to the given time.Time and assigns it to the RecurringExpiry field.
func (o *Recurring) SetRecurringExpiry(v time.Time) {
	o.RecurringExpiry = &v
}

// GetRecurringFrequency returns the RecurringFrequency field value if set, zero value otherwise.
func (o *Recurring) GetRecurringFrequency() string {
	if o == nil || common.IsNil(o.RecurringFrequency) {
		var ret string
		return ret
	}
	return *o.RecurringFrequency
}

// GetRecurringFrequencyOk returns a tuple with the RecurringFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recurring) GetRecurringFrequencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringFrequency) {
		return nil, false
	}
	return o.RecurringFrequency, true
}

// HasRecurringFrequency returns a boolean if a field has been set.
func (o *Recurring) HasRecurringFrequency() bool {
	if o != nil && !common.IsNil(o.RecurringFrequency) {
		return true
	}

	return false
}

// SetRecurringFrequency gets a reference to the given string and assigns it to the RecurringFrequency field.
func (o *Recurring) SetRecurringFrequency(v string) {
	o.RecurringFrequency = &v
}

// GetTokenService returns the TokenService field value if set, zero value otherwise.
func (o *Recurring) GetTokenService() string {
	if o == nil || common.IsNil(o.TokenService) {
		var ret string
		return ret
	}
	return *o.TokenService
}

// GetTokenServiceOk returns a tuple with the TokenService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recurring) GetTokenServiceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenService) {
		return nil, false
	}
	return o.TokenService, true
}

// HasTokenService returns a boolean if a field has been set.
func (o *Recurring) HasTokenService() bool {
	if o != nil && !common.IsNil(o.TokenService) {
		return true
	}

	return false
}

// SetTokenService gets a reference to the given string and assigns it to the TokenService field.
func (o *Recurring) SetTokenService(v string) {
	o.TokenService = &v
}

func (o Recurring) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Recurring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Contract) {
		toSerialize["contract"] = o.Contract
	}
	if !common.IsNil(o.RecurringDetailName) {
		toSerialize["recurringDetailName"] = o.RecurringDetailName
	}
	if !common.IsNil(o.RecurringExpiry) {
		toSerialize["recurringExpiry"] = o.RecurringExpiry
	}
	if !common.IsNil(o.RecurringFrequency) {
		toSerialize["recurringFrequency"] = o.RecurringFrequency
	}
	if !common.IsNil(o.TokenService) {
		toSerialize["tokenService"] = o.TokenService
	}
	return toSerialize, nil
}

type NullableRecurring struct {
	value *Recurring
	isSet bool
}

func (v NullableRecurring) Get() *Recurring {
	return v.value
}

func (v *NullableRecurring) Set(val *Recurring) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurring) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurring(val *Recurring) *NullableRecurring {
	return &NullableRecurring{value: val, isSet: true}
}

func (v NullableRecurring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Recurring) isValidContract() bool {
	var allowedEnumValues = []string{"ONECLICK", "RECURRING", "PAYOUT"}
	for _, allowed := range allowedEnumValues {
		if o.GetContract() == allowed {
			return true
		}
	}
	return false
}
func (o *Recurring) isValidTokenService() bool {
	var allowedEnumValues = []string{"VISATOKENSERVICE", "MCTOKENSERVICE", "AMEXTOKENSERVICE", "TOKEN_SHARING"}
	for _, allowed := range allowedEnumValues {
		if o.GetTokenService() == allowed {
			return true
		}
	}
	return false
}
