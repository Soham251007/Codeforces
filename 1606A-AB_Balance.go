package main
import "fmt"

func solve(){
	var s string
	fmt.Scan(&s)
	if string(s[0]) == string(s[len(s)-1]){
		fmt.Println(s)
	} else {
		b := []byte(s)
		b[len(b)-1] = b[0]
		s = string(b)
		fmt.Println(s)
	}

}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}