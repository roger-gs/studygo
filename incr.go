//对C语言++v的模拟

package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

var v int = 1

func main() {
	incr(&v)
	incr(&v)
	fmt.Println(incr(&v))
}
