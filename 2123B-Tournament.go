package main 
import ("fmt"
	"sort")

func solve(){
	var n, j, k int
	fmt.Scan(&n, &j, &k)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	if k>1 {
		fmt.Println("YES")
		return
	}
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i]
	}
	sort.Ints(a)
	 for i, k := 0, len(a)-1; i < k; i, k = i+1, k-1 {
        a[i], a[k] = a[k], a[i]
    }
	if a[0] == b[j-1]{
		fmt.Println("YES")
		return
	} else {
		fmt.Println("NO")
		return
	}
}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}