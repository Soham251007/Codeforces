package main
import "fmt"

func solve(){
	var x,n int
	fmt.Scan(&x, &n)
	if n%2==0{
		fmt.Println(0)
		return
	}else{
		fmt.Println(x)
		return
	}
	
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}