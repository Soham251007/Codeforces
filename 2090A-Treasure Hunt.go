package main
import "fmt"

func solve(){
	var x,y,a int
	fmt.Scan(&x,&y,&a)
	if a%(x+y) < x {
		fmt.Println("NO \n")
	}else {
		fmt.Println("YES \n")
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}