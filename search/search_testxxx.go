package search

import "testing"

func TestLinearSearch1(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want int
	// }{
	// 	{"test", args{[]int{1, 2, 3, 4, 5}, 0}, -1},
	// 	{"test", args{[]int{1, 2, 3, 4, 5}, 1}, 0},
	// 	{"test", args{[]int{1, 2, 3, 4, 5}, 2}, 1},
	// 	{"test", args{[]int{1, 2, 3, 4, 5}, 3}, 2},
	// 	{"test", args{[]int{1, 2, 3, 4, 5}, 4}, 3},
	// 	{"test", args{[]int{1, 2, 3, 4, 5}, 5}, 4},
	// 	{"test", args{[]int{1, 2, 3, 4, 5}, 6}, -1},
	// }

	type testStruct struct {
		name string
		args args
		want int
	}
	var test testStruct

	test = testStruct{"test", args{[]int{1, 2, 3, 4, 5}, 6}, -1}

	t.Run(test.name, func(t *testing.T) {
		if got := LinearSearch(test.args.arr, test.args.n); got != test.want {
			t.Errorf("LinearSearch(%v, %v) = %v, want %v", test.args.arr, test.args.n, got, test.want)
		}
	})
}
