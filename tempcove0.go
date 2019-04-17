// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
	AbsoluteZeroK Kelvin  = 273.15
)

type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度
type Kelvin float64     //开尔文温度

//华氏温度转换为摄氏温度
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//开尔文温度转换为摄氏温度
func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

//摄氏温度转换为开尔文温度
func CtoK(c Celsius) Kelvin {
	return Kelvin(273.15 + c)
}

//摄氏温度转换为华氏温度
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%v = %v, %v = %v\n",
			f, FtoC(f), c, CtoF(c))
	}
}
