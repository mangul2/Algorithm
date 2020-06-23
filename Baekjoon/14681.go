package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Scanln(&x)
	fmt.Scanf("%d", &y)
	if x>=1 && y>=1 {
		fmt.Println(1)
	}else if x<=-1 && y>=1 {
		fmt.Println(2)
	}else if x<=-1 && y<=-1 {
		fmt.Println(3)
	}else if x>=1 && y<=-1 {
		fmt.Println(4)
	}
}