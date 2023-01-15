package main

import "fmt"

func main() {
	s := make([]string, 4)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3
	fmt.Println("cap:", cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]
	fmt.Println("new len:", len(s))
	fmt.Println("new cap:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	// 切片声明方式2
	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]
}
