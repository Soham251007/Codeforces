package main
import ("fmt")

func solve(){
	var n int
	fmt.Scan(&n)
	var a []int 
	for i := 0; i<n; i++{
		if i == n-1{
			if n%2 == 0{
				a = append(a, 2)
				break
			}else{
				a = append(a, -1)
				break
			}
		}
		if i % 2 == 0{
			a = append(a, -1)
		}else{
			a = append(a, 3)
		}
	}
	for i := 0; i<n; i++{
		if i == n-1{
			fmt.Print(a[i], "\n")
			break
		}
		fmt.Print(a[i], " ")
	}
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}
}