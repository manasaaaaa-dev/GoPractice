package main
import "fmt"

type Item struct {
    Price    float64
    Stock  bool
}

var products = map[string]Item{
    "Pen":    Item{10, true},
    "Pencil": Item{5, false},
}

func main() {
    fmt.Println(products)
    fmt.Println("Pen details:", products["Pen"])
}