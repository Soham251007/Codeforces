package main
import "fmt"

func solve(){
	var x string
	fmt.Scan(&x)
	var z bool = false
	var y int = 0
	var n int = len(x)
	for i:= n - 1 ; i >= 0; i-- {
		if string(x[i]) != "0" {
			z = true
		} else if z {
			y++
		}
	}
	fmt.Println(n - (y + 1))
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}