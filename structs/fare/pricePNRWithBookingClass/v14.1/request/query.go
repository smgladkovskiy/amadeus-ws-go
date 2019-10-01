package Fare_PricePNRWithBookingClassRequest_v14_1 // tpcbrq141

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRQ_14_1_1A Fare_PricePNRWithBookingClass"`

	PricingOptionGroup []*PricingOptionGroup `xml:"pricingOptionGroup"` // maxOccurs="999"
}

type PricingOptionGroup struct {
	PricingOptionKey *PricingOptionKey `xml:"pricingOptionKey"`

	OptionDetail *AttributeType `xml:"optionDetail,omitempty"`

	CarrierInformation *TransportIdentifierType `xml:"carrierInformation,omitempty"`

	Currency *CurrenciesType `xml:"currency,omitempty"`

	PenDisInformation *DiscountAndPenaltyInformationType `xml:"penDisInformation,omitempty"`

	MonetaryInformation *MonetaryInformationType `xml:"monetaryInformation,omitempty"`

	TaxInformation []*DutyTaxFeeDetailsType `xml:"taxInformation,omitempty"` // maxOccurs="99"

	DateInformation []*StructuredDateTimeInformationType `xml:"dateInformation,omitempty"` // maxOccurs="2"

	FrequentFlyerInformation *FrequentTravellerIdentificationCodeType `xml:"frequentFlyerInformation,omitempty"`

	FormOfPaymentInformation *FormOfPaymentType `xml:"formOfPaymentInformation,omitempty"`

	LocationInformation *PlaceLocationIdentificationType `xml:"locationInformation,omitempty"`

	PaxSegTstReference *ReferenceInfoType `xml:"paxSegTstReference,omitempty"`
}

//
// Complex structs
//

type AttributeInformationTypeU struct {
	// Used for attribute value rather than attributeType
	AttributeType formats.AlphaNumericString_Length1To25 `xml:"attributeType"`

	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"`
}

type AttributeType struct {
	// Details for the message criteria (name, value).
	CriteriaDetails []*AttributeInformationTypeU `xml:"criteriaDetails"` // maxOccurs="99"
}

type CompanyIdentificationTypeI struct {
	OtherCompany formats.AlphaNumericString_Length1To35 `xml:"otherCompany,omitempty"`
}

type CurrenciesType struct {
	FirstCurrencyDetails *CurrencyDetailsTypeU `xml:"firstCurrencyDetails,omitempty"`
}

type CurrencyDetailsTypeU struct {
	CurrencyQualifier formats.AlphaNumericString_Length1To3 `xml:"currencyQualifier"`

	CurrencyIsoCode formats.AlphaNumericString_Length1To3 `xml:"currencyIsoCode,omitempty"`
}

type DiscountAndPenaltyInformationType struct {
	DiscountPenaltyQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"discountPenaltyQualifier,omitempty"`

	DiscountPenaltyDetails []*DiscountPenaltyMonetaryInformationType `xml:"discountPenaltyDetails,omitempty"` // maxOccurs="9"
}

type DiscountPenaltyMonetaryInformationType struct {
	Function formats.AMA_EDICodesetType_Length1to3 `xml:"function,omitempty"`

	AmountType formats.AMA_EDICodesetType_Length1to3 `xml:"amountType,omitempty"`

	Amount formats.AlphaNumericString_Length1To18 `xml:"amount,omitempty"`

	Rate formats.AlphaNumericString_Length1To35 `xml:"rate,omitempty"`

	Currency formats.AMA_EDICodesetType_Length1to3 `xml:"currency,omitempty"`
}

type DutyTaxFeeAccountDetailType struct {
	IsoCountry formats.AlphaNumericString_Length1To6 `xml:"isoCountry"`
}

type DutyTaxFeeDetailType struct {
	TaxRate formats.AlphaNumericString_Length1To17 `xml:"taxRate,omitempty"`

	TaxValueQualifier formats.AlphaNumericString_Length1To1 `xml:"taxValueQualifier,omitempty"`
}

type DutyTaxFeeDetailsType struct {
	TaxQualifier formats.AlphaNumericString_Length1To3 `xml:"taxQualifier"`

	TaxType *DutyTaxFeeAccountDetailType `xml:"taxType,omitempty"`

	TaxNature formats.AlphaNumericString_Length1To15 `xml:"taxNature,omitempty"`

	TaxData *DutyTaxFeeDetailType `xml:"taxData,omitempty"`
}

type FormOfPaymentDetailsType struct {
	Type formats.AlphaNumericString_Length1To10 `xml:"type"`

	Amount *formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	CreditCardNumber formats.AlphaNumericString_Length1To35 `xml:"creditCardNumber,omitempty"`
}

type FormOfPaymentType struct {
	// Details on the form of payment
	FormOfPayment *FormOfPaymentDetailsType `xml:"formOfPayment"`

	OtherFormOfPayment []*FormOfPaymentDetailsType `xml:"otherFormOfPayment,omitempty"` // maxOccurs="2"
}

type FrequentTravellerIdentificationCodeType struct {
	// Frequent Traveller Info
	FrequentTravellerDetails []*FrequentTravellerIdentificationType `xml:"frequentTravellerDetails"` // maxOccurs="99"
}

type FrequentTravellerIdentificationType struct {
	// Carrier where the FQTV is registered.
	Carrier formats.AlphaNumericString_Length1To35 `xml:"carrier,omitempty"`

	Number formats.AlphaNumericString_Length1To28 `xml:"number,omitempty"`

	// To specify a Tier linked to the FQTV
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`

	// For example : priority code
	PriorityCode formats.AlphaNumericString_Length1To12 `xml:"priorityCode,omitempty"`
}

type MonetaryInformationDetailsType struct {
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount *formats.NumericDecimal_Length1To35 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"`
}

type MonetaryInformationType struct {
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"` // maxOccurs="19"
}

type PlaceLocationIdentificationType struct {
	LocationType formats.AlphaNumericString_Length1To3 `xml:"locationType"`

	FirstLocationDetails *RelatedLocationOneIdentificationType `xml:"firstLocationDetails,omitempty"`

	SecondLocationDetails *RelatedLocationTwoIdentificationType `xml:"secondLocationDetails,omitempty"`
}

type PricingOptionKey struct {
	PricingOptionKey formats.AlphaNumericString_Length1To3 `xml:"pricingOptionKey"`
}

type ReferenceInfoType struct {
	ReferenceDetails []*ReferencingDetailsType `xml:"referenceDetails,omitempty"` // maxOccurs="99"
}

type ReferencingDetailsType struct {
	Type formats.AlphaNumericString_Length1To10 `xml:"type,omitempty"`

	Value formats.AlphaNumericString_Length1To60 `xml:"value,omitempty"`
}

type RelatedLocationOneIdentificationType struct {
	Code formats.AlphaNumericString_Length1To25 `xml:"code,omitempty"`
}

type RelatedLocationTwoIdentificationType struct {
	Code formats.AlphaNumericString_Length1To25 `xml:"code,omitempty"`
}

type StructuredDateTimeInformationType struct {
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeType struct {
	// Year number.
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type TransportIdentifierType struct {
	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"`
}
