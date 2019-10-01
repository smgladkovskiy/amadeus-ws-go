package Ticket_DeleteTST_v04_1 // ttstdr041

//import "encoding/xml"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTSTDR_04_1_1A Ticket_DeleteTSTReply"`
	ProcessStatus    *ResponseAnalysisDetailsTypeI       `xml:"processStatus,omitempty"` // Process status after the TST delete
	ApplicationError *ApplicationError                   `xml:"applicationError,omitempty"`
	PnrLocatorData   *ReservationControlInformationTypeI `xml:"pnrLocatorData,omitempty"` // PNR record locator information for this transaction. This PNR record locator is used for tracing purpose.
}

type ApplicationError struct {
	ApplicationErrorInfo *ApplicationErrorInformationType `xml:"applicationErrorInfo"` // General error information returned by ticketing application
	ErrorText            *InteractiveFreeTextTypeI        `xml:"errorText,omitempty"`  // Description in free flow text of the error returned by ticketing application
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

type ReservationControlInformationDetailsTypeI struct {
	// Record locator.
	ControlNumber string `xml:"controlNumber"`
}

type ResponseAnalysisDetailsTypeI struct {
	ResponseType     string `xml:"responseType"`     // Type of the response (update/cancel request)
	ProcessingStatus string `xml:"processingStatus"` // Status of the process ran.
}
