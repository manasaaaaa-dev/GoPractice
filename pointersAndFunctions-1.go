package main
//pointers and functions


import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64 //lowercases are unexported and they won't work outside package , that is why use caps
}


//passing the value 
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


//passing the address
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}