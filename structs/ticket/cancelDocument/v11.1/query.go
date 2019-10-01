package Ticket_CancelDocument_v11_1 // trcanq111

import (
	"encoding/xml"
	//"github.com/tmconsulting/amadeus-golang-sdk/sdk/formats"
)

type Request struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TRCANQ_11_1_1A Ticket_CancelDocument"`

	// primary ticket number of document to cancel
	DocumentNumberDetails *TicketNumberTypeI `xml:"documentNumberDetails,omitempty"`

	// up to four sequence number ranges
	SequenceNumberRanges []*ItemNumberTypeI `xml:"sequenceNumberRanges,omitempty"` // maxOccurs="20"

	// Drive specific void process
	VoidOption *StatusType `xml:"voidOption,omitempty"`

	// stock provider of the document to cancel
	StockProviderDetails *OfficeSettingsDetailsType `xml:"stockProviderDetails"`

	// office the document to cancel belongs to
	TargetOfficeDetails *AdditionalBusinessSourceInformationType `xml:"targetOfficeDetails,omitempty"`
}

//
// Complex structs
//

type AdditionalBusinessSourceInformationType struct {
	// ORIGINATOR DETAILS
	OriginatorDetails *OriginatorIdentificationDetailsType `xml:"originatorDetails"`
}

type DocumentInfoFromOfficeSettingType struct {
	// airline alphabetic code
	// xmlType: AlphaNumericString_Length1To35
	MarketIataCode string `xml:"marketIataCode,omitempty"`

	// Stock Provider Code
	// xmlType: AlphaNumericString_Length1To35
	StockProviderCode string `xml:"stockProviderCode,omitempty"`
}

type ItemNumberIdentificationTypeI struct {
	// document sequence number
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"`

	// value used: FROM or TO
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`
}

type ItemNumberTypeI struct {
	// range of sequence numbers
	ItemNumberDetails []*ItemNumberIdentificationTypeI `xml:"itemNumberDetails"` // maxOccurs="2"
}

type OfficeSettingsDetailsType struct {
	// Office settings
	OfficeSettingsDetails *DocumentInfoFromOfficeSettingType `xml:"officeSettingsDetails"`
}

type OriginatorIdentificationDetailsType struct {
	// amid number of the office
	// xmlType: AlphaNumericString_Length1To9
	InHouseIdentification2 string `xml:"inHouseIdentification2"`
}

type StatusDetailsType struct {
	// list of status/qualifiers Either His for Historical or     Crt for Current
	// xmlType: AlphaNumericString_Length1To3
	Indicator string `xml:"indicator,omitempty"`
}

type StatusType struct {
	// STATUS DETAILS
	StatusInformation *StatusDetailsType `xml:"statusInformation"`
}

type TicketNumberDetailsTypeI struct {
	// document number
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number"`
}

type TicketNumberTypeI struct {
	// document identifier
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`

	// status
	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"`
}
