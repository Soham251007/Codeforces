package main
import "fmt"

func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	done := make(map[int]bool)
	for _, num:= range a{
		if done[num] == true{
			fmt.Println("YES")
			return
		}
		done[num] = true
	}
	fmt.Println("NO")
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}