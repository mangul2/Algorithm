package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var num = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &num[i])
	}

	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})

	fmt.Printf("%d %d\n", num[0], num[len(num)-1])
}
