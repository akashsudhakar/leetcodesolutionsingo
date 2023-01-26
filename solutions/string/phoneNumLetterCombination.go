package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	digitMap := make(map[rune][]string)
	digitMap['0'] = []string{}
	digitMap['1'] = []string{}
	digitMap['2'] = []string{"a", "b", "c"}
	digitMap['3'] = []string{"d", "e", "f"}
	digitMap['4'] = []string{"g", "h", "i"}
	digitMap['5'] = []string{"j", "k", "l"}
	digitMap['6'] = []string{"m", "n", "o"}
	digitMap['7'] = []string{"p", "q", "r", "s"}
	digitMap['8'] = []string{"t", "u", "v"}
	digitMap['9'] = []string{"w", "x", "y", "z"}

	outList := []string{}
	findCombinations(&outList, []rune(digits), "", 0, digitMap)
	return outList
}

func findCombinations(outList *[]string, digitArr []rune, previous string, index int, digitMap map[rune][]string) {
	if index == len(digitArr) {
		*outList = append(*outList, previous)
		return
	}
	letters := digitArr[index]
	for _, character := range digitMap[letters] {
		word := strings.Join([]string{previous, character}, "")
		findCombinations(outList, digitArr, word, index+1, digitMap)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}
