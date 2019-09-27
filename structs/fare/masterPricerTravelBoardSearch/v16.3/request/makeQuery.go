package Fare_MasterPricerTravelBoardSearchRequest_v16_3 // fmptbq143

import (
	search "github.com/tmconsulting/amadeus-golang-sdk/structs/fare/masterPricerTravelBoardSearch"
	"github.com/tmconsulting/amadeus-golang-sdk/structs/formats"
	"github.com/tmconsulting/amadeus-golang-sdk/utils/convert"
)

func MakeRequest(request *search.SearchRequest) *Request {

	var query = Request{
		NumberOfUnit: &NumberOfUnitsType{
			UnitNumberDetail: []*NumberOfUnitDetailsType_270113C{
				{
					NumberOfUnits: formats.NumericInteger_Length1To6(request.Passengers.ADT + request.Passengers.CHD),
					TypeOfUnit:    formats.AlphaNumericString_Length1To3("PX"),
				},
				{
					NumberOfUnits: formats.NumericInteger_Length1To6(200),
					TypeOfUnit:    formats.AlphaNumericString_Length1To3("RC"),
				},
			},
		},
		FareOptions: &FareOptions{
			PricingTickInfo: &PricingTicketingDetailsType{
				PricingTicketing: &PricingTicketingInformationType{
					PriceType: []formats.AlphaNumericString_Length0To3{
						formats.AlphaNumericString_Length0To3("XLC"), // No LCC fares
						formats.AlphaNumericString_Length0To3("CUC"),
						formats.AlphaNumericString_Length0To3("ET"),
						formats.AlphaNumericString_Length0To3("RP"),
						formats.AlphaNumericString_Length0To3("RU"),  // Unifares
						formats.AlphaNumericString_Length0To3("TAC"), // Ticket ability check
						formats.AlphaNumericString_Length0To3("NAD"), // No Airline Distribution
					},
				},
			},
			ConversionRate: &ConversionRateType{
				ConversionRateDetail: []*ConversionRateDetailsType{
					{
						Currency: formats.AlphaString_Length1To3(request.Currency),
					},
				},
			},
			FeeIdDescription: &CodedAttributeType_199259S{
				FeeId: []*CodedAttributeInformationType_277155C{
					{FeeType: formats.AlphaNumericString_Length1To5("FFI"), FeeIdNumber: formats.AlphaNumericString_Length1To50("2")},
					{FeeType: formats.AlphaNumericString_Length1To5("UPH"), FeeIdNumber: formats.AlphaNumericString_Length1To50("6")},
				},
			},
		},
		TravelFlightInfo: &TravelFlightInformationType_199258S{
			CabinId: &CabinIdentificationType_233500C{
				CabinQualifier: formats.AlphaNumericString_Length1To2("MD"),
				Cabin: []formats.AlphaString_Length0To1{
					formats.AlphaString_Length0To1("Y"),
				},
			},
		},
	}

	paxID := 1
	if request.Passengers.ADT > 0 {
		var travellers []*TravellerDetailsType
		for i := 0; i < request.Passengers.ADT; i++ {
			travellers = append(travellers, &TravellerDetailsType{
				Ref: formats.NumericInteger_Length1To3(paxID),
			})
			paxID++
		}
		query.PaxReference = append(query.PaxReference, &TravellerReferenceInformationType{
			Ptc: []formats.AlphaNumericString_Length1To6{
				"ADT",
			},
			Traveller: travellers,
		})
	}

	if request.Passengers.CHD > 0 {
		var travellers []*TravellerDetailsType
		for i := 0; i < request.Passengers.CHD; i++ {
			travellers = append(travellers, &TravellerDetailsType{
				Ref: formats.NumericInteger_Length1To3(paxID),
			})
			paxID++
		}
		query.PaxReference = append(query.PaxReference, &TravellerReferenceInformationType{
			Ptc: []formats.AlphaNumericString_Length1To6{
				"CH",
			},
			Traveller: travellers,
		})
	}
	if request.Passengers.INF > 0 {
		paxInfID := 1
		var travellers []*TravellerDetailsType
		for i := 0; i < request.Passengers.INF; i++ {
			infantIndicator := formats.NumericInteger_Length1To1(paxInfID)
			travellers = append(travellers, &TravellerDetailsType{
				Ref:             formats.NumericInteger_Length1To3(paxInfID),
				InfantIndicator: &infantIndicator,
			})
			paxInfID++
		}
		query.PaxReference = append(query.PaxReference, &TravellerReferenceInformationType{
			Ptc: []formats.AlphaNumericString_Length1To6{
				"INF",
			},
			Traveller: travellers,
		})
	}

	var routeCount = len(request.Itineraries)
	var useMultiCity = false
	//if routeCount > 3 && len(request.BaseClass) > 0 {
	//	useMultiCity = true
	//} else
	if routeCount < 4 {
		useMultiCity = true
	}
	for routeID := 1; routeID <= routeCount; routeID++ {
		var route = request.Itineraries[routeID]
		var firstDateTimeDetail = DateAndTimeDetailsTypeI{}
		if route.DepartureDate.Error == nil {
			firstDateTimeDetail.Date = formats.Date_DDMMYY(convert.DateToAmadeusDate(route.DepartureDate.Date))
			if route.DepartureDate.TimeFlag {
				firstDateTimeDetail.Time = formats.Time24_HHMM(convert.DateToAmadeusTime(route.DepartureDate.Date))
				firstDateTimeDetail.TimeQualifier = formats.AlphaNumericString_Length1To3("TD")
			}
		} else if route.ArrivalDate.Error == nil {
			firstDateTimeDetail.Date = formats.Date_DDMMYY(convert.DateToAmadeusDate(route.ArrivalDate.Date))
			if route.ArrivalDate.TimeFlag {
				firstDateTimeDetail.Time = formats.Time24_HHMM(convert.DateToAmadeusTime(route.ArrivalDate.Date))
				firstDateTimeDetail.TimeQualifier = formats.AlphaNumericString_Length1To3("TA")
			}
		}

		var departureLocation = route.DepartureLocation.AirportCode
		if departureLocation == "" {
			departureLocation = route.DepartureLocation.CityCode
		}
		var arrivalLocation = route.ArrivalLocation.AirportCode
		if arrivalLocation == "" {
			arrivalLocation = route.ArrivalLocation.CityCode
		}

		var itinerary = Itinerary{
			RequestedSegmentRef: &OriginAndDestinationRequestType{
				SegRef: formats.NumericInteger_Length1To2(routeID),
			},
			TimeDetails: &DateAndTimeInformationType_181295S{
				FirstDateTimeDetail: &firstDateTimeDetail,
			},
		}
		if useMultiCity {
			itinerary.DepartureLocalization = &DepartureLocationType{
				DepMultiCity: []*MultiCityOptionType{
					{
						LocationId: formats.AlphaString_Length3To5(departureLocation),
					},
				},
			}
			itinerary.ArrivalLocalization = &ArrivalLocalizationType{
				ArrivalMultiCity: []*MultiCityOptionType{
					{
						LocationId: formats.AlphaString_Length3To5(arrivalLocation),
					},
				},
			}
		} else {
			itinerary.DepartureLocalization = &DepartureLocationType{
				DeparturePoint: &ArrivalLocationDetailsType_120834C{
					LocationId: formats.AlphaString_Length3To5(departureLocation),
				},
			}
			itinerary.ArrivalLocalization = &ArrivalLocalizationType{
				ArrivalPointDetails: &ArrivalLocationDetailsType{
					LocationId: formats.AlphaString_Length3To5(arrivalLocation),
				},
			}
		}

		query.Itinerary = append(query.Itinerary, &itinerary)
	}

	return &query
}