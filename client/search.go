package client

import (
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v16.3/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v16.3/response"
)

func (client *AmadeusClient) FareMasterPricerTravelBoardSearchV143(query *Fare_MasterPricerTravelBoardSearchRequest_v14_3.Request) (*Fare_MasterPricerTravelBoardSearchResponse_v14_3.Response, *ResponseSOAPHeader, error) {
	var reply Fare_MasterPricerTravelBoardSearchResponse_v14_3.Response
	header, err := client.service.Call(soapUrl, "FMPTBQ_14_3_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) FareMasterPricerTravelBoardSearchV163(query *Fare_MasterPricerTravelBoardSearchRequest_v16_3.Request) (*Fare_MasterPricerTravelBoardSearchResponse_v16_3.Response, *ResponseSOAPHeader, error) {
	var reply Fare_MasterPricerTravelBoardSearchResponse_v16_3.Response
	header, err := client.service.Call(soapUrl, "FMPTBQ_14_3_1A", query, &reply, client)
	return &reply, header, err
}

func (client *AmadeusClient) FareInformativeBestPricingWithoutPNRV124(query *Fare_InformativeBestPricingWithoutPNRRequest_v12_4.Request) (*Fare_InformativeBestPricingWithoutPNRResponse_v12_4.Response, *ResponseSOAPHeader, error) {
	var reply Fare_InformativeBestPricingWithoutPNRResponse_v12_4.Response
	header, err := client.service.Call(soapUrl, "TIBNRQ_12_4_1A", query, &reply, client)
	return &reply, header, err
}
