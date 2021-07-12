package examples

import "fmt"

func SliceEx1() {

	// var x []float64
	// x has no length in this condition
	//x := make([]float64, 5)

	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[1:5] // [2, 3, 4, 5]
	y := arr[3:4]
	fmt.Println(x)
	fmt.Println(y)
}

func SliceEx2() {

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, slice2)

}

func SliceEx3() {

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)

	fmt.Println(slice2, slice1)
}
