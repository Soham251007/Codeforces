package main
import ("fmt"
	"strings"
	"sort")

func solve(){
	s := ""
	fmt.Scan(&s)
	if strings.Contains(s, "FFT") == false && strings.Contains(s, "NTT") == false{
		fmt.Println(s)
		return
	}
	runeStr:=[]rune(s)
	sort.Slice(runeStr, func(i, j int) bool {
		return runeStr[i] > runeStr[j]
	})
	sortS := string(runeStr)
	fmt.Println(sortS)
}

func main(){
	var t int
	fmt.Scan(&t)
	for i:=0; i<t; i++{
		solve()
	}

}
