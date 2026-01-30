package main
import "fmt"

func solve(){
	var n,q int
	fmt.Scan(&n, &q)
	a := make([]int, n)
	b := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&a[i])
	}
	for i:=0; i<n; i++{
		fmt.Scan(&b[i])
	}
	c := make([][]int, q)
	for i:=0; i<q; i++{
		c[i] = make([]int, 2)
		fmt.Scan(&c[i][0], &c[i][1])
	}
	for i:=0; i<n; i++{
		if b[i]>a[i]{
			a[i] = b[i]
		}
	}
	for i := n-2; i >= 0; i-- {
    	if a[i+1] > a[i] {
        	a[i] = a[i+1]
    	}
	}
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
    	pref[i+1] = pref[i] + a[i]
	}
	for i := 0; i < q; i++ {
    	l, r := c[i][0], c[i][1]
    	fmt.Print(pref[r]-pref[l-1], " ")
	}
	fmt.Println()
}
func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}