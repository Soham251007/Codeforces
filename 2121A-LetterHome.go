package main
import ("fmt"
	"math")

func solve(){
	var n,s int 
	fmt.Scan(&n, &s)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var first int
	first = int(math.Abs(float64(s - a[0]))) 
	var last int 
	last = int(math.Abs(float64(s - a[n-1])))
	var ans int
	ans = min(first, last) + a[n-1] - a[0]
	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}