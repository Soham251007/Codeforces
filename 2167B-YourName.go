package main
import ("fmt")

func solve(){
	var n int
	fmt.Scan(&n)
	var s,t string
	fmt.Scan(&s, &t)
	var done = make(map[string]int)
	var done1 = make(map[string]int)
	for i:=0; i<n; i++{
		done[string(s[i])]++
		done1[string(t[i])]++
	}
	for i:=0; i<n; i++{
		if done[string(s[i])] == done1[string(s[i])]{
			continue
		}else{
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}