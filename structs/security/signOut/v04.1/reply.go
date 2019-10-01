package SecuritySignOut_v04_1 // vlssor041

//import "encoding/xml"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSOR_04_1_1A Security_SignOutReply"`
	ErrorSection  *ErrorSection                `xml:"errorSection,omitempty"`
	ProcessStatus *ResponseAnalysisDetailsType `xml:"processStatus,omitempty"` // This segment is only used if process is OK. In that case P is specified.
}

type ErrorSection struct {
	ApplicationError    *ApplicationErrorInformationType `xml:"applicationError"`              // Application Error
	InteractiveFreeText *InteractiveFreeTextTypeI        `xml:"interactiveFreeText,omitempty"` // Supplementary Info on the Error.
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
	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type FreeTextQualificationTypeI struct {
	TextSubjectQualifier string `xml:"textSubjectQualifier"`      // Subject
	InformationType      string `xml:"informationType,omitempty"` // Info Type
	Language             string `xml:"language,omitempty"`        // Language
}

type InteractiveFreeTextTypeI struct {
	FreeTextQualification *FreeTextQualificationTypeI `xml:"freeTextQualification,omitempty"` // Free Text Qualifier
	FreeText              []string                    `xml:"freeText,omitempty"`              // maxOccurs="99"// Free Text
}

type ResponseAnalysisDetailsType struct {
	// P must be specified when status of the process is OK.
	StatusCode string `xml:"statusCode"`
}
