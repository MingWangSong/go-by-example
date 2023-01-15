package main

import "fmt"

// 结构体定义
type user struct {
	name     string
	password string
}

func main() {
	// 结构体赋值
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "wang"}
	c.password = "1024"
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d)               // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "haha")) // false
	fmt.Println(a)
	fmt.Println(checkPassword2(&a, "haha")) // false
	fmt.Println(a)
}

func checkPassword(u user, password string) bool {
	u.name = "song_no"
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	u.name = "song_*"
	return u.password == password
}
