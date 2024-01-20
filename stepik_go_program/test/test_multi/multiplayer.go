package multi

func ExpMulti(n ...int) (res int) {
	res = 1
	for _, item := range n {
		res *= item
	}
	return res
}
