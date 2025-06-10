package main
import "fmt"

 

func solve(){
	var n,d int
	fmt.Scan(&n, &d)
	fmt.Print(1 , " ")
	if n>=3 || d % 3 == 0 {
		fmt.Print(3 , " ")
	}
	if d == 5 {
		fmt.Print(5, " ")
	}
	if n >= 3 || d == 7 {
		fmt.Print(7, " ")
	} 
	
	if n>=6 || d % 9 == 0 || (n>=3 && d%3 == 0){
		fmt.Print(9 , " ")
	}
}

func main(){
	var t int 
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}