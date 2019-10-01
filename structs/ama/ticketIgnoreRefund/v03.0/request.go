package AMA_TicketIgnoreRefund_v03_0 // amtigrf30

import "encoding/xml"

/*
<AMA_TicketIgnoreRefundRQ Version="3.000"
                          xsi:schemaLocation="http://xml.amadeus.com/2010/06/TicketGTP_v3"
                          xmlns="http://xml.amadeus.com/2010/06/TicketGTP_v3"/>
*/
type Request struct {
	XMLName xml.Name `xml:"http://xml.amadeus.com/2010/06/TicketGTP_v3 AMA_TicketIgnoreRefundRQ"`
	Version string   `xml:"Version,attr"`
	//SchemaLocation string   `xml:"xsi:schemaLocation,attr,omitempty"`
}
