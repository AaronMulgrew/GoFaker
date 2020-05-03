package bank

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
)

func calculateCheckDigit(BankCode []byte, AccountNumber []byte) int64 {

	// convert the byte array to string
	var t2 = []int{}
	for _, i := range BankCode {
		j, err := strconv.Atoi(string(i))
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}

	for _, i := range AccountNumber {
		j, err := strconv.Atoi(string(i))
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}

	// make sure we add the letters GB converted to IBAN numbers
	t2 = append(t2, 16)
	t2 = append(t2, 11)
	t2 = append(t2, 0)
	t2 = append(t2, 0)

	var buf bytes.Buffer
	for i := range t2 {
		buf.WriteString(fmt.Sprintf("%d", t2[i]))
	}

	n := new(big.Int)
	allNums, _ := n.SetString(buf.String(), 10)
	// MOD by 97 and do 98 - result... in line with IBAN standards
	checkdigit := 98 - (allNums.Mod(allNums, big.NewInt(97)).Int64())
	return checkdigit
}

func randomDigits(length int) []byte {
	digits := "0123456789"
	//characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//allCharactersAndDigits := digits + characters

	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	//buf[1] = characters[rand.Intn(len(characters))]

	for i := 0; i < length; i++ {
		buf[i] = digits[rand.Intn(len(digits))]
	}
	return buf
}

// GenerateIBAN is the main entry point for the banking generation
// and returns the IBAN, Bank code and account number.
func GenerateIBAN(seed int64) (string, string, string) {
	rand.Seed(seed)
	bankCode := randomDigits(8)
	accountNumber := randomDigits(10)

	checkdigits := calculateCheckDigit(bankCode, accountNumber)

	IBAN := "GB" + strconv.FormatInt(checkdigits, 10) + string(bankCode) + string(accountNumber)

	return IBAN, string(bankCode), string(accountNumber)
}
