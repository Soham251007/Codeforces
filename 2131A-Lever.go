package main
import "fmt"

func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	ans := 0
	for i:=0; i<len(a); i++{
		temp := a[i] - b[i]
		if temp>0{
			ans += temp
		}
	}
	ans = ans+1
	fmt.Println(ans)
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}