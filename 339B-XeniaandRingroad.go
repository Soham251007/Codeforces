package main
import "fmt"


func main(){
	var  n, m int
	fmt.Scan(&n, &m)
	var a = make([]int, m+1)
	a[0] = 1
	for i := 1; i < m+1; i++ {
		fmt.Scan(&a[i])
	}
	var temp, ans int
	for i := 0; i < m; i++ {
		if a[i]<= a[i+1] {
			temp = a[i+1] - a[i]
			ans += temp
		} else {
			temp = n - a[i] + a[i+1]
			ans += temp
		}
	}
	fmt.Println(ans)
}