package search

import "testing"

func TestLinearSearch(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{[]int{1, 2, 3, 4, 5}, 0}, -1},
		{"test", args{[]int{1, 2, 3, 4, 5}, 1}, 0},
		{"test", args{[]int{1, 2, 3, 4, 5}, 2}, 1},
		{"test", args{[]int{1, 2, 3, 4, 5}, 3}, 2},
		{"test", args{[]int{1, 2, 3, 4, 5}, 4}, 3},
		{"test", args{[]int{1, 2, 3, 4, 5}, 5}, 4},
		{"test", args{[]int{1, 2, 3, 4, 5}, 6}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearch(tt.args.arr, tt.args.n); got != tt.want {
				t.Errorf("LinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
