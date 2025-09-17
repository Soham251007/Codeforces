package main
import "fmt"

func solve(){
	var n,m,x,y int
	fmt.Scan(&n,&m,&x,&y)
	var a = make([]int, n)
	var b = make([]int, m)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	for i:=0;i<m;i++{
		fmt.Scan(&b[i])
	}
	fmt.Println(m+n)
	return
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}