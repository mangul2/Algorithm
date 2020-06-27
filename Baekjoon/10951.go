package main

import "fmt"

func main() {
	var a, b int
	for {
		val, _ := fmt.Scan(&a, &b)
		if val != 2 {
			break
		}
		fmt.Println(a + b)
	}
}
