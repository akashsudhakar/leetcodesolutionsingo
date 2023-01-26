package main

import (
	"sort"
)

type data struct {
	word string
	mod  int
}

func RemainderSorting(strArr []string) []string {
	var dataArr []data
	for _, value := range strArr {
		modulus := len(value) % 3
		dataArr = append(dataArr, data{value, modulus})
	}
	sort.Slice(dataArr, func(i, j int) bool {
		if dataArr[i].mod != dataArr[j].mod {
			return dataArr[i].mod < dataArr[j].mod
		}
		return dataArr[i].word < dataArr[j].word
	})
	outArr := make([]string, len(dataArr))
	for index, value := range dataArr {
		outArr[index] = value.word
	}
	return outArr
}
