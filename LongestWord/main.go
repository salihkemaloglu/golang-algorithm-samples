package main

import (
	"fmt"
	"regexp"
	"strings"
)

//LongestWord ...
func LongestWord(sen string) string {
	words := strings.Fields(sen)
	longestWord := ""
	for _, element := range words {
		word := removeSpecialChar(element)
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	return longestWord
}

func removeSpecialChar(input string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(input, "")
}

func main() {

	fmt.Println(LongestWord("a confusing /:sentence:/[ this is not!!!!!!!~"))

}
