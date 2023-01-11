package array

func findDifference(nums1 []int, nums2 []int) [][]int {
	existMapFir := map[int]bool{}

	for _, num := range nums1 {
		existMapFir[num] = true
	}
	var onlyInSec []int
	existMapSec := map[int]bool{}
	for _, num := range nums2 {
		existMapSec[num] = true
		if !existMapFir[num] {
			onlyInSec = append(onlyInSec, num)
			existMapFir[num] = true
		}
	}
	var onlyInFir []int
	for _, num := range nums1 {
		if !existMapSec[num] {
			onlyInFir = append(onlyInFir, num)
			existMapSec[num] = true
		}
	}
	out := make([][]int, 2)
	out[0] = onlyInFir
	out[1] = onlyInSec
	return out
}
