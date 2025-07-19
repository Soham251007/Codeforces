package main
import "fmt"

func solve(){
	var n, k int
	fmt.Scan(&n, &k)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ans := 0
	j := 0

	for i := 0; i < n; i++ {
		if a[i] == 0{
			j++
			if j == k {
				ans ++ 
				j = 0
				i ++
			}
		}else {
			j = 0
		}
			
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