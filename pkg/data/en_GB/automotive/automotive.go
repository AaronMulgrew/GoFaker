package automotive

import (
	"fmt"
	"math/rand"
)

// RandomDigits is a function to generate the
// semi-random character string for a valid registration
// numberplate
func RandomDigits(length int, digits bool) string {
	partial := ""
	if digits {
		// these are the only years in which cars can have valid registration plates.
		var registrationYears [6]int
		registrationYears[0] = 0
		registrationYears[1] = 1
		registrationYears[2] = 2
		registrationYears[3] = 5
		registrationYears[4] = 6
		registrationYears[5] = 7
		min := 0
		max := 5
		randIndex := rand.Intn(max-min) + min
		num1 := registrationYears[randIndex]
		num2 := rand.Intn(9-0) + min
		partial = fmt.Sprintf("%d%d", num1, num2)
	} else {
		base := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		buf := make([]byte, length)
		buf[0] = base[rand.Intn(len(base))]
		for i := 0; i < length; i++ {
			buf[i] = base[rand.Intn(len(base))]
		}
		partial = string(buf)
	}

	//characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//allCharactersAndDigits := digits + characters

	return partial
}

// GenerateLicensePlate puts the three components together -
// two letters, two year digits and three letters
func GenerateLicensePlate(seed int64) string {
	// License Plates come in 3 parts
	// 2 letters, 2 digits (year), 3 letters
	part1 := RandomDigits(2, false)
	part2 := RandomDigits(2, true)
	part3 := RandomDigits(3, false)
	return part1 + part2 + part3
}

func randomSpeed(seed int64) string {
	base := "012345678"
	length := 2
	buf := make([]byte, length)
	buf[0] = base[rand.Intn(len(base))]
	for i := 0; i < length; i++ {
		buf[i] = base[rand.Intn(len(base))]
	}
	mph := string(buf) + " mph"

	return mph
}

// GenerateRandomSpeed returns a random miles per hour speed
func GenerateRandomSpeed(seed int64) string {
	speed := randomSpeed(seed)
	return speed
}
