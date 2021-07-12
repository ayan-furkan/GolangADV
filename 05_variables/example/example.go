package example

import "fmt"

func Celcius() float32 {
	fmt.Print("Fahrenheit = ")
	var input float32
	fmt.Scanf("%f", &input)
	cel := ((input - 32) * 5 / 9)
	return cel
}

func FeetMeter() float32 {

	fmt.Print("Feet = ")
	var feet float32
	fmt.Scanf("%f", &feet)
	meter := feet * 0.3048

	return meter

}
