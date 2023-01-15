package main

import "fmt"

/*
	&变量名    表示   取该变量的地址（实现深拷贝）
	*数据类型  表示  	该数据类型的指针类型
	*变量名    表示   取该地址的数值，因此该变量必须为指针类型
*/

func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2ptr(&n)
	fmt.Println(n) // 7
}
