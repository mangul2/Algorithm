package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &a)

	var b, c int

	for i := 1; i <= a; i++ {
		fmt.Fscan(reader, &b, &c)
		fmt.Fprintf(writer, "Case #%d: %d\n", i, b+c)
	}
	writer.Flush()
}
