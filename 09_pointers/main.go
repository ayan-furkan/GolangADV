package main

import "fmt"

// func zero(x int) {
// 	x = 0
//   }
//   func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x) // x is still 5
//   }

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	// x := 5
	// zero(&x)
	// fmt.Println(x) // x is 0

	a := 1

	b := 2

	fmt.Println(a, b)

	swap(&a, &b)

	fmt.Println(a, b)
}

// func one(xPtr *int) {
// 	*xPtr = 1
//   }
//   func main() {
// 	xPtr := new(int)
// 	one(xPtr)
// 	fmt.Println(*xPtr) // x is 1
//   }

func swap(a *int, b *int) {

	// *a ,*b = *b, *a	OR

	tempa := *a

	*a = *b

	*b = tempa

}
