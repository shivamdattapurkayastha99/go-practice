package main
import ("fmt"
		"math"		
	)
type Geometry interface{
	Area() float64
	

}
type Rectangle struct{
	Width float64
	Height float64

}
type Circle struct {
    Radius float64
}
func (r Rectangle) Area() float64  {
	return r.Width*r.Height
	
}
func (c Circle) Area() float64  {
	return math.Pi*c.Radius*c.Radius
	
}
func measure(g Geometry)  {
	fmt.Println(g.Area())
}
func main()  {
	rectangle:=Rectangle{Width:3,Height:4}
	circle:=Circle{Radius: 5}
	measure(rectangle)
	measure(circle)
}