package main
import "fmt"

func solve(){
	var n int
	a := ""
	var m int
	b := ""
	c := ""
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&m)
	fmt.Scan(&b)
	fmt.Scan(&c)
	for i:=0; i<len(b); i++{
		if string(c[i]) == "D"{
			a = a + string(b[i])
		}
		if string(c[i]) == "V"{
			a = string(b[i]) + a
		}
	}
	fmt.Println(a)
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}