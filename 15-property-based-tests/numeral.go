package main

import (
	"errors"
	"fmt"
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

const maxConvertableArabicNumber uint16 = 3999

var ErrMaxConvertableArabicNumberExceeded = errors.New(fmt.Sprintf("can't convert a number greater than %d", maxConvertableArabicNumber))

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

func ConvertToRoman(arabic uint16) (string, error) {
	if arabic > maxConvertableArabicNumber {
		return "", ErrMaxConvertableArabicNumberExceeded
	}
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
