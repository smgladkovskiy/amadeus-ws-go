package client

import (
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayQueryReport/v10.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayQueryReport/v10.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/cancelDocument/v11.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/deleteTST/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/displayTST/v07.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/displayTST/v07.1/response"
)

func (client *AmadeusClient) SalesReportsDisplayQueryReportV101(query *SalesReports_QueryReportRequest_v10_1.Request) (*SalesReports_QueryReportReply_v10_1.Response, *ResponseSOAPHeader, error) {
	var reply SalesReports_QueryReportReply_v10_1.Response
	header, err := client.service.Call(soapUrl, "TSRQRQ_10_1_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) TicketCancelDocumentV111(query *Ticket_CancelDocument_v11_1.Request) (*Ticket_CancelDocument_v11_1.Response, *ResponseSOAPHeader, error) {
	var reply Ticket_CancelDocument_v11_1.Response
	header, err := client.service.Call(soapUrl, "TRCANQ_11_1_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) TicketDeleteTSTV041(query *Ticket_DeleteTST_v04_1.Request) (*Ticket_DeleteTST_v04_1.Response, *ResponseSOAPHeader, error) {
	var reply Ticket_DeleteTST_v04_1.Response
	header, err := client.service.Call(soapUrl, "TTSTDQ_04_1_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) TicketDisplayTSTV071(query *Ticket_DisplayTSTRequest_v07_1.Request) (*Ticket_DisplayTSTResponse_v07_1.Response, *ResponseSOAPHeader, error) {
	var reply Ticket_DisplayTSTResponse_v07_1.Response
	header, err := client.service.Call(soapUrl, "TTSTRQ_07_1_1A", query, &reply, client)
	return &reply, header, err
}

//func (client *AmadeusClient) TicketDisplayTSTV1511A(query *Ticket_DisplayTST_v15_1_1A.Request) (*Ticket_DisplayTSTReply_v15_1_1A.Response, *ResponseSOAPHeader, error) {
//	var reply Ticket_DisplayTSTReply_v15_1_1A.Response
//	header, err := client.service.Call(soapUrl, "TTSTRQ_15_1_1A", query, &reply, client)
//	return &reply, header, err
//}
