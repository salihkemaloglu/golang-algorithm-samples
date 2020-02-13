package main

import (
	"fmt"
	"strconv"
)

// QuestionsMarks ..
func QuestionsMarks(str string) string {

	for i := 0; i < len(str); i++ {
		if isNumeric(str[i : i+1]) {
			count := 1
			for j := i + 1; j < len(str); j++ {
				if str[j:j+1] != "?" && count < 3 && isNumeric(str[j:j+1]) {
					return "false"
				}
				count++
			}
		}
	}
	return "true"

}
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(QuestionsMarks("aa6?9"))

}
