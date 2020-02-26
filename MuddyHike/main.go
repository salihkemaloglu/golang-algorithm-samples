package main

import "fmt"

func readEntries(n int32) ([]int32, error) {
	in := make([]int32, n)
	for i := range in {
		_, err := fmt.Scanf("%d", &in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}

func findDeepestMud(arr [][]int32) int32 {
	var deepestMud int32 = 0
	var mudArr []int32
	for i := 0; i < len(arr); i++ {
		var mud = arr[i][0]
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] < mud {
				mud = arr[i][j]
			}
		}
		mudArr = append(mudArr, mud)
	}
	fmt.Println(mudArr)
	for _, mud := range mudArr {
		if mud > deepestMud {
			deepestMud = mud
		}
	}
	return deepestMud
}

func main() {

	var r, c, i int32
	var arr [][]int32
	fmt.Scanf("%d%d", &r, &c)
	for i = 0; i < r; i++ {
		test, err := readEntries(c)
		if err != nil {
			break
		}
		arr = append(arr, test)
	}
	fmt.Println(findDeepestMud(arr))

}
