package main
import "fmt"

func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var m int
	m = 0
	for i := 0; i < n - 1; i++ {
		if a[i] == a[i+1] && a[i] == 0 {
			m = 1
		}
	}
	var l int
	l = 0
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			l = 1
		}
	}
	if m == 1 {
		fmt.Println("YES")
	} else if l == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}



func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}