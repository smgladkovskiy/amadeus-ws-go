package Ticket_ProcessEDocResponse_v15_2 // tatres152

import "github.com/tmconsulting/amadeus-golang-sdk/structs/formats"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TATRES_15_2_1A Ticket_ProcessEDocReply"`

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

	// Tour code at Level 0
	TourInfo *TourInformationTypeI `xml:"tourInfo,omitempty"`

	// Form of payment at Level 0
	Fop *FormOfPaymentTypeI `xml:"fop,omitempty"`

	// Document count
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits,omitempty"`

	// tax details at Level 0
	TaxInfo []*TaxTypeI `xml:"taxInfo,omitempty"` // maxOccurs="6"

	// Error at message level
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"`

	// Coupon details
	CouponInfo *CouponInformationTypeI `xml:"couponInfo,omitempty"`

	// Document details
	DocumentInformation []*DocumentInformationDetailsTypeI `xml:"documentInformation,omitempty"` // maxOccurs="99"

	// Freeflow text at Level 0 POD Payment On Demand indicator
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // maxOccurs="99"

	// Fare details
	FareDetails *FareInformationTypeI `xml:"fareDetails,omitempty"`

	// Document group at level 1
	GeneralDocGroup []*GeneralDocGroup `xml:"generalDocGroup,omitempty"` // maxOccurs="999"

	// Passenger Group
	DocGroup []*DocGroup `xml:"docGroup,omitempty"` // maxOccurs="99"
}

type GeneralDocGroup struct {
	// Document details of group 1
	DocInfo *TicketNumberTypeI `xml:"docInfo"`

	// Originator details group 1
	OriginatorInfo *OriginatorOfRequestDetailsTypeI `xml:"originatorInfo,omitempty"`

	// error at TKT level
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"`

	// Coupon details of group 1
	CouponInfo []*CouponInformationTypeI `xml:"couponInfo,omitempty"` // maxOccurs="4"
}

type DocGroup struct {
	// Passenger Details
	PaxInfo *TravellerInformationTypeI `xml:"paxInfo"`

	EnhancedPaxInfo *EnhancedTravellerInformationType `xml:"enhancedPaxInfo,omitempty"`

	// Travel agent details at level 2
	SysProvider *TicketAgentInfoTypeI `xml:"sysProvider,omitempty"`

	// Reservation details at level 2
	ReferenceInfo *ReservationControlInformationTypeI `xml:"referenceInfo,omitempty"`

	// Monetary information at level 1
	FareInfo *MonetaryInformationTypeI `xml:"fareInfo,omitempty"`

	// Form of payment at level 2
	Fop *FormOfPaymentTypeI `xml:"fop,omitempty"`

	// Pricing details at level 1
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"`

	// Origin - Destination at Level 2
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"`

	// Frequent traveller info at Level 2
	FrequentTravellerInfo *FrequentTravellerInformationTypeI `xml:"frequentTravellerInfo,omitempty"`

	// Tour code L2
	TourInfo *TourInformationTypeI `xml:"tourInfo,omitempty"`

	// Originator details group 2
	OriginatorInfo *OriginatorOfRequestDetailsTypeI `xml:"originatorInfo,omitempty"`

	// Coupon count
	NumberOfUnits *NumberOfUnitsTypeI `xml:"numberOfUnits,omitempty"`

	// tax details at Level 2
	TaxInfo []*TaxTypeI `xml:"taxInfo,omitempty"` // maxOccurs="5"

	// Error at Pax level
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"`

	// Document details at Level 2
	DocumentInformation []*DocumentInformationDetailsTypeI `xml:"documentInformation,omitempty"` // maxOccurs="99"

	// Freeflow text at Level 1 RA Revenue Attribution indicator
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // maxOccurs="99"

	OtherTextInfo *FreeTextInformationType `xml:"otherTextInfo,omitempty"`

	// form of identification at Level 2
	CustomerReference *ConsumerReferenceInformationTypeI `xml:"customerReference,omitempty"`

	// Fare details
	FareDetails *FareInformationTypeI `xml:"fareDetails,omitempty"`

	// contains the RFIC/RFISC
	PricingInfo *PricingTicketingSubsequentTypeI `xml:"pricingInfo,omitempty"`

	// Document Group
	DocDetailsGroup []*DocDetailsGroup `xml:"docDetailsGroup,omitempty"` // maxOccurs="99"

	// Structured fare calc group
	FareElementsGroup []*FareElementsGroup `xml:"fareElementsGroup,omitempty"` // maxOccurs="2"

	// Reissued Flown Group
	ReissuedFlownDetails []*ReissuedFlownDetails `xml:"reissuedFlownDetails,omitempty"` // maxOccurs="99"

	// Staff group
	StaffTravellerGroup *StaffTravellerGroup `xml:"staffTravellerGroup,omitempty"`

	// Exchanged information Group
	OriginalIssuanceGroup []*OriginalIssuanceGroup `xml:"originalIssuanceGroup,omitempty"` // maxOccurs="2"

	// Fee group
	CarrierFeeGroup []*CarrierFeeGroup `xml:"carrierFeeGroup,omitempty"` // maxOccurs="9"
}

type DocDetailsGroup struct {
	// Document details of group 2 Conjunction Box indicator
	DocInfo *TicketNumberTypeI `xml:"docInfo"`

	// error at level 2
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"`

	// Originator details group 3
	OriginatorInfo *OriginatorOfRequestDetailsTypeI `xml:"originatorInfo,omitempty"`

	// Date information at level 3
	ValidityDates *DateAndTimeInformationTypeI `xml:"validityDates,omitempty"`

	// Freeflow text at Level 2
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // maxOccurs="2"

	// Pricing details at level 2
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"`

	// form of identification at Level 3
	CustomerReference *ConsumerReferenceInformationTypeI `xml:"customerReference,omitempty"`

	// Coupon Group
	CouponGroup []*CouponGroup `xml:"couponGroup,omitempty"` // maxOccurs="99"

	FareComponentDetailsGroup []*FareComponentDetailsGroup `xml:"fareComponentDetailsGroup,omitempty"` // maxOccurs="4"
}

type CouponGroup struct {
	// Coupon details of group 4
	CouponInfo *CouponInformationTypeI `xml:"couponInfo"`

	// Travel details at Level 4
	Leg []*TravelProductInformationTypeI_29340S `xml:"leg,omitempty"` // maxOccurs="2"

	// Reservation details at level 4
	ReferenceInfo *ReservationControlInformationTypeI `xml:"referenceInfo,omitempty"`

	// Related product information at level 4
	BookingStatus *RelatedProductInformationTypeI `xml:"bookingStatus,omitempty"`

	// Pricing subsequent at L4
	PricingInfo *PricingTicketingSubsequentTypeI `xml:"pricingInfo,omitempty"`

	BaggageInfo *ExcessBaggageTypeI `xml:"baggageInfo,omitempty"`

	// Frequent traveller info at Level 4
	FrequentTravellerInfo *FrequentTravellerInformationTypeI `xml:"frequentTravellerInfo,omitempty"`

	// Date information at level 4
	ValidityDates *DateAndTimeInformationTypeI `xml:"validityDates,omitempty"`

	// Error at coupon level
	Error *ApplicationErrorInformationTypeI `xml:"error,omitempty"`

	// Originator details group 3
	OriginatorInfo *OriginatorOfRequestDetailsTypeI `xml:"originatorInfo,omitempty"`

	// Freeflow text at Level 3
	TextInfo []*InteractiveFreeTextTypeI `xml:"textInfo,omitempty"` // maxOccurs="99"

	// Pricing details at level 3 Noshow indicator
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"`

	FareBasisDetails *FareQualifierDetailsType_174783S `xml:"fareBasisDetails,omitempty"`
}

type FareComponentDetailsGroup struct {
	// fare Component identification
	FareComponentID *ItemNumberType `xml:"fareComponentID"`

	// Market information related to fare component
	MarketFareComponent *TravelProductInformationTypeI `xml:"marketFareComponent,omitempty"`

	// Monetary Information
	MonetaryInformation *MonetaryInformationType `xml:"monetaryInformation,omitempty"`

	// Component Class information
	ComponentClassInfo *PricingOrTicketingSubsequentType `xml:"componentClassInfo,omitempty"`

	// Fare Qualifier Detail
	FareQualifiersDetail *FareQualifierDetailsType `xml:"fareQualifiersDetail,omitempty"`

	// Details of the fare family used for this fare component
	FareFamilyDetails *FareFamilyType `xml:"fareFamilyDetails,omitempty"`

	// Carrier owner of the fare family
	FareFamilyOwner *TransportIdentifierType `xml:"fareFamilyOwner,omitempty"`

	// The segments included in Coupon Fare information group (Tax/Fee/Charge Data at Ticket Coupon Level) have been developed for Amadeus Revenue Accounting User only as not mandatory in Reso 722f
	CouponDetailsGroup []*CouponDetailsGroup `xml:"couponDetailsGroup"` // maxOccurs="99"
}

type CouponDetailsGroup struct {
	// Tattoo + type of the product identifying the coupon.
	ProductId *ReferenceInfoType `xml:"productId"`

	// Flight Connection Type
	FlightConnectionType *TravelProductInformationType `xml:"flightConnectionType,omitempty"`

	// Tax Details Airport Tax included
	TaxDetails *TaxType `xml:"taxDetails,omitempty"`
}

type FareElementsGroup struct {
	// Fare component information
	FareComponentInfo *FareComponentInformationTypeI `xml:"fareComponentInfo"`

	// structured Fare Calc Group6
	FareComponentsGroup []*FareComponentsGroup `xml:"fareComponentsGroup,omitempty"` // maxOccurs="99"

	// structured Fare Calc Monetary Information 1
	StructuredFareCalcMonInfo1 *MonetaryInformationTypeI `xml:"structuredFareCalcMonInfo1,omitempty"`

	// Structured Fare Calc taxe information
	StructuredFareCalcTaxInfo []*TaxTypeI `xml:"structuredFareCalcTaxInfo,omitempty"` // maxOccurs="99"

	// structured Fare Calc Conversion Info
	StructuredFareCalcConversionInfo *ConversionRateTypeI `xml:"structuredFareCalcConversionInfo,omitempty"`
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
	StructuredFareCalcMonInfo2 *MonetaryInformationTypeI `xml:"structuredFareCalcMonInfo2,omitempty"`

	// structured Fare Calc Pricing subsequent
	StructuredFareCalcPricingInfo *PricingTicketingSubsequentTypeI `xml:"structuredFareCalcPricingInfo,omitempty"`

	// structured Fare Calculation details
	StructuredFareCalcDetails *FareCalculationCodeDetailsTypeI `xml:"structuredFareCalcDetails,omitempty"`

	// structured Fare Calc rate
	StructuredFareCalcRate *FareRulesInformationTypeI `xml:"structuredFareCalcRate,omitempty"`

	// Pricing details
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo"`

	// Fare details
	StructuredFareCalcFareDetails *FareInformationTypeI `xml:"structuredFareCalcFareDetails,omitempty"`
}

type FareCouponGroup struct {
	// structured Fare Calc Action
	StructuredFareCalcAction *ActionDetailsTypeI `xml:"structuredFareCalcAction"`

	StructuredFareCalcLeg *TravelProductInformationTypeI_29340S `xml:"structuredFareCalcLeg,omitempty"`
}

type ReissuedFlownDetails struct {
	// Refund status
	RefundStatus *StatusTypeI `xml:"refundStatus"`

	// Refund traveller details
	RefundLeg []*TravelProductInformationTypeI_29340S `xml:"refundLeg,omitempty"` // maxOccurs="2"

	// Pricing details
	ProductInfo *PricingTicketingDetailsTypeI `xml:"productInfo,omitempty"`

	// Refund document details
	RefundDocumentInfo *TicketNumberTypeI `xml:"refundDocumentInfo,omitempty"`

	// Refund coupon details
	CouponInfo []*CouponInformationTypeI `xml:"couponInfo,omitempty"` // maxOccurs="9"

	// Refund pricing subsequent
	RefundPricingInfo *PricingTicketingSubsequentTypeI `xml:"refundPricingInfo,omitempty"`

	// Refund routing details
	RefundRoutingDetails []*RoutingInformationTypeI `xml:"refundRoutingDetails,omitempty"` // maxOccurs="99"

	// Refund product details
	RefundProductDetails *AdditionalProductDetailsTypeI `xml:"refundProductDetails,omitempty"`
}

type StaffTravellerGroup struct {
	// To indicate the details associated with a traveller's status
	TravellerStatus *TravellerPriorityDetailsTypeI `xml:"travellerStatus"`

	// Staff information
	StaffInfo *SpecificDataInformationTypeI `xml:"staffInfo,omitempty"`
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

	LocationDetails *LocationTypeI `xml:"locationDetails,omitempty"`
}

type AdditionalFareQualifierDetailsType struct {
	// xmlType: AlphaNumericString_Length1To35
	RateClass string `xml:"rateClass,omitempty"`

	// xmlType: AlphaNumericString_Length1To18
	CommodityCategory string `xml:"commodityCategory,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	PricingGroup string `xml:"pricingGroup,omitempty"`
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
	NumberOfStops *int32 `xml:"numberOfStops,omitempty"`
}

type ApplicationErrorDetailTypeI struct {
	// xmlType: AlphaNumericString_Length1To5
	ErrorCode string `xml:"errorCode"`

	// xmlType: AlphaNumericString_Length1To3
	ErrorCategory string `xml:"errorCategory,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationTypeI struct {
	ErrorDetails *ApplicationErrorDetailTypeI `xml:"errorDetails"`
}

type BaggageDetailsTypeI struct {
	// Total number of units.
	FreeAllowance *int32 `xml:"freeAllowance,omitempty"`

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
	// Carrier owner fo the fare family
	// xmlType: AlphaNumericString_Length1To35
	OtherCompany string `xml:"otherCompany,omitempty"`
}

type CompanyIdentificationTypeI_51997C struct {
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

type DiscountPenaltyInformationType struct {
	FareQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"fareQualifier,omitempty"`
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
	TravellerNameInfo *TravellerNameInfoType `xml:"travellerNameInfo,omitempty"`

	// 5 possible types of names, for 1 passenger.
	OtherPaxNamesDetails []*TravellerNameDetailsType `xml:"otherPaxNamesDetails"` // maxOccurs="5"
}

type ExcessBaggageDetailsTypeI struct {
	// Currency of the excess baggage charge.
	// xmlType: AlphaNumericString_Length1To3
	Currency string `xml:"currency,omitempty"`

	// Excess baggage rate per unit.
	// xmlType: NumericDecimal_Length1To18
	Amount *float64 `xml:"amount,omitempty"`
}

type ExcessBaggageTypeI struct {
	ExcessBaggageDetails *ExcessBaggageDetailsTypeI `xml:"excessBaggageDetails,omitempty"`

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
	Count *int32 `xml:"count,omitempty"`

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

type FareFamilyDetailsType struct {
	// Commercial fare Family Short name
	// xmlType: AlphaNumericString_Length1To30
	CommercialFamily string `xml:"commercialFamily"`
}

type FareFamilyType struct {
	// Fare Family Short Name
	// xmlType: AlphaNumericString_Length1To30
	FareFamilyname string `xml:"fareFamilyname,omitempty"`

	// HIERARCHICAL ORDER WITHIN FARE FAMILY
	Hierarchy *int32 `xml:"hierarchy,omitempty"`

	// Indicates Commercial Fare Family Short names
	CommercialFamilyDetails []*FareFamilyDetailsType `xml:"commercialFamilyDetails,omitempty"` // maxOccurs="20"
}

type FareInformationTypeI struct {
	// Specifies an industry defined fare component priced passenger type code (PTC).
	// xmlType: AlphaNumericString_Length1To3
	ValueQualifier string `xml:"valueQualifier,omitempty"`

	FareTypeGrouping *FareTypeGroupingInformationTypeI `xml:"fareTypeGrouping,omitempty"`
}

type FareQualifierDetailsType struct {
	DiscountDetails []*DiscountPenaltyInformationType `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type FareQualifierDetailsType_174783S struct {
	AdditionalFareDetails *AdditionalFareQualifierDetailsType `xml:"additionalFareDetails,omitempty"`
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

type FreeTextDetailsType struct {
	// xmlType: AlphaNumericString_Length1To3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// xmlType: AlphaNumericString_Length1To4
	InformationType string `xml:"informationType,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	Language string `xml:"language,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	Source string `xml:"source"`

	// xmlType: AlphaNumericString_Length1To3
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	// xmlType: AlphaNumericString_Length1To320
	FreeText []string `xml:"freeText"` // maxOccurs="99"
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

type ItemNumberIdentificationType struct {
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"`
}

type ItemNumberIdentificationTypeI struct {
	// Reference number or Fare component number
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"`
}

type ItemNumberType struct {
	ItemNumberDetails []*ItemNumberIdentificationType `xml:"itemNumberDetails"` // maxOccurs="99"
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

type LocationTypeI struct {
	// xmlType: AlphaNumericString_Length1To25
	TrueLocationId string `xml:"trueLocationId,omitempty"`
}

type LocationTypeI_52002C struct {
	// xmlType: AlphaNumericString_Length1To25
	TrueLocationId string `xml:"trueLocationId,omitempty"`

	// xmlType: AlphaNumericString_Length1To17
	TrueLocation string `xml:"trueLocation,omitempty"`
}

type MarriageControlDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To3
	Relation string `xml:"relation,omitempty"`

	MarriageIdentifier *int32 `xml:"marriageIdentifier,omitempty"`

	LineNumber *int32 `xml:"lineNumber,omitempty"`

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

type MonetaryInformationDetailsType struct {
	// xmlType: AlphaNumericString_Length1To3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	// xmlType: AlphaNumericString_Length1To35
	Amount string `xml:"amount,omitempty"`

	// Currency
	// xmlType: AlphaNumericString_Length1To3
	Currency string `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsTypeI struct {
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

type MonetaryInformationType struct {
	// Monetary information per fare component
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails"`

	// Other monetary information per fare component
	OtherMonetaryDetails []*MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"` // maxOccurs="19"
}

type MonetaryInformationTypeI struct {
	MonetaryDetails []*MonetaryInformationDetailsTypeI `xml:"monetaryDetails"` // maxOccurs="99"
}

type NumberOfUnitDetailsTypeI struct {
	// Total number of ticket/document numbers.
	NumberOfUnit int32 `xml:"numberOfUnit"`

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
	OriginatorId *int32 `xml:"originatorId,omitempty"`
}

type OriginatorOfRequestDetailsTypeI struct {
	DeliveringSystem *SystemDetailsTypeI `xml:"deliveringSystem"`

	OriginIdentification *OriginatorIdentificationDetailsTypeI `xml:"originIdentification,omitempty"`

	LocationDetails *LocationTypeI `xml:"locationDetails,omitempty"`

	CascadingSystem *SystemDetailsTypeI_83771C `xml:"cascadingSystem,omitempty"`

	// 1 character code for airline agent, travel agent, etc.
	// xmlType: AlphaNumericString_Length1To1
	OriginatorTypeCode string `xml:"originatorTypeCode,omitempty"`

	OriginDetails *OriginatorDetailsTypeI `xml:"originDetails,omitempty"`

	// A code identifying the issuing agent or if the transaction is being initiated by a computer programme then the word 'system'.
	// xmlType: AlphaNumericString_Length1To9
	Originator string `xml:"originator"`
}

type PricingOrTicketingSubsequentType struct {
	// RATE OR TARIFF CLASS INFORMATION
	FareBasisDetails *RateTariffClassInformationType `xml:"fareBasisDetails,omitempty"`
}

type PricingTicketingDetailsTypeI struct {
	PriceTicketDetails *PricingTicketingInformationTypeI `xml:"priceTicketDetails,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	PriceTariffType string `xml:"priceTariffType,omitempty"`

	ProductDateTimeDetails *ProductDateTimeTypeI_52035C `xml:"productDateTimeDetails,omitempty"`

	CompanyDetails *CompanyIdentificationTypeI_51997C `xml:"companyDetails,omitempty"`

	CompanyNumberDetails *CompanyIdentificationNumbersTypeI `xml:"companyNumberDetails,omitempty"`

	LocationDetails *LocationDetailsTypeI `xml:"locationDetails,omitempty"`

	OtherLocationDetails *LocationDetailsTypeI `xml:"otherLocationDetails,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	IdNumber string `xml:"idNumber,omitempty"`

	// xmlType: NumericDecimal_Length1To18
	MonetaryAmount *float64 `xml:"monetaryAmount,omitempty"`
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

	// Noshow indicator
	// xmlType: AlphaNumericString_Length1To3
	Indicators []string `xml:"indicators,omitempty"` // maxOccurs="11"
}

type PricingTicketingSubsequentTypeI struct {
	FareBasisDetails *RateTariffClassInformationTypeI `xml:"fareBasisDetails,omitempty"`

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

	DateVariation *int32 `xml:"dateVariation,omitempty"`
}

type ProductDateTimeTypeI_52035C struct {
	// xmlType: AlphaNumericString_Length1To35
	DepartureDate string `xml:"departureDate,omitempty"`

	//formats: Time24_HHMM
	DepartureTime string `xml:"departureTime,omitempty"`

	// xmlType: AlphaNumericString_Length1To35
	ArrivalDate string `xml:"arrivalDate,omitempty"`

	ArrivalTime *int32 `xml:"arrivalTime,omitempty"`

	DateVariation *int32 `xml:"dateVariation,omitempty"`
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

type ProductTypeDetailsType struct {
	// TST Connection Type
	// xmlType: AlphaNumericString_Length1To1
	FlightIndicator string `xml:"flightIndicator"`
}

type ProductTypeDetailsTypeI struct {
	// xmlType: AlphaNumericString_Length1To6
	FlightIndicator []string `xml:"flightIndicator"` // maxOccurs="9"
}

type RateTariffClassInformationType struct {
	// Fare Basis Code
	// xmlType: AlphaNumericString_Length1To35
	RateTariffClass string `xml:"rateTariffClass,omitempty"`

	// Ticket Designator
	// xmlType: AlphaNumericString_Length1To35
	OtherRateTariffClass string `xml:"otherRateTariffClass,omitempty"`
}

type RateTariffClassInformationTypeI struct {
	// Fare basis/ticket designator.
	// xmlType: AlphaNumericString_Length1To35
	RateTariffClass string `xml:"rateTariffClass,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	RateTariffIndicator string `xml:"rateTariffIndicator,omitempty"`
}

type ReferenceInfoType struct {
	ReferenceDetails *ReferencingDetailsType `xml:"referenceDetails"`
}

type ReferenceTypeI struct {
	// The coupon sequence number (1of3, 2of5)
	// xmlType: AlphaNumericString_Length1To6
	NumberOfItems string `xml:"numberOfItems,omitempty"`

	// The coupon/itinerary sequence number.
	// xmlType: AlphaNumericString_Length1To35
	LastItemIdentifier []string `xml:"lastItemIdentifier,omitempty"` // maxOccurs="99"
}

type ReferencingDetailsType struct {
	// xmlType: AlphaNumericString_Length1To10
	Type string `xml:"type"`

	// xmlType: AlphaNumericString_Length1To60
	Value string `xml:"value"`
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

type TaxDetailsType struct {
	// xmlType: AlphaNumericString_Length1To17
	Rate string `xml:"rate,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	CountryCode string `xml:"countryCode,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	CurrencyCode string `xml:"currencyCode,omitempty"`

	// xmlType: AlphaNumericString_Length1To3
	Type []string `xml:"type,omitempty"` // maxOccurs="99"
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

type TaxType struct {
	// xmlType: AlphaNumericString_Length1To3
	TaxCategory string `xml:"taxCategory,omitempty"`

	TaxDetails []*TaxDetailsType `xml:"taxDetails,omitempty"` // maxOccurs="99"
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
	BookingIataNumber *int32 `xml:"bookingIataNumber,omitempty"`
}

type TicketNumberDetailsTypeI struct {
	// Document number.
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"`

	// Document type.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`

	// Total number of booklets issued.
	NumberOfBooklets *int32 `xml:"numberOfBooklets,omitempty"`

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

type TransportIdentifierType struct {
	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"`
}

type TravelProductInformationType struct {
	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`

	// TST Connection Type
	FlightTypeDetails *ProductTypeDetailsType `xml:"flightTypeDetails,omitempty"`
}

type TravelProductInformationTypeI struct {
	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`
}

type TravelProductInformationTypeI_29340S struct {
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`

	BoardPointDetails *LocationTypeI_52002C `xml:"boardPointDetails,omitempty"`

	OffpointDetails *LocationTypeI_52002C `xml:"offpointDetails,omitempty"`

	CompanyDetails *CompanyIdentificationTypeI_51997C `xml:"companyDetails,omitempty"`

	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"`

	FlightTypeDetails *ProductTypeDetailsTypeI `xml:"flightTypeDetails,omitempty"`

	ItemNumber *int32 `xml:"itemNumber,omitempty"`

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

	Quantity *int32 `xml:"quantity,omitempty"`

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

	Age *int32 `xml:"age,omitempty"`
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

	// Specifies passenger type (adult, infant, military, etc.).
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`

	// Specify age of unaccompanied minor.
	Quantity *int32 `xml:"quantity,omitempty"`
}
