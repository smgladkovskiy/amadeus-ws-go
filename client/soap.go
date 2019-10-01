package client

import (
	"encoding/xml"
	"fmt"

	"github.com/tmconsulting/amadeus-golang-sdk/structs/session/v03.0"
)

const (
	Ns    = "http://schemas.xmlsoap.org/soap/envelope/"
	XsiNs = "http://www.w3.org/2001/XMLSchema-instance"
	XsdNs = "http://www.w3.org/2001/XMLSchema"
	WasNs = "http://www.w3.org/2005/08/addressing"
)

type Body struct {
	XMLName xml.Name    `xml:"soap:Body"`
	Fault   *Fault      `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type Header struct {
	XMLName xml.Name      `xml:"soap:Header"`
	Items   []interface{} `xml:",omitempty"`
}

type AddressingFrom struct {
	// XMLName          xml.Name `xml:"wsa:From"`
	Address string `xml:"Address"`
}

type ResponseSOAPBody struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   *Fault      `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

func (b *ResponseSOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &Fault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

type Fault struct {
	XMLName xml.Name `xml:"Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}

func (f *Fault) Error() string {
	return fmt.Sprintf("code: %s; error: %s", f.Code, f.String)
}

type RequestSOAP4Envelope struct {
	XMLName  xml.Name `xml:"soap:Envelope"`
	SOAPAttr string   `xml:"xmlns:soap,attr"`
	XSIAttr  string   `xml:"xmlns:xsi,attr"`
	XSDAttr  string   `xml:"xmlns:xsd,attr"`
	Header   *RequestSOAP4Header
	Body     Body
}

type ResponseSOAP4Envelope struct {
	// XMLName        xml.Name                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header ResponseSOAPHeader
	Body   ResponseSOAPBody
}

type RequestSOAP4Header struct {
	XMLName     xml.Name                      `xml:"soap:Header"`
	WSAAttr     string                        `xml:"xmlns:wsa,attr"`
	To          string                        `xml:"wsa:To"`
	Action      string                        `xml:"wsa:Action"`
	MessageId   string                        `xml:"wsa:MessageID"`
	Security    *WSSSecurityHeader            // `xml:"wsse:Security"`
	AMASecurity *AMASecurityHostedUser        // `xml:"AMA_SecurityHostedUser"`
	Session     *Session_v03_0.RequestSession // `xml:"Session"`
}

type ResponseSOAPHeader struct {
	// XMLName        xml.Name                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ soap:Header"`
	To        string                `xml:"To"`
	From      AddressingFrom        `xml:"From"`
	Action    string                `xml:"Action"`
	MessageId string                `xml:"MessageID"`
	RelatesTo string                `xml:"RelatesTo"`
	Session   Session_v03_0.Session `xml:"Session"`
	Items     []interface{}         `xml:",omitempty"`
}
