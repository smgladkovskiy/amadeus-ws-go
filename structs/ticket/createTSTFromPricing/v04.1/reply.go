package Ticket_CreateTSTFromPricing_v04_1 // tautcr041

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TAUTCR_04_1_1A Ticket_CreateTSTFromPricingReply"`
	ApplicationError *ApplicationError                   `xml:"applicationError,omitempty"`
	PnrLocatorData   *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"` // PNR record locator information for this transaction. This PNR record locator is used for tracing purpose.
	TstList          []*TstList                          `xml:"tstList,omitempty"`        // maxOccurs="1980"
}

type ApplicationError struct {
	ApplicationErrorInfo *ApplicationErrorInformationType `xml:"applicationErrorInfo"` // General error information returned by ticketing application
	ErrorText            *InteractiveFreeTextTypeI        `xml:"errorText,omitempty"`  // Description in free flow text of the error returned by ticketing application
}

type TstList struct {
	TstReference   *ItemReferencesAndVersionsType `xml:"tstReference"`             // TST tattoo number created by the transaction.
	PaxInformation *ReferenceInformationTypeI     `xml:"paxInformation,omitempty"` // Reference information on passengers.
}

//
// Complex structs
//

type ApplicationErrorDetailType struct {
	ApplicationErrorCode      string `xml:"applicationErrorCode"`                // Code identifying the data validation error condition.
	CodeListQualifier         string `xml:"codeListQualifier,omitempty"`         // Identification of a code list.
	CodeListResponsibleAgency string `xml:"codeListResponsibleAgency,omitempty"` // Code identifying the agency responsible for a code list.
}

type ApplicationErrorInformationType struct {
	// Application error details.
	ApplicationErrorDetail *ApplicationErrorDetailType `xml:"applicationErrorDetail"`
}

type InteractiveFreeTextTypeI struct {
	// Free flow text describing the error
	ErrorFreeText string `xml:"errorFreeText,omitempty"`
}

type ItemReferencesAndVersionsType struct {
	ReferenceType   string                   `xml:"referenceType,omitempty"`   // qualifies the type of the reference used. Code set to define
	UniqueReference *int32                   `xml:"uniqueReference,omitempty"` // Tattoo number (It is in fact the Tst Display Number)
	IDDescription   *UniqueIdDescriptionType `xml:"iDDescription,omitempty"`   // Gives the TST ID number
}

type ReferenceInformationTypeI struct {
	// Passenger/segment/TST reference details
	RefDetails []*ReferencingDetailsTypeI `xml:"refDetails,omitempty"` // maxOccurs="99"
}

type ReferencingDetailsTypeI struct {
	RefQualifier string `xml:"refQualifier,omitempty"` // Qualifyer of the reference (Pax/Seg/Tst)
	RefNumber    *int32 `xml:"refNumber,omitempty"`    // Passenger/segment/TST reference number
}

type ReservationControlInformationDetailsTypeI struct {
	// Record locator.
	ControlNumber string `xml:"controlNumber"`
}

type ReservationControlInformationTypeI struct {
	// Reservation control information
	ReservationInformation *ReservationControlInformationDetailsTypeI `xml:"reservationInformation"`
}

type UniqueIdDescriptionType struct {
	// The TST Id Number : The Id number allows to determine a TST in the single manner.
	IDSequenceNumber int32 `xml:"iDSequenceNumber"`
}
