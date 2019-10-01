package Fare_InformativeBestPricingWithoutPNRRequest_v12_4 // tibnrq124

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TIBNRQ_12_4_1A Fare_InformativeBestPricingWithoutPNR"`

	// Contains general information about the message, especially the use case.
	MessageDetails *MessageActionDetailsTypeI `xml:"messageDetails"`

	OriginatorGroup *OriginatorGroup `xml:"originatorGroup,omitempty"`

	ConsumerReferenceInformation *ConsumerReferenceInformationType `xml:"consumerReferenceInformation,omitempty"`

	// Handles: * Fare conversion option * Fare selection option
	CurrencyOverride *ConversionRateTypeI `xml:"currencyOverride,omitempty"`

	// Used for Unifare pricing
	CorporateFareInfo *CorporateFareInformationType `xml:"corporateFareInfo,omitempty"`

	TaxExemptGroup *TaxExemptGroup `xml:"taxExemptGroup,omitempty"`

	// Form of payment at query level:  9888: type of form of payment 5004: amount (amount charged on the current FOP in case of multi FOP) 1154: Bin Number, i.e. the first six digits of any credit card
	GeneralFormOfPayment *FormOfPaymentTypeI `xml:"generalFormOfPayment,omitempty"`

	PassengersGroup []*PassengersGroup `xml:"passengersGroup,omitempty"` // maxOccurs="99"

	CabinPreferenceOption []*CabinPreferenceOption `xml:"cabinPreferenceOption,omitempty"` // maxOccurs="15"

	PricingOptionsGroup []*PricingOptionsGroup `xml:"pricingOptionsGroup,omitempty"` // maxOccurs="99"

	TripsGroup *TripsGroup `xml:"tripsGroup"`

	ObFeeRequestGroup *ObFeeRequestGroup `xml:"obFeeRequestGroup,omitempty"`
}

type OriginatorGroup struct {
	// Supplementary protocol or business related information.
	AdditionalBusinessInformation *AdditionalBusinessSourceInformationTypeI `xml:"additionalBusinessInformation,omitempty"`
}

type TaxExemptGroup struct {
	// Handles tax exemption option.
	TaxExempt *TaxTypeI `xml:"taxExempt"`
}

type PassengersGroup struct {
	// Contains: * Number of passengers in the group * Group tattoo
	SegmentRepetitionControl *SegmentRepetitionControlTypeI `xml:"segmentRepetitionControl"`

	// Passengers' tattoos provided by the carrier in case of LCC pricing with U2. NOT USED FOR FSC.
	TravellersID *SpecificTravellerTypeI `xml:"travellersID,omitempty"`

	PtcGroup []*PtcGroup `xml:"ptcGroup,omitempty"` // maxOccurs="3"
}

type PtcGroup struct {
	// PTC/Discount Code
	DiscountPtc *FareInformationTypeI `xml:"discountPtc"`

	// Form of payment at passenger level  9888: type of form of payment 5004: amount (amount charged on the current FOP in case of multi FOP) 1154: Bin Number, i.e. the first six digits of any credit card
	PassengerFormOfPayment *FormOfPaymentTypeI `xml:"passengerFormOfPayment,omitempty"`
}

type CabinPreferenceOption struct {
	// * Contains Preference priority
	PreferenceLevel *ProductInformationTypeI `xml:"preferenceLevel"`

	// * Contains Cabin option information
	CabinPreference *SeatRequestParametersType `xml:"cabinPreference,omitempty"`
}

type PricingOptionsGroup struct {
	// Handles the following pricing options: * Point of Sell override * Point of Ticketing override * Fare list * P/E ticket * Price by fare basis indicator * Validating Carrier * Award Publishing Carrier * No PNR Split (Best Pricer only)
	PricingDetails *PricingTicketingDetailsType `xml:"pricingDetails"`

	// * Expanded parameters * Forced PTC flag
	ExtPricingDetails *FareQualifierDetailsTypeI `xml:"extPricingDetails,omitempty"`
}

type TripsGroup struct {
	// This segment can be left empty.
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination"`

	SegmentGroup []*SegmentGroup `xml:"segmentGroup"` // maxOccurs="99"
}

type SegmentGroup struct {
	// Convey the information related to a segment (company, flight number, origin, destination...).
	SegmentInformation *TravelProductInformationTypeI `xml:"segmentInformation"`

	// Used for technical stops, even if it is currently deprecated.
	AdditionnalSegmentDetails *AdditionalProductDetailsTypeI `xml:"additionnalSegmentDetails,omitempty"`

	// *** deprecated ***  * A segment tattoo provided by the carrier in case of LCC pricing with U2. USELESS FOR FSC.
	SegmentPricingOptions *PricingTicketingDetailsTypeI `xml:"segmentPricingOptions,omitempty"`

	// Contain Zap Off details
	ZapOffDetails *FareQualifierDetailsTypeI `xml:"zapOffDetails,omitempty"`

	InventoryGroup *InventoryGroup `xml:"inventoryGroup,omitempty"`

	// Options at passenger level: PTC/Discount Tier Level
	PsgDetailsGroup []*GroupPassengerDetailsType `xml:"psgDetailsGroup,omitempty"` // maxOccurs="99"
}

type InventoryGroup struct {
	// Deprecated.  To store the flight inventory (opened classes and number of remaining seats) if known.
	Inventory *ProductInformationTypeI `xml:"inventory"`
}

type ObFeeRequestGroup struct {
	// Marker fee options
	MarkerFeeOptions *DummySegmentTypeI `xml:"markerFeeOptions"`

	FeeOptionInfoGroup []*FeeOptionInfoGroup `xml:"feeOptionInfoGroup"` // maxOccurs="5"
}

type FeeOptionInfoGroup struct {
	// Nature of the fee ( OB, OC, ...)
	FeeTypeInfo *SelectionDetailsType `xml:"feeTypeInfo"`

	// Associated tax rate
	RateTaxInfo *MonetaryInformationType `xml:"rateTaxInfo,omitempty"`

	FeeDetailsInfoGroup []*FeeDetailsInfoGroup `xml:"feeDetailsInfoGroup,omitempty"` // maxOccurs="99"
}

type FeeDetailsInfoGroup struct {
	// Fee information
	FeeInfo *SpecificDataInformationTypeI `xml:"feeInfo"`

	// Include, exclude
	FeeProcessingInfo *SelectionDetailsType `xml:"feeProcessingInfo"`

	// Associated amounts: amount to take into account to calculate the fee, ...
	AssociatedAmountsInfo *MonetaryInformationType_61169S `xml:"associatedAmountsInfo,omitempty"`
}

//
// Complex structs
//

type AdditionalBusinessSourceInformationTypeI struct {
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType"`

	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails,omitempty"`

	LocationDetails *LocationTypeI `xml:"locationDetails,omitempty"`

	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"`

	SystemCode formats.AlphaNumericString_Length1To35 `xml:"systemCode,omitempty"`
}

type AdditionalFareQualifierDetailsTypeI struct {
	RateClass formats.AlphaNumericString_Length1To35 `xml:"rateClass,omitempty"`

	CommodityCategory formats.AlphaNumericString_Length1To18 `xml:"commodityCategory,omitempty"`

	PricingGroup formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"`

	SecondRateClass []formats.AlphaNumericString_Length1To35 `xml:"secondRateClass,omitempty"` // maxOccurs="29"
}

type AdditionalProductDetailsTypeI struct {
	LegDetails *AdditionalProductTypeI `xml:"legDetails,omitempty"`

	DepartureStationInfo *StationInformationTypeI `xml:"departureStationInfo,omitempty"`

	ArrivalStationInfo *StationInformationTypeI `xml:"arrivalStationInfo,omitempty"`

	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"`

	TravellerTimeDetails *TravellerTimeDetailsTypeI `xml:"travellerTimeDetails,omitempty"`

	FacilitiesInformation []*ProductFacilitiesTypeI `xml:"facilitiesInformation,omitempty"` // maxOccurs="10"
}

type AdditionalProductTypeI struct {
	Equipment formats.AlphaNumericString_Length1To8 `xml:"equipment,omitempty"`

	NumberOfStops *formats.NumericInteger_Length1To3 `xml:"numberOfStops,omitempty"`

	Duration *formats.NumericInteger_Length1To6 `xml:"duration,omitempty"`

	Percentage *formats.NumericInteger_Length1To8 `xml:"percentage,omitempty"`

	DaysOfOperation formats.AlphaNumericString_Length1To7 `xml:"daysOfOperation,omitempty"`

	DateTimePeriod formats.AlphaNumericString_Length1To35 `xml:"dateTimePeriod,omitempty"`

	ComplexingFlightIndicator formats.AlphaNumericString_Length1To1 `xml:"complexingFlightIndicator,omitempty"`

	Locations []formats.AlphaNumericString_Length1To25 `xml:"locations,omitempty"` // maxOccurs="3"
}

type CompanyIdentificationNumbersTypeI struct {
	Identifier formats.AlphaNumericString_Length1To15 `xml:"identifier"`

	OtherIdentifier formats.AlphaNumericString_Length1To15 `xml:"otherIdentifier,omitempty"`
}

type CompanyIdentificationTypeI struct {
	MarketingCompany formats.AlphaNumericString_Length1To35 `xml:"marketingCompany,omitempty"`

	OperatingCompany formats.AlphaNumericString_Length1To35 `xml:"operatingCompany,omitempty"`

	OtherCompany formats.AlphaNumericString_Length1To35 `xml:"otherCompany,omitempty"`
}

type ConsumerReferenceIdentificationType struct {
	ReferenceQualifier formats.AMA_EDICodesetType_Length1to3 `xml:"referenceQualifier"`

	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber"`
}

type ConsumerReferenceInformationType struct {
	// Customer references
	CustomerReferences []*ConsumerReferenceIdentificationType `xml:"customerReferences"` // maxOccurs="20"
}

type ConversionRateDetailsTypeI struct {
	ConversionType formats.AlphaNumericString_Length1To3 `xml:"conversionType,omitempty"`

	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Not used
	RateType formats.AlphaNumericString_Length1To3 `xml:"rateType,omitempty"`

	// Not used
	PricingAmount *formats.NumericDecimal_Length1To18 `xml:"pricingAmount,omitempty"`

	// Not used
	ConvertedValueAmount *formats.NumericInteger_Length1To18 `xml:"convertedValueAmount,omitempty"`

	// Not used
	DutyTaxFeeType formats.AlphaNumericString_Length1To3 `xml:"dutyTaxFeeType,omitempty"`

	// Not used
	MeasurementValue *formats.NumericInteger_Length1To18 `xml:"measurementValue,omitempty"`

	// Not used
	MeasurementSignificance formats.AlphaNumericString_Length1To3 `xml:"measurementSignificance,omitempty"`
}

type ConversionRateTypeI struct {
	// Fare Conversion option (/R,FC-xxx): 6345 = currencty  Fare Selection option (/R,FS-xxx): 9875 = 700 6345 = currencty
	ConversionRateDetails *ConversionRateDetailsTypeI `xml:"conversionRateDetails"`

	// Fare Conversion option (/R,FC-xxx): 6345 = currencty  Fare Selection option (/R,FS-xxx): 9875 = 700 6345 = currencty
	OtherConvRateDetails []*ConversionRateDetailsTypeI `xml:"otherConvRateDetails,omitempty"` // maxOccurs="19"
}

type CorporateFareIdentifiersTypeI struct {
	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier,omitempty"`

	CorporateID []formats.AlphaNumericString_Length1To35 `xml:"corporateID,omitempty"` // maxOccurs="20"
}

type CorporateFareInformationType struct {
	CorporateFareIdentifiers []*CorporateFareIdentifiersTypeI `xml:"corporateFareIdentifiers"` // maxOccurs="20"
}

type DataTypeInformationTypeI struct {
	// Carrier fee code
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// Status event
	StatusEvent formats.AlphaNumericString_Length1To3 `xml:"statusEvent,omitempty"`
}

type DiscountPenaltyInformationTypeI struct {
	FareQualifier formats.AlphaNumericString_Length1To3 `xml:"fareQualifier"`

	RateCategory formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"`

	Amount *formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	Percentage *formats.NumericInteger_Length1To8 `xml:"percentage,omitempty"`
}

type DummySegmentTypeI struct {
}

type FareCategoryCodesTypeI struct {
	FareType formats.AlphaNumericString_Length1To20 `xml:"fareType"`

	OtherFareType []formats.AlphaNumericString_Length1To20 `xml:"otherFareType,omitempty"` // maxOccurs="8"
}

type FareDetailsTypeI struct {
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	Rate *formats.NumericInteger_Length1To8 `xml:"rate,omitempty"`

	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`

	FareCategory formats.AlphaNumericString_Length1To3 `xml:"fareCategory,omitempty"`
}

type FareDetailsTypeI_58070C struct {
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	Rate *formats.NumericDecimal_Length1To8 `xml:"rate,omitempty"`

	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`

	FareCategory formats.AlphaNumericString_Length1To3 `xml:"fareCategory,omitempty"`
}

type FareInformationTypeI struct {
	// PTC or Fare discount
	ValueQualifier formats.AlphaNumericString_Length1To3 `xml:"valueQualifier,omitempty"`

	Value *formats.NumericInteger_Length1To15 `xml:"value,omitempty"`

	FareDetails *FareDetailsTypeI_58070C `xml:"fareDetails,omitempty"`

	IdentityNumber formats.AlphaNumericString_Length1To35 `xml:"identityNumber,omitempty"`

	FareTypeGrouping *FareTypeGroupingInformationTypeI `xml:"fareTypeGrouping,omitempty"`

	RateCategory []formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"` // maxOccurs="9"
}

type FareInformationTypeI_133432S struct {
	ValueQualifier formats.AlphaNumericString_Length1To3 `xml:"valueQualifier,omitempty"`

	Value *formats.NumericInteger_Length1To15 `xml:"value,omitempty"`

	FareDetails *FareDetailsTypeI `xml:"fareDetails,omitempty"`

	IdentityNumber formats.AlphaNumericString_Length1To35 `xml:"identityNumber,omitempty"`

	FareTypeGrouping *FareTypeGroupingInformationTypeI `xml:"fareTypeGrouping,omitempty"`

	RateCategory []formats.AlphaNumericString_Length1To35 `xml:"rateCategory,omitempty"` // maxOccurs="9"
}

type FareQualifierDetailsTypeI struct {
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType,omitempty"`

	FareCategories *FareCategoryCodesTypeI `xml:"fareCategories,omitempty"`

	FareDetails *FareDetailsTypeI_58070C `xml:"fareDetails,omitempty"`

	AdditionalFareDetails *AdditionalFareQualifierDetailsTypeI `xml:"additionalFareDetails,omitempty"`

	DiscountDetails []*DiscountPenaltyInformationTypeI `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type FareTypeGroupingInformationTypeI struct {
	PricingGroup []formats.AlphaNumericString_Length1To35 `xml:"pricingGroup,omitempty"` // maxOccurs="5"
}

type FormOfPaymentDetailsTypeI struct {
	Type formats.AlphaNumericString_Length1To10 `xml:"type"`

	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	Amount *formats.NumericDecimal_Length1To18 `xml:"amount,omitempty"`

	VendorCode formats.AlphaNumericString_Length1To35 `xml:"vendorCode,omitempty"`

	CreditCardNumber formats.AlphaNumericString_Length1To35 `xml:"creditCardNumber,omitempty"`

	ExpiryDate formats.AlphaNumericString_Length1To35 `xml:"expiryDate,omitempty"`

	ApprovalCode formats.AlphaNumericString_Length1To17 `xml:"approvalCode,omitempty"`

	SourceOfApproval formats.AlphaNumericString_Length1To3 `xml:"sourceOfApproval,omitempty"`

	AuthorisedAmount *formats.NumericDecimal_Length1To18 `xml:"authorisedAmount,omitempty"`

	AddressVerification formats.AlphaNumericString_Length1To3 `xml:"addressVerification,omitempty"`

	CustomerAccount formats.AlphaNumericString_Length1To35 `xml:"customerAccount,omitempty"`

	ExtendedPayment formats.AlphaNumericString_Length1To3 `xml:"extendedPayment,omitempty"`

	FopFreeText formats.AlphaNumericString_Length1To70 `xml:"fopFreeText,omitempty"`

	MembershipStatus formats.AlphaNumericString_Length1To3 `xml:"membershipStatus,omitempty"`

	TransactionInfo formats.AlphaNumericString_Length1To35 `xml:"transactionInfo,omitempty"`
}

type FormOfPaymentTypeI struct {
	FormOfPayment *FormOfPaymentDetailsTypeI `xml:"formOfPayment"`

	OtherFormOfPayment []*FormOfPaymentDetailsTypeI `xml:"otherFormOfPayment,omitempty"` // maxOccurs="98"
}

type FrequentTravellerIdentificationCodeType struct {
	// Frequent Traveller Info
	FrequentTravellerDetails []*FrequentTravellerIdentificationType `xml:"frequentTravellerDetails"` // maxOccurs="99"
}

type FrequentTravellerIdentificationType struct {
	// Carrier where the FQTV is registered.
	Carrier formats.AlphaNumericString_Length1To35 `xml:"carrier,omitempty"`

	Number formats.AlphaNumericString_Length1To28 `xml:"number,omitempty"`

	// Specifies which traveller in the TIF segment the frequent traveller number applies (same as 9944 in TIF).
	CustomerReference formats.AlphaNumericString_Length1To10 `xml:"customerReference,omitempty"`

	// status code: 'OK' if the frequent flyer card has been validated
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// To specify a Tier linked to the FQTV
	TierLevel formats.AlphaNumericString_Length1To35 `xml:"tierLevel,omitempty"`

	// For example : priority code
	PriorityCode formats.AlphaNumericString_Length1To12 `xml:"priorityCode,omitempty"`

	// For example : Level description
	TierDescription formats.AlphaNumericString_Length1To35 `xml:"tierDescription,omitempty"`

	// For example : Company code of alliance
	CompanyCode formats.AlphaNumericString_Length1To35 `xml:"companyCode,omitempty"`

	CustomerValue *formats.NumericInteger_Length1To4 `xml:"customerValue,omitempty"`
}

type GenericDetailsTypeI struct {
	// not used
	CabinClassDesignator formats.AlphaString_Length1To1 `xml:"cabinClassDesignator,omitempty"`

	// Not used
	NoSmokingIndicator formats.AlphaString_Length1To1 `xml:"noSmokingIndicator,omitempty"`

	// Cabin class
	CabinClass formats.AlphaNumericString_Length1To1 `xml:"cabinClass,omitempty"`

	// Compartment designator
	CompartmentDesignator formats.AlphaString_Length1To1 `xml:"compartmentDesignator,omitempty"`

	// not used
	SeatCharacteristic []formats.AlphaNumericString_Length1To2 `xml:"seatCharacteristic,omitempty"` // maxOccurs="5"
}

type GroupPassengerDetailsType struct {
	// Trigger
	PassengerReference *SegmentRepetitionControlTypeI `xml:"passengerReference"`

	// Passengers' tattoos
	TravellersID *SpecificTravellerTypeI `xml:"travellersID,omitempty"`

	// pricing option at passenger level
	PsgDetailsInfo *PsgDetailsInfo `xml:"psgDetailsInfo"`
}

type PsgDetailsInfo struct {
	// PTC/Discount Code
	DiscountPtc *FareInformationTypeI_133432S `xml:"discountPtc"`

	// Tier level information
	FlequentFlyerDetails *FrequentTravellerIdentificationCodeType `xml:"flequentFlyerDetails,omitempty"`
}

type LocationDetailsTypeI struct {
	City formats.AlphaNumericString_Length1To25 `xml:"city,omitempty"`

	Country formats.AlphaNumericString_Length1To3 `xml:"country,omitempty"`
}

type LocationTypeI struct {
	TrueLocationId formats.AlphaNumericString_Length1To25 `xml:"trueLocationId,omitempty"`

	TrueLocation formats.AlphaNumericString_Length1To17 `xml:"trueLocation,omitempty"`
}

type MarriageControlDetailsTypeI struct {
	Relation formats.AlphaNumericString_Length1To3 `xml:"relation,omitempty"`

	MarriageIdentifier *formats.NumericInteger_Length1To10 `xml:"marriageIdentifier,omitempty"`

	LineNumber *formats.NumericInteger_Length1To6 `xml:"lineNumber,omitempty"`

	OtherRelation formats.AlphaNumericString_Length1To3 `xml:"otherRelation,omitempty"`

	CarrierCode formats.AlphaNumericString_Length1To35 `xml:"carrierCode,omitempty"`
}

type MessageActionDetailsTypeI struct {
	MessageFunctionDetails *MessageFunctionBusinessDetailsTypeI `xml:"messageFunctionDetails,omitempty"`

	ResponseType formats.AlphaNumericString_Length1To3 `xml:"responseType,omitempty"`
}

type MessageFunctionBusinessDetailsTypeI struct {
	BusinessFunction formats.AlphaNumericString_Length1To3 `xml:"businessFunction,omitempty"`

	MessageFunction formats.AlphaNumericString_Length1To3 `xml:"messageFunction,omitempty"`

	ResponsibleAgency formats.AlphaNumericString_Length1To3 `xml:"responsibleAgency,omitempty"`

	AdditionalMessageFunction []formats.AlphaNumericString_Length1To3 `xml:"additionalMessageFunction,omitempty"` // maxOccurs="20"
}

type MileageTimeDetailsTypeI struct {
	FlightLegMileage *formats.NumericInteger_Length1To18 `xml:"flightLegMileage,omitempty"`

	UnitQualifier formats.AlphaNumericString_Length1To3 `xml:"unitQualifier,omitempty"`

	ElapsedGroundTime *formats.NumericInteger_Length1To4 `xml:"elapsedGroundTime,omitempty"`
}

type MonetaryInformationDetailsType struct {
	// Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount formats.AlphaNumericString_Length1To12 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsType_96941C struct {
	// Qualifier
	TypeQualifier formats.AlphaNumericString_Length1To3 `xml:"typeQualifier"`

	// Amount
	Amount formats.AlphaNumericString_Length1To12 `xml:"amount,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Location
	Location formats.AlphaNumericString_Length1To25 `xml:"location,omitempty"`
}

type MonetaryInformationType struct {
	// Monetary info
	MonetaryDetails []*MonetaryInformationDetailsType `xml:"monetaryDetails"` // maxOccurs="20"
}

type MonetaryInformationType_61169S struct {
	// Monetary info
	MonetaryDetails []*MonetaryInformationDetailsType_96941C `xml:"monetaryDetails"` // maxOccurs="20"
}

type OriginAndDestinationDetailsTypeI struct {
	Origin formats.AlphaNumericString_Length1To25 `xml:"origin,omitempty"`

	Destination formats.AlphaNumericString_Length1To25 `xml:"destination,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	OriginatorId *formats.NumericInteger_Length1To9 `xml:"originatorId,omitempty"`

	InHouseIdentification1 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification1,omitempty"`

	InHouseIdentification2 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification2,omitempty"`

	InHouseIdentification3 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification3,omitempty"`
}

type PricingTicketingDetailsType struct {
	PriceTicketDetails *PricingTicketingInformationType `xml:"priceTicketDetails,omitempty"`

	PriceTariffType formats.AlphaNumericString_Length1To3 `xml:"priceTariffType,omitempty"`

	// Not used
	ProductDateTimeDetails *ProductDateTimeTypeI `xml:"productDateTimeDetails,omitempty"`

	// * Validating Carrier : 9906 : MarketingCompany  * Award Publishing Carrier :  9906 : OtherCompany
	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`

	// Controlling Carrier
	CompanyNumberDetails *CompanyIdentificationNumbersTypeI `xml:"companyNumberDetails,omitempty"`

	// Point Of Sale Override Information (3225 : City)
	LocationDetails *LocationDetailsTypeI `xml:"locationDetails,omitempty"`

	// Point Of Ticketing Override Information (3225 : City)
	OtherLocationDetails *LocationDetailsTypeI `xml:"otherLocationDetails,omitempty"`

	// PSR
	IdNumber formats.AlphaNumericString_Length1To35 `xml:"idNumber,omitempty"`

	MonetaryAmount *formats.NumericDecimal_Length1To18 `xml:"monetaryAmount,omitempty"`
}

type PricingTicketingDetailsTypeI struct {
	PriceTicketDetails *PricingTicketingInformationTypeI `xml:"priceTicketDetails,omitempty"`

	PriceTariffType formats.AlphaNumericString_Length1To3 `xml:"priceTariffType,omitempty"`

	ProductDateTimeDetails *ProductDateTimeTypeI `xml:"productDateTimeDetails,omitempty"`

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`

	CompanyNumberDetails *CompanyIdentificationNumbersTypeI `xml:"companyNumberDetails,omitempty"`

	LocationDetails *LocationDetailsTypeI `xml:"locationDetails,omitempty"`

	OtherLocationDetails *LocationDetailsTypeI `xml:"otherLocationDetails,omitempty"`

	IdNumber formats.AlphaNumericString_Length1To35 `xml:"idNumber,omitempty"`

	MonetaryAmount *formats.NumericDecimal_Length1To18 `xml:"monetaryAmount,omitempty"`
}

type PricingTicketingInformationType struct {
	Indicators []formats.AlphaNumericString_Length1To3 `xml:"indicators,omitempty"` // maxOccurs="20"
}

type PricingTicketingInformationTypeI struct {
	Indicators []formats.AlphaNumericString_Length1To3 `xml:"indicators,omitempty"` // maxOccurs="20"
}

type ProductDateTimeTypeI struct {
	DepartureDate formats.AlphaNumericString_Length1To35 `xml:"departureDate,omitempty"`

	DepartureTime *formats.NumericInteger_Length1To4 `xml:"departureTime,omitempty"`

	ArrivalDate formats.AlphaNumericString_Length1To35 `xml:"arrivalDate,omitempty"`

	ArrivalTime *formats.NumericInteger_Length1To4 `xml:"arrivalTime,omitempty"`

	DateVariation *formats.NumericInteger_Length1To1 `xml:"dateVariation,omitempty"`
}

type ProductDetailsTypeI struct {
	Designator formats.AlphaNumericString_Length1To17 `xml:"designator"`

	AvailabilityStatus formats.AlphaNumericString_Length1To3 `xml:"availabilityStatus,omitempty"`

	SpecialService formats.AlphaNumericString_Length1To3 `xml:"specialService,omitempty"`

	Option []formats.AlphaNumericString_Length1To7 `xml:"option,omitempty"` // maxOccurs="3"
}

type ProductFacilitiesTypeI struct {
	Code formats.AlphaNumericString_Length1To3 `xml:"code,omitempty"`

	Description formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`

	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	ExtensionCode []formats.AlphaNumericString_Length1To17 `xml:"extensionCode,omitempty"` // maxOccurs="26"
}

type ProductIdentificationDetailsTypeI struct {
	FlightNumber formats.AlphaNumericString_Length1To35 `xml:"flightNumber"`

	BookingClass formats.AlphaNumericString_Length1To17 `xml:"bookingClass,omitempty"`

	OperationalSuffix formats.AlphaNumericString_Length1To3 `xml:"operationalSuffix,omitempty"`

	Modifier []formats.AlphaNumericString_Length1To7 `xml:"modifier,omitempty"` // maxOccurs="3"
}

type ProductInformationTypeI struct {
	ProductDetailsQualifier formats.AlphaNumericString_Length1To3 `xml:"productDetailsQualifier,omitempty"`

	// Not used
	BookingClassDetails []*ProductDetailsTypeI `xml:"bookingClassDetails,omitempty"` // maxOccurs="26"
}

type ProductTypeDetailsTypeI struct {
	FlightIndicator []formats.AlphaNumericString_Length1To6 `xml:"flightIndicator"` // maxOccurs="9"
}

type RangeOfRowsDetailsTypeI struct {
	// Seat row number
	SeatRowNumber formats.NumericInteger_Length1To3 `xml:"seatRowNumber"`

	// Number of rows
	NumberOfRows *formats.NumericInteger_Length1To3 `xml:"numberOfRows,omitempty"`

	// Seat column
	SeatColumn []formats.AlphaNumericString_Length1To1 `xml:"seatColumn,omitempty"` // maxOccurs="20"
}

type SeatRequestParametersType struct {
	// Generic details
	GenericDetails *GenericDetailsTypeI `xml:"genericDetails,omitempty"`

	// not used
	RangeOfRowsDetails *RangeOfRowsDetailsTypeI `xml:"rangeOfRowsDetails,omitempty"`

	// not used
	ProcessingIndicator formats.AlphaNumericString_Length1To3 `xml:"processingIndicator,omitempty"`

	// not used
	ReferenceNumber formats.AlphaNumericString_Length1To35 `xml:"referenceNumber,omitempty"`

	// not used
	Description formats.AlphaNumericString_Length1To70 `xml:"description,omitempty"`
}

type SegmentRepetitionControlDetailsTypeI struct {
	Quantity *formats.NumericInteger_Length1To15 `xml:"quantity,omitempty"`

	NumberOfUnits *formats.NumericInteger_Length1To15 `xml:"numberOfUnits,omitempty"`

	TotalNumberOfItems *formats.NumericInteger_Length1To15 `xml:"totalNumberOfItems,omitempty"`
}

type SegmentRepetitionControlTypeI struct {
	SegmentControlDetails []*SegmentRepetitionControlDetailsTypeI `xml:"segmentControlDetails,omitempty"` // maxOccurs="9"
}

type SelectionDetailsInformationType struct {
	// Carrier fee type
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// Carrier fee status
	OptionInformation formats.AlphaNumericString_Length1To3 `xml:"optionInformation,omitempty"`
}

type SelectionDetailsType struct {
	// Carrier fees options
	CarrierFeeDetails *SelectionDetailsInformationType `xml:"carrierFeeDetails"`
}

type SourceTypeDetailsTypeI struct {
	SourceQualifier1 formats.AlphaNumericString_Length1To3 `xml:"sourceQualifier1"`

	SourceQualifier2 formats.AlphaNumericString_Length1To3 `xml:"sourceQualifier2,omitempty"`
}

type SpecificDataInformationTypeI struct {
	// Carrier fee information
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation"`
}

type SpecificTravellerDetailsTypeI struct {
	ReferenceNumber formats.AlphaNumericString_Length1To10 `xml:"referenceNumber,omitempty"`

	MeasurementValue *formats.NumericInteger_Length1To18 `xml:"measurementValue,omitempty"`

	FirstDate formats.AlphaNumericString_Length1To35 `xml:"firstDate,omitempty"`

	Surname formats.AlphaNumericString_Length1To70 `xml:"surname,omitempty"`

	FirstName formats.AlphaNumericString_Length1To70 `xml:"firstName,omitempty"`
}

type SpecificTravellerTypeI struct {
	TravellerDetails []*SpecificTravellerDetailsTypeI `xml:"travellerDetails,omitempty"` // maxOccurs="99"
}

type StationInformationTypeI struct {
	GateDescription formats.AlphaNumericString_Length1To6 `xml:"gateDescription,omitempty"`

	Terminal formats.AlphaNumericString_Length1To25 `xml:"terminal,omitempty"`

	Concourse formats.AlphaNumericString_Length1To25 `xml:"concourse,omitempty"`
}

type TaxDetailsTypeI struct {
	Rate formats.AlphaNumericString_Length1To17 `xml:"rate,omitempty"`

	CountryCode formats.AlphaNumericString_Length1To3 `xml:"countryCode,omitempty"`

	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`

	Type []formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"` // maxOccurs="99"
}

type TaxTypeI struct {
	TaxCategory formats.AlphaNumericString_Length1To3 `xml:"taxCategory,omitempty"`

	TaxDetails []*TaxDetailsTypeI `xml:"taxDetails,omitempty"` // maxOccurs="99"
}

type TravelProductInformationTypeI struct {
	FlightDate *ProductDateTimeTypeI `xml:"flightDate,omitempty"`

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`

	CompanyDetails *CompanyIdentificationTypeI `xml:"companyDetails,omitempty"`

	FlightIdentification *ProductIdentificationDetailsTypeI `xml:"flightIdentification,omitempty"`

	FlightTypeDetails *ProductTypeDetailsTypeI `xml:"flightTypeDetails,omitempty"`

	ItemNumber *formats.NumericInteger_Length1To6 `xml:"itemNumber,omitempty"`

	SpecialSegment formats.AlphaNumericString_Length1To3 `xml:"specialSegment,omitempty"`

	MarriageDetails []*MarriageControlDetailsTypeI `xml:"marriageDetails,omitempty"` // maxOccurs="99"
}

type TravellerTimeDetailsTypeI struct {
	DepartureTime *formats.NumericInteger_Length1To4 `xml:"departureTime,omitempty"`

	ArrivalTime *formats.NumericInteger_Length1To4 `xml:"arrivalTime,omitempty"`

	CheckInDateTime formats.AlphaNumericString_Length1To10 `xml:"checkInDateTime,omitempty"`
}
