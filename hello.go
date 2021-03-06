package main

import (
	"fmt"
	. "github.com/jacksnow_wolf/learn-go/internal"
	"reflect"
	"strings"
)

const (
	WAITING = uint8(iota)
	COMPLETED
	PROCRESSING
)

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("1st line a type = %T\n", a)
	fmt.Printf("2nd line b type = %T\n", b)
	fmt.Printf("3rd line c type = %T\n", c)

	ptr = &a
	fmt.Printf("value of a : %d\n", a)
	fmt.Printf("value of *ptr : %d\n", *ptr)
	fmt.Printf("value of ptr : %d\n", ptr)

	logger := GetLogger()
	logger.Info("Hello World!")

	str := "(address)"

	fmt.Println(strings.Trim(str, "()"))
	fmt.Println(reflect.TypeOf(WAITING).String())
	fmt.Println(reflect.TypeOf(COMPLETED).String())
	fmt.Println(reflect.TypeOf(PROCRESSING).String())
}
