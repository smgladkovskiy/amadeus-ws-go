package DocIssuance_IssueTicket_v09_1 // ttktir091

//import "encoding/xml"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TTKTIR_09_1_1A DocIssuance_IssueTicketReply"`
	ProcessingStatus *ResponseAnalysisDetailsType `xml:"processingStatus"`     // The type of the answer : warning, aknowledgement, PNR display...
	ErrorGroup       *ErrorGroupType              `xml:"errorGroup,omitempty"` // Contains warning and errors.
}

//
// Complex structs
//

type ApplicationErrorDetailType struct {
	ErrorCode      string `xml:"errorCode"`                // Code identifying the data validation error condition.
	ErrorCategory  string `xml:"errorCategory,omitempty"`  // Identification of a code list.
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"` // Code identifying the agency responsible for a code list.
}

type ApplicationErrorInformationType struct {
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"` // Application error details.
}

type ErrorGroupType struct {
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType `xml:"errorOrWarningCodeDetails"`         // The details of error/warning code.
	ErrorWarningDescription   *FreeTextInformationType         `xml:"errorWarningDescription,omitempty"` // The desciption of warning or error.
}

type FreeTextDetailsType struct {
	TextSubjectQualifier string `xml:"textSubjectQualifier"` // textSubjectQualifier
	Source               string `xml:"source"`               // source
	Encoding             string `xml:"encoding"`             // encoding
}

type FreeTextInformationType struct {
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"` // warning/error message details.
	FreeText        string               `xml:"freeText"`        // Free text and message sequence numbers of the remarks.
}

type ResponseAnalysisDetailsType struct {
	// type of the answer
	StatusCode string `xml:"statusCode"`
}
