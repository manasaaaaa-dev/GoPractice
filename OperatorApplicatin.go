package main
import "fmt"

//let us take a function that accepts any function that takes
//2 int as input and returns int 

func apply (fn func(int, int) int ,a,b int)int{
	return fn(a,b)
}

func main() {
    add := func(x, y int) int {
        return x + y
    }

    multiply := func(x, y int) int {
        return x * y
    }

    fmt.Println(apply(add, 5, 3))      // 8
    fmt.Println(apply(multiply, 5, 3)) // 15
}