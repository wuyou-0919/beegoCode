package main

import "fmt"

func main() {
	var date int
	fmt.Println("该月份属于什么季节:")
	fmt.Scanln(&date)
	switch date {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	}
}