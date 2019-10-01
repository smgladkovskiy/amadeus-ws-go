package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/client"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/response"
)

func (s *service) FareMasterPricerTravelBoardSearch(query *Fare_MasterPricerTravelBoardSearchRequest_v14_3.Request) (*Fare_MasterPricerTravelBoardSearchResponse_v14_3.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.FareMasterPricerTravelBoardSearchV143(query)
}

func (s *service) FareInformativeBestPricingWithout(query *Fare_InformativeBestPricingWithoutPNRRequest_v12_4.Request) (*Fare_InformativeBestPricingWithoutPNRResponse_v12_4.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.FareInformativeBestPricingWithoutPNRV124(query)
}

func (s *service) FareInformativePricingWithoutPNR(query *Fare_InformativePricingWithoutPNR_v12_4.Request) (*Fare_InformativePricingWithoutPNRReply_v12_4.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.FareInformativePricingWithoutPNRV124(query)
}
