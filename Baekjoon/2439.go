package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		for j := 1; j <= a; j++ {
			if j <= a-i {
				fmt.Printf(" ")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
}
