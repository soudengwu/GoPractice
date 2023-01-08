package romanNumeral

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Description string
	Number      uint16
	Roman       string
}{
	{Number: 1, Roman: "I"},
	{Number: 2, Roman: "II"},
	{Number: 3, Roman: "III"},
	{Number: 4, Roman: "IV"},
	{Number: 5, Roman: "V"},
	{Number: 6, Roman: "VI"},
	{Number: 7, Roman: "VII"},
	{Number: 8, Roman: "VIII"},
	{Number: 9, Roman: "IX"},
	{Number: 10, Roman: "X"},
	{Number: 14, Roman: "XIV"},
	{Number: 18, Roman: "XVIII"},
	{Number: 20, Roman: "XX"},
	{Number: 39, Roman: "XXXIX"},
	{Number: 40, Roman: "XL"},
	{Number: 47, Roman: "XLVII"},
	{Number: 49, Roman: "XLIX"},
	{Number: 50, Roman: "L"},
	{Number: 100, Roman: "C"},
	{Number: 90, Roman: "XC"},
	{Number: 400, Roman: "CD"},
	{Number: 500, Roman: "D"},
	{Number: 900, Roman: "CM"},
	{Number: 1000, Roman: "M"},
	{Number: 1984, Roman: "MCMLXXXIV"},
	{Number: 3999, Roman: "MMMCMXCIX"},
	{Number: 2014, Roman: "MMXIV"},
	{Number: 1006, Roman: "MVI"},
	{Number: 798, Roman: "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Number, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Number)
			want := test.Roman

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestConvertRomanToNum(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Number), func(t *testing.T) {
			got := ConvertRomanToNum(test.Roman)
			want := test.Number

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			log.Println(arabic)
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertRomanToNum(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
