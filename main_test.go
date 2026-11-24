package main

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	for i, v := range nums {
		if got := Search(nums, v); got != i {
			t.Errorf("Search(%d) = %d, want %d", v, got, i)
		}
	}
	for _, v := range []int{2, -5, 13} {
		if got := Search(nums, v); got != -1 {
			t.Errorf("Search(%d) = %d, want -1", v, got)
		}
	}
	if got := Search([]int{}, 1); got != -1 {
		t.Errorf("empty = %d", got)
	}
	if got := Search([]int{5}, 5); got != 0 {
		t.Errorf("single hit = %d", got)
	}
	if got := Search([]int{5}, 4); got != -1 {
		t.Errorf("single miss = %d", got)
	}
}
