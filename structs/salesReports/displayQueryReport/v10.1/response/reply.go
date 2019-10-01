package SalesReports_QueryReportReply_v10_1 // tsrqrr101

//import "encoding/xml"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/TSRQRR_10_1_1A SalesReports_DisplayQueryReportReply"`

	// Conveys a potential error and its associated Free flow text.
	ErrorGroup *ErrorGroupType `xml:"errorGroup,omitempty"`

	// Conveys all the details associated to the query report.
	QueryReportDataDetails *QueryReportDataDetails `xml:"queryReportDataDetails,omitempty"`
}

type QueryReportDataDetails struct {
	// Currency
	CurrencyInfo *CurrenciesTypeU `xml:"currencyInfo"`

	// Conveys the system date and the specific requested date if present.
	DateDetails []*StructuredDateTimeInformationType `xml:"dateDetails"` // maxOccurs="2"

	// Conveys by office the documents net details.
	QueryReportDataOfficeGroup []*QueryReportDataOfficeGroup `xml:"queryReportDataOfficeGroup"` // maxOccurs="9"

	// Conveys details related to the agent or user associated to the requested document.
	AgentDetails *UserIdentificationType `xml:"agentDetails,omitempty"`

	// Conveys details used for handling scrolling requests.
	ActionDetails *ActionDetailsTypeI `xml:"actionDetails,omitempty"`

	// Sales report period dates: From and To dates.
	SalesPeriodDetails *StructuredPeriodInformationType `xml:"salesPeriodDetails,omitempty"`
}

type QueryReportDataOfficeGroup struct {
	// Conveys the data related to the requestor's agency.
	RequestorAgencyDetails *AdditionalBusinessSourceInformationTypeI `xml:"requestorAgencyDetails"`

	// Conveys data related to a document.
	DocumentData []*DocumentData `xml:"documentData"` // maxOccurs="9999"
}

type DocumentData struct {
	// This conveys the sequence number of the document within the sales report.
	SequenceIdentification *ItemNumberTypeI `xml:"sequenceIdentification"`

	// Logical document number with airline code and related details.
	DocumentNumber *TicketNumberTypeI `xml:"documentNumber"`

	// Various total amounts of the document.
	MonetaryInformation *MonetaryInformationTypeI `xml:"monetaryInformation"`

	// Booking agent sine
	BookingAgent *UserIdentificationType `xml:"bookingAgent"`

	// Transaction details:  - transaction code - transaction type
	TransactionDataDetails *TransactionInformationForTicketingType `xml:"transactionDataDetails"`

	// This group conveys data related to the FOP.
	FopDetails *GeneralFopRepresentationType `xml:"fopDetails"`

	// Conveys the passenger name.
	PassengerName *TravellerInformationTypeI `xml:"passengerName,omitempty"`

	// Reservation information: RLOC of the PNR.
	ReservationInformation *ReservationControlInformationTypeI `xml:"reservationInformation,omitempty"`
}

//
// Complex structs
//

type ActionDetailsTypeI struct {
	// Conveys the number of lines to retrieve.
	NumberOfItemsDetails *ProcessingInformationTypeI `xml:"numberOfItemsDetails,omitempty"`

	// Identification of the last element retrieved.
	LastItemsDetails *ReferenceTypeI `xml:"lastItemsDetails,omitempty"`
}

type AdditionalBusinessSourceInformationTypeI struct {
	// Source type
	SourceType *SourceTypeDetailsTypeI `xml:"sourceType"`

	// details of the sales report originator.
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails"`
}

type ApplicationErrorDetailType struct {
	// Code identifying the data validation error condition.
	// xmlType: AlphaNumericString_Length1To5
	ErrorCode string `xml:"errorCode"`

	// Identification of a code list.
	// xmlType: AlphaNumericString_Length1To3
	ErrorCategory string `xml:"errorCategory,omitempty"`

	// Code identifying the agency responsible for a code list.
	// xmlType: AlphaNumericString_Length1To3
	ErrorCodeOwner string `xml:"errorCodeOwner,omitempty"`
}

type ApplicationErrorInformationType struct {
	// Application error details.
	ErrorDetails *ApplicationErrorDetailType `xml:"errorDetails"`
}

type CurrenciesTypeU struct {
	// Details of the currency.
	CurrencyDetails *CurrencyDetailsTypeU `xml:"currencyDetails,omitempty"`
}

type CurrencyDetailsTypeU struct {
	// Conveys the type of currency.
	// xmlType: AlphaNumericString_Length1To3
	CurrencyQualifier string `xml:"currencyQualifier"`

	// ISO code of the currency.
	// xmlType: AlphaNumericString_Length1To3
	CurrencyIsoCode string `xml:"currencyIsoCode,omitempty"`
}

type ErrorGroupType struct {
	// The details of error/warning code.
	ErrorOrWarningCodeDetails *ApplicationErrorInformationType `xml:"errorOrWarningCodeDetails"`

	// The desciption of warning or error.
	ErrorWarningDescription *FreeTextInformationType `xml:"errorWarningDescription,omitempty"`
}

type FormOfPaymentDetailsTypeI struct {
	// Form of payment code.
	// xmlType: AlphaNumericString_Length1To10
	Type string `xml:"type"`

	// Form of payment indicator
	// xmlType: AlphaNumericString_Length1To3
	Indicator string `xml:"indicator,omitempty"`

	// Amount paid by the form of payment.
	// xmlType: NumericDecimal_Length1To18
	Amount *float64 `xml:"amount,omitempty"`

	// Credit card company identification (VI, CA, AX...)
	// xmlType: AlphaNumericString_Length1To35
	VendorCode string `xml:"vendorCode,omitempty"`

	// Credit card number
	// xmlType: AlphaNumericString_Length1To35
	CreditCardNumber string `xml:"creditCardNumber,omitempty"`

	// Expiry date of the credit card.
	// xmlType: AlphaNumericString_Length1To35
	ExpiryDate string `xml:"expiryDate,omitempty"`

	// Approval code of the credit card.
	// xmlType: AlphaNumericString_Length1To17
	ApprovalCode string `xml:"approvalCode,omitempty"`

	// Source of the approval code: manual or automatic.
	// xmlType: AlphaNumericString_Length1To3
	SourceOfApproval string `xml:"sourceOfApproval,omitempty"`
}

type FormOfPaymentTypeI struct {
	// Form of payment details
	FormOfPayment *FormOfPaymentDetailsTypeI `xml:"formOfPayment"`
}

type FreeTextDetailsType struct {
	// Qualifier of the subject
	// xmlType: AlphaNumericString_Length1To3
	TextSubjectQualifier string `xml:"textSubjectQualifier"`

	// Type of information
	// xmlType: AlphaNumericString_Length1To4
	InformationType string `xml:"informationType,omitempty"`

	// Status
	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"`

	// Company ID
	// xmlType: AlphaNumericString_Length1To35
	CompanyId string `xml:"companyId,omitempty"`

	// Language
	// xmlType: AlphaNumericString_Length1To3
	Language string `xml:"language,omitempty"`

	// Source of the free text
	// xmlType: AlphaNumericString_Length1To3
	Source string `xml:"source"`

	// Encoding type
	// xmlType: AlphaNumericString_Length1To3
	Encoding string `xml:"encoding"`
}

type FreeTextInformationType struct {
	// Error free text details
	FreeTextDetails *FreeTextDetailsType `xml:"freeTextDetails"`

	// Free text and message sequence numbers of the remarks.
	// xmlType: AlphaNumericString_Length1To199
	FreeText []string `xml:"freeText"` // maxOccurs="99"
}

type GeneralFopRepresentationType struct {
	// Form of payment details.
	FopDescription *FormOfPaymentTypeI `xml:"fopDescription"`

	// To convey monetary amounts, currency...
	MonetaryInfo *MonetaryInformationTypeI `xml:"monetaryInfo"`
}

type ItemNumberIdentificationTypeI struct {
	// Conveys a document sequence number within the sales report.
	// xmlType: AlphaNumericString_Length1To6
	Number string `xml:"number,omitempty"`

	// Type of the number.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`
}

type ItemNumberTypeI struct {
	// Conveys the sequence number of a transaction within a sales report display.
	ItemNumberDetails *ItemNumberIdentificationTypeI `xml:"itemNumberDetails"`
}

type MonetaryInformationDetailsTypeI struct {
	// Type of the amount
	// xmlType: AlphaNumericString_Length1To3
	TypeQualifier string `xml:"typeQualifier"`

	// Amount.
	// xmlType: AlphaNumericString_Length1To35
	Amount string `xml:"amount,omitempty"`

	// Currency of the amount.
	// xmlType: AlphaNumericString_Length1To3
	Currency string `xml:"currency,omitempty"`
}

type MonetaryInformationTypeI struct {
	// Details on the amounts
	MonetaryDetails *MonetaryInformationDetailsTypeI `xml:"monetaryDetails"`

	OtherMonetaryDetails []*MonetaryInformationDetailsTypeI `xml:"otherMonetaryDetails,omitempty"` // maxOccurs="19"
}

type OriginatorIdentificationDetailsTypeI struct {
	// IATA number of the agency.
	OriginatorId int32 `xml:"originatorId"`

	// Office ID of the agency.
	// xmlType: AlphaNumericString_Length1To9
	InHouseIdentification1 string `xml:"inHouseIdentification1,omitempty"`
}

type OriginatorIdentificationDetailsTypeI_85102C struct {
	// Agent sine (Numeric Sine)+(Agent Initials)+(Duty Code) ex : 0001XVSU).
	// xmlType: AlphaNumericString_Length1To9
	OriginatorId string `xml:"originatorId,omitempty"`
}

type ProcessingInformationTypeI struct {
	// - In a query: number of rows requested. - In a reply: 0 if no more remaining rows, else empty.
	NumberOfItems *int32 `xml:"numberOfItems,omitempty"`
}

type ReferenceTypeI struct {
	// - In a reply: conveys  the key on the last item sent in case there are more items.  - In a request: when it is not the first call, conveys the last ticket key received.  This key can be a ticket number or a sales report number, depending on the type of the report.
	// xmlType: AlphaNumericString_Length1To35
	LastItemIdentifier string `xml:"lastItemIdentifier,omitempty"`
}

type ReservationControlInformationDetailsTypeI struct {
	// Reservation control number : RLoc of the PNR.
	// xmlType: AlphaNumericString_Length1To20
	ControlNumber string `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// Reservation details
	Reservation []*ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"` // maxOccurs="9"
}

type SourceTypeDetailsTypeI struct {
	// Source qualifier: reporting office.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier1 string `xml:"sourceQualifier1"`

	// Indicates that data are associated to all agencies sharing the same IATA number.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier2 string `xml:"sourceQualifier2,omitempty"`
}

type StructuredDateTimeInformationType struct {
	// This data element can be used to provide the semantic of the information provided. Examples : - Current date - Requested specific date
	// xmlType: AlphaNumericString_Length1To3
	BusinessSemantic string `xml:"businessSemantic,omitempty"`

	// Convey date and/or time.
	DateTime *StructuredDateTimeType `xml:"dateTime,omitempty"`
}

type StructuredDateTimeType struct {
	// Year number.
	//formats: Year_YYYY
	Year string `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	//formats: Month_mM
	Month string `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	//formats: Day_nN
	Day string `xml:"day,omitempty"`
}

type StructuredPeriodInformationType struct {
	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType `xml:"beginDateTime,omitempty"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType `xml:"endDateTime,omitempty"`
}

type TicketNumberDetailsTypeI struct {
	// Document number.
	// xmlType: AlphaNumericString_Length1To35
	Number string `xml:"number,omitempty"`

	// In case of conjunctiv, indicates the total number of tickets.
	NumberOfBooklets *int32 `xml:"numberOfBooklets,omitempty"`
}

type TicketNumberTypeI struct {
	// Details on the document transaction.
	DocumentDetails *TicketNumberDetailsTypeI `xml:"documentDetails"`

	// Provide status on the document in the sales report: it can be confirmed or not.
	// xmlType: AlphaNumericString_Length1To3
	Status string `xml:"status,omitempty"`
}

type TransactionInformationForTicketingType struct {
	// Reporting transaction details
	TransactionDetails *TransactionInformationsType `xml:"transactionDetails"`
}

type TransactionInformationsType struct {
	// Transaction Code, coded : CANR  MCOA  MCOM  MDnn MPnn PTAM TKTA  TKTB  TKTM  TKTT  TORM  XSBA  XSBM  ACMR   RENM  RFND  ACMA  SSAC  TAAD  ADMA  RCSM  SSAD  BPAS  CANX CANN  PSCN  VSCN  RSCN RENA TASF
	// xmlType: AlphaNumericString_Length1To4
	Code string `xml:"code,omitempty"`

	// Transaction Type, coded :     SALE     INVT     REFD      ADJA     ADJP     AUTS     CCAS     CCCS     MANS     VOID
	// xmlType: AlphaNumericString_Length1To4
	Type string `xml:"type,omitempty"`

	// 'O' --) Original transaction code 'C' --) Current transaction code
	// xmlType: AlphaNumericString_Length1To1
	IssueIndicator string `xml:"issueIndicator,omitempty"`
}

type TravellerInformationTypeI struct {
	// Conveys details on the passenger.
	PaxDetails *TravellerSurnameInformationTypeI `xml:"paxDetails"`
}

type TravellerSurnameInformationTypeI struct {
	// Passenger surname followed by the name.
	// xmlType: AlphaNumericString_Length1To70
	Surname string `xml:"surname"`

	// Qualifier
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`
}

type UserIdentificationType struct {
	// Originator Identification Details: reporting office.
	OriginIdentification *OriginatorIdentificationDetailsTypeI_85102C `xml:"originIdentification,omitempty"`

	// User Id of the requester
	// xmlType: AlphaNumericString_Length1To30
	Originator string `xml:"originator,omitempty"`
}
