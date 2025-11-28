package main
import "fmt"
func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&a[i])
	}
	var cnt = make(map[int]int)
	for i:=0;i<n;i++{
		cnt[a[i]]++
	}
	ans:=0
	var done = make(map[int]bool)
	for i:=0;i<n;i++{
		if cnt[a[i]]>a[i]{
			if done[a[i]] == true{
				continue
			}
			ans += cnt[a[i]] - a[i]
			done[a[i]] = true
		}else if cnt[a[i]]<a[i]{
			if done[a[i]] == true{
				continue
			}
			ans += cnt[a[i]]
			done[a[i]] = true
		}
	}
	fmt.Println(ans)
}
func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}