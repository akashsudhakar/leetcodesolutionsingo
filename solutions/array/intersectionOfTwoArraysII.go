package array

func intersect(nums1 []int, nums2 []int) []int {
	countMap := map[int]int{}
	for _, num := range nums1 {
		if countMap[num] == 0 {
			countMap[num] = 1
		} else {
			countMap[num] = countMap[num] + 1
		}
	}
	var out []int
	for _, num := range nums2 {
		if countMap[num] > 0 {
			out = append(out, num)
			countMap[num] = countMap[num] - 1
		}
	}
	return out
}
