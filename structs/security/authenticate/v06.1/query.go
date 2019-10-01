package Security_Authenticate_v06_1 // vlsslq061

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type SecurityAuthenticate struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/VLSSLQ_06_1_1A Security_Authenticate"`

	// It contains conversation properties between the SI and the JFE.
	ConversationClt *ConversationIDType `xml:"conversationClt,omitempty"`

	// This segment is dedicated to specify a user identifier. The first one  for the main logon and the second one to the delegate user if any. Either a SignId or a UserId can be used, depending on DE 9972 value (Z for Sign Id, U for User Id). The value itself is specified in DE 9904.
	UserIdentifier []*UserIdentificationType `xml:"userIdentifier"` // maxOccurs="2"

	// This segment can be used to specify an AMADEUS Duty Code for the TPF sign in. DUT must be specified in the qualifier.
	DutyCode *ReferenceInformationTypeI `xml:"dutyCode,omitempty"`

	// This segment is used to specify the system which is used by the user. The workstation Id (DE 3148) can be specified to determine the location in terms of office of the request (if no workstation Id, the source office must be specified in the first UID segment). Moreover, if a User Id is used instead of a Sign Id, the organization must be specified in DE 9906. WorkstationId is actually the Atid
	SystemDetails *SystemDetailsInfoType `xml:"systemDetails,omitempty"`

	// These segments contain the password information. Two segments can be used in case of a New password is required.
	PasswordInfo []*BinaryDataType `xml:"passwordInfo,omitempty"` // maxOccurs="2"

	FullLocation *FullLocation `xml:"fullLocation,omitempty"`

	// Conatins Baleb and INdex of application. ex: JFE 1
	ApplicationId *ApplicationType `xml:"applicationId,omitempty"`
}

type FullLocation struct {
	// Stores the location of the workstation.
	WorkstationPos *PlaceLocationIdentificationTypeU `xml:"workstationPos"`

	// Contains terminal and facility or worstation id.
	LocationInfo *TerminalLocationType `xml:"locationInfo,omitempty"`
}

//
// Complex structs
//

type ApplicationIdentificationType struct {
	// Application Label. Exemple : NGD. Label is the first part of the ApplicationId. Label is considered as an Internal Id, as it is the key of the application.
	InternalId formats.AlphaNumericString_Length1To35 `xml:"internalId"`

	// Application Index. Index is part of the ApplicationId.
	SeqNumber formats.AlphaNumericString_Length1To6 `xml:"seqNumber,omitempty"`
}

type ApplicationType struct {
	// Identification of the application.
	ApplicationDetails *ApplicationIdentificationType `xml:"applicationDetails"`
}

type BinaryDataType struct {
	// Length of the data element 114Z. The unit is given in number of binary characters (bytes).
	DataLength formats.NumericInteger_Length1To15 `xml:"dataLength"`

	// type of the data. When E is specified, this means that the password is not crypted. When B specified, this means that the password is crypted.
	DataType formats.AlphaNumericString_Length1To1 `xml:"dataType"`

	// used to store binary data
	BinaryData formats.AlphaNumericString_Length1To99999 `xml:"binaryData"`
}

type ConversationIDType struct {
	// Sender identification
	SenderIdentification formats.AlphaNumericString_Length1To35 `xml:"senderIdentification"`

	// Recipient identification
	RecipientIdentification formats.AlphaNumericString_Length1To35 `xml:"recipientIdentification"`

	// Sender's interchange control reference
	SenderInterchangeControlReference formats.AlphaNumericString_Length1To14 `xml:"senderInterchangeControlReference"`

	// Recipient's interchange control reference
	RecipientInterchangeControlReference formats.AlphaNumericString_Length1To14 `xml:"recipientInterchangeControlReference"`
}

type FacilityInformationType struct {
	// Type of the Facility, coded. E.g.: - Check-In Desk - Gate... The codeset is not yet defined.
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// Unique Reference to a Facility of a given Type in a terminal. This can be several kind of values. E.g.: - 25 (means Gate 25 when associated  Facility Type= Gate, or Check-in Desk 25 when associated Facility Type= Check-In Desk) - BAEXC: means Ba-Executive Club Lounge when associated to
	Identifier formats.AlphaNumericString_Length1To5 `xml:"identifier,omitempty"`
}

type LocationIdentificationBatchTypeU struct {
	// Airport(139) or City (227) code
	Code formats.AlphaNumericString_Length1To35 `xml:"code,omitempty"`

	// Discriminator between airport or city.
	Qualifier formats.AlphaNumericString_Length1To17 `xml:"qualifier,omitempty"`
}

type OriginatorIdentificationDetailsTypeI struct {
	// AMADEUS Office Id the workstation belongs to. Must be empty if a workstation Id is specified in the SYS segment.
	SourceOffice formats.AlphaNumericString_Length1To9 `xml:"sourceOffice"`
}

type PlaceLocationIdentificationTypeU struct {
	// Type of location
	LocationType formats.AlphaNumericString_Length1To3 `xml:"locationType"`

	// Description of the location
	LocationDescription *LocationIdentificationBatchTypeU `xml:"locationDescription,omitempty"`

	// Details on the location
	FirstLocationDetails *RelatedLocationOneIdentificationTypeU `xml:"firstLocationDetails,omitempty"`
}

type ReferenceInformationTypeI struct {
	// This composite  duty code information. For Duty code info, DUT must be specified in the Qualifier.
	DutyCodeDetails *ReferencingDetailsTypeI `xml:"dutyCodeDetails"`
}

type ReferencingDetailsTypeI struct {
	// Code determining if the data is a tree identifier, a category identifier or a parent category identifier.
	ReferenceQualifier formats.AlphaNumericString_Length1To3 `xml:"referenceQualifier"`

	// Data value.
	ReferenceIdentifier formats.AlphaNumericString_Length1To35 `xml:"referenceIdentifier"`
}

type RelatedLocationOneIdentificationTypeU struct {
	// Terminal (180) or Building(300) code
	Code formats.AlphaNumericString_Length1To25 `xml:"code,omitempty"`

	// Discriminator between airport or city.
	Qualifier formats.AlphaNumericString_Length1To17 `xml:"qualifier,omitempty"`
}

type SystemDetailsInfoType struct {
	// This field contains a workstation Identifier. It is used to retrieve the physical origin of the request (mainly for printing purposes) .
	WorkstationId formats.AlphaNumericString_Length1To35 `xml:"workstationId,omitempty"`

	// Used to specify an organization when User Id logon is used.
	OrganizationDetails *SystemDetailsTypeI `xml:"organizationDetails,omitempty"`

	// Explain what type of workstation ID is stored in data element 3148
	IdQualifier formats.AlphaNumericString_Length1To1 `xml:"idQualifier,omitempty"`
}

type SystemDetailsTypeI struct {
	// This DE is used to specify an organization Id such as BA, 1A or QF.
	OrganizationId formats.AlphaNumericString_Length1To35 `xml:"organizationId"`
}

type TerminalLocationType struct {
	// To convey information related to a specific Facility.
	FacilityDetails *FacilityInformationType `xml:"facilityDetails"`
}

type UserIdentificationType struct {
	// To specify the source office the workstation belongs to. Not used in the second repetition of the segment (if any, it will not be taken into account). Used when no workstation Id is specified in SYS segment.
	OriginIdentification *OriginatorIdentificationDetailsTypeI `xml:"originIdentification,omitempty"`

	// Used to specify which kind of info is given in DE 9900.
	OriginatorTypeCode formats.AlphaNumericString_Length1To1 `xml:"originatorTypeCode"`

	// Authority code of the requester If Sign Id : (Numeric Sine)+(Agent Initials) ex : 0001XV). If user Id : Logon User Id.
	Originator formats.AlphaNumericString_Length1To30 `xml:"originator"`
}
