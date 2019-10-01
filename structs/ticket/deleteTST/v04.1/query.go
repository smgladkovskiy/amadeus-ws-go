package Ticket_DeleteTST_v04_1 // ttstdq041

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName        xml.Name                            `xml:"http://xml.amadeus.com/TTSTDQ_04_1_1A Ticket_DeleteTST"`
	DeleteMode     *CodedAttributeType                 `xml:"deleteMode"`               // It can be: - 'ALL' - all TSTs are deleted; - 'SEL' - only the information corresponding to the TST/passenger selection is deleted.
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"` // PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	PsaList        []*PsaList                          `xml:"psaList,omitempty"`        // maxOccurs="1980"
}

type PsaList struct {
	ItemReference *ItemReferencesAndVersionsType `xml:"itemReference"`          // TST reference number. Order number is here meaningless.
	PaxReference  *ReferenceInformationTypeI     `xml:"paxReference,omitempty"` // Reference information about passenger(s) to delete from the TST(s)
}

//
// Complex structs
//

type CodedAttributeInformationType struct {
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType"` // provides the attribute Type
}

type CodedAttributeType struct {
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails"` // provides details for the Attribute
}

type ItemReferencesAndVersionsType struct {
	ReferenceType   formats.AlphaNumericString_Length1To3 `xml:"referenceType,omitempty"`   // qualifies the type of the reference used. Code set to define
	UniqueReference *formats.NumericInteger_Length1To5    `xml:"uniqueReference,omitempty"` // Tattoo number (It is in fact the Tst Display Number)
	IDDescription   *UniqueIdDescriptionType              `xml:"iDDescription,omitempty"`   // Gives the TST ID number
}

type ReferenceInformationTypeI struct {
	// Passenger/segment/TST reference details
	RefDetails []*ReferencingDetailsTypeI `xml:"refDetails,omitempty"` // maxOccurs="99"
}

type ReferencingDetailsTypeI struct {
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"` // Qualifyer of the reference (Pax/Seg/Tst)
	RefNumber    *formats.NumericInteger_Length1To5    `xml:"refNumber,omitempty"`    // Passenger/segment/TST reference number
}

type ReservationControlInformationTypeI struct {
	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation"`
}

type UniqueIdDescriptionType struct {
	// The TST Id Number : The Id number allows to determine a TST in the single manner.
	IDSequenceNumber formats.NumericInteger_Length1To11 `xml:"iDSequenceNumber"`
}
