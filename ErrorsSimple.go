package main

import "fmt"

//Create a custom error type
type DivisionError struct {
    Message string
}

//Give it an Error() method — now it satisfies the error interface
func (e *DivisionError) Error() string {
    return e.Message
}

//A function that might fail
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, &DivisionError{"cannot divide by zero!"}
    }
    return a / b, nil  // nil means no error = success
}

func main() {
    //This works fine
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)  // Result: 5
    }

    //This will fail
    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)  // Error: cannot divide by zero!
    } else {
        fmt.Println("Result:", result)
    }
}