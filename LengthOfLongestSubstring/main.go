package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	result := ""
	for i := 0; i < len(s); i++ {
		temp := ""
		for j := i; j < len(s); j++ {
			if !contains(temp, s[j:j+1]) {
				temp += s[j : j+1]
			} else {
				break
			}
		}
		if len(temp) > len(result) {
			result = temp
		}
	}
	return len(result)

}
func contains(s string, e string) bool {
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == e {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
