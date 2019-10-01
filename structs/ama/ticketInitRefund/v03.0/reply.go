package AMA_TicketInitRefund_v03_0 // amtinrf30

//import "encoding/xml"

/*
<AMA_TicketInitRefundRS xmlns="http://xml.amadeus.com/2010/06/TicketGTP_v3"
                        xmlns:ama="http://xml.amadeus.com/2010/06/Types_v2"
                        xmlns:att="http://xml.amadeus.com/2010/06/TicketTypes_v2"
                        xmlns:iata="http://www.iata.org/IATA/2007/00/IATA2010.1"
                        xmlns:ota="http://www.opentravel.org/OTA/2003/05/OTA2010B"
                        xmlns:xsd="http://www.w3.org/2001/XMLSchema" Version="3.000">
	<GeneralReply>
		<Success />
	</GeneralReply>
	<FunctionalData>
		<ContractBundle ID="1">
			<Success />
			<RefundDetails>
				<att:RuleID>
					<att:ReferenceDetails>
						<att:Type>RID</att:Type>
						<att:Value>36</att:Value>
					</att:ReferenceDetails>
				</att:RuleID>
				<att:Contracts>
					<att:Contract ID="1" IssueDate="2017-06-02">
						<att:Segments>
							<att:Segment Tattoo="1" />
						</att:Segments>
						<att:Passengers>
							<att:Passenger Tattoo="1">
								<att:FirstName>ONE PAX</att:FirstName>
								<att:LastName>ATC REFUND</att:LastName>
								<att:FullName>ATC REFUND/ONE PAX</att:FullName>
							</att:Passenger>
						</att:Passengers>
						<att:Taxes>
							<att:Tax Amount="560" CurrencyCode="USD" DecimalPlaces="2" Category="R" ISO_Code="AY" />
							<att:Tax Amount="1800" CurrencyCode="USD" DecimalPlaces="2" Category="R" ISO_Code="US" />
							<att:Tax Amount="450" CurrencyCode="USD" DecimalPlaces="2" Category="R" ISO_Code="XF">
								<att:AssociatedPFC_Tax Amount="450" CurrencyCode="USD" DecimalPlaces="2" CountryCode="MIA" />
							</att:Tax>
							<att:Tax Amount="12500" CurrencyCode="USD" DecimalPlaces="2" Category="R" ISO_Code="YQ" />
							<att:Tax Amount="560" CurrencyCode="USD" DecimalPlaces="2" Category="PT" ISO_Code="AY" />
							<att:Tax Amount="1800" CurrencyCode="USD" DecimalPlaces="2" Category="PT" ISO_Code="US" />
							<att:Tax Amount="450" CurrencyCode="USD" DecimalPlaces="2" Category="PT" ISO_Code="XF">
								<att:AssociatedPFC_Tax Amount="450" CurrencyCode="USD" DecimalPlaces="2" CountryCode="MIA" />
							</att:Tax>
							<att:Tax Amount="12500" CurrencyCode="USD" DecimalPlaces="2" Category="PT" ISO_Code="YQ" />
							<att:Tax Amount="15310" CurrencyCode="USD" DecimalPlaces="2" Category="701">
								<ota:TaxDescription Name="XT" />
							</att:Tax>
						</att:Taxes>
						<att:Commissions>
							<att:Commission Percent="0.00">
								<ota:UniqueID Type="NEW" ID="1" />
								<ota:CommissionPayableAmount Amount="0" CurrencyCode="USD" DecimalPlaces="2" />
								<ota:Comment Name="P" />
							</att:Commission>
						</att:Commissions>
						<att:MonetaryInformations>
							<att:MonetaryInformation Amount="270000" CurrencyCode="USD" DecimalPlaces="2" Qualifier="B" />
							<att:MonetaryInformation Amount="0" CurrencyCode="USD" DecimalPlaces="2" Qualifier="RFU" />
							<att:MonetaryInformation Amount="270000" CurrencyCode="USD" DecimalPlaces="2" Qualifier="FRF" />
							<att:MonetaryInformation Amount="285310" CurrencyCode="USD" DecimalPlaces="2" Qualifier="RFT" />
							<att:MonetaryInformation Amount="15310" CurrencyCode="USD" DecimalPlaces="2" Qualifier="TXT" />
							<att:MonetaryInformation Amount="15310" CurrencyCode="USD" DecimalPlaces="2" Qualifier="TP" />
							<att:MonetaryInformation Amount="0" CurrencyCode="USD" DecimalPlaces="2" Qualifier="NRC" />
							<att:MonetaryInformation Amount="285310" CurrencyCode="USD" DecimalPlaces="2" Qualifier="RFA" />
							<att:MonetaryInformation Amount="0" CurrencyCode="USD" DecimalPlaces="2" Qualifier="RUA" />
						</att:MonetaryInformations>
						<att:PricingDetails>
							<att:PriceTicketDetails>
								<att:Indicator>I</att:Indicator>
							</att:PriceTicketDetails>
						</att:PricingDetails>
						<att:DocumentAndCouponInformation>
							<att:DocumentNumber Number="12579854912560" />  ticket number to include the check digit
							<att:CouponGroup>
								<att:CouponInformationDetails>
									<att:CouponNumber>1</att:CouponNumber>
									<att:CouponStatus>RF</att:CouponStatus>
								</att:CouponInformationDetails>
							</att:CouponGroup>
						</att:DocumentAndCouponInformation>
						<att:RefundedRoute>
							<att:Station>MIA</att:Station>
							<att:Station>LHR</att:Station>
						</att:RefundedRoute>
						<att:Refundable Amount="285310" CurrencyCode="USD" DecimalPlaces="2" />
					</att:Contract>
				</att:Contracts>
				<att:DateTime>
					<att:BusinessSemantic Code="DR" />
					<att:StructuredDateTime>
						<att:Year>2017</att:Year>
						<att:Month>06</att:Month>
						<att:Day>02</att:Day>
					</att:StructuredDateTime>
				</att:DateTime>
				<att:DateTime>
					<att:BusinessSemantic Code="710" />
					<att:StructuredDateTime>
						<att:Year>2017</att:Year>
						<att:Month>06</att:Month>
						<att:Day>02</att:Day>
					</att:StructuredDateTime>
				</att:DateTime>
				<att:FormOfPayments>
					<att:FormOfPayment ID="1" Type="CA" Amount="285310" CurrencyCode="USD" DecimalPlaces="2">
						<att:FreeFlow>CHECK</att:FreeFlow>
					</att:FormOfPayment>
				</att:FormOfPayments>
				<att:ReportingOffice>
					<att:AgentCode>33895934</att:AgentCode>
					<att:OfficeID>MIA1AXXXX</att:OfficeID>
					<att:Originator>9998WSSU</att:Originator>
				</att:ReportingOffice>
				<att:TransactionCode>TKTT</att:TransactionCode>
				<att:ReferenceDetails>
					<att:ReferenceDetail>
						<att:Type>TKT</att:Type>
						<att:Value>Y</att:Value>
					</att:ReferenceDetail>
					<att:ReferenceDetail>
						<att:Type>DIS</att:Type>
						<att:Value>C</att:Value>
					</att:ReferenceDetail>
				</att:ReferenceDetails>
			</RefundDetails>
		</ContractBundle>
	</FunctionalData>
</AMA_TicketInitRefundRS>

<AMA_TicketInitRefundRS xmlns="http://xml.amadeus.com/2010/06/TicketGTP_v3"
                            xmlns:ama="http://xml.amadeus.com/2010/06/Types_v2"
                            xmlns:att="http://xml.amadeus.com/2010/06/TicketTypes_v2"
                            xmlns:iata="http://www.iata.org/IATA/2007/00/IATA2010.1"
                            xmlns:ota="http://www.opentravel.org/OTA/2003/05/OTA2010B"
                            xmlns:xsd="http://www.w3.org/2001/XMLSchema" Version="3.000">
	<GeneralReply>
		<Success/>
	</GeneralReply>
	<FunctionalData>
		<ContractBundle ID="1">
			<Errors>
				<ama:Errors>
					<ama:Error Type="000"
                               Code="25677"
                               RecordID="6281085809716">ATC REFUND NOT AUTHORIZED</ama:Error>
				</ama:Errors>
			</Errors>
		</ContractBundle>
	</FunctionalData>
</AMA_TicketInitRefundRS>
*/

/*
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:awsse="http://xml.amadeus.com/2010/06/Session_v3" xmlns:wsa="http://www.w3.org/2005/08/addressing">
	<soap:Header>
		<wsa:To>http://www.w3.org/2005/08/addressing/anonymous</wsa:To>
		<wsa:From>
			<wsa:Address>https://nodeD1.test.webservices.amadeus.com/1ASIWIBETMN</wsa:Address>
		</wsa:From>
		<wsa:Action>http://webservices.amadeus.com/Ticket_InitRefund_3.0</wsa:Action>
		<wsa:MessageID>urn:uuid:14692eda-6a3e-0804-5965-3ff835fc6352</wsa:MessageID>
		<wsa:RelatesTo RelationshipType="http://www.w3.org/2005/08/addressing/reply">RVTWUVQCLIU8SBCJB2E43I</wsa:RelatesTo>
		<awsse:Session TransactionStatusCode="InSeries">
			<awsse:SessionId>01Q8WYACMG</awsse:SessionId>
			<awsse:SequenceNumber>4</awsse:SequenceNumber>
			<awsse:SecurityToken>3EKO6RQZFLD3B22U097H15OIJI</awsse:SecurityToken>
		</awsse:Session>
	</soap:Header>
	<soap:Body>
		<AMA_TicketInitRefundRS xmlns="http://xml.amadeus.com/2010/06/TicketGTP_v3" xmlns:ama="http://xml.amadeus.com/2010/06/Types_v2" xmlns:att="http://xml.amadeus.com/2010/06/TicketTypes_v2" xmlns:iata="http://www.iata.org/IATA/2007/00/IATA2010.1" xmlns:ota="http://www.opentravel.org/OTA/2003/05/OTA2010B" xmlns:xsd="http://www.w3.org/2001/XMLSchema" Version="3.000">
			<GeneralReply>
				<Success/>
			</GeneralReply>
			<FunctionalData>
				<ContractBundle ID="1">
					<Success/>
					<ListDisplay>
						<Ticket Number="08210858097326" Type="T"/>
						<Documents>
							<Document>
								<SequenceNumber>1</SequenceNumber>
								<ReportingOffice>
									<att:AgentCode>92227166</att:AgentCode>
									<att:OfficeID>MOWR224PW</att:OfficeID>
									<att:Originator>9999WSSU</att:Originator>
								</ReportingOffice>
								<TransactionCode>RFND</TransactionCode>
								<DocumentStatus>UNK</DocumentStatus>
								<Passenger Tattoo="3">
									<att:FirstName>SOFIA MRS</att:FirstName>
									<att:LastName>IVANOVA</att:LastName>
									<att:FullName>IVANOVA/SOFIA MRS</att:FullName>
								</Passenger>
							</Document>
							<Document>
								<SequenceNumber>2</SequenceNumber>
								<ReportingOffice>
									<att:AgentCode>92227166</att:AgentCode>
									<att:OfficeID>MOWR224PW</att:OfficeID>
									<att:Originator>9999WSSU</att:Originator>
								</ReportingOffice>
								<TransactionCode>TKTT</TransactionCode>
								<DocumentStatus>PND</DocumentStatus>
								<Passenger Tattoo="3">
									<att:FirstName>SOFIA MRS</att:FirstName>
									<att:LastName>IVANOVA</att:LastName>
									<att:FullName>IVANOVA/SOFIA MRS</att:FullName>
								</Passenger>
							</Document>
						</Documents>
					</ListDisplay>
				</ContractBundle>
			</FunctionalData>
		</AMA_TicketInitRefundRS>
	</soap:Body>
</soap:Envelope>
*/
type Response struct {
	//XMLName        xml.Name        `xml:"http://xml.amadeus.com/2010/06/TicketGTP_v3 AMA_TicketInitRefundRS"`
	//AMAPAttr       string          `xml:"xmlns:ama,attr"`
	//ATTPAttr       string          `xml:"xmlns:att,attr"`
	//IATAAttr       string          `xml:"xmlns:iata,attr"`
	//OTAAttr        string          `xml:"xmlns:ota,attr"`
	//XSDAttr        string          `xml:"xmlns:xsd,attr"`
	//Version        string          `xml:"Version,attr"`

	GeneralReply   *GeneralReply   `xml:"GeneralReply"`
	FunctionalData *FunctionalData `xml:"FunctionalData"`
}

type GeneralReply struct {
	Success *Success `xml:"Success,omitempty"`
}

type Success struct {
}

type FunctionalData struct {
	ContractBundle []*ContractBundle `xml:"ContractBundle"`
}

type ContractBundle struct {
	ID string `xml:"ID,attr"`

	Success       *Success       `xml:"Success,omitempty"`
	Errors        *Errors        `xml:"Errors,omitempty"`
	RefundDetails *RefundDetails `xml:"RefundDetails,omitempty"`
	ListDisplay   []*ListDisplay `xml:"ListDisplay,omitempty"`
}

type Errors struct {
	AMAErrors *AMAErrors `xml:"Errors,omitempty"`
}

type AMAErrors struct {
	AMAError []*AMAError `xml:"Error"`
}

type AMAError struct {
	Type     string `xml:"Type,attr"`
	Code     string `xml:"Code,attr"`
	RecordID string `xml:"RecordID,attr"`
	Message  string `xml:",chardata"`
}

type RefundDetails struct {
	RuleID          *RuleID          `xml:"RuleID,omitempty"`
	Contracts       *ContractsReply  `xml:"Contracts,omitempty"`
	DateTime        []*DateTime      `xml:"DateTime,omitempty"`
	FormOfPayments  *FormOfPayments  `xml:"FormOfPayments,omitempty"`
	ReportingOffice *ReportingOffice `xml:"ReportingOffice,omitempty"`
	//TransactionCode *TransactionCode   `xml:"TransactionCode,omitempty"`
	TransactionCode  string            `xml:"TransactionCode,omitempty"`
	ReferenceDetails *ReferenceDetails `xml:"ReferenceDetails,omitempty"`
}

type RuleID struct {
	ReferenceDetails *ReferenceDetail `xml:"ReferenceDetails,omitempty"`
}

type ReferenceDetail struct {
	Type  string `xml:"Type,omitempty"`
	Value string `xml:"Value,omitempty"`
}

type ContractsReply struct {
	Contract []*ContractReply `xml:"Contract"`
}

type ContractReply struct {
	ID        string `xml:"ID,attr"`
	IssueDate string `xml:"IssueDate,attr"` // example: "2017-06-02"

	Segments                     *Segments                     `xml:"Segments,omitempty"`
	Passengers                   *Passengers                   `xml:"Passengers,omitempty"`
	Taxes                        *Taxes                        `xml:"Taxes,omitempty"`
	Commissions                  *Commissions                  `xml:"Commissions,omitempty"`
	MonetaryInformations         *MonetaryInformations         `xml:"MonetaryInformations,omitempty"`
	PricingDetails               *PricingDetails               `xml:"PricingDetails,omitempty"`
	DocumentAndCouponInformation *DocumentAndCouponInformation `xml:"DocumentAndCouponInformation,omitempty"`
	RefundedRoute                *RefundedRoute                `xml:"RefundedRoute,omitempty"`
	Refundable                   *Refundable                   `xml:"Refundable,omitempty"`
}

type Segments struct {
	Segment []*Segment `xml:"Segments"`
}

type Segment struct {
	Tattoo string `xml:"Tattoo,attr"`
}

type Passengers struct {
	Passenger []*Passenger `xml:"Passenger"`
}

type Passenger struct {
	Tattoo string `xml:"Tattoo,attr"`

	FirstName string `xml:"FirstName,omitempty"`
	LastName  string `xml:"LastName,omitempty"`
	FullName  string `xml:"FullName,omitempty"`
}

type Taxes struct {
	Tax []*Tax `xml:"Tax"`
}

type Tax struct {
	Amount        string `xml:"Amount,attr"`
	CurrencyCode  string `xml:"CurrencyCode,attr"`
	DecimalPlaces string `xml:"DecimalPlaces,attr"`
	Category      string `xml:"Category,attr"`
	ISOCode       string `xml:"ISO_Code,attr,omitempty"`

	AssociatedPFCTax *AssociatedPFCTax `xml:"AssociatedPFC_Tax,omitempty"`
	TaxDescription   *TaxDescription   `xml:"TaxDescription,omitempty"`
}

type AssociatedPFCTax struct {
	Amount        string `xml:"Amount,attr"`
	CurrencyCode  string `xml:"CurrencyCode,attr"`
	DecimalPlaces string `xml:"DecimalPlaces,attr"`
	CountryCode   string `xml:"CountryCode,attr"`
}

type TaxDescription struct {
	Name string `xml:"Name,attr"`
}

type Commissions struct {
	Commission []*Commission `xml:"Commission"`
}

type Commission struct {
	Percent float64 `xml:"Percent,attr"`

	UniqueID                *UniqueID                `xml:"UniqueID,omitempty"`
	CommissionPayableAmount *CommissionPayableAmount `xml:"CommissionPayableAmount,omitempty"`
	Comment                 *Comment                 `xml:"Comment,omitempty"`
}

type UniqueID struct {
	Type string `xml:"Type,attr"`
	ID   string `xml:"ID,attr"`
}

type CommissionPayableAmount struct {
	Amount        string `xml:"Amount,attr"`
	CurrencyCode  string `xml:"CurrencyCode,attr"`
	DecimalPlaces string `xml:"DecimalPlaces,attr"`
}

type Comment struct {
	Name string `xml:"Name,attr"`
}

type MonetaryInformations struct {
	MonetaryInformation []*MonetaryInformation `xml:"MonetaryInformation"`
}

type MonetaryInformation struct {
	Amount        string `xml:"Amount,attr"`
	CurrencyCode  string `xml:"CurrencyCode,attr"`
	DecimalPlaces string `xml:"DecimalPlaces,attr"`
	Qualifier     string `xml:"Qualifier,attr"`
}

type PricingDetails struct {
	PriceTicketDetails *PriceTicketDetails `xml:"PriceTicketDetails,omitempty"`
}

type PriceTicketDetails struct {
	Indicator string `xml:"Indicator,omitempty"`
}

type DocumentAndCouponInformation struct {
	DocumentNumber *DocumentNumber `xml:"DocumentNumber,omitempty"`
	CouponGroup    *CouponGroup    `xml:"CouponGroup,omitempty"`
}

type DocumentNumber struct {
	Number string `xml:"Number,attr"`
}

type CouponGroup struct {
	CouponInformationDetails *CouponInformationDetails `xml:"CouponInformationDetails,omitempty"`
}

type CouponInformationDetails struct {
	CouponNumber string `xml:"CouponNumber,omitempty"`
	CouponStatus string `xml:"CouponStatus,omitempty"`
}

type RefundedRoute struct {
	//Station []*Station `xml:"Station,omitempty"`
	Station []string `xml:"Station,omitempty"`
}

//type Station struct {
//	Value string `xml:",chardata"`
//}

type Refundable struct {
	Amount        string `xml:"Amount,attr"`
	CurrencyCode  string `xml:"CurrencyCode,attr"`
	DecimalPlaces string `xml:"DecimalPlaces,attr"`
}

type DateTime struct {
	BusinessSemantic   *BusinessSemantic   `xml:"BusinessSemantic,omitempty"`
	StructuredDateTime *StructuredDateTime `xml:"StructuredDateTime,omitempty"`
}

type BusinessSemantic struct {
	Code string `xml:"Code,attr"`
}

type StructuredDateTime struct {
	Year  string `xml:"Year"`
	Month string `xml:"Month"`
	Day   string `xml:"Day"`
}

type FormOfPayments struct {
	FormOfPayment []*FormOfPayment `xml:"FormOfPayment"`
}

type FormOfPayment struct {
	ID            string `xml:"ID,attr"`
	Type          string `xml:"Type,attr"`
	Amount        string `xml:"Amount,attr"`
	CurrencyCode  string `xml:"CurrencyCode,attr"`
	DecimalPlaces string `xml:"DecimalPlaces,attr"`

	//FreeFlow      *FreeFlow `xml:"FreeFlow,omitempty"`
	FreeFlow string `xml:",FreeFlow,omitempty"`
}

//type FreeFlow struct {
//	Value string `xml:",chardata"`
//}

type ReportingOffice struct {
	AgentCode  string `xml:"AgentCode,omitempty"`
	OfficeID   string `xml:"OfficeID,omitempty"`
	Originator string `xml:"Originator,omitempty"`
}

//type TransactionCode struct {
//	Value string `xml:",chardata"`
//}

type ReferenceDetails struct {
	ReferenceDetail []*ReferenceDetail `xml:"ReferenceDetail,omitempty"`
}

type ListDisplay struct {
	Ticket    *Ticket    `xml:"Ticket"`
	Documents *Documents `xml:"Documents"`
}

type Ticket struct {
	Number string `xml:"Number,attr"`
	Type   string `xml:"Type,attr"`
}

type Documents struct {
	Document []*Document `xml:"Document"`
}

type Document struct {
	SequenceNumber  string           `xml:"SequenceNumber"`
	ReportingOffice *ReportingOffice `xml:"ReportingOffice,omitempty"`
	//TransactionCode *TransactionCode `xml:"TransactionCode,omitempty"`
	TransactionCode string       `xml:"TransactionCode,omitempty"`
	Passenger       []*Passenger `xml:"Passenger"`
}
