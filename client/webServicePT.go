package client

import (
	"strings"
)

type WebServicePT struct {
	Client *SOAP4Client
	// wsap   string
}

func (w *WebServicePT) Call(soapUrl, soapAction string, query, reply interface{}, client *AmadeusClient) (*ResponseSOAPHeader, error) {
	messageId := strings.ToUpper(RandStringBytesMaskImprSrc(22))
	return w.Client.Call(soapUrl, soapAction, messageId, query, reply, client)
}

func (w *WebServicePT) AddHeader(header interface{}) {
	w.Client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
// Deprecated
func (w *WebServicePT) SetHeader(header interface{}) {
	w.AddHeader(header)
}

func (w *WebServicePT) UpdateHeader(header interface{}) {
	w.Client.UpdateHeader(header)
}
