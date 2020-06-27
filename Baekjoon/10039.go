package main

import "fmt"

func main() {
	var a, b int
	for i := 0; i <= 4; i++ {
		fmt.Scan(&a)
		if a < 40 {
			a = 40
		}
		b = b + a
	}
	fmt.Println(b / 5)
}
