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

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(AlphabetSoup("hooplah"))

}
