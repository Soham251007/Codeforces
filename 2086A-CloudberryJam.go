package main
import "fmt"

func solve() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n*2)
}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}