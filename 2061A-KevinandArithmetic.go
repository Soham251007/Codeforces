package main
import "fmt"

func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	z:=0
	x:=0
	for i := 0; i < n; i++ {
		if a[i] %2 == 0 {
			z++
		} else {
			x++
		}
	}
	ans := 0
	if z > 0{
		ans = x+1
	}else{
		ans = x-1
	}
	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}