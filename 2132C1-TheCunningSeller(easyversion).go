package main
import ("fmt"
	"math")

func solve(){
	var n int
	fmt.Scan(&n)
	x := 0
	w := -1
	var cost[]int
	for int(math.Pow(3, float64(x)))<=(n){
		w++
		val := int(math.Pow(3, float64(w+1))) + int(float64(w)*math.Pow(3, float64(w-1)))
		cost = append(cost, val)
		x++
	}
	ans := 0
	z:=0
	for i := len(cost)-1; i>=0; i--{
		for z<n{
			z += int(math.Pow(3, float64(i)))
			if z > n{
				z -= int(math.Pow(3, float64(i)))
				break
			}
			ans += cost[i]
		}
		if z == n{
			break
		}
		if z>n{
			continue
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