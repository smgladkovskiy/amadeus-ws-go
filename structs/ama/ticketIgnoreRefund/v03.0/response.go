package AMA_TicketIgnoreRefund_v03_0

import "encoding/xml"

/*
<AMA_TicketIgnoreRefundRS Version="3.000"
                          xsi:schemaLocation="http://xml.amadeus.com/2010/06/TicketGTP_v3"
                          xmlns="http://xml.amadeus.com/2010/06/TicketGTP_v3">
	<Success/>
</AMA_TicketIgnoreRefundRS>
*/
type Response struct {
	XMLName        xml.Name `xml:"http://xml.amadeus.com/2010/06/TicketGTP_v3 AMA_TicketIgnoreRefundRS"`
	Version        string   `xml:"Version,attr"`
	SchemaLocation string   `xml:"xsi:schemaLocation,attr,omitempty"`

	Success *Success `xml:"Success,omitempty"`
}

type Success struct {
}
