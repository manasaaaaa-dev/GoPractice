package main 

import (
	"fmt"
)

type MyTemperature float64

func (tf MyTemperature) ToCelsius() float64 {
	
	return float64 ((tf-32)*5/9)
	
}

func (tf MyTemperature) IsHot() bool {
	cel := tf.ToCelsius()
	if cel > 35 {
		return true
	}	
	return false
}

func main () {
	tf := MyTemperature (98.6)
	
	fmt.Println(tf.ToCelsius())
	fmt.Println(tf.IsHot())
}
