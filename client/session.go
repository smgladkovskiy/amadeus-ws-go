package client

import (
	"errors"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/security/signOut/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/session/v03.0"
)

func (client *AmadeusClient) SecuritySignOutV041() (*SecuritySignOut_v04_1.Response, *ResponseSOAPHeader, error) {
	var query SecuritySignOut_v04_1.Request
	var reply SecuritySignOut_v04_1.Response
	header, err := client.service.Call(soapUrl, "VLSSOQ_04_1_1A", &query, &reply, client)
	return &reply, header, err
}

func CreateSession() *Session_v03_0.Session {
	return &Session_v03_0.Session{
		TransactionStatusCode: Session_v03_0.TransactionStatusCode[Session_v03_0.Start],
	}
}

func (client *AmadeusClient) GetSession() *Session_v03_0.Session {
	return client.Session
}

func (client *AmadeusClient) IncSessionSequenceNumber(header *ResponseSOAPHeader) {
	if header != nil {
		client.Session = &header.Session
	}
	client.Session.SequenceNumber++
}

func (client *AmadeusClient) CheckIfSessionIsClosed() bool {
	if client == nil || client.Session == nil || client.Session.TransactionStatusCode == Session_v03_0.TransactionStatusCode[Session_v03_0.End] {
		return true
	}
	return false
}

func (client *AmadeusClient) SetSessionEndTransaction() bool {
	if client == nil || client.Session == nil || client.Session.TransactionStatusCode == Session_v03_0.TransactionStatusCode[Session_v03_0.End] {
		return false
	}
	client.Session.TransactionStatusCode = Session_v03_0.TransactionStatusCode[Session_v03_0.End]
	return true
}

func (client *AmadeusClient) UpdateSessionV030(session *Session_v03_0.Session) bool {
	if client == nil {
		return false
	}
	client.Session = session
	return true
}

func (client *AmadeusClient) CloseSessionV041() (*SecuritySignOut_v04_1.Response, error) {
	if client == nil || client.service == nil {
		return nil, errors.New("nil client or service")
	}

	if client.Session != nil && client.Session.TransactionStatusCode != Session_v03_0.TransactionStatusCode[Session_v03_0.End] {
		client.Session.TransactionStatusCode = Session_v03_0.TransactionStatusCode[Session_v03_0.End]
		reply, header, err := client.SecuritySignOutV041()
		if err == nil {
			client.Session = nil
		} else {
			client.Session = &header.Session
		}
		return reply, err
	}
	return nil, errors.New("can't close session: Session is nil or Session.TransactionStatusCode == End")
}
