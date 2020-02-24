package main

import "fmt"

//LetterChanges ...
func LetterChanges(str string) string {

	letterChanges := ""
	for _, char := range str {
		if char >= 97 && char <= 122 {
			if char == 122 {
				letterChanges += string(97)
			} else {
				char++
				if char == 97 {
					letterChanges += string(65)
				} else if char == 101 {
					letterChanges += string(69)
				} else if char == 105 {
					letterChanges += string(73)
				} else if char == 111 {
					letterChanges += string(79)
				} else if char == 117 {
					letterChanges += string(85)
				} else {
					letterChanges += string(char)
				}
			}
		} else {
			letterChanges += string(char)
		}

	}
	return letterChanges

}

func main() {

	fmt.Println(LetterChanges("fun times!"))

}
