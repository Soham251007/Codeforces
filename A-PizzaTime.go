package main

import "fmt"

func solve() {
	var n int
	fmt.Scan(&n)
	ans := 0
	for n > 2 {
		ans += int(n / 3)
		n = n/3 + n%3
	}
	fmt.Println(ans)
}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
