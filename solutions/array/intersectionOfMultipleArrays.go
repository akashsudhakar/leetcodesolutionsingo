package array

import "sort"

func intersection(nums [][]int) []int {
	countMap := map[int]int{}

	for _, array := range nums {
		for _, num := range array {
			if countMap[num] == 0 {
				countMap[num] = 1
			} else {
				countMap[num] = countMap[num] + 1
			}
		}
	}

	var out []int
	for key, value := range countMap {
		if value == len(nums) {
			out = append(out, key)
		}
	}
	sort.Ints(out)
	return out
}
