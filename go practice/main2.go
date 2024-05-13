package main
import "fmt"
import "unicode/utf8"
func main()  {
	var intNum int=32
	intNum+=1
	fmt.Println(intNum)
	var floatNum float32=123.9
	fmt.Println(floatNum)
	var myString string="hello"+ " "+"shivam"
	fmt.Println(myString)
	fmt.Println(len("test"))
	fmt.Println(utf8.RuneCountInString("y"))
	var Myboolean bool=false
	fmt.Println(Myboolean)
	myVar:="text"
	fmt.Println(myVar)

	}