package calc

func ExtGCD(a, b int) (l, m, n int) {
	/*
		拡張ユークリッドの互除法
		ax+by=1
	*/
	if b > 0 {
		var d, y, x int
		d, y, x = ExtGCD(b, a%b)
		y -= (a / b) * x
		//fmt.Println(y, x)
		return d, x, y
	}
	return a, 1, 0
}

func Inverse(p, N int) int {
	/*
		pの逆元を出力する
	*/
	_, a, _ := ExtGCD(p, N)
	return a
}
