package main
import "fmt"
func solve(){
	var a,b,c,d,e int
	fmt.Scan(&a,&b,&c,&d)
	if b>=a{
		fmt.Println(b)
	} else if d>=c{
		fmt.Println(-1)
	} else {
		a -= b
		e = c-d
		fmt.Println(b + (a+e-1)/e *c)
	}
}
func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}