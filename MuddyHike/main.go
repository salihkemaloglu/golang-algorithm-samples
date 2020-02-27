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
	revArr:=reverseArray(arr)
	var deepestMud int32 = 0
	jmud := 0
	var mudArr []int32
	for i := 0; i < len(revArr); i++ {
		var mud int32 = 0
		if i == 0 {
			fmt.Println("jmud1:", jmud)
			mud = revArr[i][0]
		} else {
			fmt.Println("jmud1:", jmud)
			mud = revArr[i][jmud]
		}
		for j := 0; j < len(revArr[i]); j++ {
			if i == 0 {
				if revArr[i][j] < mud {
					mud = revArr[i][j]
					fmt.Println("mud5", mud)
					jmud = j
				}
				fmt.Println("jmud1:", jmud)
				fmt.Println("-------")
			} else {
				fmt.Printf("jmud+1:%v ,jmud-1:%v ,j:%v", jmud+1, jmud-1, j)
				fmt.Println("")
				if j != jmud+1 && j != jmud-1 {
					fmt.Println("geldi")
					if revArr[i][j] < mud {
						mud = revArr[i][j]
						fmt.Println("mud3", mud)
						jmud = j
					}
				}
				fmt.Println("-------")

			}
		}
		fmt.Println("mudlast", mud)
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
