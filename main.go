package main

import (
	"fmt"
	"goalgorithms/randomdata"
	"goalgorithms/search"
)

func main() {
	const alphaLower = "abcdefghijklmnopqrstuvwxyz"

	fmt.Println("hello algorithms")
	sl := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(search.LinearSearch(sl, 3))
	fmt.Println(randomdata.RandomIntArray(10, 10, false))
	fmt.Println("123456789")
	// fmt.Println(randomdata.RandomString(10, alphaLower))
	for i := 0; i < 1000000; i++ {
		s2 := randomdata.RandomString(10, alphaLower)
		// fmt.Println(s2, len(s2))
		if len(s2) == 1 {
			fmt.Println("==1")
		}
	}
}
