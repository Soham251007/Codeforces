package main
import "fmt"

func solve(){
	var k int
	fmt.Scan(&k)
	if k % 3 == 1{
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