package main
import "fmt"

func solve(){
	var n,k int
	fmt.Scan(&n, &k)
	s := ""
	fmt.Scan(&s)
	ans := 0
	j:=0
	for i:=0; i<n; i++{
		if j != 0{
			if string(s[i]) == "1"{
				j = k
				continue
			}
			j--
			continue
		}
		if string(s[i]) == "0"{
			ans++
		}else{
			j = k
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