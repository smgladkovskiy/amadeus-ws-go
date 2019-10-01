package Ticket_ProcessEDocRequest_v15_2 // tatreq152

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TATREQ_15_2_1A Ticket_ProcessEDoc"`

	MsgActionDetails *MessageActionDetailsTypeI `xml:"msgActionDetails"`

	// Travel agent details at level 0
	SysProvider *TicketAgentInfoTypeI `xml:"sysProvider,omitempty"`

	// Reservation details at level 0
	ReferenceInfo *ReservationControlInformationTypeI `xml:"referenceInfo,omitempty"`

	// Monetary information at level 0
	FareInfo *MonetaryInformationTypeI `xml:"fareInfo,omitempty"`

	// Pricing details at level 0
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"`

	// Origin - Destination at Level 0
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"`

	EnhancedPaxInfoCriteria *EnhancedTravellerInformationType_199238S `xml:"enhancedPaxInfoCriteria,omitempty"`

	// Travel details at Level 0
	Leg *TravelProductInformationTypeI `xml:"leg,omitempty"`

	// Form of payment at Level 0
	Fop *FormOfPaymentTypeI `xml:"fop,omitempty"`

	// Frequent traveller info at Level 0
	FrequentTravellerInfo *FrequentTravellerInformationTypeI `xml:"frequentTravellerInfo,omitempty"`

	// Tour code at Level 0
	TourInfo *TourInformationTypeI `xml:"tourInfo,omitempty"`

	// Document count
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits,omitempty"`

	// tax details at Level 0
	TaxInfo []*TaxTypeI `xml:"taxInfo,omitempty"` // maxOccurs="6"

	// Freeflow text at Level 0
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // maxOccurs="99"

	// form of identification at Level 0
	CustomerReference *ConsumerReferenceInformationTypeI `xml:"customerReference,omitempty"`

	// Fare details at Level 0
	FareDetails *FareInformationTypeI `xml:"fareDetails,omitempty"`

	// For MCO, reason for issuance codes
	PricingInfo *PricingTicketingSubsequentTypeI `xml:"pricingInfo,omitempty"`

	// Document group at level 1
	InfoGroup []*InfoGroup `xml:"infoGroup,omitempty"` // maxOccurs="99"

	// Passenger Group
	DocGroup []*DocGroup `xml:"docGroup,omitempty"` // maxOccurs="99"
}

type InfoGroup struct {
	// Document details of group 1
	DocInfo *TicketNumberTypeI `xml:"docInfo"`

	// coupon details of Group 1
	CouponInfo []*CouponInformationTypeI `xml:"couponInfo,omitempty"` // maxOccurs="4"

	// Originator details group 1
	OriginatorInfo *OriginatorOfRequestDetailsTypeI `xml:"originatorInfo,omitempty"`
}

type DocGroup struct {
	PaxInfo *TravellerInformationTypeI `xml:"paxInfo"`

	EnhancedPaxInfo *EnhancedTravellerInformationType `xml:"enhancedPaxInfo,omitempty"`

	// Travel agent details at level 2
	SysProvider *TicketAgentInfoTypeI `xml:"sysProvider,omitempty"`

	// Reservation details at level 2
	ReferenceInfo *ReservationControlInformationTypeI `xml:"referenceInfo,omitempty"`

	// Monetary information at level 2
	FareInfo *MonetaryInformationTypeI `xml:"fareInfo,omitempty"`

	// Form of payment at level 2
	Fop *FormOfPaymentTypeI `xml:"fop,omitempty"`

	// Pricing details at level 2
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"`

	// Origin - Destination at Level 2
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"`

	// Travel details at Level 2
	Leg *TravelProductInformationTypeI `xml:"leg,omitempty"`

	// Pricing subsequent at level 2
	PricingInfo *PricingTicketingSubsequentTypeI_51235S `xml:"pricingInfo,omitempty"`

	// Frequent traveller info at Level 2
	FrequentTravellerInfo *FrequentTravellerInformationTypeI `xml:"frequentTravellerInfo,omitempty"`

	// Tour code L2
	TourInfo *TourInformationTypeI `xml:"tourInfo,omitempty"`

	// Coupon count
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits,omitempty"`

	// tax details at Level 2
	TaxInfo []*TaxTypeI `xml:"taxInfo,omitempty"` // maxOccurs="5"

	// Document details at Level 2
	DocumentInformation []*DocumentInformationDetailsTypeI `xml:"documentInformation,omitempty"` // maxOccurs="99"

	// Freeflow text at Level 2
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // maxOccurs="9"

	// form of identification at Level 2
	CustomerReference *ConsumerReferenceInformationTypeI `xml:"customerReference,omitempty"`

	// Fare details at Level 2
	FareDetails *FareInformationTypeI `xml:"fareDetails,omitempty"`

	// Document Group
	DocDetailsGroup []*DocDetailsGroup `xml:"docDetailsGroup,omitempty"` // maxOccurs="99"

	// Structured fare calc group
	FareElementsGroup []*FareElementsGroup `xml:"fareElementsGroup,omitempty"` // maxOccurs="2"

	// Refund group
	RefundGroup []*RefundGroup `xml:"refundGroup,omitempty"` // maxOccurs="99"

	// Staff group
	StaffTravellerGroup *StaffTravellerGroup `xml:"staffTravellerGroup,omitempty"`

	// Exchanged information Group
	OriginalIssuanceGroup []*OriginalIssuanceGroup `xml:"originalIssuanceGroup,omitempty"` // maxOccurs="2"

	// Fee group
	CarrierFeeGroup []*CarrierFeeGroup `xml:"carrierFeeGroup,omitempty"` // maxOccurs="9"
}

type DocDetailsGroup struct {
	// Document details of group 3
	DocInfo *TicketNumberTypeI `xml:"docInfo"`

	// Pricing details at level 3
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"`

	// Freeflow text at Level 2
	TextInfo *InteractiveFreeTextTypeI `xml:"textInfo,omitempty"`

	// Coupon Group
	CouponGroup []*CouponGroup `xml:"couponGroup,omitempty"` // maxOccurs="4"
}

type CouponGroup struct {
	// Coupon details of group 4
	CouponInfo *CouponInformationTypeI `xml:"couponInfo"`

	// Travel details at Level 4
	Leg []*TravelProductInformationTypeI `xml:"leg,omitempty"` // maxOccurs="2"

	// Reservation details at level 4
	ReferenceInfo *ReservationControlInformationTypeI `xml:"referenceInfo,omitempty"`

	// Related product information at level 4
	BookingStatus *RelatedProductInformationTypeI `xml:"bookingStatus,omitempty"`

	// Pricing subsequent at L4
	PricingInfo *PricingTicketingSubsequentTypeI_51235S `xml:"pricingInfo,omitempty"`

	// Baggage Details
	BaggageInfo *ExcessBaggageTypeI `xml:"baggageInfo,omitempty"`

	// Frequent traveller info at Level 4
	FrequentTravellerInfo *FrequentTravellerInformationTypeI `xml:"frequentTravellerInfo,omitempty"`

	// Date information at level 4
	ValidityDates *DateAndTimeInformationTypeI `xml:"validityDates,omitempty"`

	// Freeflow text at Level 4
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // maxOccurs="9"

	// Pricing details at level 4
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"`
}

type FareElementsGroup struct {
	// Fare component information
	FareComponentInfo *FareComponentInformationTypeI `xml:"fareComponentInfo"`

	// structured Fare Calc Group6
	FareComponentsGroup []*FareComponentsGroup `xml:"fareComponentsGroup,omitempty"` // maxOccurs="99"

	// structured Fare Calc monetary information 1
	FareInfo *MonetaryInformationTypeI `xml:"fareInfo,omitempty"`

	// Structured Fare Calc taxe information
	TaxInfo []*TaxTypeI `xml:"taxInfo,omitempty"` // maxOccurs="99"

	// structured Fare Calc Conversion Info
	ConversionRate *ConversionRateTypeI `xml:"conversionRate,omitempty"`
}

type FareComponentsGroup struct {
	// structured Fare Calc quantity
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits"`

	// structured Fare Calc Group6
	PricedFareComponentsGroup []*PricedFareComponentsGroup `xml:"pricedFareComponentsGroup,omitempty"` // maxOccurs="99"
}

type PricedFareComponentsGroup struct {
	// structured Fare Calc item
	StructuredFareCalcItem *ItemNumberTypeI `xml:"structuredFareCalcItem"`

	// structured Fare Calc Group8
	FareCouponGroup []*FareCouponGroup `xml:"fareCouponGroup,omitempty"` // maxOccurs="99"

	// structured Fare Calc Monetary Information 2
	FareInfo *MonetaryInformationTypeI `xml:"fareInfo,omitempty"`

	// structured Fare Calc Pricing subsequent
	PricingInfo *PricingTicketingSubsequentTypeI_51235S `xml:"pricingInfo,omitempty"`

	// structured Fare Calculation details
	FareCalculationInfo *FareCalculationCodeDetailsTypeI `xml:"fareCalculationInfo,omitempty"`

	// structured Fare Calc rate
	FareRulesInfo *FareRulesInformationTypeI `xml:"fareRulesInfo,omitempty"`

	// Pricing details
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo"`

	// Fare details
	StructuredFareCalcFareDetails *FareInformationTypeI `xml:"structuredFareCalcFareDetails,omitempty"`
}

type FareCouponGroup struct {
	// structured Fare Calc Action
	CouponInfo *ActionDetailsTypeI `xml:"couponInfo"`

	Leg *TravelProductInformationTypeI `xml:"leg,omitempty"`
}

type RefundGroup struct {
	// Refund status
	RefundStatus *StatusTypeI `xml:"refundStatus"`

	// Refund traveller details
	RefundLeg []*TravelProductInformationTypeI `xml:"refundLeg,omitempty"` // maxOccurs="2"

	// refund Pricing details
	ProductInfo *PricingTicketingDetailsTypeI_51227S `xml:"productInfo,omitempty"`

	// Refund document details
	RefundDocumentInfo *TicketNumberTypeI `xml:"refundDocumentInfo,omitempty"`

	// Refund coupon details
	CouponInfo []*CouponInformationTypeI `xml:"couponInfo,omitempty"` // maxOccurs="9"

	// Refund pricing subsequent
	RefundPricingInfo *PricingTicketingSubsequentTypeI_51235S `xml:"refundPricingInfo,omitempty"`

	// Refund routing details
	RefundRoutingDetails []*RoutingInformationTypeI `xml:"refundRoutingDetails,omitempty"` // maxOccurs="99"

	// Refund product details
	RefundProductDetails *AdditionalProductDetailsTypeI `xml:"refundProductDetails,omitempty"`
}

type StaffTravellerGroup struct {
	// To indicate the details associated with a traveller's status
	TravellerPriorityInfo *TravellerPriorityDetailsTypeI `xml:"travellerPriorityInfo"`

	// Staff information
	SpecialDataInfo *SpecificDataInformationTypeI `xml:"specialDataInfo,omitempty"`
}

type OriginalIssuanceGroup struct {
	// Exchange originator details
	OfficeIdentification *AdditionalBusinessSourceInformationTypeI `xml:"officeIdentification"`

	// Exchange details
	ExchangeDocumentDetails *DocumentInformationDetailsTypeI `xml:"exchangeDocumentDetails,omitempty"`
}

type CarrierFeeGroup struct {
	// Fee type
	SelectionDetails *SelectionDetailsTypeI `xml:"selectionDetails"`

	// fee Sub Group
	CarrierFeeInfo []*CarrierFeeInfo `xml:"carrierFeeInfo"` // maxOccurs="99"
}

type CarrierFeeInfo struct {
	// fee sub type
	SpecialDataInfo *SpecificDataInformationTypeI `xml:"specialDataInfo"`

	// fee monetary Info
	FareInfo *MonetaryInformationTypeI `xml:"fareInfo"`

	// fee form of payment
	Fop *FormOfPaymentTypeI `xml:"fop"`

	// Fee taxes
	TaxInfo []*TaxTypeI `xml:"taxInfo,omitempty"` // maxOccurs="99"
}

//
// Complex structs
//

type ActionDetailsTypeI struct {
	LastItemsDetails []*ReferenceTypeI `xml:"lastItemsDetails,omitempty"` // maxOccurs="99"
}

type AdditionalBusinessSourceInformationTypeI struct {
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType"`

	OriginatorDetails *OriginatorIdentificationDetailsTypeI_83809C `xml:"originatorDetails,omitempty"`

	LocationDetails *LocationTypeI_83769C `xml:"locationDetails,omitempty"`
}

type AdditionalProductDetailsTypeI struct {
	// Additional details describing a specific means of transport
	LegDetails *AdditionalProductTypeI `xml:"legDetails,omitempty"`
}

type AdditionalProductTypeI struct {
	// UN/IATA code identifying type of aircraft (747, 737, etc.).
	// xmlType: AlphaNumericString_Length1To8
	Equipment string `xml:"equipment,omitempty"`

	// Number of stops enroute made in a journey.
	NumberOfStops *formats.NumericInteger_Length1To3 `xml:"numberOfStops,omitempty"`
}

type BaggageDetailsTypeI struct {
	// Total number of units.
	FreeAllowance *formats.NumericInteger_Length1To15 `xml:"freeAllowance,omitempty"`

	// Code to qualify unit as pieces or seats.
	// xmlType: AlphaNumericString_Length1To3
	QuantityCode string `xml:"quantityCode,omitempty"`

	// Code to qualify unit as pounds or kilos.
	// xmlType: AlphaNumericString_Length1To3
	UnitQualifier string `xml:"unitQualifier,omitempty"`
}

type CompanyIdentificationNumbersTypeI struct {
	// xmlType: AlphaNumericString_Length1To15
	Identifier string `xml:"identifier"`

	// xmlType: AlphaNumericString_Length1To15
	OtherIdentifier string `xml:"otherIdentifier,omitempty"`
}

type CompanyIdentificationTypeI struct {
	// xmlType: AlphaNumericString_Length1To35
	MarketingCompany string `xml:"marketingCompany,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	OperatingCompany string `xml:"operatingCompany,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	OtherCompany string `xml:"otherCompany,omitempty"`
}

type CompanyIdentificationTypeI_83773C struct {
	// Pricing System Code
	// xmlType: AlphaNumericString_Length1To35
	PricingSystemCode string `xml:"pricingSystemCode,omitempty"`
}

type ConsumerReferenceIdentificationTypeI struct {
	// To indicate type of passenger check-in identification number (FOID).
	// xmlType: AlphaNumericString_Length1To3
	ReferenceQualifier string `xml:"referenceQualifier"`

	// The actual form of identification number (FOID).
	// xmlType: AlphaNumericString_Length1To35
	ReferenceNumber string `xml:"referenceNumber,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	ReferencePartyName string `xml:"referencePartyName,omitempty"`
}

type ConsumerReferenceInformationTypeI struct {
	CustomerReferences []*ConsumerReferenceIdentificationTypeI `xml:"customerReferences"` // maxOccurs="20"
}

type ConversionRateDetailsTypeI struct {
	// Rate of Exchange
	// xmlType: NumericDecimal_Length1To18
	PricingAmount *float64 `xml:"pricingAmount,omitempty"`
}

type ConversionRateTypeI struct {
	ConversionRateDetails []*ConversionRateDetailsTypeI `xml:"conversionRateDetails"` // maxOccurs="20"
}

type CouponInformationDetailsTypeI struct {
	// Coupon number.
	// xmlType: AlphaNumericString_Length1To6
	CpnNumber string `xml:"cpnNumber,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	CpnStatus string `xml:"cpnStatus,omitempty"`

	// Coupon amount.
	// xmlType: NumericDecimal_Length1To18
	CpnAmount *float64 `xml:"cpnAmount,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	CpnExchangeMedia string `xml:"cpnExchangeMedia,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	SettlementAuthorization string `xml:"settlementAuthorization,omitempty"`

	// Involuntary indicator.
	// xmlType: AlphaNumericString_Length1To3
	VoluntaryIndicator string `xml:"voluntaryIndicator,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	CpnPreviousStatus string `xml:"cpnPreviousStatus,omitempty"`

	// In connection with coupon number.
	// xmlType: AlphaNumericString_Length1To6
	CpnInConnectionWith string `xml:"cpnInConnectionWith,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	CpnSequenceNumber string `xml:"cpnSequenceNumber,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	CpnInConnectionWithQualifier string `xml:"cpnInConnectionWithQualifier,omitempty"`
}

type CouponInformationTypeI struct {
	CouponDetails []*CouponInformationDetailsTypeI `xml:"couponDetails"` // maxOccurs="4"
}

type DataInformationTypeI struct {
	// The carrier fee application code
	// xmlType: AlphaNumericString_Length1To3
	Indicator string `xml:"indicator,omitempty"`
}

type DataTypeInformationTypeI struct {
	// To specify the type of data, i.e. staff reservation entitlement, status, type of travel, or ticket type
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type"`

	// Reservation entitlement/status, type of travel, ticket type
	// xmlType: AlphaNumericString_Length1To3
	StatusEvent string `xml:"statusEvent,omitempty"`
}

type DateAndTimeDetailsTypeI struct {
	// To identify type of date to follow.
	// xmlType: AlphaNumericString_Length1To3
	Qualifier string `xml:"qualifier,omitempty"`

	// Valid date (ddmmyy).
	// xmlType: AlphaNumericString_Length1To35
	Date string `xml:"date,omitempty"`

	// A time (hhmm).
	// xmlType: AlphaNumericString_Length1To4
	Time string `xml:"time,omitempty"`
}

type DateAndTimeInformationTypeI struct {
	DateAndTimeDetails []*DateAndTimeDetailsTypeI `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}

type DocumentDetailsTypeI struct {
	// Certificate number or original issue/ticket document number.
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"`

	// Original issue date (ddmmmyy)
	// xmlType: AlphaNumericString_Length1To35
	Date string `xml:"date,omitempty"`
}

type DocumentInformationDetailsTypeI struct {
	DocumentDetails *DocumentDetailsTypeI `xml:"documentDetails"`
}

type EnhancedTravellerInformationType struct {
	// Name attributes unique for one passenger.
	TravellerNameInfo *TravellerNameInfoType_276832C `xml:"travellerNameInfo,omitempty"`

	// 5 possible types of names, for 1 passenger.
	OtherPaxNamesDetails []*TravellerNameDetailsType `xml:"otherPaxNamesDetails"` // maxOccurs="5"
}

type EnhancedTravellerInformationType_199238S struct {
	// Name attributes unique for one passenger.
	TravellerNameInfo *TravellerNameInfoType `xml:"travellerNameInfo,omitempty"`

	// 5 possible types of names, for 1 passenger.
	OtherPaxNamesDetails []*TravellerNameDetailsType `xml:"otherPaxNamesDetails"` // maxOccurs="5"
}

type ExcessBaggageTypeI struct {
	BaggageDetails []*BaggageDetailsTypeI `xml:"baggageDetails,omitempty"` // maxOccurs="3"
}

type FareCalculationCodeDetailsTypeI struct {
	// Fare calculation descriptive indicators
	// xmlType: AlphaNumericString_Length1To3
	ChargeCategory string `xml:"chargeCategory,omitempty"`

	// Charge or fare calculation amount
	// xmlType: NumericDecimal_Length1To18
	Amount *float64 `xml:"amount,omitempty"`

	// May apply to a specific airport/city
	// xmlType: AlphaNumericString_Length1To25
	LocationCode []string `xml:"locationCode,omitempty"` // maxOccurs="2"

	// Percentage amount
	// xmlType: NumericDecimal_Length1To8
	Rate *float64 `xml:"rate,omitempty"`
}

type FareComponentDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	DataType string `xml:"dataType,omitempty"`

	// Fare component count
	Count *formats.NumericInteger_Length1To15 `xml:"count,omitempty"`

	// Price quote date
	// xmlType: AlphaNumericString_Length1To35
	PricingDate string `xml:"pricingDate,omitempty"`

	// Account code
	// xmlType: AlphaNumericString_Length1To35
	AccountCode string `xml:"accountCode,omitempty"`

	// Input designator
	// xmlType: AlphaNumericString_Length1To35
	InputDesignator string `xml:"inputDesignator,omitempty"`
}

type FareComponentInformationTypeI struct {
	FareComponentDetails *FareComponentDetailsTypeI `xml:"fareComponentDetails,omitempty"`

	// Ticket document number
	// xmlType: AlphaNumericString_Length1To35
	TicketNumber string `xml:"ticketNumber,omitempty"`
}

type FareInformationTypeI struct {
	// Specifies an industry defined fare component priced passenger type code (PTC).
	// xmlType: AlphaNumericString_Length1To3
	ValueQualifier string `xml:"valueQualifier,omitempty"`

	FareTypeGrouping *FareTypeGroupingInformationTypeI `xml:"fareTypeGrouping,omitempty"`
}

type FareRulesInformationTypeI struct {
	// Tariff identification
	// xmlType: AlphaNumericString_Length1To9
	TariffClassId string `xml:"tariffClassId,omitempty"`

	CompanyDetails *CompanyIdentificationTypeI_83773C `xml:"companyDetails,omitempty"`

	// Fare rule paragraph number code
	// xmlType: AlphaNumericString_Length1To7
	RuleSectionId []string `xml:"ruleSectionId,omitempty"` // maxOccurs="99"
}

type FareTypeGroupingInformationTypeI struct {
	// Account code
	// xmlType: AlphaNumericString_Length1To35
	PricingGroup string `xml:"pricingGroup,omitempty"`
}

type FormOfPaymentDetailsTypeI struct {
	// Form of payment type.
	// xmlType: AlphaNumericString_Length1To10
	Type string `xml:"type"`

	// xmlType: AlphaNumericString_Length1To3
	Indicator string `xml:"indicator,omitempty"`

	// Form of payment amount.
	// xmlType: NumericDecimal_Length1To18
	Amount *float64 `xml:"amount,omitempty"`

	// Vendor code (CC).
	// xmlType: AlphaNumericString_Length1To35
	VendorCode string `xml:"vendorCode,omitempty"`

	// Account number (CC/GR/SGR).
	// xmlType: AlphaNumericString_Length1To35
	CreditCardNumber string `xml:"creditCardNumber,omitempty"`

	// Expiration date (CC) (mmyy).
	// xmlType: AlphaNumericString_Length1To35
	ExpiryDate string `xml:"expiryDate,omitempty"`

	// Approval code (CC).
	// xmlType: AlphaNumericString_Length1To17
	ApprovalCode string `xml:"approvalCode,omitempty"`

	// Source of approval code (CC).
	// xmlType: AlphaNumericString_Length1To3
	SourceOfApproval string `xml:"sourceOfApproval,omitempty"`

	// Maximum authorized amount (CC).
	// xmlType: NumericDecimal_Length1To18
	AuthorisedAmount *float64 `xml:"authorisedAmount,omitempty"`

	// Address verification code (CC).
	// xmlType: AlphaNumericString_Length1To3
	AddressVerification string `xml:"addressVerification,omitempty"`

	// Customer file reference.
	// xmlType: AlphaNumericString_Length1To70
	CustomerAccount string `xml:"customerAccount,omitempty"`

	// Extended payment code (CC).
	// xmlType: AlphaNumericString_Length1To3
	ExtendedPayment string `xml:"extendedPayment,omitempty"`

	// Specifies the EMD document number that is being used for payment.
	// xmlType: AlphaNumericString_Length1To70
	FopFreeText string `xml:"fopFreeText,omitempty"`

	// Credit card corporate contract.
	// xmlType: AlphaNumericString_Length1To3
	MembershipStatus string `xml:"membershipStatus,omitempty"`

	// Credit card transaction information.
	// xmlType: AlphaNumericString_Length1To35
	TransactionInfo string `xml:"transactionInfo,omitempty"`
}

type FormOfPaymentTypeI struct {
	FormOfPayment []*FormOfPaymentDetailsTypeI `xml:"formOfPayment"` // maxOccurs="99"
}

type FreeTextQualificationTypeI struct {
	// Function qualifier.
	// xmlType: AlphaNumericString_Length1To3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// A code describing data in 4440.
	// xmlType: AlphaNumericString_Length1To4
	InformationType string `xml:"informationType,omitempty"`

	// Fare calculation reporting indicator or pricing indicator.
	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId,omitempty"`
}

type FrequentTravellerIdentificationTypeI struct {
	// Airline designator, coded.
	// xmlType: AlphaNumericString_Length1To35
	Carrier string `xml:"carrier"`

	// A code to identify a frequent traveller - the frequent traveller number.
	// xmlType: AlphaNumericString_Length1To25
	Number string `xml:"number"`
}

type FrequentTravellerInformationTypeI struct {
	FrequentTravellerDetails []*FrequentTravellerIdentificationTypeI `xml:"frequentTravellerDetails"` // maxOccurs="9"
}

type InteractiveFreeTextTypeI struct {
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`

	// Free text message.
	// xmlType: AlphaNumericString_Length1To70
	FreeText []string `xml:"freeText,omitempty"` // maxOccurs="99"
}

type InternalIDDetailsTypeI struct {
	// Reference number/authority code assigned to the requester as in the booking agent's initials or logon id
	// xmlType: AlphaNumericString_Length1To9
	InhouseId string `xml:"inhouseId"`

	// A code to identify type of agent identification.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`
}

type ItemNumberIdentificationTypeI struct {
	// Reference number or Fare component number
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"`
}

type ItemNumberTypeI struct {
	ItemNumberDetails []*ItemNumberIdentificationTypeI `xml:"itemNumberDetails"` // maxOccurs="99"
}

type LocationDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To25
	City string `xml:"city,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	Country string `xml:"country,omitempty"`
}

type LocationDetailsTypeI_83777C struct {
	// ISO Country code of where ticketed.
	// xmlType: AlphaNumericString_Length1To3
	Country string `xml:"country,omitempty"`
}

type LocationTypeI struct {
	// xmlType: AlphaNumericString_Length1To25
	TrueLocationId string `xml:"trueLocationId,omitempty"`

	// xmlType: AlphaNumericString_Length1To17
	TrueLocation string `xml:"trueLocation,omitempty"`
}

type LocationTypeI_83769C struct {
	// xmlType: AlphaNumericString_Length1To25
	TrueLocationId string `xml:"trueLocationId,omitempty"`
}

type MarriageControlDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	Relation string `xml:"relation,omitempty"`

	MarriageIdentifier *formats.NumericInteger_Length1To10 `xml:"marriageIdentifier,omitempty"`

	LineNumber *formats.NumericInteger_Length1To6 `xml:"lineNumber,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	OtherRelation string `xml:"otherRelation,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	CarrierCode string `xml:"carrierCode,omitempty"`
}

type MessageActionDetailsTypeI struct {
	MessageFunctionDetails *MessageFunctionBusinessDetailsTypeI `xml:"messageFunctionDetails"`

	// Indicates wether request was processed successfully.
	// xmlType: AlphaNumericString_Length1To3
	ResponseType string `xml:"responseType,omitempty"`
}

type MessageFunctionBusinessDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	MessageFunction string `xml:"messageFunction,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"` // maxOccurs="20"
}

type MonetaryInformationDetailsTypeI struct {
	// A code describing an amount.
	// xmlType: AlphaNumericString_Length1To3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount, 'FREE', 'BULK'
	// xmlType: AlphaNumericString_Length1To35
	Amount string `xml:"amount"`

	// ISO currency code.
	// xmlType: AlphaNumericString_Length1To3
	Currency string `xml:"currency,omitempty"`

	// The From city for the carrier fee
	// xmlType: AlphaNumericString_Length1To25
	LocationFrom string `xml:"locationFrom,omitempty"`

	// xmlType: AlphaNumericString_Length1To25
	LocationTo string `xml:"locationTo,omitempty"`
}

type MonetaryInformationTypeI struct {
	MonetaryDetails []*MonetaryInformationDetailsTypeI `xml:"monetaryDetails"` // maxOccurs="20"
}

type NumberOfUnitDetailsTypeI struct {
	// Total number of ticket/document numbers.
	NumberOfUnit formats.NumericInteger_Length1To15 `xml:"numberOfUnit"`

	// xmlType: AlphaNumericString_Length1To3
	UnitQualifier string `xml:"unitQualifier,omitempty"`
}

type NumberOfUnitsTypeI struct {
	QuantityDetails []*NumberOfUnitDetailsTypeI `xml:"quantityDetails"` // maxOccurs="9"
}

type OriginAndDestinationDetailsTypeI struct {
	// Origin.
	// xmlType: AlphaNumericString_Length1To25
	Origin string `xml:"origin,omitempty"`

	// Destination
	// xmlType: AlphaNumericString_Length1To25
	Destination string `xml:"destination,omitempty"`
}

type OriginatorDetailsTypeI struct {
	// ISO Country code of the agent issuing the ticket.
	// xmlType: AlphaNumericString_Length1To3
	CodedCountry string `xml:"codedCountry,omitempty"`

	// ISO currency code for currency of originator country.
	// xmlType: AlphaNumericString_Length1To3
	CodedCurrency string `xml:"codedCurrency,omitempty"`

	// ISO code of language.
	// xmlType: AlphaNumericString_Length1To3
	CodedLanguage string `xml:"codedLanguage,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// ATA/IATA ID number or pseudo IATA number.
	// xmlType: AlphaNumericString_Length1To9
	OriginatorId string `xml:"originatorId"`

	// Amadeus office identification (AMID)
	// xmlType: AlphaNumericString_Length1To9
	InHouseIdentification1 string `xml:"inHouseIdentification1,omitempty"`

	// xmlType: AlphaNumericString_Length1To9
	InHouseIdentification2 string `xml:"inHouseIdentification2,omitempty"`
}

type OriginatorIdentificationDetailsTypeI_83809C struct {
	// Original issue agent numeric code (IATA number)
	OriginatorId *formats.NumericInteger_Length1To9 `xml:"originatorId,omitempty"`
}

type OriginatorOfRequestDetailsTypeI struct {
	DeliveringSystem *SystemDetailsTypeI `xml:"deliveringSystem"`

	OriginIdentification *OriginatorIdentificationDetailsTypeI `xml:"originIdentification,omitempty"`

	LocationDetails *LocationTypeI_83769C `xml:"locationDetails,omitempty"`

	CascadingSystem *SystemDetailsTypeI_83771C `xml:"cascadingSystem,omitempty"`

	// 1 character code for airline agent, travel agent, etc.
	// xmlType: AlphaNumericString_Length1To1
	OriginatorTypeCode string `xml:"originatorTypeCode,omitempty"`

	OriginDetails *OriginatorDetailsTypeI `xml:"originDetails,omitempty"`

	// A code identifying the issuing agent or if the transaction is being initiated by a computer programme then the word 'system'.
	// xmlType: AlphaNumericString_Length1To9
	Originator string `xml:"originator"`
}

type PricingTicketingDetailsTypeI struct {
	PriceTicketDetails *PricingTicketingInformationTypeI `xml:"priceTicketDetails,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	PriceTariffType string `xml:"priceTariffType,omitempty"`

	ProductDateTimeDetails *ProductDateTimeTypeI_52035C `xml:"productDateTimeDetails,omitempty"`

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`

	CompanyNumberDetails *CompanyIdentificationNumbersTypeI `xml:"companyNumberDetails,omitempty"`

	LocationDetails *LocationDetailsTypeI `xml:"locationDetails,omitempty"`

	OtherLocationDetails *LocationDetailsTypeI `xml:"otherLocationDetails,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	IdNumber string `xml:"idNumber,omitempty"`

	// xmlType: NumericDecimal_Length1To18
	MonetaryAmount *float64 `xml:"monetaryAmount,omitempty"`
}

type PricingTicketingDetailsTypeI_51227S struct {
	PriceTicketDetails *PricingTicketingInformationTypeI_83775C `xml:"priceTicketDetails,omitempty"`

	// Fare rule code
	// xmlType: AlphaNumericString_Length1To3
	PriceTariffType string `xml:"priceTariffType,omitempty"`

	ProductDateTimeDetails *ProductDateTimeTypeI_83774C `xml:"productDateTimeDetails,omitempty"`

	CompanyDetails *CompanyIdentificationTypeI_83773C `xml:"companyDetails,omitempty"`

	LocationDetails []*LocationDetailsTypeI_83777C `xml:"locationDetails,omitempty"` // maxOccurs="2"

	// Waiver code
	// xmlType: AlphaNumericString_Length1To35
	IdNumber string `xml:"idNumber,omitempty"`
}

type PricingTicketingInformationTypeI struct {
	// Ticketing Mode Indicator
	// xmlType: AlphaNumericString_Length1To3
	TicketingModeIndicator string `xml:"ticketingModeIndicator,omitempty"`

	// International or Domestic Sales Indicator
	// xmlType: AlphaNumericString_Length1To3
	InternationalDomSalesIndicator string `xml:"internationalDomSalesIndicator,omitempty"`

	// Statistical Code
	// xmlType: AlphaNumericString_Length1To3
	StatisticalCode string `xml:"statisticalCode,omitempty"`

	// Self Sale Indicator
	// xmlType: AlphaNumericString_Length1To3
	SelfSaleIndicator string `xml:"selfSaleIndicator,omitempty"`

	// Net Reporting Indicator
	// xmlType: AlphaNumericString_Length1To3
	NetReportingIndicator string `xml:"netReportingIndicator,omitempty"`

	// Tax on Commission
	// xmlType: AlphaNumericString_Length1To3
	TaxOnCommissionIndicator string `xml:"taxOnCommissionIndicator,omitempty"`

	// Non Endorsable indicator
	// xmlType: AlphaNumericString_Length1To3
	NonEndorsableIndicator string `xml:"nonEndorsableIndicator,omitempty"`

	// Non Refundable indicator
	// xmlType: AlphaNumericString_Length1To3
	NonRefundableIndicator string `xml:"nonRefundableIndicator,omitempty"`

	// Penalty Restriction
	// xmlType: AlphaNumericString_Length1To3
	PenaltyRestrictionIndicator string `xml:"penaltyRestrictionIndicator,omitempty"`

	// Present Credit Card indicator
	// xmlType: AlphaNumericString_Length1To3
	PresentCreditCardIndicator string `xml:"presentCreditCardIndicator,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	EmergencySet string `xml:"emergencySet,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	EmergencyClear string `xml:"emergencyClear,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	NonInterlineableIndicator string `xml:"nonInterlineableIndicator,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	NonCommissionable string `xml:"nonCommissionable,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	PresentDebitCardIndicator string `xml:"presentDebitCardIndicator,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	NonReissuableIndicator string `xml:"nonReissuableIndicator,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	CarrierFeeReportingIndicator string `xml:"carrierFeeReportingIndicator,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	RefundSystemComputerCalculated string `xml:"refundSystemComputerCalculated,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	RefundManuallyCalculated string `xml:"refundManuallyCalculated,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	Indicator20 []string `xml:"indicator20,omitempty"` // maxOccurs="11"
}

type PricingTicketingInformationTypeI_83775C struct {
	// Ticketing mode indicator
	// xmlType: AlphaNumericString_Length1To3
	TicketingModeIndicator string `xml:"ticketingModeIndicator,omitempty"`

	// International or domestic sales indicator.
	// xmlType: AlphaNumericString_Length1To3
	InternationalDomSalesIndicator string `xml:"internationalDomSalesIndicator,omitempty"`

	// statistical Code
	// xmlType: AlphaNumericString_Length1To3
	StatisticalCodeIndicator string `xml:"statisticalCodeIndicator,omitempty"`

	// Self Sale Indicator
	// xmlType: AlphaNumericString_Length1To3
	SelfSaleIndicator string `xml:"selfSaleIndicator,omitempty"`

	// Net Reporting Indicator
	// xmlType: AlphaNumericString_Length1To3
	NetReportingIndicator string `xml:"netReportingIndicator,omitempty"`

	// Tax on commission indicator
	// xmlType: AlphaNumericString_Length1To3
	TaxOnCommissionIndicator string `xml:"taxOnCommissionIndicator,omitempty"`

	// Non-Endorsable indicator
	// xmlType: AlphaNumericString_Length1To3
	NonEndorsableIndicator string `xml:"nonEndorsableIndicator,omitempty"`

	// Non-Refundable indicator
	// xmlType: AlphaNumericString_Length1To3
	NonRefundableIndicator string `xml:"nonRefundableIndicator,omitempty"`

	// Penalty restriction indicator.
	// xmlType: AlphaNumericString_Length1To3
	PenaltyRestrictionIndicator string `xml:"penaltyRestrictionIndicator,omitempty"`

	// Present Credit Card indicator
	// xmlType: AlphaNumericString_Length1To3
	PresentCreditCardIndicator string `xml:"presentCreditCardIndicator,omitempty"`

	// Non-interlineable indicator.
	// xmlType: AlphaNumericString_Length1To3
	NonInterlineableIndicator string `xml:"nonInterlineableIndicator,omitempty"`

	// Non-commissionable indicator.
	// xmlType: AlphaNumericString_Length1To3
	NonCommissionableIndicator string `xml:"nonCommissionableIndicator,omitempty"`

	// Non exchangeable indicator
	// xmlType: AlphaNumericString_Length1To3
	NonExchangeableIndicator string `xml:"nonExchangeableIndicator,omitempty"`
}

type PricingTicketingSubsequentTypeI struct {
	// xmlType: AlphaNumericString_Length1To35
	ItemNumber string `xml:"itemNumber,omitempty"`

	FareBasisDetails *RateTariffClassInformationTypeI `xml:"fareBasisDetails,omitempty"`

	// xmlType: NumericDecimal_Length1To18
	FareValue *float64 `xml:"fareValue,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	PriceType string `xml:"priceType,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	SpecialCondition string `xml:"specialCondition,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	OtherSpecialCondition string `xml:"otherSpecialCondition,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	AdditionalSpecialCondition string `xml:"additionalSpecialCondition,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	TaxCategory []string `xml:"taxCategory,omitempty"` // maxOccurs="2"
}

type PricingTicketingSubsequentTypeI_51235S struct {
	FareBasisDetails *RateTariffClassInformationTypeI_83791C `xml:"fareBasisDetails,omitempty"`

	// Reason for issuance code.
	// xmlType: AlphaNumericString_Length1To3
	SpecialCondition string `xml:"specialCondition,omitempty"`

	// Reason for issuance sub code.
	// xmlType: AlphaNumericString_Length1To3
	OtherSpecialCondition string `xml:"otherSpecialCondition,omitempty"`
}

type ProductDateTimeTypeI struct {
	// xmlType: AlphaNumericString_Length1To35
	DepartureDate string `xml:"departureDate,omitempty"`

	// xmlType: AlphaNumericString_Length1To4
	DepartureTime string `xml:"departureTime,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	ArrivalDate string `xml:"arrivalDate,omitempty"`

	// xmlType: AlphaNumericString_Length1To4
	ArrivalTime string `xml:"arrivalTime,omitempty"`

	DateVariation *formats.NumericInteger_Length1To1 `xml:"dateVariation,omitempty"`
}

type ProductDateTimeTypeI_52035C struct {
	// xmlType: AlphaNumericString_Length1To35
	DepartureDate string `xml:"departureDate,omitempty"`

	DepartureTime formats.Time24_HHMM `xml:"departureTime,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	ArrivalDate string `xml:"arrivalDate,omitempty"`

	ArrivalTime *formats.NumericInteger_Length1To4 `xml:"arrivalTime,omitempty"`

	DateVariation *formats.NumericInteger_Length1To1 `xml:"dateVariation,omitempty"`
}

type ProductDateTimeTypeI_83774C struct {
	// Date of issue (ddmmyy).
	// xmlType: AlphaNumericString_Length1To35
	DepartureDate string `xml:"departureDate,omitempty"`

	// xmlType: AlphaNumericString_Length1To4
	DepartureTime string `xml:"departureTime,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To35
	FlightNumber string `xml:"flightNumber"`

	// xmlType: AlphaNumericString_Length1To17
	BookingClass string `xml:"bookingClass,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`

	// xmlType: AlphaNumericString_Length1To7
	Modifier []string `xml:"modifier,omitempty"` // maxOccurs="3"
}

type ProductLocationDetailsTypeI struct {
	// Enroute city/airport
	// xmlType: AlphaNumericString_Length1To25
	Station string `xml:"station,omitempty"`
}

type ProductTypeDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To6
	FlightIndicator []string `xml:"flightIndicator"` // maxOccurs="9"
}

type RateTariffClassInformationTypeI struct {
	// xmlType: AlphaNumericString_Length1To35
	RateTariffClass string `xml:"rateTariffClass,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	RateTariffIndicator string `xml:"rateTariffIndicator,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	OtherRateTariffClass string `xml:"otherRateTariffClass,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	OtherRateTariffIndicator string `xml:"otherRateTariffIndicator,omitempty"`
}

type RateTariffClassInformationTypeI_83791C struct {
	// Fare basis/ticket designator.
	// xmlType: AlphaNumericString_Length1To35
	RateTariffClass string `xml:"rateTariffClass,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	RateTariffIndicator string `xml:"rateTariffIndicator,omitempty"`
}

type ReferenceTypeI struct {
	// The coupon sequence number (1of3, 2of5)
	// xmlType: AlphaNumericString_Length1To6
	NumberOfItems string `xml:"numberOfItems,omitempty"`

	// The coupon/itinerary sequence number.
	// xmlType: AlphaNumericString_Length1To35
	LastItemIdentifier []string `xml:"lastItemIdentifier,omitempty"` // maxOccurs="99"
}

type RelatedProductInformationTypeI struct {
	// Sold reservation status code
	// xmlType: AlphaNumericString_Length1To3
	StatusCode []string `xml:"statusCode,omitempty"` // maxOccurs="10"
}

type ReservationControlInformationDetailsTypeI struct {
	// A 2-3 character airline/CRS code of the following record reference.
	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId,omitempty"`

	// A reference to a record.
	// xmlType: AlphaNumericString_Length1To20
	ControlNumber string `xml:"controlNumber,omitempty"`

	// A code to identify type of record reference.
	// xmlType: AlphaNumericString_Length1To1
	ControlType string `xml:"controlType,omitempty"`
}

type ReservationControlInformationTypeI struct {
	Reservation []*ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"` // maxOccurs="9"
}

type RoutingInformationTypeI struct {
	RoutingDetails []*ProductLocationDetailsTypeI `xml:"routingDetails,omitempty"` // maxOccurs="99"
}

type SelectionDetailsInformationTypeI struct {
	// To specify the type of carrier fee.
	// xmlType: AlphaNumericString_Length1To3
	Option string `xml:"option"`
}

type SelectionDetailsTypeI struct {
	SelectionDetails []*SelectionDetailsInformationTypeI `xml:"selectionDetails,omitempty"` // maxOccurs="99"
}

type SourceTypeDetailsTypeI struct {
	// To specify this is the original issuer of the ticket.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier1 []string `xml:"sourceQualifier1"` // maxOccurs="3"
}

type SpecificDataInformationTypeI struct {
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation"`

	DataInformation []*DataInformationTypeI `xml:"dataInformation,omitempty"` // maxOccurs="99"
}

type StatusDetailsTypeI struct {
	// To specify that this is reissued flown flight coupon data.
	// xmlType: AlphaNumericString_Length1To3
	Indicator string `xml:"indicator,omitempty"`
}

type StatusTypeI struct {
	StatusDetails []*StatusDetailsTypeI `xml:"statusDetails"` // maxOccurs="99"
}

type SystemDetailsTypeI struct {
	// A 2-3 character airline/CRS code, or bilaterally agreed code, of the system that delivers the message.
	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId"`

	// 3 character ATA/IATA airport/city code of the delivering system physical location.
	// xmlType: AlphaNumericString_Length1To25
	LocationId string `xml:"locationId,omitempty"`
}

type SystemDetailsTypeI_83771C struct {
	// A 2-3 character airline/CRS code, or bilaterally agree code, of the system originating the message, WHEN DIFFERENT FROM THE DELIVERING SYSTEM.
	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId,omitempty"`

	// 3 character ATA/IATA airport/city code of the system that originates the message.
	// xmlType: AlphaNumericString_Length1To25
	LocationId string `xml:"locationId,omitempty"`
}

type TaxDetailsTypeI struct {
	// Tax/fee/charge amount.
	// xmlType: AlphaNumericString_Length1To17
	Rate string `xml:"rate,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	CountryCode string `xml:"countryCode,omitempty"`

	// ISO code identifying currency.
	// xmlType: AlphaNumericString_Length1To3
	CurrencyCode string `xml:"currencyCode,omitempty"`

	// Tax/fee/charge type.
	// xmlType: AlphaNumericString_Length1To3
	Type []string `xml:"type,omitempty"` // maxOccurs="99"
}

type TaxTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	TaxCategory string `xml:"taxCategory,omitempty"`

	TaxDetails []*TaxDetailsTypeI `xml:"taxDetails"` // maxOccurs="99"
}

type TicketAgentInfoTypeI struct {
	// Service airline/system provider identifier e.g. the airline passenger accounting and modulus 7
	// xmlType: AlphaNumericString_Length1To15
	CompanyIdNumber string `xml:"companyIdNumber,omitempty"`

	InternalIdDetails []*InternalIDDetailsTypeI `xml:"internalIdDetails,omitempty"` // maxOccurs="5"

	// Booking agent IATA number
	BookingIataNumber *formats.NumericInteger_Length1To9 `xml:"bookingIataNumber,omitempty"`
}

type TicketNumberDetailsTypeI struct {
	// Document number.
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"`

	// Document type.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`

	// Total number of booklets issued.
	NumberOfBooklets *formats.NumericInteger_Length1To15 `xml:"numberOfBooklets,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	DataIndicator string `xml:"dataIndicator,omitempty"`

	// In connection with document number
	// xmlType: AlphaNumericString_Length1To35
	InConnectionWith string `xml:"inConnectionWith,omitempty"`
}

type TicketNumberTypeI struct {
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`

	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"`
}

type TourDetailsTypeI struct {
	// Tour code.
	// xmlType: AlphaNumericString_Length1To35
	TourCode string `xml:"tourCode,omitempty"`
}

type TourInformationTypeI struct {
	TourInformationDetails *TourDetailsTypeI `xml:"tourInformationDetails,omitempty"`
}

type TravelProductInformationTypeI struct {
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`

	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"`

	FlightTypeDetails *ProductTypeDetailsTypeI `xml:"flightTypeDetails,omitempty"`

	ItemNumber *formats.NumericInteger_Length1To6 `xml:"itemNumber,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	SpecialSegment string `xml:"specialSegment,omitempty"`

	MarriageDetails []*MarriageControlDetailsTypeI `xml:"marriageDetails,omitempty"` // maxOccurs="99"
}

type TravellerDetailsTypeI struct {
	// Specifies passenger given name and title.
	// xmlType: AlphaNumericString_Length1To70
	GivenName string `xml:"givenName"`
}

type TravellerInformationTypeI struct {
	PaxDetails *TravellerSurnameInformationTypeI `xml:"paxDetails"`

	OtherPaxDetails []*TravellerDetailsTypeI `xml:"otherPaxDetails,omitempty"` // maxOccurs="99"
}

type TravellerNameDetailsType struct {
	// xmlType: AlphaNumericString_Length1To5
	NameType string `xml:"nameType,omitempty"`

	// xmlType: AlphaNumericString_Length1To1
	ReferenceName string `xml:"referenceName,omitempty"`

	// xmlType: AlphaNumericString_Length1To1
	DisplayedName string `xml:"displayedName,omitempty"`

	// xmlType: AlphaNumericString_Length1To4
	RomanizationMethod string `xml:"romanizationMethod,omitempty"`

	// Passenger surname
	// xmlType: AlphaNumericString_Length1To70
	Surname string `xml:"surname"`

	// Passenger firstname
	// xmlType: AlphaNumericString_Length1To70
	GivenName string `xml:"givenName,omitempty"`

	// xmlType: AlphaNumericString_Length1To70
	Title []string `xml:"title,omitempty"` // maxOccurs="2"
}

type TravellerNameInfoType struct {
	// PAX = PAX IN = Infant
	Qualifier formats.AMA_EDICodesetType_Length1to3 `xml:"qualifier,omitempty"`

	Quantity *formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`

	// Passenger type (PTC).
	Type formats.AMA_EDICodesetType_Length1to3 `xml:"type,omitempty"`

	// Passenger type (PTC).
	OtherType formats.AMA_EDICodesetType_Length1to3 `xml:"otherType,omitempty"`

	// xmlType: AlphaNumericString_Length1To1
	InfantIndicator string `xml:"infantIndicator,omitempty"`

	// Identification code, 2 cases: ID<1 to 51 char free text) or CR<1 to 40 char free text).
	// xmlType: AlphaNumericString_Length1To70
	TravellerIdentificationCode string `xml:"travellerIdentificationCode,omitempty"`
}

type TravellerNameInfoType_276832C struct {
	// PAX = PAX IN = Infant
	Qualifier formats.AMA_EDICodesetType_Length1to3 `xml:"qualifier,omitempty"`

	Quantity *formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`

	// Passenger type (PTC).
	Type formats.AMA_EDICodesetType_Length1to3 `xml:"type,omitempty"`

	// Passenger type (PTC).
	OtherType formats.AMA_EDICodesetType_Length1to3 `xml:"otherType,omitempty"`

	// xmlType: AlphaNumericString_Length1To1
	InfantIndicator string `xml:"infantIndicator,omitempty"`

	// Identification code, 2 cases: ID<1 to 51 char free text) or CR<1 to 40 char free text).
	// xmlType: AlphaNumericString_Length1To70
	TravellerIdentificationCode string `xml:"travellerIdentificationCode,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	Gender string `xml:"gender,omitempty"`

	Age *formats.NumericInteger_Length1To3 `xml:"age,omitempty"`
}

type TravellerPriorityDetailsTypeI struct {
	// Staff airline of employment
	// xmlType: AlphaNumericString_Length1To35
	Company string `xml:"company,omitempty"`

	// Staff date of joining (ddmmmyy)
	// xmlType: AlphaNumericString_Length1To35
	DateOfJoining string `xml:"dateOfJoining,omitempty"`

	// Staff id number
	// xmlType: AlphaNumericString_Length1To10
	TravellerReference string `xml:"travellerReference,omitempty"`
}

type TravellerSurnameInformationTypeI struct {
	// Passenger name.
	// xmlType: AlphaNumericString_Length1To70
	Surname string `xml:"surname"`

	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`

	// Specify age of unaccompanied minor.
	Quantity *formats.NumericInteger_Length1To15 `xml:"quantity,omitempty"`
}
