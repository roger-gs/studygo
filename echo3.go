package main

import (
	"fmt"
	"os"
)

func main() {
	//	fmt.Println(strings.Join(os.Args[0:], " "))
	//	fmt.Println(os.Args[1:])
	for i, sep := range os.Args[1:] {
		fmt.Println(i, sep)
	}
}
