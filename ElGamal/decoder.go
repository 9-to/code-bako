package main

import "fmt"

func main() {
	var c1, c2, x, p int
	fmt.Println("write p")
	fmt.Scanf("%d", &p)
	fmt.Println("write cipher-code c1, c2")
	fmt.Scanf("%d %d", &c1, &c2)
	fmt.Println("write secret-key x")
	fmt.Scanf("%d", &x)
	m := 1
	for i := 0; i < (p - 1 - x); i++ {
		m = (m * c1) % p
	}
	m = (m * c2) % p
	fmt.Println("message is ", m)
}
