package randomdata

import (
	"math/rand"
	"time"
)

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
