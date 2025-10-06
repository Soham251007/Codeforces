package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	input := make([][]int, n)
	for i:=0; i<n; i++{
		input[i] = make([]int, 3)
		fmt.Scan(&input[i][0], &input[i][1], &input[i][2])
	}
	x := 0
	y := 0
	z := 0
	for i:=0; i<n; i++{
		x += input[i][0]
		y += input[i][1]
		z += input[i][2]
	}
	if x!=0 || y!=0 || z!=0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}