package PNR_Cancel_v11_3 // pnrxcl113

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName         xml.Name                           `xml:"http://xml.amadeus.com/PNRXCL_11_3_1A PNR_Cancel"`
	ReservationInfo *ReservationControlInformationType `xml:"reservationInfo,omitempty"` // reservation control informations
	PnrActions      *OptionalPNRActionsType            `xml:"pnrActions"`                // specific actions to be processed on PNR
	CancelElements  []*CancelPNRElementType            `xml:"cancelElements,omitempty"`  // maxOccurs="4"// CANCEL PNR ELEMENT
}

//
// Complex structs
//

type CancelPNRElementType struct {
	EntryType formats.AlphaNumericString_Length1To1 `xml:"entryType"`         // E : XE    I : XI D : XD G : Name integration P : Priority line
	Element   []*ElementIdentificationType          `xml:"element,omitempty"` // maxOccurs="999"// Element data
}

type ElementIdentificationType struct {
	Identifier formats.AlphaNumericString_Length1To3 `xml:"identifier,omitempty"` // ELEMENT DATA
	Number     *formats.NumericInteger_Length1To5    `xml:"number,omitempty"`     // element number
	SubElement *formats.NumericInteger_Length1To5    `xml:"subElement,omitempty"` // sub element number
}

type OptionalPNRActionsType struct {
	// Processing options. Only the option 0, 10, 11 and 20 are supported.
	OptionCode []formats.NumericInteger_Length1To3 `xml:"optionCode"` // maxOccurs="40"
}

type ReservationControlInformationDetailsTypeI struct {
	// profile or PNR record locator
	ControlNumber formats.AlphaNumericString_Length1To20 `xml:"controlNumber"`
}

type ReservationControlInformationType struct {
	// record information
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation"`
}
