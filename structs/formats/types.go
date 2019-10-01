package formats

// Used for codes in the AMADEUS code tables. Code Length is one alphanumeric character.
// pattern = "[0-9A-Z]"
type AMA_EDICodesetType_Length1 string

// Used for codes in the AMADEUS code tables. Code Length is three alphanumeric characters.
// pattern = "[0-9A-Z]{1,3}"
type AMA_EDICodesetType_Length1to3 string

// Format limitations: an..1
type AlphaNumericString_Length0To1 string

// Format limitations: an..18
type AlphaNumericString_Length0To18 string

// Format limitations: an..3
type AlphaNumericString_Length0To3 string

// Format limitations: an1
type AlphaNumericString_Length1To1 string

// Format limitations: an..10
type AlphaNumericString_Length1To10 string

// Format limitations: an..109
type AlphaNumericString_Length1To109 string

// Format limitations: an..11
type AlphaNumericString_Length1To11 string

// Format limitations: an..12
type AlphaNumericString_Length1To12 string

// Format limitations: an..126
type AlphaNumericString_Length1To126 string

// Format limitations: an..127
type AlphaNumericString_Length1To127 string

// Format limitations: an..13
type AlphaNumericString_Length1To13 string

// Format limitations: an..14
type AlphaNumericString_Length1To14 string

// Format limitations: an..15
type AlphaNumericString_Length1To15 string

// Format limitations: an..17
type AlphaNumericString_Length1To17 string

// Format limitations: an..18
type AlphaNumericString_Length1To18 string

// Format limitations: an..19
type AlphaNumericString_Length1To19 string

// Format limitations: an..199
type AlphaNumericString_Length1To199 string

// Format limitations: an..2
type AlphaNumericString_Length1To2 string

// Format limitations: an..20
type AlphaNumericString_Length1To20 string

// Format limitations: an..200
type AlphaNumericString_Length1To200 string

// Format limitations: an..25
type AlphaNumericString_Length1To25 string

// Format limitations: an..250
type AlphaNumericString_Length1To250 string

// Format limitations: an..256
type AlphaNumericString_Length1To256 string

// Format limitations: an..27
type AlphaNumericString_Length1To27 string

// Format limitations: an..28
type AlphaNumericString_Length1To28 string

// Format limitations: an..3
type AlphaNumericString_Length1To3 string

// Format limitations: an..30
type AlphaNumericString_Length1To30 string

// Format limitations: an..35
type AlphaNumericString_Length1To35 string

// Format limitations: an..4
type AlphaNumericString_Length1To4 string

// Format limitations: an..40
type AlphaNumericString_Length1To40 string

// Format limitations: an..49
type AlphaNumericString_Length1To49 string

// Format limitations: an..5
type AlphaNumericString_Length1To5 string

// Format limitations: an..50
type AlphaNumericString_Length1To50 string

// Format limitations: an..56
type AlphaNumericString_Length1To56 string

// Format limitations: an..57
type AlphaNumericString_Length1To57 string

// Format limitations: an..6
type AlphaNumericString_Length1To6 string

// Format limitations: an..60
type AlphaNumericString_Length1To60 string

// Format limitations: an..7
type AlphaNumericString_Length1To7 string

// Format limitations: an..70
type AlphaNumericString_Length1To70 string

// Format limitations: an..8
type AlphaNumericString_Length1To8 string

// Format limitations: an..9
type AlphaNumericString_Length1To9 string

// Format limitations: an..99999
type AlphaNumericString_Length1To99999 string

// Format limitations: an2
type AlphaNumericString_Length2To2 string

// Format limitations: an2..3
type AlphaNumericString_Length2To3 string

// Format limitations: an3
type AlphaNumericString_Length3To3 string

// Format limitations: an3..5
type AlphaNumericString_Length3To5 string

// Format limitations: an4
type AlphaNumericString_Length4To4 string

// Format limitations: an5..6
type AlphaNumericString_Length5To6 string

// Format limitations: an6
type AlphaNumericString_Length6To6 string

// Format limitations: an7
type AlphaNumericString_Length7To7 string

// Format limitations: an9
type AlphaNumericString_Length9To9 string

// Format limitations: a..1
type AlphaString_Length0To1 string

// Format limitations: a1
type AlphaString_Length1To1 string

// Format limitations: a..2
type AlphaString_Length1To2 string

// Format limitations: a..3
type AlphaString_Length1To3 string

// Format limitations: a..30
type AlphaString_Length1To30 string

// Format limitations: a..56
type AlphaString_Length1To56 string

// Format limitations: a..57
type AlphaString_Length1To57 string

// Format limitations: a2
type AlphaString_Length2To2 string

// Format limitations: a3
type AlphaString_Length3To3 string

// Format limitations: a3..5
type AlphaString_Length3To5 string

// Date format: DDMMYY
// pattern = "(0[1-9]|[1-2][0-9]|3[0-1])(0[1-9]|1[0-2])[0-9]{2}"
type Date_DDMMYY string

// Date  format: MMYY
// pattern = "(0[1-9]|1[0-2])([0-9][0-9])"
type Date_MMYY string

// Date format: YYYYMMDD
// pattern = "[0-9]{4}(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])"
type Date_YYYYMMDD string

// Value of the day in the  month. Only significant digits are mandatory. Example: 7
// pattern = "(0?[1-9]|[1-2][0-9]|3[0-1])"
type Day_nN string

// Format limitations: n..10
type DecimalLengthNTo10 float64

// Format limitations: n..15
type DecimalLengthNTo15 float64

// Format limitations: n..18
type DecimalLengthNTo18 float64

// Format limitations: n..3
type DecimalLengthNTo3 float64

// Format limitations: n..4
type DecimalLengthNTo4 float64

// Format limitations: n..6
type DecimalLengthNTo6 float64

// Format limitations: n..8
type DecimalLengthNTo8 float64

// Format limitations: n..9
type DecimalLengthNTo9 float64

// Value of the hours in the  time. Only significant digits are mandatory. Example: 7
// pattern = "[0-1]?[0-9]|2[0-3]"
type Hour_hH string

// Value of the month. Only significant digits are mandatory. Example: 7
// pattern = "(0?[1-9]|1[0-2])"
type Month_mM string

// Format limitations: n..10
type NumericDecimal_Length1To10 float64

// Format limitations: n..11
type NumericDecimal_Length1To11 float64

// Format limitations: n..12
type NumericDecimal_Length1To12 float64

// Format limitations: n..18
type NumericDecimal_Length1To18 float64

// Format limitations: n..35
type NumericDecimal_Length1To35 float64

// Format limitations: n..5
type NumericDecimal_Length1To5 float64

// Format limitations: n..8
type NumericDecimal_Length1To8 float64

// Format limitations: n..9
type NumericDecimal_Length1To9 float64

// Format limitations: n..3
// pattern = "-?[0-9]{0,3}"
type NumericInteger_Length0To3 int32

// Format limitations: n..35
// pattern = "-?[0-9]{0,35}"
type NumericInteger_Length0To35 int32

// Format limitations: n10
// pattern = "-?[0-9]{10,10}"
type NumericInteger_Length10To10 int32

// Format limitations: n1
// pattern = "-?[0-9]{1,1}"
type NumericInteger_Length1To1 int32

// Format limitations: n..10
// pattern = "-?[0-9]{1,10}"
type NumericInteger_Length1To10 int32

// Format limitations: n..11
// pattern = "-?[0-9]{1,11}"
type NumericInteger_Length1To11 int32

// Format limitations: n..12
// pattern = "-?[0-9]{1,12}"
type NumericInteger_Length1To12 int32

// Format limitations: n..13
// pattern = "-?[0-9]{1,13}"
type NumericInteger_Length1To13 int32

// Format limitations: n..15
// pattern = "-?[0-9]{1,15}"
type NumericInteger_Length1To15 int32

// Format limitations: n..18
// pattern = "-?[0-9]{1,18}"
type NumericInteger_Length1To18 int32

// Format limitations: n..2
// pattern = "-?[0-9]{1,2}"
type NumericInteger_Length1To2 int32

// Format limitations: n..3
// pattern = "-?[0-9]{1,3}"
type NumericInteger_Length1To3 int32

// Format limitations: n..4
// pattern = "-?[0-9]{1,4}"
type NumericInteger_Length1To4 int32

// Format limitations: n..5
// pattern = "-?[0-9]{1,5}"
type NumericInteger_Length1To5 int32

// Format limitations: n..6
// pattern = "-?[0-9]{1,6}"
type NumericInteger_Length1To6 int32

// Format limitations: n..8
// pattern = "-?[0-9]{1,8}"
type NumericInteger_Length1To8 int32

// Format limitations: n..9
// pattern = "-?[0-9]{1,9}"
type NumericInteger_Length1To9 int32

// Format limitations: n2
// pattern = "-?[0-9]{2,2}"
type NumericInteger_Length2To2 int32

// Format limitations: n3
// pattern = "-?[0-9]{3,3}"
type NumericInteger_Length3To3 int32

// Format limitations: n4
// pattern = "-?[0-9]{4,4}"
type NumericInteger_Length4To4 int32

// Format limitations: an10
type NumericString_Length10To10 string

// Format limitations: an3
type NumericString_Length3To3 string

// Format limitations: n1
type StringLength1To1 string

// Format limitations: an..12
type StringLength1To12 string

// Format limitations: an..18
type StringLength1To18 string

// Format limitations: an..2
type StringLength1To2 string

// Format limitations: an..20
type StringLength1To20 string

// Format limitations: an..3
type StringLength1To3 string

// Format limitations: an..35
type StringLength1To35 string

// Format limitations: an..4
type StringLength1To4 string

// Format limitations: an..5
type StringLength1To5 string

// Format limitations: an..6
type StringLength1To6 string

// Format limitations: an..7
type StringLength1To7 string

// Format limitations: an..9
type StringLength1To9 string

// Format limitations: an..9999
type StringLength1To9999 string

// Format limitations: an3..5
type StringLength3To5 string

// Format limitations: n6
type StringLength6To6 string

// Time format: 24H. All digits are mandatory . Example: from 0000 to 2359
// pattern = "([0-1][0-9]|2[0-3])[0-5][0-9]"
type Time24_HHMM string

// Time format: 24H.Only significant digits are mandatory. Example: from 0 to 2359
// pattern = "([0-1]?[0-9]|2[0-3])?[0-5]?[0-9]"
type Time24_hhmM string

// Value of the year. Example: 2003
// pattern = "[0-9]{4}"
type Year_YYYY string

//  **********************************
//	Formats for backward compatibility
//	**********************************

// Format limitations: an..35
type AlphaNumericStringLength1To35 string

// Format limitations: an..14
type AlphaNumericStringLength1To14 string

// Format limitations: an1
type AlphaNumericStringLength1To1 string

// Format limitations: an..30
type AlphaNumericStringLength1To30 string

// Format limitations: n..15
type NumericIntegerLength1To15 int32

// Format limitations: an..99999
type AlphaNumericStringLength1To99999 string

// Format limitations: an..3
type AlphaNumericStringLength1To3 string

// Format limitations: an..9
type AlphaNumericStringLength1To9 string

// Format limitations: an..17
type AlphaNumericStringLength1To17 string

// Format limitations: an..25
type AlphaNumericStringLength1To25 string

// Format limitations: an..5
type AlphaNumericStringLength1To5 string

// Format limitations: an..6
type AlphaNumericStringLength1To6 string

// Format limitations: an..70
type AlphaNumericStringLength1To70 string

// Format limitations: a..6
type AlphaStringLength1To6 string

// Format limitations: an..10
type AlphaNumericStringLength1To10 string

// Format limitations: an..4
type AlphaNumericStringLength1To4 string

//  **************************
//	END backward compatibility
//	**************************

// Value of the minutes in the  time. Only significant digits are mandatory. Example: 7
type Minute_mM string

// Format limitations: n..2
type NumericDecimal_Length1To2 float64

// Format limitations: an..547
type AlphaNumericString_Length1To547 string

// Format limitations: a..6
type AlphaString_Length1To6 string

// Format limitations: an..100
type AlphaNumericString_Length1To100 string

// Format limitations: an..500
type AlphaNumericString_Length1To500 string

// Format limitations: an..999
type AlphaNumericString_Length1To999 string

// Format limitations: an..255
type AlphaNumericString_Length1To255 string

// Format limitations: an..320
type AlphaNumericString_Length1To320 string

// Format limitations: an..64
type AlphaNumericString_Length1To64 string

// Format limitations: an..90
type AlphaNumericString_Length1To90 string

// Format limitations: an..2
type AlphaNumericString_Length0To2 string

// Format limitations: an..55
type AlphaNumericString_Length1To55 string

// Format limitations: n2..4
type NumericInteger_Length2To4 int32

// Format limitations: an..63
type AlphaNumericString_Length1To63 string

// Format limitations: an..32
type AlphaNumericString_Length1To32 string

// Format limitations: an..16
type AlphaNumericString_Length1To16 string

// Format limitations: an8
type AlphaNumericString_Length8To8 string

// Format limitations: n8
type NumericInteger_Length8To8 int32

// Format limitations: an..24
type AlphaNumericString_Length1To24 string

// Format limitations: n6
type NumericInteger_Length6To6 int32

// Format limitations: n..4
type NumericInteger_Length0To4 int32

// Format limitations: n..3
type NumericDecimal_Length1To3 float64

// Format limitations: a..60
type AlphaString_Length1To60 string

// Format limitations: an..22
type AlphaNumericString_Length1To22 string

// Format limitations: n5..8
type NumericInteger_Length5To8 int32

// Format limitations: an..31
type AlphaNumericString_Length1To31 string

// Format limitations: an..120
type AlphaNumericString_Length1To120 string

// Format limitations: n..35
type NumericInteger_Length1To35 int32

// Format limitations: an..188
type AlphaNumericString_Length1To188 string

// Format limitations: n..30
type NumericDecimal_Length1To30 float64

// Format limitations: n..20
type NumericDecimal_Length1To20 float64

// Format limitations: n..20
type NumericInteger_Length1To20 int32

// Format limitations: an..61
type AlphaNumericString_Length1To61 string

// Format limitations: an..512
type AlphaNumericString_Length1To512 string

// Format limitations: an..99
type AlphaNumericString_Length1To99 string

// Format limitations: an..108
type AlphaNumericString_Length0To108 string

// Format limitations: an..56
type AlphaNumericString_Length0To56 string

// Format limitations: an..144
type AlphaNumericString_Length0To144 string

// Format limitations: an..400
type AlphaNumericString_Length1To400 string

// Format limitations: an6..9
type AlphaNumericString_Length6To9 string

// Format limitations: an..70
type AlphaNumericString_Length0To70 string

// Format limitations: an..6
type AlphaNumericString_Length0To6 string

// Format limitations: a..3
type AlphaString_Length0To3 string

// Format limitations: an..12
type AlphaNumericString_Length0To12 string

// Format limitations: a..9
type AlphaString_Length0To9 string

// Format limitations: an..5
type AlphaNumericString_Length0To5 string

// Format limitations: n..6
type NumericInteger_Length0To6 int32
