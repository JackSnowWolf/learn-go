package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var str string = fmt.Sprintf("%s%d%s", "Alice is ", 10, " years old.")
	fmt.Println(str)

	var b, c = 1, 2
	fmt.Println(b, c)

	var tag = false
	fmt.Println(tag)
}
