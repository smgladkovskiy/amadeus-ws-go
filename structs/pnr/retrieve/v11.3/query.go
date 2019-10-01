package PNR_Retrieve_v11_3 // pnrret113

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/PNRRET_11_3_1A PNR_Retrieve"`

	Settings *Settings `xml:"settings,omitempty"`

	RetrievalFacts *RetrievalFacts `xml:"retrievalFacts"`
}

type Settings struct {
	// optional information on what needs to be returned in the PNR , hardcopy print or PNR  mode.
	Options *OptionalPNRActionsType `xml:"options"`

	// Identification of the printer when hardcopy is requested
	Printer *PrinterIdentificationType `xml:"printer,omitempty"`
}

type RetrievalFacts struct {
	// retrieval type , options , and references related to the PNR content .
	Retrieve *RetrievePNRType `xml:"retrieve"`

	// Informations needed for retreive by record locator or retreive by customer profile.
	ReservationOrProfileIdentifier *ReservationControlInformationType `xml:"reservationOrProfileIdentifier,omitempty"`

	PersonalFacts *PersonalFacts `xml:"personalFacts,omitempty"`

	// Informations needed for a retreive by frequent flyer .
	FrequentFlyer *FrequentTravellerIdentificationCodeType `xml:"frequentFlyer,omitempty"`

	// Informations needed for a retreive by account number
	Accounting *AccountingInformationElementType `xml:"accounting,omitempty"`
}

type PersonalFacts struct {
	// Informations needed for retreive by office and name or retreive by service and name .
	TravellerInformation *TravellerInformationType `xml:"travellerInformation"`

	// Informations on the travel product
	ProductInformation *TravelProductInformationType `xml:"productInformation,omitempty"`

	// Ticket number that can be used  optionally when retreive by record locator is done .
	Ticket *TicketNumberType `xml:"ticket,omitempty"`
}

//
// Complex structs
//

type AccountingElementType struct {
	// account number
	Number formats.AlphaNumericString_Length1To30 `xml:"number"`
}

type AccountingInformationElementType struct {
	// accounting element needed when retreive by account number
	Account *AccountingElementType `xml:"account"`
}

type CompanyIdentificationTypeI struct {
	// Airline/ provider code
	Code formats.AlphaNumericString_Length1To3 `xml:"code"`
}

type FrequentTravellerIdentificationCodeType struct {
	// frequent traveler identification needed when retreive by frequent traveller
	FrequentTraveller *FrequentTravellerIdentificationTypeI `xml:"frequentTraveller"`
}

type FrequentTravellerIdentificationTypeI struct {
	// airline code
	CompanyId formats.AlphaNumericString_Length2To2 `xml:"companyId"`

	// frequent traveller reference number
	MembershipNumber formats.AlphaNumericString_Length1To27 `xml:"membershipNumber"`
}

type LocationTypeI struct {
	// Board point or off point
	CityCode formats.AlphaString_Length3To3 `xml:"cityCode"`
}

type OptionalPNRActionsType struct {
	// 51 - return all RTSVC 52 - return line numbers 53 - return RLRs 55 - do not return individual names of a group 101 - hard copy print 201 - return RTSVC for car 231 - return RTSVC for hotel 261 - return RTSVC for air 300 - TY mode
	OptionCode []formats.NumericInteger_Length1To3 `xml:"optionCode"` // maxOccurs="40"
}

type PrinterIdentificationDetailsType struct {
	// name of the printer
	Name formats.AlphaNumericString_Length5To6 `xml:"name"`

	// network id of the printer
	Network formats.AlphaNumericString_Length2To2 `xml:"network,omitempty"`
}

type PrinterIdentificationType struct {
	// printer identification
	IdentifierDetail *PrinterIdentificationDetailsType `xml:"identifierDetail,omitempty"`

	// amadeus office id
	Office formats.AlphaNumericString_Length9To9 `xml:"office,omitempty"`

	// IATA teletype address
	TeletypeAddress formats.AlphaNumericString_Length7To7 `xml:"teletypeAddress,omitempty"`
}

type ProductDateTimeTypeI struct {
	// Departure/ pick-up/ check-in date
	DepDate formats.Date_DDMMYY `xml:"depDate"`

	// Only for retrieve by flight/departure time, needed.
	DepTime formats.Time24_HHMM `xml:"depTime,omitempty"`

	// Return/ drop-off/ check-out date
	ArrDate formats.Date_DDMMYY `xml:"arrDate,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	// Flight number
	Identification formats.NumericInteger_Length1To4 `xml:"identification"`

	// Flight number alpha suffix
	Subtype formats.AlphaString_Length1To1 `xml:"subtype,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// profile or PNR record locator
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber"`
}

type ReservationControlInformationType struct {
	// record information
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation"`
}

type RetrievePNRType struct {
	// 1 - redisplay 2 - retrieve by record locator 3 - retrieve by office and name 4 - retrieve by service and name 5 - retrieve by frequent traveller 6 - retrieve by account number 7 - retrieve by customer profile 8 - retrieve by Insurance policy number 9 - retrieve by numeric record locator
	Type formats.NumericInteger_Length1To1 `xml:"type"`

	// information needed for redisplay if RTSVC (AIR, HTL, CAR). Needed for retrieve by service and name . Optional for retrieve by office and name
	Service formats.AlphaString_Length3To3 `xml:"service,omitempty"`

	// Element tattoo. Optional for redisplay when RTSVC.
	Tattoo formats.AlphaNumericString_Length1To5 `xml:"tattoo,omitempty"`

	// PNR owner office id
	Office formats.AlphaNumericString_Length9To9 `xml:"office,omitempty"`

	// Target system for retrieve PNR via claim. Only for retrieve by record locator and retrieve by service and name
	TargetSystem formats.AlphaNumericString_Length2To2 `xml:"targetSystem,omitempty"`

	// X - RTAXR V - RV Optional for redisplay and retrieval types 2, 3 and 4. N/A otherwise.
	Option1 formats.AlphaString_Length1To1 `xml:"option1,omitempty"`

	// A - active PNRs only Optional for retrieve by office and name. N/A otherwise.
	Option2 formats.AlphaString_Length1To1 `xml:"option2,omitempty"`
}

type TicketNumberType struct {
	// airline code
	Airline formats.NumericInteger_Length3To3 `xml:"airline"`

	// airline ticket number
	TicketNumber formats.NumericInteger_Length10To10 `xml:"ticketNumber"`
}

type TravelProductInformationType struct {
	// Defaults to current date
	Product *ProductDateTimeTypeI `xml:"product,omitempty"`

	// Needed for retrieval by flight, optional otherwise.
	BoardpointDetail *LocationTypeI `xml:"boardpointDetail,omitempty"`

	// needed only if board point is provided.
	OffpointDetail *LocationTypeI `xml:"offpointDetail,omitempty"`

	// company identification
	Company *CompanyIdentificationTypeI `xml:"company,omitempty"`

	// needed only for retrieve by flight
	ProductDetails *ProductIdentificationDetailsTypeI `xml:"productDetails,omitempty"`
}

type TravellerDetailsTypeI struct {
	// Traveler first name
	FirstName formats.AlphaString_Length1To56 `xml:"firstName"`
}

type TravellerInformationType struct {
	// traveler surname information
	Traveller *TravellerSurnameInformationTypeI `xml:"traveller"`

	// N/A for retrieve by record locator
	Passenger *TravellerDetailsTypeI `xml:"passenger,omitempty"`
}

type TravellerSurnameInformationTypeI struct {
	// traveler last name or group name
	Surname formats.AlphaString_Length1To57 `xml:"surname"`
}
