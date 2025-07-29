package main
import "fmt"

func solve(){
	var n int 
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	j := len(a) - 1
	i := 0
	ans := ""
	for z := 0; z < n; z++ {
		if z % 2 == 0 {
			if a[i] < a[j]{
				ans += "L"
				i++
			}else {
				ans += "R"
				j--
			
			}
			continue
		}
		if a[i] > a[j]{
			ans += "L"
			i++
		}else {
			ans += "R"
			j--
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