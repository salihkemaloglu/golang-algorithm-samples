package main

import (
	"fmt"
	"strconv"
)

func main() {
	test := Solution(0)
	fmt.Println(test)
}

//Solution ...
func Solution(N int) int {
	maxTotalNumberString := ""
	useOnce := true
	if N == 0 {
		maxTotalNumberString += "5" + strconv.FormatInt(int64(N), 10)

	} else if N > 0 {
		stringNumber := strconv.FormatInt(int64(N), 10)
		for i := 0; i < len(stringNumber); i++ {
			num, err := strconv.Atoi(stringNumber[i : i+1])
			if err != nil {
				fmt.Println(err)
			}
			if num > 5 {
				maxTotalNumberString += strconv.FormatInt(int64(num), 10)
			} else if useOnce {
				maxTotalNumberString += "5" + strconv.FormatInt(int64(num), 10)
				useOnce = false
			} else {
				maxTotalNumberString += strconv.FormatInt(int64(num), 10)
			}
		}
	} else {
		NPositive := N * -1
		stringNumber := strconv.FormatInt(int64(NPositive), 10)
		for i := 0; i < len(stringNumber); i++ {
			num, err := strconv.Atoi(stringNumber[i : i+1])
			if err != nil {
				fmt.Println(err)
			}
			if num < 5 {
				maxTotalNumberString += strconv.FormatInt(int64(num), 10)
			} else if useOnce {
				maxTotalNumberString += "5" + strconv.FormatInt(int64(num), 10)
				useOnce = false
			} else {
				maxTotalNumberString += strconv.FormatInt(int64(num), 10)
			}
		}
	}
	if useOnce && N != 0 {
		maxTotalNumberString += "5"
	}
	maxTotalNumberInt, err := strconv.Atoi(maxTotalNumberString)
	if err != nil {
		fmt.Println(err)
	}

	if N < 0 {
		maxTotalNumberInt = maxTotalNumberInt * -1
	}
	return maxTotalNumberInt
}
