package AMA_TicketInitRefund_v03_0 // amtinrf30

import "encoding/xml"

/*
<AMA_TicketInitRefundRQ xmlns="http://xml.amadeus.com/2010/06/TicketGTP_v3" Version="3.000">
	<Contracts>
		<Contract Number="1257985491256" />  ticket number
	</Contracts>
	<ActionDetails>
		<ActionDetail Indicator="ATC" />  “ATC” action detail
	</ActionDetails>
</AMA_TicketInitRefundRQ>


<AMA_TicketInitRefundRQ Version="3.000" xsi:schemaLocation="http://xml.amadeus.com/2010/06/TicketGTP_v3 AMA_TicketInitRefundRQ.xsd">
	<Contracts>
		<Contract Number="78546321"></Contract>
		<Contract Number="65231492"></Contract>
	</Contracts>
	<StockProvider StockProviderCode="FRR" StockTypeCode="R"></StockProvider>
	<AdditionalDataList>
		<AdditionalData>
			<att:Data Key="RLOC" Value="6UMR2N"></att:Data>
		</AdditionalData>
	</AdditionalDataList>
</AMA_TicketInitRefundRQ>
*/
type Request struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/2010/06/TicketGTP_v3 AMA_TicketInitRefundRQ"`
	Version string   `xml:"Version,attr"`
	//SchemaLocation     string              `xml:"xsi:schemaLocation,attr,omitempty"`

	Contracts          *Contracts          `xml:"Contracts,omitempty"`
	StockProvider      *StockProvider      `xml:"StockProvider,omitempty"`
	ActionDetails      *ActionDetails      `xml:"ActionDetails,omitempty"`
	AdditionalDataList *AdditionalDataList `xml:"AdditionalDataList,omitempty"`
}

type Contracts struct {
	Contract []*Contract `xml:"Contract"`
}

type Contract struct {
	Number string `xml:"Number,attr"`
}

type StockProvider struct {
	StockProviderCode string `xml:"StockProviderCode,attr"`
	StockTypeCode     string `xml:"StockTypeCode,attr"`
}

type ActionDetails struct {
	ActionDetail *ActionDetail `xml:"ActionDetail"`
}

type ActionDetail struct {
	Indicator string `xml:"Indicator,attr"`
}

type AdditionalDataList struct {
	AdditionalData *AdditionalData `xml:"AdditionalData,omitempty"`
}

type AdditionalData struct {
	Data *Data `xml:"att:Data,omitempty"`
}

type Data struct {
	Key   string `xml:"Key,attr"`
	Value string `xml:"Value,attr"`
}
