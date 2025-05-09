package main
import ("fmt")

func solve(){
	var x string
	var y string
	var z string
	fmt.Scan(&x, &y, &z)
	fmt.Println(string(x[0]) + string(y[0]) + string(z[0]))
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
