/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// checks if the FindTerminalResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FindTerminalResponse{}

// FindTerminalResponse struct for FindTerminalResponse
type FindTerminalResponse struct {
	// The company account that the terminal is associated with. If this is the only account level shown in the response, the terminal is assigned to the inventory of the company account.
	CompanyAccount string `json:"companyAccount"`
	// The merchant account that the terminal is associated with. If the response doesn't contain a `store` the terminal is assigned to this merchant account.
	MerchantAccount *string `json:"merchantAccount,omitempty"`
	// Boolean that indicates if the terminal is assigned to the merchant inventory. This is returned when the terminal is assigned to a merchant account.  - If **true**, this indicates that the terminal is in the merchant inventory. This also means that the terminal cannot be boarded.  - If **false**, this indicates that the terminal is assigned to the merchant account as an in-store terminal. This means that the terminal is ready to be boarded, or is already boarded.
	MerchantInventory *bool `json:"merchantInventory,omitempty"`
	// The store code of the store that the terminal is assigned to.
	Store *string `json:"store,omitempty"`
	// The unique terminal ID.
	Terminal string `json:"terminal"`
}

// NewFindTerminalResponse instantiates a new FindTerminalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindTerminalResponse(companyAccount string, terminal string) *FindTerminalResponse {
	this := FindTerminalResponse{}
	this.CompanyAccount = companyAccount
	this.Terminal = terminal
	return &this
}

// NewFindTerminalResponseWithDefaults instantiates a new FindTerminalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindTerminalResponseWithDefaults() *FindTerminalResponse {
	this := FindTerminalResponse{}
	return &this
}

// GetCompanyAccount returns the CompanyAccount field value
func (o *FindTerminalResponse) GetCompanyAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyAccount
}

// GetCompanyAccountOk returns a tuple with the CompanyAccount field value
// and a boolean to check if the value has been set.
func (o *FindTerminalResponse) GetCompanyAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyAccount, true
}

// SetCompanyAccount sets field value
func (o *FindTerminalResponse) SetCompanyAccount(v string) {
	o.CompanyAccount = v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *FindTerminalResponse) GetMerchantAccount() string {
	if o == nil || common.IsNil(o.MerchantAccount) {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindTerminalResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccount) {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *FindTerminalResponse) HasMerchantAccount() bool {
	if o != nil && !common.IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *FindTerminalResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetMerchantInventory returns the MerchantInventory field value if set, zero value otherwise.
func (o *FindTerminalResponse) GetMerchantInventory() bool {
	if o == nil || common.IsNil(o.MerchantInventory) {
		var ret bool
		return ret
	}
	return *o.MerchantInventory
}

// GetMerchantInventoryOk returns a tuple with the MerchantInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindTerminalResponse) GetMerchantInventoryOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantInventory) {
		return nil, false
	}
	return o.MerchantInventory, true
}

// HasMerchantInventory returns a boolean if a field has been set.
func (o *FindTerminalResponse) HasMerchantInventory() bool {
	if o != nil && !common.IsNil(o.MerchantInventory) {
		return true
	}

	return false
}

// SetMerchantInventory gets a reference to the given bool and assigns it to the MerchantInventory field.
func (o *FindTerminalResponse) SetMerchantInventory(v bool) {
	o.MerchantInventory = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *FindTerminalResponse) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindTerminalResponse) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *FindTerminalResponse) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *FindTerminalResponse) SetStore(v string) {
	o.Store = &v
}

// GetTerminal returns the Terminal field value
func (o *FindTerminalResponse) GetTerminal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Terminal
}

// GetTerminalOk returns a tuple with the Terminal field value
// and a boolean to check if the value has been set.
func (o *FindTerminalResponse) GetTerminalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terminal, true
}

// SetTerminal sets field value
func (o *FindTerminalResponse) SetTerminal(v string) {
	o.Terminal = v
}

func (o FindTerminalResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindTerminalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyAccount"] = o.CompanyAccount
	if !common.IsNil(o.MerchantAccount) {
		toSerialize["merchantAccount"] = o.MerchantAccount
	}
	if !common.IsNil(o.MerchantInventory) {
		toSerialize["merchantInventory"] = o.MerchantInventory
	}
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	toSerialize["terminal"] = o.Terminal
	return toSerialize, nil
}

type NullableFindTerminalResponse struct {
	value *FindTerminalResponse
	isSet bool
}

func (v NullableFindTerminalResponse) Get() *FindTerminalResponse {
	return v.value
}

func (v *NullableFindTerminalResponse) Set(val *FindTerminalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFindTerminalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFindTerminalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindTerminalResponse(val *FindTerminalResponse) *NullableFindTerminalResponse {
	return &NullableFindTerminalResponse{value: val, isSet: true}
}

func (v NullableFindTerminalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindTerminalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
