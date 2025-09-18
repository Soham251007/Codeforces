package main
import ("fmt"
	"sort")

func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	var b[]int
	e := 0
	o := 0
	j := 0
	for i:=0; i<n;i++{
		if a[i] % 2 == 0{
			e += a[i]
		}else{
			o += a[i]
			b = append(b,a[i])
			j++
		}
	}
	sort.Ints(b)
	if o == 0{
		fmt.Println(0)
		return
	}
	ans := 0
	
	ans += e
	for i:=len(b)-1; i>=len(b)/2; i--{
		ans += b[i]
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