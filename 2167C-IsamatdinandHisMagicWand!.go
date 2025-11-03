package main
import ("fmt"
		"sort")

func solve(){
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&a[i])
	}
	flag := 0
	for i:=0; i<n-1; i++{
		if a[i]%2 == a[i+1]%2{
			flag = 1
		}else{
			flag = 0
			break
		}
	}
	if flag == 1{
		for i:=0; i<n; i++{
			if i == n-1{
				fmt.Print(a[i], "\n")
				break
			}
			fmt.Print(a[i], " ")
		}
		return
	}else{
		sort.Ints(a)
		for i:=0; i<n; i++{
			if i == n-1{
				fmt.Print(a[i], "\n")
				break
			}
			fmt.Print(a[i], " ")
		}
		return
	}
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}