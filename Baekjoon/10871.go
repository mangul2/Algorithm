package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)

	for i := 0; i < a; i++ {
		fmt.Scan(&c)
		if c < b {
			fmt.Printf("%d ", c)
		}
	}
}
