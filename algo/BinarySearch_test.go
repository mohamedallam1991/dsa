package algo

import (
	"testing"
)

var arr = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arg    []int
		needle int
		want   bool
	}{
		{arr, 8, false},
		{arr, 13, true},
		{arr, 42, false},
		{arr, 83, true},
	}
	for _, tt := range tests {
		t.Run("array test", func(t *testing.T) {
			if tt.want != BinarySearch(tt.arg, tt.needle) {
				if LinearSearch(tt.arg, tt.needle) == false {
					t.Error("Using LinearSearch Your binary search isnt working properly")
				}
				t.Errorf("Your binary search isnt working properly")
			}
		})
	}
}

func TestBinarySearchTest(t *testing.T) {
	tt := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	for _, v := range tt {
		expect := BinarySearch(tt, v)
		actual := LinearSearch(tt, v)
		if expect != actual {
			t.Error("Your binary search isnt working properly")
		}
	}

}

func LinearSearch(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
