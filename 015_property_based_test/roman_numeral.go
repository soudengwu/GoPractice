package romanNumeral

import (
	"strings"
)

// No roman numeral must appear more than 3 times,
// if that happens, take the next highest and subtract by the next highest symbol
// only [I, X and C] can be used as subtractor

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral
type windowedRoman string

var allRomanNumerals = RomanNumerals{
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

func ConvertToRoman(num uint16) string {
	var result strings.Builder

	// for i := 0; i < num; i++ {
	// 	result.WriteString("I")
	// }

	// for num > 0 {
	// 	switch {
	// 	case num > 9:
	// 		result.WriteString("X")
	// 		num -= 10
	// 	case num > 8:
	// 		result.WriteString("IX")
	// 		num -= 9
	// 	case num > 4:
	// 		result.WriteString("V")
	// 		num -= 5
	// 	case num > 3:
	// 		result.WriteString("IV")
	// 		num -= 4
	// 	default:
	// 		result.WriteString("I")
	// 		num--
	// 	}
	// }

	for _, numeral := range allRomanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func ConvertRomanToNum(roman string) (total uint16) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
