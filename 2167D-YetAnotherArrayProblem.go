package main
import ("fmt"
		"sort")

func gcd(a int, b int) int{
	for b!=0{
		a,b = b, a%b
	}
	return a
}
func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&a[i])
	}
	var ans[]int
	j := 0
	for j < n{
		i:=2
		for i>0{
			if gcd(a[j], i) == 1{
				ans = append(ans, i)
				break
			}
			i++
		}
		j++
	}
	if len(ans) == 0{
		fmt.Println(-1)
		return
	}
	sort.Ints(ans)
	fmt.Println(ans[0])
	return
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}