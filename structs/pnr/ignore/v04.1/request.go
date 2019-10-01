package PNR_Ignore_v04_1 // cltreq041

import "encoding/xml"

type Request struct {
	XMLName          xml.Name          `xml:"http://xml.amadeus.com/CLTREQ_04_1_IA PNR_Ignore"`
	ClearInformation *ClearInformation `xml:"clearInformation"`
}
