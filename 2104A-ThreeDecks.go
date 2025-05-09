package main

import "fmt"

func solve() {
	var a,b,c int
	fmt.Scan(&a, &b, &c)
	var x int = (a+b+c) / 3
	if (a+b+c) % 3 != 0 {
		fmt.Println("NO")
	} else if b <= x{
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
