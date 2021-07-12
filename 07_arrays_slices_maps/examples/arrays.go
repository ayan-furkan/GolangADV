package examples

import "fmt"

func Ex1() {

	var X [4]float64

	X[0] = 54
	X[1] = 9
	X[2] = 91
	X[3] = 72

	var total float64 = 0
	for i := 0; i <= 3; i++ {

		total += X[i]
	}
	fmt.Println("Avarage1 = ", total/float64(len(X)))
}

func Ex2() {

	var X [4]float64

	X[0] = 54
	X[1] = 9
	X[2] = 91
	X[3] = 72

	var total1 float64

	for _, value := range X {
		total1 += value

	}
	fmt.Println("Avarage2 = ", total1/float64(len(X)))
}

func Ex3() {

	X := [4]float64{5, 6, 7, 8}
	Y := [4]float64{
		1,
		2,
		3,
		4,
	}

	fmt.Print(X, "\n", Y, "\n")
}
