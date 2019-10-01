package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/client"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayQueryReport/v10.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayQueryReport/v10.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/cancelDocument/v11.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/deleteTST/v04.1"
)

func (s *service) SalesReportsDisplayQueryReport(query *SalesReports_QueryReportRequest_v10_1.Request) (*SalesReports_QueryReportReply_v10_1.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.SalesReportsDisplayQueryReportV101(query)
}

func (s *service) TicketCancelDocument(query *Ticket_CancelDocument_v11_1.Request) (*Ticket_CancelDocument_v11_1.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.TicketCancelDocumentV111(query)
}

func (s *service) TicketDeleteTST(query *Ticket_DeleteTST_v04_1.Request) (*Ticket_DeleteTST_v04_1.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.TicketDeleteTSTV041(query)
}
