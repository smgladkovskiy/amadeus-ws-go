package Air_SellFromRecommendationRequest_v05_2 // itareq052

import "encoding/xml"

type Request struct {
	XMLName              xml.Name              `xml:"http://xml.amadeus.com/ITAREQ_05_2_IA Air_SellFromRecommendation"`
	MessageActionDetails *MessageActionDetails `xml:"messageActionDetails,omitempty"`
	RecordLocator        *RecordLocator        `xml:"recordLocator,omitempty"`
	ItineraryDetails     []*ItineraryDetails   `xml:"itineraryDetails,omitempty"` // maxOccurs="99"
}

type MessageActionDetails struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`
}

type MessageFunctionDetails struct {
	MessageFunction           string   `xml:"messageFunction,omitempty"`
	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"` // maxOccurs="20"
}

type RecordLocator struct {
	Reservation []*Reservation `xml:"reservation,omitempty"` // maxOccurs="9"
}

type Reservation struct {
	CompanyId     string   `xml:"companyId,omitempty"`
	ControlNumber string   `xml:"controlNumber,omitempty"`
	ControlType   string   `xml:"controlType,omitempty"`
	Date          string   `xml:"date,omitempty"`
	Time          *float64 `xml:"time,omitempty"`
}

type ItineraryDetails struct {
	OriginDestinationDetails *OriginDestinationDetails `xml:"originDestinationDetails"`
	Message                  *Message                  `xml:"message,omitempty"`
	SegmentInformation       []*SegmentInformation     `xml:"segmentInformation"` // maxOccurs="9"
}

type OriginDestinationDetails struct {
	Origin      string `xml:"origin,omitempty"`
	Destination string `xml:"destination,omitempty"`
}

type Message struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`
}

type SegmentInformation struct {
	TravelProductInformation  *TravelProductInformation  `xml:"travelProductInformation"`
	RelatedproductInformation *RelatedproductInformation `xml:"relatedproductInformation"`
}

type TravelProductInformation struct {
	FlightDate           *FlightDate           `xml:"flightDate"`
	BoardPointDetails    *BoardPointDetails    `xml:"boardPointDetails"`
	OffpointDetails      *OffpointDetails      `xml:"offpointDetails"`
	CompanyDetails       *CompanyDetails       `xml:"companyDetails"`
	FlightIdentification *FlightIdentification `xml:"flightIdentification"`
	FlightTypeDetails    *FlightTypeDetails    `xml:"flightTypeDetails,omitempty"`
	SpecialSegment       string                `xml:"specialSegment,omitempty"`
	MarriageDetails      []*MarriageDetails    `xml:"marriageDetails,omitempty"` // maxOccurs="99"
}

type FlightDate struct {
	DepartureDate string   `xml:"departureDate,omitempty"`
	DepartureTime *float64 `xml:"departureTime,omitempty"`
	ArrivalDate   string   `xml:"arrivalDate,omitempty"`
	ArrivalTime   *float64 `xml:"arrivalTime,omitempty"`
	DateVariation string   `xml:"dateVariation,omitempty"`
}

type BoardPointDetails struct {
	TrueLocationId string `xml:"trueLocationId"`
}

type OffpointDetails struct {
	TrueLocationId string `xml:"trueLocationId"`
}

type CompanyDetails struct {
	MarketingCompany string `xml:"marketingCompany"`
}

type FlightIdentification struct {
	FlightNumber      string `xml:"flightNumber"`
	BookingClass      string `xml:"bookingClass"`
	OperationalSuffix string `xml:"operationalSuffix,omitempty"`
}

type FlightTypeDetails struct {
	FlightIndicator []string `xml:"flightIndicator"` // maxOccurs="9"
}

type MarriageDetails struct {
	Relation           string   `xml:"relation,omitempty"`
	MarriageIdentifier *float64 `xml:"marriageIdentifier,omitempty"`
	LineNumber         *float64 `xml:"lineNumber,omitempty"`
}

type RelatedproductInformation struct {
	Quantity   float64  `xml:"quantity"`
	StatusCode []string `xml:"statusCode,omitempty"` // maxOccurs="10"
}
