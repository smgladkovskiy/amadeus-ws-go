package client

import (
	"github.com/tmconsulting/amadeus-golang-sdk/structs/docIssuance/issueTicket/v09.1"
)

func (client *AmadeusClient) DocIssuanceIssueTicketV091(query *DocIssuance_IssueTicket_v09_1.Request) (*DocIssuance_IssueTicket_v09_1.Response, *ResponseSOAPHeader, error) {
	var reply DocIssuance_IssueTicket_v09_1.Response
	header, err := client.service.Call(soapUrl, "TTKTIQ_09_1_1A", query, &reply, client)
	return &reply, header, err
}
