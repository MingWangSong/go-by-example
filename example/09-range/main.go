package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// 不用range遍历
	len := len(nums)
	for i := 0; i < len; i++ {
		sum += nums[i]
		if nums[i] == 2 {
			fmt.Println("index:", i, "num:", nums[i]) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9

	// range遍历
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 18

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v) // b 8; a A
	}
	for k := range m {
		fmt.Println("key", k) // key a; key b
	}
}
