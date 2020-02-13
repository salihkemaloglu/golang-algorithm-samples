package main

import "fmt"

//FirstFactorial ...
func FirstFactorial(num int) int {

	// code goes here
	var total = 1
	for i := 1; i <= num; i++ {
		total = total * i
	}
	return total

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FirstFactorial(8))

}
