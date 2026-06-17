package main
import "fmt"

type Car struct {
	Brand string
	Year int
}

func (c Car) String() string {
	return fmt.Sprintf("%v,(%v model)",c.Brand,c.Year)
}


func main() {
	a := Car{"Toyota",200}
	b := Car{"Harrier",900}
	fmt.Println(a, b)
}