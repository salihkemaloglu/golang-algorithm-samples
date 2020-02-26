package main

import (
	"fmt"
	"strconv"
	"strings"
)

//MissingDigit ...
func MissingDigit(str string) string {

	items := strings.Split(str, " ")
	total := 0
	if isNumeric(items[4]) {
		switch items[1] {
		case "+":
			if isNumeric(items[0]) {
				total = convert(items[4]) - convert(items[0])
				totalStr := strconv.Itoa(total)
				for i := 0; i < len(totalStr); i++ {
					if totalStr[i:i+1] != items[2][i:i+1] {
						return totalStr[i : i+1]
					}
				}
			} else {
				total = convert(items[4]) - convert(items[2])
				totalStr := strconv.Itoa(total)
				for i := 0; i < len(totalStr); i++ {
					if totalStr[i:i+1] != items[0][i:i+1] {
						return totalStr[i : i+1]
					}
				}
			}
		case "-":
			if isNumeric(items[0]) {
				total = convert(items[4]) + (convert(items[0]) * -1)
				if total < 0 {
					total = total * -1
				}
				totalStr := strconv.Itoa(total)
				for i := 0; i < len(totalStr); i++ {
					if totalStr[i:i+1] != items[2][i:i+1] {
						return totalStr[i : i+1]
					}
				}
			} else {
				total = convert(items[4]) + convert(items[2])
				if total < 0 {
					total = total * -1
				}
				totalStr := strconv.Itoa(total)
				for i := 0; i < len(totalStr); i++ {
					if totalStr[i:i+1] != items[0][i:i+1] {
						return totalStr[i : i+1]
					}
				}
			}
		case "*":
			if isNumeric(items[0]) {
				total = convert(items[4]) / convert(items[0])
				totalStr := strconv.Itoa(total)
				for i := 0; i < len(totalStr); i++ {
					if totalStr[i:i+1] != items[2][i:i+1] {
						return totalStr[i : i+1]
					}
				}
			} else {
				total = convert(items[4]) / convert(items[2])
				totalStr := strconv.Itoa(total)
				for i := 0; i < len(totalStr); i++ {
					if totalStr[i:i+1] != items[0][i:i+1] {
						return totalStr[i : i+1]
					}
				}
			}
		case "/":
			if isNumeric(items[0]) {
				total = convert(items[4]) * convert(items[0])
				totalStr := strconv.Itoa(total)
				for i := 0; i < len(totalStr); i++ {
					if totalStr[i:i+1] != items[2][i:i+1] {
						return totalStr[i : i+1]
					}
				}
			} else {
				total = convert(items[4]) * convert(items[2])
				totalStr := strconv.Itoa(total)
				for i := 0; i < len(totalStr); i++ {
					if totalStr[i:i+1] != items[0][i:i+1] {
						return totalStr[i : i+1]
					}
				}
			}
		}
	} else {
		switch items[1] {
		case "+":
			total = convert(items[0]) + convert(items[2])
			totalStr := strconv.Itoa(total)
			for i := 0; i < len(totalStr); i++ {
				if totalStr[i:i+1] != items[4][i:i+1] {
					return totalStr[i : i+1]
				}
			}
		case "-":
			total = convert(items[0]) - convert(items[2])
			totalStr := strconv.Itoa(total)
			for i := 0; i < len(totalStr); i++ {
				if totalStr[i:i+1] != items[4][i:i+1] {
					return totalStr[i : i+1]
				}
			}

		case "*":
			total = convert(items[0]) * convert(items[2])
			totalStr := strconv.Itoa(total)
			for i := 0; i < len(totalStr); i++ {
				if totalStr[i:i+1] != items[4][i:i+1] {
					return totalStr[i : i+1]
				}
			}

		case "/":
			total = convert(items[0]) / convert(items[2])
			totalStr := strconv.Itoa(total)
			for i := 0; i < len(totalStr); i++ {
				if totalStr[i:i+1] != items[4][i:i+1] {
					return totalStr[i : i+1]
				}
			}
		}
	}

	return strconv.Itoa(total)

}

func convert(s string) int {
	i1, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i1
}
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
func main() {

	fmt.Println(MissingDigit("710 - 100 = 6x0"))

}
