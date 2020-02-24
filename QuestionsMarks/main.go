package main

import (
	"fmt"
	"strconv"
)

// QuestionsMarks ..
func QuestionsMarks(str string) string {

	questionsMarks := 0
	for i := 0; i < len(str); i++ {
		if isNumeric(str[i : i+1]) {

			for j := i + 1; j < i+3; j++ {
				if str[j:j+1] == "?" {
					questionsMarks++
				} else {
					return ""
				}
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

	fmt.Println(QuestionsMarks("9???1???9???1???9"))

}
