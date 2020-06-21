package main

import "fmt"

var (
	tag bool
)

func main() {
	fmt.Println("Hello, World!")
	var str string = fmt.Sprintf("%s%d%s", "Alice is ", 10, " years old.")
	fmt.Println(str)

	b, c := 5, 7
	fmt.Println(b, c)

	tag = false
	fmt.Println(tag)
}
