package examples

import "fmt"

func Problem1() {

	arr := [4]int32{1, 2, 3, 4}

	fmt.Println(arr[3])

}

func Problem2() {

	slc := make([]int, 3, 9)
	fmt.Println(slc)
}

func Problem3() {

	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
}

func Problem4() {

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := x[0]
	for _, element := range x {

		if element < smallest {
			smallest = element
		}

	}
	fmt.Println(smallest)
}
