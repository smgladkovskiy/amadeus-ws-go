package AMA_TicketProcessRefund_v03_0

import "encoding/xml"

/*
<AMA_TicketProcessRefundRQ Version="3.00">
</AMA_TicketProcessRefundRQ>
*/
type Request struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/2010/06/TicketGTP_v3 AMA_TicketProcessRefundRQ"`
	Version string   `xml:"Version,attr"`
	//SchemaLocation string   `xml:"xsi:schemaLocation,attr,omitempty"`
}
