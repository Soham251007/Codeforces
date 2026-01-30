package main
import "fmt"

func rev(s []int, i, j int) {
    if i < 0 || j >= len(s) || i >= j {
        return
    }

    for i < j {
        s[i], s[j] = s[j], s[i]
        i++
        j--
    }
}

func solve(){
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&p[i])
	}
	for i:=0; i<n; i++{
		if p[i] == n - i {
			continue
		}
		for j := i+1; j < n; j++{
			if p[j] == n-i{
				rev(p, i, j)
				break
			}
		}
		break
	}
	for i := 0; i < n; i++ {
   	 	if i > 0 {
        	fmt.Print(" ")
    	}
    	fmt.Print(p[i])
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