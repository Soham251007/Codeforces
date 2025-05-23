package main
import ("fmt"
	"sort")

func solve() {
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
	sort.Ints(a)
	sort.Ints(b)
	var c int
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			c += 1
		} else if a[i] + 1 == b[i]{
			c += 1
		}
	}
	if c == n {
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