package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var a int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &a)

	for i:=0; i<=1; i++ {
		fmt.Fprintln(writer, a)
		a--;
	}
	writer.Flush()
}