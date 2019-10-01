package Ticket_CreateTSTFromPricing_v04_1 // tautcq041

import (
	"encoding/xml"
)

type Request struct {
	XMLName        xml.Name                            `xml:"http://xml.amadeus.com/TAUTCQ_04_1_1A Ticket_CreateTSTFromPricing"`
	PnrLocatorData *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"` // PNR record locator information for this transaction. This PNR record locator is used for tracing purpose, no internal retrieve.
	PsaList        []*PsaList                          `xml:"psaList"`                  // maxOccurs="1980"
}

type PsaList struct {
	ItemReference *ItemReferencesAndVersionsType `xml:"itemReference"`          // Reference of the fare selected. A fare may have been calculated by Fare Quote for several passengers but there is still the possibility to create a TST only for a part of these passengers.
	PaxReference  *ReferenceInformationTypeI     `xml:"paxReference,omitempty"` // Reference information on passengers.
}
