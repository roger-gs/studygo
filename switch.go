package main

import (
	"fmt"
)

func eval(a, b int, op string) int {
	var result int
	switch op {   // switch后面带参数
	case "+":
		result = a+b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	switch {    // switch后面不带参数
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score:%d",score)) // 会中断程序的执行
	case score < 60:
		g = "E"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
func main() {
	fmt.Println(eval(3,4,"*"))
	fmt.Println(
		grade(99),
		grade(100),
		grade(77),
		grade(80),
		grade(79),
		grade(64),
		grade(43),
		)
}
