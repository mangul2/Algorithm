package main

import "fmt"

func main() {
	a := 0
	b := 0
	c := 0
	fmt.Scan(&c)
	for a < c {
		a += 1
		b += a
	}
	fmt.Println(b)
}
