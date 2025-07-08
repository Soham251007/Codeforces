package main
import "fmt"

func solve(){
	var n int 
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i:=0 ; i<n; i++{
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				fmt.Println("YES")
				fmt.Println(2) 
				fmt.Println(a[i], a[j])
				return
			}
		}
	}
	fmt.Println("NO")
}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}