/*
乱数rを流用した場合、2つの暗号文C1,C2と、
1つの平文m2よりm1の内容も攻撃者にばれてしまう
c12 = m1*y^r
c22 = m2*y^r
c12/c22 = m1/m2
*/
package main

import (
	calc "code-bako/sandbox"
	"fmt"
)

func main() {
	fmt.Println("write cipher-code C1,C2")
	var c1, c2, p, m2 int
	fmt.Scanf("%d %d", &c1, &c2)
	fmt.Println("write p")
	fmt.Scanf("%d", &p)
	m1m2 := (c1 * calc.Inverse(c2, p)) % p
	fmt.Println("write m2")
	fmt.Scanf("%d", &m2)
	fmt.Println("m1 = ", (m1m2*m2)%p)
}
