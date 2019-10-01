package CommandCryptic_v07_3 // hsfreq073

import (
	"encoding/xml"
)

type Request struct {
	XMLName                    xml.Name                    `xml:"http://xml.amadeus.com/HSFREQ_07_3_1A Command_Cryptic"`
	MessageAction              *MessageAction              `xml:"messageAction"`
	NumberOfUnits              *NumberOfUnits              `xml:"numberOfUnits,omitempty"`
	IntelligentWorkstationInfo *IntelligentWorkstationInfo `xml:"intelligentWorkstationInfo,omitempty"`
	LongTextString             *LongTextString             `xml:"longTextString"`
}

type MessageAction struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails"`
	ResponseType           string                  `xml:"responseType,omitempty"`
}

type MessageFunctionDetails struct {
	BusinessFunction          string   `xml:"businessFunction,omitempty"`
	MessageFunction           string   `xml:"messageFunction"`
	AdditionalMessageFunction []string `xml:"additionalMessageFunction,omitempty"` // maxOccurs="20"
}

type NumberOfUnits struct {
	NumberOfUnitsDetails1 *NumberOfUnitsDetails1   `xml:"numberOfUnitsDetails1"`
	NumberOfUnitsDetails2 []*NumberOfUnitsDetails2 `xml:"numberOfUnitsDetails2,omitempty"` // maxOccurs="8"
}

type NumberOfUnitsDetails1 struct {
	Units          float64 `xml:"units"`
	UnitsQualifier string  `xml:"unitsQualifier,omitempty"`
}

type NumberOfUnitsDetails2 struct {
	Units          *float64 `xml:"units,omitempty"`
	UnitsQualifier string   `xml:"unitsQualifier,omitempty"`
}

type IntelligentWorkstationInfo struct {
	CompanyIdentification string `xml:"companyIdentification,omitempty"`
}
