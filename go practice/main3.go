package main

import (
	"errors"
	"fmt"
)
func main()  {
	var printValue string="hello"
	printMe(printValue)
	var numerator int =11
	var denominator int =12
	var result,remainder,err=intDivision(numerator,denominator)
// 	if err!=nil{
// 		fmt.Printf(err.Error())
// 	}
// 	else if remainder==0{
// 		fmt.Printf("the result of integer division is %v",result)
// 	}
// 	else{

// 	fmt.Printf("the result of the  division is %v and remainder is %v",result,remainder)
// }
switch  {
case err!=nil:
	fmt.Printf(err.Error())
case remainder==0:
	fmt.Printf("the result of integer division is %v",result)
default:
	fmt.Printf("the result of the  division is %v and remainder is %v",result,remainder)
}
}	
func printMe(printValue string){
	fmt.Println(printValue)
}
func intDivision(numerator int,denominator int) (int,int,error){
	var err error
	if denominator==0 {
		err=errors.New("cannot divide by zero")
		return 0,0,err
	}
	var result int=numerator/denominator
	var remainder int=numerator%denominator

	return result,remainder,err
}