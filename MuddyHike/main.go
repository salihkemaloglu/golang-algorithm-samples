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

func findDeepestMud(arr [][]int32, r int32, c int32) int32 {
	var path []int32
	var totalPaths [][]int32
	var deepestMud int32 = 0
	var i, j int32 = 0, 0

	for i = 0; i < c; i++ {
		for j = 0; j < r; j++ {
			if j == 0 {
				path = append(path, arr[i][j])
			} else {
				var mud int32 = arr[i][0]
				if isValid(i, j-1, r, c) {
					if arr[i][j-1] < mud {
						mud = arr[i][j-1]
					}
				}
				if isValid(i-1, j, r, c) {
					if arr[i-1][j] < mud {
						mud = arr[i-1][j]
					}
				}
				if isValid(i, j+1, r, c) {
					if arr[i][j+1] < mud {
						mud = arr[i][j+1]
					}
				}
				if isValid(i+1, j, r, c) {
					if arr[i+1][j] < mud {
						mud = arr[i+1][j]
					}
				}
				path = append(path, mud)
			}
		}
		totalPaths = append(totalPaths, path)
	}

	if j == c {
		for _, mud := range path {
			if mud > deepestMud {
				deepestMud = mud
			}
		}
	}
	return deepestMud
}

func isValid(i int32, j int32, r int32, c int32) bool {

	if (i >= 0 && i < r) && (j >= 0 && j < c) {
		return true
	}
	return false
}
func main() {

	// var r, c, i int32
	// var arr [][]int32
	// fmt.Scanf("%d%d", &r, &c)
	// for i = 0; i < r; i++ {
	// 	test, err := readEntries(c)
	// 	if err != nil {
	// 		break
	// 	}
	// 	arr = append(arr, test)
	// }
	var r, c int32 = 5, 4
	arr := [][]int32{{2, 1, 0, 8}, {3, 7, 3, 5}, {3, 1, 2, 4}, {9, 0, 4, 6}, {5, 3, 2, 3}}
	fmt.Println(findDeepestMud(arr, r, c))

}
