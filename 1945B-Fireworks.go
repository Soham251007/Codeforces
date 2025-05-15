package main
import "fmt"

func solve(){
	var a,b,m int
	fmt.Scan(&a,&b,&m)
	sum := m/a + m/b + 2
	fmt.Println(sum)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}