package DocIssuance_IssueTicket_v09_1 // ttktiq091

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName                  xml.Name                           `xml:"http://xml.amadeus.com/TTKTIQ_09_1_1A DocIssuance_IssueTicket"`
	AgentInfo                *TicketAgentInfoTypeI              `xml:"agentInfo,omitempty"`                // agent code
	OverrideDate             *StructuredDateTimeInformationType `xml:"overrideDate,omitempty"`             // override past date TST
	Selection                []*ReferenceInfoType               `xml:"selection,omitempty"`                // maxOccurs = "3"              // Segment, TST, Coupon or line selection
	PaxSelection             []*ReferenceInformationType        `xml:"paxSelection,omitempty"`             // maxOccurs = "99" // Passenger Selection
	Stock                    *StockInformationType              `xml:"stock,omitempty"`                    // stock reference
	OptionGroup              []*OptionGroup                     `xml:"optionGroup,omitempty"`              // maxOccurs = "99"
	InfantOrAdultAssociation *TravellerInformationType          `xml:"infantOrAdultAssociation,omitempty"` // Infant or Adult passenger
	OtherCompoundOptions     []*CodedAttributeType              `xml:"otherCompoundOptions,omitempty"`     // maxOccurs = "99" // Contains other options that do take parameters.
}

type OptionGroup struct {
	Switches                *StatusTypeI                       `xml:"switches"`                          // contains a list of switches, i.e. options that do not take any parameters.
	SubCompoundOptions      []*AttributeType                   `xml:"subCompoundOptions,omitempty"`      // maxOccurs = "99"      // Contains sub-options that do take parameters.
	OverrideAlternativeDate *StructuredDateTimeInformationType `xml:"overrideAlternativeDate,omitempty"` // To override Alternative date for Invoice printing
}

//
// Complex structs
//

type AttributeInformationTypeU struct {
	AttributeType        formats.AlphaNumericString_Length1To25  `xml:"attributeType"`                  // provides the sub-option name: EMPR Email to the specified address EMPRA Email to the address specified in the APE element EMPRN Email to addresses associated to nickname FAXR Fax to Specified Number FAXA Fax to Number specified in AP element of the PNR FAXN Fax to number associated to nickname LP Issuance langage PALT Printer Options / Alternate printer PCIH Printer Options / Customize Itinerary Header PDCT Printer Options / Document Type PDPR Printer Options / Distribution Profile PDUP Printer Options / Duplicate Itinerary PMNE Printer Options / Mnemonic PNUM Printer Options / Number POFP Printer Options / Office profile
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"` // provides the sub-option value
}

type AttributeType struct {
	CriteriaDetails *AttributeInformationTypeU `xml:"criteriaDetails"`
}

type CodedAttributeInformationType struct {
	AttributeType        formats.AlphaNumericString_Length1To5   `xml:"attributeType"`                  // provides the option name
	AttributeDescription formats.AlphaNumericString_Length1To256 `xml:"attributeDescription,omitempty"` // provides a description for the attribute (option value)
}

type CodedAttributeType struct {
	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails"`
}

type InternalIDDetailsTypeI struct {
	InhouseId formats.AlphaNumericString_Length1To9 `xml:"inhouseId"` // agent code or agent sine
	Type      formats.AlphaNumericString_Length1To3 `xml:"type"`      // Agent qualifier
}

type ReferenceInfoType struct {
	ReferenceDetails []*ReferencingDetailsType `xml:"referenceDetails"` // maxOccurs="99" // Reference list (segment tattoos or PNR lines or TST number or Coupon number)
}

type ReferenceInformationType struct {
	// Used to convey the passenger tatoo or display number.
	PassengerReference *ReferencingDetailsType_108978C `xml:"passengerReference"`
}

type ReferencingDetailsType struct {
	Type  formats.AlphaNumericString_Length1To3  `xml:"type"`            // Reference qualifier
	Value formats.AlphaNumericString_Length1To35 `xml:"value,omitempty"` // Reference value
}

type ReferencingDetailsType_108978C struct {
	Type  formats.AlphaNumericString_Length1To3  `xml:"type,omitempty"` // Reference qualifier
	Value formats.AlphaNumericString_Length1To35 `xml:"value"`          // Passenger tattoo.
}

type StatusDetailsTypeI struct {
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator"`
}

type StatusTypeI struct {
	StatusDetails *StatusDetailsTypeI `xml:"statusDetails"`
}

type StockInformationType struct {
	// To convey Stock type
	StockTicketNumberDetails *StockTicketNumberDetailsType `xml:"stockTicketNumberDetails"`
}

type StockTicketNumberDetailsType struct {
	Qualifier     formats.AlphaNumericString_Length1To1  `xml:"qualifier"`               // Stock Qualifier: 'S' stock control number
	ControlNumber formats.AlphaNumericString_Length1To35 `xml:"controlNumber,omitempty"` // Stock control number
}

type StructuredDateTimeInformationType struct {
	BusinessSemantic formats.AlphaNumericString_Length1To3 `xml:"businessSemantic"`   // Date format
	DateTime         *StructuredDateTimeType               `xml:"dateTime,omitempty"` // Convey date and/or time.
}

type StructuredDateTimeType struct {
	Year  formats.Year_YYYY `xml:"year"`  // Year number.
	Month formats.Month_mM  `xml:"month"` // Month number in the year ( begins to 1 )
	Day   formats.Day_nN    `xml:"day"`   // Day number in the month ( begins to 1 )
}

type TicketAgentInfoTypeI struct {
	// agent ID details
	InternalIdDetails *InternalIDDetailsTypeI `xml:"internalIdDetails"`
}

type TravellerInformationType struct {
	// Detail if the discount is adult only or infant only
	PaxDetails *TravellerSurnameInformationType `xml:"paxDetails"`
}

type TravellerSurnameInformationType struct {
	// Customer type: A=adult IN = infant
	Type formats.AlphaNumericString_Length1To2 `xml:"type"`
}
