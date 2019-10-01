package client

import (
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketIgnoreRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketInitRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketProcessRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/ignore/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayTransactionReport/v13.2"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/processEDoc/v15.2/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/processEDoc/v15.2/response"
)

func (client *AmadeusClient) AMATicketIgnoreRefundV030(query *AMA_TicketIgnoreRefund_v03_0.Request) (*AMA_TicketIgnoreRefund_v03_0.Response, *ResponseSOAPHeader, error) {
	var reply AMA_TicketIgnoreRefund_v03_0.Response
	header, err := client.service.Call(amaUrl, "Ticket_IgnoreRefund_3.0", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) AMATicketInitRefundV030(query *AMA_TicketInitRefund_v03_0.Request) (*AMA_TicketInitRefund_v03_0.Response, *ResponseSOAPHeader, error) {
	var reply AMA_TicketInitRefund_v03_0.Response
	header, err := client.service.Call(amaUrl, "Ticket_InitRefund_3.0", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) AMATicketProcessRefundV030(query *AMA_TicketProcessRefund_v03_0.Request) (*AMA_TicketProcessRefund_v03_0.Response, *ResponseSOAPHeader, error) {
	var reply AMA_TicketProcessRefund_v03_0.Response
	header, err := client.service.Call(amaUrl, "Ticket_ProcessRefund_3.0", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) TicketProcessEDocV152(query *Ticket_ProcessEDocRequest_v15_2.Request) (*Ticket_ProcessEDocResponse_v15_2.Response, *ResponseSOAPHeader, error) {
	var reply Ticket_ProcessEDocResponse_v15_2.Response
	header, err := client.service.Call(soapUrl, "TATREQ_15_2_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) SalesReportsDisplayTransactionReportV132(query *SalesReports_DisplayTransactionReport_v13_2.Request) (*SalesReports_DisplayTransactionReport_v13_2.Response, *ResponseSOAPHeader, error) {
	var reply SalesReports_DisplayTransactionReport_v13_2.Response
	header, err := client.service.Call(soapUrl, "TSRTRQ_13_2_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) PNRIgnoreV041(query *PNR_Ignore_v04_1.Request) (*PNR_Ignore_v04_1.Response, *ResponseSOAPHeader, error) {
	var reply PNR_Ignore_v04_1.Response
	header, err := client.service.Call(soapUrl, "CLTREQ_04_1_IA", query, &reply, client)
	return &reply, header, err
}
