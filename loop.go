package main

import (
	"bufio"
	`fmt`
	"os"
	`strconv`
)
func suma(n int) int {
	sum := 0
	for ;n >= 0;n-- {    // 省略初始条件
		sum += n
	}
	return sum
}
func intToBin(a int) string {  // 将整数转换为二进制
	result := ""
	for ;a > 0;a /= 2{   // 省略初始条件
		lsb :=a % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func printFile(files string){
	file,err := os.Open(files)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func forever() {
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(intToBin(20))
	fmt.Println(suma(10000))
	fmt.Println(intToBin(200))
	printFile("abc.txt")
	forever()
}