package main

func commonChars(words []string) []string {
	countArray := make([]int, 26)
	for _, char := range words[0] {
		countArray[char-97] = countArray[char-97] + 1
	}
	for _, word := range words {
		iterArray := make([]int, 26)
		for _, char := range word {
			iterArray[char-97] = iterArray[char-97] + 1
		}
		for i := 0; i < 26; i++ {
			countArray[i] = min(iterArray[i], countArray[i])
		}
	}
	var outArr []string
	for index, num := range countArray {
		if num > 0 {
			for i := 0; i < num; i++ {
				outArr = append(outArr, string(index+97))
			}
		}
	}
	return outArr
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	commonChars([]string{"bella", "label", "roller"})
}
