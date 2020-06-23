package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d", &b)
	fmt.Println(a * (b % 10))
	fmt.Println(a * ((b % 100) / 10))
	fmt.Println(a * (b / 100))
	fmt.Println(a * b)
}
