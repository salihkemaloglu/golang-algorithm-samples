package main

import (
	"fmt"
	"regexp"
	"strings"
)

//LongestWord ...
func LongestWord(sen string) string {

	// code goes here
	// Note: feel free to modify the return type of this function
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

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LongestWord("a confusing /:sentence:/[ this is not!!!!!!!~"))

}
