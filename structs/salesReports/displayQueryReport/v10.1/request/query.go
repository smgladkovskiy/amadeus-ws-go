package SalesReports_QueryReportRequest_v10_1 // tsrqrq101

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/TSRQRQ_10_1_1A SalesReports_DisplayQueryReport"`

	// Conveys details used for handling scrolling requests.
	ActionDetails *ActionDetailsTypeI `xml:"actionDetails,omitempty"`

	// Conveys optional details related to an agent or a user to target a specific sales report.
	AgentUserDetails *UserIdentificationType `xml:"agentUserDetails,omitempty"`

	// Conveys an optional sales report date.
	DateDetails *StructuredDateTimeInformationType `xml:"dateDetails,omitempty"`

	// Conveys an optional sales report currency to select sales data in this currency.
	CurrencyInfo *CurrenciesTypeU `xml:"currencyInfo,omitempty"`

	// Conveys optional agency reference.
	AgencyDetails *AdditionalBusinessSourceInformationTypeI `xml:"agencyDetails,omitempty"`

	// Conveys optional sales report period dates: From and To dates.
	SalesPeriodDetails *StructuredPeriodInformationType `xml:"salesPeriodDetails,omitempty"`

	// Conveys optional transaction code and/or transaction type to target given documents in a sales report.
	TransactionData []*TransactionInformationForTicketingType `xml:"transactionData,omitempty"` // maxOccurs="2"

	// Conveys optional details on the form of payment, or the credit card company name.
	FormOfPaymentDetails *FormOfPaymentTypeI `xml:"formOfPaymentDetails,omitempty"`

	// Conveys the validating carrier details.
	ValidatingCarrierDetails *TransportIdentifierType `xml:"validatingCarrierDetails,omitempty"`

	// Conveys an option to request a sales report for: - all agents in the office. - all offices sharing the same IATA number - all STP/STDO offices - a STP/STDO office - all TDO offices - a TDO office. - only net remit documents.  Some of those options are combinables.
	RequestOption *SelectionDetailsTypeI `xml:"requestOption,omitempty"`

	// Conveys optional sales indicator to target given documents in a sales report, domestic or international sales.
	SalesIndicator *StatusType `xml:"salesIndicator,omitempty"`

	// Conveys an optional sequence or document number to sort the transactions report.
	FromSequenceDocumentNumber *ItemNumberTypeI `xml:"fromSequenceDocumentNumber,omitempty"`

	// Conveys optional client ID data to target given documents in a sales report.
	AttributeInfo *CodedAttributeType `xml:"attributeInfo,omitempty"`
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

	// Details (office/ID, IATA number) of the target office
	OriginatorDetails *OriginatorIdentificationDetailsTypeI `xml:"originatorDetails,omitempty"`
}

type CodedAttributeInformationType struct {
	// provides the attribute Type
	// xmlType: AlphaNumericString_Length1To5
	AttributeType string `xml:"attributeType"`

	// provides a description for the attribute
	// xmlType: AlphaNumericString_Length1To256
	AttributeDescription string `xml:"attributeDescription,omitempty"`
}

type CodedAttributeType struct {
	// provides details for the Attribute
	AttributeDetails []*CodedAttributeInformationType `xml:"attributeDetails"` // maxOccurs="99"
}

type CompanyIdentificationTypeI struct {
	// Validating carrier code on 2 alpha characters.
	// xmlType: AlphaNumericString_Length1To3
	MarketingCompany string `xml:"marketingCompany,omitempty"`
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

type FormOfPaymentDetailsTypeI struct {
	// Form of payment type
	// xmlType: AlphaNumericString_Length1To10
	Type string `xml:"type"`

	// Credit card vendor code
	// xmlType: AlphaNumericString_Length1To2
	VendorCode string `xml:"vendorCode,omitempty"`
}

type FormOfPaymentTypeI struct {
	// Form of payment.
	FormOfPayment *FormOfPaymentDetailsTypeI `xml:"formOfPayment"`
}

type ItemNumberIdentificationTypeI struct {
	// Conveys a number which can be: - up to 6 figures: a sequence number - from 7 to 10 figures: a document number.
	// xmlType: AlphaNumericString_Length1To10
	Number string `xml:"number,omitempty"`

	// Type of the number.
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`
}

type ItemNumberTypeI struct {
	// Conveys a document number or a sequence number.
	ItemNumberDetails []*ItemNumberIdentificationTypeI `xml:"itemNumberDetails"` // maxOccurs="99"
}

type OriginatorIdentificationDetailsTypeI struct {
	// IATA number of the agency.
	OriginatorId *formats.NumericInteger_Length1To9 `xml:"originatorId,omitempty"`

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
	NumberOfItems *formats.NumericInteger_Length1To6 `xml:"numberOfItems,omitempty"`
}

type ReferenceTypeI struct {
	// - In a reply: conveys  the key on the last item sent in case there are more items.  - In a request: when it is not the first call, conveys the last ticket key received.  This key can be a ticket number or a sales report number, depending on the type of the report.
	// xmlType: AlphaNumericString_Length1To35
	LastItemIdentifier string `xml:"lastItemIdentifier,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	// Conveys an option related to the request.
	// xmlType: AlphaNumericString_Length1To3
	Option string `xml:"option"`
}

type SelectionDetailsTypeI struct {
	// Details of the option.
	SelectionDetails *SelectionDetailsInformationTypeI `xml:"selectionDetails"`

	// Details of other options, related to the office type.
	OtherSelectionDetails []*SelectionDetailsInformationTypeI `xml:"otherSelectionDetails,omitempty"` // maxOccurs="98"
}

type SourceTypeDetailsTypeI struct {
	// Source qualifier: reporting office.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier1 string `xml:"sourceQualifier1"`

	// Indicates that data are associated to all agencies sharing the same IATA number.
	// xmlType: AlphaNumericString_Length1To3
	SourceQualifier2 string `xml:"sourceQualifier2,omitempty"`
}

type StatusDetailsTypeI struct {
	// Sales indicator:  - 'DOM' for a domestic sale - 'INT' for an international sale
	// xmlType: AlphaNumericString_Length1To3
	Type string `xml:"type,omitempty"`
}

type StatusType struct {
	// Conveys details on the sales indicator.
	StatusInformation []*StatusDetailsTypeI `xml:"statusInformation"` // maxOccurs="99"
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
	Year formats.Year_YYYY `xml:"year,omitempty"`

	// Month number in the year ( begins to 1 )
	Month formats.Month_mM `xml:"month,omitempty"`

	// Day number in the month ( begins to 1 )
	Day formats.Day_nN `xml:"day,omitempty"`
}

type StructuredPeriodInformationType struct {
	// Convey the begin date/time of a period.
	BeginDateTime *StructuredDateTimeType `xml:"beginDateTime,omitempty"`

	// Convey the end date/time of a period.
	EndDateTime *StructuredDateTimeType `xml:"endDateTime,omitempty"`
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

type TransportIdentifierType struct {
	// Carrier details
	CompanyIdentification *CompanyIdentificationTypeI `xml:"companyIdentification,omitempty"`
}

type UserIdentificationType struct {
	// Originator Identification Details: reporting office.
	OriginIdentification *OriginatorIdentificationDetailsTypeI_85102C `xml:"originIdentification,omitempty"`

	// User Id of the requester
	// xmlType: AlphaNumericString_Length1To30
	Originator string `xml:"originator,omitempty"`
}
