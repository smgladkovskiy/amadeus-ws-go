package Air_SellFromRecommendationResponse_v05_2 // itares052

//import "encoding/xml"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/ITARES_05_2_IA Air_SellFromRecommendationReply"`
	Message             *Message               `xml:"message,omitempty"`
	ErrorAtMessageLevel []*ErrorAtMessageLevel `xml:"errorAtMessageLevel,omitempty"` // maxOccurs="5"
	ItineraryDetails    []*ItineraryDetails    `xml:"itineraryDetails,omitempty"`    // maxOccurs="99"
}

type Message struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails,omitempty"`
	ResponseType           string                  `xml:"responseType,omitempty"` // Format limitations: an..3
}

type MessageFunctionDetails struct {
	BusinessFunction          string   `xml:"businessFunction,omitempty"`          // Format limitations: an..3
	MessageFunction           string   `xml:"messageFunction,omitempty"`           // Format limitations: an..3
	ResponsibleAgency         string   `xml:"responsibleAgency,omitempty"`         // Format limitations: an..3
	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"` // maxOccurs="20" // Format limitations: an..3
}

type ErrorAtMessageLevel struct {
	ErrorSegment    *ErrorSegment    `xml:"errorSegment"`
	InformationText *InformationText `xml:"informationText,omitempty"`
}

type ErrorSegment struct {
	ErrorDetails *ErrorDetails `xml:"errorDetails"`
}

type ErrorDetails struct {
	ErrorCode      string `xml:"errorCode"`                // Format limitations: an..5
	ErrorCategory  string `xml:"errorCategory"`            // Format limitations: an..3
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"` // Format limitations: an..3
}

type InformationText struct {
	FreeTextQualification *FreeTextQualification `xml:"freeTextQualification,omitempty"`
	FreeText              []string               `xml:"freeText,omitempty"` // maxOccurs="99" // Format limitations: an..70
}

type FreeTextQualification struct {
	TextSubjectQualifier string `xml:"textSubjectQualifier"`      // Format limitations: an..3
	InformationType      string `xml:"informationType,omitempty"` // Format limitations: an..4
	Status               string `xml:"status,omitempty"`          // Format limitations: an..3
	CompanyId            string `xml:"companyId,omitempty"`       // Format limitations: an..35
	Language             string `xml:"language,omitempty"`        // Format limitations: an..3
}

type ItineraryDetails struct {
	OriginDestination   *OriginDestination     `xml:"originDestination"`
	ErrorItinerarylevel []*ErrorItinerarylevel `xml:"errorItinerarylevel,omitempty"` // maxOccurs="5"
	SegmentInformation  []*SegmentInformation  `xml:"segmentInformation,omitempty"`  // maxOccurs="9"
}

type OriginDestination struct {
	Origin      string `xml:"origin,omitempty"`      // Format limitations: an..3
	Destination string `xml:"destination,omitempty"` // Format limitations: an..3
}

type ErrorItinerarylevel struct {
	ErrorSegment    *ErrorSegment    `xml:"errorSegment"`
	InformationText *InformationText `xml:"informationText,omitempty"`
}

type SegmentInformation struct {
	FlightDetails       *FlightDetails         `xml:"flightDetails"`
	ApdSegment          *ApdSegment            `xml:"apdSegment,omitempty"`
	ActionDetails       *ActionDetails         `xml:"actionDetails"`
	InformationText     *InformationText       `xml:"informationText,omitempty"`
	ErrorAtSegmentLevel []*ErrorAtSegmentLevel `xml:"errorAtSegmentLevel,omitempty"` // maxOccurs="5"
}

type FlightDetails struct {
	FlightDate           *FlightDate           `xml:"flightDate,omitempty"`
	BoardPointDetails    *BoardPointDetails    `xml:"boardPointDetails,omitempty"`
	OffpointDetails      *OffpointDetails      `xml:"offpointDetails,omitempty"`
	CompanyDetails       *CompanyDetails       `xml:"companyDetails,omitempty"`
	FlightIdentification *FlightIdentification `xml:"flightIdentification,omitempty"`
	FlightTypeDetails    *FlightTypeDetails    `xml:"flightTypeDetails,omitempty"`
	SpecialSegment       string                `xml:"specialSegment,omitempty"`  // Format limitations: an..3
	MarriageDetails      []*MarriageDetails    `xml:"marriageDetails,omitempty"` // maxOccurs="99"
}

type FlightDate struct {
	DepartureDate string   `xml:"departureDate,omitempty"` // Format limitations: an..35
	DepartureTime *float64 `xml:"departureTime,omitempty"` // Format limitations: n..4
	ArrivalDate   string   `xml:"arrivalDate,omitempty"`   // Format limitations: an..35
	ArrivalTime   *float64 `xml:"arrivalTime,omitempty"`   // Format limitations: n..4
	DateVariation string   `xml:"dateVariation,omitempty"` // Format limitations: n1
}

type BoardPointDetails struct {
	TrueLocationId string `xml:"trueLocationId,omitempty"` // Format limitations: an..25
	TrueLocation   string `xml:"trueLocation,omitempty"`   // Format limitations: an..17
}

type OffpointDetails struct {
	TrueLocationId string `xml:"trueLocationId,omitempty"` // Format limitations: an..25
	TrueLocation   string `xml:"trueLocation,omitempty"`   // Format limitations: an..17
}

type CompanyDetails struct {
	MarketingCompany string `xml:"marketingCompany,omitempty"` // Format limitations: an..35
	OperatingCompany string `xml:"operatingCompany,omitempty"` // Format limitations: an..35
	OtherCompany     string `xml:"otherCompany,omitempty"`     // Format limitations: an..35
}

type FlightIdentification struct {
	FlightNumber      string   `xml:"flightNumber"`                // Format limitations: an..4
	BookingClass      string   `xml:"bookingClass,omitempty"`      // Format limitations: an..2
	OperationalSuffix string   `xml:"operationalSuffix,omitempty"` // Format limitations: an..3
	Modifier          []string `xml:"modifier,omitempty"`          // maxOccurs="3" // Format limitations: an..7
}

type FlightTypeDetails struct {
	// Format limitations: an..6
	FlightIndicator []string `xml:"flightIndicator"` // maxOccurs="9"
}

type MarriageDetails struct {
	Relation           string   `xml:"relation,omitempty"`           // Format limitations: an..3
	MarriageIdentifier *float64 `xml:"marriageIdentifier,omitempty"` // Format limitations: n..10
	LineNumber         *float64 `xml:"lineNumber,omitempty"`         // Format limitations: n..6
	OtherRelation      string   `xml:"otherRelation,omitempty"`      // Format limitations: an..3
	CarrierCode        string   `xml:"carrierCode,omitempty"`        // Format limitations: an..35
}

type ApdSegment struct {
	LegDetails            *LegDetails              `xml:"legDetails,omitempty"`
	DepartureStationInfo  *DepartureStationInfo    `xml:"departureStationInfo,omitempty"`
	ArrivalStationInfo    *ArrivalStationInfo      `xml:"arrivalStationInfo,omitempty"`
	FacilitiesInformation []*FacilitiesInformation `xml:"facilitiesInformation,omitempty"` // maxOccurs="10"
}

type LegDetails struct {
	Equipment                 string   `xml:"equipment,omitempty"`                 // Format limitations: an..8
	NumberOfStops             *float64 `xml:"numberOfStops,omitempty"`             // Format limitations: n..3
	Duration                  *float64 `xml:"duration,omitempty"`                  // Format limitations: n..6
	Percentage                *float64 `xml:"percentage,omitempty"`                // Format limitations: n..8
	DaysOfOperation           string   `xml:"daysOfOperation,omitempty"`           // Format limitations: an..7
	DateTimePeriod            string   `xml:"dateTimePeriod,omitempty"`            // Format limitations: an..35
	ComplexingFlightIndicator string   `xml:"complexingFlightIndicator,omitempty"` // Format limitations: n1
	Locations                 []string `xml:"locations,omitempty"`                 // maxOccurs="3" // Format limitations: an..25
}

type DepartureStationInfo struct {
	GateDescription string `xml:"gateDescription,omitempty"` // Format limitations: an..6
	Terminal        string `xml:"terminal,omitempty"`        // Format limitations: an..25
	Concourse       string `xml:"concourse,omitempty"`       // Format limitations: an..25
}

type ArrivalStationInfo struct {
	GateDescription string `xml:"gateDescription,omitempty"` // Format limitations: an..6
	Terminal        string `xml:"terminal,omitempty"`        // Format limitations: an..25
	Concourse       string `xml:"concourse,omitempty"`       // Format limitations: an..25
}

type FacilitiesInformation struct {
	Code          string   `xml:"code,omitempty"`          // Format limitations: an..3
	Description   string   `xml:"description,omitempty"`   // Format limitations: an..70
	Qualifier     string   `xml:"qualifier,omitempty"`     // Format limitations: an..3
	ExtensionCode []string `xml:"extensionCode,omitempty"` // maxOccurs="26" // Format limitations: an..17
}

type ActionDetails struct {
	Quantity   float64  `xml:"quantity"`   // Format limitations: n..15
	StatusCode []string `xml:"statusCode"` // maxOccurs="10" // Format limitations: an..3
}

type ErrorAtSegmentLevel struct {
	ErrorSegment    *ErrorSegment    `xml:"errorSegment"`
	InformationText *InformationText `xml:"informationText,omitempty"`
}
