package Fare_CheckRulesRequest_v07_1 // farqnq071

import "encoding/xml"

type Request struct {
	XMLName              xml.Name                `xml:"http://xml.amadeus.com/FARQNQ_07_1_1A Fare_CheckRules"`
	MsgType              *MsgType                `xml:"msgType"`
	AvailcabinStatus     *AvailcabinStatus       `xml:"availcabinStatus,omitempty"`
	ConversionRate       *ConversionRate         `xml:"conversionRate,omitempty"`
	PricingTickInfo      *PricingTickInfo        `xml:"pricingTickInfo,omitempty"`
	MultiCorporate       *MultiCorporate         `xml:"multiCorporate,omitempty"`
	ItemNumber           *ItemNumber             `xml:"itemNumber,omitempty"`
	DateOfFlight         *DateOfFlight           `xml:"dateOfFlight,omitempty"`
	FlightQualification  []*FlightQualification  `xml:"flightQualification,omitempty"`  // maxOccurs="99"
	TransportInformation []*TransportInformation `xml:"transportInformation,omitempty"` // maxOccurs="99"
	TripDescription      []*TripDescription      `xml:"tripDescription,omitempty"`      // maxOccurs="99"
	PricingInfo          []*PricingInfo          `xml:"pricingInfo,omitempty"`          // maxOccurs="9"
	FareRule             *FareRule               `xml:"fareRule,omitempty"`
}

type MsgType struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`
}

type MessageFunctionDetails struct {
	MessageFunction string `xml:"messageFunction,omitempty"`
}

type AvailcabinStatus struct {
	ProductDetailsQualifier string                 `xml:"productDetailsQualifier,omitempty"`
	BookingClassDetails     []*BookingClassDetails `xml:"bookingClassDetails"` // maxOccurs="26"
}

type BookingClassDetails struct {
	Designator string   `xml:"designator"`
	Option     []string `xml:"option,omitempty"` // maxOccurs="3"
}

type ConversionRate struct {
	ConversionRateDetails *ConversionRateDetails  `xml:"conversionRateDetails"`
	OtherConvRateDetails  []*OtherConvRateDetails `xml:"otherConvRateDetails,omitempty"` // maxOccurs="19"
}

type ConversionRateDetails struct {
	ConversionType          string   `xml:"conversionType,omitempty"`
	Currency                string   `xml:"currency,omitempty"`
	RateType                string   `xml:"rateType,omitempty"`
	PricingAmount           *float64 `xml:"pricingAmount,omitempty"`
	ConvertedValueAmount    *float64 `xml:"convertedValueAmount,omitempty"`
	DutyTaxFeeType          string   `xml:"dutyTaxFeeType,omitempty"`
	MeasurementValue        *float64 `xml:"measurementValue,omitempty"`
	MeasurementSignificance string   `xml:"measurementSignificance,omitempty"`
}

type OtherConvRateDetails struct {
	ConversionType          string   `xml:"conversionType,omitempty"`
	Currency                string   `xml:"currency,omitempty"`
	RateType                string   `xml:"rateType,omitempty"`
	PricingAmount           *float64 `xml:"pricingAmount,omitempty"`
	ConvertedValueAmount    *float64 `xml:"convertedValueAmount,omitempty"`
	DutyTaxFeeType          string   `xml:"dutyTaxFeeType,omitempty"`
	MeasurementValue        *float64 `xml:"measurementValue,omitempty"`
	MeasurementSignificance string   `xml:"measurementSignificance,omitempty"`
}

type PricingTickInfo struct {
	ProductDateTimeDetails *ProductDateTimeDetails `xml:"productDateTimeDetails,omitempty"`
	LocationDetails        *LocationDetails        `xml:"locationDetails,omitempty"`
	OtherLocationDetails   *OtherLocationDetails   `xml:"otherLocationDetails,omitempty"`
	IdNumber               string                  `xml:"idNumber,omitempty"`
}

type ProductDateTimeDetails struct {
	DepartureDate string `xml:"departureDate,omitempty"`
	ArrivalDate   string `xml:"arrivalDate,omitempty"`
}

type LocationDetails struct {
	City    string `xml:"city,omitempty"`
	Country string `xml:"country,omitempty"`
}

type OtherLocationDetails struct {
	City    string `xml:"city,omitempty"`
	Country string `xml:"country,omitempty"`
}

type MultiCorporate struct {
	CorporateId []*CorporateId `xml:"corporateId"` // maxOccurs="20"
}

type CorporateId struct {
	CorporateQualifier string   `xml:"corporateQualifier"`
	Identity           []string `xml:"identity"` // maxOccurs="9"
}

type ItemNumber struct {
	ItemNumberDetails []*ItemNumberDetails `xml:"itemNumberDetails"` // maxOccurs="99"
}

type ItemNumberDetails struct {
	Number string `xml:"number,omitempty"`
	Type   string `xml:"type,omitempty"`
}

type DateOfFlight struct {
	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}

type DateAndTimeDetails struct {
	Qualifier string   `xml:"qualifier,omitempty"`
	Date      *float64 `xml:"date,omitempty"`
}

type FlightQualification struct {
	MovementType          string                 `xml:"movementType,omitempty"`
	FareCategories        *FareCategories        `xml:"fareCategories,omitempty"`
	FareDetails           *FareDetails           `xml:"fareDetails,omitempty"`
	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`
	DiscountDetails       []*DiscountDetails     `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type FareCategories struct {
	FareType      string   `xml:"fareType"`
	OtherFareType []string `xml:"otherFareType,omitempty"` // maxOccurs="8"
}

type FareDetails struct {
	Qualifier    string   `xml:"qualifier,omitempty"`
	Rate         *float64 `xml:"rate,omitempty"`
	Country      string   `xml:"country,omitempty"`
	FareCategory string   `xml:"fareCategory,omitempty"`
}

type AdditionalFareDetails struct {
	RateClass         string   `xml:"rateClass,omitempty"`
	CommodityCategory string   `xml:"commodityCategory,omitempty"`
	PricingGroup      string   `xml:"pricingGroup,omitempty"`
	SecondRateClass   []string `xml:"secondRateClass,omitempty"` // maxOccurs="29"
}

type DiscountDetails struct {
	FareQualifier string   `xml:"fareQualifier"`
	RateCategory  string   `xml:"rateCategory,omitempty"`
	Amount        *float64 `xml:"amount,omitempty"`
	Percentage    *float64 `xml:"percentage,omitempty"`
}

type TransportInformation struct {
	TransportService *TransportService `xml:"transportService"`
	AvailCabinConf   *AvailCabinConf   `xml:"availCabinConf,omitempty"`
	RoutingInfo      *RoutingInfo      `xml:"routingInfo,omitempty"`
	SelectionDetail  *SelectionDetail  `xml:"selectionDetail,omitempty"`
}

type TransportService struct {
	CompanyIdentification        *CompanyIdentification          `xml:"companyIdentification"`
	ProductIdentificationDetails []*ProductIdentificationDetails `xml:"productIdentificationDetails,omitempty"` // maxOccurs="99"
}

type CompanyIdentification struct {
	MarketingCompany string `xml:"marketingCompany,omitempty"`
	Operatingcompany string `xml:"operatingcompany,omitempty"`
	OtherCompany     string `xml:"otherCompany,omitempty"`
}

type ProductIdentificationDetails struct {
	FlightNumber      string `xml:"flightNumber"`
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`
}

type AvailCabinConf struct {
	ProductDetailsQualifier string                 `xml:"productDetailsQualifier,omitempty"`
	BookingClassDetails     []*BookingClassDetails `xml:"bookingClassDetails"` // maxOccurs="26"
}

type RoutingInfo struct {
	RoutingDetails []*RoutingDetails `xml:"routingDetails,omitempty"` // maxOccurs="99"
}

type RoutingDetails struct {
	Station      string `xml:"station,omitempty"`
	OtherStation string `xml:"otherStation,omitempty"`
	Qualifier    string `xml:"qualifier,omitempty"`
}

type SelectionDetail struct {
	SelectionDetails    *SelectionDetails      `xml:"selectionDetails"`
	SelectionDetailsTwo []*SelectionDetailsTwo `xml:"selectionDetailsTwo,omitempty"` // maxOccurs="98"
}

type SelectionDetails struct {
	Option            string `xml:"option"`
	OptionInformation string `xml:"optionInformation,omitempty"`
}

type SelectionDetailsTwo struct {
	Option            string `xml:"option"`
	OptionInformation string `xml:"optionInformation,omitempty"`
}

type TripDescription struct {
	OrigDest           *OrigDest           `xml:"origDest"`
	DateFlightMovement *DateFlightMovement `xml:"dateFlightMovement,omitempty"`
	Routing            []*Routing          `xml:"routing,omitempty"` // maxOccurs="99"
}

type OrigDest struct {
	Origin      string `xml:"origin,omitempty"`
	Destination string `xml:"destination,omitempty"`
}

type DateFlightMovement struct {
	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}

type Routing struct {
	RoutingInfo            *RoutingInfo            `xml:"routingInfo"`
	TransportService       *TransportService       `xml:"transportService,omitempty"`
	SegFareDetails         *SegFareDetails         `xml:"segFareDetails,omitempty"`
	PertinentQuantity      *PertinentQuantity      `xml:"pertinentQuantity,omitempty"`
	SelectionMakingDetails *SelectionMakingDetails `xml:"selectionMakingDetails,omitempty"`
}

type SegFareDetails struct {
	MovementType          string                 `xml:"movementType,omitempty"`
	FareCategories        *FareCategories        `xml:"fareCategories,omitempty"`
	FareDetails           *FareDetails           `xml:"fareDetails,omitempty"`
	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`
	DiscountDetails       []*DiscountDetails     `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type PertinentQuantity struct {
	QuantityDetails      *QuantityDetails        `xml:"quantityDetails"`
	OtherQuantityDetails []*OtherQuantityDetails `xml:"otherQuantityDetails,omitempty"` // maxOccurs="8"
}

type QuantityDetails struct {
	Qualifier string  `xml:"qualifier"`
	Value     float64 `xml:"value"`
	Unit      string  `xml:"unit,omitempty"`
}

type OtherQuantityDetails struct {
	Qualifier string  `xml:"qualifier"`
	Value     float64 `xml:"value"`
	Unit      string  `xml:"unit,omitempty"`
}

type SelectionMakingDetails struct {
	SelectionDetails    *SelectionDetails      `xml:"selectionDetails"`
	SelectionDetailsTwo []*SelectionDetailsTwo `xml:"selectionDetailsTwo,omitempty"` // maxOccurs="98"
}

type PricingInfo struct {
	NumberOfUnits     *NumberOfUnits       `xml:"numberOfUnits"`
	TicketPricingDate []*TicketPricingDate `xml:"ticketPricingDate,omitempty"` // maxOccurs="3"
	Fare              []*Fare              `xml:"fare,omitempty"`              // maxOccurs="99"
}

type NumberOfUnits struct {
	QuantityDetails      *QuantityDetails1        `xml:"quantityDetails"`
	OtherQuantityDetails []*OtherQuantityDetails1 `xml:"otherQuantityDetails,omitempty"` // maxOccurs="8"
}

type QuantityDetails1 struct {
	NumberOfUnit  *float64 `xml:"numberOfUnit,omitempty"`
	UnitQualifier string   `xml:"unitQualifier,omitempty"`
}

type OtherQuantityDetails1 struct {
	NumberOfUnit  *float64 `xml:"numberOfUnit,omitempty"`
	UnitQualifier string   `xml:"unitQualifier,omitempty"`
}

type TicketPricingDate struct {
	ProductDateTimeDetails *ProductDateTimeDetails `xml:"productDateTimeDetails,omitempty"`
	LocationDetails        *LocationDetails        `xml:"locationDetails,omitempty"`
	OtherLocationDetails   *OtherLocationDetails   `xml:"otherLocationDetails,omitempty"`
	IdNumber               string                  `xml:"idNumber,omitempty"`
}

type Fare struct {
	DetailsOfFare            *DetailsOfFare              `xml:"detailsOfFare"`
	FareQualificationDetails []*FareQualificationDetails `xml:"fareQualificationDetails,omitempty"` // maxOccurs="99"
}

type DetailsOfFare struct {
	FareTypeGrouping *FareTypeGrouping `xml:"fareTypeGrouping,omitempty"`
}

type FareTypeGrouping struct {
	PricingGroup []string `xml:"pricingGroup,omitempty"` // maxOccurs="5"
}

type FareQualificationDetails struct {
	MovementType          string                 `xml:"movementType,omitempty"`
	FareCategories        *FareCategories        `xml:"fareCategories,omitempty"`
	FareDetails           *FareDetails           `xml:"fareDetails,omitempty"`
	AdditionalFareDetails *AdditionalFareDetails `xml:"additionalFareDetails,omitempty"`
	DiscountDetails       []*DiscountDetails     `xml:"discountDetails,omitempty"` // maxOccurs="9"
}

type FareRule struct {
	TarifFareRule           *TarifFareRule           `xml:"tarifFareRule"`
	TravellerIdentification *TravellerIdentification `xml:"travellerIdentification,omitempty"`
	TravellerDate           *TravellerDate           `xml:"travellerDate,omitempty"`
}

type TarifFareRule struct {
	TariffClassId  string          `xml:"tariffClassId,omitempty"`
	CompanyDetails *CompanyDetails `xml:"companyDetails,omitempty"`
	RuleSectionId  []string        `xml:"ruleSectionId,omitempty"` // maxOccurs="99"
}

type CompanyDetails struct {
	MarketingCompany string `xml:"marketingCompany,omitempty"`
	Operatingcompany string `xml:"operatingcompany,omitempty"`
	OtherCompany     string `xml:"otherCompany,omitempty"`
}

type TravellerIdentification struct {
	ReferenceDetails []*ReferenceDetails `xml:"referenceDetails"` // maxOccurs="99"
}

type ReferenceDetails struct {
	Type  string `xml:"type,omitempty"`
	Value string `xml:"value,omitempty"`
}

type TravellerDate struct {
	DateAndTimeDetails []*DateAndTimeDetails `xml:"dateAndTimeDetails,omitempty"` // maxOccurs="99"
}
