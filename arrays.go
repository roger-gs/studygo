package main

import (
	"fmt"
)

type Currency int  // 定义货币类型
const (
	A Currency = iota  // 0
	B // 1
	C  // 2
)
func main() {
	var a = [7]int{1, 2, 3, 4, 5, 6, 7}
	var c = [...]string{"a","b","c","d"}
	b := a[3:5]
	var d = [...]string{A:"RMB",B:"USD",C:"EUR"}   // 同时定义索引与值
	e := [...]int{6:100}
	fmt.Println(a[2:3],a[1:],a[:5])
	fmt.Println(c)
	for i,v := range a {
		fmt.Println(i,v)
	}
	fmt.Println(B,d[A])
	fmt.Println(b,b[1:4],len(b),cap(b))
	fmt.Println(e,len(e),cap(e))
}

