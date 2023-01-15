package main

import "fmt"

/*type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user{name: "wang", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048")) // true
}*/

type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

// 如需修改结构体内容，则需写结构体方法
func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user{"song", "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048")) // true
	fmt.Println(a.checkPassword("1024")) // false
}
