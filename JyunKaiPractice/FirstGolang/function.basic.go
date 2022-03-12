package main 
import "fmt"
func main(){
	test("test")
	var result int = sum(1,2)
	fmt.Println(result)
	fmt.Println(test_many())
}
func test(text string){
	fmt.Println(text)
}
func sum(a int,b int) int{
	return a+b
}
func test_many() (int,string){
	return 1,"hello"
}