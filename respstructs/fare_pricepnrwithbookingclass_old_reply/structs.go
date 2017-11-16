package fare_pricepnrwithbookingclass_old_reply

//import "encoding/xml"

type FarePricePNRWithBookingClassOldReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TPCBRR_12_4_1A Fare_PricePNRWithBookingClassReply"`

	ApplicationError *ErrorGroupType `xml:"applicationError,omitempty"`  // minOccurs="0"

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`  // minOccurs="0"

	FareList []*FareList `xml:"fareList,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FareList struct {

	// Pricing information such as pricing rule and sales indicator.
	PricingInformation *PricingTicketingSubsequentTypeI `xml:"pricingInformation"`

	// Fare reference number. Ordering information is not relevant here.
	FareReference *ItemReferencesAndVersionsType_94584S `xml:"fareReference"`

	// Fare Indicators
	FareIndicators *FareInformationType `xml:"fareIndicators,omitempty"`  // minOccurs="0"

	// Last date to ticket the fare.
	LastTktDate *StructuredDateTimeInformationType `xml:"lastTktDate,omitempty"`  // minOccurs="0"

	// Validating carrier identification.
	ValidatingCarrier *TransportIdentifierType `xml:"validatingCarrier,omitempty"`  // minOccurs="0"

	// Passenger/segment association of the fare is specified here.
	PaxSegReference *ReferenceInformationTypeI `xml:"paxSegReference"`

	FareDataInformation *MonetaryInformationType `xml:"fareDataInformation,omitempty"`  // minOccurs="0"

	TaxInformation []*TaxInformation `xml:"taxInformation,omitempty"`  // minOccurs="0" maxOccurs="120"

	// Banker's rates are used to convert amounts of the TST (converts base fare to equivalent fare) 1st C661 : 1st bankers' rate which is a percentage (no currency) 2nd C661 : 2nd bankers' rate which is currency+amount.
	BankerRates *ConversionRateTypeI `xml:"bankerRates,omitempty"`  // minOccurs="0"

	PassengerInformation []*PassengerInformation `xml:"passengerInformation,omitempty"`  // minOccurs="0" maxOccurs="99"

	// Origin and destination of the fare. 1st C3225 occurence : origin city. 2nd C3225 occurence : destination city
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination,omitempty"`  // minOccurs="0"

	SegmentInformation []*SegmentInformation `xml:"segmentInformation,omitempty"`  // minOccurs="0" maxOccurs="96"

	// Other pricing information such as endorsement, tour name...
	OtherPricingInfo []*CodedAttributeType_39223S `xml:"otherPricingInfo,omitempty"`  // minOccurs="0" maxOccurs="99"

	WarningInformation []*WarningInformation `xml:"warningInformation,omitempty"`  // minOccurs="0" maxOccurs="99"

	AutomaticReissueInfo *AutomaticReissueInfo `xml:"automaticReissueInfo,omitempty"`  // minOccurs="0"

	// Corporate number
	CorporateInfo *CorporateFareInformationType `xml:"corporateInfo,omitempty"`  // minOccurs="0"

	FeeBreakdown []*FeeBreakdown `xml:"feeBreakdown,omitempty"`  // minOccurs="0" maxOccurs="9"

	// convey the mileage information
	Mileage *AdditionalProductDetailsTypeI `xml:"mileage,omitempty"`  // minOccurs="0"

	FareComponentDetailsGroup []*FareComponentDetailsType `xml:"fareComponentDetailsGroup,omitempty"`  // minOccurs="0" maxOccurs="99"

	EndFareList *DummySegmentTypeI `xml:"endFareList"`
}

type TaxInformation struct {

	// Tax details
	TaxDetails *DutyTaxFeeDetailsTypeU `xml:"taxDetails"`

	// Amount details. If the tax is a passenger facility charge (PFC) the detail of the airports related taxes is given here.
	AmountDetails *MonetaryInformationTypeI `xml:"amountDetails,omitempty"`  // minOccurs="0"
}

type PassengerInformation struct {

	// Penalty/discount details specified in the request.
	PenDisInformation *DiscountAndPenaltyInformationTypeI_6128S `xml:"penDisInformation"`

	// Reference of passengers that have a type code.
	PassengerReference *ReferenceInformationTypeI `xml:"passengerReference,omitempty"`  // minOccurs="0"
}

type SegmentInformation struct {

	// Connection information.
	ConnexInformation *ConnectionTypeI `xml:"connexInformation"`

	// Details on open segments added to the price calculation. These open segments exist only in the fare calculated, they have no equivalent in the PNR itinerary. This segment gives also information on booking class for best buy transactions.
	SegDetails *TravelProductInformationTypeI_26322S `xml:"segDetails,omitempty"`  // minOccurs="0"

	// Fare basis information
	FareQualifier *FareQualifierDetailsTypeI `xml:"fareQualifier,omitempty"`  // minOccurs="0"

	// Validity information for this fare
	ValidityInformation []*StructuredDateTimeInformationType `xml:"validityInformation,omitempty"`  // minOccurs="0" maxOccurs="2"

	// Baggage allowance information
	BagAllowanceInformation *ExcessBaggageTypeI `xml:"bagAllowanceInformation,omitempty"`  // minOccurs="0"

	// Reference of the segment associated to the group.
	SegmentReference *ReferenceInformationTypeI `xml:"segmentReference,omitempty"`  // minOccurs="0"

	// The segment order in the pricing response can be different than the one of the PNR itinerary (segments are reordered at price calculation time). This order inform,ation is conveyed by the sequence number. If this order information is not present then the order is by default the one of the PNR.
	SequenceInformation *ItemReferencesAndVersionsType `xml:"sequenceInformation,omitempty"`  // minOccurs="0"
}

type WarningInformation struct {

	// Fare warning information code.
	WarningCode *ApplicationErrorInformationType `xml:"warningCode"`

	// Description in free flow text of the warning concerning the fare.
	WarningText *InteractiveFreeTextTypeI_6759S `xml:"warningText,omitempty"`  // minOccurs="0"
}

type AutomaticReissueInfo struct {

	// This segment contains the original ticket number.
	TicketInfo *TicketNumberTypeI `xml:"ticketInfo"`

	// This segment contains the coupon number (in absolute) corresponding to the first coupon for use from the last flawn segment.
	CouponInfo *CouponInformationTypeI `xml:"couponInfo"`

	PaperCouponRange *PaperCouponRange `xml:"paperCouponRange,omitempty"`  // minOccurs="0"

	// Base fare Information
	BaseFareInfo *MonetaryInformationTypeI_20897S `xml:"baseFareInfo"`

	FirstDpiGroup *FirstDpiGroup `xml:"firstDpiGroup"`

	SecondDpiGroup *SecondDpiGroup `xml:"secondDpiGroup"`

	// this segment conveys specific reissue attributes like Revalidation flag.
	ReissueAttributes *CodedAttributeType `xml:"reissueAttributes,omitempty"`  // minOccurs="0"
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

	FeeDetails []*FeeDetails `xml:"feeDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type FeeDetails struct {

	// Fee information
	FeeInfo *SpecificDataInformationTypeI `xml:"feeInfo"`

	// Attributes of this fee (commercial description)
	FeeDescription *InteractiveFreeTextTypeI `xml:"feeDescription,omitempty"`  // minOccurs="0"

	// Fee associated amounts: amount with/without tax, total tax amount
	FeeAmounts *MonetaryInformationTypeI_39230S `xml:"feeAmounts,omitempty"`  // minOccurs="0"

	// taxes related to this fee
	FeeTaxes []*TaxTypeI `xml:"feeTaxes,omitempty"`  // minOccurs="0" maxOccurs="99"
}

//
// Complex structs
//

type AdditionalFareQualifierDetailsTypeI struct {

	// Primary code of the fare basis. This is not a codeset but a free flow text field.
	PrimaryCode string `xml:"primaryCode,omitempty"`  // minOccurs="0"

	// Fare basis code of the fare basis. This is not a codeset but a free flow text field.
	FareBasisCode string `xml:"fareBasisCode,omitempty"`  // minOccurs="0"

	// Ticket designator of the fare basis
	TicketDesignator string `xml:"ticketDesignator,omitempty"`  // minOccurs="0"

	// For any query : discount ticket designator to be assigned by Fare Quote server. For any response : priced PTCs
	DiscTktDesignator string `xml:"discTktDesignator,omitempty"`  // minOccurs="0"
}

type AdditionalProductDetailsTypeI struct {

	MileageTimeDetails *MileageTimeDetailsTypeI `xml:"mileageTimeDetails,omitempty"`  // minOccurs="0"
}

type ApplicationErrorDetailType struct {

	// Code identifying the data validation error condition.
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	ErrorCategory string `xml:"errorCategory,omitempty"`  // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`  // minOccurs="0"
}

type ApplicationErrorDetailType_48648C struct {

	// Code identifying the data validation error condition.
	ApplicationErrorCode string `xml:"applicationErrorCode"`

	// Identification of a code list.
	CodeListQualifier string `xml:"codeListQualifier,omitempty"`  // minOccurs="0"

	// Code identifying the agency responsible for a code list.
	CodeListResponsibleAgency string `xml:"codeListResponsibleAgency,omitempty"`  // minOccurs="0"
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
	BaggageQuantity *int32 `xml:"baggageQuantity,omitempty"`  // minOccurs="0"

	// Baggage allowance weight
	BaggageWeight *int32 `xml:"baggageWeight,omitempty"`  // minOccurs="0"

	// Baggage allowance type (weight/number)
	BaggageType string `xml:"baggageType,omitempty"`  // minOccurs="0"

	// Measurement unit for weighing baggage allowance
	MeasureUnit string `xml:"measureUnit,omitempty"`  // minOccurs="0"
}

type CodedAttributeInformationType struct {

	// provides the attribute Type
	AttributeType string `xml:"attributeType"`

	// provides a description for the attribute
	AttributeDescription string `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeInformationType_66047C struct {

	// provides the attribute Type
	AttributeType string `xml:"attributeType"`

	// provides a description for the attribute
	AttributeDescription string `xml:"attributeDescription,omitempty"`  // minOccurs="0"
}

type CodedAttributeType struct {

	// provides details for the Attribute
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"`  // maxOccurs="99"
}

type CodedAttributeType_39223S struct {

	// provides details for the Attribute
	AttributeDetails []*CodedAttributeInformationType_66047C `xml:"attributeDetails"`  // maxOccurs="5"
}

type CompanyIdentificationTypeI struct {

	// Carrier code
	CarrierCode string `xml:"carrierCode,omitempty"`  // minOccurs="0"
}

type ConnectionDetailsTypeI struct {

	// Specify ARNK and surface segments  not included in the fare routing.
	RoutingInformation string `xml:"routingInformation,omitempty"`  // minOccurs="0"

	// Type of connection for the flight
	ConnexType string `xml:"connexType,omitempty"`  // minOccurs="0"
}

type ConnectionTypeI struct {

	// Connection details
	ConnecDetails *ConnectionDetailsTypeI `xml:"connecDetails"`
}

type ConversionRateDetailsTypeI struct {

	// Currency of the rate
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// Amount/percentage
	Amount *float64 `xml:"amount,omitempty"`  // minOccurs="0"
}

type ConversionRateTypeI struct {

	// First rate detail.
	FirstRateDetail *ConversionRateDetailsTypeI `xml:"firstRateDetail"`

	// Second rate detail.
	SecondRateDetail *ConversionRateDetailsTypeI `xml:"secondRateDetail,omitempty"`  // minOccurs="0"
}

type CorporateFareIdentifiersTypeI struct {

	// Format limitations: an..3
	FareQualifier string `xml:"fareQualifier,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	CorporateID []string `xml:"corporateID,omitempty"`  // minOccurs="0" maxOccurs="20"
}

type CorporateFareInformationType struct {

	CorporateFareIdentifiers []*CorporateFareIdentifiersTypeI `xml:"corporateFareIdentifiers"`  // maxOccurs="20"
}

type CouponDetailsType struct {

	// Tattoo + type of the product identifying the coupon.
	ProductId *ReferenceInfoType `xml:"productId"`

	// Flight Connection Type
	FlightConnectionType *TravelProductInformationType `xml:"flightConnectionType,omitempty"`  // minOccurs="0"
}

type CouponInformationDetailsTypeI struct {

	// Coupon number
	CpnNumber string `xml:"cpnNumber"`
}

type CouponInformationTypeI struct {

	// Details on coupon
	CouponDetails *CouponInformationDetailsTypeI `xml:"couponDetails"`

	// Details on coupon
	OtherCouponDetails []*CouponInformationDetailsTypeI `xml:"otherCouponDetails,omitempty"`  // minOccurs="0" maxOccurs="3"
}

type DataInformationTypeI struct {

	// fee attribute
	Indicator string `xml:"indicator,omitempty"`  // minOccurs="0"
}

type DataTypeInformationTypeI struct {

	// fee subcode
	Type string `xml:"type"`
}

type DiscountAndPenaltyInformationTypeI struct {

	// Used to specify penalty information
	PenDisData *DiscountPenaltyMonetaryInformationTypeI_29792C `xml:"penDisData,omitempty"`  // minOccurs="0"
}

type DiscountAndPenaltyInformationTypeI_6128S struct {

	// Qualify the type of information.  Penalties are not passenger associated and are pure monetary information. Discount are passenger associated but only discount code is specified.
	InfoQualifier string `xml:"infoQualifier,omitempty"`  // minOccurs="0"

	// Used to specify penalty information.
	PenDisData []*DiscountPenaltyMonetaryInformationTypeI `xml:"penDisData,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type DiscountPenaltyInformationType struct {

	// Used for codes in the AMADEUS code tables. Code Length is three alphanumeric characters.
	FareQualifier string `xml:"fareQualifier,omitempty"`  // minOccurs="0"
}

type DiscountPenaltyInformationTypeI struct {

	// Discount off type.
	ZapOffType string `xml:"zapOffType"`

	// Discount amount
	ZapOffAmount *float64 `xml:"zapOffAmount,omitempty"`  // minOccurs="0"

	// Discount percentage.
	ZapOffPercentage *int32 `xml:"zapOffPercentage,omitempty"`  // minOccurs="0"
}

type DiscountPenaltyMonetaryInformationTypeI struct {

	// Type of penalty.
	PenaltyType string `xml:"penaltyType,omitempty"`  // minOccurs="0"

	// The penalty amount can be described differently: amount/percentage.
	PenaltyQualifier string `xml:"penaltyQualifier,omitempty"`  // minOccurs="0"

	// Amount of the penalty.
	PenaltyAmount *float64 `xml:"penaltyAmount,omitempty"`  // minOccurs="0"

	// This discount code is defined by the airlines. This cannot be coded as airlines might apply any combination of letters for their discounts.
	DiscountCode string `xml:"discountCode,omitempty"`  // minOccurs="0"

	// Penalty currency code.
	PenaltyCurrency string `xml:"penaltyCurrency,omitempty"`  // minOccurs="0"
}

type DiscountPenaltyMonetaryInformationTypeI_29792C struct {

	// The amount Type can be a percentage or an amount
	PenaltyQualifier string `xml:"penaltyQualifier,omitempty"`  // minOccurs="0"

	// specify the value
	PenaltyAmount *float64 `xml:"penaltyAmount,omitempty"`  // minOccurs="0"

	// penalty currency code
	PenaltyCurrency string `xml:"penaltyCurrency,omitempty"`  // minOccurs="0"
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
	TaxType *DutyTaxFeeAccountDetailTypeU `xml:"taxType,omitempty"`  // minOccurs="0"

	// Nature of the tax
	TaxNature string `xml:"taxNature,omitempty"`  // minOccurs="0"

	// Exempt tax indicator. If an tax is Exempted no amount is provided for this tax.
	TaxExempt string `xml:"taxExempt,omitempty"`  // minOccurs="0"
}

type DutyTaxFeeTypeDetailsTypeU struct {

	// Tax type identifier
	TaxIdentifier string `xml:"taxIdentifier"`
}

type ErrorGroupType struct {

	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType_84497S `xml:"errorOrWarningCodeDetails"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType `xml:"errorWarningDescription,omitempty"`  // minOccurs="0"
}

type ExcessBaggageTypeI struct {

	// Baggage allowance information details
	BagAllowanceDetails *BaggageDetailsTypeI `xml:"bagAllowanceDetails,omitempty"`  // minOccurs="0"
}

type FareComponentDetailsType struct {

	// fare Component identification
	FareComponentID *ItemNumberType `xml:"fareComponentID"`

	// Market information related to fare component
	MarketFareComponent *TravelProductInformationTypeI `xml:"marketFareComponent,omitempty"`  // minOccurs="0"

	// Monetary Information
	MonetaryInformation *MonetaryInformationType_157196S `xml:"monetaryInformation,omitempty"`  // minOccurs="0"

	// Component Class information
	ComponentClassInfo *PricingOrTicketingSubsequentType `xml:"componentClassInfo,omitempty"`  // minOccurs="0"

	// Fare Qualifier Detail
	FareQualifiersDetail *FareQualifierDetailsType `xml:"fareQualifiersDetail,omitempty"`  // minOccurs="0"

	CouponDetailsGroup []*CouponDetailsType `xml:"couponDetailsGroup"`  // maxOccurs="99"
}

type FareDetailsType struct {

	// fare indicators
	FareCategory string `xml:"fareCategory,omitempty"`  // minOccurs="0"
}

type FareInformationType struct {

	FareDetails *FareDetailsType `xml:"fareDetails,omitempty"`  // minOccurs="0"
}

type FareQualifierDetailsType struct {

	DiscountDetails []*DiscountPenaltyInformationType `xml:"discountDetails,omitempty"`  // minOccurs="0" maxOccurs="9"
}

type FareQualifierDetailsTypeI struct {

	// Type of movement for this segment to take into account by Fare Quote to calculate the fare.
	MovementType string `xml:"movementType,omitempty"`  // minOccurs="0"

	// Fare basis detail
	FareBasisDetails *AdditionalFareQualifierDetailsTypeI `xml:"fareBasisDetails,omitempty"`  // minOccurs="0"

	// Discount data for zap off to apply to price calculation.
	ZapOffDetails *DiscountPenaltyInformationTypeI `xml:"zapOffDetails,omitempty"`  // minOccurs="0"
}

type FreeTextDetailsType struct {

	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Format limitations: an..4
	InformationType string `xml:"informationType,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Status string `xml:"status,omitempty"`  // minOccurs="0"

	// Format limitations: an..35
	CompanyId string `xml:"companyId,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Language string `xml:"language,omitempty"`  // minOccurs="0"

	// Format limitations: an..3
	Source string `xml:"source"`

	// Format limitations: an..3
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {

	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails,omitempty"`  // minOccurs="0"

	// Free text and message sequence numbers of the remarks.
	FreeText []string `xml:"freeText"`  // maxOccurs="99"
}

type FreeTextQualificationTypeI struct {

	// Format limitations: an..3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`
}

type InteractiveFreeTextTypeI struct {

	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`  // minOccurs="0"

	// Format limitations: an..10
	FreeText string `xml:"freeText"`
}

type InteractiveFreeTextTypeI_6759S struct {

	// Free flow text describing the error
	ErrorFreeText string `xml:"errorFreeText,omitempty"`  // minOccurs="0"
}

type ItemNumberIdentificationType struct {

	// Format limitations: an..35
	Number string `xml:"number,omitempty"`  // minOccurs="0"
}

type ItemNumberType struct {

	ItemNumberDetails []*ItemNumberIdentificationType `xml:"itemNumberDetails"`  // maxOccurs="99"
}

type ItemReferencesAndVersionsType struct {

	// Identification details : order number
	SequenceSection *UniqueIdDescriptionType `xml:"sequenceSection,omitempty"`  // minOccurs="0"
}

type ItemReferencesAndVersionsType_94584S struct {

	// qualifies the type of the reference used. Code set to define
	ReferenceType string `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Tattoo number
	UniqueReference *int32 `xml:"uniqueReference,omitempty"`  // minOccurs="0"
}

type LocationTypeI struct {

	// Format limitations: an..25
	TrueLocationId string `xml:"trueLocationId,omitempty"`  // minOccurs="0"
}

type LocationTypeI_47688C struct {

	// Code of the city.
	CityCode string `xml:"cityCode,omitempty"`  // minOccurs="0"
}

type MileageTimeDetailsTypeI struct {

	// mileage total associated to the TST
	TotalMileage int32 `xml:"totalMileage"`
}

type MonetaryInformationDetailsType struct {

	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	Amount string `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI struct {

	// Qualify the type of fare defined in this composite
	FareDataQualifier string `xml:"fareDataQualifier"`

	// Fare data amount
	FareAmount string `xml:"fareAmount,omitempty"`  // minOccurs="0"

	// Fare data currency code
	FareCurrency string `xml:"fareCurrency,omitempty"`  // minOccurs="0"

	// Location of the fare data (PFCs specific)
	FareLocation string `xml:"fareLocation,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_37257C struct {

	// Type qualifier
	TypeQualifier string `xml:"typeQualifier"`

	// amount
	Amount string `xml:"amount"`

	// currency
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// location
	Location string `xml:"location,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsTypeI_63727C struct {

	// Qualifier
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	Amount string `xml:"amount,omitempty"`  // minOccurs="0"

	// Currency
	Currency string `xml:"currency,omitempty"`  // minOccurs="0"

	// Location
	Location string `xml:"location,omitempty"`  // minOccurs="0"
}

type MonetaryInformationDetailsType_223826C struct {

	// Format limitations: an..3
	FareDataQualifier string `xml:"fareDataQualifier"`

	// Amount
	FareAmount string `xml:"fareAmount,omitempty"`  // minOccurs="0"

	// Currency
	FareCurrency string `xml:"fareCurrency,omitempty"`  // minOccurs="0"

	// location
	FareLocation string `xml:"fareLocation,omitempty"`  // minOccurs="0"
}

type MonetaryInformationType struct {

	FareDataMainInformation *MonetaryInformationDetailsType_223826C `xml:"fareDataMainInformation"`

	FareDataSupInformation []*MonetaryInformationDetailsType_223826C `xml:"fareDataSupInformation,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type MonetaryInformationTypeI struct {

	// Main fare data infomation, can b thee base or  the total fare information which are mandatory  anyway
	FareDataMainInformation *MonetaryInformationDetailsTypeI `xml:"fareDataMainInformation"`

	// Supplementary fare data information
	FareDataSupInformation []*MonetaryInformationDetailsTypeI `xml:"fareDataSupInformation,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type MonetaryInformationTypeI_20897S struct {

	// monetaryDetails
	MonetaryDetails *MonetaryInformationDetailsTypeI_37257C `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsTypeI_37257C `xml:"otherMonetaryDetails,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type MonetaryInformationTypeI_39230S struct {

	// Monetary info
	MonetaryDetails []*MonetaryInformationDetailsTypeI_63727C `xml:"monetaryDetails"`  // maxOccurs="20"
}

type MonetaryInformationType_157196S struct {

	// Monetary information per fare component
	MonetaryDetails *MonetaryInformationDetailsType `xml:"monetaryDetails"`

	// Other monetary information per fare component
	OtherMonetaryDetails []*MonetaryInformationDetailsType `xml:"otherMonetaryDetails,omitempty"`  // minOccurs="0" maxOccurs="19"
}

type OriginAndDestinationDetailsTypeI struct {

	// Code of the city.
	CityCode []string `xml:"cityCode"`  // maxOccurs="2"
}

type PricingOrTicketingSubsequentType struct {

	// RATE OR TARIFF CLASS INFORMATION
	FareBasisDetails *RateTariffClassInformationType `xml:"fareBasisDetails,omitempty"`  // minOccurs="0"
}

type PricingTicketingSubsequentTypeI struct {

	// Information on TST type.
	TstInformation *RateTariffClassInformationTypeI `xml:"tstInformation"`

	// International sales indicator
	SalesIndicator string `xml:"salesIndicator,omitempty"`  // minOccurs="0"

	// Fare calculation mode indicator. This indicator specifies the type fare.
	Fcmi string `xml:"fcmi"`

	// Information of original fare used to create TST. The TST is created from Best Fare ( possible or available).
	BestFareType string `xml:"bestFareType,omitempty"`  // minOccurs="0"
}

type ProductIdentificationDetailsTypeI struct {

	// OPEN or AIR are the two identifications accepted.  OPEN means the segment described here is an open segment. AIR means that it is a valid AIR segment.
	Identification string `xml:"identification"`

	// to describe the transportation class.
	BookingClass string `xml:"bookingClass,omitempty"`  // minOccurs="0"

	// Class of service to use in order to price the extra segment.
	ClassOfService string `xml:"classOfService,omitempty"`  // minOccurs="0"
}

type ProductTypeDetailsType struct {

	// TST Connection Type
	FlightIndicator string `xml:"flightIndicator"`
}

type RateTariffClassInformationType struct {

	// Fare Basis Code
	RateTariffClass string `xml:"rateTariffClass,omitempty"`  // minOccurs="0"

	// Ticket Designator
	OtherRateTariffClass string `xml:"otherRateTariffClass,omitempty"`  // minOccurs="0"
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
	RefDetails []*ReferencingDetailsTypeI `xml:"refDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferencingDetailsType struct {

	// Format limitations: an..10
	Type string `xml:"type"`

	// Format limitations: an..60
	Value string `xml:"value"`
}

type ReferencingDetailsTypeI struct {

	// Qualifyer of the reference (Pax/Seg/Tst/Fare tattoo)
	RefQualifier string `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Passenger/segment/TST/fare tattoo reference number
	RefNumber *int32 `xml:"refNumber,omitempty"`  // minOccurs="0"
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
	DataInformation []*DataInformationTypeI `xml:"dataInformation,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type StructuredDateTimeInformationType struct {

	// This data element can be used to provide the semantic of the information provided. Examples : - Impacted period - Departure date - Estimated arrival date and time
	BusinessSemantic string `xml:"businessSemantic,omitempty"`  // minOccurs="0"

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`  // minOccurs="0"
}

type StructuredDateTimeType struct {

	// Year number. The format is a little long for short term usage but it can be reduced by implementation if required.
	Year *int32 `xml:"year,omitempty"`  // minOccurs="0"

	// Month number in the year ( begins to 1 )
	Month string `xml:"month,omitempty"`  // minOccurs="0"

	// Day number in the month ( begins to 1 )
	Day string `xml:"day,omitempty"`  // minOccurs="0"
}

type TaxDetailsTypeI struct {

	// Tax Amount
	Rate string `xml:"rate,omitempty"`  // minOccurs="0"

	// ISO code identifying Country
	CountryCode string `xml:"countryCode,omitempty"`  // minOccurs="0"

	// ISO code identifying currency
	CurrencyCode string `xml:"currencyCode,omitempty"`  // minOccurs="0"

	// Tax designator code
	Type string `xml:"type,omitempty"`  // minOccurs="0"

	// tax designator code.
	SecondType string `xml:"secondType,omitempty"`  // minOccurs="0"
}

type TaxTypeI struct {

	// Tax details
	TaxDetails []*TaxDetailsTypeI `xml:"taxDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type TicketNumberDetailsTypeI struct {

	// Ticket number
	Number string `xml:"number"`

	// ticket type
	Type string `xml:"type,omitempty"`  // minOccurs="0"
}

type TicketNumberTypeI struct {

	// Details on the document
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`
}

type TransportIdentifierType struct {

	// Information related to validating carrier.
	CarrierInformation *CompanyIdentificationTypeI `xml:"carrierInformation,omitempty"`  // minOccurs="0"
}

type TravelProductInformationType struct {

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`  // minOccurs="0"

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`  // minOccurs="0"

	// TST Connection Type
	FlightTypeDetails *ProductTypeDetailsType `xml:"flightTypeDetails,omitempty"`  // minOccurs="0"
}

type TravelProductInformationTypeI struct {

	BoardPointDetails *LocationTypeI `xml:"boardPointDetails,omitempty"`  // minOccurs="0"

	OffpointDetails *LocationTypeI `xml:"offpointDetails,omitempty"`  // minOccurs="0"
}

type TravelProductInformationTypeI_26322S struct {

	// City of departure for this extra segment.
	DepartureCity *LocationTypeI_47688C `xml:"departureCity,omitempty"`  // minOccurs="0"

	// City of arrival for this extra segment.
	ArrivalCity *LocationTypeI_47688C `xml:"arrivalCity,omitempty"`  // minOccurs="0"

	// Airline detail information of the extra segment.
	AirlineDetail *CompanyIdentificationTypeI `xml:"airlineDetail,omitempty"`  // minOccurs="0"

	// Segment detail information.
	SegmentDetail *ProductIdentificationDetailsTypeI `xml:"segmentDetail,omitempty"`  // minOccurs="0"

	// Ticketing status for this segment. Relevant only in case of reply.
	TicketingStatus string `xml:"ticketingStatus,omitempty"`  // minOccurs="0"
}

type UniqueIdDescriptionType struct {

	// Number specifying the ordering information of the item described within a group.
	SequenceNumber *int32 `xml:"sequenceNumber,omitempty"`  // minOccurs="0"
}
