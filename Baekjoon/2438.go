package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	for i := a; i > 0; i-- {
		for j := a; j >= i; j-- {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}