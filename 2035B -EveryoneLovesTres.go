package main
import "fmt"

func solve() {
	var n int
	fmt.Scan(&n)
	if n == 1 || n == 3{
		fmt.Println(-1)
	} else if n % 2 == 0 {
		for i := 1; i <= n-2 ; i++ {
			fmt.Print(3)
		}
		fmt.Print(66, " ")
	} else if n % 2 != 0 && n >=5 {
		for i := 1; i <= n-5 ; i++ {
			fmt.Print(3)
		}
		fmt.Print(36366, " ")
	}
}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}