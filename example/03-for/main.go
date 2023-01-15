package main

import "fmt"

func main() {

	// 死循环
	i := 1
	for {
		fmt.Println("loop")
		break
	}

	// 常规循环
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	// continue和break
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		if n == 3 {
			break
		}
		fmt.Println(n)
	}

	// 类似while循环
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}
