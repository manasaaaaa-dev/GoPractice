package main

import "fmt"

type Shape interface {
	Area () float64
}
type Rectangle struct {
    L, B float64
}

type Circle struct {
    Radius float64
}
func (r Rectangle) Area() float64 {
    return r.L * r.B
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// this function accepts ANY shape
func PrintArea(s Shape) {
    fmt.Println("Area is:", s.Area())
}

func main() {
    r := Rectangle{L: 3, B: 4}
    c := Circle{Radius: 5}

    PrintArea(r)  // Area is: 12
    PrintArea(c)  // Area is: 78.5
}



// What just happened?

// PrintArea accepts a Shape interface
// Both Rectangle and Circle have Area() method
// So both automatically satisfy Shape
// You can pass both to PrintArea without any extra code!
/*
why are they useful:

they are useful because without an interface we would have written 
PrintRectangleArea(r)
PrintCircleArea(c)
these are  separate functions for every shape (shape means  -> type, type rectangle , circle )
