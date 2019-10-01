package service

import "errors"

type MethodVersion int
type MethodName int
type MethodsMap map[MethodName]MethodVersion

var ErrNoRealisation = errors.New("sorry, method has no realisation")

// Methods versions sdk realisation
const (
	PNRRetrieveV113 MethodVersion = iota
	TicketDisplayTSTV071
	FareInformativePricingWithoutPNRV124
	FareCheckRulesV071
	CommandCrypticV073
	SecuritySignOutV041
	UpdateSessionV030
	CloseSessionV041
	FareMasterPricerTravelBoardSearchV143
	FareMasterPricerTravelBoardSearchV163
	FareInformativeBestPricingWithoutPNRV124
	AirSellFromRecommendationV052
	PNRAddMultiElementsV113
	FarePricePNRWithBookingClassV141
	TicketCreateTSTFromPricingV041
	PNRCancelV113
	DocIssuanceIssueTicketV091
	SalesReportsDisplayQueryReportV101
	TicketCancelDocumentV111
	TicketDeleteTSTV041
	AMATicketInitRefundV030
	AMATicketIgnoreRefundV030
	AMATicketProcessRefundV030
	TicketProcessEDocV152
	SalesReportsDisplayTransactionReportV132
	PNRIgnoreV041
)

//MethodsMap Represents methods that have realisation in current version of sdk
const (
	PNRRetrieve MethodName = iota
	TicketDisplayTST
	FareInformativePricingWithoutPNR
	FareCheckRules
	CommandCryptic
	SecuritySignOut
	UpdateSession
	CloseSession
	FareMasterPricerTravelBoardSearch
	FareInformativeBestPricingWithoutPNR
	AirSellFromRecommendation
	PNRAddMultiElements
	FarePricePNRWithBookingClass
	TicketCreateTSTFromPricing
	PNRCancel
	DocIssuanceIssueTicket
	SalesReportsDisplayQueryReport
	TicketCancelDocument
	TicketDeleteTST
	AMATicketInitRefund
	AMATicketIgnoreRefund
	AMATicketProcessRefund
	TicketProcessEDoc
	SalesReportsDisplayTransactionReport
	PNRIgnore
)

//GetLatest Generates latest actual methods map
func GetLatestMethodsMap() MethodsMap {
	return MethodsMap{
		PNRRetrieve:                          PNRRetrieveV113,
		TicketDisplayTST:                     TicketDisplayTSTV071,
		FareInformativePricingWithoutPNR:     FareInformativePricingWithoutPNRV124,
		FareCheckRules:                       FareCheckRulesV071,
		CommandCryptic:                       CommandCrypticV073,
		SecuritySignOut:                      SecuritySignOutV041,
		UpdateSession:                        UpdateSessionV030,
		CloseSession:                         CloseSessionV041,
		FareMasterPricerTravelBoardSearch:    FareMasterPricerTravelBoardSearchV143,
		FareInformativeBestPricingWithoutPNR: FareInformativeBestPricingWithoutPNRV124,
		AirSellFromRecommendation:            AirSellFromRecommendationV052,
		PNRAddMultiElements:                  PNRAddMultiElementsV113,
		FarePricePNRWithBookingClass:         FarePricePNRWithBookingClassV141,
		TicketCreateTSTFromPricing:           TicketCreateTSTFromPricingV041,
		PNRCancel:                            PNRCancelV113,
		DocIssuanceIssueTicket:               DocIssuanceIssueTicketV091,
		SalesReportsDisplayQueryReport:       SalesReportsDisplayQueryReportV101,
		TicketCancelDocument:                 TicketCancelDocumentV111,
		TicketDeleteTST:                      TicketDeleteTSTV041,
		AMATicketInitRefund:                  AMATicketInitRefundV030,
		AMATicketIgnoreRefund:                AMATicketIgnoreRefundV030,
		AMATicketProcessRefund:               AMATicketProcessRefundV030,
		TicketProcessEDoc:                    TicketProcessEDocV152,
		SalesReportsDisplayTransactionReport: SalesReportsDisplayTransactionReportV132,
		PNRIgnore:                            PNRIgnoreV041,
	}
}
