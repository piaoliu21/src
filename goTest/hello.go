package main

import "fmt"

var a, b int

func main() {
	fmt.Println("输入第1个整数")
	_, _ = fmt.Scanln(&a)
	fmt.Println("输入第2个整数")
	_, _ = fmt.Scanln(&b)
	sum, avg := calc(a, b)

	fmt.Println("sum=", sum, ":", "avg=", avg)
}
