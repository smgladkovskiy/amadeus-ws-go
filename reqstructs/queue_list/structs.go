package queue_list

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/formats"
)

type QueueList struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/QDQLRQ_11_1_1A Queue_List"`

	// presence implies that this is a follow up scrolling entry to a previous entry. Absence implies start of a new search
	Scroll *ActionDetailsTypeI `xml:"scroll,omitempty"`  // minOccurs="0"

	// used to specify the target office for which the queue count is to be displayed
	TargetOffice *AdditionalBusinessSourceInformationTypeI `xml:"targetOffice,omitempty"`  // minOccurs="0"

	// used to specify the queue if required
	QueueNumber *QueueInformationTypeI `xml:"queueNumber,omitempty"`  // minOccurs="0"

	// used to select the category
	CategoryDetails *SubQueueInformationTypeI `xml:"categoryDetails,omitempty"`  // minOccurs="0"

	// date range as system defined
	Date *StructuredDateTimeInformationType `xml:"date,omitempty"`  // minOccurs="0"

	// defines the start point for the search and may also define the end point of the search
	ScanRange *RangeDetailsTypeI `xml:"scanRange,omitempty"`  // minOccurs="0"

	SearchCriteria []*SearchCriteria `xml:"searchCriteria,omitempty"`  // minOccurs="0" maxOccurs="3"

	// Passenger name list (all the names in the PNR).
	PassengerName []*TravellerInformationTypeI `xml:"passengerName,omitempty"`  // minOccurs="0" maxOccurs="5"

	// The last 2 characters of the sine of the agent who placed the PNR on queue.
	AgentSine *UserIdentificationType `xml:"agentSine,omitempty"`  // minOccurs="0"

	// Account number issue from AIAN entry in the PNR.
	AccountNumber *AccountingInformationElementType `xml:"accountNumber,omitempty"`  // minOccurs="0"

	FlightInformation []*FlightInformation `xml:"flightInformation,omitempty"`  // minOccurs="0" maxOccurs="5"

	// This is the point of sale of segments in PNRs: - 9 char Amadeus Office ID. - OR 2 char GDS code for OA PNRs  PNRs containing a segment sold in any Amadeus Office ID matching pattern NCE6X*** or ***BA0*** or sold in Sabre (1S) or Gallileo (1G).
	Pos []*PointOfSaleInformationType `xml:"pos,omitempty"`  // minOccurs="0" maxOccurs="5"

	// The repetition is 10 because we can transport: - until 5 tierLevel - until 5 customerValue, including possibly range of customerValue.  If we have tierLevel in the FTI, the customerValue must not be present. If we have customerValue in the FTI, the tierLevel must not be present.
	TierLevelAndCustomerValue []*FrequentTravellerIdentificationCodeType `xml:"tierLevelAndCustomerValue,omitempty"`  // minOccurs="0" maxOccurs="10"

	SortCriteria *SortCriteria `xml:"sortCriteria,omitempty"`  // minOccurs="0"
}

type SearchCriteria struct {

	// used to specify if ticketing, departure or creation options
	SearchOption *SelectionDetailsTypeI `xml:"searchOption"`

	// used to specify the dates to be searched on
	Dates []*StructuredPeriodInformationType `xml:"dates"`  // maxOccurs="5"
}

type FlightInformation struct {

	// It transport the type of flight information that will follow.
	FlightInformationType *StatusTypeI `xml:"flightInformationType"`

	// Board point or Off Point.
	BoardPointOrOffPoint []*OriginAndDestinationDetailsTypeI `xml:"boardPointOrOffPoint,omitempty"`  // minOccurs="0" maxOccurs="5"

	// Airline code or Flight Number (in fact, airline + flight number)
	AirlineCodeOrFlightNumber []*TransportIdentifierType `xml:"airlineCodeOrFlightNumber,omitempty"`  // minOccurs="0" maxOccurs="5"

	// Booking class.
	ClassOfService []*ProductInformationTypeI `xml:"classOfService,omitempty"`  // minOccurs="0" maxOccurs="5"

	// Segment status code.
	SegmentStatus []*RelatedProductInformationTypeI `xml:"segmentStatus,omitempty"`  // minOccurs="0" maxOccurs="5"
}

type SortCriteria struct {

	// dummy for SDT clash
	Dumbo *DummySegmentTypeI `xml:"dumbo"`

	// Determine the order of the display.
	SortOption []*SelectionDetailsTypeI `xml:"sortOption"`  // maxOccurs="3"
}

//
// Complex structs
//

type AccountingElementType struct {

	// Account number
	Number formats.AlphaNumericString_Length1To10 `xml:"number"`
}

type AccountingInformationElementType struct {

	// One of these 4 data elements is mandatory , but non in particular
	Account *AccountingElementType `xml:"account"`
}

type ActionDetailsTypeI struct {

	// used for scrollling purposes
	NumberOfItemsDetails *ProcessingInformationTypeI `xml:"numberOfItemsDetails"`
}

type AdditionalBusinessSourceInformationTypeI struct {

	// the office we are targetting
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType"`

	// contains the office ID
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails"`
}

type CompanyIdentificationTypeI struct {

	// Marketing company.
	MarketingCompany formats.AlphaNumericString_Length1To3 `xml:"marketingCompany,omitempty"`  // minOccurs="0"
}

type DummySegmentTypeI struct {
}

type FrequentTravellerIdentificationCodeType struct {

	// Frequent Traveller Info. Repetition 2 is used only in the case we provide a customer value range (only one is accepted).
	FrequentTravellerDetails []*FrequentTravellerIdentificationType `xml:"frequentTravellerDetails"`  // maxOccurs="2"
}

type FrequentTravellerIdentificationType struct {

	// This field specifies the Tier Level.   This is a 4 letter string indicating the airline's ranking of frequent flyers. It is not to be confused with Alliance tier.  If tierLevel is filled in a given FTI segment, customerValue must not be filled.
	TierLevel formats.AlphaNumericString_Length1To4 `xml:"tierLevel,omitempty"`  // minOccurs="0"

	// This field specifies the Customer value.   This is a 4 letter string indicating the airline's ranking of frequent flyers. It is not to be confused with Alliance tier.  If customerValue is filled in a given FTI segment, tierLevel field must not be filled.
	CustomerValue *formats.NumericInteger_Length1To4 `xml:"customerValue,omitempty"`  // minOccurs="0"
}

type LocationTypeU struct {

	// Office identification. It can contain wildcards.
	Name formats.AlphaNumericString_Length1To9 `xml:"name,omitempty"`  // minOccurs="0"
}

type OriginAndDestinationDetailsTypeI struct {

	// Board point
	Origin formats.AlphaNumericString_Length3To3 `xml:"origin,omitempty"`  // minOccurs="0"

	// Off point
	Destination formats.AlphaNumericString_Length3To3 `xml:"destination,omitempty"`  // minOccurs="0"
}

type OriginatorIdentificationDetailsTypeI struct {

	// the office that is being targetted
	InHouseIdentification1 formats.AlphaNumericString_Length1To9 `xml:"inHouseIdentification1"`
}

type PartyIdentifierTypeU struct {

	// GDS identifier: 1A, 1S, 1G.
	PartyIdentifier formats.AlphaNumericString_Length1To3 `xml:"partyIdentifier,omitempty"`  // minOccurs="0"
}

type PointOfSaleInformationType struct {

	// Party identification.
	PointOfSale *PartyIdentifierTypeU `xml:"pointOfSale,omitempty"`  // minOccurs="0"

	// Office id in case the party identifier is 1A.
	LocationDetails *LocationTypeU `xml:"locationDetails,omitempty"`  // minOccurs="0"
}

type ProcessingInformationTypeI struct {

	// determine if move up or move down required
	ActionQualifier formats.AlphaNumericString_Length1To3 `xml:"actionQualifier,omitempty"`  // minOccurs="0"
}

type ProductDetailsTypeI struct {

	// Class designator.
	Designator formats.AlphaNumericString_Length1To1 `xml:"designator"`
}

type ProductIdentificationDetailsTypeI struct {

	// Flight number.
	FlightNumber formats.AlphaNumericString_Length1To4 `xml:"flightNumber"`
}

type ProductInformationTypeI struct {

	// Booking class details.
	BookingClassDetails *ProductDetailsTypeI `xml:"bookingClassDetails"`
}

type QueueInformationDetailsTypeI struct {

	// queue number
	Number formats.NumericInteger_Length1To2 `xml:"number"`
}

type QueueInformationTypeI struct {

	// queue identification
	QueueDetails *QueueInformationDetailsTypeI `xml:"queueDetails"`
}

type RangeDetailsTypeI struct {

	// define is a range or not
	RangeQualifier formats.AlphaNumericString_Length1To3 `xml:"rangeQualifier,omitempty"`  // minOccurs="0"

	// define the start and possible end point of the scan
	RangeDetails []*RangeTypeI `xml:"rangeDetails"`  // maxOccurs="5"
}

type RangeTypeI struct {

	// starting point of the scan
	Min formats.NumericInteger_Length1To18 `xml:"min"`

	// ending point of the scan
	Max *formats.NumericInteger_Length1To18 `xml:"max,omitempty"`  // minOccurs="0"
}

type RelatedProductInformationTypeI struct {

	// Status code
	StatusCode formats.AlphaNumericString_Length2To2 `xml:"statusCode"`
}

type SelectionDetailsInformationTypeI struct {

	// used to determine if a new start or a continuation Also used for search and sort criteria on the ticketing, departure and creation dates
	Option formats.AlphaNumericString_Length1To3 `xml:"option"`
}

type SelectionDetailsTypeI struct {

	// used for search and sort criteria
	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails"`
}

type SourceTypeDetailsTypeI struct {

	// not needed - but mandatory field So just stick a 4 in it !!
	SourceQualifier1 formats.AlphaNumericString_Length1To3 `xml:"sourceQualifier1"`
}

type StatusDetailsTypeI struct {

	// Indicator showing what flight information will be transported.
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator"`
}

type StatusTypeI struct {

	// Flight status details.
	StatusDetails *StatusDetailsTypeI `xml:"statusDetails"`
}

type StructuredDateTimeInformationType struct {

	// used for date range only The date ranges are defined on central system as 1,2,3,4 The actual values of the ranges are set in the office profile
	TimeMode formats.NumericInteger_Length1To3 `xml:"timeMode"`
}

type StructuredDateTimeType struct {

	// Year number.
	Year formats.Year_YYYY `xml:"year"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day"`
}

type StructuredPeriodInformationType struct {

	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType `xml:"beginDateTime"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType `xml:"endDateTime,omitempty"`  // minOccurs="0"
}

type SubQueueInformationDetailsTypeI struct {

	// E for every category    A for cats with items to be worked C for category number N for nickname CN for both category number and nickname numeric for date range
	IdentificationType formats.AlphaNumericString_Length1To3 `xml:"identificationType"`

	// category number
	ItemNumber formats.AlphaNumericString_Length1To3 `xml:"itemNumber,omitempty"`  // minOccurs="0"

	// used for nickname on inbound used for category name on outbound
	ItemDescription formats.AlphaNumericString_Length1To35 `xml:"itemDescription,omitempty"`  // minOccurs="0"
}

type SubQueueInformationTypeI struct {

	// identifies the category or categories.
	SubQueueInfoDetails *SubQueueInformationDetailsTypeI `xml:"subQueueInfoDetails"`
}

type TransportIdentifierType struct {

	// Company identification.
	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"`  // minOccurs="0"

	// Flight details.
	FlightDetails []*ProductIdentificationDetailsTypeI `xml:"flightDetails,omitempty"`  // minOccurs="0" maxOccurs="2"
}

type TravellerInformationTypeI struct {

	// Traveller surname information.
	PaxDetails *TravellerSurnameInformationTypeI `xml:"paxDetails"`
}

type TravellerSurnameInformationTypeI struct {

	// Passenger surname.
	Surname formats.AlphaNumericString_Length1To70 `xml:"surname"`
}

type UserIdentificationType struct {

	// The last 2 characters of the sine of the agent who placed the PNR on queue.
	Originator formats.AlphaNumericString_Length1To2 `xml:"originator"`
}
