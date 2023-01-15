package main

import (
	"fmt"
	"math"
)

/*func main() {
	var a = "initial"
	var b, c int = 1, 2
	var d = true
	var e float64
	f := float32(e)
	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple
	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}*/

func main() {
	// 常用数据类型有字符串，整数，浮点型，布尔型；字符串是内置类型
	// 声明方式
	var x1 = 1
	x2 := 2
	x3 := int8(3)
	var x4 int8 = 4
	fmt.Println("输出：", x1, x2, x3, x4) // 个人偏向方式2和3；因为方式2推断类型省事，方式3可以明确指定类型

	//var a = 1
	a := 1
	c := int8(3)
	var e float64
	f := float32(e)
	var b = "www."
	g := b + "foo"
	fmt.Println(a, c, f, g)

	// 常量
	const s string = "constant"
	const h = 50000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
