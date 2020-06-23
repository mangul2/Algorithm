package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	var b string
	if a>=90 {
		b="A"
	}else if a>=80 {
		b="B"
	}else if a>=70 {
		b="C"
	}else if a>=60 {
		b="D"
	}else {
		b="F"
	}
	fmt.Println(b)
}