package main

import "fmt"

// Search returns the index of target in a sorted slice, or -1. The loop
// invariant: if target exists, it lies within [lo, hi].
func Search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2 // overflow-safe midpoint
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(Search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(Search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
