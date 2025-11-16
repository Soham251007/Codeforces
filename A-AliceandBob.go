package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, a int
	fmt.Fscan(in, &n, &a)

	small, big := 0, 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		if x < a {
			small++
		} else if x > a {
			big++
		}
	}

	if small > big {
		fmt.Fprintln(out, a-1)
	} else {
		fmt.Fprintln(out, a+1)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		solve(in, out)
	}
}
