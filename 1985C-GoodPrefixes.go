package main
import "fmt"

func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var sum, x , y int
	sum = 0
	x = 0
	y = 0
	for i := 0; i<n;i++{
		sum += a[i]
		x = max(x, a[i])
		if sum - x == x {
			y++
		}
	}
	fmt.Println(y)
}

func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}