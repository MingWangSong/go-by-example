package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	// 必须之前设置好随机种子，不然程序启动使用默认随机种子，导致后续的随机数不会变化
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("Please input your guess")
	// 输入流获取
	reader := bufio.NewReader(os.Stdin)
	// while循环
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		// 去除后缀
		input = strings.Trim(input, "\r\n")
		// 字符转数字
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("猜大了！！！！请再猜")
		} else if guess < secretNumber {
			fmt.Println("猜小了！！！！请再猜")
		} else {
			fmt.Println("猜对了！！！！！！！")
			break
		}
	}
}
