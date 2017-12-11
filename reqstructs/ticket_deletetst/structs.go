package ticket_deletetst

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-ws-go/formats"
)

type TicketDeleteTST struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTDQ_04_1_1A Ticket_DeleteTST"`

	// It can be: - 'ALL' - all TSTs are deleted; - 'SEL' - only the information corresponding to the TST/passenger selection is deleted.
	DeleteMode *CodedAttributeType `xml:"deleteMode"`

	// PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`  // minOccurs="0"

	PsaList []*PsaList `xml:"psaList,omitempty"`  // minOccurs="0" maxOccurs="1980"
}

type PsaList struct {

	// TST reference number. Order number is here meaningless.
	ItemReference *ItemReferencesAndVersionsType `xml:"itemReference"`

	// Reference information about passenger(s) to delete from the TST(s)
	PaxReference *ReferenceInformationTypeI `xml:"paxReference,omitempty"`  // minOccurs="0"
}

//
// Complex structs
//

type CodedAttributeInformationType struct {

	// provides the attribute Type
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType"`
}

type CodedAttributeType struct {

	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails"`
}

type ItemReferencesAndVersionsType struct {

	// qualifies the type of the reference used. Code set to define
	ReferenceType formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`  // minOccurs="0"

	// Tattoo number (It is in fact the Tst Display Number)
	UniqueReference *formats.NumericInteger_Length1To5 `xml:"uniqueReference,omitempty"`  // minOccurs="0"

	// Gives the TST ID number
	IDDescription *UniqueIdDescriptionType `xml:"iDDescription,omitempty"`  // minOccurs="0"
}

type ReferenceInformationTypeI struct {

	// Passenger/segment/TST reference details
	RefDetails []*ReferencingDetailsTypeI `xml:"refDetails,omitempty"`  // minOccurs="0" maxOccurs="99"
}

type ReferencingDetailsTypeI struct {

	// Qualifyer of the reference (Pax/Seg/Tst)
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"`  // minOccurs="0"

	// Passenger/segment/TST reference number
	RefNumber *formats.NumericInteger_Length1To5 `xml:"refNumber,omitempty"`  // minOccurs="0"
}

type ReservationControlInformationDetailsTypeI struct {

	// Record locator.
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {

	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation"`
}

type UniqueIdDescriptionType struct {

	// The TST Id Number : The Id number allows to determine a TST in the single manner.
	IDSequenceNumber formats.NumericInteger_Length1To11 `xml:"iDSequenceNumber"`
}
