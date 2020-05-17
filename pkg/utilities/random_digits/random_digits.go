package random_digits

import "math/rand"

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
