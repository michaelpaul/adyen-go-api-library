/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
	"fmt"
)

// DonationPaymentMethod - The type and required details of a payment method to use.
type DonationPaymentMethod struct {
	ApplePayDetails      *ApplePayDetails
	CardDetails          *CardDetails
	GooglePayDetails     *GooglePayDetails
	IdealDetails         *IdealDetails
	PayWithGoogleDetails *PayWithGoogleDetails
}

// ApplePayDetailsAsDonationPaymentMethod is a convenience function that returns ApplePayDetails wrapped in DonationPaymentMethod
func ApplePayDetailsAsDonationPaymentMethod(v *ApplePayDetails) DonationPaymentMethod {
	return DonationPaymentMethod{
		ApplePayDetails: v,
	}
}

// CardDetailsAsDonationPaymentMethod is a convenience function that returns CardDetails wrapped in DonationPaymentMethod
func CardDetailsAsDonationPaymentMethod(v *CardDetails) DonationPaymentMethod {
	return DonationPaymentMethod{
		CardDetails: v,
	}
}

// GooglePayDetailsAsDonationPaymentMethod is a convenience function that returns GooglePayDetails wrapped in DonationPaymentMethod
func GooglePayDetailsAsDonationPaymentMethod(v *GooglePayDetails) DonationPaymentMethod {
	return DonationPaymentMethod{
		GooglePayDetails: v,
	}
}

// IdealDetailsAsDonationPaymentMethod is a convenience function that returns IdealDetails wrapped in DonationPaymentMethod
func IdealDetailsAsDonationPaymentMethod(v *IdealDetails) DonationPaymentMethod {
	return DonationPaymentMethod{
		IdealDetails: v,
	}
}

// PayWithGoogleDetailsAsDonationPaymentMethod is a convenience function that returns PayWithGoogleDetails wrapped in DonationPaymentMethod
func PayWithGoogleDetailsAsDonationPaymentMethod(v *PayWithGoogleDetails) DonationPaymentMethod {
	return DonationPaymentMethod{
		PayWithGoogleDetails: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DonationPaymentMethod) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApplePayDetails
	err = json.Unmarshal(data, &dst.ApplePayDetails)
	if err == nil {
		jsonApplePayDetails, _ := json.Marshal(dst.ApplePayDetails)
		if string(jsonApplePayDetails) == "{}" || !dst.ApplePayDetails.isValidType() { // empty struct
			dst.ApplePayDetails = nil
		} else {
			match++
		}
	} else {
		dst.ApplePayDetails = nil
	}

	// try to unmarshal data into CardDetails
	err = json.Unmarshal(data, &dst.CardDetails)
	if err == nil {
		jsonCardDetails, _ := json.Marshal(dst.CardDetails)
		if string(jsonCardDetails) == "{}" || !dst.CardDetails.isValidType() { // empty struct
			dst.CardDetails = nil
		} else {
			match++
		}
	} else {
		dst.CardDetails = nil
	}

	// try to unmarshal data into GooglePayDetails
	err = json.Unmarshal(data, &dst.GooglePayDetails)
	if err == nil {
		jsonGooglePayDetails, _ := json.Marshal(dst.GooglePayDetails)
		if string(jsonGooglePayDetails) == "{}" || !dst.GooglePayDetails.isValidType() { // empty struct
			dst.GooglePayDetails = nil
		} else {
			match++
		}
	} else {
		dst.GooglePayDetails = nil
	}

	// try to unmarshal data into IdealDetails
	err = json.Unmarshal(data, &dst.IdealDetails)
	if err == nil {
		jsonIdealDetails, _ := json.Marshal(dst.IdealDetails)
		if string(jsonIdealDetails) == "{}" || !dst.IdealDetails.isValidType() { // empty struct
			dst.IdealDetails = nil
		} else {
			match++
		}
	} else {
		dst.IdealDetails = nil
	}

	// try to unmarshal data into PayWithGoogleDetails
	err = json.Unmarshal(data, &dst.PayWithGoogleDetails)
	if err == nil {
		jsonPayWithGoogleDetails, _ := json.Marshal(dst.PayWithGoogleDetails)
		if string(jsonPayWithGoogleDetails) == "{}" || !dst.PayWithGoogleDetails.isValidType() { // empty struct
			dst.PayWithGoogleDetails = nil
		} else {
			match++
		}
	} else {
		dst.PayWithGoogleDetails = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApplePayDetails = nil
		dst.CardDetails = nil
		dst.GooglePayDetails = nil
		dst.IdealDetails = nil
		dst.PayWithGoogleDetails = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DonationPaymentMethod)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DonationPaymentMethod)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DonationPaymentMethod) MarshalJSON() ([]byte, error) {
	if src.ApplePayDetails != nil {
		return json.Marshal(&src.ApplePayDetails)
	}

	if src.CardDetails != nil {
		return json.Marshal(&src.CardDetails)
	}

	if src.GooglePayDetails != nil {
		return json.Marshal(&src.GooglePayDetails)
	}

	if src.IdealDetails != nil {
		return json.Marshal(&src.IdealDetails)
	}

	if src.PayWithGoogleDetails != nil {
		return json.Marshal(&src.PayWithGoogleDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DonationPaymentMethod) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApplePayDetails != nil {
		return obj.ApplePayDetails
	}

	if obj.CardDetails != nil {
		return obj.CardDetails
	}

	if obj.GooglePayDetails != nil {
		return obj.GooglePayDetails
	}

	if obj.IdealDetails != nil {
		return obj.IdealDetails
	}

	if obj.PayWithGoogleDetails != nil {
		return obj.PayWithGoogleDetails
	}

	// all schemas are nil
	return nil
}

type NullableDonationPaymentMethod struct {
	value *DonationPaymentMethod
	isSet bool
}

func (v NullableDonationPaymentMethod) Get() *DonationPaymentMethod {
	return v.value
}

func (v *NullableDonationPaymentMethod) Set(val *DonationPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDonationPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDonationPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDonationPaymentMethod(val *DonationPaymentMethod) *NullableDonationPaymentMethod {
	return &NullableDonationPaymentMethod{value: val, isSet: true}
}

func (v NullableDonationPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDonationPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
