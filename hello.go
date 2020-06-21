package main

import "fmt"

var (
	tag bool
)

func main() {
	fmt.Println("Hello, World!")
	var str string = fmt.Sprintf("%s%d%s", "Alice is ", 10, " years old.")
	fmt.Println(str)

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("1st line c value: %d\n", c)
	c = a - b
	fmt.Printf("2nd line c value: %d\n", c)
	c = a * b
	fmt.Printf("3rd line c value:%d\n", c)
	c = a / b
	fmt.Printf("4th line c value: %d\n", c)
	c = a % b
	fmt.Printf("5th line c value: %d\n", c)
	a++
	fmt.Printf("6th line c value: %d\n", a)
	a = 21
	a--
	fmt.Printf("7th line c value: %d\n", a)

	tag = false
	fmt.Println(tag)
}
