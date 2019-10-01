package Fare_PricePNRWithBookingClassResponse_v14_1 // tpcbrr141

//import "encoding/xml"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_14_1_1A Fare_PricePNRWithBookingClassReply"`

	ApplicationError *ErrorGroupType `xml:"applicationError,omitempty"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`

	FareList []*FareList `xml:"fareList,omitempty"` // maxOccurs="99"
}

type FareList struct {
	// Pricing information such as pricing rule and sales indicator.
	PricingInformation *PricingTicketingSubsequentTypeI `xml:"pricingInformation"`

	// Fare reference number. Ordering information is not relevant here.
	FareReference *ItemReferencesAndVersionsType_94584S `xml:"fareReference"`

	// Fare Indicators
	FareIndicators *FareInformationType `xml:"fareIndicators,omitempty"`

	// Last date to ticket the fare.
	LastTktDate *StructuredDateTimeInformationType `xml:"lastTktDate,omitempty"`

	// Validating carrier identification.
	ValidatingCarrier *TransportIdentifierType `xml:"validatingCarrier,omitempty"`

	// Passenger/segment association of the fare is specified here.
	PaxSegReference *ReferenceInformationTypeI `xml:"paxSegReference"`

	FareDataInformation *MonetaryInformationType_187640S `xml:"fareDataInformation,omitempty"`

	TaxInformation []*TaxInformation `xml:"taxInformation,omitempty"` // maxOccurs="120"

	// Banker's rates are used to convert amounts of the TST (converts base fare to equivalent fare) 1st C661 : 1st bankers' rate which is a percentage (no currency) 2nd C661 : 2nd bankers' rate which is currency+amount.
	BankerRates *ConversionRateTypeI `xml:"bankerRates,omitempty"`

	PassengerInformation []*PassengerInformation `xml:"passengerInformation,omitempty"` // maxOccurs="99"

	// Origin and destination of the fare. 1st C3225 occurence : origin city. 2nd C3225 occurence : destination city
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"`

	SegmentInformation []*SegmentInformation `xml:"segmentInformation,omitempty"` // maxOccurs="96"

	// Other pricing information such as endorsement, tour name...
	OtherPricingInfo []*CodedAttributeType_39223S `xml:"otherPricingInfo,omitempty"` // maxOccurs="99"

	WarningInformation []*WarningInformation `xml:"warningInformation,omitempty"` // maxOccurs="99"

	AutomaticReissueInfo *AutomaticReissueInfo `xml:"automaticReissueInfo,omitempty"`

	// Corporate number
	CorporateInfo *CorporateFareInformationType `xml:"corporateInfo,omitempty"`

	FeeBreakdown []*FeeBreakdown `xml:"feeBreakdown,omitempty"` // maxOccurs="9"

	// convey the mileage information
	Mileage *AdditionalProductDetailsTypeI `xml:"mileage,omitempty"`

	// Details at fare component or at bound level.
	FareComponentDetailsGroup []*FareComponentDetailsType `xml:"fareComponentDetailsGroup,omitempty"` // maxOccurs="99"

	EndFareList *DummySegmentTypeI `xml:"endFareList"`
}

type TaxInformation struct {
	// Tax details
	TaxDetails *DutyTaxFeeDetailsTypeU `xml:"taxDetails"`

	// Amount details. If the tax is a passenger facility charge (PFC) the detail of the airports related taxes is given here.
	AmountDetails *MonetaryInformationTypeI `xml:"amountDetails,omitempty"`
}

type PassengerInformation struct {
	// Penalty/discount details specified in the request.
	PenDisInformation *DiscountAndPenaltyInformationTypeI_6128S `xml:"penDisInformation"`

	// Reference of passengers that have a type code.
	PassengerReference *ReferenceInformationTypeI `xml:"passengerReference,omitempty"`
}

type SegmentInformation struct {
	// Connection information.
	ConnexInformation *ConnectionTypeI `xml:"connexInformation"`

	// Details on open segments added to the price calculation. These open segments exist only in the fare calculated, they have no equivalent in the PNR itinerary. This segment gives also information on booking class for best buy transactions.
	SegDetails *TravelProductInformationTypeI_26322S `xml:"segDetails,omitempty"`

	// Fare basis information
	FareQualifier *FareQualifierDetailsTypeI `xml:"fareQualifier,omitempty"`

	// Validity information for this fare
	ValidityInformation []*StructuredDateTimeInformationType `xml:"validityInformation,omitempty"` // maxOccurs="2"

	// Baggage allowance information
	BagAllowanceInformation *ExcessBaggageTypeI `xml:"bagAllowanceInformation,omitempty"`

	// Reference of the segment associated to the group.
	SegmentReference *ReferenceInformationTypeI `xml:"segmentReference,omitempty"`

	// The segment order in the pricing response can be different than the one of the PNR itinerary (segments are reordered at price calculation time). This order inform,ation is conveyed by the sequence number. If this order information is not present then the order is by default the one of the PNR.
	SequenceInformation *ItemReferencesAndVersionsType `xml:"sequenceInformation,omitempty"`
}

type WarningInformation struct {
	// Fare warning information code.
	WarningCode *ApplicationErrorInformationType `xml:"warningCode"`

	// Description in free flow text of the warning concerning the fare.
	WarningText *InteractiveFreeTextTypeI_6759S `xml:"warningText,omitempty"`
}

type AutomaticReissueInfo struct {
	// This segment contains the original ticket number.
	TicketInfo *TicketNumberTypeI `xml:"ticketInfo"`

	// This segment contains the coupon number (in absolute) corresponding to the first coupon for use from the last flawn segment.
	CouponInfo *CouponInformationTypeI `xml:"couponInfo"`

	PaperCouponRange *PaperCouponRange `xml:"paperCouponRange,omitempty"`

	// Base fare Information
	BaseFareInfo *MonetaryInformationTypeI_20897S `xml:"baseFareInfo"`

	FirstDpiGroup *FirstDpiGroup `xml:"firstDpiGroup"`

	SecondDpiGroup *SecondDpiGroup `xml:"secondDpiGroup"`

	// this segment conveys specific reissue attributes like Revalidation flag.
	ReissueAttributes *CodedAttributeType `xml:"reissueAttributes,omitempty"`
}

type PaperCouponRange struct {
	// This segment contains the original ticket number.
	TicketInfo *TicketNumberTypeI `xml:"ticketInfo"`

	// This segment contains the coupon number (in absolute) corresponding to the first coupon for use from the last flawn segment.
	CouponInfo *CouponInformationTypeI `xml:"couponInfo"`
}

type FirstDpiGroup struct {
	// Penalty amount in reissue currency
	ReIssuePenalty *DiscountAndPenaltyInformationTypeI `xml:"reIssuePenalty"`

	// Reissue Informations
	ReissueInfo *MonetaryInformationTypeI_20897S `xml:"reissueInfo"`

	// Old Tax informations
	OldTaxInfo *MonetaryInformationTypeI_20897S `xml:"oldTaxInfo"`

	// Balance Reissue Informations
	ReissueBalanceInfo *MonetaryInformationTypeI_20897S `xml:"reissueBalanceInfo"`
}

type SecondDpiGroup struct {
	// Discount and penalty info.
	Penalty *DiscountAndPenaltyInformationTypeI `xml:"penalty"`

	// Residual Value information
	ResidualValueInfo *MonetaryInformationTypeI_20897S `xml:"residualValueInfo"`

	// Old Tax informations
	OldTaxInfo *MonetaryInformationTypeI_20897S `xml:"oldTaxInfo"`

	// Balance issue Informations
	IssueBalanceInfo *MonetaryInformationTypeI_20897S `xml:"issueBalanceInfo"`
}

type FeeBreakdown struct {
	// Nature of the fee (OB, OC)
	FeeType *SelectionDetailsTypeI `xml:"feeType"`

	FeeDetails []*FeeDetails `xml:"feeDetails,omitempty"` // maxOccurs="99"
}

type FeeDetails struct {
	// Fee information
	FeeInfo *SpecificDataInformationTypeI `xml:"feeInfo"`

	// Attributes of this fee (commercial description)
	FeeDescription *InteractiveFreeTextTypeI `xml:"feeDescription,omitempty"`

	// Fee associated amounts: amount with/without tax, total tax amount
	FeeAmounts *MonetaryInformationTypeI_39230S `xml:"feeAmounts,omitempty"`

	// taxes related to this fee
	FeeTaxes []*TaxTypeI `xml:"feeTaxes,omitempty"` // maxOccurs="99"
}

//
// Complex structs
//

type AdditionalFareQualifierDetailsTypeI struct {
	// Primary code of the fare basis. This is not a codeset but a free flow text field.
	PrimaryCode string `xml:"primaryCode,omitempty"`

	// Fare basis code of the fare basis. This is not a codeset but a free flow text field.
	FareBasisCode string `xml:"fareBasisCode,omitempty"`

	// Ticket designator of the fare basis
	TicketDesignator string `xml:"ticketDesignator,omitempty"`

	// For any query : discount ticket designator to be assigned by Fare Quote server. For any response : priced PTCs
	DiscTktDesignator string `xml:"discTktDesignator,omitempty"`
}

type AdditionalProductDetailsTypeI struct {
	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"`
}

type ApplicationErrorDetailType struct {
	// Code identifying the data validation error condition.
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	ErrorCategory string `xml:"errorCategory,omitempty"`

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorDetailType_48648C struct {
	// Code identifying the data validation error condition.
	ApplicationErrorCode string `xml:"applicationErrorCode"`

	// Identification of a code list.
	CodeListQualifier string `xml:"codeListQualifier,omitempty"`

	// Code identifying the agency responsible for a code list.
	CodeListResponsibleAgency string `xml:"codeListResponsibleAgency,omitempty"`
}

type ApplicationErrorInformationType struct {
	// Application error details.
	ApplicationErrorDetail *ApplicationErrorDetailType_48648C `xml:"applicationErrorDetail"`
}

type ApplicationErrorInformationType_84497S struct {
	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type BaggageDetailsTypeI struct {
	// Baggage allowance quantity (piece concept)
	BaggageQuantity *int32 `xml:"baggageQuantity,omitempty"`

	// Baggage allowance weight
	BaggageWeight *int32 `xml:"baggageWeight,omitempty"`

	// Baggage allowance type (weight/number)
	BaggageType string `xml:"baggageType,omitempty"`

	// Measurement unit for weighing baggage allowance
	MeasureUnit string `xml:"measureUnit,omitempty"`
}

type CodedAttributeInformationType struct {
	// provides the attribute Type
	AttributeType string `xml:"attributeType"`

	// provides a description for the attribute
	AttributeDescription string `xml:"attributeDescription,omitempty"`
}

type CodedAttributeInformationType_66047C struct {
	// provides the attribute Type
	AttributeType string `xml:"attributeType"`

	// provides a description for the attribute
	AttributeDescription string `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	// provides details for the Attribute
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"` // maxOccurs="99"
}

type CodedAttributeType_39223S struct {
	// provides details for the Attribute
	AttributeDetails []*CodedAttributeInformationType_66047C `xml:"attributeDetails"` // maxOccurs="5"
}

type CompanyIdentificationTypeI struct {
	// Carrier code
	CarrierCode string `xml:"carrierCode,omitempty"`
}

type CompanyIdentificationTypeI_222513C struct {
	// Carrier owner fo the fare family
	OtherCompany string `xml:"otherCompany,omitempty"`
}

type ConnectionDetailsTypeI struct {
	// Specify ARNK and surface segments  not included in the fare routing.
	RoutingInformation string `xml:"routingInformation,omitempty"`

	// Type of connection for the flight
	ConnexType string `xml:"connexType,omitempty"`
}

type ConnectionTypeI struct {
	// Connection details
	ConnecDetails *ConnectionDetailsTypeI `xml:"connecDetails"`
}

type ConversionRateDetailsTypeI struct {
	// Currency of the rate
	CurrencyCode string `xml:"currencyCode,omitempty"`

	// Amount/percentage
	Amount *float64 `xml:"amount,omitempty"`
}

type ConversionRateTypeI struct {
	// First rate detail.
	FirstRateDetail *ConversionRateDetailsTypeI `xml:"firstRateDetail"`

	// Second rate detail.
	SecondRateDetail *ConversionRateDetailsTypeI `xml:"secondRateDetail,omitempty"`
}

type CorporateFareIdentifiersTypeI struct {
	// Format limitations: an..3
	FareQualifier string `xml:"fareQualifier,omitempty"`

	// Format limitations: an..35
	CorporateID []string `xml:"corporateID,omitempty"` // maxOccurs="20"
}

type CorporateFareInformationType struct {
	CorporateFareIdentifiers []*CorporateFareIdentifiersTypeI `xml:"corporateFareIdentifiers"` // maxOccurs="20"
}

type CouponDetailsType struct {
	// Tattoo + type of the product identifying the coupon.
	ProductId *ReferenceInfoType `xml:"productId"`

	// Flight Connection Type
	FlightConnectionType *TravelProductInformationType `xml:"flightConnectionType,omitempty"`
}

type CouponInformationDetailsTypeI struct {
	// Coupon number
	CpnNumber string `xml:"cpnNumber"`
}

type CouponInformationTypeI struct {
	// Details on coupon
	CouponDetails *CouponInformationDetailsTypeI `xml:"couponDetails"`

	// Details on coupon
	OtherCouponDetails []*CouponInformationDetailsTypeI `xml:"otherCouponDetails,omitempty"` // maxOccurs="3"
}

type DataInformationTypeI struct {
	// fee attribute
	Indicator string `xml:"indicator,omitempty"`
}

type DataTypeInformationTypeI struct {
	// fee subcode
	Type string `xml:"type"`
}

type DiscountAndPenaltyInformationTypeI struct {
	// Used to specify penalty information
	PenDisData *DiscountPenaltyMonetaryInformationTypeI_29792C `xml:"penDisData,omitempty"`
}

type DiscountAndPenaltyInformationTypeI_6128S struct {
	// Qualify the type of information.  Penalties are not passenger associated and are pure monetary information. Discount are passenger associated but only discount code is specified.
	InfoQualifier string `xml:"infoQualifier,omitempty"`

	// Used to specify penalty information.
	PenDisData []*DiscountPenaltyMonetaryInformationTypeI `xml:"penDisData,omitempty"` // maxOccurs="9"
}

type DiscountPenaltyInformationType struct {
	// Used for codes in the AMADEUS code tables. Code Length is three alphanumeric characters.
	FareQualifier string `xml:"fareQualifier,omitempty"`
}

type DiscountPenaltyInformationTypeI struct {
	// Discount off type.
	ZapOffType string `xml:"zapOffType"`

	// Discount amount
	ZapOffAmount *float64 `xml:"zapOffAmount,omitempty"`

	// Discount percentage.
	ZapOffPercentage *int32 `xml:"zapOffPercentage,omitempty"`
}

type DiscountPenaltyMonetaryInformationTypeI struct {
	// Type of penalty.
	PenaltyType string `xml:"penaltyType,omitempty"`

	// The penalty amount can be described differently: amount/percentage.
	PenaltyQualifier string `xml:"penaltyQualifier,omitempty"`

	// Amount of the penalty.
	PenaltyAmount *float64 `xml:"penaltyAmount,omitempty"`

	// This discount code is defined by the airlines. This cannot be coded as airlines might apply any combination of letters for their discounts.
	DiscountCode string `xml:"discountCode,omitempty"`

	// Penalty currency code.
	PenaltyCurrency string `xml:"penaltyCurrency,omitempty"`
}

type DiscountPenaltyMonetaryInformationTypeI_29792C struct {
	// The amount Type can be a percentage or an amount
	PenaltyQualifier string `xml:"penaltyQualifier,omitempty"`

	// specify the value
	PenaltyAmount *float64 `xml:"penaltyAmount,omitempty"`

	// penalty currency code
	PenaltyCurrency string `xml:"penaltyCurrency,omitempty"`
}

type DummySegmentTypeI struct {
}

type DutyTaxFeeAccountDetailTypeU struct {
	// Iso country of the tax
	IsoCountry string `xml:"isoCountry"`
}

type DutyTaxFeeDetailsTypeU struct {
	// Tax data qualifier
	TaxQualifier string `xml:"taxQualifier"`

	// Tax type identification
	TaxIdentification *DutyTaxFeeTypeDetailsTypeU `xml:"taxIdentification"`

	// Type of the tax
	TaxType *DutyTaxFeeAccountDetailTypeU `xml:"taxType,omitempty"`

	// Nature of the tax
	TaxNature string `xml:"taxNature,omitempty"`

	// Exempt tax indicator. If an tax is Exempted no amount is provided for this tax.
	TaxExempt string `xml:"taxExempt,omitempty"`
}

type DutyTaxFeeTypeDetailsTypeU struct {
	// Tax type identifier
	TaxIdentifier string `xml:"taxIdentifier"`
}

type ErrorGroupType struct {
	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType_84497S `xml:"errorOrWarningCodeDetails"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType `xml:"errorWarningDescription,omitempty"`
}

type ExcessBaggageTypeI struct {
	// Baggage allowance information details
	BagAllowanceDetails *BaggageDetailsTypeI `xml:"bagAllowanceDetails,omitempty"`
}

type FareComponentDetailsType struct {
	FareComponentID *ItemNumberType `xml:"fareComponentID"`

	// Market information related to the fare component or to the bound.
	MarketFareComponent *TravelProductInformationTypeI `xml:"marketFareComponent,omitempty"`

	// Monetary Information.
	MonetaryInformation *MonetaryInformationType `xml:"monetaryInformation,omitempty"`

	// Component Class information
	ComponentClassInfo *PricingOrTicketingSubsequentType `xml:"componentClassInfo,omitempty"`

	// Fare Qualifier Detail
	FareQualifiersDetail *FareQualifierDetailsType `xml:"fareQualifiersDetail,omitempty"`

	// Details of the fare family used for this fare component
	FareFamilyDetails *FareFamilyType `xml:"fareFamilyDetails,omitempty"`

	// Carrier owner of the fare family
	FareFamilyOwner *TransportIdentifierType_156079S `xml:"fareFamilyOwner,omitempty"`

	// Used to specify coupons included in the fare component or in the bound.
	CouponDetailsGroup []*CouponDetailsType `xml:"couponDetailsGroup"` // maxOccurs="99"
}

type FareDetailsType struct {
	// fare indicators
	FareCategory string `xml:"fareCategory,omitempty"`
}

type FareFamilyDetailsType struct {
	// Commercial fare Family Short name
	CommercialFamily string `xml:"commercialFamily"`
}

type FareFamilyType struct {
	// Fare Family Short Name
	FareFamilyname string `xml:"fareFamilyname,omitempty"`

	// HIERARCHICAL ORDER WITHIN FARE FAMILY
	Hierarchy *int32 `xml:"hierarchy,omitempty"`

	// Indicates Commercial Fare Family Short names
	CommercialFamilyDetails []*FareFamilyDetailsType `xml:"commercialFamilyDetails,omitempty"` // maxOccurs="20"
}

type FareInformationType struct {
	FareDetails *FareDetailsType `xml:"fareDetails,omitempty"`
}

type FareQualifierDetailsType struct {
	DiscountDetails []*DiscountPenaltyInformationType `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type FareQualifierDetailsTypeI struct {
	// Type of movement for this segment to take into account by Fare Quote to calculate the fare.
	MovementType string `xml:"movementType,omitempty"`

	// Fare basis detail
	FareBasisDetails *AdditionalFareQualifierDetailsTypeI `xml:"fareBasisDetails,omitempty"`

	// Discount data for zap off to apply to price calculation.
	ZapOffDetails *DiscountPenaltyInformationTypeI `xml:"zapOffDetails,omitempty"`
}

type FreeTextDetailsType struct {
	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..4
	InformationType string `xml:"informationType,omitempty"`

	// Format limitations: an..3
	Status string `xml:"status,omitempty"`

	// Format limitations: an..35
	CompanyId string `xml:"companyId,omitempty"`

	// Format limitations: an..3
	Language string `xml:"language,omitempty"`

	// Format limitations: an..3
	Source string `xml:"source"`

	// Format limitations: an..3
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`

	// Free text and message sequence numbers of the remarks.
	FreeText []string `xml:"freeText"` // maxOccurs="99"
}

type FreeTextQualificationTypeI struct {
	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`
}

type InteractiveFreeTextTypeI struct {
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`

	// Format limitations: an..10
	FreeText string `xml:"freeText"`
}

type InteractiveFreeTextTypeI_6759S struct {
	// Free flow text describing the error
	ErrorFreeText string `xml:"errorFreeText,omitempty"`
}

type ItemNumberIdentificationType struct {
	// Item identification: number of the fare component or of the bound.
	Number string `xml:"number,omitempty"`

	// Item type: fare component (FC) or bound (BND).
	Type string `xml:"type,omitempty"`
}

type ItemNumberType struct {
	// Item identification: number of the fare component or of the bound.
	ItemNumberDetails []*ItemNumberIdentificationType `xml:"itemNumberDetails"` // maxOccurs="99"
}

type ItemReferencesAndVersionsType struct {
	// Identification details : order number
	SequenceSection *UniqueIdDescriptionType `xml:"sequenceSection,omitempty"`
}

type ItemReferencesAndVersionsType_94584S struct {
	// qualifies the type of the reference used. Code set to define
	ReferenceType string `xml:"referenceType,omitempty"`

	// Tattoo number
	UniqueReference *int32 `xml:"uniqueReference,omitempty"`
}

type LocationTypeI struct {
	// Format limitations: an..25
	TrueLocationId string `xml:"trueLocationId,omitempty"`
}

type LocationTypeI_47688C struct {
	// Code of the city.
	CityCode string `xml:"cityCode,omitempty"`
}

type MileageTimeDetailsTypeI struct {
	// mileage total associated to the TST
	TotalMileage int32 `xml:"totalMileage"`
}

type MonetaryInformationDetailsType struct {
	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	Amount string `xml:"amount,omitempty"`

	// Currency
	Currency string `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsTypeI struct {
	// Qualify the type of fare defined in this composite
	FareDataQualifier string `xml:"fareDataQualifier"`

	// Fare data amount
	FareAmount string `xml:"fareAmount,omitempty"`

	// Fare data currency code
	FareCurrency string `xml:"fareCurrency,omitempty"`

	// Location of the fare data (PFCs specific)
	FareLocation string `xml:"fareLocation,omitempty"`
}

type MonetaryInformationDetailsTypeI_37257C struct {
	// Type qualifier
	TypeQualifier string `xml:"typeQualifier"`

	// amount
	Amount string `xml:"amount"`

	// currency
	Currency string `xml:"currency,omitempty"`

	// location
	Location string `xml:"location,omitempty"`
}

type MonetaryInformationDetailsTypeI_63727C struct {
	// Qualifier
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	Amount string `xml:"amount,omitempty"`

	// Currency
	Currency string `xml:"currency,omitempty"`

	// Location
	Location string `xml:"location,omitempty"`
}

type MonetaryInformationDetailsType_223832C struct {
	// Format limitations: an..3
	FareDataQualifier string `xml:"fareDataQualifier"`

	// Amount
	FareAmount string `xml:"fareAmount,omitempty"`

	// Currency
	FareCurrency string `xml:"fareCurrency,omitempty"`

	// location
	FareLocation string `xml:"fareLocation,omitempty"`
}

type MonetaryInformationType struct {
	// Monetary information per fare component
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails"`

	// Other monetary information per fare component
	OtherMonetaryDetails []*MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"` // maxOccurs="19"
}

type MonetaryInformationTypeI struct {
	// Main fare data infomation, can b thee base or  the total fare information which are mandatory  anyway
	FareDataMainInformation *MonetaryInformationDetailsTypeI `xml:"fareDataMainInformation"`

	// Supplementary fare data information
	FareDataSupInformation []*MonetaryInformationDetailsTypeI `xml:"fareDataSupInformation,omitempty"` // maxOccurs="19"
}

type MonetaryInformationTypeI_20897S struct {
	// monetaryDetails
	MonetaryDetails *MonetaryInformationDetailsTypeI_37257C `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsTypeI_37257C `xml:"otherMonetaryDetails,omitempty"` // maxOccurs="5"
}

type MonetaryInformationTypeI_39230S struct {
	// Monetary info
	MonetaryDetails []*MonetaryInformationDetailsTypeI_63727C `xml:"monetaryDetails"` // maxOccurs="20"
}

type MonetaryInformationType_187640S struct {
	FareDataMainInformation *MonetaryInformationDetailsType_223832C `xml:"fareDataMainInformation"`

	FareDataSupInformation []*MonetaryInformationDetailsType_223832C `xml:"fareDataSupInformation,omitempty"` // maxOccurs="99"
}

type OriginAndDestinationDetailsTypeI struct {
	// Code of the city.
	CityCode []string `xml:"cityCode"` // maxOccurs="2"
}

type PricingOrTicketingSubsequentType struct {
	// RATE OR TARIFF CLASS INFORMATION
	FareBasisDetails *RateTariffClassInformationType `xml:"fareBasisDetails,omitempty"`
}

type PricingTicketingSubsequentTypeI struct {
	// Information on TST type.
	TstInformation *RateTariffClassInformationTypeI `xml:"tstInformation"`

	// International sales indicator
	SalesIndicator string `xml:"salesIndicator,omitempty"`

	// Fare calculation mode indicator. This indicator specifies the type fare.
	Fcmi string `xml:"fcmi"`

	// Information of original fare used to create TST. The TST is created from Best Fare ( possible or available).
	BestFareType string `xml:"bestFareType,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	// OPEN or AIR are the two identifications accepted.  OPEN means the segment described here is an open segment. AIR means that it is a valid AIR segment.
	Identification string `xml:"identification"`

	// to describe the transportation class.
	BookingClass string `xml:"bookingClass,omitempty"`

	// Class of service to use in order to price the extra segment.
	ClassOfService string `xml:"classOfService,omitempty"`
}

type ProductTypeDetailsType struct {
	// TST Connection Type
	FlightIndicator string `xml:"flightIndicator"`
}

type RateTariffClassInformationType struct {
	// Fare Basis Code
	RateTariffClass string `xml:"rateTariffClass,omitempty"`

	// Ticket Designator
	OtherRateTariffClass string `xml:"otherRateTariffClass,omitempty"`
}

type RateTariffClassInformationTypeI struct {
	// Indicator qualifying the type of TST (basically manual or automatic)
	TstIndicator string `xml:"tstIndicator"`
}

type ReferenceInfoType struct {
	ReferenceDetails *ReferencingDetailsType `xml:"referenceDetails"`
}

type ReferenceInformationTypeI struct {
	// Passenger/segment/TST/fare reference details
	RefDetails []*ReferencingDetailsTypeI `xml:"refDetails,omitempty"` // maxOccurs="99"
}

type ReferencingDetailsType struct {
	// Format limitations: an..10
	Type string `xml:"type"`

	// Format limitations: an..60
	Value string `xml:"value"`
}

type ReferencingDetailsTypeI struct {
	// Qualifyer of the reference (Pax/Seg/Tst/Fare tattoo)
	RefQualifier string `xml:"refQualifier,omitempty"`

	// Passenger/segment/TST/fare tattoo reference number
	RefNumber *int32 `xml:"refNumber,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// Record locator.
	ControlNumber string `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {
	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation"`
}

type SelectionDetailsInformationTypeI struct {
	// Format limitations: an..2
	Option string `xml:"option"`
}

type SelectionDetailsTypeI struct {
	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails"`
}

type SpecificDataInformationTypeI struct {
	// Carrier fee code
	DataTypeInformation *DataTypeInformationTypeI `xml:"dataTypeInformation"`

	// Carrier fee application code  (NI, NR, CM, NC)
	DataInformation []*DataInformationTypeI `xml:"dataInformation,omitempty"` // maxOccurs="99"
}

type StructuredDateTimeInformationType struct {
	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic string `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeType struct {
	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year *int32 `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month string `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day string `xml:"day,omitempty"`
}

type TaxDetailsTypeI struct {
	// Tax Amount
	Rate string `xml:"rate,omitempty"`

	// ISO code identifying Country
	CountryCode string `xml:"countryCode,omitempty"`

	// ISO code identifying currency
	CurrencyCode string `xml:"currencyCode,omitempty"`

	// Tax designator code
	Type string `xml:"type,omitempty"`

	// tax designator code.
	SecondType string `xml:"secondType,omitempty"`
}

type TaxTypeI struct {
	// Tax details
	TaxDetails []*TaxDetailsTypeI `xml:"taxDetails,omitempty"` // maxOccurs="99"
}

type TicketNumberDetailsTypeI struct {
	// Ticket number
	Number string `xml:"number"`

	// ticket type
	Type string `xml:"type,omitempty"`
}

type TicketNumberTypeI struct {
	// Details on the document
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`
}

type TransportIdentifierType struct {
	// Information related to validating carrier.
	CarrierInformation *CompanyIdentificationTypeI `xml:"carrierInformation,omitempty"`
}

type TransportIdentifierType_156079S struct {
	CompanyIdentification *CompanyIdentificationTypeI_222513C `xml:"companyIdentification,omitempty"`
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

type TravelProductInformationTypeI_26322S struct {
	// City of departure for this extra segment.
	DepartureCity *LocationTypeI_47688C `xml:"departureCity,omitempty"`

	// City of arrival for this extra segment.
	ArrivalCity *LocationTypeI_47688C `xml:"arrivalCity,omitempty"`

	// Airline detail information of the extra segment.
	AirlineDetail *CompanyIdentificationTypeI `xml:"airlineDetail,omitempty"`

	// Segment detail information.
	SegmentDetail *ProductIdentificationDetailsTypeI `xml:"segmentDetail,omitempty"`

	// Ticketing status for this segment. Relevant only in case of reply.
	TicketingStatus string `xml:"ticketingStatus,omitempty"`
}

type UniqueIdDescriptionType struct {
	// Number specifying the ordering information of the item described within a group.
	SequenceNumber *int32 `xml:"sequenceNumber,omitempty"`
}
