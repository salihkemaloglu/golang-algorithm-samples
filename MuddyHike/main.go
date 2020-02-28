package main

import (
	"fmt"
)

var arr [][]int32
var arrVisit []string
var i, j int32 = 0, 0
var path []int32

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

func findDeepestMud(r int32, c int32) int32 {

	var deepestMud, xi, xj int32 = 0, 0, 0
	if j == 0 {
		j = findSmallest(arr[0])
		visit := fmt.Sprint(i) + "," + fmt.Sprint(j)
		arrVisit = append(arrVisit, visit)
		path = append(path, arr[i][j])
	} else {
		var mud int32 = 0
		if isValid(i, j-1, r, c) && isVisited(arrVisit, i, j-1) {
			fmt.Println(arr[i][j-1])
			if mud == 0 && arr[i][j-1] > 0 {
				mud = arr[i][j-1]
				xi = i
				xj = j - 1
			} else if arr[i-1][j] <= mud {
				mud = arr[i][j-1]
				xi = i
				xj = j - 1
			}
		}
		if isValid(i-1, j, r, c) && isVisited(arrVisit, i-1, j) {
			fmt.Println(arr[i-1][j])
			if mud == 0 && arr[i][j-1] > 0 {
				mud = arr[i][j-1]
				xi = i
				xj = j - 1
			} else if arr[i-1][j] <= mud {
				mud = arr[i][j-1]
				xi = i
				xj = j - 1
			}
		}
		if isValid(i, j+1, r, c) && isVisited(arrVisit, i, j+1) {
			fmt.Println(arr[i][j+1])
			if mud == 0 && arr[i][j-1] > 0 {
				mud = arr[i][j-1]
				xi = i
				xj = j - 1
			} else if arr[i-1][j] <= mud {
				mud = arr[i][j-1]
				xi = i
				xj = j - 1
			}
		}
		if isValid(i+1, j, r, c) && isVisited(arrVisit, i+1, j) {
			fmt.Println(arr[i+1][j])
			if mud == 0 && arr[i][j-1] > 0 {
				mud = arr[i][j-1]
				xi = i
				xj = j - 1
			} else if arr[i-1][j] <= mud {
				mud = arr[i][j-1]
				xi = i
				xj = j - 1
			}
		}
		i = xi
		j = xj
		visit := fmt.Sprint(i) + "," + fmt.Sprint(j)
		arrVisit = append(arrVisit, visit)
		path = append(path, mud)
	}

	if j == c {
		for _, mud := range path {
			if mud > deepestMud {
				deepestMud = mud
			}
		}
		return deepestMud
	}

	return findDeepestMud(r, c)
}

func isValid(i int32, j int32, r int32, c int32) bool {

	if (i >= 0 && i < r) && (j >= 0 && j < c) {
		return true
	}
	return false
}
func isVisited(arrVisit []string, i int32, j int32) bool {
	visit := fmt.Sprint(i) + "," + fmt.Sprint(j)
	for i := 0; i < len(arrVisit); i++ {
		if visit == arrVisit[i] {
			return false
		}
	}
	return true
}
func findSmallest(arr []int32) int32 {
	small := arr[0]
	var index, i int32
	for i = 0; i < int32(len(arr)); i++ {
		if arr[i] < small {
			small = arr[i]
			index = i
		}
	}
	return index
}
func reverseArray(arr [][]int32) [][]int32 {

	var twoD [][]int32
	for i := 0; i < len(arr[0]); i++ {
		var oneD []int32
		for j := 0; j < len(arr); j++ {
			oneD = append(oneD, arr[j][i])
		}
		twoD = append(twoD, oneD)
	}
	return twoD
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
	arr = reverseArray([][]int32{{2, 1, 0, 8}, {1, 7, 3, 5}, {3, 1, 2, 4}, {9, 0, 4, 6}, {5, 3, 2, 3}})
	fmt.Println(findDeepestMud(r, c))

}
