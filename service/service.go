package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/client"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/air/sellFromRecommendation/v05.2/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/air/sellFromRecommendation/v05.2/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketIgnoreRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketInitRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ama/ticketProcessRefund/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/commandCryptic"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/commandCryptic/v07.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/docIssuance/issueTicket/v09.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/checkRules/v07.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/checkRules/v07.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativeBestPricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/informativePricingWithoutPNR/v12.4/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v14.3/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v16.3/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch/v16.3/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/pricePNRWithBookingClass/v14.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/fare/pricePNRWithBookingClass/v14.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/addMultiElements/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/cancel/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/ignore/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/reply/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/pnr/retrieve/v11.3"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayQueryReport/v10.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayQueryReport/v10.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/salesReports/displayTransactionReport/v13.2"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/security/signOut/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/session/v03.0"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/cancelDocument/v11.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/createTSTFromPricing/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/deleteTST/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/displayTST/v07.1/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/displayTST/v07.1/response"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/processEDoc/v15.2/request"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/ticket/processEDoc/v15.2/response"
)

type Client struct {
	url     string
	user    string
	pass    string
	agent   string
	tls     bool
	headers []interface{}
}

func New(sdk AmadeusSDK, opts ...Option) Service {
	srv := &service{sdk: sdk}

	srv.mm = GetLatestMethodsMap()

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

type service struct {
	mm  MethodsMap
	sdk AmadeusSDK
}

type AmadeusSDK interface {
	// Information
	PNRRetrieveV113(query *PNR_Retrieve_v11_3.Request) (*PNR_Reply_v11_3.Response, *client.ResponseSOAPHeader, error)
	TicketDisplayTSTV071(query *Ticket_DisplayTSTRequest_v07_1.Request) (*Ticket_DisplayTSTResponse_v07_1.Response, *client.ResponseSOAPHeader, error)
	FareInformativePricingWithoutPNRV124(query *Fare_InformativePricingWithoutPNR_v12_4.Request) (*Fare_InformativePricingWithoutPNRReply_v12_4.Response, *client.ResponseSOAPHeader, error)
	FareCheckRulesV071(query *Fare_CheckRulesRequest_v07_1.Request) (*Fare_CheckRulesResponse_v07_1.Response, *client.ResponseSOAPHeader, error)
	CommandCrypticV073(query *CommandCryptic_v07_3.Request) (*CommandCryptic_v07_3.Response, error)

	// Session
	SecuritySignOutV041() (*SecuritySignOut_v04_1.Response, *client.ResponseSOAPHeader, error)
	GetSession() *Session_v03_0.Session
	IncSessionSequenceNumber(header *client.ResponseSOAPHeader)
	CheckIfSessionIsClosed() bool
	SetSessionEndTransaction() bool
	UpdateSessionV030(session *Session_v03_0.Session) bool
	CloseSessionV041() (reply *SecuritySignOut_v04_1.Response, err error)

	// Search
	FareMasterPricerTravelBoardSearchV143(query *Fare_MasterPricerTravelBoardSearchRequest_v14_3.Request) (*Fare_MasterPricerTravelBoardSearchResponse_v14_3.Response, *client.ResponseSOAPHeader, error)
	FareMasterPricerTravelBoardSearchV163(query *Fare_MasterPricerTravelBoardSearchRequest_v16_3.Request) (*Fare_MasterPricerTravelBoardSearchResponse_v16_3.Response, *client.ResponseSOAPHeader, error)
	FareInformativeBestPricingWithoutPNRV124(query *Fare_InformativeBestPricingWithoutPNRRequest_v12_4.Request) (*Fare_InformativeBestPricingWithoutPNRResponse_v12_4.Response, *client.ResponseSOAPHeader, error)

	// Book
	AirSellFromRecommendationV052(query *Air_SellFromRecommendationRequest_v05_2.Request) (*Air_SellFromRecommendationResponse_v05_2.Response, *client.ResponseSOAPHeader, error)
	PNRAddMultiElementsV113(query *PNR_AddMultiElementsRequest_v11_3.Request) (*PNR_Reply_v11_3.Response, *client.ResponseSOAPHeader, error)
	FarePricePNRWithBookingClassV141(query *Fare_PricePNRWithBookingClassRequest_v14_1.Request) (*Fare_PricePNRWithBookingClassResponse_v14_1.Response, *client.ResponseSOAPHeader, error)
	TicketCreateTSTFromPricingV041(query *Ticket_CreateTSTFromPricing_v04_1.Request) (*Ticket_CreateTSTFromPricing_v04_1.Response, *client.ResponseSOAPHeader, error)
	PNRCancelV113(query *PNR_Cancel_v11_3.Request) (*PNR_Reply_v11_3.Response, *client.ResponseSOAPHeader, error)

	// Issue
	DocIssuanceIssueTicketV091(query *DocIssuance_IssueTicket_v09_1.Request) (*DocIssuance_IssueTicket_v09_1.Response, *client.ResponseSOAPHeader, error)

	// Void
	SalesReportsDisplayQueryReportV101(query *SalesReports_QueryReportRequest_v10_1.Request) (*SalesReports_QueryReportReply_v10_1.Response, *client.ResponseSOAPHeader, error)
	TicketCancelDocumentV111(query *Ticket_CancelDocument_v11_1.Request) (*Ticket_CancelDocument_v11_1.Response, *client.ResponseSOAPHeader, error)
	TicketDeleteTSTV041(query *Ticket_DeleteTST_v04_1.Request) (*Ticket_DeleteTST_v04_1.Response, *client.ResponseSOAPHeader, error)

	// Refund
	AMATicketInitRefundV030(query *AMA_TicketInitRefund_v03_0.Request) (*AMA_TicketInitRefund_v03_0.Response, *client.ResponseSOAPHeader, error)
	AMATicketIgnoreRefundV030(query *AMA_TicketIgnoreRefund_v03_0.Request) (*AMA_TicketIgnoreRefund_v03_0.Response, *client.ResponseSOAPHeader, error)
	AMATicketProcessRefundV030(query *AMA_TicketProcessRefund_v03_0.Request) (*AMA_TicketProcessRefund_v03_0.Response, *client.ResponseSOAPHeader, error)
	TicketProcessEDocV152(query *Ticket_ProcessEDocRequest_v15_2.Request) (*Ticket_ProcessEDocResponse_v15_2.Response, *client.ResponseSOAPHeader, error)
	SalesReportsDisplayTransactionReportV132(query *SalesReports_DisplayTransactionReport_v13_2.Request) (*SalesReports_DisplayTransactionReport_v13_2.Response, *client.ResponseSOAPHeader, error)
	PNRIgnoreV041(query *PNR_Ignore_v04_1.Request) (*PNR_Ignore_v04_1.Response, *client.ResponseSOAPHeader, error)
}

type Service interface {
	// Information
	PNRRetrieve(query *PNR_Retrieve_v11_3.Request) (*PNR_Reply_v11_3.Response, *client.ResponseSOAPHeader, error)
	TicketDisplayTST(query *Ticket_DisplayTSTRequest_v07_1.Request) (*Ticket_DisplayTSTResponse_v07_1.Response, *client.ResponseSOAPHeader, error)
	FareCheckRules(query *Fare_CheckRulesRequest_v07_1.Request) (*Fare_CheckRulesResponse_v07_1.Response, *client.ResponseSOAPHeader, error)
	CommandCryptic(msg string) (*commandCryptic.Response, error)

	// Session
	SecuritySignOut() (*SecuritySignOut_v04_1.Response, *client.ResponseSOAPHeader, error)
	Session() *Session_v03_0.Session
	IncSequenceNumber(header *client.ResponseSOAPHeader)
	IfSessionIsClosed() bool
	SessionEndTransaction() bool
	UpdateSession(session *Session_v03_0.Session) bool
	CloseSession() (reply *SecuritySignOut_v04_1.Response, err error)

	// Search
	FareMasterPricerTravelBoardSearch(query *Fare_MasterPricerTravelBoardSearchRequest_v14_3.Request) (*Fare_MasterPricerTravelBoardSearchResponse_v14_3.Response, *client.ResponseSOAPHeader, error)
	FareInformativeBestPricingWithout(query *Fare_InformativeBestPricingWithoutPNRRequest_v12_4.Request) (*Fare_InformativeBestPricingWithoutPNRResponse_v12_4.Response, *client.ResponseSOAPHeader, error)
	FareInformativePricingWithoutPNR(query *Fare_InformativePricingWithoutPNR_v12_4.Request) (*Fare_InformativePricingWithoutPNRReply_v12_4.Response, *client.ResponseSOAPHeader, error)

	// Book
	AirSellFromRecommendation(query *Air_SellFromRecommendationRequest_v05_2.Request) (*Air_SellFromRecommendationResponse_v05_2.Response, *client.ResponseSOAPHeader, error)
	PNRAddMultiElements(query *PNR_AddMultiElementsRequest_v11_3.Request) (*PNR_Reply_v11_3.Response, *client.ResponseSOAPHeader, error)
	FarePricePNRWithBookingClass(query *Fare_PricePNRWithBookingClassRequest_v14_1.Request) (*Fare_PricePNRWithBookingClassResponse_v14_1.Response, *client.ResponseSOAPHeader, error)
	TicketCreateTSTFromPricing(query *Ticket_CreateTSTFromPricing_v04_1.Request) (*Ticket_CreateTSTFromPricing_v04_1.Response, *client.ResponseSOAPHeader, error)

	// Cancellation
	PNRCancel(query *PNR_Cancel_v11_3.Request) (*PNR_Reply_v11_3.Response, *client.ResponseSOAPHeader, error)

	// Issue
	DocIssuanceIssueTicket(query *DocIssuance_IssueTicket_v09_1.Request) (*DocIssuance_IssueTicket_v09_1.Response, *client.ResponseSOAPHeader, error)

	// Void
	SalesReportsDisplayQueryReport(query *SalesReports_QueryReportRequest_v10_1.Request) (*SalesReports_QueryReportReply_v10_1.Response, *client.ResponseSOAPHeader, error)
	TicketCancelDocument(query *Ticket_CancelDocument_v11_1.Request) (*Ticket_CancelDocument_v11_1.Response, *client.ResponseSOAPHeader, error)
	TicketDeleteTST(query *Ticket_DeleteTST_v04_1.Request) (*Ticket_DeleteTST_v04_1.Response, *client.ResponseSOAPHeader, error)

	// Refund
	RefundIgnore(query *AMA_TicketIgnoreRefund_v03_0.Request) (*AMA_TicketIgnoreRefund_v03_0.Response, *client.ResponseSOAPHeader, error)
	RefundInit(query *AMA_TicketInitRefund_v03_0.Request) (*AMA_TicketInitRefund_v03_0.Response, *client.ResponseSOAPHeader, error)
	RefundProcess(query *AMA_TicketProcessRefund_v03_0.Request) (*AMA_TicketProcessRefund_v03_0.Response, *client.ResponseSOAPHeader, error)
	TicketProcessEDoc(query *Ticket_ProcessEDocRequest_v15_2.Request) (*Ticket_ProcessEDocResponse_v15_2.Response, *client.ResponseSOAPHeader, error)
	SalesReportsDisplayTransactionReport(query *SalesReports_DisplayTransactionReport_v13_2.Request) (*SalesReports_DisplayTransactionReport_v13_2.Response, *client.ResponseSOAPHeader, error)
	PNRIgnore(query *PNR_Ignore_v04_1.Request) (*PNR_Ignore_v04_1.Response, *client.ResponseSOAPHeader, error)
}
