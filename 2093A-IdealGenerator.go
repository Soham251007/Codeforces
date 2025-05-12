package main
import "fmt"

func solve() {
	var k int
	fmt.Scan(&k)
	if k % 2 == 0{
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}