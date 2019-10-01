package AMA_TicketProcessRefund_v03_0 // amtpcrf30

//import "encoding/xml"

/*
<AMA_TicketProcessRefundRS xmlns="http://xml.amadeus.com/2010/06/TicketGTP_v3"
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
				<att:Contracts>
					<att:Contract ID="1" Status="Refunded" IssueDate="2017-06-02">
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
						<att:DocumentAndCouponInformation>
							<att:DocumentNumber Number="12579854912560" />  ticket number to include the check digit
							<att:CouponGroup>
								<att:CouponInformationDetails>
									<att:CouponNumber>1</att:CouponNumber>
									<att:CouponStatus>RF</att:CouponStatus>
									<att:SettlementAuthorization> 1250OTD830BI0</att:SettlementAuthorization>
								</att:CouponInformationDetails>
							</att:CouponGroup>
						</att:DocumentAndCouponInformation>
						<att:Refundable Amount="285310" CurrencyCode="USD" DecimalPlaces="2" />
					</att:Contract>
				</att:Contracts>
				<att:DateTime>
					<att:BusinessSemantic Code="TID" />
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
				<att:GlobalRefundReceipt> 1250OTD830BI0</att:GlobalRefundReceipt>
			</RefundDetails>
		</ContractBundle>
	</FunctionalData>
</AMA_TicketProcessRefundRS>
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
	Contracts      *Contracts      `xml:"Contracts,omitempty"`
	DateTime       []*DateTime     `xml:"DateTime,omitempty"`
	FormOfPayments *FormOfPayments `xml:"FormOfPayments,omitempty"`
	//GlobalRefundReceipt *GlobalRefundReceipt `xml:"GlobalRefundReceipt,omitempty"`
	GlobalRefundReceipt string `xml:"GlobalRefundReceipt,omitempty"`
}

type Contracts struct {
	Contract []*Contract `xml:"Contract"`
}

type Contract struct {
	ID        string `xml:"ID,attr"`
	Status    string `xml:"Status,attr"`
	IssueDate string `xml:"IssueDate,attr"` // example: "2017-06-02"

	Segments                     *Segments                     `xml:"Segments,omitempty"`
	Passengers                   *Passengers                   `xml:"Passengers,omitempty"`
	Taxes                        *Taxes                        `xml:"Taxes,omitempty"`
	Commissions                  *Commissions                  `xml:"Commissions,omitempty"`
	MonetaryInformations         *MonetaryInformations         `xml:"MonetaryInformations,omitempty"`
	DocumentAndCouponInformation *DocumentAndCouponInformation `xml:"DocumentAndCouponInformation,omitempty"`
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
	CouponNumber            string `xml:"CouponNumber,omitempty"`
	CouponStatus            string `xml:"CouponStatus,omitempty"`
	SettlementAuthorization string `xml:"SettlementAuthorization,omitempty"`
}

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
	FreeFlow string `xml:"FreeFlow,omitempty"`
}

//type GlobalRefundReceipt struct {
//	Value string `xml:",chardata"`
//}
