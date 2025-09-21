package main
import "fmt"

func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	var ans[]int
	for i:=1; i<n;i++{
		if a[i-1]<a[i]{
			ans = append(ans, i+1)
		}
	}
	fmt.Println(len(ans))
	for _, num:=range ans{
		if num == ans[len(ans)-1]{
			fmt.Print(num, "\n")
			break
		}
		fmt.Print(num, " ")
	}
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}