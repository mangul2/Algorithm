package main

import "fmt"

func main() {
	var a, b, mina, minb int
	mina = 2001
	minb = 2001
	for i := 0; i <= 2; i++ {
		fmt.Scan(&a)
		if mina > a {
			mina = a
		}
	}
	for i := 0; i <= 1; i++ {
		fmt.Scan(&b)
		if minb > b {
			minb = b
		}
	}
	fmt.Print(mina + minb - 50)
}
