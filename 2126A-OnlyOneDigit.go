package main
import "fmt"

func solve(){
	var x int
	fmt.Scan(&x)
	small := 9
	for x>0{
		z := x % 10
		if z < small {
			small = z
		}
		x /= 10
	}
	fmt.Println(small)
}

func main(){
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}