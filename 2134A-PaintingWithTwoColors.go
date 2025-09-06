package main
import "fmt"

func parity(a int) bool{
	if a%2 == 0{
		return true
	}
	return false
}

func solve(){
	var n,a,b int
	fmt.Scan(&n, &a, &b)
	if parity(n) == parity(b){
		if a>b{
			if parity(n) == parity(a){
				fmt.Println("YES")
				return
			}
		}else{
			fmt.Println("YES")
			return
		}
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