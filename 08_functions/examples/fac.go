package examples

func Factoriel(x uint) uint {

	if x == 0 {
		return 1
	}

	return x * Factoriel(x-1)
}
