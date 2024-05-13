package main
import "fmt"
func main()  {
	 intArr:=[...] int32{1,2,3}
	// intArr[1] =123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])
	fmt.Println(intArr)
	var intSlice []int32=[]int32{4,5,6}
	fmt.Println(intSlice)
	intSlice=append(intSlice, 7)

	var intSlice3 []int32=make([]int32, 3)
	var myMap map[string]uint8=make(map[string]uint8)
	fmt.Println(myMap)
	var myMap2=map[string]uint8{"Adam":23,"sarah":18}
	fmt.Println(myMap2["Adam"])
	var age,ok=myMap2["Jason"]
	if ok{
		fmt.Printf("Age is %v",age)
	}
	else{
		fmt.Println("Invalid name")
	}
	for name,age:=range myMap2{
		fmt.Printf("Name:"%v,"Age:%v\n",name,age)

	}

	for i,v:=range intArr{
		fmt.Printf("Index:%v,Value:%v\n",i,v)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}



}