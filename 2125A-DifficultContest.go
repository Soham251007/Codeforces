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
	numF:=0
	numN:=0
	numT:=0
	for i:=0; i<len(s); i++{
		if string(s[i]) == "F"{
			numF++
		}
		if string(s[i]) == "N"{
			numN++
		}
		if string(s[i]) == "T"{
			numT++
		}
	}
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