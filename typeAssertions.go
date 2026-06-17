package main
import "fmt"

func main() {
	var i interface{} = "manasa"
	//case -1
	s:=i.(string)// safe assertion : assume the interface has a string 
	fmt.Println(s)

	//case-2 
	s, ok := i.(string)
	fmt.Println(s, ok)//check if  the vaue present in the  interface 

	//case-3
	s,ok := i.(float64)//wrong but safe as we are asking to check what is there in the interface and then asking to print the vaalue 
	fmt.Println(s,ok)

	//case-4 
	//overconfident , and not even pairing with ok to check if the value present is something that they decide :this is the panic case 
	s:= i.(string)
	fmt.Println(s)
}