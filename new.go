//使用new(type)函数创建一个类型为type的匿名变量
package main

import "fmt"

func main() {
	v := new(int)
	fmt.Println(*v)
	*v = 3
	fmt.Println(*v)
}
