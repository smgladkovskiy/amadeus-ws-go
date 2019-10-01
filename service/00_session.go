package service

import (
	"github.com/tmconsulting/amadeus-golang-sdk/client"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/security/signOut/v04.1"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/session/v03.0"
)

func (s *service) SecuritySignOut() (*SecuritySignOut_v04_1.Response, *client.ResponseSOAPHeader, error) {
	return s.sdk.SecuritySignOutV041()
}

func (s *service) Session() *Session_v03_0.Session {
	return s.sdk.GetSession()
}

func (s *service) IncSequenceNumber(header *client.ResponseSOAPHeader) {
	s.sdk.IncSessionSequenceNumber(header)
}

func (s *service) IfSessionIsClosed() bool {
	return s.sdk.CheckIfSessionIsClosed()
}

func (s *service) SessionEndTransaction() bool {
	return s.sdk.SetSessionEndTransaction()
}

func (s *service) UpdateSession(session *Session_v03_0.Session) bool {
	return s.sdk.UpdateSessionV030(session)
}

func (s *service) CloseSession() (reply *SecuritySignOut_v04_1.Response, err error) {
	return s.sdk.CloseSessionV041()
}
