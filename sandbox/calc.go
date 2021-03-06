package calc

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func FastPower(a, x, N int) (y int) {
	/*
		高速べき乗法
		y = a^x mod N
	*/
	var bina []int //xの二進数表示
	var k int      //binaの長さ
	xx := strconv.FormatInt(int64(x), 2)
	for _, c := range xx {
		bina = append(bina, (int(c - '0')))
	}
	k = len(bina)
	//fmt.Println(bina, k)
	y = a
	for i := 1; i < k; i++ {
		y = (y * y) % N
		if bina[i] == 1 {
			y = (y * a) % N
		}
		//fmt.Println(i, ":", y)
	}
	return y
}

func PrimeCheck(n, t int) bool {
	/*
		Miller-Rabinの素数判定プログラム
		素数ならば true を返す
	*/
	if n == 2 {
		return true
	} else if n <= 1 {
		return false
	} else if n%2 == 0 {
		return false
	}
	r := n - 1
	s := 0
	for r%2 == 0 {
		r /= 2
		s++
	}
	//fmt.Println(s, r)
	rand.Seed(time.Now().UnixNano())
	y := make([]int, s, s)
	for i := 0; i < s; i++ {
		y[i] = 0
	}
	for i := 0; i < t; i++ {
		flag := 0
		a := rand.Intn(n-1) + 1
		y[0] = FastPower(a, r, n)
		if y[0] == 1 || y[0] == n-1 {
			break
		}
		for j := 1; j <= s-1; j++ {
			y[j] = (y[j-1] * y[j-1]) % n
			if y[j] == n-1 {
				flag = 1
				break
			}
		}
		if flag != 1 {
			return false
		}
	}
	return true
}

func MakePE(p int) (g []int) {
	/*
		素数から原始元の数列を返す
	*/
	if !PrimeCheck(p, 10000) {
		os.Exit(1)
	}
	check := make(map[int]int) // p-1の素因数分解
	pp := p - 1
	rand.Seed(time.Now().UnixNano())
	for pp != 1 {
		a := rand.Intn(pp) + 1
		if PrimeCheck(a, 10000) {
			if pp%a == 0 {
				check[a] = a
				pp /= a
			}
		}
	}
	fmt.Println(check)
	var gg []int
	for g := 2; g < p; g++ {
		flag := false
		for a := range check {
			tmp := FastPower(g, (p-1)/a, p)
			if tmp == 1 {
				flag = true
				break
			}
		}
		if flag == false {
			gg = append(gg, g)
		}
	}
	g = gg //[rand.Intn(len(gg))]
	return g
}

func MakePrime() (p int) {
	/*
		素数を生成する
	*/
	check := 0
	rand.Seed(time.Now().UnixNano())
	for check != 1 {
		p = rand.Intn(1000)
		if PrimeCheck(p, 1000) {
			check = 1
		}
	}
	return p
}

func SafePrime() (p, q, g int) {
	/*
		安全な素数 p=2q+1 とその原始元を出力する
	*/
	check := 0
	rand.Seed(time.Now().UnixNano())
	for check != 1 {
		q = rand.Intn(1000000)
		if PrimeCheck(q, 10000) {
			p = q*2 + 1
			if PrimeCheck(p, 10000) {
				check = 1
			}
		}

	}
	fmt.Println(q, p)
	gg := MakePE(p)
	check2 := 0
	for check2 != 1 {
		g = gg[rand.Intn(len(gg))]
		tmp := FastPower(g, 2, p)
		tmp2 := FastPower(g, q, p)
		if tmp != 1 && tmp2 != 1 {
			check2 = 1
		}
	}
	return p, q, g
}

func CheckQuad(a, p int) bool {
	/*
		平方剰余かどうかを確認する
		p:法
		a:=x^2 mod p

	*/
	a = a % p
	if PrimeCheck(p, 1000) {
		g := MakePE(p)
		gg := g[0]
		for i := 1; i < p; i++ {
			if a == FastPower(gg, i, p) {
				if i%2 == 0 {
					return true
				} else {
					return false
				}
			}
		}
		return false
	} else { //素数じゃない場合
		//N-pq合成数
	}
	return false
}

func Legendre(a, p int) int {
	/*
		ルジャンドルの記号
		(a/p)
	*/
	if CheckQuad(a, p) {
		return 1
	} else {
		return -1
	}
}

func MakeQ() (p, q, a int) {
	/*
		大きな素数p、位数が大きな素数qとなるZ_p^*の要素g
	*/
	check := 0
	var R int
	rand.Seed(time.Now().UnixNano())
	for check != 1 {
		q = rand.Intn(1000)
		if PrimeCheck(q, 10000) {
			R = rand.Intn(100)
			p = q*2*R + 1
			if PrimeCheck(p, 10000) {
				check = 1
			}
		}
	}
	g := MakePE(p)
	a = FastPower(g[0], 2*R, p)
	return p, q, a
}
