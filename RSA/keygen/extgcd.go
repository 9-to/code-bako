package keygen

func ExtGCD(a, b int) (l, m, n int) {
	if b > 0 {
		var d, y, x int
		d, y, x = ExtGCD(b, a%b)
		y -= (a / b) * x
		//fmt.Println(y, x)
		return d, x, y
	}
	return a, 1, 0
}
