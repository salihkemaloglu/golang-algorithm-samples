package main

import (
	"fmt"
	"io"
	"math"
)

func main() {

	var n, k, ti int
	fmt.Scanf("%d%d", &n, &k)
	var arr []int
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &ti)
		if err == io.EOF {
			break
		}
		arr = append(arr, ti)
		if ti-arr[0] >= 1000 {
			arr = removeIndex(arr, 0)
		}
	}
	fmt.Println(math.Ceil(float64(len(arr)) / float64(k)))
}
func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
