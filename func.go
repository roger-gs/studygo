package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func jisuan(a, b int,op string) (int,error) {
	switch op {
	case "+":
		return a + b,nil
	case "-":
		return a - b,nil
	case "*":
		return a * b,nil
	case "/":
		x,_ := div(a,b)
		return x,nil
	default:
		return 0,fmt.Errorf("unknown operator %s", op)
	}
}
func apply(op func(int,int) int,a,b int) int {
	p := reflect.ValueOf(op).Pointer()   // 反射
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" + "(%d, %d)\n",opName,a,b)
	return op(a,b)
}
func div(a,b int) (q,r int)  {
	q = a / b
	r = a % b
	return
}

func sum(numbers ...int) int {     // 可变参数列表
	s := 0
	for i := range numbers {
		s += i
	}
	return s
}

func pow(a,b int) int {
	return int(math.Pow(float64(a),float64(b)))
}
func main() {
	if reslut,err := jisuan(4,6,"x");err != nil {
		fmt.Println("Error: ",err)
	}else {
		fmt.Println(reslut)
	}
	fmt.Println(jisuan(5,9,"*"),
	)
	fmt.Println(div(8,6))
	fmt.Println(apply(pow,3,4))
	fmt.Println(sum(1,2,3,4,5,6,7))
}