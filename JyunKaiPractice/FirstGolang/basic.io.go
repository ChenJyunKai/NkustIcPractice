package main
import "fmt"
func main(){
	var x,y int
	fmt.Println("輸入兩個數字 用空格隔開")
	fmt.Scanln(&x,&y)
	fmt.Println("兩數相乘=",x+y)
}