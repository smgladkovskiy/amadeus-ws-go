# Amadeus WBS SDK

This package contains structures, forms, functions and SOAP handler for Amadeus WS.

## Methods implementation progress

- [ ] Air_FlightInfo (05.1)
- [x] Air_SellFromRecommendation (05.2)
- [x] Command_Cryptic (07.3)
- [x] DocIssuance_IssueTicket (09.1)
- [x] Fare_CheckRules (07.1)
- [ ] Fare_ConvertCurrency (08.1)
- [x] Fare_InformativeBestPricingWithoutPNR (12.4)
- [x] Fare_InformativePricingWithoutPNR (12.4)
- [ ] Fare_MasterPricerCalendar (14.3, 12.2)
- [x] Fare_MasterPricerTravelBoardSearch (16.3, 14.3)
- [x] Fare_PricePNRWithBookingClass (14.1)
- [x] PNR_AddMultiElements (11.3)
- [x] PNR_Cancel (11.3)
- [x] PNR_Retrieve (11.3)
- [ ] Queue_CountPlanner (03.1)
- [ ] Queue_CountTotal (03.1)
- [ ] Queue_List (11.1)
- [ ] Queue_MoveItem (03.1)
- [ ] Queue_PlacePNR (03.1)
- [ ] Queue_RemoveItem (03.1)
- [x] Security_Authenticate (06.1)
- [x] Security_SignOut (04.1)
- [x] Ticket_CreateTSTFromPricing (04.1)
- [ ] Ticket_CreditCardCheck (06.1)
- [x] Ticket_DeleteTST (04.1)
- [x] Ticket_DisplayTST (07.1)
- [x] Ticket_ProcessEDoc (15.2)

## Installation

It is go gettable and go.mod powered

    $ go get github.com/tmconsulting/amadeus-golang-sdk@latest

## Usage

Prepare log writer realisation if you need to see outgoing and incoming xmls (null logging used if nil is passed), 
 methods map that you want (and allowed by Amadeus) to run (use `GetLatestMethodsMap()` to use latest methods versions 
 that are implemented) and credentials to connect: url, originator, password (not in base64!). Initiate SDK and service:

```go
package main

import (
	"log"
	
	"github.com/tmconsulting/amadeus-golang-sdk/logger/stdoutLogger"
	"github.com/tmconsulting/amadeus-golang-sdk/sdk"
	"github.com/tmconsulting/amadeus-golang-sdk/service"
)

func main() {
	url := "https://nodeD1.test.webservices.amadeus.com/1ASIWXXXXXX"
 	originator := "WSBENXXX"
 	passwordRaw := "dGhlIHBhc3N3b3Jk"
 	officeID := "BRUXX1111"
 	client := sdk.CreateAmadeusClient(url, originator, passwordRaw, officeID, stdoutLogger.Init())

 	amadeusSDK := service.NewSKD(client, service.GetLatestMethodsMap())

 	response, err := amadeusSDK.CommandCryptic("AN20MAYMOWLED/ALH")
 	if err != nil {
  	log.Fatalf("error: %v", err)
 	}
  
 	log.Printf("response: %v\n", response)
}
```

## Testing

Create test `.env` file from [test.env-example](test.env-example), or run with ENV variables `make test` 
(or `go test ./... -v`)

## Contribution

Contribution, in any kind of way, is highly welcome!
It doesn't matter if you are not able to write code.
Creating issues or holding talks and help other people to use 
[amadeus-golang-sdk](https://github.com/tmconsulting/amadeus-golang-sdk) is contribution, too!

A few examples:

* Correct typos in the README / documentation
* Reporting bugs
* Implement a new feature or endpoint
* Sharing the love if like to use [amadeus-golang-sdk](https://github.com/tmconsulting/amadeus-golang-sdk) and help people 
to get use to it

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

## License

SDK is released under the [MIT License](./LICENSE).
