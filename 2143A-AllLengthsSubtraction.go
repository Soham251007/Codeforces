package main
import "fmt"

func solve(){
	n := 0
	fmt.Scan(&n)
	var a = make([]int, n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	b := 0
	c := n-1
	for	i:=1;i<=n; i++{
		if a[b] == i{
			b++
		}else if a[c] == i{
			c--
		}else{
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}