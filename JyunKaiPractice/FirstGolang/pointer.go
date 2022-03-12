package main
import "fmt"
func main(){
	var x int = 3
	var xPr *int = &x
	add(xPr)
	fmt.Println(*xPr)
}
func add(xPr *int){
	*xPr = *xPr + *xPr
	fmt.Println(*xPr)
}