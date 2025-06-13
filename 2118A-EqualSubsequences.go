package main

import "fmt"

func solve() {
	var n, k int
	fmt.Scan(&n, &k)
	for i := 0; i < k; i++ {
		fmt.Print(1)
	}
	for i := 0; i < (n - k); i++ {
		fmt.Print(0)
	}

}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
		fmt.Println()
	}
}
