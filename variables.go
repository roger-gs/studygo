package main

import (
	"fmt"
	"math"
	"math/cmplx"
)
// 在函数外部不能使用:=进行赋值
var aa = 4
var bb = "hello world"
var cc = true
// 多个赋值可以使用var()的方式
var (
	dd = false
	ee = 4
)

func variableZeroValue() {
	var a int  // 初值为0
	var s string  // 初值为空字符串
	fmt.Printf("%d %q\n",a,s)
}

func variableInit() {
	var b,e int  = 3,5  // 类型可以省去
	var c string = "hello"  // 类型可以省去
	fmt.Println(b,c,e)
}

func euler() {  // 欧拉公式
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))   // 取c的模
}

func triangle() {   // 三角函数
	var  a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}
func variableTypeDe() {
	var a, b, c, d = 3, 4, "abc",true   // 类型省略复制
	fmt.Println(a,b,c,d)
}

func variableShort() {
	a,b,c,d := 1,4,"world" ,false   // 只能在函数内部使用
	b = 5  // 给b重新赋值，不能使用:=
	fmt.Println(a,b,c,d)
}

func enums() {
	const (
		cpp = 0
		java = 1
		python =2
		golang =3
	)
	const (
		cpp1 = iota  // iota是自增值
		java1
		python1
		golang1
	)
	// b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,golang)
	fmt.Println(cpp1,java1,python1,golang1)
	fmt.Println(b,kb,mb,gb,tb,pb)
}
func main() {
	variableZeroValue()
	variableInit()
	variableTypeDe()
	variableShort()
	fmt.Println(aa,bb,cc,dd,ee)
	euler()
	triangle()
	enums()
}