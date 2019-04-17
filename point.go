package main

import "fmt"
func add(a int) int {
	a += 1
	return a
}
func main() {
	var a int = 5
	var pa *int   // 定义指针
	pa =  &a  // 指针指向a的地址
	*pa += 1  // 引用传递
	fmt.Println(pa,*pa,a)
	fmt.Println("After add()!")
	fmt.Println(add(a),a)   // go语言是值传递，变量在做运算时，会将值拷贝一份，原变量值不变
}

