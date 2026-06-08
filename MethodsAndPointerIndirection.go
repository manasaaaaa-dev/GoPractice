package main
import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v :=Vertex{3,4}
	fmt.Println(v.Abs()) //  v is a value, value receiver — direct match
	fmt.Println(AbsFunc(v)) //  function takes value, passing value



	p:=&Vertex{4,3}
	fmt.Println(p.Abs()) //  p is a pointer, Go auto converts to (*p).Abs()
	fmt.Println(AbsFunc(*p)) //  function needs value, *p dereferences pointer to value
}