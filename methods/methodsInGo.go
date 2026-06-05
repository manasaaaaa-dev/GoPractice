package main 

import (
	"fmt"
	"math"
)

type Vertex struct{
	x,y float64
}

func (v Vertex) Abs () float64 {
	return math.Sqrt((v.X*v.X + v.Y*v.Y))
}


// methods have the receiving type 
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
