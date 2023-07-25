/*
Transfers API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the PartyIdentification2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PartyIdentification2{}

// PartyIdentification2 struct for PartyIdentification2
type PartyIdentification2 struct {
	Address *Address2 `json:"address,omitempty"`
	// The date of birth of the individual in [ISO-8601](https://www.w3.org/TR/NOTE-datetime) format. For example, **YYYY-MM-DD**. Should not be before January 1, 1900.  Allowed only when `type` is **individual**.
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	// First name of the individual.  Allowed only when `type` is **individual**.
	FirstName *string `json:"firstName,omitempty"`
	// The name of the entity.
	FullName string `json:"fullName"`
	// Last name of the individual.  Allowed only when `type` is **individual**.
	LastName *string `json:"lastName,omitempty"`
	// A unique reference to identify the party or counterparty involved in transfers. This identifier ensures consistency and uniqueness throughout all transactions initiated to and from the same party. For example, your client's unique wallet or payee ID.
	Reference *string `json:"reference,omitempty"`
	// The type of entity that owns the bank account.   Possible values: **individual**, **organization**, or **unknown**.
	Type *string `json:"type,omitempty"`
}

// NewPartyIdentification2 instantiates a new PartyIdentification2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartyIdentification2(fullName string) *PartyIdentification2 {
	this := PartyIdentification2{}
	this.FullName = fullName
	var type_ string = "unknown"
	this.Type = &type_
	return &this
}

// NewPartyIdentification2WithDefaults instantiates a new PartyIdentification2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyIdentification2WithDefaults() *PartyIdentification2 {
	this := PartyIdentification2{}
	var type_ string = "unknown"
	this.Type = &type_
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PartyIdentification2) GetAddress() Address2 {
	if o == nil || common.IsNil(o.Address) {
		var ret Address2
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification2) GetAddressOk() (*Address2, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PartyIdentification2) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address2 and assigns it to the Address field.
func (o *PartyIdentification2) SetAddress(v Address2) {
	o.Address = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *PartyIdentification2) GetDateOfBirth() string {
	if o == nil || common.IsNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification2) GetDateOfBirthOk() (*string, bool) {
	if o == nil || common.IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *PartyIdentification2) HasDateOfBirth() bool {
	if o != nil && !common.IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *PartyIdentification2) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PartyIdentification2) GetFirstName() string {
	if o == nil || common.IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification2) GetFirstNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PartyIdentification2) HasFirstName() bool {
	if o != nil && !common.IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PartyIdentification2) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFullName returns the FullName field value
func (o *PartyIdentification2) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *PartyIdentification2) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *PartyIdentification2) SetFullName(v string) {
	o.FullName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PartyIdentification2) GetLastName() string {
	if o == nil || common.IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification2) GetLastNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PartyIdentification2) HasLastName() bool {
	if o != nil && !common.IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PartyIdentification2) SetLastName(v string) {
	o.LastName = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PartyIdentification2) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification2) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PartyIdentification2) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PartyIdentification2) SetReference(v string) {
	o.Reference = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PartyIdentification2) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification2) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PartyIdentification2) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PartyIdentification2) SetType(v string) {
	o.Type = &v
}

func (o PartyIdentification2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartyIdentification2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	if !common.IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	toSerialize["fullName"] = o.FullName
	if !common.IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePartyIdentification2 struct {
	value *PartyIdentification2
	isSet bool
}

func (v NullablePartyIdentification2) Get() *PartyIdentification2 {
	return v.value
}

func (v *NullablePartyIdentification2) Set(val *PartyIdentification2) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyIdentification2) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyIdentification2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyIdentification2(val *PartyIdentification2) *NullablePartyIdentification2 {
	return &NullablePartyIdentification2{value: val, isSet: true}
}

func (v NullablePartyIdentification2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyIdentification2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PartyIdentification2) isValidType() bool {
	var allowedEnumValues = []string{"individual", "organization", "unknown"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
