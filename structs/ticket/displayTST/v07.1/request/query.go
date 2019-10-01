package Ticket_DisplayTSTRequest_v07_1 // ttstrq071

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName              xml.Name                            `xml:"http://xml.amadeus.com/TTSTRQ_07_1_1A"`
	DisplayMode          *CodedAttributeType                 `xml:"displayMode"`                    // It can be: - 'ALL' - all TSTs are displayed; - 'SEL' - only the information corresponding to the TST/passenger selection is displayed.
	PnrLocatorData       *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"`       // PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	ScrollingInformation *ActionDetailsTypeI                 `xml:"scrollingInformation,omitempty"` // Scrolling information.
	TstReference         []*ItemReferencesAndVersionsType    `xml:"tstReference,omitempty"`         // maxOccurs="1980"// Used to get TST information for selected TST references. As we can have 10 TST per Pax, 99 passenger per PNR, and a TST split with the Infant , the max number of TST is 1980.
	PsaInformation       *ReferenceInformationTypeI          `xml:"psaInformation,omitempty"`       // Passenger/segment references information about TST(s) to retrieve. - Passenger reference specified : all the TSTs concerning this passenger reference are returned. - Passenger/segment reference : only the TST having this passenger/segment association is returned.
}

//
// Complex structs
//

type ActionDetailsTypeI struct {
	// Information on next list of TSTs to return.
	NextListInformation *ReferenceTypeI `xml:"nextListInformation,omitempty"`
}

type CodedAttributeInformationType struct {
	// provides the attribute Type
	AttributeType formats.AlphaNumericString_Length1To3 `xml:"attributeType"`
}

type CodedAttributeType struct {
	// provides details for the Attribute
	AttributeDetails *CodedAttributeInformationType `xml:"attributeDetails"`
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

type ReferenceTypeI struct {
	RemainingInformation *formats.NumericInteger_Length1To5    `xml:"remainingInformation,omitempty"` // In case of query specifies the number of TSTs to get in reply. In case of response specifies the number of TSTs remaining.
	RemainingReference   formats.AlphaNumericString_Length1To5 `xml:"remainingReference,omitempty"`   // In case of first query specifies the value of  this field in the last reply. In case of other queries specifies the last reference returned in the previous list. In case of reply specifies the last TST reference of the list. In case of last reply the value of this field set in the first query is sent.
}

type ReferencingDetailsTypeI struct {
	RefQualifier formats.AlphaNumericString_Length1To3 `xml:"refQualifier,omitempty"` // Qualifyer of the reference (Pax/Seg/Tst)
	RefNumber    *formats.NumericInteger_Length1To5    `xml:"refNumber,omitempty"`    // Passenger/segment/TST reference number
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
