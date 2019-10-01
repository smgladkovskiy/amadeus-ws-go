package Fare_CheckRulesResponse_v07_1 // farqnr071

//import "encoding/xml"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/FARQNR_07_1_1A Fare_CheckRulesReply"`
	TransactionType *TransactionType `xml:"transactionType"`
	StatusInfo      *StatusInfo      `xml:"statusInfo,omitempty"`
	FareRouteInfo   *FareRouteInfo   `xml:"fareRouteInfo,omitempty"`
	InfoText        []*InfoText      `xml:"infoText,omitempty"` // maxOccurs="999"
	ErrorInfo       *ErrorInfo       `xml:"errorInfo,omitempty"`
	TariffInfo      []*TariffInfo    `xml:"tariffInfo,omitempty"`    // maxOccurs="999"
	FlightDetails   []*FlightDetails `xml:"flightDetails,omitempty"` // maxOccurs="999"
}

type TransactionType struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`
}

type MessageFunctionDetails struct {
	// Format limitations: an..3
	MessageFunction string `xml:"messageFunction,omitempty"`
}

type StatusInfo struct {
	StatusDetails *StatusDetails  `xml:"statusDetails"`
	OtherDetails  []*OtherDetails `xml:"otherDetails,omitempty"` // maxOccurs="98"
}

type StatusDetails struct {
	// Format limitations: an..3
	Indicator string `xml:"indicator,omitempty"`
}

type OtherDetails struct {
	// Format limitations: an..3
	Indicator string `xml:"indicator,omitempty"`
}

type FareRouteInfo struct {
	DayOfWeek            string                `xml:"dayOfWeek,omitempty"` // Format limitations: an..7
	FareQualifierDetails *FareQualifierDetails `xml:"fareQualifierDetails,omitempty"`
	IdentificationNumber string                `xml:"identificationNumber,omitempty"` // Format limitations: an..35
	ValidityPeriod       *ValidityPeriod       `xml:"validityPeriod,omitempty"`
}

type FareQualifierDetails struct {
	// Format limitations: an..3
	FareQualifier []string `xml:"fareQualifier,omitempty"` // maxOccurs="3"
}

type ValidityPeriod struct {
	FirstDate  *float64 `xml:"firstDate,omitempty"`  // Format limitations: n..6
	SecondDate *float64 `xml:"secondDate,omitempty"` // Format limitations: n..6
}

type InfoText struct {
	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`
	FreeText              []string               `xml:"freeText,omitempty"` // maxOccurs="99"// Format limitations: an..70
}

type FreeTextQualification struct {
	TextSubjectQualifier string `xml:"textSubjectQualifier"`      // Format limitations: an..3
	InformationType      string `xml:"informationType,omitempty"` // Format limitations: an..4
}

type ErrorInfo struct {
	RejectErrorCode *RejectErrorCode `xml:"rejectErrorCode"`
	ErrorFreeText   *ErrorFreeText   `xml:"errorFreeText,omitempty"`
}

type RejectErrorCode struct {
	ErrorDetails *ErrorDetails `xml:"errorDetails"`
}

type ErrorDetails struct {
	ErrorCode string `xml:"errorCode"` // Format limitations: an..3
}

type ErrorFreeText struct {
	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`
	FreeText              []string               `xml:"freeText,omitempty"` // maxOccurs="99"// Format limitations: an..70
}

type TariffInfo struct {
	FareRuleInfo *FareRuleInfo   `xml:"fareRuleInfo"`
	FareRuleText []*FareRuleText `xml:"fareRuleText,omitempty"` // maxOccurs="999"
}

type FareRuleInfo struct {
	RuleSectionLocalId string          `xml:"ruleSectionLocalId,omitempty"` // Format limitations: an..9
	CompanyDetails     *CompanyDetails `xml:"companyDetails,omitempty"`
	RuleCategoryCode   string          `xml:"ruleCategoryCode,omitempty"` // Format limitations: an..7
}

type CompanyDetails struct {
	MarketingCompany string `xml:"marketingCompany,omitempty"` // Format limitations: an..3
	Operatingcompany string `xml:"operatingcompany,omitempty"` // Format limitations: an..3
	OtherCompany     string `xml:"otherCompany,omitempty"`     // Format limitations: an..3
}

type FareRuleText struct {
	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`
	FreeText              []string               `xml:"freeText,omitempty"` // maxOccurs="99"// Format limitations: an..70
}

type FlightDetails struct {
	NbOfSegments             *NbOfSegments               `xml:"nbOfSegments"`
	AmountConversion         *AmountConversion           `xml:"amountConversion,omitempty"`
	QuantityValue            *QuantityValue              `xml:"quantityValue,omitempty"`
	PricingAndDateInfo       *PricingAndDateInfo         `xml:"pricingAndDateInfo,omitempty"`
	QualificationFareDetails []*QualificationFareDetails `xml:"qualificationFareDetails,omitempty"` // maxOccurs="99"
	TransportService         []*TransportService         `xml:"transportService,omitempty"`         // maxOccurs="4"
	FlightErrorCode          []*FlightErrorCode          `xml:"flightErrorCode,omitempty"`          // maxOccurs="999"
	ProductInfo              []*ProductInfo              `xml:"productInfo,omitempty"`              // maxOccurs="99"
	PriceInfo                []*PriceInfo                `xml:"priceInfo,omitempty"`                // maxOccurs="99"
	FareDetailInfo           []*FareDetailInfo           `xml:"fareDetailInfo,omitempty"`           // maxOccurs="9"
	OdiGrp                   []*OdiGrp                   `xml:"odiGrp,omitempty"`                   // maxOccurs="999"
	TravellerGrp             []*TravellerGrp             `xml:"travellerGrp,omitempty"`             // maxOccurs="99"
	FareRouteGrp             []*FareRouteGrp             `xml:"fareRouteGrp,omitempty"`             // maxOccurs="99"
	ItemGrp                  []*ItemGrp                  `xml:"itemGrp,omitempty"`                  // maxOccurs="999"
}

type NbOfSegments struct {
	SegmentControlDetails []*SegmentControlDetails `xml:"segmentControlDetails,omitempty"` // maxOccurs="9"
}

type SegmentControlDetails struct {
	Quantity           *float64 `xml:"quantity,omitempty"`           // Format limitations: n..15
	NumberOfUnits      *float64 `xml:"numberOfUnits,omitempty"`      // Format limitations: n..15
	TotalNumberOfItems *float64 `xml:"totalNumberOfItems,omitempty"` // Format limitations: n..15
}

type AmountConversion struct {
	ConversionRateDetails      *ConversionRateDetails        `xml:"conversionRateDetails"`
	OtherConversionRateDetails []*OtherConversionRateDetails `xml:"otherConversionRateDetails,omitempty"` // maxOccurs="19"
}

type ConversionRateDetails struct {
	ConversionType          string   `xml:"conversionType,omitempty"`          // Format limitations: an..3
	Currency                string   `xml:"currency,omitempty"`                // Format limitations: an..3
	RateType                string   `xml:"rateType,omitempty"`                // Format limitations: an..3
	PricingAmount           *float64 `xml:"pricingAmount,omitempty"`           // Format limitations: n..18
	MeasurementValue        *float64 `xml:"measurementValue,omitempty"`        // Format limitations: n..18
	MeasurementSignificance string   `xml:"measurementSignificance,omitempty"` // Format limitations: an..3
}

type OtherConversionRateDetails struct {
	ConversionType          string   `xml:"conversionType,omitempty"`          // Format limitations: an..3
	Currency                string   `xml:"currency,omitempty"`                // Format limitations: an..3
	RateType                string   `xml:"rateType,omitempty"`                // Format limitations: an..3
	PricingAmount           *float64 `xml:"pricingAmount,omitempty"`           // Format limitations: n..18
	MeasurementValue        *float64 `xml:"measurementValue,omitempty"`        // Format limitations: n..18
	MeasurementSignificance string   `xml:"measurementSignificance,omitempty"` // Format limitations: an..3
}

type QuantityValue struct {
	QuantityDetails []*QuantityDetails `xml:"quantityDetails"` // maxOccurs="9"
}

type QuantityDetails struct {
	Qualifier string  `xml:"qualifier"`      // Format limitations: an..3
	Value     float64 `xml:"value"`          // Format limitations: n..5
	Unit      string  `xml:"unit,omitempty"` // Format limitations: an..3
}

type PricingAndDateInfo struct {
	ProductDateTimeDetails *ProductDateTimeDetails `xml:"productDateTimeDetails,omitempty"`
	LocationDetails        *LocationDetails        `xml:"locationDetails,omitempty"`
	OtherLocationDetails   *LocationDetails        `xml:"otherLocationDetails,omitempty"`
	IdNumber               string                  `xml:"idNumber,omitempty"` // Format limitations: an..35
}

type ProductDateTimeDetails struct {
	DepartureDate string `xml:"departureDate,omitempty"` // Format limitations: n6
	ArrivalDate   string `xml:"arrivalDate,omitempty"`   // Format limitations: n6
}

type LocationDetails struct {
	City    string `xml:"city,omitempty"`    // Format limitations: an3..5
	Country string `xml:"country,omitempty"` // Format limitations: an..3
}

type QualificationFareDetails struct {
	MovementType          string                 `xml:"movementType,omitempty"` // Format limitations: an..3
	FareCategories        *FareCategories        `xml:"fareCategories,omitempty"`
	FareDetails           *FareDetails           `xml:"fareDetails,omitempty"`
	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`
	DiscountDetails       []*DiscountDetails     `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type FareCategories struct {
	FareType []string `xml:"fareType"` // maxOccurs="9"// Format limitations: an..3
}

type FareDetails struct {
	Qualifier    string   `xml:"qualifier,omitempty"`    // Format limitations: an..3
	Rate         *float64 `xml:"rate,omitempty"`         // Format limitations: n..3
	Country      string   `xml:"country,omitempty"`      // Format limitations: an..3
	FareCategory string   `xml:"fareCategory,omitempty"` // Format limitations: an..3
}

type AdditionalFareDetails struct {
	RateClass         string   `xml:"rateClass,omitempty"`         // Format limitations: an..10
	CommodityCategory string   `xml:"commodityCategory,omitempty"` // Format limitations: an..13
	FareClass         []string `xml:"fareClass,omitempty"`         // maxOccurs="2"// Format limitations: an..10
}

type DiscountDetails struct {
	// Format limitations: an..3
	FareQualifier string `xml:"fareQualifier"`

	// Format limitations: an..35
	RateCategory string `xml:"rateCategory,omitempty"`

	// Format limitations: n..3
	Amount *float64 `xml:"amount,omitempty"`

	// Format limitations: n..3
	Percentage *float64 `xml:"percentage,omitempty"`
}

type TransportService struct {
	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"` // maxOccurs="99"
}

type CompanyIdentification struct {
	// Format limitations: an..3
	MarketingCompany string `xml:"marketingCompany,omitempty"`

	// Format limitations: an..3
	Operatingcompany string `xml:"operatingcompany,omitempty"`

	// Format limitations: an..3
	OtherCompany string `xml:"otherCompany,omitempty"`
}

type ProductIdentificationDetails struct {
	// Format limitations: an..4
	FlightNumber string `xml:"flightNumber"`

	// Format limitations: an1
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`
}

type FlightErrorCode struct {
	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"` // maxOccurs="99"
}

type ProductInfo struct {
	ProductDetails *ProductDetails `xml:"productDetails"`

	ProductErrorCode []*ProductErrorCode `xml:"productErrorCode,omitempty"` // maxOccurs="99"
}

type ProductDetails struct {
	// Format limitations: an..3
	ProductDetailsQualifier string `xml:"productDetailsQualifier,omitempty"`

	BookingClassDetails []*BookingClassDetails `xml:"bookingClassDetails"` // maxOccurs="26"
}

type BookingClassDetails struct {
	// Format limitations: an1
	Designator string `xml:"designator"`

	// Format limitations: an1
	Option []string `xml:"option,omitempty"` // maxOccurs="3"
}

type ProductErrorCode struct {
	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"` // maxOccurs="99"
}

type PriceInfo struct {
	MonetaryRates *MonetaryRates `xml:"monetaryRates"`

	TaxAmount *TaxAmount `xml:"taxAmount,omitempty"`

	FareTypeInfo []*FareTypeInfo `xml:"fareTypeInfo,omitempty"` // maxOccurs="99"
}

type MonetaryRates struct {
	MonetaryDetails *MonetaryDetails `xml:"monetaryDetails"`

	AmountTwo []*AmountTwo `xml:"amountTwo,omitempty"` // maxOccurs="19"
}

type MonetaryDetails struct {
	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Format limitations: an..18
	Amount string `xml:"amount"`

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"`
}

type AmountTwo struct {
	// Format limitations: an..3
	TypeQualifier string `xml:"typeQualifier"`

	// Format limitations: an..18
	Amount string `xml:"amount"`

	// Format limitations: an..3
	Currency string `xml:"currency,omitempty"`
}

type TaxAmount struct {
	TaxDetails []*TaxDetails `xml:"taxDetails,omitempty"` // maxOccurs="99"
}

type TaxDetails struct {
	// Format limitations: an..17
	Rate string `xml:"rate,omitempty"`

	// Format limitations: an..3
	CountryCode string `xml:"countryCode,omitempty"`

	// Format limitations: an..3
	CurrencyCode string `xml:"currencyCode,omitempty"`

	// Format limitations: an..3
	Type []string `xml:"type,omitempty"` // maxOccurs="99"
}

type FareTypeInfo struct {
	FareDetailQualif *FareDetailQualif `xml:"fareDetailQualif"`

	FlightMovementDate *FlightMovementDate `xml:"flightMovementDate,omitempty"`

	FaraRulesInfo *FaraRulesInfo `xml:"faraRulesInfo,omitempty"`

	SelectionMakingDetails *SelectionMakingDetails `xml:"selectionMakingDetails,omitempty"`

	AmountConvDetails *AmountConvDetails `xml:"amountConvDetails,omitempty"`
}

type FareDetailQualif struct {
	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type FlightMovementDate struct {
	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}

type DateAndTimeDetails struct {
	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"`

	// Format limitations: an..6
	Date string `xml:"date,omitempty"`
}

type FaraRulesInfo struct {
	// Format limitations: an..9
	RuleSectionLocalId string `xml:"ruleSectionLocalId,omitempty"`

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`

	// Format limitations: an..7
	RuleCategoryCode string `xml:"ruleCategoryCode,omitempty"`
}

type SelectionMakingDetails struct {
	SelectionDetails *SelectionDetails `xml:"selectionDetails"`

	SelectionDetailsTwo []*SelectionDetailsTwo `xml:"selectionDetailsTwo,omitempty"` // maxOccurs="98"
}

type SelectionDetails struct {
	// Format limitations: an..3
	Option string `xml:"option"`

	// Format limitations: an..3
	OptionInformation string `xml:"optionInformation,omitempty"`
}

type SelectionDetailsTwo struct {
	// Format limitations: an..3
	Option string `xml:"option"`

	// Format limitations: an..3
	OptionInformation string `xml:"optionInformation,omitempty"`
}

type AmountConvDetails struct {
	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`

	OtherConversionRateDetails []*OtherConversionRateDetails `xml:"otherConversionRateDetails,omitempty"` // maxOccurs="19"
}

type FareDetailInfo struct {
	NbOfUnits           *NbOfUnits          `xml:"nbOfUnits"`
	PricingPlusDateInfo *PricingAndDateInfo `xml:"pricingPlusDateInfo,omitempty"`
	FareDeatilInfo      *FareDeatilInfo     `xml:"fareDeatilInfo,omitempty"`
}

type NbOfUnits struct {
	QuantityDetails []*QuantityDetails1 `xml:"quantityDetails"` // maxOccurs="9"
}

type QuantityDetails1 struct {
	// Format limitations: n..15
	NumberOfUnit *float64 `xml:"numberOfUnit,omitempty"`

	// Format limitations: an..3
	UnitQualifier string `xml:"unitQualifier,omitempty"`
}

type FareDeatilInfo struct {
	FareTypeGrouping *FareTypeGrouping `xml:"fareTypeGrouping,omitempty"`
}

type FareTypeGrouping struct {
	// Format limitations: an..6
	PricingGroup []string `xml:"pricingGroup,omitempty"` // maxOccurs="5"
}

type OdiGrp struct {
	OriginDestination *OriginDestination `xml:"originDestination"`

	FlightDateAndTime []*FlightDateAndTime `xml:"flightDateAndTime,omitempty"` // maxOccurs="5"

	FlightErrorText *FlightErrorText `xml:"flightErrorText,omitempty"`

	MonGrp []*MonGrp `xml:"monGrp,omitempty"` // maxOccurs="99"

	RoutingGrp []*RoutingGrp `xml:"routingGrp,omitempty"` // maxOccurs="5"

	TravelProductGrp []*TravelProductGrp `xml:"travelProductGrp,omitempty"` // maxOccurs="999"
}

type OriginDestination struct {
	// Format limitations: an3..5
	Origin string `xml:"origin,omitempty"`

	// Format limitations: an3..5
	Destination string `xml:"destination,omitempty"`
}

type FlightDateAndTime struct {
	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}

type FlightErrorText struct {
	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"` // maxOccurs="99"
}

type MonGrp struct {
	MonetaryValues *MonetaryValues `xml:"monetaryValues"`

	FareDetailGrp []*FareDetailGrp `xml:"fareDetailGrp,omitempty"` // maxOccurs="99"
}

type MonetaryValues struct {
	MonetaryDetails *MonetaryDetails `xml:"monetaryDetails"`

	AmountTwo []*AmountTwo `xml:"amountTwo,omitempty"` // maxOccurs="19"
}

type FareDetailGrp struct {
	FareQualif *FareQualif `xml:"fareQualif"`

	AmountCvtRate *AmountCvtRate `xml:"amountCvtRate,omitempty"`
}

type FareQualif struct {
	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type AmountCvtRate struct {
	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`

	OtherConversionRateDetails []*OtherConversionRateDetails `xml:"otherConversionRateDetails,omitempty"` // maxOccurs="19"
}

type RoutingGrp struct {
	RoutingInfo *RoutingInfo `xml:"routingInfo"`

	ServiceTransport *ServiceTransport `xml:"serviceTransport,omitempty"`

	QualificationOfFare *QualificationOfFare `xml:"qualificationOfFare,omitempty"`

	PertinentQuantity *PertinentQuantity `xml:"pertinentQuantity,omitempty"`
}

type RoutingInfo struct {
	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"` // maxOccurs="99"
}

type RoutingDetails struct {
	// Format limitations: an3..5
	Station string `xml:"station,omitempty"`

	// Format limitations: an3..5
	OtherStation string `xml:"otherStation,omitempty"`

	// Format limitations: an..3
	Qualifier string `xml:"qualifier,omitempty"`
}

type ServiceTransport struct {
	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"` // maxOccurs="99"
}

type QualificationOfFare struct {
	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type PertinentQuantity struct {
	QuantityDetails []*QuantityDetails `xml:"quantityDetails"` // maxOccurs="9"
}

type TravelProductGrp struct {
	TravelProductInfo *TravelProductInfo `xml:"travelProductInfo"`

	RoutingGrp []*RoutingGrp1 `xml:"routingGrp,omitempty"` // maxOccurs="99"
}

type TravelProductInfo struct {
	FlightDate *FlightDate `xml:"flightDate,omitempty"`

	BoardPointDetails *BoardPointDetails `xml:"boardPointDetails,omitempty"`

	OffpointDetails *OffpointDetails `xml:"offpointDetails,omitempty"`

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`

	FlightIdentification *FlightIdentification `xml:"flightIdentification,omitempty"`

	FlightTypeDetails *FlightTypeDetails `xml:"flightTypeDetails,omitempty"`

	// Format limitations: n..6
	ItemNumber *float64 `xml:"itemNumber,omitempty"`

	// Format limitations: an..3
	SpecialSegment string `xml:"specialSegment,omitempty"`

	MarriageDetails *MarriageDetails `xml:"marriageDetails,omitempty"`
}

type FlightDate struct {
	// Format limitations: n6
	DepartureDate string `xml:"departureDate,omitempty"`

	// Format limitations: n6
	ArrivalDate string `xml:"arrivalDate,omitempty"`
}

type BoardPointDetails struct {
	// Format limitations: an3..5
	TrueLocationId string `xml:"trueLocationId,omitempty"`

	// Format limitations: an3..5
	TrueLocation string `xml:"trueLocation,omitempty"`
}

type OffpointDetails struct {
	// Format limitations: an3..5
	TrueLocationId string `xml:"trueLocationId,omitempty"`

	// Format limitations: an3..5
	TrueLocation string `xml:"trueLocation,omitempty"`
}

type FlightIdentification struct {
	// Format limitations: an..4
	FlightNumber string `xml:"flightNumber"`

	// Format limitations: an1
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`
}

type FlightTypeDetails struct {
	// Format limitations: an..6
	FlightIndicator string `xml:"flightIndicator"`

	// Format limitations: an..6
	SecondSequenceNb []string `xml:"secondSequenceNb,omitempty"` // maxOccurs="8"
}

type MarriageDetails struct {
	// Format limitations: an..3
	Relation string `xml:"relation,omitempty"`

	// Format limitations: n..10
	MarriageIdentifier *float64 `xml:"marriageIdentifier,omitempty"`

	// Format limitations: n..6
	LineNumber *float64 `xml:"lineNumber,omitempty"`

	// Format limitations: an..3
	OtherRelation string `xml:"otherRelation,omitempty"`

	// Format limitations: an..35
	CarrierCode string `xml:"carrierCode,omitempty"`
}

type RoutingGrp1 struct {
	RoutingInfo *RoutingInfo `xml:"routingInfo"`

	TransportServiceChange []*TransportServiceChange `xml:"transportServiceChange,omitempty"` // maxOccurs="9"
}

type TransportServiceChange struct {
	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"` // maxOccurs="99"
}

type TravellerGrp struct {
	TravellerIdentRef *TravellerIdentRef `xml:"travellerIdentRef"`

	FareRulesDetails *FareRulesDetails `xml:"fareRulesDetails,omitempty"`

	FlightMovementDateInfo *FlightMovementDateInfo `xml:"flightMovementDateInfo,omitempty"`
}

type TravellerIdentRef struct {
	ReferenceDetails []*ReferenceDetails `xml:"referenceDetails"` // maxOccurs="99"
}

type ReferenceDetails struct {
	// Format limitations: an..3
	Type string `xml:"type,omitempty"`

	// Format limitations: an..10
	Value string `xml:"value,omitempty"`
}

type FareRulesDetails struct {
	// Format limitations: an..9
	TariffClassId string `xml:"tariffClassId,omitempty"`

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`

	// Format limitations: an..7
	RuleSectionId []string `xml:"ruleSectionId,omitempty"` // maxOccurs="99"
}

type FlightMovementDateInfo struct {
	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}

type FareRouteGrp struct {
	FareRouteInfo *FareRouteInfo `xml:"fareRouteInfo"`

	JourneyGrp []*JourneyGrp `xml:"journeyGrp,omitempty"` // maxOccurs="999"
}

type JourneyGrp struct {
	JourneyOriginAndDestination *JourneyOriginAndDestination `xml:"journeyOriginAndDestination"`

	JourneyProductGrp []*JourneyProductGrp `xml:"journeyProductGrp,omitempty"` // maxOccurs="999"
}

type JourneyOriginAndDestination struct {
	// Format limitations: an3..5
	Origin string `xml:"origin,omitempty"`

	// Format limitations: an3..5
	Destination string `xml:"destination,omitempty"`
}

type JourneyProductGrp struct {
	JourneyProductInfo *JourneyProductInfo `xml:"journeyProductInfo"`

	JourneyRoutingGrp []*JourneyRoutingGrp `xml:"journeyRoutingGrp,omitempty"` // maxOccurs="99"
}

type JourneyProductInfo struct {
	FlightDate *FlightDate `xml:"flightDate,omitempty"`

	BoardPointDetails *BoardPointDetails `xml:"boardPointDetails,omitempty"`

	OffpointDetails *OffpointDetails `xml:"offpointDetails,omitempty"`

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`

	FlightIdentification *FlightIdentification `xml:"flightIdentification,omitempty"`

	FlightTypeDetails *FlightTypeDetails `xml:"flightTypeDetails,omitempty"`

	// Format limitations: n..6
	ItemNumber *float64 `xml:"itemNumber,omitempty"`

	// Format limitations: an..3
	SpecialSegment string `xml:"specialSegment,omitempty"`

	MarriageDetails *MarriageDetails `xml:"marriageDetails,omitempty"`
}

type JourneyRoutingGrp struct {
	JourneyRoutingInfo *JourneyRoutingInfo `xml:"journeyRoutingInfo"`

	JourneyTransportService *JourneyTransportService `xml:"journeyTransportService,omitempty"`
}

type JourneyRoutingInfo struct {
	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"` // maxOccurs="99"
}

type JourneyTransportService struct {
	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"` // maxOccurs="99"
}

type ItemGrp struct {
	ItemNb *ItemNb `xml:"itemNb"`

	ProductAvailabilityStatus *ProductAvailabilityStatus `xml:"productAvailabilityStatus,omitempty"`

	QuantityItem *QuantityItem `xml:"quantityItem,omitempty"`

	TransportServiceItem []*TransportServiceItem `xml:"transportServiceItem,omitempty"` // maxOccurs="4"

	FreeTextItem []*FreeTextItem `xml:"freeTextItem,omitempty"` // maxOccurs="99"

	FareQualifItem []*FareQualifItem `xml:"fareQualifItem,omitempty"` // maxOccurs="99"

	OriginDestinationGrp []*OriginDestinationGrp `xml:"originDestinationGrp,omitempty"` // maxOccurs="99"

	UnitGrp []*UnitGrp `xml:"unitGrp,omitempty"` // maxOccurs="9"

	MonetaryGrp []*MonetaryGrp `xml:"monetaryGrp,omitempty"` // maxOccurs="99"

	FarerouteGrp []*FarerouteGrp `xml:"farerouteGrp,omitempty"` // maxOccurs="99"
}

type ItemNb struct {
	ItemNumberDetails []*ItemNumberDetails `xml:"itemNumberDetails"` // maxOccurs="99"
}

type ItemNumberDetails struct {
	// Format limitations: an..3
	Number string `xml:"number,omitempty"`

	// Format limitations: an..3
	Type string `xml:"type,omitempty"`
}

type ProductAvailabilityStatus struct {
	// Format limitations: an..3
	ProductDetailsQualifier string `xml:"productDetailsQualifier,omitempty"`

	BookingClassDetails []*BookingClassDetails `xml:"bookingClassDetails"` // maxOccurs="26"
}

type QuantityItem struct {
	QuantityDetails []*QuantityDetails `xml:"quantityDetails"` // maxOccurs="9"
}

type TransportServiceItem struct {
	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"` // maxOccurs="99"
}

type FreeTextItem struct {
	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`

	// Format limitations: an..70
	FreeText []string `xml:"freeText,omitempty"` // maxOccurs="99"
}

type FareQualifItem struct {
	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type OriginDestinationGrp struct {
	OriginDestOfJourney *OriginDestOfJourney `xml:"originDestOfJourney"`

	DateForMovements *DateForMovements `xml:"dateForMovements,omitempty"`

	RoutingForJourney *RoutingForJourney `xml:"routingForJourney,omitempty"`
}

type OriginDestOfJourney struct {
	// Format limitations: an3..5
	Origin string `xml:"origin,omitempty"`

	// Format limitations: an3..5
	Destination string `xml:"destination,omitempty"`
}

type DateForMovements struct {
	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}

type RoutingForJourney struct {
	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"` // maxOccurs="99"
}

type UnitGrp struct {
	NbOfUnits              *NbOfUnits          `xml:"nbOfUnits"`
	UnitPricingAndDateInfo *PricingAndDateInfo `xml:"unitPricingAndDateInfo,omitempty"`
	UnitFareDetails        *UnitFareDetails    `xml:"unitFareDetails,omitempty"`
}

type UnitFareDetails struct {
	FareTypeGrouping *FareTypeGrouping `xml:"fareTypeGrouping,omitempty"`
}

type MonetaryGrp struct {
	MonetaryValues          *MonetaryValues      `xml:"monetaryValues"`
	MonetFareRuleValues     *MonetFareRuleValues `xml:"monetFareRuleValues,omitempty"`
	MonetTravellerRef       *MonetTravellerRef   `xml:"monetTravellerRef,omitempty"`
	MonetTicketPriceAndDate *PricingAndDateInfo  `xml:"monetTicketPriceAndDate,omitempty"`
	MonetTaxValues          *MonetTaxValues      `xml:"monetTaxValues,omitempty"`
	QualifGrp               []*QualifGrp         `xml:"qualifGrp,omitempty"` // maxOccurs="99"
}

type MonetFareRuleValues struct {
	// Format limitations: an..9
	RuleSectionLocalId string `xml:"ruleSectionLocalId,omitempty"`

	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`

	// Format limitations: an..7
	RuleCategoryCode string `xml:"ruleCategoryCode,omitempty"`
}

type MonetTravellerRef struct {
	ReferenceDetails []*ReferenceDetails `xml:"referenceDetails"` // maxOccurs="99"
}

type MonetTaxValues struct {
	TaxDetails []*TaxDetails `xml:"taxDetails,omitempty"` // maxOccurs="99"
}

type QualifGrp struct {
	QualificationFare *QualificationFare `xml:"qualificationFare"`

	QualifSelection *QualifSelection `xml:"qualifSelection,omitempty"`

	QualifDateFlightMovement *QualifDateFlightMovement `xml:"qualifDateFlightMovement,omitempty"`

	QualifConversionRate *QualifConversionRate `xml:"qualifConversionRate,omitempty"`
}

type QualificationFare struct {
	// Format limitations: an..3
	MovementType string `xml:"movementType,omitempty"`

	FareCategories *FareCategories `xml:"fareCategories,omitempty"`

	FareDetails *FareDetails `xml:"fareDetails,omitempty"`

	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`

	DiscountDetails []*DiscountDetails `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type QualifSelection struct {
	SelectionDetails *SelectionDetails `xml:"selectionDetails"`

	SelectionDetailsTwo []*SelectionDetailsTwo `xml:"selectionDetailsTwo,omitempty"` // maxOccurs="98"
}

type QualifDateFlightMovement struct {
	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}

type QualifConversionRate struct {
	ConversionRateDetails *ConversionRateDetails `xml:"conversionRateDetails"`

	OtherConversionRateDetails []*OtherConversionRateDetails `xml:"otherConversionRateDetails,omitempty"` // maxOccurs="19"
}

type FarerouteGrp struct {
	InfoForFareRoute *InfoForFareRoute `xml:"infoForFareRoute"`

	FarerouteTransportService []*FarerouteTransportService `xml:"farerouteTransportService,omitempty"` // maxOccurs="99"

	FinalOdiGrp []*FinalOdiGrp `xml:"finalOdiGrp,omitempty"` // maxOccurs="99"
}

type InfoForFareRoute struct {
	// Format limitations: an..7
	DayOfWeek string `xml:"dayOfWeek,omitempty"`

	FareQualifierDetails *FareQualifierDetails `xml:"fareQualifierDetails,omitempty"`

	// Format limitations: an..35
	IdentificationNumber string `xml:"identificationNumber,omitempty"`

	ValidityPeriod *ValidityPeriod `xml:"validityPeriod,omitempty"`
}

type FarerouteTransportService struct {
	CompanyIdentification *CompanyIdentification `xml:"companyIdentification"`

	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"` // maxOccurs="99"
}

type FinalOdiGrp struct {
	FinalOriginDestination *FinalOriginDestination `xml:"finalOriginDestination"`

	LastOdiRoutingInfo *LastOdiRoutingInfo `xml:"lastOdiRoutingInfo,omitempty"`

	LastOdiDateFlightMovement *LastOdiDateFlightMovement `xml:"lastOdiDateFlightMovement,omitempty"`
}

type FinalOriginDestination struct {
	// Format limitations: an3..5
	Origin string `xml:"origin,omitempty"`

	// Format limitations: an3..5
	Destination string `xml:"destination,omitempty"`
}

type LastOdiRoutingInfo struct {
	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"` // maxOccurs="99"
}

type LastOdiDateFlightMovement struct {
	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}
