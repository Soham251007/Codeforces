package main 
import ("fmt"
	"math")

func absDiff(x, y int) int {
	return int(math.Abs(float64(x - y)))
}

func solve() {
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var x int
	x = 0
	for i := 0; i < n-1; i++ {
		if absDiff(a[i], a[i+1]) == 5 {
			x ++
		} else if absDiff(a[i], a[i+1]) == 7 {
			x ++
		}
	}
	if x == n - 1{
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