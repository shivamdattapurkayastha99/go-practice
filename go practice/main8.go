package main
import "fmt"
func main()  {
	var num int=10
	fmt.Println(num)
	fmt.Println(&num)
	var ptr *int=&num
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr=20
	fmt.Println(num)

}