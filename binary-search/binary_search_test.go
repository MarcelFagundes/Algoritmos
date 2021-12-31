package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	type test struct {
		values []int
		target int
		want   int
	}
	tests := []test{
		{[]int{1, 2, 3, 4}, 3, 2},
		{[]int{1, 2, 3, 4}, 9, -1},
	}
	for _, test := range tests {
		got := BinarySearch(test.values, test.target)
		if got != test.want {
			t.Fatalf("got %v, want %v", got, test.want)
		}
	}
}
