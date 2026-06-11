package main 
import "fmt"

type Marks struct {
    Maths,Science int
}

var students = map[string]Marks {
    "Ravi" : Marks {90,88},
}
//short notes :map literal lets you create a map and fill it with values in one step (instead of creating empty + adding keys one by one).
func main() {
    fmt.Println(students)
    fmt.Println("Ravi's marks:", students["Ravi"])
}
    
