package main

import "fmt"

func main() {
	xy := []int{1, 5, 6, 7}
	xs := []float64{1, 2, 3, 4, 5}
	fmt.Println(average(xs))
	fmt.Println(f2())
	fmt.Println(add(1, 2, 3, 4))
	fmt.Println(add(xy...))
}

func average(xs []float64) float64 {

	var total float64
	for _, i := range xs {
		total += i

	}
	return total / float64(len(xs))
}

func f2() (r int) {
	r = 1
	return
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
