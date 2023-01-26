package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
}

type data struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	var stack []data

	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1].height > h {
			index := stack[len(stack)-1].index
			height := stack[len(stack)-1].height
			maxArea = max(maxArea, height*(i-index))
			stack = stack[:len(stack)-1]
			start = index
		}
		stack = append(stack, data{start, h})
	}
	for _, h := range stack {
		maxArea = max(maxArea, (h.height * (len(heights) - h.index)))
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
