package main

import (
	"fmt"
	"math"
)

type Myfloat float64

func (f Myfloat) Absval() Myfloat {
	if f<0{
		return -f 
	}

	return f
}

func main() {
	f:= Myfloat(-math.Sqrt(13))
	fmt.Println(f.Absval())
}
