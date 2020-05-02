package automotive

import (
	"testing"
)

func TestGenerateLicensePlate(t *testing.T) {
	if len(GenerateLicensePlate(230948320948)) != 7 {
		t.Error("Expected license plate to equal seven digits.")
	}
}
