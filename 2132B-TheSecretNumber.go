package main
import ("fmt"
	"math")

func solve(){
	var n int
	var a[]int
	fmt.Scan(&n)
	for i:=1;i<=17;i++{
		temp := n/ (int(math.Pow(float64(10), float64(i)))+1)
		if n % (int(math.Pow(float64(10), float64(i)))+1) == 0{
			a = append(a, temp)
		}
	}
	fmt.Println(len(a))
	for i := len(a)-1; i>=0; i--{
		if i == 0{
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