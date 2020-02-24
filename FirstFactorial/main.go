package main

import "fmt"

//FirstFactorial ...
func FirstFactorial(num int) int {
	var total = 1
	for i := 1; i <= num; i++ {
		total = total * i
	}
	return total

}

func main() {
	fmt.Println(FirstFactorial(8))
}
