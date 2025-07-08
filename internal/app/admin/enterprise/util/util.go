package util

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidateCNPJ(cnpj string) error {
	cnpj = strings.TrimSpace(cnpj)
	re := regexp.MustCompile(`[^0-9]`)
	cnpj = re.ReplaceAllString(cnpj, "")

	if len(cnpj) != 14 {
		return fmt.Errorf("CNPJ has an invalid length")
	}

	allDigitsEqual := true
	for i := 1; i < 14 && allDigitsEqual; i++ {
		if cnpj[i] != cnpj[0] {
			allDigitsEqual = false
		}
	}
	if allDigitsEqual {
		return fmt.Errorf("CNPJ has all identical digits")
	}

	pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i := 0; i < 12; i++ {
		sum += int(cnpj[i]-'0') * pesos1[i]
	}
	remainder := sum % 11
	dv1 := 0
	if remainder >= 2 {
		dv1 = 11 - remainder
	}

	sum = 0
	for i := 0; i < 13; i++ {
		sum += int(cnpj[i]-'0') * pesos2[i]
	}
	remainder = sum % 11
	dv2 := 0
	if remainder >= 2 {
		dv2 = 11 - remainder
	}

	if int(cnpj[12]-'0') != dv1 || int(cnpj[13]-'0') != dv2 {
		return fmt.Errorf("Invalid CNPJ: check digits are incorrect")
	}

	return nil
}

// RemoveNonDigits removes all non-numeric charecters from the input string.
func RemoveNonDigits(s string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(s, "")
}
