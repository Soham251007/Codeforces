package main
import "fmt"

func main(){
	var s string
	fmt.Scan(&s)
	var l = make(map[int]string)
	l[1] = "a"
	l[2] = "b"
	l[3] = "c"
	l[4] = "d"
	l[5] = "e"
	l[6] = "f"
	l[7] = "g"
	l[8] = "h"
	l[9] = "i"
	l[10] = "j"
	l[11] = "k"
	l[12] = "l"
	l[13] = "m"
	l[14] = "n"
	l[15] = "o"
	l[16] = "p"
	l[17] = "q"
	l[18] = "r"
	l[19] = "s"
	l[20] = "t"
	l[21] = "u"
	l[22] = "v"
	l[23] = "w"
	l[24] = "x"
	l[25] = "y"
	l[26] = "z"
	var m = make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5
	m["f"] = 6
	m["g"] = 7
	m["h"] = 8
	m["i"] = 9
	m["j"] = 10
	m["k"] = 11
	m["l"] = 12
	m["m"] = 13
	m["n"] = 14
	m["o"] = 15
	m["p"] = 16
	m["q"] = 17
	m["r"] = 18
	m["s"] = 19
	m["t"] = 20
	m["u"] = 21
	m["v"] = 22
	m["w"] = 23
	m["x"] = 24
	m["y"] = 25
	m["z"] = 26
	if s == "first"{
		var n int
		fmt.Scan(&n)
		var a = make([]int, n)
		ans := ""
		for i:=0; i<n; i++{
			fmt.Scan(&a[i])
		}
		for i:=0; i<n; i++{
			temp := a[i]
			ans = ans+l[temp]
		}
		fmt.Println(ans)
	}
	if s == "second"{
		var s string
		fmt.Scan(&s)
		var a = make([]int, len(s))
		for i:=0; i<len(s); i++{
			a[i] = m[string(s[i])]
		}
		fmt.Println(len(a))
		for i := 0; i < len(a); i++ {
			if i==len(a)-1{
				fmt.Print(a[i], "\n")
				return
			}
			fmt.Print(a[i], " ")
		}
	}
}