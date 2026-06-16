package main
import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()

    fmt.Println(pos(10))  // 10   Notebook A: 0+10=10
    fmt.Println(pos(10))  // 20   Notebook A: 10+10=20
    fmt.Println(neg(5))   // 5    Notebook B: 0+5=5, totally separate
    fmt.Println(pos(10))  // 30   Notebook A: 20+10=30, neg didn't affect it
}

//Notice that  neg(5) returning 5 did NOT mess up pos's running total. They are completely independent!