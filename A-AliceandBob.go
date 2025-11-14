package main

import (
	"fmt"
)


func solve(){
	var n, a int
	fmt.Scan(&n, &a)
	v := make([]int, n)
	small := 0
	big:=0
	for i:=0; i<n; i++{
		fmt.Scan(&v[i])
		if v[i]>a{
			big++
		}else if v[i]<a{
			small++
		}
	}
	if small>big{
		fmt.Println(a-1)
	}else{
		fmt.Println(a+1)
	}
}
func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}