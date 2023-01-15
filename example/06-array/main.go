package main

import "fmt"

func main() {

	// 数组声明方式1
	var a [5]int
	a[4] = 66666666
	fmt.Println("get a[2]:", a[2])
	fmt.Println("get a[4]:", a[4])
	fmt.Println("len:", len(a))

	// 数组声明方式2
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 二维数组声明方式
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
