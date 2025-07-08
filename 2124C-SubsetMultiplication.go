package main 
import "fmt"

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / Gcd(a, b) * b
}
func solve() {
	var n int 
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	gcd := 0
	lcm := 1
	for i:= n-1 ; i >= 0; i-- {
		gcd = Gcd(gcd, a[i])
		lcm = Lcm(lcm, a[i]/gcd)
	}
	fmt.Println(lcm)

}
func main() {
	var t int 
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}