//功能是将输入的英尺长度转换为米长度
package main

import (
	"fmt"
	"os"
	"strconv"
)

// 定义英尺和米的类型
type Foot float64
type Mile float64

func FtoM(f Foot) Mile { // 英尺转换为米
	return Mile(f / 0.3048)
}

func MtoF(m Mile) Foot { // 米转换为英尺
	return Foot(0.3048 * m)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(2)
		}
		f := Foot(t)
		m := Mile(t)
		fmt.Printf("foot：%v，mile：%v", FtoM(f), MtoF(m))
	}
}
