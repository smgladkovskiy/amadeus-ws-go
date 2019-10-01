package client

import (
	"github.com/tmconsulting/amadeus-golang-sdk/structs/air/sellFromRecommendation/v05.2/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/air/sellFromRecommendation/v05.2/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/pricePNRWithBookingClass/v14.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/pricePNRWithBookingClass/v14.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/addMultiElements/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/cancel/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/reply/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/createTSTFromPricing/v04.1"
)

func (client *AmadeusClient) AirSellFromRecommendationV052(query *Air_SellFromRecommendationRequest_v05_2.Request) (*Air_SellFromRecommendationResponse_v05_2.Response, *ResponseSOAPHeader, error) {
	var reply Air_SellFromRecommendationResponse_v05_2.Response
	header, err := client.service.Call(soapUrl, "ITAREQ_05_2_IA", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) PNRAddMultiElementsV113(query *PNR_AddMultiElementsRequest_v11_3.Request) (*PNR_Reply_v11_3.Response, *ResponseSOAPHeader, error) {
	var reply PNR_Reply_v11_3.Response
	header, err := client.service.Call(soapUrl, "PNRADD_11_3_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) FarePricePNRWithBookingClassV141(query *Fare_PricePNRWithBookingClassRequest_v14_1.Request) (*Fare_PricePNRWithBookingClassResponse_v14_1.Response, *ResponseSOAPHeader, error) {
	var reply Fare_PricePNRWithBookingClassResponse_v14_1.Response
	header, err := client.service.Call(soapUrl, "TPCBRQ_14_1_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) TicketCreateTSTFromPricingV041(query *Ticket_CreateTSTFromPricing_v04_1.Request) (*Ticket_CreateTSTFromPricing_v04_1.Response, *ResponseSOAPHeader, error) {
	var reply Ticket_CreateTSTFromPricing_v04_1.Response
	header, err := client.service.Call(soapUrl, "TAUTCQ_04_1_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) PNRCancelV113(query *PNR_Cancel_v11_3.Request) (*PNR_Reply_v11_3.Response, *ResponseSOAPHeader, error) {
	var reply PNR_Reply_v11_3.Response
	header, err := client.service.Call(soapUrl, "PNRXCL_11_3_1A", query, &reply, client)
	return &reply, header, err
}
