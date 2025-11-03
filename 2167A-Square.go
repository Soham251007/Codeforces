package main
import "fmt"

func solve(){
	var a,b,c,d int
	fmt.Scan(&a, &b, &c, &d)
	if a == b && c == d{
		if a == c{
			fmt.Println("YES")
			return
		}else{
			fmt.Println("NO")
			return
		}
	}else{
		fmt.Println("NO")
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