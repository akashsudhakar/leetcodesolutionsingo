package main

import (
	"fmt"
)

func main() {
	var x int
	arr := [3]int{3, 5, 2}
	x, arr = arr[0], arr[0:]
	fmt.Println(x, arr)
}
