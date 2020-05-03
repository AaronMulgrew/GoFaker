package automotive

import (
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

func TestGenerateLicensePlate(t *testing.T) {
	if len(GenerateLicensePlate(230948320948)) != 7 {
		t.Error("Expected license plate to equal seven alphanumeric characters.")
	}
	if IsLetter(GenerateLicensePlate(230948320948)) {
		t.Error("Expected license plate to be alphanumeric.")
	}
	if isAlphaNumeric(GenerateLicensePlate(230948320948)) != true {
		t.Error("Expected license plate to be alphanumeric.")
	}
}

func TestGenerateRandomSpeed(t *testing.T) {
	if len(GenerateRandomSpeed(230948320948)) != 6 {
		t.Error("Expected speed to be six characters, current speed (2 digits) followed by a space and miles per hour.")
	}
}
