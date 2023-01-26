package main

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	ans := []string{}

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if NumOfSetBits(h)+NumOfSetBits(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

func NumOfSetBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
