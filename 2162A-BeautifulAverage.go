package main
import ("fmt"
	"sort")
func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	fmt.Println(a[n-1])
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}