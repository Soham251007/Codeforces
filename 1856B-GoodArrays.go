package main
import "fmt"
func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var sum, count int
	for i := 0; i < n; i++ {
		sum += a[i]
		if a[i] == 1 {
			count++
		}
	}
	if sum >= count + n && n > 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}