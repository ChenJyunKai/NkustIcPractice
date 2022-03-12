package main
import "fmt"
func main(){
	fmt.Println("請輸入領取金額")
	var x int 
	fmt.Scanln(&x)
	if x<100{
		fmt.Println("太少")
	}else if x<1000{
		fmt.Println("ok")
	}else{
		fmt.Println("太多")
	}
}