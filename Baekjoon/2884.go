package main

import "fmt"

func main() {
	var h, m int
	fmt.Scanf("%d %d", &h, &m)
	var cm int
	if m>=45 {
		m=m-45
	}else{
		cm=45-m
		m=60-cm
		if h==0 {
			h=24
		}
		h=h-1
	}
	fmt.Println(h, m)
}