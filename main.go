package main

import (
	"fmt"
	"goalgorithms/randomdata"
	"goalgorithms/search"
)

func main() {
	fmt.Println("hello algorithms")
	sl := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(search.LinearSearch(sl, 3))
	fmt.Println(randomdata.RandomIntArray(10, 10, false))
}
