package SecuritySignOut_v04_1 // vlssoq041

import (
	"encoding/xml"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
)

type Request struct {
	XMLName         xml.Name            `xml:"http://xml.amadeus.com/VLSSOQ_04_1_1A Security_SignOut"`
	ConversationClt *ConversationIDType `xml:"conversationClt,omitempty"` // It contains conversation properties between the SI and the JFE.
}

//
// Complex structs
//

type ConversationIDType struct {
	SenderIdentification                 formats.AlphaNumericString_Length1To35 `xml:"senderIdentification"`                 // Sender identification
	RecipientIdentification              formats.AlphaNumericString_Length1To35 `xml:"recipientIdentification"`              // Recipient identification
	SenderInterchangeControlReference    formats.AlphaNumericString_Length1To14 `xml:"senderInterchangeControlReference"`    // Sender's interchange control reference
	RecipientInterchangeControlReference formats.AlphaNumericString_Length1To14 `xml:"recipientInterchangeControlReference"` // Recipient's interchange control reference
}
