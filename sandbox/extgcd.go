package calc

import (
	"math/rand"
)

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

func Sqrt(xx, N int) (out []int) {
	/*
		平方根を出力する
	*/
	var tmp int
	xx = xx % N
	for i := 1; i < N; i++ {
		if xx == (i*i)%N {
			tmp = i
		}
	}
	out = append(out, tmp)
	out = append(out, N-tmp)
	//fmt.Println(out)
	return out
}

func Sqrt2(xx, p int) (out []int) {
	/*
		改良版
		奇素数pが法の時、平方根を出力する
	*/
	var tmp int
	xx = xx % p
	if (p+1)%4 == 0 { //p=4k+3
		tmp = FastPower(xx, (p+1)/4, p)
	} else if (p+3)%4 == 0 { //p=4k+1
		b := rand.Intn(p-1) + 1
		for CheckQuad(b, p) {
			b = rand.Intn(p-1) + 1
		}
		t := p - 1
		s := 0
		for t%2 == 0 {
			t /= 2
			s++
		}
		c := FastPower(b, t, p)
		r := FastPower(xx, (t+1)/2, p)
		ss := 1
		for i := 1; i < s; i++ {
			ss *= 2
		}
		for i := 1; i < s; i++ {
			ss = ss / 2
			d := FastPower(Legendre(r*r, xx), ss, p)
			if d == -1 {
				r = (r * c) % p
			}
			c = (c * c) % p
		}
		tmp = r
	} else {
		for i := 1; i < p; i++ {
			if xx == (i*i)%p {
				tmp = i
			}
		}
	}
	out = append(out, tmp)
	out = append(out, p-tmp)
	return out
}
