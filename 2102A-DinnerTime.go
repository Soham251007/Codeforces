package main
import "fmt"

func solve(){
	var n,m,p,q int
	fmt.Scan(&n,&m,&p,&q)
	if n % p == 0 && (n/p)*q != m{
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}