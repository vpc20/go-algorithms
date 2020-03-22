package search

// LinearSearch function
func LinearSearch(arr []int, n int) int {
	for i, v := range arr {
		if v == n {
			return i
		}
	}
	return -1
}
