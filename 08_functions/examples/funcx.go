package examples

func Average(xs []float64) float64 {

	var total float64
	for _, i := range xs {
		total += i

	}
	return total / float64(len(xs))
}

func F2() (r int) {
	r = 1
	return
}

func Add(args ...int) int {
	var total int
	for _, x := range args {
		total += x
	}
	return total
}
