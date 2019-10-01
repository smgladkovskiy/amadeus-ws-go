package Security_Authenticate_v06_1 // vlsslr061

//import "encoding/xml"

type SecurityAuthenticateReply struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLR_06_1_1A Security_AuthenticateReply"`

	ErrorSection *ErrorSection `xml:"errorSection,omitempty"`

	// This segment is only used if process is OK. In that case P is specified.
	ProcessStatus *ResponseAnalysisDetailsType `xml:"processStatus,omitempty"`

	// This segment is used to specify organization details associated with the user.
	OrganizationInfo *OrganizationType `xml:"organizationInfo,omitempty"`

	// Identifier of a group of conversation, shared by sevreal mono-signed conversations.
	ConversationGrp *BusinessProcessIdType `xml:"conversationGrp,omitempty"`
}

type ErrorSection struct {
	// Application Error
	ApplicationError *ApplicationErrorInformationType `xml:"applicationError"`

	// Supplementary Info on the Error.
	InteractiveFreeText *InteractiveFreeTextTypeI `xml:"interactiveFreeText,omitempty"`
}

//
// Complex structs
//

type ApplicationErrorDetailType struct {
	// Code identifying the data validation error condition.
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	ErrorCategory string `xml:"errorCategory,omitempty"`

	// Code identifying the agency responsible for a code list.
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationType struct {
	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type BusinessProcessIdType struct {
	// A unique reference to identify the process/business
	ProcessIdentifier string `xml:"processIdentifier"`
}

type FreeTextQualificationTypeI struct {
	// Subject
	Subject string `xml:"subject"`

	// Info Type
	InfoType string `xml:"infoType,omitempty"`

	// Language
	Language string `xml:"language,omitempty"`
}

type InteractiveFreeTextTypeI struct {
	// Free Text Qualifier
	FreeTextQualif *FreeTextQualificationTypeI `xml:"freeTextQualif,omitempty"`

	// Free Text
	FreeText []string `xml:"freeText,omitempty"` // maxOccurs="99"
}

type OrganizationIdentificationType struct {
	// Organization label (Company Id).
	Label string `xml:"label"`
}

type OrganizationType struct {
	// This composite is used to specify an organization details
	OrganizationDetails *OrganizationIdentificationType `xml:"organizationDetails"`
}

type ResponseAnalysisDetailsType struct {
	// P must be specified when status of the process is OK.
	StatusCode string `xml:"statusCode"`
}
