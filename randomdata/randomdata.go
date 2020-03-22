package randomdata

import (
	"math/rand"
	"time"
)

const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"

// RandomIntArray returns an array of random integer values
func RandomIntArray(arrSize int, intRange int, positiveVal bool) []int {
	rand.Seed(time.Now().UnixNano())
	intArr := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		intArr[i] = rand.Intn(intRange)
		if !positiveVal {
			if rand.Intn(2) == 1 {
				intArr[i] = -intArr[i]
			}
		}
	}
	return intArr
}

// RandomString returns a random string of specified length and character set
func RandomString(maxlen int, charset string) string {
	var s string
	for i := 0; i < rand.Intn(maxlen)+1; i++ {
		s = s + string(charset[rand.Intn(len(charset))])
	}
	return s
}
