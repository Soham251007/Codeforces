package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
	return false
	}
	for i := 2; i*i <= n; i++ {
	if n%i == 0 {
	return false
	}
	}
	return true
	}

func solve() {
	var x, k int
	fmt.Scan(&x, &k)
	if k > 1 && x > 1 {
		fmt.Println("NO")
	} else if k == 1 {
		if isPrime(x) == true {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else if k == 2 && x == 1 {
		fmt.Println("YES")
	} else if k > 2 && x == 1 {
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
