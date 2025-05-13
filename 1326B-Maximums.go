package main
import "fmt"

func main(){
	var n int 
	fmt.Scan(&n)
	var x int
	x = 0
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		a[i] += x
		x = max(x, a[i])
		fmt.Println(a[i])
	}
}