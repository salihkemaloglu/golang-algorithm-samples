package main

import (
	"fmt"
	"io"
)

func solvePizzaCrust(r float64, c float64) float64 {

	totalArea := r * r
	cheeseArea := (r - c) * (r - c)
	return (cheeseArea / totalArea) * 100
}

func main() {

	var r, c float64

	for {
		_, err := fmt.Scanf("%v%v", &r, &c)
		if err == io.EOF {
			break
		}
		fmt.Println(solvePizzaCrust(r, c))
	}

}
