package main
import "fmt"

func fibonacci() func() int {
	a, b := 0, 1    
	return func() int {
        result := a
        temp := a + b   
        a = b          
        b = temp
        return result
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}