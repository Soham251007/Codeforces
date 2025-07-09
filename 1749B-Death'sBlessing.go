package main

import (
	"fmt"
	"sort"
)

func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	sort.Ints(b)
	ans := 0
	ans1 := 0
	for i := 0; i < n; i++ {
		ans += a[i]
		if i < n-1 {
			ans1 += b[i]
		}
	}
	
	fmt.Println(ans + ans1)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}