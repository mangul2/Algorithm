package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	var c float64 = a / b
	fmt.Printf("%f", c)
}
