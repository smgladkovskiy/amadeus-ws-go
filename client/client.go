package client

import (
	"github.com/tmconsulting/amadeus-golang-sdk/logger"
	"github.com/tmconsulting/amadeus-golang-sdk/logger/nilLogger"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/session/v03.0"
)

var (
	soapUrl = "http://webservices.amadeus.com/WSAP/"
	amaUrl  = "http://webservices.amadeus.com/"
)

func New(opts ...Option) *AmadeusClient {
	c := &AmadeusClient{
		Session: CreateSession(),
		service: &WebServicePT{
			Client: &SOAP4Client{
				TLSEnable: true,
				Logger:    logger.NewLogger(nilLogger.Init()),
			},
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type AmadeusClient struct {
	Session *Session_v03_0.Session
	service *WebServicePT
}

func (client *AmadeusClient) Close() error {
	if client == nil || client.service == nil {
		return nil
	}
	if client.Session != nil && client.Session.TransactionStatusCode != Session_v03_0.TransactionStatusCode[Session_v03_0.End] {
		client.Session.TransactionStatusCode = Session_v03_0.TransactionStatusCode[Session_v03_0.End]
		if _, _, err := client.SecuritySignOutV041(); err != nil {
			return err
		}
		client.Session = nil
	}
	return nil
}
