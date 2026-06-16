package main
import "fmt"

type Notifier interface {
	Send()
}

type Email struct {
	To string
}

func (e *Email) Send() {
	fmt.Println("Send email to " , e.To)
}

type SMS string

func (s SMS) Send() {
	fmt.Println("send sms to " , s.TO)
}

func describe(n Notifier) {
	fmt.Printf("(%v, %T)\n", n, n) //This prints the same variable i twice — once showing its value, once showing its type.
}


func main() {
	var n Notifier 
	n = &Email{To : manasa@gmail.com}
	describe(n)
	n.Send()


	n = SMS("9876543210") 
	describe(n)
	n.Send()
}


//%v and %T print the value as is and the type of the value 
//fmt.Printf("%T\n", 42)        // int
// fmt.Printf("%T\n", "hello")   // string
// fmt.Printf("%T\n", 3.14)      // float64
// fmt.Printf("%T\n", Square{4}) // main.Square
