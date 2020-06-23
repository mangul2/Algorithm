package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	var b int = a%4
	if b==0 {
		if (a%100)==0 {
			if (a%400)==0 {
				fmt.Println(1)
			}else{
				fmt.Println(0)
			}
		}else{
			fmt.Println(1)
		}
	}else{
		fmt.Println(0)
	}
}