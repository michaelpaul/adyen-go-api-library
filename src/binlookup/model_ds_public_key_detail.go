/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the DSPublicKeyDetail type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DSPublicKeyDetail{}

// DSPublicKeyDetail struct for DSPublicKeyDetail
type DSPublicKeyDetail struct {
	// Card brand.
	Brand *string `json:"brand,omitempty"`
	// Directory Server (DS) identifier.
	DirectoryServerId *string `json:"directoryServerId,omitempty"`
	// The version of the mobile 3D Secure 2 SDK. For the possible values, refer to the versions in [Adyen 3DS2 Android](https://github.com/Adyen/adyen-3ds2-android/releases) and [Adyen 3DS2 iOS](https://github.com/Adyen/adyen-3ds2-ios/releases).
	FromSDKVersion *string `json:"fromSDKVersion,omitempty"`
	// Public key. The 3D Secure 2 SDK encrypts the device information by using the DS public key.
	PublicKey *string `json:"publicKey,omitempty"`
}

// NewDSPublicKeyDetail instantiates a new DSPublicKeyDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDSPublicKeyDetail() *DSPublicKeyDetail {
	this := DSPublicKeyDetail{}
	return &this
}

// NewDSPublicKeyDetailWithDefaults instantiates a new DSPublicKeyDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDSPublicKeyDetailWithDefaults() *DSPublicKeyDetail {
	this := DSPublicKeyDetail{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *DSPublicKeyDetail) GetBrand() string {
	if o == nil || common.IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSPublicKeyDetail) GetBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *DSPublicKeyDetail) HasBrand() bool {
	if o != nil && !common.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *DSPublicKeyDetail) SetBrand(v string) {
	o.Brand = &v
}

// GetDirectoryServerId returns the DirectoryServerId field value if set, zero value otherwise.
func (o *DSPublicKeyDetail) GetDirectoryServerId() string {
	if o == nil || common.IsNil(o.DirectoryServerId) {
		var ret string
		return ret
	}
	return *o.DirectoryServerId
}

// GetDirectoryServerIdOk returns a tuple with the DirectoryServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSPublicKeyDetail) GetDirectoryServerIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DirectoryServerId) {
		return nil, false
	}
	return o.DirectoryServerId, true
}

// HasDirectoryServerId returns a boolean if a field has been set.
func (o *DSPublicKeyDetail) HasDirectoryServerId() bool {
	if o != nil && !common.IsNil(o.DirectoryServerId) {
		return true
	}

	return false
}

// SetDirectoryServerId gets a reference to the given string and assigns it to the DirectoryServerId field.
func (o *DSPublicKeyDetail) SetDirectoryServerId(v string) {
	o.DirectoryServerId = &v
}

// GetFromSDKVersion returns the FromSDKVersion field value if set, zero value otherwise.
func (o *DSPublicKeyDetail) GetFromSDKVersion() string {
	if o == nil || common.IsNil(o.FromSDKVersion) {
		var ret string
		return ret
	}
	return *o.FromSDKVersion
}

// GetFromSDKVersionOk returns a tuple with the FromSDKVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSPublicKeyDetail) GetFromSDKVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.FromSDKVersion) {
		return nil, false
	}
	return o.FromSDKVersion, true
}

// HasFromSDKVersion returns a boolean if a field has been set.
func (o *DSPublicKeyDetail) HasFromSDKVersion() bool {
	if o != nil && !common.IsNil(o.FromSDKVersion) {
		return true
	}

	return false
}

// SetFromSDKVersion gets a reference to the given string and assigns it to the FromSDKVersion field.
func (o *DSPublicKeyDetail) SetFromSDKVersion(v string) {
	o.FromSDKVersion = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *DSPublicKeyDetail) GetPublicKey() string {
	if o == nil || common.IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSPublicKeyDetail) GetPublicKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *DSPublicKeyDetail) HasPublicKey() bool {
	if o != nil && !common.IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *DSPublicKeyDetail) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o DSPublicKeyDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DSPublicKeyDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !common.IsNil(o.DirectoryServerId) {
		toSerialize["directoryServerId"] = o.DirectoryServerId
	}
	if !common.IsNil(o.FromSDKVersion) {
		toSerialize["fromSDKVersion"] = o.FromSDKVersion
	}
	if !common.IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullableDSPublicKeyDetail struct {
	value *DSPublicKeyDetail
	isSet bool
}

func (v NullableDSPublicKeyDetail) Get() *DSPublicKeyDetail {
	return v.value
}

func (v *NullableDSPublicKeyDetail) Set(val *DSPublicKeyDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableDSPublicKeyDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableDSPublicKeyDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDSPublicKeyDetail(val *DSPublicKeyDetail) *NullableDSPublicKeyDetail {
	return &NullableDSPublicKeyDetail{value: val, isSet: true}
}

func (v NullableDSPublicKeyDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDSPublicKeyDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
