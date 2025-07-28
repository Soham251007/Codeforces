package main
import ("fmt"
	"sort")
func solve() {
	var n , c int
	fmt.Scan(&n, &c)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	var x[]int
	for _, element := range a {
		x = append([]int{element}, x...)
	}
	temp := 0
	for i := 0; i < n; i++ {
		if x[i] <= c{
			for j := i + 1; j < n; j++ {
				x[j]=x[j]*2
			}
			temp++
		}
	}
	ans := len(x) - temp
	fmt.Println(ans)
}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}