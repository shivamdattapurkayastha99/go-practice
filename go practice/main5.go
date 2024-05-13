package main
import "fmt"
func main()  {
	var myString=[]rune("resume")
	var indexed=myString[0]
	fmt.Printf("%v,%T",indexed,indexed)
	for i,v:=range myString{
		fmt.Println(i,v)
	}
	fmt.Printf("\n The length of 'myString' is %v",len(myString))
	var myRune='a'
	fmt.Printf("\nmyRune=%v",myRune)
	var strSlice=[]string{"s","u","b"}
	var catStr=""
	for i:=range strSlice{
		catStr+=strSlice[i]

	}
	fmt.Printf("\n%v",catStr)
	
}