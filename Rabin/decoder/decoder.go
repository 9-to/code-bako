package decoder

import (
	calc "code-bako/sandbox"
	"fmt"
)

func Decoder(c, p1, p2 int) (m []int) {
	x1 := calc.Sqrt2(c, p1)
	a1 := x1[0]
	a2 := x1[1]
	x2 := calc.Sqrt2(c, p2)
	b1 := x2[0]
	b2 := x2[1]
	fmt.Println("x1,x2:", x1, x2)
	q1 := calc.Inverse(p1, p2) //中国人の剰余定理
	q2 := calc.Inverse(p2, p1)
	m = append(m, (b1*p1*q1+a1*p2*q2)%(p1*p2))
	m = append(m, (b1*p1*q1+a2*p2*q2)%(p1*p2))
	m = append(m, (b2*p1*q1+a1*p2*q2)%(p1*p2))
	m = append(m, (b2*p1*q1+a2*p2*q2)%(p1*p2))
	return m
}
