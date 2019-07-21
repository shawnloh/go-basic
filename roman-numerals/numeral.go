package roman_numerals

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

type romanNumerals []RomanNumeral

func (r romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if i+1 < len(roman) && symbol == 'I' {
			//nextSymbol := roman[i+1]
			//
			//// build the two character string
			//potentialNumber := string([]byte{symbol, nextSymbol})
			//
			//// get the value of the two character string
			////value := romanNumerals.ValueOf(potentialNumber)
			//if value != 0 {
			//	total += value
			//} else {
			//	total ++
			//}
		} else {
			total++
		}
	}
	return total
}
