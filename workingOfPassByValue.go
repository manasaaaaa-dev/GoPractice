package main
import "fmt"

type Rectangle struct {
    L, B float64
}

func AreaPassByValue(r Rectangle, change float64) float64 {
    r.L = r.L * change
    r.B = r.B * change
    return r.L * r.B
}

func main() {
    r := Rectangle{L: 3, B: 4}
    fmt.Println("Before function call — L:", r.L, "B:", r.B) // 3 4

    AreaPassByValue(r, 2)
    fmt.Println("After function call  — L:", r.L, "B:", r.B)   // still 3 4 

    
	//Even though inside the function we did r.L * 2 and r.B * 2, the original r in main is untouched.  pass by value works on a copy, and the area reurned is the area that is being worked on the copy , then again the value of l and b remain 3 and 4 
}