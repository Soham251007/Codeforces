package main
import "fmt"

func solve(){
	var n int
	var s string
	fmt.Scan(n)
	fmt.Scan(&n)
	fmt.Scan(&s)
	var a[]int
	for i:=0; i<n; i++{
		if string(s[i]) == "0"{
			a = append(a, i+1)
		}
	}
	fmt.Println(len(a))
	for i:=0; i<len(a); i++{
		if i == len(a)-1{
			fmt.Print(a[i], "\n")
			continue
		}
		fmt.Print(a[i], " ")
	}
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}