package main

import (
	"fmt"
	"strings"
)

//WordSplit ...
func WordSplit(strArr []string) string {

	secondWordList := strings.Split(strArr[1], ",")
	totalWords := ""
	fistWord := ""
	count := 0
	for i := 0; i < len(strArr[0]); i++ {
		fistWord += strArr[0][i : i+1]
		secondWord := strArr[0][i+1 : len(strArr[0])]
		for _, item := range secondWordList {
			if fistWord == item {
				count++
			} else if secondWord == item {
				count++
			}
		}

		if count == 2 {
			totalWords = fistWord + "," + secondWord
			break
		}
		count = 0
	}
	if totalWords == "" {
		return "non"
	}
	return totalWords

}

func main() {
	input := []string{"baseball", "a,all,b,ball,bas,bae,cat,code,d,e,quit,z"}
	fmt.Println(WordSplit(input))

}
