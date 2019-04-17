package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
/*	const filename = "abc,txt"
	contents,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("%s\n",contents)
	}
*/
	const filename = "abc.txt"
	if contents,err := ioutil.ReadFile(filename);err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
}
