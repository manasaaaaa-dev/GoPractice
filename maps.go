package main
import "fmt"

func main() {
	marks := map [string]int{
		"maths" :90,
		"telugu" :80,
		"english": 78,
	}

	fmt.Println ("Marks obtained in math",marks["maths"])

	for subject, mark := range marks {
        fmt.Println(subject, ":", mark)
    }
}