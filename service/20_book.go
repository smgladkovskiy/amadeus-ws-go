package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/client"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/air/sellFromRecommendation/v05.2/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/air/sellFromRecommendation/v05.2/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/pricePNRWithBookingClass/v14.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/pricePNRWithBookingClass/v14.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/addMultiElements/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/reply/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/createTSTFromPricing/v04.1"
)

func (s *service) AirSellFromRecommendation(query *Air_SellFromRecommendationRequest_v05_2.Request) (*Air_SellFromRecommendationResponse_v05_2.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.AirSellFromRecommendationV052(query)
}

func (s *service) PNRAddMultiElements(query *PNR_AddMultiElementsRequest_v11_3.Request) (*PNR_Reply_v11_3.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.PNRAddMultiElementsV113(query)
}

func (s *service) FarePricePNRWithBookingClass(query *Fare_PricePNRWithBookingClassRequest_v14_1.Request) (*Fare_PricePNRWithBookingClassResponse_v14_1.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.FarePricePNRWithBookingClassV141(query)
}

func (s *service) TicketCreateTSTFromPricing(query *Ticket_CreateTSTFromPricing_v04_1.Request) (*Ticket_CreateTSTFromPricing_v04_1.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.TicketCreateTSTFromPricingV041(query)
}
