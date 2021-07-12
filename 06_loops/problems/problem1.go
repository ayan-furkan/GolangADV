package problems

import "fmt"

func Divthree() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
