package main

func main() {

	// numbers := []int{
	// 	12, 45, 85, 98,
	// 	75, 21, 8, 73,
	// 	25, 32, 96, 101,
	// }
	// fmt.Println(biggest(numbers...))

	// fmt.Println(half(5))

	// odd := makeOddGenerator()
	// fmt.Println(odd())
	// fmt.Println(odd())
	// fmt.Println(odd())

	// for i := 0; i < 15; i++ {
	// 	fmt.Print(fib(i), " ")
	// }

	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	// panic("PANIC")

}

func sum(args ...int) int {

	total := 0
	for _, x := range args {

		total += x
	}
	return total
}

func half(x uint) (uint, bool) {

	if x%2 == 1 {
		return 0, false
	}

	return 1, true

}

func biggest(args ...int) int {

	var big int
	for _, v := range args {

		if v > big {
			big = v
		}

	}
	return big

}

func makeOddGenerator() func() uint {

	x := uint(1)

	return func() (odd uint) {

		odd = x
		x += 2
		return

	}
}

func fib(a int) int {
	// 0 1 1 2 3 5 8 13 21 34

	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	} else {
		return fib(a-2) + fib(a-1)
	}
}
