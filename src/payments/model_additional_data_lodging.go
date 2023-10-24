/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the AdditionalDataLodging type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataLodging{}

// AdditionalDataLodging struct for AdditionalDataLodging
type AdditionalDataLodging struct {
	// The arrival date. * Date format: **yyyyMmDd**. For example, for 2023 April 22, **20230422**.
	LodgingCheckInDate *string `json:"lodging.checkInDate,omitempty"`
	// The departure date. * Date format: **yyyyMmDd**. For example, for 2023 April 22, **20230422**.
	LodgingCheckOutDate *string `json:"lodging.checkOutDate,omitempty"`
	// The toll-free phone number for the lodging. * Format: numeric * Max length: 17 characters. * For US and CA numbers must be 10 characters in length * Must not start with a space * Must not be all zeros * Must not contain any special characters such as + or -
	LodgingCustomerServiceTollFreeNumber *string `json:"lodging.customerServiceTollFreeNumber,omitempty"`
	// Identifies that the facility complies with the Hotel and Motel Fire Safety Act of 1990. Must be 'Y' or 'N'. * Format: alphabetic * Max length: 1 character
	LodgingFireSafetyActIndicator *string `json:"lodging.fireSafetyActIndicator,omitempty"`
	// The folio cash advances, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: numeric * Max length: 12 characters
	LodgingFolioCashAdvances *string `json:"lodging.folioCashAdvances,omitempty"`
	// The card acceptor’s internal invoice or billing ID reference number. * Max length: 25 characters. * Must not start with a space * Must not be all zeros
	LodgingFolioNumber *string `json:"lodging.folioNumber,omitempty"`
	// Any charges for food and beverages associated with the booking, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: numeric * Max length: 12 characters
	LodgingFoodBeverageCharges *string `json:"lodging.foodBeverageCharges,omitempty"`
	// Indicates if the customer didn't check in for their booking.  Possible values:  * **Y**: the customer didn't check in  * **N**: the customer checked in
	LodgingNoShowIndicator *string `json:"lodging.noShowIndicator,omitempty"`
	// The prepaid expenses for the booking. * Format: numeric * Max length: 12 characters
	LodgingPrepaidExpenses *string `json:"lodging.prepaidExpenses,omitempty"`
	// The lodging property location's phone number. * Format: numeric. * Min length: 10 characters * Max length: 17 characters * For US and CA numbers must be 10 characters in length * Must not start with a space * Must not be all zeros * Must not contain any special characters such as + or -
	LodgingPropertyPhoneNumber *string `json:"lodging.propertyPhoneNumber,omitempty"`
	// The total number of nights the room is booked for. * Format: numeric * Must be a number between 0 and 99 * Max length: 2 characters
	LodgingRoom1NumberOfNights *string `json:"lodging.room1.numberOfNights,omitempty"`
	// The rate for the room, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: numeric * Max length: 12 characters * Must not be a negative number
	LodgingRoom1Rate *string `json:"lodging.room1.rate,omitempty"`
	// The total room tax amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: numeric * Max length: 12 characters * Must not be a negative number
	LodgingTotalRoomTax *string `json:"lodging.totalRoomTax,omitempty"`
	// The total tax amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes). * Format: numeric * Max length: 12 characters * Must not be a negative number
	LodgingTotalTax *string `json:"lodging.totalTax,omitempty"`
	// The number of nights. This should be included in the auth message. * Format: numeric * Max length: 2 characters
	TravelEntertainmentAuthDataDuration *string `json:"travelEntertainmentAuthData.duration,omitempty"`
	// Indicates what market-specific dataset will be submitted. Must be 'H' for Hotel. This should be included in the auth message.  * Format: alphanumeric * Max length: 1 character
	TravelEntertainmentAuthDataMarket *string `json:"travelEntertainmentAuthData.market,omitempty"`
}

// NewAdditionalDataLodging instantiates a new AdditionalDataLodging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataLodging() *AdditionalDataLodging {
	this := AdditionalDataLodging{}
	return &this
}

// NewAdditionalDataLodgingWithDefaults instantiates a new AdditionalDataLodging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataLodgingWithDefaults() *AdditionalDataLodging {
	this := AdditionalDataLodging{}
	return &this
}

// GetLodgingCheckInDate returns the LodgingCheckInDate field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingCheckInDate() string {
	if o == nil || common.IsNil(o.LodgingCheckInDate) {
		var ret string
		return ret
	}
	return *o.LodgingCheckInDate
}

// GetLodgingCheckInDateOk returns a tuple with the LodgingCheckInDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingCheckInDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingCheckInDate) {
		return nil, false
	}
	return o.LodgingCheckInDate, true
}

// HasLodgingCheckInDate returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingCheckInDate() bool {
	if o != nil && !common.IsNil(o.LodgingCheckInDate) {
		return true
	}

	return false
}

// SetLodgingCheckInDate gets a reference to the given string and assigns it to the LodgingCheckInDate field.
func (o *AdditionalDataLodging) SetLodgingCheckInDate(v string) {
	o.LodgingCheckInDate = &v
}

// GetLodgingCheckOutDate returns the LodgingCheckOutDate field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingCheckOutDate() string {
	if o == nil || common.IsNil(o.LodgingCheckOutDate) {
		var ret string
		return ret
	}
	return *o.LodgingCheckOutDate
}

// GetLodgingCheckOutDateOk returns a tuple with the LodgingCheckOutDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingCheckOutDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingCheckOutDate) {
		return nil, false
	}
	return o.LodgingCheckOutDate, true
}

// HasLodgingCheckOutDate returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingCheckOutDate() bool {
	if o != nil && !common.IsNil(o.LodgingCheckOutDate) {
		return true
	}

	return false
}

// SetLodgingCheckOutDate gets a reference to the given string and assigns it to the LodgingCheckOutDate field.
func (o *AdditionalDataLodging) SetLodgingCheckOutDate(v string) {
	o.LodgingCheckOutDate = &v
}

// GetLodgingCustomerServiceTollFreeNumber returns the LodgingCustomerServiceTollFreeNumber field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingCustomerServiceTollFreeNumber() string {
	if o == nil || common.IsNil(o.LodgingCustomerServiceTollFreeNumber) {
		var ret string
		return ret
	}
	return *o.LodgingCustomerServiceTollFreeNumber
}

// GetLodgingCustomerServiceTollFreeNumberOk returns a tuple with the LodgingCustomerServiceTollFreeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingCustomerServiceTollFreeNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingCustomerServiceTollFreeNumber) {
		return nil, false
	}
	return o.LodgingCustomerServiceTollFreeNumber, true
}

// HasLodgingCustomerServiceTollFreeNumber returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingCustomerServiceTollFreeNumber() bool {
	if o != nil && !common.IsNil(o.LodgingCustomerServiceTollFreeNumber) {
		return true
	}

	return false
}

// SetLodgingCustomerServiceTollFreeNumber gets a reference to the given string and assigns it to the LodgingCustomerServiceTollFreeNumber field.
func (o *AdditionalDataLodging) SetLodgingCustomerServiceTollFreeNumber(v string) {
	o.LodgingCustomerServiceTollFreeNumber = &v
}

// GetLodgingFireSafetyActIndicator returns the LodgingFireSafetyActIndicator field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingFireSafetyActIndicator() string {
	if o == nil || common.IsNil(o.LodgingFireSafetyActIndicator) {
		var ret string
		return ret
	}
	return *o.LodgingFireSafetyActIndicator
}

// GetLodgingFireSafetyActIndicatorOk returns a tuple with the LodgingFireSafetyActIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingFireSafetyActIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingFireSafetyActIndicator) {
		return nil, false
	}
	return o.LodgingFireSafetyActIndicator, true
}

// HasLodgingFireSafetyActIndicator returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingFireSafetyActIndicator() bool {
	if o != nil && !common.IsNil(o.LodgingFireSafetyActIndicator) {
		return true
	}

	return false
}

// SetLodgingFireSafetyActIndicator gets a reference to the given string and assigns it to the LodgingFireSafetyActIndicator field.
func (o *AdditionalDataLodging) SetLodgingFireSafetyActIndicator(v string) {
	o.LodgingFireSafetyActIndicator = &v
}

// GetLodgingFolioCashAdvances returns the LodgingFolioCashAdvances field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingFolioCashAdvances() string {
	if o == nil || common.IsNil(o.LodgingFolioCashAdvances) {
		var ret string
		return ret
	}
	return *o.LodgingFolioCashAdvances
}

// GetLodgingFolioCashAdvancesOk returns a tuple with the LodgingFolioCashAdvances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingFolioCashAdvancesOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingFolioCashAdvances) {
		return nil, false
	}
	return o.LodgingFolioCashAdvances, true
}

// HasLodgingFolioCashAdvances returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingFolioCashAdvances() bool {
	if o != nil && !common.IsNil(o.LodgingFolioCashAdvances) {
		return true
	}

	return false
}

// SetLodgingFolioCashAdvances gets a reference to the given string and assigns it to the LodgingFolioCashAdvances field.
func (o *AdditionalDataLodging) SetLodgingFolioCashAdvances(v string) {
	o.LodgingFolioCashAdvances = &v
}

// GetLodgingFolioNumber returns the LodgingFolioNumber field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingFolioNumber() string {
	if o == nil || common.IsNil(o.LodgingFolioNumber) {
		var ret string
		return ret
	}
	return *o.LodgingFolioNumber
}

// GetLodgingFolioNumberOk returns a tuple with the LodgingFolioNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingFolioNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingFolioNumber) {
		return nil, false
	}
	return o.LodgingFolioNumber, true
}

// HasLodgingFolioNumber returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingFolioNumber() bool {
	if o != nil && !common.IsNil(o.LodgingFolioNumber) {
		return true
	}

	return false
}

// SetLodgingFolioNumber gets a reference to the given string and assigns it to the LodgingFolioNumber field.
func (o *AdditionalDataLodging) SetLodgingFolioNumber(v string) {
	o.LodgingFolioNumber = &v
}

// GetLodgingFoodBeverageCharges returns the LodgingFoodBeverageCharges field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingFoodBeverageCharges() string {
	if o == nil || common.IsNil(o.LodgingFoodBeverageCharges) {
		var ret string
		return ret
	}
	return *o.LodgingFoodBeverageCharges
}

// GetLodgingFoodBeverageChargesOk returns a tuple with the LodgingFoodBeverageCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingFoodBeverageChargesOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingFoodBeverageCharges) {
		return nil, false
	}
	return o.LodgingFoodBeverageCharges, true
}

// HasLodgingFoodBeverageCharges returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingFoodBeverageCharges() bool {
	if o != nil && !common.IsNil(o.LodgingFoodBeverageCharges) {
		return true
	}

	return false
}

// SetLodgingFoodBeverageCharges gets a reference to the given string and assigns it to the LodgingFoodBeverageCharges field.
func (o *AdditionalDataLodging) SetLodgingFoodBeverageCharges(v string) {
	o.LodgingFoodBeverageCharges = &v
}

// GetLodgingNoShowIndicator returns the LodgingNoShowIndicator field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingNoShowIndicator() string {
	if o == nil || common.IsNil(o.LodgingNoShowIndicator) {
		var ret string
		return ret
	}
	return *o.LodgingNoShowIndicator
}

// GetLodgingNoShowIndicatorOk returns a tuple with the LodgingNoShowIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingNoShowIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingNoShowIndicator) {
		return nil, false
	}
	return o.LodgingNoShowIndicator, true
}

// HasLodgingNoShowIndicator returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingNoShowIndicator() bool {
	if o != nil && !common.IsNil(o.LodgingNoShowIndicator) {
		return true
	}

	return false
}

// SetLodgingNoShowIndicator gets a reference to the given string and assigns it to the LodgingNoShowIndicator field.
func (o *AdditionalDataLodging) SetLodgingNoShowIndicator(v string) {
	o.LodgingNoShowIndicator = &v
}

// GetLodgingPrepaidExpenses returns the LodgingPrepaidExpenses field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingPrepaidExpenses() string {
	if o == nil || common.IsNil(o.LodgingPrepaidExpenses) {
		var ret string
		return ret
	}
	return *o.LodgingPrepaidExpenses
}

// GetLodgingPrepaidExpensesOk returns a tuple with the LodgingPrepaidExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingPrepaidExpensesOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingPrepaidExpenses) {
		return nil, false
	}
	return o.LodgingPrepaidExpenses, true
}

// HasLodgingPrepaidExpenses returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingPrepaidExpenses() bool {
	if o != nil && !common.IsNil(o.LodgingPrepaidExpenses) {
		return true
	}

	return false
}

// SetLodgingPrepaidExpenses gets a reference to the given string and assigns it to the LodgingPrepaidExpenses field.
func (o *AdditionalDataLodging) SetLodgingPrepaidExpenses(v string) {
	o.LodgingPrepaidExpenses = &v
}

// GetLodgingPropertyPhoneNumber returns the LodgingPropertyPhoneNumber field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingPropertyPhoneNumber() string {
	if o == nil || common.IsNil(o.LodgingPropertyPhoneNumber) {
		var ret string
		return ret
	}
	return *o.LodgingPropertyPhoneNumber
}

// GetLodgingPropertyPhoneNumberOk returns a tuple with the LodgingPropertyPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingPropertyPhoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingPropertyPhoneNumber) {
		return nil, false
	}
	return o.LodgingPropertyPhoneNumber, true
}

// HasLodgingPropertyPhoneNumber returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingPropertyPhoneNumber() bool {
	if o != nil && !common.IsNil(o.LodgingPropertyPhoneNumber) {
		return true
	}

	return false
}

// SetLodgingPropertyPhoneNumber gets a reference to the given string and assigns it to the LodgingPropertyPhoneNumber field.
func (o *AdditionalDataLodging) SetLodgingPropertyPhoneNumber(v string) {
	o.LodgingPropertyPhoneNumber = &v
}

// GetLodgingRoom1NumberOfNights returns the LodgingRoom1NumberOfNights field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingRoom1NumberOfNights() string {
	if o == nil || common.IsNil(o.LodgingRoom1NumberOfNights) {
		var ret string
		return ret
	}
	return *o.LodgingRoom1NumberOfNights
}

// GetLodgingRoom1NumberOfNightsOk returns a tuple with the LodgingRoom1NumberOfNights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingRoom1NumberOfNightsOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingRoom1NumberOfNights) {
		return nil, false
	}
	return o.LodgingRoom1NumberOfNights, true
}

// HasLodgingRoom1NumberOfNights returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingRoom1NumberOfNights() bool {
	if o != nil && !common.IsNil(o.LodgingRoom1NumberOfNights) {
		return true
	}

	return false
}

// SetLodgingRoom1NumberOfNights gets a reference to the given string and assigns it to the LodgingRoom1NumberOfNights field.
func (o *AdditionalDataLodging) SetLodgingRoom1NumberOfNights(v string) {
	o.LodgingRoom1NumberOfNights = &v
}

// GetLodgingRoom1Rate returns the LodgingRoom1Rate field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingRoom1Rate() string {
	if o == nil || common.IsNil(o.LodgingRoom1Rate) {
		var ret string
		return ret
	}
	return *o.LodgingRoom1Rate
}

// GetLodgingRoom1RateOk returns a tuple with the LodgingRoom1Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingRoom1RateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingRoom1Rate) {
		return nil, false
	}
	return o.LodgingRoom1Rate, true
}

// HasLodgingRoom1Rate returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingRoom1Rate() bool {
	if o != nil && !common.IsNil(o.LodgingRoom1Rate) {
		return true
	}

	return false
}

// SetLodgingRoom1Rate gets a reference to the given string and assigns it to the LodgingRoom1Rate field.
func (o *AdditionalDataLodging) SetLodgingRoom1Rate(v string) {
	o.LodgingRoom1Rate = &v
}

// GetLodgingTotalRoomTax returns the LodgingTotalRoomTax field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingTotalRoomTax() string {
	if o == nil || common.IsNil(o.LodgingTotalRoomTax) {
		var ret string
		return ret
	}
	return *o.LodgingTotalRoomTax
}

// GetLodgingTotalRoomTaxOk returns a tuple with the LodgingTotalRoomTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingTotalRoomTaxOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingTotalRoomTax) {
		return nil, false
	}
	return o.LodgingTotalRoomTax, true
}

// HasLodgingTotalRoomTax returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingTotalRoomTax() bool {
	if o != nil && !common.IsNil(o.LodgingTotalRoomTax) {
		return true
	}

	return false
}

// SetLodgingTotalRoomTax gets a reference to the given string and assigns it to the LodgingTotalRoomTax field.
func (o *AdditionalDataLodging) SetLodgingTotalRoomTax(v string) {
	o.LodgingTotalRoomTax = &v
}

// GetLodgingTotalTax returns the LodgingTotalTax field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetLodgingTotalTax() string {
	if o == nil || common.IsNil(o.LodgingTotalTax) {
		var ret string
		return ret
	}
	return *o.LodgingTotalTax
}

// GetLodgingTotalTaxOk returns a tuple with the LodgingTotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetLodgingTotalTaxOk() (*string, bool) {
	if o == nil || common.IsNil(o.LodgingTotalTax) {
		return nil, false
	}
	return o.LodgingTotalTax, true
}

// HasLodgingTotalTax returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasLodgingTotalTax() bool {
	if o != nil && !common.IsNil(o.LodgingTotalTax) {
		return true
	}

	return false
}

// SetLodgingTotalTax gets a reference to the given string and assigns it to the LodgingTotalTax field.
func (o *AdditionalDataLodging) SetLodgingTotalTax(v string) {
	o.LodgingTotalTax = &v
}

// GetTravelEntertainmentAuthDataDuration returns the TravelEntertainmentAuthDataDuration field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetTravelEntertainmentAuthDataDuration() string {
	if o == nil || common.IsNil(o.TravelEntertainmentAuthDataDuration) {
		var ret string
		return ret
	}
	return *o.TravelEntertainmentAuthDataDuration
}

// GetTravelEntertainmentAuthDataDurationOk returns a tuple with the TravelEntertainmentAuthDataDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetTravelEntertainmentAuthDataDurationOk() (*string, bool) {
	if o == nil || common.IsNil(o.TravelEntertainmentAuthDataDuration) {
		return nil, false
	}
	return o.TravelEntertainmentAuthDataDuration, true
}

// HasTravelEntertainmentAuthDataDuration returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasTravelEntertainmentAuthDataDuration() bool {
	if o != nil && !common.IsNil(o.TravelEntertainmentAuthDataDuration) {
		return true
	}

	return false
}

// SetTravelEntertainmentAuthDataDuration gets a reference to the given string and assigns it to the TravelEntertainmentAuthDataDuration field.
func (o *AdditionalDataLodging) SetTravelEntertainmentAuthDataDuration(v string) {
	o.TravelEntertainmentAuthDataDuration = &v
}

// GetTravelEntertainmentAuthDataMarket returns the TravelEntertainmentAuthDataMarket field value if set, zero value otherwise.
func (o *AdditionalDataLodging) GetTravelEntertainmentAuthDataMarket() string {
	if o == nil || common.IsNil(o.TravelEntertainmentAuthDataMarket) {
		var ret string
		return ret
	}
	return *o.TravelEntertainmentAuthDataMarket
}

// GetTravelEntertainmentAuthDataMarketOk returns a tuple with the TravelEntertainmentAuthDataMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataLodging) GetTravelEntertainmentAuthDataMarketOk() (*string, bool) {
	if o == nil || common.IsNil(o.TravelEntertainmentAuthDataMarket) {
		return nil, false
	}
	return o.TravelEntertainmentAuthDataMarket, true
}

// HasTravelEntertainmentAuthDataMarket returns a boolean if a field has been set.
func (o *AdditionalDataLodging) HasTravelEntertainmentAuthDataMarket() bool {
	if o != nil && !common.IsNil(o.TravelEntertainmentAuthDataMarket) {
		return true
	}

	return false
}

// SetTravelEntertainmentAuthDataMarket gets a reference to the given string and assigns it to the TravelEntertainmentAuthDataMarket field.
func (o *AdditionalDataLodging) SetTravelEntertainmentAuthDataMarket(v string) {
	o.TravelEntertainmentAuthDataMarket = &v
}

func (o AdditionalDataLodging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataLodging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LodgingCheckInDate) {
		toSerialize["lodging.checkInDate"] = o.LodgingCheckInDate
	}
	if !common.IsNil(o.LodgingCheckOutDate) {
		toSerialize["lodging.checkOutDate"] = o.LodgingCheckOutDate
	}
	if !common.IsNil(o.LodgingCustomerServiceTollFreeNumber) {
		toSerialize["lodging.customerServiceTollFreeNumber"] = o.LodgingCustomerServiceTollFreeNumber
	}
	if !common.IsNil(o.LodgingFireSafetyActIndicator) {
		toSerialize["lodging.fireSafetyActIndicator"] = o.LodgingFireSafetyActIndicator
	}
	if !common.IsNil(o.LodgingFolioCashAdvances) {
		toSerialize["lodging.folioCashAdvances"] = o.LodgingFolioCashAdvances
	}
	if !common.IsNil(o.LodgingFolioNumber) {
		toSerialize["lodging.folioNumber"] = o.LodgingFolioNumber
	}
	if !common.IsNil(o.LodgingFoodBeverageCharges) {
		toSerialize["lodging.foodBeverageCharges"] = o.LodgingFoodBeverageCharges
	}
	if !common.IsNil(o.LodgingNoShowIndicator) {
		toSerialize["lodging.noShowIndicator"] = o.LodgingNoShowIndicator
	}
	if !common.IsNil(o.LodgingPrepaidExpenses) {
		toSerialize["lodging.prepaidExpenses"] = o.LodgingPrepaidExpenses
	}
	if !common.IsNil(o.LodgingPropertyPhoneNumber) {
		toSerialize["lodging.propertyPhoneNumber"] = o.LodgingPropertyPhoneNumber
	}
	if !common.IsNil(o.LodgingRoom1NumberOfNights) {
		toSerialize["lodging.room1.numberOfNights"] = o.LodgingRoom1NumberOfNights
	}
	if !common.IsNil(o.LodgingRoom1Rate) {
		toSerialize["lodging.room1.rate"] = o.LodgingRoom1Rate
	}
	if !common.IsNil(o.LodgingTotalRoomTax) {
		toSerialize["lodging.totalRoomTax"] = o.LodgingTotalRoomTax
	}
	if !common.IsNil(o.LodgingTotalTax) {
		toSerialize["lodging.totalTax"] = o.LodgingTotalTax
	}
	if !common.IsNil(o.TravelEntertainmentAuthDataDuration) {
		toSerialize["travelEntertainmentAuthData.duration"] = o.TravelEntertainmentAuthDataDuration
	}
	if !common.IsNil(o.TravelEntertainmentAuthDataMarket) {
		toSerialize["travelEntertainmentAuthData.market"] = o.TravelEntertainmentAuthDataMarket
	}
	return toSerialize, nil
}

type NullableAdditionalDataLodging struct {
	value *AdditionalDataLodging
	isSet bool
}

func (v NullableAdditionalDataLodging) Get() *AdditionalDataLodging {
	return v.value
}

func (v *NullableAdditionalDataLodging) Set(val *AdditionalDataLodging) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataLodging) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataLodging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataLodging(val *AdditionalDataLodging) *NullableAdditionalDataLodging {
	return &NullableAdditionalDataLodging{value: val, isSet: true}
}

func (v NullableAdditionalDataLodging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataLodging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
