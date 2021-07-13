package examples

func F() (int, int) {
	return 5, 6
}

func MakeEvengenerator() func() uint {
	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
