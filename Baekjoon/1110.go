package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g int
	fmt.Scan(&a)
	f = a
	for {
		g = g + 1
		b = a / 10
		c = a % 10
		d = b + c
		if d >= 10 {
			d = d % 10
		}
		e = c*10 + d
		if a < 10 {
			a = a * 11
		} else {
			a = e
		}
		if a == f {
			fmt.Print(g)
			return
		}
	}
}
