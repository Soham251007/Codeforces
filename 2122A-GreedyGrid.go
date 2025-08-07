package main
import "fmt"

func solve(){
	var n,m int
	fmt.Scan(&n,&m)
	if n == 2 && m == 2 {
		fmt.Println("No")
	}else if n == 1 || m == 1{
		fmt.Println("No")
	}else{
		fmt.Println("Yes")
	}
}

func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
