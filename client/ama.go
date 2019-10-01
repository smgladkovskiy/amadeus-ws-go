package client

import "encoding/xml"

type AMASecurityHostedUser struct {
	XMLName xml.Name                    `xml:"http://xml.amadeus.com/2010/06/Security_v1 AMA_SecurityHostedUser"`
	UserId  AMASecurityHostedUserUserID `xml:"UserID"`
}

type AMASecurityHostedUserUserID struct {
	XMLName        xml.Name `xml:"UserID"`
	AgentDutyCode  string   `xml:"AgentDutyCode,attr"`
	RequestorType  string   `xml:"RequestorType,attr"`
	PseudoCityCode string   `xml:"PseudoCityCode,attr"`
	POSType        string   `xml:"POS_Type,attr"`
}
