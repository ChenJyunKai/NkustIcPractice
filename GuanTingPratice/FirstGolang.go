package main
import (
	"fmt" 
	"strings"
	"strconv"
	"bufio"
	"os"
) 

func main() {
	var equation string 
	fmt.Println("Please enter your equation")
	Reader := bufio.NewReader(os.Stdin)
	equation,_ = Reader.ReadString('\n');equation = strings.TrimSpace(equation)
	fmt.Println("The answer to this equation is：",equationprocessing(equation))
}  

func equationprocessing(equation string) int{
	equation = strings.Replace(equation," ","", -1)
	equation = strings.Replace(equation,"+"," + ", -1)
	equation = strings.Replace(equation,"-"," - ", -1)
	equation = strings.Replace(equation,"*"," * ", -1)
	equation = strings.Replace(equation,"/"," / ", -1)
	var splitarray []string = strings.Split(equation," ")
	if strings.ContainsRune(equation, '*') ||  strings.ContainsRune(equation, '/') {
		splitarray = strings.Split(muldiv(splitarray)," ") 
	}
	var sum int = 0
	return Addsub(sum,splitarray)
}

func Addsub(sum int,equation []string) int{
	var i int
	var preadd int
	var err error
	for i = 0; i < len(equation); i += 2{
		preadd, err = strconv.Atoi(equation[i])
		if err != nil {
			fmt.Println(err.Error())
		}
		if i == 0 {
			sum = preadd
		}else{
			if equation[i-1] == "+"{
				sum = sum + preadd
			}else{
				sum = sum - preadd
			}
		}	
	}
	return sum
}

func muldiv(equation []string) string{
	var n,i int
	var x,y int
	var err error
	var str string = ""
	for i = 1; i < len(equation)-1 ; i += 2{
		n = 0
		x, err = strconv.Atoi(equation[i-1])
		y, err = strconv.Atoi(equation[i+1])
		if err != nil {
			fmt.Println(err.Error())
		}
		if equation[i] == "*"{
			n = x * y
			equation[i-1] = "";equation[i] = ""
			equation[i+1] = strconv.Itoa(n)
			
		}else if equation[i] == "/"{
			equation[i-1] = "";equation[i] = ""
			n = x / y
			equation[i+1] = strconv.Itoa(n)
		}
	}
	for i = 0 ; i < len(equation); i++ {
		if equation[i] != "" {
			str = str + equation[i] + " "
		}
	}
	str = strings.TrimSpace(str)
	return str
}
