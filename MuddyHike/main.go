package main

import (
	"fmt"
	"strings"
)

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

var multiPaths []string
var isMultiPath bool
var isParent bool = false
var isMultiPathCount int32 = 0

type point struct {
	xi int32
	ji int32
}

func findDeepestMud(r int32, c int32, i int32, j int32, arr [][]int32, path []int32, arrVisit []string) []int32 {

	var xi, xj, mud int32 = 0, 0, 0
	var paths []int32
	var check bool = true
	count := 0
	//up
	if isValid(i-1, j, r, c) && isVisited(arrVisit, i-1, j) && isMultipathVisited(i-1, j) {
		if check && arr[i-1][j] >= 0 {
			mud = arr[i-1][j]
			xi = i - 1
			xj = j
			paths = append(paths, mud)
			check = false
			count++
		} else if arr[i-1][j] <= mud {
			mud = arr[i-1][j]
			xi = i - 1
			xj = j
			paths = append(paths, mud)
			count++
		}
	}
	//right
	if isValid(i, j+1, r, c) && isVisited(arrVisit, i, j+1) && isMultipathVisited(i, j+1) {
		if check && arr[i][j+1] >= 0 {
			mud = arr[i][j+1]
			xi = i
			xj = j + 1
			paths = append(paths, mud)
			check = false
			count++
		} else if arr[i][j+1] <= mud {
			mud = arr[i][j+1]
			xi = i
			xj = j + 1
			paths = append(paths, mud)
			count++
		}
	}
	//left
	if isValid(i, j-1, r, c) && isVisited(arrVisit, i, j-1) && isMultipathVisited(i, j-1) {
		if check && arr[i][j-1] >= 0 {
			mud = arr[i][j-1]
			xi = i
			xj = j - 1
			paths = append(paths, mud)
			check = false
			count++
		} else if arr[i][j-1] <= mud {
			mud = arr[i][j-1]
			xi = i
			xj = j - 1
			paths = append(paths, mud)
			count++
		}
	}
	//down
	if isValid(i+1, j, r, c) && isVisited(arrVisit, i+1, j) && isMultipathVisited(i+1, j) {
		if check && arr[i+1][j] >= 0 {
			mud = arr[i+1][j]
			xi = i + 1
			xj = j
			paths = append(paths, mud)
			check = false
			count++
		} else if arr[i+1][j] <= mud {
			mud = arr[i+1][j]
			xi = i + 1
			xj = j
			paths = append(paths, mud)
			count++
		}
	}
	i = xi
	j = xj

	if findMultiPathCount(paths, mud) && isMultipathVisited(i, j) {
		isMultiPathCount++
		if isMultiPathCount > 1 {
			makeNoneVisited()
		}
		visit := fmt.Sprint(i) + "," + fmt.Sprint(j) + ":" + "1"
		multiPaths = append(multiPaths, visit)
		isMultiPath = true

	}
	visit := fmt.Sprint(i) + "," + fmt.Sprint(j)
	arrVisit = append(arrVisit, visit)

	path = append(path, mud)
	if i == c-1 {
		return path
	} else if i == 0 && j == 0 && count == 0 {
		return nil
	}

	return findDeepestMud(r, c, i, j, arr, path, arrVisit)
}
func makeNoneVisited() {
	for i := 0; i < len(multiPaths); i++ {
		arr := strings.Split(multiPaths[i], ":")

	}
}
func findMultiPathCount(paths []int32, mud int32) bool {
	var count int32
	for _, path := range paths {
		if path == mud {
			count++
		}
	}
	if count >= 2 {
		return true
	}
	return false
}

func isValid(i int32, j int32, r int32, c int32) bool {

	if (i >= 0 && i < c) && (j >= 0 && j < r) {
		return true
	}
	return false
}

func isMultipathVisited(i int32, j int32) bool {
	visit := fmt.Sprint(i) + "," + fmt.Sprint(j)
	for i := 0; i < len(multiPaths); i++ {
		arr := strings.Split(multiPaths[i], ":")

	}
	return true
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
	smallest := arr[0]
	var i, j int32
	for i = 0; i < int32(len(arr)); i++ {
		if arr[i] < smallest {
			j = i
		}
	}
	return j
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
func findTotalMudPath(path []int32) (int32, int32, int32) {
	var deepestMud, totalMudPath, count int32 = 0, 0, 0
	for _, mud := range path {
		if mud > deepestMud {
			deepestMud = mud
		}
		count++
		totalMudPath += mud
	}
	return deepestMud, totalMudPath, count
}
func main() {

	// var r, c, i int32
	// var tempArr [][]int32
	// fmt.Scanf("%d%d", &r, &c)
	// for i = 0; i < r; i++ {
	// 	test, err := readEntries(c)
	// 	if err != nil {
	// 		break
	// 	}
	// 	tempArr = append(tempArr, test)
	// }
	var r, c int32 = 5, 4
	tempArr := [][]int32{
		{3, 1, 0, 3},
		{3, 2, 3, 5},
		{2, 1, 2, 0},
		{9, 2, 4, 6},
		{5, 3, 2, 3}}
	arr := reverseArray(tempArr)
	var j, minMudPath, deepestMud, count int32 = 0, 0, 0, 0
	check := true
	indexj := findSmallest(arr[0])
	smalletsMud := arr[0][indexj]
	for j = 0; j < int32(len(arr[0])); j++ {
		if arr[0][j] <= smalletsMud {
			isMultiPath = false
			visit := fmt.Sprint(0) + "," + fmt.Sprint(j)
			arrVisited := []string{visit}
			path := []int32{arr[0][j]}
			path = append(path, findDeepestMud(r, c, 0, j, arr, []int32{}, arrVisited)...)
			fmt.Println(path)
			if len(path) > 1 {
				mud, totalMudPath, totalCount := findTotalMudPath(path)
				if check {
					minMudPath = totalMudPath
					deepestMud = mud
					check = false
					count = totalCount
				} else if totalMudPath == minMudPath && count <= totalCount {
					minMudPath = totalMudPath
					deepestMud = mud
				} else if totalMudPath < minMudPath {
					minMudPath = totalMudPath
					deepestMud = mud
				}
				if isMultiPath {
					j--
				} else {
					multiPaths = nil
				}
			} else {
				multiPaths = nil
			}

		}
	}
	fmt.Println(deepestMud)
}
