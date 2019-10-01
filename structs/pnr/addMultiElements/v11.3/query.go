package PNR_AddMultiElementsRequest_v11_3 // pnradd113

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName                  xml.Name                            `xml:"http://xml.amadeus.com/PNRADD_11_3_1A PNR_AddMultiElements"`
	ReservationInfo          *ReservationControlInformationTypeI `xml:"reservationInfo,omitempty"`          // To specify a reference to a reservation
	PnrActions               *OptionalPNRActionsType             `xml:"pnrActions"`                         // To specify specific Actions to be processed on PNR
	TravellerInfo            []*TravellerInfo                    `xml:"travellerInfo,omitempty"`            // maxOccurs="100"
	OriginDestinationDetails []*OriginDestinationDetails         `xml:"originDestinationDetails,omitempty"` // maxOccurs="50"
	DataElementsMaster       *DataElementsMaster                 `xml:"dataElementsMaster,omitempty"`
}

type TravellerInfo struct {
	// To specify the PNR segments/elements references and action to apply
	ElementManagementPassenger *ElementManagementSegmentType `xml:"elementManagementPassenger"`

	PassengerData []*PassengerData `xml:"passengerData"` // maxOccurs="2"
}

type PassengerData struct {
	// To specify a traveler(s) and personal details relating to the traveler(s).  In values with 'X00' (X being any letter), 00 should be replaced by any value between 01 and 99.
	TravellerInformation *TravellerInformationTypeI `xml:"travellerInformation"`

	// Passenger date of birth (DDMMYYYY) If the passenger has an infant, not in a separate TIF, then the date is used for the infant date of birth.
	DateOfBirth *DateAndTimeInformationType `xml:"dateOfBirth,omitempty"`
}

type OriginDestinationDetails struct {
	// To convey the origin and destination of a journey
	OriginDestination *OriginAndDestinationDetailsTypeI `xml:"originDestination"`

	ItineraryInfo []*ItineraryInfo `xml:"itineraryInfo,omitempty"` // maxOccurs="99"
}

type ItineraryInfo struct {
	// To specify the PNR segments/elements references and action to apply
	ElementManagementItinerary *ElementManagementSegmentType `xml:"elementManagementItinerary"`

	AirAuxItinerary *AirAuxItinerary `xml:"airAuxItinerary,omitempty"`

	// To provide specific reference identification
	ReferenceForSegment *ReferenceInfoType `xml:"referenceForSegment,omitempty"`
}

type AirAuxItinerary struct {
	// To specify details related to a product
	TravelProduct *TravelProductInformationType `xml:"travelProduct"`

	// To specify the message type and business function
	MessageAction *MessageActionDetailsTypeI `xml:"messageAction"`

	// To indicate quantity and action required in relation to a product
	RelatedProduct *RelatedProductInformationTypeI `xml:"relatedProduct,omitempty"`

	// To specify the details for making a selection
	SelectionDetailsAir *SelectionDetailsTypeI `xml:"selectionDetailsAir,omitempty"`

	// To specify a reference to a reservation
	ReservationInfoSell *ReservationControlInformationTypeI `xml:"reservationInfoSell,omitempty"`

	// To provide free form or coded long text information.
	FreetextItinerary *LongFreeTextType `xml:"freetextItinerary,omitempty"`
}

type DataElementsMaster struct {
	// marker
	Marker1 *DummySegmentTypeI `xml:"marker1"`

	DataElementsIndiv []*DataElementsIndiv `xml:"dataElementsIndiv,omitempty"` // maxOccurs="250"
}

type DataElementsIndiv struct {
	// To specify the PNR segments/elements references and action to apply
	ElementManagementData *ElementManagementSegmentType `xml:"elementManagementData"`

	// To specify the Amadeus PNR Individual Security element Note: up to 3 ISI segments may be needed to represent the ES element and its 3 types of receiver
	PnrSecurity *IndividualPnrSecurityInformationType `xml:"pnrSecurity,omitempty"`

	// specify amadeus accounting informations
	Acccounting *AccountingInformationElementType `xml:"acccounting,omitempty"`

	// To specify different kinds of remarks
	MiscellaneousRemark *MiscellaneousRemarksType `xml:"miscellaneousRemark,omitempty"`

	// To specify special requests or services information relating to a traveller.
	ServiceRequest *SpecialRequirementsDetailsTypeI `xml:"serviceRequest,omitempty"`

	// To provide date and time details relative to flight movements
	DateAndTimeInformation *DateAndTimeInformationTypeI `xml:"dateAndTimeInformation,omitempty"`

	// To specify the Amadeus PNR Tour Code element
	TourCode *TourCodeType `xml:"tourCode,omitempty"`

	// To specify an Amadeus PNR Ticket element
	TicketElement *TicketElementType `xml:"ticketElement,omitempty"`

	// To provide free form or coded long text information.
	FreetextData *LongFreeTextType `xml:"freetextData,omitempty"`

	// To specify the way data are mapped for structured
	StructuredAddress *StructuredAddressType `xml:"structuredAddress,omitempty"`

	// To specify an Amadeus PNR Option element
	OptionElement *OptionElementType `xml:"optionElement,omitempty"`

	// To request a Hard Copy Print
	Printer *PrinterIdentificationType `xml:"printer,omitempty"`

	// The group handles Seat Request with possibly rail preferences
	SeatGroup *SeatEntityType `xml:"seatGroup,omitempty"`

	// To specify the Amadeus PNR Fare elements
	FareElement *FareElementsType `xml:"fareElement,omitempty"`

	// To specify the Amadeus PNR Fare discount element
	FareDiscount *FareDiscountElementType `xml:"fareDiscount,omitempty"`

	// To specify the Amadeus PNR Manual Document Registration element
	ManualFareDocument *ManualDocumentRegistrationType `xml:"manualFareDocument,omitempty"`

	// To specify the Amadeus PNR Commission element
	Commission *CommissionElementType `xml:"commission,omitempty"`

	// To specify the Amadeus PNR Original Issue / Issue in Exchange For element
	OriginalIssue *OriginalIssueType `xml:"originalIssue,omitempty"`

	// To convey details describing the form of payment
	FormOfPayment *FormOfPaymentTypeI `xml:"formOfPayment,omitempty"`

	// To convey additional details of the form of payment
	FopExtension []*MarketSpecificDataType `xml:"fopExtension,omitempty"` // maxOccurs="3"

	// To convey the FOP service details
	ServiceDetails []*StatusTypeI `xml:"serviceDetails,omitempty"` // maxOccurs="3"

	// To specify frequent traveler verification
	FrequentTravellerVerification *FrequentTravellerVerificationType `xml:"frequentTravellerVerification,omitempty"`

	// To specify the Amadeus PNR Ticketing Carrier Designator element
	TicketingCarrier *TicketingCarrierDesignatorType `xml:"ticketingCarrier,omitempty"`

	// To specify the Amadeus PNR Fare Print Override element
	FarePrintOverride *FarePrintOverrideType `xml:"farePrintOverride,omitempty"`

	// To convey frequent traveler program information relating to a specific traveller or group.
	FrequentTravellerData *FrequentTravellerInformationTypeU `xml:"frequentTravellerData,omitempty"`

	// To specify access level of an entity (office) to the element.
	AccessLevel *ExtendedOwnershipSecurityDetailsType `xml:"accessLevel,omitempty"`

	// To provide specific reference identification
	ReferenceForDataElement *ReferenceInfoType `xml:"referenceForDataElement,omitempty"`
}

//
// Complex structs
//

type AccountingElementType struct {
	// Account number
	Number formats.AlphaNumericString_Length1To10 `xml:"number,omitempty"`

	// Cost Number
	CostNumber formats.AlphaNumericString_Length1To60 `xml:"costNumber,omitempty"`

	// IATA company number
	CompanyNumber formats.AlphaNumericString_Length1To12 `xml:"companyNumber,omitempty"`

	// Client Reference Number
	ClientReference formats.AlphaNumericString_Length1To60 `xml:"clientReference,omitempty"`

	GSTTaxDetails formats.AlphaNumericString_Length1To109 `xml:"gSTTaxDetails,omitempty"`
}

type AccountingInformationElementType struct {
	// One of these 4 data elements is mandatory , but non in particular
	Account *AccountingElementType `xml:"account,omitempty"`

	// Number of units qualifier
	AccountNumberOfUnits formats.AlphaNumericString_Length1To3 `xml:"accountNumberOfUnits,omitempty"`
}

type ClassDetailsType struct {
	// For the booking class code.
	Code formats.AlphaNumericString_Length1To1 `xml:"code,omitempty"`

	BookingClass formats.AlphaNumericString_Length2To2 `xml:"bookingClass,omitempty"`
}

type CommissionElementType struct {
	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// Commission indicator  M, C, P, CR, PR
	Indicator formats.AlphaNumericString_Length1To2 `xml:"indicator,omitempty"`

	// Commission
	CommissionInfo *CommissionInformationType `xml:"commissionInfo"`

	// Old commission
	OldCommission *CommissionInformationType_6428C `xml:"oldCommission,omitempty"`

	// Manual capping Amount (after tag C)
	ManualCapping *formats.NumericDecimal_Length1To10 `xml:"manualCapping,omitempty"`
}

type CommissionInformationType struct {
	// Percentage (max 2 decimals)
	Percentage *formats.NumericDecimal_Length1To5 `xml:"percentage,omitempty"`

	// Amount (before tag A)
	Amount *formats.NumericDecimal_Length1To10 `xml:"amount,omitempty"`

	// VAT indicator  V
	VatIndicator formats.AlphaNumericString_Length1To2 `xml:"vatIndicator,omitempty"`

	// Net remit indicator  N
	RemitIndicator formats.AlphaNumericString_Length1To2 `xml:"remitIndicator,omitempty"`
}

type CommissionInformationType_6428C struct {
	// Percentage (max 2 decimals)
	Percentage *formats.NumericInteger_Length1To5 `xml:"percentage,omitempty"`

	// Amount (before tag A)
	Amount *formats.NumericDecimal_Length1To10 `xml:"amount,omitempty"`

	// VAT indicator  V
	VatIndicator formats.AlphaNumericString_Length1To2 `xml:"vatIndicator,omitempty"`

	// Net remit indicator  N
	RemitIndicator formats.AlphaNumericString_Length1To2 `xml:"remitIndicator,omitempty"`
}

type CompanyIdentificationType struct {
	// To specify the frequent traveller program company code
	Code formats.AlphaNumericString_Length1To3 `xml:"code"`

	// To specify the frequent traveller program partnership company  code
	PartnerCode formats.AlphaNumericString_Length1To3 `xml:"partnerCode,omitempty"`

	// To specify the frequent traveller program other partnership company  code
	OtherPartnerCode formats.AlphaNumericString_Length1To3 `xml:"otherPartnerCode,omitempty"`
}

type CompanyIdentificationTypeI struct {
	// 1. Air segment: Airline code 2. ATX segment: Airline code of the airline to take action 3. CAR segment: Airline code of the airline to take action 4. HTL segment: Airline code of the airline to take action
	Identification formats.AlphaNumericString_Length1To3 `xml:"identification"`

	// To convey a second carrier (e.g. in case of multi airline open segments)
	SecondIdentification formats.AlphaNumericString_Length1To3 `xml:"secondIdentification,omitempty"`
}

type DateAndTimeDetailsTypeI struct {
	// Miscellaneous Charge Order element : date.- SEAT ssr : date of cahnge of gauge - gourp seat SSR : date of change of gauge
	FirstDate formats.Date_DDMMYY `xml:"firstDate"`

	// Miscellaneous Charge Order element : ARC carrier code code.
	MovementType formats.AlphaNumericString_Length1To3 `xml:"movementType"`

	// Miscellaneous Charge Order element : ARC city code.
	LocationIdentification formats.AlphaNumericString_Length1To3 `xml:"locationIdentification"`
}

type DateAndTimeDetailsTypeI_56946C struct {
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Inf/Child date of birth
	Date formats.AlphaNumericString_Length1To8 `xml:"date,omitempty"`
}

type DateAndTimeInformationType struct {
	// DATE AND TIME DETAILS
	DateAndTimeDetails *DateAndTimeDetailsTypeI_56946C `xml:"dateAndTimeDetails,omitempty"`
}

type DateAndTimeInformationTypeI struct {
	// date and time details
	DateAndTime *DateAndTimeDetailsTypeI `xml:"dateAndTime"`
}

type DiscountInformationType struct {
	// Discount code
	AdjustmentReason formats.AlphaNumericString_Length1To6 `xml:"adjustmentReason,omitempty"`

	// Discount percentage
	Percentage *formats.NumericInteger_Length2To2 `xml:"percentage,omitempty"`

	// Status code
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Staff employee number
	StaffNumber formats.AlphaNumericString_Length1To7 `xml:"staffNumber,omitempty"`

	// Staff employee name
	StaffName formats.AlphaString_Length1To30 `xml:"staffName,omitempty"`
}

type DummySegmentTypeI struct {
}

type ElementManagementSegmentType struct {
	// segments/elements references - type and number
	Reference *ReferencingDetailsType `xml:"reference,omitempty"`

	// PNR segment or element name
	SegmentName formats.AlphaNumericString_Length1To3 `xml:"segmentName"`
}

type ExtendedOwnershipSecurityDetailsType struct {
	// This composite contains security data between entities or related to one entity.
	SecurityDetails []*ExtendedSecurityDetailsType `xml:"securityDetails"` // maxOccurs="5"
}

type ExtendedSecurityDetailsType struct {
	// Used to specify which kind of entities is specified. F - for family
	TypeOfEntity formats.AlphaNumericString_Length1To3 `xml:"typeOfEntity,omitempty"`

	// Used to specify the access mode regarding agreement and entities
	AccessMode formats.AlphaNumericString_Length1To1 `xml:"accessMode"`

	// Used to specify entity on which the detailed security applied. Mask is specified useing *, for instance, corporate 1A0 is specified as follows : ***1A0***.
	InhouseIdentification formats.AlphaNumericString_Length1To9 `xml:"inhouseIdentification"`
}

type FareDiscountElementType struct {
	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// To specify the discount details. Only 1 repetition must be used. The Fare Discount element cannot process multiple discounts. If you wish to enter multiple discounts for a passenger, you should enter several FD elements and associate them to the same passenger.
	Discount []*DiscountInformationType `xml:"discount,omitempty"` // maxOccurs="9"

	// Date of birth
	BirthDate formats.AlphaNumericString_Length6To6 `xml:"birthDate,omitempty"`

	// number of years or months
	NumberDetail *NumberOfUnitDetailsType `xml:"numberDetail,omitempty"`

	// company ID and reference number (should not be used)
	RpInformation *RpInformationType `xml:"rpInformation,omitempty"`

	// Customer info
	Customer *ItemDetailsType `xml:"customer,omitempty"`

	// Resident Discount applicable for discount code RM,RC,DC,BP or BI
	ResidentDisc *ItemDetailsType_64822C `xml:"residentDisc,omitempty"`
}

type FareElementsType struct {
	// E for FE - Endorsements / Restrictions element  K for FK - Shadow AIR office ID element  S for FS - Miscellaneous Ticketing Information element  Z for FZ - Miscellaneous Information element
	GeneralIndicator formats.AlphaNumericString_Length1To1 `xml:"generalIndicator,omitempty"`

	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// 1. FK element: Office identification
	OfficeId []formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"` // maxOccurs="5"

	// 1. FE element: Free text 2. FS element: Free text 3. FZ element: Free text
	FreetextLong formats.AlphaNumericString_Length1To126 `xml:"freetextLong,omitempty"`
}

type FarePrintOverrideDetailsType struct {
	// Base fare override
	BaseFare formats.AlphaNumericString_Length1To11 `xml:"baseFare"`

	// Total fare override
	TotalFare formats.AlphaNumericString_Length1To11 `xml:"totalFare,omitempty"`

	// equivalent fare override
	EquivalentFare formats.AlphaNumericString_Length1To11 `xml:"equivalentFare,omitempty"`

	// Tax amount override
	TaxAmount []formats.AlphaNumericString_Length1To11 `xml:"taxAmount,omitempty"` // maxOccurs="3"
}

type FarePrintOverrideType struct {
	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// fare override detaild
	Override *FarePrintOverrideDetailsType `xml:"override,omitempty"`
}

type FormOfPaymentDetailsTypeI struct {
	// Form(s) of payment   CA for Cash (Amadeus code CASH)  CK for Check (Amadeus code CHECK)  CC for Credit card (Amadeus code CC)  MS for Miscellaneous (Amadeus code MS)
	Identification formats.AlphaNumericString_Length1To3 `xml:"identification"`

	// Amount
	Amount *formats.NumericDecimal_Length1To9 `xml:"amount,omitempty"`

	// Credit card code
	CreditCardCode formats.AlphaNumericString_Length1To3 `xml:"creditCardCode,omitempty"`

	// Credit card account number
	AccountNumber formats.AlphaNumericString_Length1To35 `xml:"accountNumber,omitempty"`

	// Credit card expiry date
	ExpiryDate formats.Date_MMYY `xml:"expiryDate,omitempty"`

	// Credit card approval code
	ApprovalCode formats.AlphaNumericString_Length1To8 `xml:"approvalCode,omitempty"`

	// Customer account number
	CustomerAccountNumber formats.AlphaNumericString_Length1To10 `xml:"customerAccountNumber,omitempty"`

	// Extended credit payment request  E for extended payment
	PaymentTimeReference formats.AlphaNumericString_Length1To3 `xml:"paymentTimeReference,omitempty"`

	// free text
	Freetext formats.AlphaNumericString_Length1To70 `xml:"freetext,omitempty"`

	// Currency (as there is no place elsewhere)
	CurrencyCode formats.AlphaNumericString_Length1To3 `xml:"currencyCode,omitempty"`
}

type FormOfPaymentTypeI struct {
	// form of payment details
	Fop []*FormOfPaymentDetailsTypeI `xml:"fop"` // maxOccurs="3"
}

type FormatedTourCodeType struct {
	// Tour type  IT, BT
	ProductId formats.AlphaNumericString_Length1To2 `xml:"productId,omitempty"`

	// Last digit of year in which tour code became effective
	Year *formats.NumericInteger_Length1To1 `xml:"year,omitempty"`

	// Airline code of sponsor
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// Conference area approval code
	ApprovalCode formats.AlphaNumericString_Length1To1 `xml:"approvalCode,omitempty"`

	// Tour identification
	PartyId formats.AlphaNumericString_Length1To8 `xml:"partyId,omitempty"`
}

type FreeFormatTourCodeType struct {
	// Free format indicator  FF
	Indicator formats.AlphaNumericString_Length1To2 `xml:"indicator,omitempty"`

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To14 `xml:"freetext,omitempty"`
}

type FreeTextQualificationType struct {
	// Identifies whether the free text is coded or not coded  3 for Literal text
	SubjectQualifier formats.AlphaNumericString_Length1To3 `xml:"subjectQualifier"`

	// 1. AP element: 2. AQ element: 3. OS element
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// Transmittable/non-transmittable indicator (S or X). Codeset list not applicable.
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Airline or system code
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`
}

type FrequentTravellerIdentificationTypeU struct {
	// Airline Code
	CompanyId formats.AlphaNumericString_Length2To3 `xml:"companyId"`

	// Frequent Traveller Number
	MembershipNumber formats.AlphaNumericString_Length1To25 `xml:"membershipNumber"`
}

type FrequentTravellerInformationTypeU struct {
	// frequent traveller identification
	FrequentTraveller *FrequentTravellerIdentificationTypeU `xml:"frequentTraveller"`
}

type FrequentTravellerVerificationType struct {
	// A  Add FT name  N  Name verification  R   Redeem miles
	ActionRequest formats.AlphaNumericString_Length1To3 `xml:"actionRequest"`

	// company identification
	Company *CompanyIdentificationType `xml:"company,omitempty"`

	// frequent flyer number
	Account *ProductAccountDetailsType `xml:"account"`
}

type IndividualPnrSecurityInformationType struct {
	// Returned before EOT or after retrieve by RTE
	Security []*IndividualSecurityType `xml:"security,omitempty"` // maxOccurs="5"

	// Returned when retrieved
	SecurityInfo *SecurityInformationType `xml:"securityInfo,omitempty"`

	// Code as in the display:  G for Amadeus Global Core Office Identification  I for IATA number  P for Pseudo-Office Identification  Default is G.
	Indicator formats.AlphaNumericString_Length1To1 `xml:"indicator,omitempty"`
}

type IndividualSecurityType struct {
	// Type of receiver G:  Type of receiver I:  Type of receiver P:
	Identification formats.AlphaNumericString_Length1To14 `xml:"identification"`

	// R for Read  B for Both read and write  N for None
	AccessMode formats.AlphaString_Length1To1 `xml:"accessMode"`
}

type IssueInformationType struct {
	// Numeric airline code
	AirlineCode formats.AlphaNumericString_Length3To3 `xml:"airlineCode"`

	// {Original issue , exchange for} document number
	DocumentNumber formats.AlphaNumericString_Length1To10 `xml:"documentNumber"`

	// {Original issue , exchange for} document check digit
	DocumentCD *formats.NumericInteger_Length1To1 `xml:"documentCD,omitempty"`

	// 1st coupon number
	Coupon1 formats.AlphaNumericString_Length1To1 `xml:"coupon1,omitempty"`

	// 2nd coupon number
	Coupon2 formats.AlphaNumericString_Length1To1 `xml:"coupon2,omitempty"`

	// 3rd coupon number
	Coupon3 formats.AlphaNumericString_Length1To1 `xml:"coupon3,omitempty"`

	// 4th coupon number
	Coupon4 formats.AlphaNumericString_Length1To1 `xml:"coupon4,omitempty"`

	// Last 2 digits of the last conjunction document of the {original issue , exchange for} document
	LastConjunction *formats.NumericInteger_Length2To2 `xml:"lastConjunction,omitempty"`

	// Last conjunction document of the {original issue , exchange for} document check digit
	ExchangeDocumentCD *formats.NumericInteger_Length1To1 `xml:"exchangeDocumentCD,omitempty"`

	// 1st last conjunction document of the {original issue , exchange for} document coupon number
	LastConjunction1 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction1,omitempty"`

	// 2nd last conjunction document of the {original issue , exchange for} document coupon number
	LastConjunction2 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction2,omitempty"`

	// 3rd last conjunction document of the {original issue , exchange for} document 4th coupon number
	LastConjunction3 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction3,omitempty"`

	// 4th last conjunction document of the {original issue , exchange for} document 1st coupon number
	LastConjunction4 formats.AlphaNumericString_Length1To1 `xml:"lastConjunction4,omitempty"`

	// City code of the issue
	CityCode formats.AlphaString_Length3To3 `xml:"cityCode"`

	// Date of the {original , new} issue
	DateOfIssue formats.Date_DDMMYY `xml:"dateOfIssue"`

	// IATA number
	IataNumber formats.AlphaNumericString_Length1To9 `xml:"iataNumber,omitempty"`

	// Currency
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// Amount
	Amount *formats.NumericDecimal_Length1To10 `xml:"amount,omitempty"`
}

type ItemDetailsType struct {
	// Only applicable for some discount codes
	CompanyId formats.AlphaString_Length2To2 `xml:"companyId,omitempty"`

	// Only applicable for some discount codes
	CardType *formats.NumericInteger_Length1To1 `xml:"cardType,omitempty"`

	// Only applicable for some discount codes
	CardNumber *formats.NumericInteger_Length1To13 `xml:"cardNumber,omitempty"`

	// for PC - alpha value for other - numeric value
	CardCheck formats.AlphaNumericString_Length1To1 `xml:"cardCheck,omitempty"`

	// for PC only
	Owner *formats.NumericInteger_Length2To2 `xml:"owner,omitempty"`

	// for PC only
	Version *formats.NumericInteger_Length1To1 `xml:"version,omitempty"`
}

type ItemDetailsType_64822C struct {
	// DN, TR,GR,AM,CR,MR
	IdCardCode formats.AlphaString_Length2To2 `xml:"idCardCode,omitempty"`

	// T - used for TR resident discount only
	IdCardType formats.AlphaString_Length1To1 `xml:"idCardType,omitempty"`

	// Card Number
	CardNumber *formats.NumericInteger_Length1To13 `xml:"cardNumber,omitempty"`

	// Card alpha check
	AlphaCheck formats.AlphaNumericString_Length1To1 `xml:"alphaCheck,omitempty"`

	// Zip code
	ZipCode formats.AlphaNumericString_Length1To9 `xml:"zipCode,omitempty"`

	// Credential certificate number
	CertificateNumber formats.AlphaNumericString_Length1To20 `xml:"certificateNumber,omitempty"`
}

type LocationTypeI struct {
	// 1. Air segment: Boarding point 2. ATX segment: Boarding point 3. CAR segment: Pick-up point city 4. HTL segment: Check-in city 5. MIS segment: City code
	CityCode formats.AlphaNumericString_Length1To5 `xml:"cityCode"`

	// 1. TRN SNCF segment: Off point city name
	CityName formats.AlphaNumericString_Length1To17 `xml:"cityName,omitempty"`
}

type LongFreeTextType struct {
	// free text details
	FreetextDetail *FreeTextQualificationType `xml:"freetextDetail,omitempty"`

	// Long free text information
	LongFreetext formats.AlphaNumericString_Length1To199 `xml:"longFreetext,omitempty"`
}

type ManualDocumentRegistrationType struct {
	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// documentation details
	Document *ManualDocumentType `xml:"document,omitempty"`

	// Free text
	Freeflow formats.AlphaNumericString_Length1To49 `xml:"freeflow,omitempty"`
}

type ManualDocumentType struct {
	// Numeric airline code
	CompanyId formats.NumericString_Length3To3 `xml:"companyId"`

	// Ticket number
	TicketNumber formats.NumericString_Length10To10 `xml:"ticketNumber"`

	// Ticket number check digit
	TicketNumberCd *formats.NumericInteger_Length1To1 `xml:"ticketNumberCd,omitempty"`

	// Last conjunction ticket number
	LastConjuction *formats.NumericInteger_Length2To2 `xml:"lastConjuction,omitempty"`

	// Last conjunction ticket number check digit
	LastConjuctionCD *formats.NumericInteger_Length1To1 `xml:"lastConjuctionCD,omitempty"`
}

type MarketSpecificDataDetailsType struct {
	// Credit Card Verification data (*CV data)
	CvData formats.AlphaNumericString_Length1To10 `xml:"cvData,omitempty"`

	// Printed and reported freeflow
	PrintedFreeflow formats.AlphaNumericString_Length1To70 `xml:"printedFreeflow,omitempty"`

	// Reported but not printed freeflow
	ReportedFreeflow formats.AlphaNumericString_Length1To70 `xml:"reportedFreeflow,omitempty"`

	// Credit Card ONO data.
	OnoData formats.AlphaNumericString_Length1To15 `xml:"onoData,omitempty"`

	// Credit Card GWT data
	GwtData formats.AlphaNumericString_Length1To15 `xml:"gwtData,omitempty"`

	// Credit Card Holder name.
	ChdData formats.AlphaNumericString_Length1To40 `xml:"chdData,omitempty"`

	// Delegation code.
	DelegationCode formats.AlphaNumericString_Length1To3 `xml:"delegationCode,omitempty"`

	// MCO Fop Document Number
	McoDocNumber formats.AlphaNumericString_Length1To13 `xml:"mcoDocNumber,omitempty"`

	// MCO Fop Coupon Number
	McoCouponNumber formats.AlphaNumericString_Length1To4 `xml:"mcoCouponNumber,omitempty"`

	// MCO Fop Iata Number
	McoIataNumber *formats.NumericInteger_Length1To9 `xml:"mcoIataNumber,omitempty"`

	// MCO Fop Place of Issue
	McoPlaceOfIssue formats.AlphaNumericString_Length1To3 `xml:"mcoPlaceOfIssue,omitempty"`

	// MCO Fop date of Issue. DDMMMYY
	McoDateOfIssue formats.AlphaNumericString_Length7To7 `xml:"mcoDateOfIssue,omitempty"`

	// Standard Fop Iata Number
	IataNumber *formats.NumericInteger_Length1To9 `xml:"iataNumber,omitempty"`
}

type MarketSpecificDataType struct {
	// FOP detail Sequence Number.
	FopSequenceNumber formats.NumericInteger_Length1To2 `xml:"fopSequenceNumber"`

	// Passenger type: PAX or INF.
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// Form of Payment Market Specific Data
	NewFopsDetails *MarketSpecificDataDetailsType `xml:"newFopsDetails,omitempty"`

	// To provide extended FOP details.
	ExtFOP []*ReferencingDetailsTypeI `xml:"extFOP,omitempty"` // maxOccurs="99"
}

type MessageActionDetailsTypeI struct {
	// type of segment
	Business *MessageFunctionBusinessDetailsTypeI `xml:"business"`
}

type MessageFunctionBusinessDetailsTypeI struct {
	// 1. Air segment:
	Function formats.AlphaNumericString_Length1To3 `xml:"function"`
}

type MiscellaneousRemarkType struct {
	// RC for confidential remark  RI for invoice remark  RM for miscellaneous remark  RQ for quality control remark
	Type formats.AlphaNumericString_Length1To3 `xml:"type"`

	// This is the 3rd character (x) of the remark title RIx or RMx, or 2 letter code for RMxx, conditional for RM, not applicable for RC and RQ
	Category formats.AlphaNumericString_Length1To2 `xml:"category,omitempty"`

	// Free text and message sequence numbers of the remarks.
	Freetext formats.AlphaNumericString_Length1To127 `xml:"freetext,omitempty"`

	// Provider type (element RIA):  1 for Air provider   2 for Car provider (CCR)  3 for Hotel Provider (HHL)  M for Miscellaneous
	ProviderType formats.AlphaNumericString_Length1To3 `xml:"providerType,omitempty"`

	// MCO element : Currency code
	Currency formats.AlphaNumericString_Length1To3 `xml:"currency,omitempty"`

	// MCO element : total fee amount
	Amount *formats.NumericDecimal_Length1To11 `xml:"amount,omitempty"`

	// Office Id (confidential remark RC)
	OfficeId []formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"` // maxOccurs="5"
}

type MiscellaneousRemarksType struct {
	// miscellaneous remarks
	Remarks *MiscellaneousRemarkType `xml:"remarks,omitempty"`
}

type NetRemitTourCodeType struct {
	// Net remit indicator  N
	Indicator formats.AlphaNumericString_Length1To2 `xml:"indicator,omitempty"`

	// Free flow text of next remit
	Freetext formats.AlphaNumericString_Length1To14 `xml:"freetext,omitempty"`
}

type NumberOfUnitDetailsType struct {
	// Number of  years or months
	Units formats.NumericInteger_Length1To2 `xml:"units"`

	// YRS for year  MTH for month
	UnitsQualifier formats.AlphaNumericString_Length1To3 `xml:"unitsQualifier"`
}

type OptionElementInformationType struct {
	// Office Id
	OfficeId formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`

	// Date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// Queue number
	Queue *formats.NumericInteger_Length1To3 `xml:"queue,omitempty"`

	// Category number
	Category *formats.NumericInteger_Length1To3 `xml:"category,omitempty"`

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To200 `xml:"freetext,omitempty"`
}

type OptionElementType struct {
	// queuing option details
	OptionDetail *OptionElementInformationType `xml:"optionDetail,omitempty"`
}

type OptionalPNRActionsType struct {
	// 0       No Special Processing 10 - 49  PNR processing options 50 - 99  PNRACC options 100 - 149  Other Opt 150 - 199  Error Processing 200 - 229  Car Opt 230 - 259  Hotel Opt 260 - 299  Air + AUX Opt 300 - 329  Ticketing Opt
	OptionCode []formats.NumericInteger_Length1To3 `xml:"optionCode"` // maxOccurs="40"
}

type OriginAndDestinationDetailsTypeI struct {
	// Airport/city code of  Origin In a Client request message, a non-blank ODI is used in an air sell request to advise that the following segments (TVL etc...) are connected. There is a maximum of 6 TVLs following a non-blank ODI.
	Origin formats.AlphaString_Length3To3 `xml:"origin,omitempty"`

	// Airport/city code of  Destination
	Destination formats.AlphaString_Length3To3 `xml:"destination,omitempty"`
}

type OriginalIssueType struct {
	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// 8 for Voucher/Certificate indicator   RN for RN indicator
	VoucherIndicator formats.AlphaNumericString_Length1To2 `xml:"voucherIndicator,omitempty"`

	// 1st occurrence for original issue 2 occurrences for issues in exchange
	Issue []*IssueInformationType `xml:"issue"` // maxOccurs="3"

	// Base fare amount
	BaseFare *formats.NumericDecimal_Length1To10 `xml:"baseFare,omitempty"`

	// Total tax amount
	TotalTax *formats.NumericDecimal_Length1To10 `xml:"totalTax,omitempty"`

	// Penalty amount
	Penalty *formats.NumericDecimal_Length1To10 `xml:"penalty,omitempty"`

	// Free flow text
	Freeflow formats.AlphaNumericString_Length1To126 `xml:"freeflow,omitempty"`
}

type PrinterIdentificationDetailsType struct {
	// name of the printer
	Name formats.AlphaNumericString_Length5To6 `xml:"name"`

	// netwrok ID of the printer
	Network formats.AlphaNumericString_Length2To2 `xml:"network,omitempty"`
}

type PrinterIdentificationType struct {
	// identification details
	IdentifierDetail *PrinterIdentificationDetailsType `xml:"identifierDetail,omitempty"`

	// 1A office id
	Office formats.AlphaNumericString_Length9To9 `xml:"office,omitempty"`

	// IATA teletype address
	TeletypeAddress formats.AlphaNumericString_Length7To7 `xml:"teletypeAddress,omitempty"`
}

type ProductAccountDetailsType struct {
	// To specify the product/account number qualifier.
	NumberQualifier formats.AlphaNumericString_Length1To3 `xml:"numberQualifier,omitempty"`

	// A code to identify a frequent traveller (e.g. a frequent traveller number)
	Number formats.AlphaNumericString_Length1To27 `xml:"number"`
}

type ProductDateTimeTypeI struct {
	// 1. Air segment: Departure date 2. ATX segment: Requested date 3. CAR segment: Pick-up date 4. HTL segment: Check-in date 5. MIS segment: Date for service required
	DepDate formats.Date_DDMMYY `xml:"depDate"`

	// 1. Air segment Departure time 2. SUR segment: Pick-up time 3. TRN Amtrack segment: Departure time 4. TRN SNCF segment: Departure time
	DepTime formats.Time24_HHMM `xml:"depTime,omitempty"`

	// 1. Air segment  Arrival date (not in the display) 2. CAR segment Drop-off date 3. HTL segment: Check-out date 4. TTO segment: Return date of the Tour
	ArrDate formats.Date_DDMMYY `xml:"arrDate,omitempty"`

	// 1. Air segment  Arrival time 2. TRN Amtrack segment: Arrival time 3. TRN SNCF segment: Arrival time
	ArrTime formats.Time24_HHMM `xml:"arrTime,omitempty"`
}

type ProductIdentificationDetailsTypeI struct {
	// 1. Air segment: Flight number or  OPEN  ARNK - air segment arrival unknown 2. CAR se 3. segment: Car type 4. SUR segment: Transportation type  (refer to VGTVD transaction)
	Identification formats.AlphaNumericString_Length1To6 `xml:"identification"`

	// 1. Air segment: Class of service 2. TRN Amtrack segment: Class of service 3. TRN SNCF segment: Class of service
	ClassOfService formats.AlphaString_Length1To1 `xml:"classOfService,omitempty"`

	// 1. Air segment  Flight number alpha suffix  A, B, C, D, E 2. SUR segment: Departure code  A or D 3. TRN SNCF segment: Train type
	Subtype formats.AlphaString_Length1To1 `xml:"subtype,omitempty"`

	// 1. Air segment:  N for Night class 2. TRN Amtrack segment:  N for Night class 3. TRN SNCF segment:  N for Night class
	Description formats.AlphaNumericString_Length1To1 `xml:"description,omitempty"`
}

type ProductTypeDetailsType struct {
	// Used to convey availibility context.
	FlightIndicator formats.AlphaNumericString_Length1To2 `xml:"flightIndicator"`
}

type RailSeatConfigurationType struct {
	// Seat space.
	SeatSpace formats.AlphaNumericString_Length2To2 `xml:"seatSpace,omitempty"`

	// Coach type.
	CoachType formats.AlphaNumericString_Length2To2 `xml:"coachType,omitempty"`

	// Seat equipment.
	SeatEquipment formats.AlphaNumericString_Length2To2 `xml:"seatEquipment,omitempty"`

	// Seat position.
	SeatPosition formats.AlphaNumericString_Length1To1 `xml:"seatPosition,omitempty"`

	// Seat direction.
	SeatDirection formats.AlphaNumericString_Length1To1 `xml:"seatDirection,omitempty"`

	// Seat deck.
	SeatDeck formats.AlphaNumericString_Length1To1 `xml:"seatDeck,omitempty"`

	// Special passenger information.
	SpecialPassengerType []formats.AlphaNumericString_Length1To1 `xml:"specialPassengerType,omitempty"` // maxOccurs="2"
}

type RailSeatPreferencesType struct {
	// Selection of the type of seat request.
	SeatRequestFunction formats.AlphaNumericString_Length1To1 `xml:"seatRequestFunction,omitempty"`

	// Seat smoking zone indicator.
	SmokingIndicator formats.AlphaString_Length1To1 `xml:"smokingIndicator,omitempty"`

	// Seat class details.
	ClassDetails *ClassDetailsType `xml:"classDetails,omitempty"`

	// Seat configuration details.
	SeatConfiguration *RailSeatConfigurationType `xml:"seatConfiguration,omitempty"`
}

type RailSeatReferenceInformationType struct {
	// Rail seat reference information.
	RailSeatReferenceDetails *SeatReferenceInformationType `xml:"railSeatReferenceDetails,omitempty"`
}

type ReferenceInfoType struct {
	// This composite is used to transmit association information
	Reference []*ReferencingDetailsType `xml:"reference,omitempty"` // maxOccurs="198"
}

type ReferencingDetailsType struct {
	// Amadeus codes are used here.  PT for Passenger Tatoo reference number  PR for Passenger Client- request-message-defined reference number  ST for Segment Tatoo reference number  SR for Segment Client- request-message-defined reference number
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier"`

	// refers to a PNR segment/element that had this number in its related EMS segment in the same message (qualifier PT, ST)
	Number formats.AlphaNumericString_Length1To5 `xml:"number"`
}

type ReferencingDetailsTypeI struct {
	// Reference qualifier
	ReferenceQualifier formats.AlphaNumericString_Length1To10 `xml:"referenceQualifier"`

	// Reference value.
	DataValue formats.AlphaNumericString_Length1To35 `xml:"dataValue,omitempty"`
}

type RelatedProductInformationTypeI struct {
	// 1. Air segment: Number of passengers 2. ATX segment: Number of passengers 3. CAR segment: Number of cars 4. CCR segment: Number of cars 5. HHL segment: Number of  rooms 6. HTL segment: Number of  rooms
	Quantity *formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`

	// status
	Status formats.AlphaString_Length1To2 `xml:"status"`
}

type ReservationControlInformationDetailsTypeI struct {
	// 1A
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// 1. RR element: Record locator of the original PNR  2. Record locator information: Airline record locator 3. Profile record locator information: Profile record locator 4. Air segment: Passive segment airline record locator  Due to technical limitations, RCI for air segment is truncated to 7 characters.
	ControlNumber formats.AlphaNumericString_Length1To19 `xml:"controlNumber,omitempty"`
}

type ReservationControlInformationTypeI struct {
	// reservation control information - i.e. record locator
	Reservation *ReservationControlInformationDetailsTypeI `xml:"reservation,omitempty"`
}

type RpInformationType struct {
	// Airline code (should not be used)
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// Numeric value (should not be used)
	ReferenceNumber *formats.NumericInteger_Length1To12 `xml:"referenceNumber,omitempty"`
}

type SeatEntityType struct {
	// To make Seat requests on flights within the PNR
	SeatRequest *SeatRequestType `xml:"seatRequest"`

	// Used to convey specific seat details relative to Train for a specific request or the "near-to" seat details for a "next-to" request.
	RailSeatReferenceInformation []*RailSeatReferenceInformationType `xml:"railSeatReferenceInformation,omitempty"` // maxOccurs="9"

	// Rail Seat Preferences
	RailSeatPreferences *RailSeatPreferencesType `xml:"railSeatPreferences,omitempty"`
}

type SeatReferenceInformationType struct {
	// Coach number.
	CoachNumber formats.AlphaNumericString_Length1To3 `xml:"coachNumber,omitempty"`

	// Deck number.
	DeckNumber formats.AlphaNumericString_Length1To3 `xml:"deckNumber,omitempty"`

	// Seat number.
	SeatNumber formats.AlphaNumericString_Length1To3 `xml:"seatNumber,omitempty"`
}

type SeatRequestType struct {
	// seat requirement details
	Seat *SeatRequierementsType `xml:"seat,omitempty"`

	// Maximum 99 for Group Seat request
	Special []*SeatRequierementsDataType `xml:"special,omitempty"` // maxOccurs="99"
}

type SeatRequierementsDataType struct {
	// Seat number + row (seat request)  Number of seats (Group seat request)
	Data formats.AlphaNumericString_Length1To4 `xml:"data,omitempty"`

	// 3 occurrences may be used for in Amadeus seat request to indicate:  1. smoking/no smoking  2. 1st area preference  2nd area preference or passenger type
	SeatType []formats.AlphaNumericString_Length1To2 `xml:"seatType,omitempty"` // maxOccurs="3"
}

type SeatRequierementsType struct {
	// G for group
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// Type of Seat requested. S - Smoking (SMST) N - No Smoking (NSST) RQST
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// Board point
	Boardpoint formats.AlphaString_Length3To3 `xml:"boardpoint,omitempty"`

	// Off point
	Offpoint formats.AlphaString_Length3To3 `xml:"offpoint,omitempty"`
}

type SecurityInformationType struct {
	// Date of creation
	CreationDate formats.AlphaNumericString_Length6To6 `xml:"creationDate"`

	// Agent initials and duty code as in ORG (eg: AASU)
	AgentCode formats.AlphaNumericString_Length4To4 `xml:"agentCode"`

	// Office Id of creation/update
	OfficeId formats.AlphaNumericString_Length9To9 `xml:"officeId,omitempty"`
}

type SelectionDetailsInformationTypeI struct {
	// See comment below
	Option formats.AlphaNumericString_Length1To3 `xml:"option"`
}

type SelectionDetailsTypeI struct {
	// level of sell to be processed
	Selection []*SelectionDetailsInformationTypeI `xml:"selection"` // maxOccurs="10"
}

type SpecialRequirementsDataDetailsTypeI struct {
	// Seat number + row (seat SSR)  Number of seats (Group seat SSR)
	Data formats.AlphaNumericString_Length1To4 `xml:"data,omitempty"`

	// 3 occurrences may be used for in Amadeus seat SSR to indicate: 1. smoking/no smoking  2. 1st area preference 3. 2nd area preference or passenger type
	SeatType []formats.AlphaNumericString_Length1To2 `xml:"seatType,omitempty"` // maxOccurs="3"
}

type SpecialRequirementsDetailsTypeI struct {
	// special requirement type details
	Ssr *SpecialRequirementsTypeDetailsTypeI `xml:"ssr"`

	// Group seat SSR cannot ask for specific seats but only smoking and/or non-smoking (see Group seat SSR). Therefore the maximum repetitions here is 9 seats (1 per passenger of non-group PNR).
	Ssrb []*SpecialRequirementsDataDetailsTypeI `xml:"ssrb,omitempty"` // maxOccurs="9"
}

type SpecialRequirementsTypeDetailsTypeI struct {
	// - ATA/IATA defined Special Service Requirement code.  (refer to IATA AIRIMP documentation) - SK element : Keyword
	Type formats.AlphaNumericString_Length1To4 `xml:"type,omitempty"`

	// ATA/IATA status code Codeset list not applicable.
	Status formats.AlphaNumericString_Length1To3 `xml:"status,omitempty"`

	// Number of services requested
	Quantity *formats.NumericInteger_Length1To3 `xml:"quantity,omitempty"`

	// Airline code or YY
	CompanyId formats.AlphaNumericString_Length1To3 `xml:"companyId,omitempty"`

	// Processing indicator, coded. - Normal SSR   P01 request for SSR explosion at EOT ...
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`

	// Board point
	Boardpoint formats.AlphaString_Length3To3 `xml:"boardpoint,omitempty"`

	// Off point
	Offpoint formats.AlphaString_Length3To3 `xml:"offpoint,omitempty"`

	// free text data
	Freetext []formats.AlphaNumericString_Length1To70 `xml:"freetext,omitempty"` // maxOccurs="2"
}

type StatusDetailsTypeI struct {
	// Service indicator SV : Service Fee
	Indicator formats.AlphaNumericString_Length1To3 `xml:"indicator,omitempty"`
}

type StatusTypeI struct {
	// To convey the status details
	StatusDetails *StatusDetailsTypeI `xml:"statusDetails"`
}

type StructuredAddressInformationType struct {
	// A1 for Address line 1
	OptionA1 formats.AlphaNumericString_Length1To2 `xml:"optionA1"`

	// A1 50 char
	OptionTextA1 formats.AlphaNumericString_Length1To50 `xml:"optionTextA1"`
}

type StructuredAddressInformationType_5063C struct {
	// CY for company - NA for name - A2 for addr line - PO for P.O box-ZP for postacl code - CI for city - ST for state-CO for country
	Option formats.AlphaNumericString_Length1To2 `xml:"option"`

	// CY-NA-CI -  30char : A2 - 50 char: ST-CT- 25 char : PO 8 char - ZP 9 char.
	OptionText formats.AlphaNumericString_Length1To50 `xml:"optionText"`
}

type StructuredAddressType struct {
	// Information type, coded  2  for billing address  P08  for general mailing address  P19  for miscellaneous mailing address  P24  for home mailing address  P25  for delivery mailing address
	InformationType formats.AlphaNumericString_Length1To4 `xml:"informationType,omitempty"`

	// address line
	Address *StructuredAddressInformationType `xml:"address"`

	// upto 8 possible address options
	OptionalData []*StructuredAddressInformationType_5063C `xml:"optionalData,omitempty"` // maxOccurs="8"
}

type TicketElementType struct {
	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// general ticketing information
	Ticket *TicketInformationType `xml:"ticket"`

	// Print options (//print options after double slash)
	PrintOptions formats.AlphaNumericString_Length1To127 `xml:"printOptions,omitempty"`
}

type TicketInformationType struct {
	// Ticketing type  TL, OK, DO, IN, MA, TR, AT, PT, XL, ST, SS
	Indicator formats.AlphaString_Length2To2 `xml:"indicator"`

	// Ticketing date
	Date formats.Date_DDMMYY `xml:"date,omitempty"`

	// Ticketing time
	Time formats.Time24_HHMM `xml:"time,omitempty"`

	// Office Id
	OfficeId formats.AlphaNumericString_Length1To9 `xml:"officeId,omitempty"`

	// Free flow text
	Freetext formats.AlphaNumericString_Length1To15 `xml:"freetext,omitempty"`

	// Airline code
	AirlineCode formats.AlphaNumericString_Length1To3 `xml:"airlineCode,omitempty"`

	// Queue number
	QueueNumber formats.AlphaNumericString_Length1To3 `xml:"queueNumber,omitempty"`

	// Category number
	QueueCategory formats.AlphaNumericString_Length1To3 `xml:"queueCategory,omitempty"`

	// SITA Addresses
	SitaAddress []formats.AlphaNumericString_Length7To7 `xml:"sitaAddress,omitempty"` // maxOccurs="5"
}

type TicketingCarrierDesignatorType struct {
	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// airline code and printer details
	Carrier *TicketingCarrierType `xml:"carrier,omitempty"`
}

type TicketingCarrierType struct {
	// Airline code of sponsor
	AirlineCode formats.AlphaNumericString_Length1To3 `xml:"airlineCode"`

	// Print itinerary option   IBP, IEP, IBPJ, IEPJ
	OptionInfo formats.AlphaNumericString_Length1To4 `xml:"optionInfo,omitempty"`

	// Printer number
	PrinterNumber formats.AlphaNumericString_Length1To8 `xml:"printerNumber,omitempty"`

	// ISO code 639 - 1988
	Language formats.AlphaNumericString_Length1To3 `xml:"language,omitempty"`
}

type TourCodeType struct {
	// Passenger type  PAX for Passenger  INF for Infant not occupying a seat
	PassengerType formats.AlphaNumericString_Length1To3 `xml:"passengerType,omitempty"`

	// Formatted tour code
	FormatedTour *FormatedTourCodeType `xml:"formatedTour,omitempty"`

	// Net remit
	NetRemit *NetRemitTourCodeType `xml:"netRemit,omitempty"`

	// Freeformat Tour information
	FreeFormatTour *FreeFormatTourCodeType `xml:"freeFormatTour,omitempty"`
}

type TravelProductInformationType struct {
	// date and time details
	Product *ProductDateTimeTypeI `xml:"product,omitempty"`

	// boardpoint details
	BoardpointDetail *LocationTypeI `xml:"boardpointDetail,omitempty"`

	// offpoint details
	OffpointDetail *LocationTypeI `xml:"offpointDetail,omitempty"`

	// airline or system code
	Company *CompanyIdentificationTypeI `xml:"company,omitempty"`

	// product details - number and class
	ProductDetails *ProductIdentificationDetailsTypeI `xml:"productDetails,omitempty"`

	// Product type details. Here: availibility context.
	FlightTypeDetails *ProductTypeDetailsType `xml:"flightTypeDetails,omitempty"`

	// 1. Air segment  To indicate an Informational Air segment:  N for No action required
	ProcessingIndicator formats.AlphaNumericString_Length1To3 `xml:"processingIndicator,omitempty"`
}

type TravellerDetailsTypeI struct {
	// Traveler First Name + titel
	FirstName formats.AlphaNumericString_Length1To56 `xml:"firstName,omitempty"`

	// Traveler Type using  Amadeus codification. In values with 'X00' (X being any letter), 00 should be replaced by any value between 01 and 99.
	Type formats.AlphaNumericString_Length1To3 `xml:"type,omitempty"`

	// 1. Infant (INF) No more info in Edifact. 2. Infant given name only (INF/BILL) Infant given name will be placed in a 2nd occurence of C324 of this (adult) passenger TIF. The 2nd C324/6353 element will contain INF. 3. Infant given and last name (INFGATES/BILL) Infant is treated as a separate TIF following immediately this (adult) passenger TIF. This following TIF C324/6353 element will contain INF.
	InfantIndicator formats.AlphaNumericString_Length1To1 `xml:"infantIndicator,omitempty"`

	// Identification code, 2 cases: ID < 1 to 51 char free text ) or CR < 1 to 40 char free text )
	IdentificationCode formats.AlphaNumericString_Length1To70 `xml:"identificationCode,omitempty"`
}

type TravellerInformationTypeI struct {
	// traveller surname,type and quantity
	Traveller *TravellerSurnameInformationTypeI `xml:"traveller"`

	// Occurrence one relates to the traveler.  Occurrence 2 relates only to an infant accompanying the traveler for whom only the given name is present.
	Passenger []*TravellerDetailsTypeI `xml:"passenger,omitempty"` // maxOccurs="2"
}

type TravellerSurnameInformationTypeI struct {
	// Traveler Last Name or Group name
	Surname formats.AlphaNumericString_Length1To57 `xml:"surname"`

	// G for a group. (The traveler type is in C324/6353)
	Qualifier formats.AlphaNumericString_Length1To3 `xml:"qualifier,omitempty"`

	// - 1 : only one traveler defined by TIFwith exceptions below. - 2 : the traveler is accompanied by  an infant for whom only the given name is present.
	Quantity *formats.NumericInteger_Length1To2 `xml:"quantity,omitempty"`
}
