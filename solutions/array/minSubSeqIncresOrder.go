package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	leftIdx := 0
	rightIdx := len(nums)
	sumLft := 0
	sumRgt := 0

	for leftIdx < rightIdx {
		if sumLft+nums[leftIdx] < sumRgt {
			sumLft = sumLft + nums[leftIdx]
			leftIdx = leftIdx + 1
		} else {
			rightIdx = rightIdx - 1
			sumRgt = sumRgt + nums[rightIdx]
		}
	}
	var out []int
	for i := len(nums) - 1; i >= rightIdx; i-- {
		out = append(out, nums[i])
	}
	return out
}

func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8})) // 3 4 8 9 10
	fmt.Println(minSubsequence([]int{4, 4, 7, 6, 7}))  // 4 4 6 7 7
}
