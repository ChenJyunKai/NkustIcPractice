package main
import "fmt"

func main(){
	var a1 Animal=Animal{"貓",4}
	var a2 Animal=Animal{"雞",2}
	a2.name = "人"
	fmt.Println(a1,a2) 
}

type Animal struct{
	name string
	leg int
}