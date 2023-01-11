package array

func intersection(nums1 []int, nums2 []int) []int {
	countMap := map[int]int{}
	for _, num := range nums1 {
		if countMap[num] == 0 {
			countMap[num] = 1
		}
	}
	var out []int
	existMap := make(map[int]bool)
	for _, num := range nums2 {
		if countMap[num] == 1 && !existMap[num] {
			existMap[num] = true
			out = append(out, num)
		}
	}
	return out
}
