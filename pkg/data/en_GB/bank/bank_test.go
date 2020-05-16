package bank

import (
	"fmt"
	"strconv"
	"testing"
	"unicode"
)

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isNumbers(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func isAlphaNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			if !unicode.IsNumber(r) {
				return false
			}
		}
	}
	return true
}

func TestGenerateIBAN(t *testing.T) {
	IBAN, AccountNumber, BankCode := GenerateIBAN(230948320948)
	fmt.Println(IBAN)
	fmt.Println(BankCode)
	fmt.Println(AccountNumber)

	if len(IBAN) != 22 {
		t.Error("Expected IBAN to equal exactly 22 numbers.")
	}
	if isAlphaNumeric(IBAN) != true {
		t.Error("Expected IBAN to be alphanumeric")
	}
	if isNumbers(AccountNumber) != true {
		t.Error("Expected Account Number to be digits")
	}

	if isNumbers(BankCode) != true {
		t.Error("Expected Bank Code to be digits")
	}
}

func TestCalculateCheckDigit(t *testing.T) {
	BankCode := 0207214234
	AccountNumber := 70817960
	BankCodeBuf := []byte(strconv.Itoa(BankCode))
	AccountNumberBuf := []byte(strconv.Itoa(AccountNumber))

	res := calculateCheckDigit(BankCodeBuf, AccountNumberBuf)
	if res != 75 {
		t.Error("Expected res to equal 75")
	}
	fmt.Println(res)
}
