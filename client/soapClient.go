package client

import (
	"bytes"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/tmconsulting/amadeus-golang-sdk/logger"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/session/v03.0"
)

var timeout = 20 * time.Second

type SOAP4Client struct {
	Url       string
	User      string
	Pass      string
	Agent     string
	TLSEnable bool
	Headers   []interface{}
	Logger    logger.Service
}

func (s *SOAP4Client) AddHeader(header interface{}) {
	s.Headers = append(s.Headers, header)
}

func (s *SOAP4Client) UpdateHeader(header interface{}) {
	s.Headers = []interface{}{header}
}

func (s *SOAP4Client) GetHashedPassword(nonceRaw string, timestamp string) string {
	// HashedPassword = Base64 ( SHA-1 ( nonce + created + SHA-1 ( password ) ) )
	sha := sha1.New()
	sha.Write([]byte(s.Pass))
	pwd := string(sha.Sum(nil))
	sha.Reset()
	sha.Write([]byte(nonceRaw + timestamp + pwd))
	return base64.StdEncoding.EncodeToString(sha.Sum(nil))
}

func (s *SOAP4Client) MakeNewWSSSecurityHeader() *WSSSecurityHeader {
	nonce := RandStringBytesMaskImprSrc(16)
	nonceBase64 := base64.StdEncoding.EncodeToString([]byte(nonce))
	timestamp := time.Now().UTC().Format(time.RFC3339)

	return &WSSSecurityHeader{
		XmlNSWsse:      WssNsWSSE,
		XmlNSWsu:       WssNsWSU,
		MustUnderstand: "",
		Token: &WSSUsernameToken{
			Id:       "UsernameToken-" + RandStringBytesMaskImprSrc(9),
			Username: s.User,
			Nonce:    &WSSNonce{EncodingType: WssNsEncodingType, Nonce: nonceBase64},
			Password: &WSSPassword{Type: WssNsPasswordType, Password: s.GetHashedPassword(nonce, timestamp)},
			Created:  timestamp,
		},
	}
}

func (s *SOAP4Client) MakeNewAMASecurityHostedUser() *AMASecurityHostedUser {
	return &AMASecurityHostedUser{UserId: AMASecurityHostedUserUserID{
		AgentDutyCode:  "SU",
		RequestorType:  "U",
		PseudoCityCode: s.Agent,
		POSType:        "1",
	}}
}

func (s *SOAP4Client) Call(soapUrl, soapAction, messageId string, query, reply interface{}, amadeusClient *AmadeusClient) (*ResponseSOAPHeader, error) {
	session := amadeusClient.Session
	envelope := RequestSOAP4Envelope{
		SOAPAttr: Ns, XSIAttr: XsiNs, XSDAttr: XsdNs,
		Header: &RequestSOAP4Header{WSAAttr: WasNs, To: s.Url, Action: soapUrl + soapAction, MessageId: messageId},
	}
	if session == nil || session.TransactionStatusCode == Session_v03_0.TransactionStatusCode[Session_v03_0.Start] {
		envelope.Header.Security = s.MakeNewWSSSecurityHeader()
		envelope.Header.AMASecurity = s.MakeNewAMASecurityHostedUser()
	}
	if session != nil {
		envelope.Header.Session = &Session_v03_0.RequestSession{Session: session}
	}

	envelope.Body.Content = query
	buffer := new(bytes.Buffer)
	buffer.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>"))

	encoder := xml.NewEncoder(buffer)
	// encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return nil, err
	}

	if err := encoder.Flush(); err != nil {
		return nil, err
	}
	savebuf := buffer.Bytes()

	if err := s.Logger.Push("out", soapAction, string(savebuf)); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to fire hook log_request: %v\n", err)
	}

	req, err := http.NewRequest("POST", s.Url, buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapUrl+soapAction)

	req.Header.Set("User-Agent", "amadeus-golang-sdk")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.TLSEnable,
		},
		Dial: dialTimeout,
	}

	numberOfAttempts := 3

	var res *http.Response
	httpClient := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	res, err = httpClient.Do(req)
	for err != nil {
		time.Sleep(1 * time.Second)

		buffer := new(bytes.Buffer)
		buffer.Write(savebuf)
		rc := ioutil.NopCloser(buffer)
		req.Body = rc // .(io.ReadCloser)

		httpClient = &http.Client{
			Transport: tr,
			Timeout:   timeout,
		}
		res, err = httpClient.Do(req)

		numberOfAttempts--
		if numberOfAttempts == 0 {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = s.Logger.Push("inc", soapAction, string(rawbody))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to fire hook log_response: %v\n", err)
	}

	respEnvelope := new(ResponseSOAP4Envelope)
	respEnvelope.Body = ResponseSOAPBody{Content: reply}

	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return nil, err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return &respEnvelope.Header, fault
	}

	return &respEnvelope.Header, nil
}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}
