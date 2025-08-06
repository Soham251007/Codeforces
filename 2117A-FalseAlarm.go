package main
import "fmt"

func solve(){
	var n,x int
	fmt.Scan(&n, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	b := 0
	c := 0
	for i := 0; i < n; i++ {
		if a[i] == 1{
			for j := i; j < n; j++ {
				b++
				if a[j] == 1 {
					c = j
				}
			}
			break
		}
	}
	c = n - 1 - c
	b = b - c
	if b <= x{
		fmt.Println("YES")
	}else{
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