package client

import "encoding/xml"

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE         = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU          = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsEncodingType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"
	WssNsPasswordType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest"
)

type WSSSecurityHeader struct {
	XMLName        xml.Name          `xml:"wsse:Security"`
	XmlNSWsse      string            `xml:"xmlns:wsse,attr"`
	XmlNSWsu       string            `xml:"xmlns:wsu,attr"`
	MustUnderstand string            `xml:"mustUnderstand,attr,omitempty"`
	Token          *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName  xml.Name     `xml:"wsse:UsernameToken"`
	Id       string       `xml:"wsu:Id,attr,omitempty"`
	Username string       `xml:"wsse:Username"`
	Nonce    *WSSNonce    //`xml:"wsse:Nonce"`
	Password *WSSPassword //`xml:"wsse:Password"`
	Created  string       `xml:"wsu:Created"`
}

type WSSNonce struct {
	XMLName      xml.Name `xml:"wsse:Nonce"`
	EncodingType string   `xml:"EncodingType,attr"`
	Nonce        string   `xml:",chardata"`
}

type WSSPassword struct {
	XMLName  xml.Name `xml:"wsse:Password"`
	Type     string   `xml:"Type,attr"`
	Password string   `xml:",chardata"`
}
