package PNR_Ignore_v04_1 // cltres041

//import "encoding/xml"

type Response struct {
	//XMLName         xml.Name          `xml:"http://xml.amadeus.com/CLTRES_04_1_IA PNR_IgnoreReply"`
	SbrRecLoc        *SbrRecLoc        `xml:"sbrRecLoc,omitempty"`
	ClearInformation *ClearInformation `xml:"clearInformation"`
}

type SbrRecLoc struct {
	Reservation *Reservation `xml:"reservation"`
}

type Reservation struct {
	ControlNumber string `xml:"controlNumber"`
}

type ClearInformation struct {
	ActionRequest string `xml:"actionRequest"`
}
