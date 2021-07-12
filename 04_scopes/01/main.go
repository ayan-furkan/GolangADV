package main

import "fmt"

var a = 0

func foo() int {
	a++
	return a
}

func main() {

	fmt.Println(foo())
	fmt.Println(foo())

}
