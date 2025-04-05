package main
import "fmt"

func jumDigit(n int)int{
	var hasil, i int
	if n == 0{
		return 0
	}
	i = n%10
	hasil += i
	return hasil + jumDigit(n/10)
}
func main(){
	var num int
	fmt.Scan(&num)
	fmt.Print(jumDigit(num))
}