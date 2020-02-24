package main

import (
	"fmt"
	"sort"
	"strings"
)

//AlphabetSoup ...
func AlphabetSoup(str string) string {

	strs := []string{}
	for i := 0; i < len(str); i++ {
		strs = append(strs, str[i:i+1])
	}
	sort.Strings(strs)
	return strings.Join(strs, "")

}

func main() {

	fmt.Println(AlphabetSoup("hooplah"))

}
