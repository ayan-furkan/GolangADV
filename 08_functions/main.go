package main

import (
	"fmt"
)

func main() {

	// X := []float64{1, 2, 3, 4, 5, 6}

	// x, y := examples.F()
	// fmt.Println(x, y)
	// fmt.Println(examples.Average(X))

	// fmt.Println(examples.Add(1, 2, 3))

	// a, b := examples.F()
	// fmt.Println(a, b)

	// xs := []int{1, 2, 4}

	// fmt.Println(examples.Add(xs...))

	// add := func(args ...int) int {
	// 	total := 0
	// 	for _, i := range args {
	// 		total += i
	// 	}
	// 	return total
	// }
	// fmt.Println(add(1, 1))

	// makeeven := examples.MakeEvengenerator()
	// fmt.Println(makeeven())
	// fmt.Println(makeeven())
	// fmt.Println(makeeven())

	// fmt.Println(examples.Factoriel(5))

	// defer examples.Second()
	// examples.First()

	// panic("PANIC!")
	// str := recover()
	// fmt.Println(str)

	defer func() {

		str := recover()
		fmt.Println(str)

	}()
	panic("PANIC!")
}
