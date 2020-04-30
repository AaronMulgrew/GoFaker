package bank

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
)

func calculateCheckDigit(BankCode []byte, AccountNumber []byte) int64 {
	fmt.Println(string(BankCode))
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
	t2 = append(t2, 15)
	t2 = append(t2, 11)
	t2 = append(t2, 0)
	t2 = append(t2, 0)
	fmt.Println(t2)

	var buf bytes.Buffer
	for i := range t2 {
		buf.WriteString(fmt.Sprintf("%d", t2[i]))
	}

	n := new(big.Int)
	allNums, _ := n.SetString(buf.String(), 10)
	//allNums, _ := n.SetString("100100100987654321131400", 10)

	fmt.Println(allNums)

	sumAllNumbers := 0
	for _, i := range t2 {
		sumAllNumbers = sumAllNumbers + i
	}
	//strSumAllNumbers := strconv.Itoa(sumAllNumbers)

	fmt.Println(allNums.Mod(allNums, big.NewInt(97)))
	checkdigit := 98 - (allNums.Mod(allNums, big.NewInt(97)).Int64())
	return checkdigit
}

func RandomDigits(length int) []byte {
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

func GenerateIBAN(seed int64) string {
	rand.Seed(seed)
	bankCode := RandomDigits(8)
	accountNumber := RandomDigits(10)

	fmt.Println(bankCode)
	checkdigits := calculateCheckDigit(bankCode, accountNumber)

	fmt.Println(checkdigits)

	str := "GB" + strconv.FormatInt(checkdigits, 10) + string(bankCode) + string(accountNumber)

	fmt.Println(str)
	return str
}
