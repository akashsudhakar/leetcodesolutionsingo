package string

import "strings"

func lengthOfLastWord(s string) int {
	strArry := strings.Split(strings.Trim(s, " "), " ")
	length := len(strArry)
	lastWord := strArry[length-1]
	return len(lastWord)
}
