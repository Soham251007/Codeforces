package main
import "fmt"

func solve(){
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	fmt.Println(1 - (b+c)%2, 1 - (a+c)%2, 1 - (a+b)%2)

}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}