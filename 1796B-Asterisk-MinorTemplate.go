package main

import ("fmt"
	"strings")

func solve() {
	var a,b string
	fmt.Scan(&a)
	fmt.Scan(&b)
	if string(a[0]) == string(b[0]) {
		fmt.Println("YES")
		fmt.Print(string(a[0]), "*", "\n")
		return
	}
	if string(a[len(a)-1]) == string(b[len(b)-1]){
		fmt.Println("YES")
		fmt.Print("*", string(a[len(a)-1]), "\n")
		return
	}
	
	for i := 0; i < len(b)-1; i++ {
		c := string(b[i]) + string(b[i+1])
		if strings.Contains(a, c) {
			fmt.Println("YES")
			fmt.Print("*", string(b[i]), string(b[i+1]), "*", "\n")
			return
		}
	}
	fmt.Println("NO")
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}