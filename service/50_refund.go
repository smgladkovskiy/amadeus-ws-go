package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/client"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketIgnoreRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketInitRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketProcessRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/ignore/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayTransactionReport/v13.2"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/processEDoc/v15.2/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/processEDoc/v15.2/response"
)

func (s *service) RefundInit(query *AMA_TicketInitRefund_v03_0.Request) (*AMA_TicketInitRefund_v03_0.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.AMATicketInitRefundV030(query)
}

func (s *service) RefundIgnore(query *AMA_TicketIgnoreRefund_v03_0.Request) (*AMA_TicketIgnoreRefund_v03_0.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.AMATicketIgnoreRefundV030(query)
}

func (s *service) RefundProcess(query *AMA_TicketProcessRefund_v03_0.Request) (*AMA_TicketProcessRefund_v03_0.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.AMATicketProcessRefundV030(query)
}

func (s *service) TicketProcessEDoc(query *Ticket_ProcessEDocRequest_v15_2.Request) (*Ticket_ProcessEDocResponse_v15_2.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.TicketProcessEDocV152(query)
}

func (s *service) SalesReportsDisplayTransactionReport(query *SalesReports_DisplayTransactionReport_v13_2.Request) (*SalesReports_DisplayTransactionReport_v13_2.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.SalesReportsDisplayTransactionReportV132(query)
}

func (s *service) PNRIgnore(query *PNR_Ignore_v04_1.Request) (*PNR_Ignore_v04_1.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.PNRIgnoreV041(query)
}
