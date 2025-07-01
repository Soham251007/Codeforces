package main 
import ("fmt"
	"strconv")

func solve() {
	var s string
	fmt.Scan(&s)
	var intArray []int
	for i:= 0; i < len(s); i++ {
		digit,err := strconv.Atoi(string(s[i]))
		if err != nil {
			fmt.Println("Error converting character to integer:", err)
			return 
		}
		intArray = append(intArray, digit)
	}
	var cnt = make([]int,2)
	cnt[0] = 0
	cnt[1] = 0
	for i := 0; i < len(s); i++ {
		if intArray[i] == 0 {
			cnt[0]++
		} else {
			cnt[1]++
		}
	}
	for i:=0; i< len(s) + 1; i++ {
		if i == len(s) || cnt[1-intArray[i]] == 0{
			fmt.Println(len(s) - i)
			return
		}
		cnt[1-intArray[i]]--
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}