package main
import (
    "fmt"
    "strings"
)

// format takes a string and a function that changes strings
func format(s string, fn func(string) string) string {
    return fn(s)
}

func main() {
    shout := func(s string) string {
        return strings.ToUpper(s) + "!"
    }

    whisper := func(s string) string {
        return strings.ToLower(s) + "..."
    }

    fmt.Println(format("hello", shout))   // HELLO!
    fmt.Println(format("HELLO", whisper)) // hello...
}