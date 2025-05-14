package main 

import "fmt"

func solve(){
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)
	var s string
	for i := 1; i <= 12; i++ {
		if i== a || i == b{
			s+= "a"
		}
		if i == c || i == d{
			s+= "b"
		}
	}
	if string(s) == "abab" || string(s) == "baba" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}