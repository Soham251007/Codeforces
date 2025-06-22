package main
import "fmt"

func solve(){
	var n int 
	fmt.Scan(&n)
	var ans int
	var end int 
	end = -1
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a - end > 1{
			ans ++
			end = a
		}
	}
	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}