package math

import "math"

func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	maxNum := nums[0]
	secMax := math.MinInt64
	thiMax := math.MinInt64
	for _, num := range nums {
		if num >= maxNum {
			if num != maxNum {
				thiMax = secMax
				secMax = maxNum
				maxNum = num
			}
		} else if num >= secMax {
			if num != secMax {
				thiMax = secMax
				secMax = num
			}
		} else if num > thiMax {
			thiMax = num
		}
	}
	if thiMax != math.MinInt64 {
		return thiMax
	}
	return maxNum
}
