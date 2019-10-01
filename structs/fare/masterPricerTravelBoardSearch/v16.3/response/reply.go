package Fare_MasterPricerTravelBoardSearchResponse_v16_3 // fmptbr163

//import "encoding/xml"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FMPTBR_16_3_1A Fare_MasterPricerTravelBoardSearchReply"`

	// Gives status about reply : type of process, region , CPU etc..
	ReplyStatus *StatusType `xml:"replyStatus,omitempty"`

	ErrorMessage *ErrorMessage `xml:"errorMessage,omitempty"`

	// Specifies the currency used for pricing.
	ConversionRate *ConversionRateTypeI `xml:"conversionRate,omitempty"`

	// Solution Family
	SolutionFamily []*FareInformationType `xml:"solutionFamily,omitempty"` // maxOccurs="20"

	// Details of the fare families processed
	FamilyInformation []*FareFamilyType `xml:"familyInformation,omitempty"` // maxOccurs="200"

	AmountInfoForAllPax *AmountInfoForAllPax `xml:"amountInfoForAllPax,omitempty"`

	AmountInfoPerPax []*AmountInfoPerPax `xml:"amountInfoPerPax,omitempty"` // maxOccurs="20"

	FeeDetails []*FeeDetails `xml:"feeDetails,omitempty"` // maxOccurs="2099"

	// Bucket information
	BucketInfo []*BucketInformationType `xml:"bucketInfo,omitempty"` // maxOccurs="10"

	// Theme identification text
	ThemeIdText []*ThemeText `xml:"themeIdText,omitempty"` // maxOccurs="1000"

	AdditionalInfo []*AdditionalInfo `xml:"additionalInfo,omitempty"` // maxOccurs="1000"

	// Free text identifying an airline in a code share.
	CompanyIdText []*CompanyIdentificationTextType `xml:"companyIdText,omitempty"` // maxOccurs="5000"

	OfficeIdDetails []*OfficeIdDetails `xml:"officeIdDetails,omitempty"` // maxOccurs="20"

	FlightIndex []*FlightIndex `xml:"flightIndex,omitempty"` // maxOccurs="6"

	Recommendation []*Recommendation `xml:"recommendation,omitempty"` // maxOccurs="100000"

	OtherSolutions []*OtherSolutions `xml:"otherSolutions,omitempty"` // maxOccurs="100009"

	WarningInfo []*WarningInfo `xml:"warningInfo,omitempty"` // maxOccurs="9"

	GlobalInformation []*GlobalInformation `xml:"globalInformation,omitempty"` // maxOccurs="9"

	ServiceFeesGrp []*ServiceFeesGrp `xml:"serviceFeesGrp,omitempty"` // maxOccurs="3"

	// Multi dimension references
	MultiDimensionRef []*MultiDimensionalValueType `xml:"multiDimensionRef,omitempty"` // maxOccurs="100000"

	Value []*ValueSearchCriteriaType `xml:"value,omitempty"` // maxOccurs="100009"

	MnrGrp *MnrGrp `xml:"mnrGrp,omitempty"`

	OffersGroup *OffersGroup `xml:"offersGroup,omitempty"`
}

type ErrorMessage struct {
	// Application error details.
	ApplicationError *ApplicationErrorInformationType_78543S `xml:"applicationError"`

	// Type of error message and free text
	ErrorMessageText *InteractiveFreeTextType_78544S `xml:"errorMessageText,omitempty"`
}

type AmountInfoForAllPax struct {
	// Itinerary amounts for all passengers
	ItineraryAmounts *MonetaryInformationType `xml:"itineraryAmounts"`

	AmountsPerSgt []*AmountsPerSgt `xml:"amountsPerSgt,omitempty"` // maxOccurs="9"
}

type AmountsPerSgt struct {
	// Requested segment reference
	SgtRef *ReferenceInfoType_133176S `xml:"sgtRef"`

	// Amounts : Issue total amount, issue taxes amount, non refundable taxes amount
	Amounts *MonetaryInformationType `xml:"amounts,omitempty"`
}

type AmountInfoPerPax struct {
	// Passenger references
	PaxRef *SpecificTravellerType `xml:"paxRef"`

	// Passenger attributes : Infant indicator
	PaxAttributes *FareInformationType_80868S `xml:"paxAttributes,omitempty"`

	// Itinerary amounts information
	ItineraryAmounts *MonetaryInformationType `xml:"itineraryAmounts"`

	AmountsPerSgt []*AmountsPerSgt `xml:"amountsPerSgt,omitempty"` // maxOccurs="9"
}

type FeeDetails struct {
	// Fee/Reduction Reference number.
	FeeReference *ItemReferencesAndVersionsType_78564S `xml:"feeReference"`

	// Fee/Reduction information.
	FeeInformation *DiscountAndPenaltyInformationType `xml:"feeInformation,omitempty"`

	// Fee/Reduction parameters.
	FeeParameters *AttributeType_78561S `xml:"feeParameters,omitempty"`

	// To specify conversion rate details
	ConvertedOrOriginalInfo *ConversionRateTypeI_78562S `xml:"convertedOrOriginalInfo,omitempty"`
}

type AdditionalInfo struct {
	// Identifier
	IdentDetails *ProductIdentificationType `xml:"identDetails"`

	// Date information
	DateInfo *DateTimePeriodType `xml:"dateInfo,omitempty"`
}

type OfficeIdDetails struct {
	// Office Id Information
	OfficeIdInformation *UserIdentificationType `xml:"officeIdInformation"`

	// Office Id Reference Number
	OfficeIdReference *ItemReferencesAndVersionsType_78536S `xml:"officeIdReference"`
}

type FlightIndex struct {
	// Indicates references and details about requested segments
	RequestedSegmentRef *OriginAndDestinationRequestType `xml:"requestedSegmentRef"`

	GroupOfFlights []*GroupOfFlights `xml:"groupOfFlights"` // maxOccurs="100000"
}

type GroupOfFlights struct {
	// To indicate parameters for proposed flight group.
	PropFlightGrDetail *ProposedSegmentType `xml:"propFlightGrDetail"`

	FlightDetails []*FlightDetails `xml:"flightDetails"` // maxOccurs="4"
}

type FlightDetails struct {
	// Specification of details on the flight and posting availability
	FlightInformation *TravelProductType `xml:"flightInformation"`

	// returns booking class and availability context
	AvlInfo []*FlightProductInformationType_141442S `xml:"avlInfo,omitempty"` // maxOccurs="6"

	// Details on Flight date, time and location of technical stop or change of gauge
	TechnicalStop []*DateAndTimeInformationType `xml:"technicalStop,omitempty"` // maxOccurs="5"

	// Code Share Agreement description for current flight.
	CommercialAgreement *CommercialAgreementsType `xml:"commercialAgreement,omitempty"`

	// Additional Info about flight, such as Reference number, and several options
	AddInfo *HeaderInformationTypeI `xml:"addInfo,omitempty"`

	// Flight characteristics
	FlightCharacteristics *FlightCharacteristicsType `xml:"flightCharacteristics,omitempty"`

	// Flight Services by cabin/rbd
	FlightServices []*FlightServicesType `xml:"flightServices,omitempty"` // maxOccurs="9"

	// Meal services
	MealServices *MealServicesType `xml:"mealServices,omitempty"`
}

type Recommendation struct {
	// Specification of the item number
	ItemNumber *ItemNumberType_161497S `xml:"itemNumber"`

	// To describe type of recommendation
	WarningMessage []*InteractiveFreeTextType_78544S `xml:"warningMessage,omitempty"` // maxOccurs="4"

	// Indicates the Fare family reference.
	FareFamilyRef *ReferenceInfoType_133176S `xml:"fareFamilyRef,omitempty"`

	// Recommendation Price and Taxes.
	RecPriceInfo *MonetaryInformationType_193024S `xml:"recPriceInfo"`

	// Mini rules
	MiniRule []*MiniRulesType_78547S `xml:"miniRule,omitempty"` // maxOccurs="9"

	// Indicates reference of Flight or the fee reference valid for all pax (usage:start with the 1 possible Fee reference, then provide the segments references)
	SegmentFlightRef []*ReferenceInfoType `xml:"segmentFlightRef,omitempty"` // maxOccurs="100009"

	RecommandationSegmentsFareDetails []*RecommandationSegmentsFareDetails `xml:"recommandationSegmentsFareDetails,omitempty"` // maxOccurs="6"

	PaxFareProduct []*PaxFareProduct `xml:"paxFareProduct"` // maxOccurs="10"

	SpecificRecDetails []*SpecificRecDetails `xml:"specificRecDetails,omitempty"` // maxOccurs="100000"
}

type RecommandationSegmentsFareDetails struct {
	// Reference and details about requested segments.
	RecommendationSegRef *OriginAndDestinationRequestType `xml:"recommendationSegRef"`

	// Amounts per requested segment.
	SegmentMonetaryInformation *MonetaryInformationType `xml:"segmentMonetaryInformation,omitempty"`
}

type PaxFareProduct struct {
	// Passenger Fare Details.
	PaxFareDetail *PricingTicketingSubsequentType_193023S `xml:"paxFareDetail"`

	// Indicates Fee references (usage: start with the 1 possible Fee reference, then provide the segments references)
	FeeRef *ReferenceInfoType_134839S `xml:"feeRef,omitempty"`

	// Passenger Reference
	PaxReference []*TravellerReferenceInformationType `xml:"paxReference"` // maxOccurs="6"

	// add tax details for each passenger of each recommendations.
	PassengerTaxDetails *TaxType `xml:"passengerTaxDetails,omitempty"`

	Fare []*Fare `xml:"fare,omitempty"` // maxOccurs="7"

	FareDetails []*FareDetails `xml:"fareDetails"` // maxOccurs="6"
}

type Fare struct {
	// Last Date to Ticket, Penalties
	PricingMessage *InteractiveFreeTextType_78559S `xml:"pricingMessage"`

	// Amount of penalties, Surcharges...
	MonetaryInformation *MonetaryInformationType_199534S `xml:"monetaryInformation,omitempty"`
}

type FareDetails struct {
	// Reference of the Requested Segment
	SegmentRef *OriginAndDestinationRequestType `xml:"segmentRef"`

	GroupOfFares []*GroupOfFares `xml:"groupOfFares,omitempty"` // maxOccurs="4"

	// Amounts per passenger per requested segment.
	PsgSegMonetaryInformation *MonetaryInformationType `xml:"psgSegMonetaryInformation,omitempty"`

	// Majority Cabin Info
	MajCabin []*ProductInformationType `xml:"majCabin,omitempty"` // maxOccurs="10"
}

type GroupOfFares struct {
	// Contains details of Flight and Fare
	ProductInformation *FlightProductInformationType_176659S `xml:"productInformation"`

	// Fare calculation code details
	FareCalculationCodeDetails []*FareCalculationCodeDetailsType `xml:"fareCalculationCodeDetails,omitempty"` // maxOccurs="9"

	// Ticket designator, ticket code and fare basis.
	TicketInfos *FareQualifierDetailsType `xml:"ticketInfos,omitempty"`

	// Reference of Fare Family for each Fare Component
	FareFamiliesRef *ReferenceInfoType_176658S `xml:"fareFamiliesRef,omitempty"`
}

type SpecificRecDetails struct {
	// Recommendation details
	SpecificRecItem *ItemReferencesAndVersionsType `xml:"specificRecItem"`

	SpecificProductDetails []*SpecificProductDetails `xml:"specificProductDetails,omitempty"` // maxOccurs="10"
}

type SpecificProductDetails struct {
	// Product details
	ProductReferences *PricingTicketingSubsequentType `xml:"productReferences"`

	FareContextDetails []*FareContextDetails `xml:"fareContextDetails,omitempty"` // maxOccurs="6"
}

type FareContextDetails struct {
	// Reference of requested segment
	RequestedSegmentInfo *OriginAndDestinationRequestType_134833S `xml:"requestedSegmentInfo"`

	CnxContextDetails []*CnxContextDetails `xml:"cnxContextDetails,omitempty"` // maxOccurs="4"
}

type CnxContextDetails struct {
	// Fare connection context details
	FareCnxInfo *FlightProductInformationType `xml:"fareCnxInfo"`
}

type OtherSolutions struct {
	// Reference to the current solution
	Reference *SequenceDetailsTypeU `xml:"reference"`

	AmtGroup []*AmtGroup `xml:"amtGroup,omitempty"` // maxOccurs="10"

	PsgInfo []*PsgInfo `xml:"psgInfo,omitempty"` // maxOccurs="99"
}

type AmtGroup struct {
	// reference to the current amount (per bound, per segment...)
	Ref *ReferenceInfoType_165972S `xml:"ref"`

	// Amount Description
	Amount *MonetaryInformationTypeI `xml:"amount,omitempty"`
}

type PsgInfo struct {
	// passenger reference
	Ref *SegmentRepetitionControlTypeI `xml:"ref"`

	// Passenger Description Info
	Description *FareInformationTypeI `xml:"description,omitempty"`

	// Passenger frequent traveler info
	FreqTraveller *FrequentTravellerIdentificationCodeType `xml:"freqTraveller,omitempty"`

	// amount per passenger or group of passenger
	Amount *MonetaryInformationTypeI `xml:"amount,omitempty"`

	// Fare description
	Fare *FlightProductInformationType_161491S `xml:"fare,omitempty"`

	// Additional Information
	Attribute []*AttributeTypeU `xml:"attribute,omitempty"` // maxOccurs="10"
}

type WarningInfo struct {
	// Dummy Segment
	GlobalMessageMarker *DummySegmentTypeI `xml:"globalMessageMarker"`

	// Informative free text information
	GlobalMessage *InteractiveFreeTextType_78534S `xml:"globalMessage"`
}

type GlobalInformation struct {
	// Coded attributes
	Attributes *CodedAttributeType_78535S `xml:"attributes"`
}

type ServiceFeesGrp struct {
	// Service fee type (OC)
	ServiceTypeInfo *SelectionDetailsType `xml:"serviceTypeInfo"`

	ServiceFeeRefGrp []*ServiceFeeRefGrp `xml:"serviceFeeRefGrp,omitempty"` // maxOccurs="100000"

	ServiceCoverageInfoGrp []*ServiceCoverageInfoGrp `xml:"serviceCoverageInfoGrp,omitempty"` // maxOccurs="100000"

	// Globalmessage marker
	GlobalMessageMarker *DummySegmentTypeI `xml:"globalMessageMarker"`

	ServiceFeeInfoGrp []*ServiceFeeInfoGrp `xml:"serviceFeeInfoGrp,omitempty"` // maxOccurs="100000"

	ServiceDetailsGrp []*ServiceDetailsGrp1 `xml:"serviceDetailsGrp,omitempty"` // maxOccurs="200"

	FreeBagAllowanceGrp []*FreeBagAllowanceGrp `xml:"freeBagAllowanceGrp,omitempty"` // maxOccurs="100000"
}

type ServiceFeeRefGrp struct {
	// Reference of service fee global information
	RefInfo *ReferenceInfoType `xml:"refInfo"`
}

type ServiceCoverageInfoGrp struct {
	// Item reference number for service coverage details
	ItemNumberInfo *ItemNumberType `xml:"itemNumberInfo"`

	ServiceCovInfoGrp []*ServiceCovInfoGrp `xml:"serviceCovInfoGrp,omitempty"` // maxOccurs="200"
}

type ServiceCovInfoGrp struct {
	// Passenger reference number
	PaxRefInfo *SpecificTravellerType `xml:"paxRefInfo"`

	// Service coverage information at flight level Matched seat characteristics
	CoveragePerFlightsInfo []*ActionDetailsType `xml:"coveragePerFlightsInfo,omitempty"` // maxOccurs="6"

	// Carrier information
	CarrierInfo *TransportIdentifierType `xml:"carrierInfo,omitempty"`

	// Service reference number
	RefInfo *ReferenceInfoType_134840S `xml:"refInfo,omitempty"`
}

type ServiceFeeInfoGrp struct {
	// Item number details
	ItemNumberInfo *ItemNumberType `xml:"itemNumberInfo"`

	ServiceDetailsGrp []*ServiceDetailsGrp `xml:"serviceDetailsGrp,omitempty"` // maxOccurs="200"
}

type ServiceDetailsGrp struct {
	// Service reference number
	RefInfo *ReferenceInfoType_134840S `xml:"refInfo"`

	ServiceMatchedInfoGroup []*ServiceMatchedInfoGroup `xml:"serviceMatchedInfoGroup,omitempty"` // maxOccurs="99"
}

type ServiceMatchedInfoGroup struct {
	// Reference on pax number
	PaxRefInfo *SpecificTravellerType `xml:"paxRefInfo"`

	// Pricing oriented service matched information
	PricingInfo *FareInformationType_80868S `xml:"pricingInfo,omitempty"`

	// Informative Service amount
	AmountInfo *MonetaryInformationType_193024S `xml:"amountInfo,omitempty"`

	// Attribute information
	AttributeInfo []*CodedAttributeType `xml:"attributeInfo,omitempty"` // maxOccurs="10"
}

type ServiceDetailsGrp1 struct {
	// Service sub-code and options (exclusion,inclusion, mode pushed,polled)
	ServiceOptionInfo *SpecificDataInformationType `xml:"serviceOptionInfo"`

	FeeDescriptionGrp *FeeDescriptionGrp `xml:"feeDescriptionGrp,omitempty"`
}

type FeeDescriptionGrp struct {
	// Specification of the item number
	ItemNumberInfo *ItemNumberType_80866S `xml:"itemNumberInfo"`

	// Attributes  (SSR code EMD, RFIC, SSIM)
	ServiceAttributesInfo *AttributeType `xml:"serviceAttributesInfo,omitempty"`

	// Other service information (service description, ...)
	ServiceDescriptionInfo *SpecialRequirementsDetailsType `xml:"serviceDescriptionInfo,omitempty"`

	// Commercial name
	CommercialName *InteractiveFreeTextType `xml:"commercialName,omitempty"`
}

type FreeBagAllowanceGrp struct {
	// Free baggage allownce information
	FreeBagAllownceInfo *ExcessBaggageType `xml:"freeBagAllownceInfo"`

	// Item number information
	ItemNumberInfo *ItemNumberType_166130S `xml:"itemNumberInfo,omitempty"`
}

type MnrGrp struct {
	Mnr *MiniRulesType `xml:"mnr"`

	MnrDetails []*MnrDetails `xml:"mnrDetails,omitempty"` // maxOccurs="999"
}

type MnrDetails struct {
	MnrRef *ItemNumberType_176648S `xml:"mnrRef"`

	DateInfo []*DateAndTimeInformationType_182345S `xml:"dateInfo,omitempty"` // maxOccurs="16"

	CatGrp []*CatGrp `xml:"catGrp,omitempty"` // maxOccurs="5"
}

type CatGrp struct {
	// Category information
	CatInfo *CategDescrType `xml:"catInfo"`

	// Monetary information
	MonInfo *MonetaryInformationType_174241S `xml:"monInfo,omitempty"`

	// Status information
	StatusInfo *StatusType_182386S `xml:"statusInfo,omitempty"`
}

type OffersGroup struct {
	// Response information
	Response *ApplicationType `xml:"response"`

	Offers []*Offers `xml:"offers,omitempty"` // maxOccurs="100000"
}

type Offers struct {
	// Offer details
	OfferDtetails *OfferType `xml:"offerDtetails"`

	OfferItems []*OfferItems `xml:"offerItems,omitempty"` // maxOccurs="1000"
}

type OfferItems struct {
	// Offer item details
	OfferItemDetails *OfferItemType `xml:"offerItemDetails"`

	// References
	References []*ReferenceInfoType_165972S `xml:"references,omitempty"` // maxOccurs="10"
}

//
// Complex structs
//

type ActionDetailsType struct {
	// Number of items details
	NumberOfItemsDetails *ProcessingInformationType `xml:"numberOfItemsDetails,omitempty"`

	// Range of segments
	LastItemsDetails []*ReferenceType `xml:"lastItemsDetails,omitempty"` // maxOccurs="99"
}

type AdditionalFareQualifierDetailsTypeI struct {
	// Rate class
	RateClass string `xml:"rateClass,omitempty"`

	// Ticket designator.
	TicketDesignator string `xml:"ticketDesignator,omitempty"`

	// Pricing group
	PricingGroup string `xml:"pricingGroup,omitempty"`

	// Second rate class
	SecondRateClass []string `xml:"secondRateClass,omitempty"` // maxOccurs="29"
}

type AdditionalProductDetailsType struct {
	// Type of aircraft
	EquipmentType string `xml:"equipmentType,omitempty"`

	// Day number of the week
	OperatingDay string `xml:"operatingDay,omitempty"`

	// Number of stops made in a journey if different from 0
	TechStopNumber *int32 `xml:"techStopNumber,omitempty"`

	// Location places of the stops
	LocationId []string `xml:"locationId,omitempty"` // maxOccurs="3"
}

type ApplicationErrorInformationType struct {
	// The code assigned by the receiver of a message for identification of a data validation error condition.
	Error string `xml:"error"`
}

type ApplicationErrorInformationType_78543S struct {
	// Details on application error.
	ApplicationErrorDetail *ApplicationErrorInformationType `xml:"applicationErrorDetail"`
}

type ApplicationIdentificationType struct {
	// Unique identifier of the item.
	InternalId string `xml:"internalId,omitempty"`
}

type ApplicationType struct {
	// Application details
	ApplicationDetails *ApplicationIdentificationType `xml:"applicationDetails,omitempty"`
}

type AttributeInformationType struct {
	// Type of parameter.
	FeeParameterType string `xml:"feeParameterType,omitempty"`

	// Reference to company Id.
	FeeParameterDescription string `xml:"feeParameterDescription,omitempty"`
}

type AttributeInformationTypeU struct {
	// Attribute type
	AttributeType string `xml:"attributeType"`

	// Attribute description
	AttributeDescription string `xml:"attributeDescription,omitempty"`
}

type AttributeInformationType_97181C struct {
	// Attribute type
	AttributeType string `xml:"attributeType"`

	// Attribute description
	AttributeDescription string `xml:"attributeDescription,omitempty"`
}

type AttributeType struct {
	// Criteria Set Type
	AttributeQualifier string `xml:"attributeQualifier,omitempty"`

	// Criteria details
	AttributeDetails []*AttributeInformationType_97181C `xml:"attributeDetails"` // maxOccurs="99"
}

type AttributeTypeU struct {
	// provides the function of the attribute
	AttributeFunction string `xml:"attributeFunction,omitempty"`

	// provides details for the Attribute
	AttributeDetails *AttributeInformationTypeU `xml:"attributeDetails"`
}

type AttributeType_78561S struct {
	// Fee/reduction parameters.
	FeeParameter []*AttributeInformationType `xml:"feeParameter,omitempty"` // maxOccurs="20"
}

type BaggageDetailsType struct {
	// Number of pieces or weight
	FreeAllowance *int32 `xml:"freeAllowance,omitempty"`

	// Nature of the free allowance ( Number of pieces or weight)
	QuantityCode string `xml:"quantityCode,omitempty"`

	// Unit qualifier
	UnitQualifier string `xml:"unitQualifier,omitempty"`
}

type BagtagDetailsType struct {
	// Identifier
	Identifier string `xml:"identifier,omitempty"`

	// Number
	Number *int32 `xml:"number,omitempty"`
}

type BucketInformationType struct {
	// Number
	Number string `xml:"number,omitempty"`

	// Name
	Name string `xml:"name,omitempty"`

	// Mode
	Mode string `xml:"mode,omitempty"`
}

type CabinInformationType struct {
	// Identify the features associated to the cabin/class
	Service string `xml:"service"`

	// Cabin code designator
	Cabin []string `xml:"cabin,omitempty"` // maxOccurs="9"
}

type CabinProductDetailsType struct {
	// Reservation booking designator - RBD
	Rbd string `xml:"rbd"`

	// Reservation Booking Modifier
	BookingModifier string `xml:"bookingModifier,omitempty"`

	// Indicates the cabin related to the Booking code
	Cabin string `xml:"cabin,omitempty"`

	// Availibility status : posting level
	AvlStatus string `xml:"avlStatus,omitempty"`
}

type CabinProductDetailsType_195516C struct {
	// Reservation booking designator - RBD
	Rbd string `xml:"rbd,omitempty"`

	// Reservation Booking Modifier
	BookingModifier string `xml:"bookingModifier,omitempty"`

	// Indicates the cabin related to the Booking code
	Cabin string `xml:"cabin,omitempty"`

	// Availibility status : posting level
	AvlStatus string `xml:"avlStatus"`
}

type CabinProductDetailsType_205138C struct {
	// Reservation booking designator - RBD
	Rbd string `xml:"rbd"`

	// Reservation Booking Modifier
	BookingModifier string `xml:"bookingModifier,omitempty"`

	// Indicates the cabin related to the Booking code
	Cabin string `xml:"cabin,omitempty"`

	// Availibility status : posting level
	AvlStatus string `xml:"avlStatus,omitempty"`
}

type CabinProductDetailsType_229142C struct {
	// Reservation booking designator - RBD
	Rbd string `xml:"rbd"`

	// Indicates the cabin related to the Booking code
	Cabin string `xml:"cabin,omitempty"`

	// Availibility status : posting level
	AvlStatus string `xml:"avlStatus,omitempty"`
}

type CategDescrType struct {
	// Category description information
	DescriptionInfo *CategoryDescriptionType `xml:"descriptionInfo"`

	// Category processing indicator
	ProcessIndicator string `xml:"processIndicator,omitempty"`
}

type CategoryDescriptionType struct {
	// Category number from ATPCO naming conventions (C05 for Advance Purchase restrictions, C06 for Minimun stay ...)
	Number int32 `xml:"number"`

	// Category Code (ATPCO component code, e.g ADV for Advance purchase, STP for stopover restrictions, ELG for eligibility restrictions...)
	Code string `xml:"code,omitempty"`
}

type ClassInformationType struct {
	// Identify the features associated to the cabin/class
	Service string `xml:"service"`

	// Class designator
	Rbd []string `xml:"rbd,omitempty"` // maxOccurs="26"
}

type CodedAttributeInformationType struct {
	// Type of fee/reduction
	AttributeType string `xml:"attributeType"`

	// Fee Id Number
	AttributeDescription string `xml:"attributeDescription,omitempty"`
}

type CodedAttributeInformationType_270108C struct {
	// Format limitations: an..5
	AttributeType string `xml:"attributeType"`

	// Attribute description
	AttributeDescription string `xml:"attributeDescription,omitempty"`
}

type CodedAttributeInformationType_283620C struct {
	// Format limitations: an..5
	AttributeType string `xml:"attributeType"`

	// Attribute description
	AttributeDescription string `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	// Attribute function
	AttributeFunction string `xml:"attributeFunction,omitempty"`

	// Attribute details
	AttributeDetails []*CodedAttributeInformationType_283620C `xml:"attributeDetails"` // maxOccurs="10"
}

type CodedAttributeType_78535S struct {
	// Fee/reduction Id
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"` // maxOccurs="9"
}

type CommercialAgreementsType struct {
	// Codeshare Details
	CodeshareDetails *CompanyRoleIdentificationType `xml:"codeshareDetails,omitempty"`

	// Other codeshare details
	OtherCodeshareDetails []*CompanyRoleIdentificationType `xml:"otherCodeshareDetails,omitempty"` // maxOccurs="9"
}

type CompanyIdentificationTextType struct {
	// Company Id Text reference.
	TextRefNumber *int32 `xml:"textRefNumber,omitempty"`

	// Company id free text.
	CompanyText string `xml:"companyText,omitempty"`
}

type CompanyIdentificationType struct {
	// Marketing carrier
	MarketingCarrier string `xml:"marketingCarrier"`

	// Operating carrier
	OperatingCarrier string `xml:"operatingCarrier,omitempty"`

	// airline alliance code
	Alliance string `xml:"alliance,omitempty"`
}

type CompanyIdentificationTypeI struct {
	// Company
	MarketingCompany string `xml:"marketingCompany,omitempty"`

	// Company
	OperatingCompany string `xml:"operatingCompany,omitempty"`

	// Company
	OtherCompany string `xml:"otherCompany,omitempty"`
}

type CompanyRoleIdentificationType struct {
	// Type of code share agreement.
	CodeShareType string `xml:"codeShareType,omitempty"`

	// company identification
	AirlineDesignator string `xml:"airlineDesignator,omitempty"`

	// flight number
	FlightNumber *int32 `xml:"flightNumber,omitempty"`
}

type CompanyRoleIdentificationType_120771C struct {
	// Type of code share agreement.
	TransportStageQualifier string `xml:"transportStageQualifier,omitempty"`

	// company identification
	Company string `xml:"company,omitempty"`
}

type ConversionRateDetailsTypeI struct {
	// Conversion type
	ConversionType string `xml:"conversionType,omitempty"`

	// Currency
	Currency string `xml:"currency,omitempty"`

	// amount
	Amount string `xml:"amount,omitempty"`
}

type ConversionRateDetailsTypeI_179848C struct {
	// Conversion type
	ConversionType string `xml:"conversionType,omitempty"`

	// Currency
	Currency string `xml:"currency"`

	// Conversion rate for pricing
	Rate string `xml:"rate,omitempty"`

	// Converted value amount
	ConvertedAmountLink string `xml:"convertedAmountLink,omitempty"`

	// Applicable ISO country code or Tax designator code.
	TaxQualifier string `xml:"taxQualifier,omitempty"`
}

type ConversionRateTypeI struct {
	// Detail of conversion rate of First Monetary Unit.
	ConversionRateDetail []*ConversionRateDetailsTypeI_179848C `xml:"conversionRateDetail"` // maxOccurs="9"
}

type ConversionRateTypeI_78562S struct {
	// Details of conversion
	ConversionRateDetail []*ConversionRateDetailsTypeI `xml:"conversionRateDetail"` // maxOccurs="9"
}

type CriteriaiDetaislType struct {
	// Format limitations: an..3
	Type string `xml:"type,omitempty"`

	// Format limitations: an..18
	Value string `xml:"value,omitempty"`

	// Attribute
	Attribute string `xml:"attribute,omitempty"`
}

type DataInformationType struct {
	// Ancillary services options
	Indicator string `xml:"indicator,omitempty"`
}

type DataTypeInformationType struct {
	// service group/sub-group/sub-code information
	SubType string `xml:"subType"`

	// Status (automated, manually added, exempted). Default is automated
	Option string `xml:"option,omitempty"`
}

type DateAndTimeDetailsType struct {
	// Date time period qualifier
	DateQualifier string `xml:"dateQualifier,omitempty"`

	// First Date
	Date string `xml:"date,omitempty"`

	// First Time
	FirstTime string `xml:"firstTime,omitempty"`

	// Movement type.
	EquipementType string `xml:"equipementType,omitempty"`

	// Place/location identification.
	LocationId string `xml:"locationId,omitempty"`
}

type DateAndTimeDetailsType_256192C struct {
	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"`

	// Format limitations: an..35
	Date string `xml:"date,omitempty"`

	// Time
	Time string `xml:"time,omitempty"`

	// Location
	Location string `xml:"location,omitempty"`
}

type DateAndTimeInformationType struct {
	// Details on date and time
	StopDetails []*DateAndTimeDetailsType `xml:"stopDetails"` // maxOccurs="2"
}

type DateAndTimeInformationType_182345S struct {
	// DATE AND TIME DETAILS
	DateAndTimeDetails []*DateAndTimeDetailsType_256192C `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="400"
}

type DateTimePeriodDetailsBatchType struct {
	// Date time qualifier
	DateTimeQualifier string `xml:"dateTimeQualifier,omitempty"`

	// Date time details
	DateTimeDetails string `xml:"dateTimeDetails,omitempty"`
}

type DateTimePeriodDetailsTypeI struct {
	// Qualifier
	Qualifier string `xml:"qualifier"`

	// Value
	Value string `xml:"value,omitempty"`
}

type DateTimePeriodType struct {
	// Date Time Description
	DateTimeDescription *DateTimePeriodDetailsBatchType `xml:"dateTimeDescription"`
}

type DimensionDetailType struct {
	// Bucket reference
	BucketRef string `xml:"bucketRef,omitempty"`

	// Value reference
	ValueRef string `xml:"valueRef,omitempty"`
}

type DiscountAndPenaltyInformationType struct {
	// Used to specify airline collected fee or agent collected fee.
	FeeIdentification string `xml:"feeIdentification,omitempty"`

	// Used to specify penalty information
	FeeInformation *DiscountPenaltyMonetaryInformationType `xml:"feeInformation,omitempty"`
}

type DiscountPenaltyInformationType struct {
	// Discounted fare,...
	FareQualifier string `xml:"fareQualifier"`

	// Dicount code,...
	RateCategory string `xml:"rateCategory,omitempty"`

	// Amount
	Amount *float64 `xml:"amount,omitempty"`

	// Percentage
	Percentage *float64 `xml:"percentage,omitempty"`
}

type DiscountPenaltyMonetaryInformationType struct {
	// Type of discount/penalty
	FeeType string `xml:"feeType,omitempty"`

	// The amount Type can be a percentage or an amount
	FeeAmountType string `xml:"feeAmountType,omitempty"`

	// specify the value
	FeeAmount *float64 `xml:"feeAmount,omitempty"`

	// Fee currency code.
	FeeCurrency string `xml:"feeCurrency,omitempty"`
}

type DummySegmentTypeI struct {
}

type ExcessBaggageType struct {
	// Baggage details
	BaggageDetails *BaggageDetailsType `xml:"baggageDetails,omitempty"`

	// Free baggage allowance details
	BagTagDetails []*BagtagDetailsType `xml:"bagTagDetails,omitempty"` // maxOccurs="99"
}

type FareCalculationCodeDetailsType struct {
	// Qualifier of the amout or rate
	Qualifier string `xml:"qualifier,omitempty"`

	// Amount
	Amount *float64 `xml:"amount,omitempty"`

	// Location code
	LocationCode string `xml:"locationCode,omitempty"`

	// Other location code
	OtherLocationCode string `xml:"otherLocationCode,omitempty"`

	// Rate
	Rate *float64 `xml:"rate,omitempty"`
}

type FareCategoryCodesTypeI struct {
	// Fare type
	FareType string `xml:"fareType"`

	// Other fare type
	OtherFareType []string `xml:"otherFareType,omitempty"` // maxOccurs="8"
}

type FareDetailsType struct {
	// Passenger Type qualifier
	PassengerTypeQualifier string `xml:"passengerTypeQualifier,omitempty"`
}

type FareDetailsTypeI struct {
	// Qualifier
	Qualifier string `xml:"qualifier,omitempty"`

	// Rate
	Rate *float64 `xml:"rate,omitempty"`

	// Country
	Country string `xml:"country,omitempty"`

	// Fare category
	FareCategory string `xml:"fareCategory,omitempty"`
}

type FareDetailsType_193037C struct {
	// Qualifier
	Qualifier string `xml:"qualifier,omitempty"`

	// Rate
	Rate *int32 `xml:"rate,omitempty"`

	// Country
	Country string `xml:"country,omitempty"`

	// Fare Category
	FareCategory string `xml:"fareCategory,omitempty"`
}

type FareFamilyDetailsType struct {
	// Commercial fare Family Short name
	CommercialFamily string `xml:"commercialFamily"`
}

type FareFamilyType struct {
	// Fare Family Reference Number
	RefNumber int32 `xml:"refNumber"`

	// Fare Family Short Name
	FareFamilyname string `xml:"fareFamilyname,omitempty"`

	// HIERARCHICAL ORDER WITHIN FARE FAMILY
	Hierarchy *int32 `xml:"hierarchy,omitempty"`

	// CABIN USED FOR FARE FAMILY
	Cabin string `xml:"cabin,omitempty"`

	// Indicates Commercial Fare Family Short names
	CommercialFamilyDetails []*FareFamilyDetailsType `xml:"commercialFamilyDetails,omitempty"` // maxOccurs="20"

	// Short description of the fare family
	Description string `xml:"description,omitempty"`

	// Carrier code
	Carrier string `xml:"carrier,omitempty"`

	// Reference to the services details
	Services []*ServicesReferences `xml:"services,omitempty"` // maxOccurs="20"

	// Reservation booking designator
	BookingClass []string `xml:"bookingClass,omitempty"` // maxOccurs="26"
}

type FareInformationType struct {
	// Value Qualifier
	ValueQualifier string `xml:"valueQualifier,omitempty"`

	// Value
	Value *int32 `xml:"value,omitempty"`

	// Fare Details
	FareDetails *FareDetailsType_193037C `xml:"fareDetails,omitempty"`

	// Identity Number
	IdentityNumber string `xml:"identityNumber,omitempty"`

	// Fare Type Grouping
	FareTypeGrouping *FareTypeGroupingInformationType `xml:"fareTypeGrouping,omitempty"`

	// Rate Category
	RateCategory string `xml:"rateCategory,omitempty"`
}

type FareInformationTypeI struct {
	// Value qualifier
	ValueQualifier string `xml:"valueQualifier,omitempty"`

	// Value
	Value *int32 `xml:"value,omitempty"`
}

type FareInformationType_80868S struct {
	// Fare details
	FareDetails *FareDetailsType `xml:"fareDetails,omitempty"`
}

type FareProductDetailsType struct {
	// Fare basis code
	FareBasis string `xml:"fareBasis,omitempty"`
}

type FareProductDetailsType_248552C struct {
	// Fare basis code
	FareBasis string `xml:"fareBasis,omitempty"`

	// PTC priced
	PassengerType string `xml:"passengerType,omitempty"`

	// Type of fare
	FareType []string `xml:"fareType,omitempty"` // maxOccurs="9"
}

type FareQualifierDetailsType struct {
	// Route Code
	MovementType string `xml:"movementType,omitempty"`

	// Fare categories
	FareCategories *FareCategoryCodesTypeI `xml:"fareCategories,omitempty"`

	// Fare details
	FareDetails *FareDetailsTypeI `xml:"fareDetails,omitempty"`

	// Additional fare details
	AdditionalFareDetails *AdditionalFareQualifierDetailsTypeI `xml:"additionalFareDetails,omitempty"`

	// Discount details
	DiscountDetails []*DiscountPenaltyInformationType `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type FareTypeGroupingInformationType struct {
	// Pricing Group
	PricingGroup string `xml:"pricingGroup,omitempty"`
}

type FlightCharacteristicsType struct {
	// On-Time Performance
	OnTimePerformance *OnTimePerformanceType `xml:"onTimePerformance,omitempty"`

	// In flight services
	InFlightSrv []string `xml:"inFlightSrv,omitempty"` // maxOccurs="99"
}

type FlightProductInformationType struct {
	// Indicates flight cabin details
	CabinProduct []*CabinProductDetailsType_195516C `xml:"cabinProduct,omitempty"` // maxOccurs="6"

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType `xml:"contextDetails,omitempty"`
}

type FlightProductInformationType_141442S struct {
	// Indicates flight cabin details
	CabinProduct []*CabinProductDetailsType_205138C `xml:"cabinProduct,omitempty"` // maxOccurs="26"

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType_205137C `xml:"contextDetails,omitempty"`
}

type FlightProductInformationType_161491S struct {
	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType_229142C `xml:"cabinProduct,omitempty"`

	// Fare product details
	FareProductDetail *FareProductDetailsType `xml:"fareProductDetail,omitempty"`
}

type FlightProductInformationType_176659S struct {
	// Indicates flight cabin details
	CabinProduct *CabinProductDetailsType `xml:"cabinProduct,omitempty"`

	// Fare product details
	FareProductDetail *FareProductDetailsType_248552C `xml:"fareProductDetail,omitempty"`

	// Corporate number or name and number
	CorporateId []string `xml:"corporateId,omitempty"` // maxOccurs="2"

	// To determine if Fare Breaks at this segment
	BreakPoint string `xml:"breakPoint,omitempty"`

	// To specify additional characteristics.
	ContextDetails *ProductTypeDetailsType `xml:"contextDetails,omitempty"`
}

type FlightServicesType struct {
	// Type of service used
	ServiceType string `xml:"serviceType"`

	CabinInfo []*CabinInformationType `xml:"cabinInfo,omitempty"` // maxOccurs="99"

	ClassInfo []*ClassInformationType `xml:"classInfo,omitempty"` // maxOccurs="99"
}

type FreeTextQualificationType struct {
	// Type of message
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Coded Text or type of information in 4440 (e.g. type of OSI or free text, canned message value)
	InformationType string `xml:"informationType,omitempty"`
}

type FreeTextQualificationTypeI struct {
	// Text subject qualifier
	TextSubjectQualifier string `xml:"textSubjectQualifier"`
}

type FreeTextQualificationType_120769C struct {
	// Type of message
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Coded Text or type of information in 4440 (e.g. type of OSI or free text, canned message value)
	InformationType string `xml:"informationType,omitempty"`

	// ISO code for language of free text (default is English)
	Language string `xml:"language,omitempty"`
}

type FrequentTravellerIdentificationCodeType struct {
	// Frequent Traveller Info
	FrequentTravellerDetails []*FrequentTravellerIdentificationType `xml:"frequentTravellerDetails"` // maxOccurs="99"
}

type FrequentTravellerIdentificationType struct {
	// Carrier where the FQTV is registered.
	Carrier string `xml:"carrier,omitempty"`

	// Number
	Number string `xml:"number,omitempty"`

	// To specify a Tier linked to the FQTV
	TierLevel string `xml:"tierLevel,omitempty"`

	// For example : priority code
	PriorityCode string `xml:"priorityCode,omitempty"`
}

type HeaderInformationTypeI struct {
	// Status
	Status []string `xml:"status,omitempty"` // maxOccurs="2"

	// Date and Time info
	DateTimePeriodDetails *DateTimePeriodDetailsTypeI `xml:"dateTimePeriodDetails,omitempty"`

	// Reference number
	ReferenceNumber string `xml:"referenceNumber,omitempty"`

	// Contains product identification such as UIC code...
	ProductIdentification []string `xml:"productIdentification,omitempty"` // maxOccurs="2"
}

type InteractiveFreeTextType struct {
	// Free text qualification
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"`

	// Free text
	FreeText string `xml:"freeText,omitempty"`
}

type InteractiveFreeTextType_78534S struct {
	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType `xml:"freeTextQualification,omitempty"`

	// Free text
	Description []string `xml:"description,omitempty"` // maxOccurs="9"
}

type InteractiveFreeTextType_78544S struct {
	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType_120769C `xml:"freeTextQualification,omitempty"`

	// Free text
	Description []string `xml:"description,omitempty"` // maxOccurs="9"
}

type InteractiveFreeTextType_78559S struct {
	// Details on interactive free text
	FreeTextQualification *FreeTextQualificationType_120769C `xml:"freeTextQualification,omitempty"`

	// Free text
	Description []string `xml:"description,omitempty"` // maxOccurs="9"
}

type ItemNumberIdentificationType struct {
	// Ancillary Service number
	Number string `xml:"number,omitempty"`

	// Type
	Type string `xml:"type,omitempty"`

	// Qualifier
	Qualifier string `xml:"qualifier,omitempty"`

	// Responsible agency
	ResponsibleAgency string `xml:"responsibleAgency,omitempty"`
}

type ItemNumberIdentificationType_191597C struct {
	// Item number.
	Number string `xml:"number,omitempty"`

	// Indicates the item type .
	NumberType string `xml:"numberType,omitempty"`
}

type ItemNumberIdentificationType_192331C struct {
	// Service coverage number
	Number string `xml:"number,omitempty"`

	// Type
	Type string `xml:"type,omitempty"`

	// Qualifier
	Qualifier string `xml:"qualifier,omitempty"`

	// Responsible agency
	ResponsibleAgency string `xml:"responsibleAgency,omitempty"`
}

type ItemNumberIdentificationType_234878C struct {
	// Number
	Number *int32 `xml:"number,omitempty"`

	// Type
	Type string `xml:"type,omitempty"`
}

type ItemNumberIdentificationType_248537C struct {
	// Format limitations: an..35
	Number string `xml:"number,omitempty"`
}

type ItemNumberType struct {
	// Item number details
	ItemNumber *ItemNumberIdentificationType_192331C `xml:"itemNumber"`
}

type ItemNumberType_161497S struct {
	// Indicates the recommendation number.
	ItemNumberId *ItemNumberIdentificationType_191597C `xml:"itemNumberId,omitempty"`

	// Code share details.
	CodeShareDetails []*CompanyRoleIdentificationType_120771C `xml:"codeShareDetails,omitempty"` // maxOccurs="6"

	// Pricing ticketind details.
	PriceTicketing *PricingTicketingInformationType `xml:"priceTicketing,omitempty"`
}

type ItemNumberType_166130S struct {
	// Item number details
	ItemNumberDetails []*ItemNumberIdentificationType_234878C `xml:"itemNumberDetails"` // maxOccurs="99"
}

type ItemNumberType_176648S struct {
	ItemNumberDetails []*ItemNumberIdentificationType_248537C `xml:"itemNumberDetails"` // maxOccurs="99"
}

type ItemNumberType_80866S struct {
	// Item number details
	ItemNumberDetails *ItemNumberIdentificationType `xml:"itemNumberDetails"`
}

type ItemReferencesAndVersionsType struct {
	// Qualifies the type of the reference used.
	ReferenceType string `xml:"referenceType,omitempty"`

	// Unique fee reference.
	RefNumber *int32 `xml:"refNumber,omitempty"`
}

type ItemReferencesAndVersionsType_78536S struct {
	// Qualifies the type of the reference used.
	ReferenceType string `xml:"referenceType,omitempty"`

	// Unique fee reference.
	RefNumber *int32 `xml:"refNumber,omitempty"`
}

type ItemReferencesAndVersionsType_78564S struct {
	// Qualifies the type of the reference used.
	ReferenceType string `xml:"referenceType,omitempty"`

	// Unique fee reference.
	FeeRefNumber *int32 `xml:"feeRefNumber,omitempty"`
}

type ItineraryDetailsType struct {
	// Airport/City Qualifier: the passenger wants to depart/arrive from/to the same airport or city in the specified requested segment
	AirportCityQualifier string `xml:"airportCityQualifier"`

	// Requested segment number
	SegmentNumber int32 `xml:"segmentNumber"`
}

type LocationIdentificationDetailsType struct {
	// 3 characters ATA/IATA airport/city code
	LocationId string `xml:"locationId"`

	// Airport/city qualifier: the requested point is an airport when ambiguity exists (e.g. HOU)
	AirportCityQualifier string `xml:"airportCityQualifier,omitempty"`

	// Terminal information
	Terminal string `xml:"terminal,omitempty"`
}

type MealServicesType struct {
	// Service details
	ServiceDetails []*ServiceDetailsType `xml:"serviceDetails,omitempty"` // maxOccurs="26"
}

type MiniRulesDetailsType struct {
	// Coded text (period or day)
	Interpretation string `xml:"interpretation,omitempty"`

	// Data type coded or value of interpretation
	Value []string `xml:"value,omitempty"` // maxOccurs="10"
}

type MiniRulesIndicatorType struct {
	// See rule indicator and free form text indicator
	RuleIndicator []string `xml:"ruleIndicator,omitempty"` // maxOccurs="2"
}

type MiniRulesType struct {
	// Categoty of restriction: PTC, Max Adv Pur, Days, ...
	Category string `xml:"category"`
}

type MiniRulesType_78547S struct {
	// Type of restriction: PTC, Max Adv Res, Max Ticketing After Res, ...
	RestrictionType string `xml:"restrictionType,omitempty"`

	// Categoty of restriction: PTC, Max Adv Pur, Days, ...
	Category string `xml:"category"`

	// Indicators
	Indicator *MiniRulesIndicatorType `xml:"indicator,omitempty"`

	// Mini rules
	MiniRules []*MiniRulesDetailsType `xml:"miniRules,omitempty"` // maxOccurs="5"
}

type MonetaryInformationDetailsType struct {
	// To specify amount and percentage.
	AmountType string `xml:"amountType,omitempty"`

	// Amount
	Amount float64 `xml:"amount"`

	// ISO currency code
	Currency string `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsTypeI struct {
	// type Qualifier
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	Amount string `xml:"amount,omitempty"`

	// Currency
	Currency string `xml:"currency,omitempty"`
}

type MonetaryInformationDetailsType_245528C struct {
	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount
	Amount *float64 `xml:"amount,omitempty"`

	// Currency
	Currency string `xml:"currency,omitempty"`

	// location
	Location string `xml:"location,omitempty"`
}

type MonetaryInformationType struct {
	// Monetary information.
	MonetaryDetail []*MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"` // maxOccurs="20"
}

type MonetaryInformationTypeI struct {
	// Monetary details
	MonetaryDetails []*MonetaryInformationDetailsTypeI `xml:"monetaryDetails"` // maxOccurs="99"
}

type MonetaryInformationType_174241S struct {
	MonetaryDetails *MonetaryInformationDetailsType_245528C `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsType_245528C `xml:"otherMonetaryDetails,omitempty"` // maxOccurs="19"
}

type MonetaryInformationType_193024S struct {
	// Monetary information.
	MonetaryDetail []*MonetaryInformationDetailsType `xml:"monetaryDetail,omitempty"` // maxOccurs="30"
}

type MonetaryInformationType_199534S struct {
	// Monetary information
	MonetaryDetail []*MonetaryInformationDetailsType `xml:"monetaryDetail"` // maxOccurs="2"
}

type MultiDimensionalValueType struct {
	// Identifier
	Identifier string `xml:"identifier,omitempty"`

	// Dimension detail
	DimensionDetail []*DimensionDetailType `xml:"dimensionDetail,omitempty"` // maxOccurs="10"
}

type OfferItemType struct {
	// Offer item identifier
	OfferItemId string `xml:"offerItemId,omitempty"`

	// Status
	Status string `xml:"status,omitempty"`
}

type OfferType struct {
	// Reference
	Reference string `xml:"reference,omitempty"`

	// Offer identification
	OfferId string `xml:"offerId,omitempty"`
}

type OnTimePerformanceType struct {
	// Date time period
	DateTimePeriod string `xml:"dateTimePeriod,omitempty"`

	// Percentage
	Percentage *float64 `xml:"percentage,omitempty"`

	// Accuracy
	Accuracy string `xml:"accuracy,omitempty"`
}

type OriginAndDestinationRequestType struct {
	// Requested segment number
	SegRef int32 `xml:"segRef"`

	// Forces arrival or departure, from/to the same airport/city
	LocationForcing []*ItineraryDetailsType `xml:"locationForcing,omitempty"` // maxOccurs="2"
}

type OriginAndDestinationRequestType_134833S struct {
	// Requested segment number
	SegRef int32 `xml:"segRef"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// Office Name.
	OfficeName *int32 `xml:"officeName,omitempty"`

	// Agent Sign In .
	AgentSignin string `xml:"agentSignin,omitempty"`

	// Confidential Office Name.
	ConfidentialOffice string `xml:"confidentialOffice,omitempty"`

	// Other Office Name
	OtherOffice string `xml:"otherOffice,omitempty"`
}

type PricingTicketingInformationType struct {
	// Price type qualifier
	PriceType []string `xml:"priceType"` // maxOccurs="20"
}

type PricingTicketingSubsequentType struct {
	// Passenger fare product number
	PaxFareNum []string `xml:"paxFareNum"` // maxOccurs="10"
}

type PricingTicketingSubsequentType_193023S struct {
	// Passenger fare product number
	PaxFareNum string `xml:"paxFareNum"`

	// Total fare amount
	TotalFareAmount float64 `xml:"totalFareAmount"`

	// Total tax amount
	TotalTaxAmount *float64 `xml:"totalTaxAmount,omitempty"`

	// Code share details.
	CodeShareDetails []*CompanyRoleIdentificationType_120771C `xml:"codeShareDetails,omitempty"` // maxOccurs="6"

	// Monetary information.
	MonetaryDetails []*MonetaryInformationDetailsType `xml:"monetaryDetails,omitempty"` // maxOccurs="20"

	// Pricing ticketing details.
	PricingTicketing *PricingTicketingInformationType `xml:"pricingTicketing,omitempty"`
}

type ProcessingInformationType struct {
	// Action qualifier
	ActionQualifier string `xml:"actionQualifier,omitempty"`

	// Reference qualifier
	ReferenceQualifier string `xml:"referenceQualifier,omitempty"`

	// Reference number
	RefNum string `xml:"refNum,omitempty"`
}

type ProductDateTimeType struct {
	// Departure date
	DateOfDeparture string `xml:"dateOfDeparture"`

	// Departure time
	TimeOfDeparture string `xml:"timeOfDeparture,omitempty"`

	// Arrival date
	DateOfArrival string `xml:"dateOfArrival,omitempty"`

	// Arrival time
	TimeOfArrival string `xml:"timeOfArrival,omitempty"`

	// Arrival date compared to departure date, only if different from 0
	DateVariation *int32 `xml:"dateVariation,omitempty"`
}

type ProductDetailsType struct {
	// Format limitations: an..17
	Designator string `xml:"designator"`

	// Format limitations: an..3
	AvailabilityStatus string `xml:"availabilityStatus,omitempty"`

	// Format limitations: an..3
	SpecialService string `xml:"specialService,omitempty"`

	// Format limitations: an..7
	Option []string `xml:"option,omitempty"` // maxOccurs="9"
}

type ProductFacilitiesType struct {
	// Yes-No indicator whether Last Seat Available
	LastSeatAvailable string `xml:"lastSeatAvailable,omitempty"`

	// Level of access
	LevelOfAccess string `xml:"levelOfAccess,omitempty"`

	// Yes-No indicator whether electronic ticketing
	ElectronicTicketing string `xml:"electronicTicketing,omitempty"`

	// Product identification suffix
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`

	// Define whether a flight has been polled or not
	ProductDetailQualifier string `xml:"productDetailQualifier,omitempty"`

	// Add some flight restrictions (See code set list)
	FlightCharacteristic []string `xml:"flightCharacteristic,omitempty"` // maxOccurs="9"
}

type ProductIdentDetailsType struct {
	// Number
	Number string `xml:"number,omitempty"`
}

type ProductIdentificationType struct {
	// Product identification details
	ProductData []*ProductIdentDetailsType `xml:"productData,omitempty"` // maxOccurs="9"
}

type ProductInformationType struct {
	// value of the Qualifier: INT for International DOM for Domestic EUR for European  otherwise CM#10569 INVALID INTERNATIONAL INDICATOR is returned.
	ProductDetailsQualifier string `xml:"productDetailsQualifier,omitempty"`

	BookingClassDetails []*ProductDetailsType `xml:"bookingClassDetails,omitempty"` // maxOccurs="26"
}

type ProductTypeDetailsType struct {
	// Availability connection type.
	AvailabilityCnxType []string `xml:"availabilityCnxType"` // maxOccurs="9"
}

type ProductTypeDetailsType_205137C struct {
	// indicates whether the flight is domestic or international
	Avl []string `xml:"avl"` // maxOccurs="9"
}

type ProposedSegmentDetailsType struct {
	// Flight proposal reference
	Ref string `xml:"ref,omitempty"`

	// Elapse Flying Time
	UnitQualifier string `xml:"unitQualifier,omitempty"`
}

type ProposedSegmentType struct {
	// Parameters for proposed flight group
	FlightProposal []*ProposedSegmentDetailsType `xml:"flightProposal"` // maxOccurs="9"

	// Flight characteristics.
	FlightCharacteristic string `xml:"flightCharacteristic,omitempty"`

	// Majority cabin
	MajCabin string `xml:"majCabin,omitempty"`
}

type ReferenceInfoType struct {
	// Referencing details
	ReferencingDetail []*ReferencingDetailsType_191583C `xml:"referencingDetail,omitempty"` // maxOccurs="200"
}

type ReferenceInfoType_133176S struct {
	// Referencing details
	ReferencingDetail []*ReferencingDetailsType `xml:"referencingDetail,omitempty"` // maxOccurs="99"
}

type ReferenceInfoType_134839S struct {
	// Referencing details
	ReferencingDetail []*ReferencingDetailsType_195561C `xml:"referencingDetail,omitempty"` // maxOccurs="99"
}

type ReferenceInfoType_134840S struct {
	// Referencing details
	ReferencingDetail []*ReferencingDetailsType_195561C `xml:"referencingDetail,omitempty"` // maxOccurs="200"
}

type ReferenceInfoType_165972S struct {
	// Reference details
	ReferenceDetails []*ReferencingDetailsType_234704C `xml:"referenceDetails,omitempty"` // maxOccurs="20"
}

type ReferenceInfoType_176658S struct {
	// Referencing details
	ReferencingDetail []*ReferencingDetailsType `xml:"referencingDetail,omitempty"` // maxOccurs="6"
}

type ReferenceType struct {
	// Reference  of leg
	RefOfLeg string `xml:"refOfLeg,omitempty"`

	// Reference of segment starting range
	FirstItemIdentifier *int32 `xml:"firstItemIdentifier,omitempty"`

	// Reference of segment ending range
	LastItemIdentifier *int32 `xml:"lastItemIdentifier,omitempty"`
}

type ReferencingDetailsType struct {
	// Reference qualifier
	RefQualifier string `xml:"refQualifier,omitempty"`

	// Requested segment reference
	RefNumber int32 `xml:"refNumber"`
}

type ReferencingDetailsType_191583C struct {
	// Service reference qualifier
	RefQualifier string `xml:"refQualifier,omitempty"`

	// Service reference
	RefNumber int32 `xml:"refNumber"`
}

type ReferencingDetailsType_195561C struct {
	// Segment reference qualifier
	RefQualifier string `xml:"refQualifier,omitempty"`

	// Flight or flight group reference
	RefNumber int32 `xml:"refNumber"`
}

type ReferencingDetailsType_234704C struct {
	// Type
	Type string `xml:"type,omitempty"`

	// Value
	Value string `xml:"value,omitempty"`
}

type SegmentRepetitionControlDetailsTypeI struct {
	// traveller number
	Quantity *int32 `xml:"quantity,omitempty"`

	// range of traveller
	NumberOfUnits *int32 `xml:"numberOfUnits,omitempty"`
}

type SegmentRepetitionControlTypeI struct {
	// Segment control details
	SegmentControlDetails []*SegmentRepetitionControlDetailsTypeI `xml:"segmentControlDetails,omitempty"` // maxOccurs="9"
}

type SelectionDetailsInformationType struct {
	// Carrier fee type
	Type string `xml:"type"`

	// Carrier fee status
	OptionInformation string `xml:"optionInformation,omitempty"`
}

type SelectionDetailsType struct {
	// Carrier fees options
	CarrierFeeDetails *SelectionDetailsInformationType `xml:"carrierFeeDetails"`
}

type SequenceDetailsTypeU struct {
	// Sequence details
	SequenceDetails *SequenceInformationTypeU `xml:"sequenceDetails,omitempty"`
}

type SequenceInformationTypeU struct {
	// Number
	Number string `xml:"number"`

	// Identification code
	IdentificationCode string `xml:"identificationCode,omitempty"`
}

type ServiceDetailsType struct {
	// Reservation booking designator
	BookingClass string `xml:"bookingClass,omitempty"`

	// Service
	Service []string `xml:"service,omitempty"` // maxOccurs="20"
}

type ServicesReferences struct {
	// Reference of the service
	Reference string `xml:"reference,omitempty"`

	// Status of the service
	Status string `xml:"status,omitempty"`

	// Service lowest price
	FromPrice string `xml:"fromPrice,omitempty"`
}

type SpecialRequirementsDataDetailsType struct {
	// SSR seat characteristic
	SeatCharacteristics []string `xml:"seatCharacteristics,omitempty"` // maxOccurs="5"
}

type SpecialRequirementsDetailsType struct {
	// To specify the Service Requirement of the customer
	ServiceRequirementsInfo *SpecialRequirementsTypeDetailsType `xml:"serviceRequirementsInfo"`

	// Seat details
	SeatDetails []*SpecialRequirementsDataDetailsType `xml:"seatDetails,omitempty"` // maxOccurs="999"
}

type SpecialRequirementsTypeDetailsType struct {
	// To specify the Service Classification of the Service Requirement.
	ServiceClassification string `xml:"serviceClassification"`

	// Status
	ServiceStatus string `xml:"serviceStatus,omitempty"`

	// To specify the number of items involved
	ServiceNumberOfInstances *int32 `xml:"serviceNumberOfInstances,omitempty"`

	// To specify to which marketing carrier the service applies
	ServiceMarketingCarrier string `xml:"serviceMarketingCarrier,omitempty"`

	// Specify the Service group
	ServiceGroup string `xml:"serviceGroup,omitempty"`

	// Specify the Service Sub-Group
	ServiceSubGroup string `xml:"serviceSubGroup,omitempty"`

	// Free Text attached to the Service.
	ServiceFreeText []string `xml:"serviceFreeText,omitempty"` // maxOccurs="99"
}

type SpecificDataInformationType struct {
	// Carrier fee description
	DataTypeInformation *DataTypeInformationType `xml:"dataTypeInformation"`

	// Data information
	DataInformation []*DataInformationType `xml:"dataInformation,omitempty"` // maxOccurs="99"
}

type SpecificTravellerDetailsType struct {
	// Reference number
	ReferenceNumber string `xml:"referenceNumber,omitempty"`
}

type SpecificTravellerType struct {
	// Traveller details
	TravellerDetails []*SpecificTravellerDetailsType `xml:"travellerDetails,omitempty"` // maxOccurs="99"
}

type StatusDetailsType struct {
	// Advisory type information, Fare Server
	AdvisoryTypeInfo string `xml:"advisoryTypeInfo,omitempty"`

	// CPU time, user type
	Notification string `xml:"notification,omitempty"`

	// CPU time,user type
	Notification2 string `xml:"notification2,omitempty"`

	// Capture and trace information
	Description string `xml:"description,omitempty"`
}

type StatusDetailsType_256255C struct {
	// list of status/qualifiers Either His for Historical or     Crt for Current
	Indicator string `xml:"indicator,omitempty"`

	// Format limitations: an..3
	Action string `xml:"action,omitempty"`
}

type StatusType struct {
	// Status details
	Status []*StatusDetailsType `xml:"status"` // maxOccurs="10"
}

type StatusType_182386S struct {
	// STATUS DETAILS
	StatusInformation []*StatusDetailsType_256255C `xml:"statusInformation"` // maxOccurs="99"
}

type TaxDetailsType struct {
	// Amount
	Rate string `xml:"rate,omitempty"`

	// Country code
	CountryCode string `xml:"countryCode,omitempty"`

	// Currency code
	CurrencyCode string `xml:"currencyCode,omitempty"`

	// Type
	Type string `xml:"type,omitempty"`

	// Indicator
	Indicator []string `xml:"indicator,omitempty"` // maxOccurs="98"
}

type TaxType struct {
	// Tax category
	TaxCategory string `xml:"taxCategory,omitempty"`

	// Tax details
	TaxDetails []*TaxDetailsType `xml:"taxDetails,omitempty"` // maxOccurs="99"
}

type ThemeText struct {
	// Reference
	Reference string `xml:"reference,omitempty"`

	// Text
	Text string `xml:"text,omitempty"`
}

type TransportIdentifierType struct {
	// Company identification
	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"`
}

type TravelProductType struct {
	// Date and time of departure and arrival
	ProductDateTime *ProductDateTimeType `xml:"productDateTime"`

	// Location of departure and arrival
	Location []*LocationIdentificationDetailsType `xml:"location"` // maxOccurs="2"

	CompanyId *CompanyIdentificationType `xml:"companyId,omitempty"`

	// Flight number or trainNumber
	FlightOrtrainNumber string `xml:"flightOrtrainNumber,omitempty"`

	// Product details
	ProductDetail *AdditionalProductDetailsType `xml:"productDetail,omitempty"`

	// Additional product details
	AddProductDetail *ProductFacilitiesType `xml:"addProductDetail,omitempty"`

	// Attribute details
	AttributeDetails []*CodedAttributeInformationType_270108C `xml:"attributeDetails,omitempty"` // maxOccurs="20"
}

type TravellerDetailsType struct {
	// Direct reference of passenger assigned by requesting system.
	Ref *int32 `xml:"ref,omitempty"`

	// Traveller is an infant
	InfantIndicator *int32 `xml:"infantIndicator,omitempty"`
}

type TravellerReferenceInformationType struct {
	// Requested passenger type
	Ptc []string `xml:"ptc,omitempty"` // maxOccurs="4"

	// Traveller details
	Traveller []*TravellerDetailsType `xml:"traveller,omitempty"` // maxOccurs="9"
}

type UserIdentificationType struct {
	// Originator Identification Details
	OfficeIdentification *OriginatorIdentificationDetailsTypeI `xml:"officeIdentification,omitempty"`

	// Used to specify which kind of info is given in DE 9900.
	OfficeType string `xml:"officeType,omitempty"`

	// The code given to an agent by the originating reservation system.
	OfficeCode string `xml:"officeCode,omitempty"`
}

type ValueSearchCriteriaType struct {
	// Format limitations: an..35
	Ref string `xml:"ref,omitempty"`

	// Format limitations: an..18
	Value string `xml:"value,omitempty"`

	CriteriaDetails []*CriteriaiDetaislType `xml:"criteriaDetails,omitempty"` // maxOccurs="10"
}
